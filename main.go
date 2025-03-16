package main

import (
	trinity "NeoLang/Trinity"
	"flag"
	"fmt"
)

func main() {
	pathPtr := flag.String("path", "", "path to file")

	flag.Parse()

	fmt.Println(*pathPtr)

	data, err := trinity.ReadFile(*pathPtr)
	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := trinity.Scanner{
		Source: data,
	}

	scanner.ScanTokens()

	fmt.Println(data)
	fmt.Println("------- TOKENS --------")

	for _, val := range scanner.Tokens {
		fmt.Println(val)
	}
}
