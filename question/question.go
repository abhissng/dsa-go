package question

import (
	"github.com/abhissng/dsa-go/helper"
	"github.com/abhissng/dsa-go/question/easy"
)

func Init() {
	helper.Execute(easy.TwoNumSum, []int{5, 1, 3, 11, 7, 15}, 10)
	helper.Execute(easy.ValidateSubSequence, []int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10})
	helper.Execute(easy.SortedSquareArray, []int{-7, -5, -4, 3, 6, 8, 9})
	helper.Execute(easy.EfficientSortedSquareArray, []int{-7, -5, -4, 3, 6, 8, 9})
	competitions := [][]string{
		{"HTML", "Javascript"},
		{"C#", "Go"},
		{"Go", "Javascript"},
	}
	helper.Execute(easy.TournamentWinner, competitions, []int{0, 0, 1})
}
