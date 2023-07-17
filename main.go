package main

import (
	"flag"
	"fmt"
	"os"

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
	fmt.Println(fileListResponse)

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
