package easy

import "github.com/abhissng/dsa-go/helper"

// IsAnagramBrute - Anagram means if s and t are of the same length and every character counts are the same
func IsAnagramBrute(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	// creating the map to store the counting of all the alphabets in s anf t
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	// converting string to runes
	sRunes := []rune(s)
	tRunes := []rune(t)

	for i := 0; i < len(sRunes); i++ {
		// for every value we are incrementing sMap and tMap
		sMap[sRunes[i]]++
		tMap[tRunes[i]]++
	}
	for key, value := range sMap {
		// if the value in the sMap is not equal to the value in the tMap at any place then it is not an anagram
		if tMap[key] != value {
			return false
		}
	}
	return true
}
func IsAnagram(s, t string) bool {

	if len(s) != len(t) {
		return false
	}
	// here one integer can be declared for the alphabets

	var alp [26]int
	for i := 0; i < len(s); i++ {
		// in case of any one of the string we can increment the value  and decrease for the other
		// in this case as the alp is of type int,
		// and we needed an index and in unicode character a = 0 and b = 1 and so on,
		// while we subtract the value a to it, we can get to the exact unicode index  a-a =0 and b-a = 1

		alp[s[i]-'a']++ // increasing the value
		alp[t[i]-'a']-- // decreasing the value
	}
	for _, value := range alp {
		if value != 0 {
			return false
		}
	}
	return true
}

func IsAnagramUsingSorting(s, t string) bool {
	return helper.SortString(s) == helper.SortString(t)
}
