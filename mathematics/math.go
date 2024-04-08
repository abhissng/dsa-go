package mathematics

import (
	"github.com/abhissng/dsa-go/helper"
)

func Init() {
	helper.Print(countDigits, 1234)
	helper.Print(isPalindrome, 78987)
	helper.Print(factorial, 20)
}

// counts the digit in a number
func countDigits(num int) int {

	res := 0
	for num > 0 {
		num /= 10
		res++
	}

	return res
}

// check if the number is palindrome
func isPalindrome(num int) bool {
	res, original := int(0), num
	for num > 0 {
		res = (res * 10) + (num % 10)
		num /= 10
	}
	return res == original
}

// Check factorial
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
