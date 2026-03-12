package main

import (
	"asciiart-color/branch"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	var colorCode, subString, text string

	switch len(args) {
	case 1:
		text = args[0]
	case 2:
		subString = args[0]
		text = args[1]
	case 3:
		if !strings.HasPrefix(args[0], "--color=") {
			log.Fatal("Invalid arguments")
		}
		colorName := strings.TrimPrefix(args[0], "--color=")
		colorCode = branch.Color(colorName)
		subString = args[1]
		text = args[2]
	default:
		log.Fatal("Exceeded Length of Argument!")
	}

	banner := branch.Banner("standard.txt")
	result := branch.AsciiArt(text, subString, colorCode, banner)
	fmt.Print(result)
}