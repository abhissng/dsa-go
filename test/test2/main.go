package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func main() {
	s := "my interview is scheduled today, my interview is going on"
	fmt.Println("Output ")
	fmt.Printf("%+v", findRepeatingWords(s))

}

func findRepeatingWords(sentence string) map[string]int {

	wordCount := make(map[string]int)

	words := strings.Fields(sentence)

	for _, word := range words {
		wordCount[word]++
	}

	max := 0

	for _, count := range wordCount {
		if count > max {
			max = count
		}
	}
	output := make(map[string]int)
	for word, count := range wordCount {
		if count == max {
			output[word] = count
		}
	}

	return output

}

func testFindRepeatigWords(t *testing.T) {
	tests := []struct {
		sentence string
		expected map[string]int
	}{{

		"my interview is scheduled today, my interview is going on",
		map[string]int{"interview": 2, "is": 2, "my": 2},
	}, {
		"Go is fun, Go is used for backend",
		map[string]int{"Go": 2, "is": 2},
	}, {
		"This is a test",
		map[string]int{"This": 1, "is": 1, "a": 1, "test": 1},
	},
	}

	for _, test := range tests {
		t.Run(test.sentence, func(t *testing.T) {
			actual := findRepeatingWords(test.sentence)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Expected %v, but receiving  %v", test.expected, actual)
			}
		})
	}

}
