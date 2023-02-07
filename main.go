package main

import (
	"context"
	"embed"
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

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <matchid>\n", os.Args[0]) // TODO with options:		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <url>\n", os.Args[0])
		os.Exit(1)
	}

	matchID := os.Args[1]

	client := grassroots.NewClient(os.Getenv("GRASSROOTS_KEY"))
	match, err := client.GetMatch(context.Background(), matchID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get match error: %v\n", err)
		os.Exit(1)
	}

	template, err := template.ParseFS(templates, "templates/matchresults_text.go.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Newspaper template error: %v\n", err)
		os.Exit(1)
	}
	template.Execute(os.Stdout, match)
}
