package topics

func LargestElement(nums []int) int {
	largest := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[largest] {
			largest = i
		}
	}
	return nums[largest]
}

func SecondLargestElement(nums []int) (int, int) {
	largest := 0
	secondLargest := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[largest] {
			secondLargest = largest
			largest = i
		} else if nums[i] < nums[largest] {
			if secondLargest == -1 || nums[i] > nums[secondLargest] {
				secondLargest = i
			}
		}
	}
	return nums[largest], nums[secondLargest]
}
