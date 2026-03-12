package main

import (
	"strings"
)

func AsciiArt(text, subString, color string, banner []string) string {
	strBuilder := strings.Builder{}
	Reset := "\033[0m"

	lines := strings.Split(text, "\\n")

	for _, line := range lines {
		for row := 1; row < 9; row++ {
			for i, ch := range line {
				ascii := int(ch)
				index := (ascii - 32) * 9
				ascii_art := banner[index+row]

				if SubString(line, subString, i) && color != "" {
					strBuilder.WriteString(color + ascii_art + Reset)
				} else {
					strBuilder.WriteString(ascii_art)
				}
			}
			strBuilder.WriteByte('\n')
		}
		strBuilder.WriteByte('\n')
	}
	return strBuilder.String()
}




func SubString(text, subString string, position int) bool{
	if subString == "" {
		return true
	}
	for i := 0; i <= len(text) - len(subString); i++{
		end := i + len(subString)
		if text[i:end] == subString{
			if position >= i && position < end {
				return true
			}
		}

	}
	return false
}