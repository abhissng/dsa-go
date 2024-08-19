package runner

import (
	"fmt"
	"github.com/abhissng/dsa-go/MachineRound/snakeLadder"
	"github.com/abhissng/dsa-go/arrays"
	"github.com/abhissng/dsa-go/blind150"
	"github.com/abhissng/dsa-go/mathematics"
	"github.com/abhissng/dsa-go/patterns"
	"github.com/abhissng/dsa-go/question"
	"github.com/abhissng/dsa-go/structures"
)

var (
	isSnakeLadder = false
)

func init() {

	fmt.Println("Started DSA - Monday 8th April 2024")

	// Chapters runner
	patterns.Init()
	question.Init()
	mathematics.Init()
	arrays.Init()

	///////////// machine round question
	if isSnakeLadder {
		snakeLadder.Init()
	}
	// Basic data structures here
	structures.Init()

	// Blind 150 question started here
	blind150.Init()
}
