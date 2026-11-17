package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const apiURL = "https://f1-companion-api-ba5k.onrender.com/drivers"

// fetchDrivers récupère et décode la liste des pilotes depuis l'API.
func fetchDrivers(client *http.Client) (APIResponse, error) {
	var result APIResponse
	response, err := client.Get(apiURL)
	if err != nil {
		return result, fmt.Errorf("appel API impossible : %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return result, fmt.Errorf("statut HTTP inattendu : %d", response.StatusCode)
	}

	decoder := json.NewDecoder(io.LimitReader(response.Body, 5<<20))
	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("JSON invalide : %w", err)
	}
	if result.Drivers == nil {
		return result, errors.New("réponse JSON sans liste de pilotes")
	}
	return result, nil
}