package wavyword

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type WavyWord struct {
	Word  string
	Lines [][]string
}

func New(word string, waviness int) (*WavyWord, error) {
	if word == "" {
		return nil, errors.New("cannot create a wavy word from an empty string")
	}

	if waviness < 0 {
		return nil, fmt.Errorf("waviness cannot be negative, received: %d", waviness)
	}

	linesAmount := waviness*2 + 1
	initialY := linesAmount / 2
	waveLength := waviness * 4

	// Create resulting lines matrix of the correct size containing just spaces.
	lines := make([][]string, linesAmount)
	for y := range lines {
		line := make([]string, len(word))
		for x := range line {
			line[x] = " "
		}
		lines[y] = line
	}

	// Set characters to their positions according to triangular wave function:
	// y = |(x + y0 + L/2)%L - L/2|
	for x, char := range strings.Split(word, "") {
		var y int
		if waviness != 0 {
			y = int(math.Abs(float64((x+initialY+waveLength/2)%waveLength - waveLength/2)))
		} else {
			y = 0
		}

		if len(lines) <= y {
			return nil, fmt.Errorf("y is outside of waving range: f(%d)=%d", x, y)
		}

		lines[y][x] = char
	}

	wavyWord := WavyWord{
		Word:  word,
		Lines: lines,
	}

	return &wavyWord, nil
}

func (ww *WavyWord) Print() {
	for _, line := range ww.Lines {
		fmt.Println(strings.Join(line, ""))
	}
}
