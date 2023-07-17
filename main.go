package main

import (
	"flag"
	"fmt"
	"os"
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
