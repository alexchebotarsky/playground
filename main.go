package main

import (
	"fmt"
	"log"

	"github.com/goodleby/playground/christmas"
	"github.com/goodleby/playground/graph"
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

	// Graph example
	g, err := graph.New(map[int][]int{
		0: {4, 6, 8},
		1: {2, 5, 8},
		2: {1, 3},
		3: {2, 7},
		4: {0, 5},
		5: {1, 4},
		6: {0, 7},
		7: {3, 6},
		8: {0, 1},
	})
	if err != nil {
		log.Fatalf("Error creating a graph: %v", err)
	}

	log.Print(g.FindShortestPath(0, 3))
}
