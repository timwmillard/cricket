package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

const configFile = ".cricket"

func promptAuth() {
	var apiKey string
	fmt.Print("Enter API Key: ")
	_, err := fmt.Scanln(&apiKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = saveAuth(apiKey)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func saveAuth(apiKey string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filepath := path.Join(homeDir, configFile)

	config, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer config.Close()

	_, err = config.WriteString("AUTH=" + apiKey + "\n")
	if err != nil {
		return err
	}

	return nil
}

func readAuth() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filepath := path.Join(homeDir, configFile)

	config, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer config.Close()

	scanner := bufio.NewScanner(config)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		if len(line) < 2 {
			continue
		}
		if line[0] == "AUTH" {
			return line[1], nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}
