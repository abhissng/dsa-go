package topics

func SortedArray(nums []int) bool {
	/*
		[7,20,13,12]

	*/
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}

	return true
}
