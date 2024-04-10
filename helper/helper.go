package helper

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		in[i] = castVariableType(arg)

	}
	fmt.Println()
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+ | " + getFunctionName(inputFunc) + " | -+-+-+-+-+-+-+-+-+-+-+-+-+-+")
	now := startLatency()

	// Call the function with provided arguments
	result := fnValue.Call(in)

	latency := checkLatency(now)

	fmt.Printf("Input : ")
	fmt.Printf("%+v\n", args)

	// Print the result (if any)
	for _, r := range result {
		fmt.Printf("Output : ")
		fmt.Printf("%+v\n", r.Interface())
	}
	fmt.Printf("Execution latency time : %s\n", latency)
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
	// Extract just the function name
	lastSlashIndex := len(fullName) - 1
	for lastSlashIndex >= 0 && fullName[lastSlashIndex] != '/' {
		lastSlashIndex--
	}
	// Format the function name
	var funcName string
	splitName := strings.Split(fullName[lastSlashIndex+1:], ".")
	caser := cases.Title(language.Und)
	if len(splitName) == 2 {
		splitName[0] = caser.String(splitName[0])
		funcName = strings.Join(splitName, ".")
	} else {
		funcName = caser.String(fullName[lastSlashIndex+1:])
	}

	return funcName
}

// this is to cast the value of int and float32 so that reflect does not throw any exception
func castVariableType(value interface{}) reflect.Value {
	valueRef := reflect.ValueOf(value)

	switch valueRef.Kind() {
	case reflect.Int:
		return reflect.ValueOf(int64(valueRef.Int()))
	case reflect.Float32:
		return reflect.ValueOf(float64(valueRef.Float()))
	default:
		return valueRef
	}
}
