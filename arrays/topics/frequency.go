package topics

import "fmt"

func Frequency(arr []int) {

	/*

		input
		10, 10, 10, 30, 30, 40
	*/

	freq := 1
	i := 1
	for i < len(arr) {
		for i < len(arr) && arr[i] == arr[i-1] {
			freq++
			i++
		}
		fmt.Println(arr[i-1], " ", freq)
		i++
		freq = 1
	}
	if len(arr) == 1 || arr[len(arr)-1] != arr[len(arr)-2] {
		fmt.Println(arr[i-1], " ", 1)
	}
}
func FrequencyBetter(arr []int) map[int]int {
	seen := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		seen[arr[i]]++
	}
	return seen
}
