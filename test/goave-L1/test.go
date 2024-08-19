/*package main

import "fmt"

/*

Inputs

array/slice[ 1 2 3 4 5 6 7 8 9 10 ]

stack factor 4 [] [] [] []

all the left over elements should be assigned to last array



output

[1 2 ] [3 4] [5 6] [7 8 9 10]
*/
/*
func SplitArr(arr []int, stackFactor int) [][]int {
	var result [][]int
	length := len(arr)

	// size for all the sub aray
	size := length / stackFactor

	for i := 0; i < stackFactor; i++ {
		left := i * size
		right := left + size

		if i == stackFactor-1 { // for the leftover elements
			right = length
		}
		result = append(result, arr[left:right])

	}
	return result
}

func main() {
	output := SplitArr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	fmt.Printf("%+v", output)
*/
package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {

		close(ch)

		ch <- 1

		ch <- 2

	}()

	for v := range ch {

		fmt.Println(v)

	}

}
