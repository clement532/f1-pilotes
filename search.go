package main

import "strings"

// findDrivers recherche par numéro, code, nom ou prénom + nom.
func findDrivers(drivers []Driver, search string) []Driver {
	var matches []Driver
	for _, driver := range drivers {
		if driver.Number == search ||
			strings.EqualFold(driver.Code, search) ||
			strings.EqualFold(driver.LastName, search) ||
			strings.EqualFold(driver.FirstName+" "+driver.LastName, search) {
			matches = append(matches, driver)
		}
	}
	return matches
}
