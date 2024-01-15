package main

import "strconv"

func Create_matrix(N int) [][]int {
	a := make([][]int, N)
	for i := range a {
		a[i] = make([]int, N)
	}
	return a
}

func tab_to_str(data [][]int) string {
	var tab string
	n := len(data)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tab += strconv.Itoa(data[i][j]) + " "
		}
		tab += "\r\n"
	}
	return tab
}

/*
func str_to_tab(data string)) [][]int {
	var tab [][]int
	n := len(data)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tab += strconv.Itoa(data[i][j]) + " "
		}
		tab += "\r\n"
	}
	return tab
}
*/
