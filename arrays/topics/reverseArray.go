package topics

func ReverseArray(arr []int) []int {
	/*

		10,5,7,30
	*/
	low := 0
	high := len(arr) - 1
	for low < high {
		temp := arr[low]
		arr[low] = arr[high]
		arr[high] = temp
		low++
		high--
	}
	return arr
}

func ReverseArrayByLowAndHigh(arr []int, low, high int) []int {
	for low < high {
		temp := arr[low]
		arr[low] = arr[high]
		arr[high] = temp
		low++
		high--
	}
	return arr
}
