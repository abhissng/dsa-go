package main

import (
	"fmt"
	"github.com/abhissng/dsa-go/patterns"

	snakeLadder "github.com/abhissng/dsa-go/MachineRound/SnakeLadder"
	"github.com/abhissng/dsa-go/mathematics"
	"github.com/abhissng/dsa-go/question"
)

var (
	isSnakeLadder = false
)

func main() {
	fmt.Println("Started DSA - Monday 8th April 2024")

	// Chapters runner
	mathematics.Init()
	question.Init()
	patterns.Init()

	///////////// machine round question
	if isSnakeLadder {
		snakeLadder.Init()
	}

}
