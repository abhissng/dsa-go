package mathematics

import (
	"math"

	"github.com/abhissng/dsa-go/helper"
)

func Init() {
	helper.Execute(countDigits, 1234)
	helper.Execute(isPalindrome, 78987)
	helper.Execute(factorial, 20)
	helper.Execute(naiveTrailingZero, 20)
	helper.Execute(efficientTrailingZero, 20)
	helper.Execute(naiveGreatestCommonDivisor, 10, 15)
	helper.Execute(EfficientGreatestCommonDivisor, 12, 15)
}

// counts the digit in a number
func countDigits(num int64) int64 {

	res := int64(0)
	for num > 0 {
		num /= 10
		res++
	}

	return res
}

// check if the number is palindrome
func isPalindrome(num int64) bool {
	res, original := int64(0), num
	for num > 0 {
		res = (res * 10) + (num % 10)
		num /= 10
	}
	return res == original
}

// Check factorial
func factorial(num int64) int64 {
	if num == 0 {
		return 1
	}
	return int64(num) * int64(factorial(num-1))
}

// Naive CountTrailing zero
func naiveTrailingZero(val int64) int {
	fact := factorial(val)
	res := 0
	for fact%10 == 0 {
		res++
		fact /= 10
	}
	return res
}

// Efficeint CountTrailing zero
func efficientTrailingZero(val int64) int64 {
	res := int64(0)
	for i := int64(5); i <= val; i = i * 5 {
		res = res + (val / i)
	}
	return res
}

// Naive Greatest common divisor
func naiveGreatestCommonDivisor(a, b int64) int64 {
	res := int64(math.Min(float64(a), float64(b)))
	for res > 0 {
		if a%res == 0 && b%res == 0 {
			break
		}
		res--
	}
	return res
}

// Efficient Greatest common divisor
func EfficientGreatestCommonDivisor(a, b int64) int64 {
	// Efficient but needed a lot of substraction
	/*
		for a != b {
			if a > b {
				a = a - b
			} else {
				b = b - a
			}
		}

		return a
	*/
	if b == 0 {
		return a
	} else {
		return EfficientGreatestCommonDivisor(b, a%b)
	}
}
