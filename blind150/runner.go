package blind150

import (
	"github.com/abhissng/dsa-go/blind150/array/easy"
	"github.com/abhissng/dsa-go/helper"
)

func Init() {
	helper.Execute(easy.ContainsDuplicate, []int{1, 2, 3, 6, 7, 1})
	helper.Execute(easy.IsAnagramBrute, "anagram", "nagaram")
	helper.Execute(easy.IsAnagram, "anagram", "nagara")
	helper.Execute(easy.IsAnagramUsingSorting, "anagram", "nagaram")
	helper.Execute(easy.TwoSum, []int{3, 2, 4}, 6)
}
