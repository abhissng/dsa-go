package easy

import (
	"sort"

	"github.com/abhissng/dsa-go/helper"
)

// O(nlogn) time and O(n)  space complexity
func SortedSquareArray(arr []int) []int {
	output := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		output[i] = arr[i] * arr[i]
	}
	sort.Ints(output)
	return output

}

// O(n) time complexity and O(n) space complexity
func EfficientSortedSquareArray(arr []int) []int {
	smallerValueIdx, largerValueIdx := 0, len(arr)-1
	output := make([]int, len(arr))

	for idx := len(arr) - 1; idx >= 0; idx-- {
		smallerValue := arr[smallerValueIdx]
		largerValue := arr[largerValueIdx]

		if helper.Abs(int64(smallerValue)) > helper.Abs(int64(largerValue)) {
			output[idx] = smallerValue * smallerValue
			smallerValueIdx++
		} else {
			output[idx] = largerValue * largerValue
			largerValueIdx--
		}
	}
	return output
}
