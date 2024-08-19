package topics

func RemoveDuplicates(arr []int) ([]int, int) {

	distinctElement := 1
	//res := make([]int, 0)
	//res = append(res, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			arr[distinctElement] = arr[i]
			distinctElement++
		}
	}
	for i := distinctElement; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr, distinctElement
}
