package easy

func ContainsDuplicate(nums []int) bool {
	// here we define a map which will store
	// the value of the num if we have already seen it
	seen := make(map[int]bool)

	// iterating over the numbers array
	for _, num := range nums {
		// we are checking if we have seen the number in the map
		if ok := seen[num]; ok {
			return true
		}
		// we have not yet seen this number in the map so better add it
		seen[num] = true
	}
	return false
}
