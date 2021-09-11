package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	envs := os.Environ()
	envMap := make(map[string]string)
	for _, env := range envs {
		parts := strings.Split(env, "=")
		envMap[parts[0]] = parts[1]
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	tmpl, err := template.ParseFiles(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	err = tmpl.Execute(os.Stdout, envMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
