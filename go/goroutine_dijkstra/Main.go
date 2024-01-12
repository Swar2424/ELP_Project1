package main

import (
	"fmt"
	"time"
)

type Table_int struct {
	Table  *[][]int
	Origin int
}

func main() {
	time_start := time.Now()
	Canal1 := make(chan Table_int)
	Canal2 := make(chan (int))
	n := 300
	dat_numb := Load_data("./data", n)
	table := make([]int, n)

	go Launch_Dijkstra(dat_numb, Canal1)
	go Launch_Counting(dat_numb, Canal1, Canal2)

	for i := 0; i < n*n; i++ {
		table[<-Canal2] += 1
	}

	//fmt.Println(table)

	time_end := time.Now()
	fmt.Println(time_end.Sub(time_start))
}
