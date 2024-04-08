package helper

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Print(inputFunc any, args ...any) {
	// Check if inputFunc is a function
	fnValue := reflect.ValueOf(inputFunc)
	if fnValue.Kind() != reflect.Func {
		fmt.Println("Invalid function type")
		return
	}

	// Prepare arguments
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}
	funcName := getFunctionName(inputFunc)
	caser := cases.Title(language.Und)
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+ | " + caser.String(funcName) + " | -+-+-+-+-+-+-+-+-+-+-+-+-+-+")
	fmt.Printf("Input : ")
	fmt.Printf("%+v\n", args)

	now := startLatency()

	// Call the function with provided arguments
	result := fnValue.Call(in)

	checkLatency(now)

	// Print the result (if any)
	for _, r := range result {
		fmt.Printf("Output : ")
		fmt.Printf("%+v\n", r.Interface())
	}
	fmt.Println()
}
func startLatency() time.Time {
	return time.Now()
}
func checkLatency(t time.Time) {
	fmt.Printf("Execution latency time : %s\n", time.Since(t))
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
	return fullName[lastSlashIndex+1:]
}
