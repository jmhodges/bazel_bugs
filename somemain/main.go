package main

import (
	"fmt"
	"log"

	"github.com/jmhodges/gocld3/cld3"
)

func main() {
	langIder, err := cld3.NewLanguageIdentifier(0, 512)
	if err != nil {
		log.Fatalf("unable to make LanguageIdentifier: %s", err)
	}
	fmt.Println("Hello, World! is ", langIder.FindLanguage("Hello, World!").Language)
}
