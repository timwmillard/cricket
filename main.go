package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/timwmillard/cricket/grassroots"
)

var (
	//go:embed templates/*
	templates embed.FS
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Printf("%s [OPTIONS] <matchid>\n", os.Args[0])
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
	auth := flag.Bool("auth", false, "configure API Key")
	flag.Parse()

	// Get Auth Key
	if *auth {
		promptAuth()
	}
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	apiKey, err := readAuth()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading config:", err)
		os.Exit(1)
	}
	if apiKey == "" {
		promptAuth()
	}

	// Get Match ID
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	matchID := os.Args[1]

	// Grassroots API call
	client := grassroots.NewClient(apiKey)
	match, err := client.GetMatch(context.Background(), matchID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get match error: %v\n", err)
		os.Exit(1)
	}

	// Run template
	template, err := template.ParseFS(templates, "templates/matchresults_text.go.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Newspaper template error: %v\n", err)
		os.Exit(1)
	}
	template.Execute(os.Stdout, match)
}
