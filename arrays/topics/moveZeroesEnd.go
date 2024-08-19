package topics

import "github.com/abhissng/dsa-go/helper"

func MoveZeroes(nums []int) []int {

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			helper.SwapArr(nums, i, count)
			count++
		}
	}
	return nums
}
