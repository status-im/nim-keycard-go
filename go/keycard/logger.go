package main

import "fmt"

func l(format string, args ...interface{}) {
	f := fmt.Sprintf("keycard - %s\n", format)
	fmt.Printf(f, args...)
}
