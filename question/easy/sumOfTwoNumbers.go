package easy

import "sort"

/*
//naive O(n^2) time / O(1) space
func twoNumSum(arr []int,target int)[]int{
	for i := range arr{
		for j:=i+1;j<len(arr);j++{
			if arr[i] +arr[j] == target{
				return []int{i,j}
			}
		}
	}
	return nil
}
*/

/*
// better solution O(n) time / O(n) space

	func twoNumSum(arr []int, target int) []int {
		seen := make(map[int]int)
		for i := 0; i < len(arr); i++ {
			potentialMatch := target - arr[i]
			if index, ok := seen[potentialMatch]; ok {
				return []int{index, i}
			}
			seen[arr[i]] = i
		}
		return nil
	}
*/

// better solution in case of space but in this case we will not be able to get the index, rather we will get the numbers
//
//	O(nlog(n)) time /  better solution O(1) space
func TwoNumSum(arr []int, target int) []int {
	sort.Ints(arr)
	left := 0
	right := len(arr) - 1

	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == target {
			return []int{arr[left], arr[right]}
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
