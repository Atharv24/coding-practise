package main

import "fmt"

var m map[int]int

// addition: +, subtraction: -, division: /, multiplication: *
// a function that takes in 2 numbers and returns their product
func mup(num1 float32, num2 float32) float32 {
	return num1 / num2
}

func fact(n int) int {
	// base case
	// return 1 when n is 1, dont calculate further
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

// fib seq - fib(n) = fib(n-1) + fib(n-2)
// fib seq - 1 1 2 3 5 8 13 21 34

func fib(n int) int {
	fmt.Printf("calculating fib(%d)\n", n)

	// implement this
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	if ans, ok := m[n]; ok {
		return ans
	}
	answer := fib(n-1) + fib(n-2)
	m[n] = answer
	return answer
}

func main() {
	m = make(map[int]int)

	var n int
	n = 100

	fmt.Println(fib(n))
}
