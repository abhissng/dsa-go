package arrays

import (
	"github.com/abhissng/dsa-go/arrays/topics"
	"github.com/abhissng/dsa-go/helper"
)

func Init() {
	helper.Execute(topics.LargestElement, []int{5, 8, 29, 10})
	helper.Execute(topics.SecondLargestElement, []int{5, 8, 29, 10})
	helper.Execute(topics.SortedArray, []int{5, 8, 29, 10})
	helper.Execute(topics.ReverseArray, []int{5, 8, 29, 10})
	helper.Execute(topics.RemoveDuplicates, []int{10, 10, 20, 20, 20, 30, 30, 30, 30})
	helper.Execute(topics.MoveZeroes, []int{10, 20, 0, 0, 30, 0, 40, 0, 0, 30})
	helper.Execute(topics.LeftRotateByOne, []int{10, 12, 18, 1, 3})
	helper.Execute(topics.LeftRotateByD, []int{10, 12, 18, 1, 3, 6}, 3)
	helper.Execute(topics.LeftRotateByDBetter, []int{10, 12, 18, 1, 3, 6}, 3)
	helper.Execute(topics.LeadersBasic, []int{7, 10, 4, 10, 6, 5, 2})
	helper.Execute(topics.LeadersBetter, []int{7, 10, 4, 10, 6, 5, 2})
	helper.Execute(topics.MaxDifferenceBasic, []int{2, 3, 10, 6, 4, 8, 1})
	helper.Execute(topics.MaxDifference, []int{2, 3, 10, 6, 4, 8, 1})
	helper.Execute(topics.Frequency, []int{10, 10, 10, 30, 30, 40})
	helper.Execute(topics.FrequencyBetter, []int{10, 10, 10, 30, 30, 40})
	helper.Execute(topics.StockBasic, []int{1, 5, 3, 8, 12})
	helper.Execute(topics.StockBetter, []int{1, 5, 3, 8, 12})
}
