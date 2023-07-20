package main

import "fmt"

func A000124(n int) int {
	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result += i
	}

	return result
}

func main() {
	var n int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print(A000124(i))
		if i != n-1 {
			fmt.Print("-")
		}
	}
}