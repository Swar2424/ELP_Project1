package main

import (
	"fmt"
	"time"
)

func main() {

	time_start := time.Now()
	n := 10
	dat_numb := Load_data("./data", n)
	table := make([]int, n)
	//fmt.Println(dat_numb)
	for i := 0; i < n; i++ {
		dij := Dijkstra(dat_numb, i)
		//fmt.Println(dij)
		for _, slice := range dij {
			table = Counting(slice, dij, i, table)
		}
	}

	fmt.Println(table)

	time_end := time.Now()
	fmt.Println(time_end.Sub(time_start))
}
