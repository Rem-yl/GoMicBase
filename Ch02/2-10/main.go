package main

import "fmt"

func Sum(n int) int {
	total := 0

	for i := 0; i <= n; i++ {
		total += i
	}

	return total
}

func main() {
	n := 2
	res := Sum(n)
	fmt.Println(res)
}
