package main

import (
	"C"
)

//export Fibonacci
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// does not work at arm64
// func Snark_run() {
// 	groth16.Groth16_run()

//		plonk.Plonk_run()
//	}
func main() {}
