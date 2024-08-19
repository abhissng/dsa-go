package topics

func LeftRotateByOne(arr []int) []int {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
	return arr
}

// naive Solution

func LeftRotateByD(arr []int, d int) []int {
	temp := make([]int, d)

	for i := 0; i < d; i++ {
		temp[i] = arr[i]
	}

	for j := d; j < len(arr); j++ {
		arr[j-d] = arr[j]
	}

	for k := 0; k < d; k++ {
		arr[len(arr)-d+k] = temp[k]
	}
	return arr
}

func LeftRotateByDBetter(arr []int, d int) []int {
	ReverseArrayByLowAndHigh(arr, 0, d-1)
	ReverseArrayByLowAndHigh(arr, d, len(arr)-1)
	ReverseArrayByLowAndHigh(arr, 0, len(arr)-1)

	return arr
}
