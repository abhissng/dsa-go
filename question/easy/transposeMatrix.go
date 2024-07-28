package easy

func TransposeMatrix(matrix [][]int) [][]int {
	var transposeMatrix [][]int
	for col := 0; col < len(matrix[0]); col++ {
		var newRow []int
		for row := 0; row < len(matrix); row++ {
			newRow = append(newRow, matrix[row][col])
		}
		transposeMatrix = append(transposeMatrix, newRow)
	}
	return transposeMatrix

}
