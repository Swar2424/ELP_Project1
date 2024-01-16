package main

import (
	"fmt"
	"slices"
	"time"
)

func main() {

	time_start := time.Now()
	n := 2000
	dat_numb := Load_data("./data", n)
	table := make([]int, n)

	for i := 0; i < n; i++ {
		dij := Dijkstra(dat_numb, i)
		for _, slice := range dij {
			table = Counting(slice, dij, i, table)
		}
	}

	time_end := time.Now()
	fmt.Println(slices.Max(table), slices.Min(table))
	fmt.Println(time_end.Sub(time_start))
}
