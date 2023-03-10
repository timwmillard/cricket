package main

import (
	"context"
	"embed"
	"encoding/json"
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

const (
	FormatJSON = "json"
	FormatText = "text"
)

const version = "v0.1.4"

func main() {
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Printf("%s [OPTIONS] <matchid>\n", os.Args[0])
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
	auth := flag.Bool("auth", false, "configure API Key")
	j := flag.Bool("j", false, "format output as JSON")
	v := flag.Bool("version", false, "print the version")
	flag.Parse()

	if *v {
		fmt.Println(version)
		os.Exit(0)
	}

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
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	matchID := flag.Arg(0)

	// Format
	format := FormatText
	if *j {
		format = FormatJSON
	}

	// Grassroots API call
	client := grassroots.NewClient(apiKey)
	match, err := client.GetMatch(context.Background(), matchID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get match error: %v\n", err)
		os.Exit(1)
	}

	if format == FormatJSON {
		output, err := json.Marshal(match)
		if err != nil {
			fmt.Fprintf(os.Stderr, "JSON error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
	} else { // Text format

		funcMap := template.FuncMap{
			"FormatSchedule": grassroots.ScheduleTime,
			"FallOfWickets":  grassroots.FallOfWicketList,
		}

		// Run template
		tmpl, err := template.New("").Funcs(funcMap).ParseFS(templates, "templates/matchresults_text.go.tmpl")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Newspaper template error: %v\n", err)
			os.Exit(1)
		}
		tmpl.ExecuteTemplate(os.Stdout, "matchresults_text.go.tmpl", match)
	}
}
