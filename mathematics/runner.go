package mathematics

import "github.com/abhissng/dsa-go/helper"

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
	helper.Execute(EfficientGetDivisor, 182)

}
