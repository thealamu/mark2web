package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		// No argument specified
		printHelp()
	}
}

func printHelp() {
	fmt.Println(strings.TrimSpace(usageString))
}

const usageString = `
Usage: mark2web FILE
Renders markdown FILE as webpage, returning a URL to the webpage.
`
