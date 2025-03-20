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

	trinity.RunLexer(*pathPtr)
}
