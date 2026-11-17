package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// validateSearch accepte un numéro positif, 'all', 'q', un code ou un nom.
func validateSearch(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("saisissez un numéro positif, un code, un nom, 'all' ou 'q'")
	}
	if strings.EqualFold(input, "all") || strings.EqualFold(input, "q") {
		return strings.ToLower(input), nil
	}
	if number, err := strconv.Atoi(input); err == nil {
		if number <= 0 {
			return "", errors.New("le numéro doit être un entier strictement positif")
		}
		return strconv.Itoa(number), nil
	}
	if strings.ContainsRune("+-0123456789", rune(input[0])) {
		return "", errors.New("numéro invalide : saisissez un entier strictement positif")
	}
	for _, r := range input {
		if !unicode.IsLetter(r) && r != ' ' && r != '-' && r != '\'' {
			return "", errors.New("recherche invalide : utilisez un numéro positif, un code ou un nom")
		}
	}
	return input, nil
}