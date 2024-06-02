package easy

/*
// O(n) Time & O(1) Space
func validateSubSequence(array, sequence []int) bool {
	arrIdx, seqIdx := 0, 0

	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}
		arrIdx++
	}
	return seqIdx == len(sequence)
}
*/

// O(n) Time & O(1) Space
func ValidateSubSequence(array, sequence []int) bool {
	seqIdx := 0

	for i := 0; i < len(array); i++ {
		if seqIdx == len(sequence) {
			return true
		}
		if sequence[seqIdx] == array[i] {
			seqIdx++
		}
	}
	return seqIdx == len(sequence)
}
