package main

import "fmt"

const NMAX int = 100

type comment struct {
	sentimen [NMAX]string
}

func main() {
	fmt.Print("ki")
}

func InsertionSort(A *tabInt, N int) {
	var i, pass int
	var temp int

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp < A[i-1] {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass = pass + 1
	}
}
