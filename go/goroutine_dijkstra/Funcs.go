package main

func Create_matrix(N int) [][]int {
	a := make([][]int, N)
	for i := range a {
		a[i] = make([]int, N)
	}
	return a
}
