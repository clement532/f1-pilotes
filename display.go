package main

import (
	"fmt"
	"strconv"
)

func displayDriver(driver Driver) {
	fmt.Println("--------------------------")
	fmt.Printf("Numéro      : %s\n", driver.Number)
	fmt.Printf("Prénom      : %s\n", driver.FirstName)
	fmt.Printf("Nom         : %s\n", driver.LastName)
	fmt.Printf("Nationalité : %s\n", driver.Nationality)
	fmt.Printf("Code        : %s\n", driver.Code)
	fmt.Println("--------------------------")
}

func displayAllDrivers(data APIResponse) {
	fmt.Printf("\nPilotes de la saison %s (%d) :\n", data.Season, len(data.Drivers))
	for _, driver := range data.Drivers {
		displayDriver(driver)
	}
}

func displayNoResult(search string) {
	if _, err := strconv.Atoi(search); err == nil {
		fmt.Printf("Aucun pilote ne porte le numéro %s.\n", search)
	} else {
		fmt.Printf("Aucun pilote trouvé pour %q.\n", search)
	}
}

func executeSearch(data APIResponse, search string) {
	if search == "all" {
		displayAllDrivers(data)
		return
	}
	matches := findDrivers(data.Drivers, search)
	if len(matches) == 0 {
		displayNoResult(search)
		return
	}
	for _, driver := range matches {
		displayDriver(driver)
	}
}
