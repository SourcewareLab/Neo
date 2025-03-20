package trinity

import (
	lexscanner "NeoLang/Trinity/Scanner"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func RunLexer(path string) error {
	data, err := ReadFile(path)
	if err != nil {
		return err
	}

	scanner := lexscanner.NewScanner(data)

	scanner.ScanTokens()

	fmt.Println(data)
	fmt.Println("------- TOKENS --------")

	for _, val := range scanner.Tokens {
		fmt.Println(val)
	}

	return nil
}

func ReadFile(path string) (string, error) {
	fileExtension := filepath.Ext(path)
	if fileExtension != ".neo" {
		return "", errors.New("File is not a Neo file.")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
