package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	flag.CommandLine.Usage = usage
	name := flag.String("n", "", "The `name` of the template to be rendered. If empty defaults to the first specified argument")
	flag.Parse()

	if flag.NArg() < 1 {
		die(fmt.Errorf("no files provided"))
	}

	envs := os.Environ()
	envMap := make(map[string]string)
	for _, env := range envs {
		parts := strings.Split(env, "=")
		envMap[parts[0]] = parts[1]
	}

	tmpl, err := template.ParseFiles(flag.Args()...)
	if err != nil {
		die(err)
	}

	if *name == "" {
		err = tmpl.Execute(os.Stdout, envMap)
	} else {
		err = tmpl.ExecuteTemplate(os.Stdout, *name, envMap)
	}
	if err != nil {
		die(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "USAGE: %s [flag]... [file]...\nAvailable flags:\n", os.Args[0])
	flag.PrintDefaults()
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
