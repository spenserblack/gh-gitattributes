package main

import (
	"flag"
	"fmt"
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
	err = client.Get(fmt.Sprintf("repos/%s/contents", cfg.Source), &fileListResponse)
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
	gitattributeKeys := make([]string, 0, len(gitattributes))
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

	answers := struct {
		UseCommon bool
		Project   string
	}{}
	err = survey.Ask(qs, &answers)
	onError(err)

	fmt.Println(answers)
	// fmt.Println("hi world, this is the gh-gitattributes extension!")
	// client, err := api.DefaultRESTClient()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// response := struct {Login string}{}
	// err = client.Get("user", &response)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("running as %s\n", response.Login)
}

func onError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
