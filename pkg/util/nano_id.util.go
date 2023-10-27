package util

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	customAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	customNumbers  = "0123456789"
)

func GenerateRandomCharacters() (string, error) {
	nanoID, err := gonanoid.Generate(customAlphabet, 8)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("MID%v", nanoID), nil
}

func GenerateRandomNumber(size int) (string, error) {
	nanoID, err := gonanoid.Generate(customNumbers, size)
	if err != nil {
		return "", err
	}
	return nanoID, nil
}
