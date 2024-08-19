package topics

// This will print the leaders in the format which has been provided
func LeadersBasic(arr []int) []int {
	/*
			[7,10,4,10,6,5,2]
		output = 10,6,5,2

	*/
	output := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		flag := true
		for j := i + 1; j < len(arr); j++ {
			if arr[j] >= arr[i] {
				flag = false
				break
			}
		}
		if flag {
			output = append(output, arr[i])
		}
	}
	return output
}

// This will print the leaders in the reverse order which has been provided
func LeadersBetter(arr []int) []int {
	/*
			[7,10,4,10,6,5,2]
		output = 2,5,6,10

	*/
	currentLeader := arr[len(arr)-1]
	output := make([]int, 0)
	output = append(output, currentLeader)
	for i := len(arr) - 2; i > 0; i-- {
		if currentLeader < arr[i] {
			currentLeader = arr[i]
			output = append(output, currentLeader)
		}
	}
	return output
}
