package main

import (
	"fmt"
	"log"

	"github.com/goodleby/playground/christmas"
	"github.com/goodleby/playground/wavyword"
)

func main() {
	// Wavy word example
	wavyWord, err := wavyword.New("reallylongwavyword", 3)
	if err != nil {
		log.Fatalf("Error creating new wavy word: %v", err)
	}
	wavyWord.Print()

	// Christmas song example
	lines := christmas.SongLines()
	for _, line := range lines {
		fmt.Println(line)
	}
}
