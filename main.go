package main

import (
	nexus "NeoLang/Nexus"
	"flag"
	"fmt"
)

func main() {
	pathPtr := flag.String("path", "", "path to file")

	flag.Parse()

	fmt.Println(*pathPtr)

	data, err := nexus.ReadFile(*pathPtr)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(data)
}
