package topics

func MaxDifferenceBasic(arr []int) int {
	/*
	   	Input : arr = {2, 3, 10, 6, 4, 8, 1}
	      Output : 8
	      Explanation : The maximum difference is between 10 and 2.
	*/
	res := arr[1] - arr[0]
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			res = max(res, arr[j]-arr[i])
		}
	}
	return res
}

func MaxDifference(arr []int) int {
	/*
	   	Input : arr = {2, 3, 10, 6, 4, 8, 1}
	      Output : 8
	      Explanation : The maximum difference is between 10 and 2.
	*/
	res, minValue := arr[1]-arr[0], arr[0]
	for j := 1; j < len(arr); j++ {
		res = max(res, arr[j]-minValue)
		minValue = min(minValue, arr[j])
	}
	return res
}
