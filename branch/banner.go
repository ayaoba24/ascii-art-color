package branch

import (
	"log"
	"os"
	"strings"
)

func Banner(s string) []string {
    content, err := os.ReadFile(s)
    if err != nil {
        log.Fatalf("Banner file not found: %v", err) 
    }
    return strings.Split(string(content), "\n")
}