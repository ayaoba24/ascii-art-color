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
		if strings.HasPrefix(args[0], "--color=") {
			colorName := strings.TrimPrefix(args[0], "--color=")
			colorCode = branch.Color(colorName)
			text = args[1]
		} else {
			subString = args[0]
			text = args[1]
		}
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


func getValidColor(colorName string) string {
	validColors := []string{"black", "red", "green", "yellow", "blue",  "cyan", "white"}

	for {
		colorCode := branch.Color(colorName)
		if colorCode != "" {
			return colorCode
		}

		fmt.Printf("NOT '%s' is not a valid color.\n", colorName)
		fmt.Printf("Available colors: %s\n", strings.Join(validColors, ", "))
		fmt.Print("Please enter a valid color: ")

		fmt.Scan(&colorName)
	}
}
