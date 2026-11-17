package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

func run() error {
	client := &http.Client{Timeout: 90 * time.Second}
	if len(os.Args) > 2 {
		return errors.New("utilisation : go run . [numéro|code|nom|all]")
	}
	if len(os.Args) == 2 {
		return runWithArgument(client, os.Args[1])
	}
	return runInteractive(client)
}

func runWithArgument(client *http.Client, input string) error {
	search, err := validateSearch(input)
	if err != nil {
		return err
	}
	if search == "q" {
		return nil
	}
	data, err := fetchDrivers(client)
	if err != nil {
		return err
	}
	executeSearch(data, search)
	return nil
}

func runInteractive(client *http.Client) error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Recherche de pilotes F1 — 'all' pour tous, 'q' pour quitter")
	var cached *APIResponse
	for {
		fmt.Print("Votre recherche : ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("lecture de la saisie : %w", err)
			}
			fmt.Println("\nAu revoir !")
			return nil
		}
		search, err := validateSearch(scanner.Text())
		if err != nil {
			fmt.Println("Erreur :", err)
			continue
		}
		if search == "q" {
			fmt.Println("Au revoir !")
			return nil
		}
		if cached == nil {
			data, err := fetchDrivers(client)
			if err != nil {
				fmt.Println("Erreur :", err)
				continue
			}
			cached = &data
		}
		executeSearch(*cached, search)
	}
}
