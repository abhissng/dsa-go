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
	helper.Execute(leastCommonDivisor, 4, 6)
	helper.Execute(EfficientLeastCommonDivisor, 4, 6)
	helper.Execute(naiveCheckPrime, 1031)
	helper.Execute(betterCheckPrime, 1031)
	helper.Execute(EfficientCheckPrime, 1031)
	helper.Execute(naiveGetDivisors, 25)
	helper.Execute(EfficientGetDivisor, 25)

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
	res, original := 0, num
	for num > 0 {
		res = (res * 10) + (num % 10)
		num /= 10
	}
	return res == original
}

// Check factorial
func factorial(num int) int64 {
	if num == 0 {
		return 1
	}
	return int64(num) * int64(factorial(num-1))
}

// Naive CountTrailing zero
func naiveTrailingZero(val int) int {
	fact := factorial(val)
	res := 0
	for fact%10 == 0 {
		res++
		fact /= 10
	}
	return res
}

// Efficient CountTrailing zero
func efficientTrailingZero(val int) int {
	res :=0
	for i := 5; i <= val; i = i * 5 {
		res = res + (val / i)
	}
	return res
}

// Naive Greatest common divisor
func naiveGreatestCommonDivisor(a, b int) int {
	res := int(math.Min(float64(a), float64(b)))
	for res > 0 {
		if a%res == 0 && b%res == 0 {
			break
		}
		res--
	}
	return res
}

// EfficientGreatestCommonDivisor
func EfficientGreatestCommonDivisor(a, b int) int {
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

// LeastCommonDivisor returns the least common multiple of a and b
func leastCommonDivisor(a, b int) int {
	res := int(math.Max(float64(a), float64(b)))
	for {
		if res%a == 0 && res%b == 0 {
			return res
		}
		res++
	}
}

// a*b = gcd(a, b) * lcm(a, b)
// using this formula we can calculate the lcm efficiently
func EfficientLeastCommonDivisor(a, b int) int {
	return (a * b) / EfficientGreatestCommonDivisor(a, b)
}

func naiveCheckPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func betterCheckPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func EfficientCheckPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func naiveGetDivisors(n int) []int{
	arr := make([]int, 0)
	arr = append(arr, 1)
	// fmt.Printf("%d ", 1)
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			arr = append(arr, i)
			// fmt.Printf("%d ", i)
		}

	}
	return arr
}

func EfficientGetDivisor(n int) []int {
	arr := make([]int, 0)
	/*
		for i := int64(1); i*i <= n; i++ {
			if n%i == 0 {
				arr = append(arr, i)

				if i != n/i {
					arr = append(arr, n/i)
				}
			}
		}
	*/

	// To print in ascending orders
	var i int
	for i = 1; i*i <= n; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}

	for ; i >= 1; i-- {
		if i != n/i && n%i == 0 {
			arr = append(arr, n/i)
		}
	}
	return arr
}
