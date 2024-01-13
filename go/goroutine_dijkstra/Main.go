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
	Canal3 := make(chan (int))
	n := 10
	dat_numb := Load_data("./data", n)
	table := make([]int, n)
	//fmt.Println(dat_numb)

	go Launch_Dijkstra(dat_numb, Canal1)
	go Launch_Counting(dat_numb, Canal1, Canal2, Canal3)

	for i := range Canal2 {
		table[i] += 1
	}

	time_end := time.Now()
	fmt.Println(table)
	fmt.Println(time_end.Sub(time_start))
}
