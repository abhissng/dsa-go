package mathematics

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
		})
	}
}

func Test_countDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigits(tt.args.num); got != tt.want {
				t.Errorf("countDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.num); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factorial(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.args.num); got != tt.want {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_naiveTrailingZero(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := naiveTrailingZero(tt.args.val); got != tt.want {
				t.Errorf("naiveTrailingZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_efficientTrailingZero(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := efficientTrailingZero(tt.args.val); got != tt.want {
				t.Errorf("efficientTrailingZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_naiveGreatestCommonDivisor(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := naiveGreatestCommonDivisor(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("naiveGreatestCommonDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEfficientGreatestCommonDivisor(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EfficientGreatestCommonDivisor(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("EfficientGreatestCommonDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leastCommonDivisor(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastCommonDivisor(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("leastCommonDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEfficientLeastCommonDivisor(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EfficientLeastCommonDivisor(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("EfficientLeastCommonDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_naiveCheckPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := naiveCheckPrime(tt.args.n); got != tt.want {
				t.Errorf("naiveCheckPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_betterCheckPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := betterCheckPrime(tt.args.n); got != tt.want {
				t.Errorf("betterCheckPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEfficientCheckPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EfficientCheckPrime(tt.args.n); got != tt.want {
				t.Errorf("EfficientCheckPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_naiveGetDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := naiveGetDivisors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("naiveGetDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEfficientGetDivisor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EfficientGetDivisor(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EfficientGetDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}
