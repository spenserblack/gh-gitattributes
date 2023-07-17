package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cli/go-gh/v2/pkg/api"
)

func main() {
	cfg := newConfigOrDefault(os.Stderr)
	setupFlags(cfg)
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage of %s\n", os.Args[0])
		fmt.Fprintln(out)
		fmt.Fprintln(out, "CONFIGURATION")
		fmt.Fprintln(out, "\tgh_gitattributes_source")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "FLAGS")
		flag.PrintDefaults()
	}
	flag.Parse()

	client, err := api.DefaultRESTClient()
	onError(err)
	fileListResponse := []struct {
		Name string
	}{}
	err = client.Get(fmt.Sprintf("repos/%s/contents", *sourceFlag), &fileListResponse)
	onError(err)

	qs := make([]*survey.Question, 0, 2)
	gitattributes := make(map[string]string, len(fileListResponse))

	for _, file := range fileListResponse {
		if file.Name == ".gitattributes" {
			// NOTE We don't want to use the repo's own gitattributes file
			continue
		}
		if file.Name == "Common.gitattributes" {
			qs = append(qs, &survey.Question{
				Name: "UseCommon",
				Prompt: &survey.Confirm{
					Message: "Use Common.gitattributes?",
					Default: false,
				},
			})
		} else if strings.HasSuffix(file.Name, ".gitattributes") {
			gitattributes[strings.TrimSuffix(file.Name, ".gitattributes")] = file.Name
		}
	}
	if len(gitattributes) == 0 {
		fmt.Println("No gitattributes files found")
		return
	}
	noProject := "None - No project type"
	gitattributeKeys := make([]string, 1, len(gitattributes) + 1)
	gitattributeKeys[0] = noProject
	for key := range gitattributes {
		gitattributeKeys = append(gitattributeKeys, key)
	}

	qs = append(qs, &survey.Question{
		Name: "Project",
		Prompt: &survey.Select{
			Message: "Select a project type",
			Options: gitattributeKeys,
		},
	})

	askOpts := []survey.AskOpt{survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)}
	answers := struct {
		UseCommon bool
		Project   string
	}{}
	err = survey.Ask(qs, &answers, askOpts...)
	onError(err)

	selected := make([]string, 0, 2)
	if answers.UseCommon {
		selected = append(selected, "Common.gitattributes")
	}
	if answers.Project != noProject {
		selected = append(selected, gitattributes[answers.Project])
	}

	if len(selected) == 0 {
		fmt.Fprintln(os.Stderr, "Nothing to write. The output will be empty.")
	}

	var out io.Writer
	if *stdoutFlag {
		out = os.Stdout
	} else {
		out, err = os.Create(".gitattributes")
		onError(err)
	}

	for _, file := range selected {
		err = writeFile(out, client, fmt.Sprintf("repos/%s/contents/%s", *sourceFlag, file))
		onError(err)
	}
}

func onError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func writeFile(w io.Writer, api *api.RESTClient, path string) error {
	response := struct {
		Content  string
		Encoding string
	}{}
	// NOTE We're always assuming the encoding is base64 right now
	err := api.Get(path, &response)
	if err != nil {
		return err
	}
	if response.Encoding != "base64" {
		return fmt.Errorf("unsupported encoding: %s", response.Encoding)
	}
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(response.Content))
	_, err = io.Copy(w, decoder)
	return err
}
