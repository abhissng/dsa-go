package helper

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"
)

func Execute(inputFunc any, args ...any) {
	// Check if inputFunc is a function
	fnValue := reflect.ValueOf(inputFunc)
	if fnValue.Kind() != reflect.Func {
		fmt.Println("Invalid function type")
		return
	}

	// Prepare arguments
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		// in[i] = castVariableType(arg)
		in[i] = reflect.ValueOf(arg)
	}
	fmt.Println()
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+ | " + getFunctionName(inputFunc) + " | -+-+-+-+-+-+-+-+-+-+-+-+-+-+")

	for i, a := range args {
		fmt.Printf("Input %d: ", i+1)
		fmt.Printf("%+v\n", a)
	}

	now := startLatency()

	// Call the function with provided arguments
	result := fnValue.Call(in)

	latency := checkLatency(now)

	// Print the result (if any)
	for _, r := range result {
		fmt.Printf("Output: ")
		fmt.Printf("%+v\n", r.Interface())
	}
	fmt.Printf("\nExecution latency time : %s\n", latency)
}
func startLatency() time.Time {
	return time.Now()
}
func checkLatency(t time.Time) time.Duration {
	return time.Since(t)
}

// getFunctionName returns the name of the given function
func getFunctionName(f interface{}) string {
	// Get the name of the function from its runtime representation
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()

	split := strings.Split(fullName, "/")

	var funcName string
	for i := 3; i < len(split); i++ {
		if strings.Contains(split[i], ".") {
			funcSplit := strings.Split(split[i], ".")
			funcName += funcSplit[0] + " -> " + funcSplit[1]
		} else {
			funcName += split[i] + " -> "
		}
	}
	return funcName
}

/*
// this is to cast the value of int and float32 so that reflect does not throw any exception
func castVariableType(value interface{}) reflect.Value {
	valueRef := reflect.ValueOf(value)

	switch valueRef.Kind() {
	case reflect.Int:
		return reflect.ValueOf(int64(valueRef.Int()))
	case reflect.Float32:
		return reflect.ValueOf(float64(valueRef.Float()))
	case reflect.Slice:
		fmt.Println("here",reflect.ValueOf(valueRef))
		// Handle slice conversion
		if valueRef.Type().Elem().Kind() == reflect.Int {
			// Convert []int to []int64
			newSlice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(int64(0))), valueRef.Len(), valueRef.Cap())
			for i := 0; i < valueRef.Len(); i++ {
				newSlice.Index(i).Set(reflect.ValueOf(int64(valueRef.Index(i).Int())))
			}
			return newSlice
		}

	default:
		return valueRef
	}
	return valueRef

}
*/

// Abs returns the absolute value of x.
func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// SwapArr Swap the value of using high and low value.
func SwapArr(arr []int, low, high int) []int {
	temp := arr[low]
	arr[low] = arr[high]
	arr[high] = temp
	return arr
}

type Number interface {
	int | int64
}

// Max gives the maximum between two integer values
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min gives the minimum between two integer values
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func SortString(s string) string {
	// convert string to rune so that we can compare the rune value as rune is int32
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		// in ascending order
		return runes[i] < runes[j]
	})
	return string(runes)
}
