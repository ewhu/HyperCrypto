// cmd/hypercrypto/main.go
package main

import (
	"flag"
	"log"
	"os"

	"hypercrypto/internal/hypercrypto"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := hypercrypto.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
