package main

import (
	"fmt"
	"time"
)

func main() {
	time_start := time.Now()
	Canal1 := make(chan ([][]int))
	Canal2 := make(chan (int))
	Canal3 := make(chan (int))
	n := 5
	dat_numb := Load_data("./data", n)
	table := make([]int, n)

	go Launch_Dijkstra(dat_numb, Canal1)
	go Launch_Counting(dat_numb, Canal1, Canal2, Canal3)

	for i := range Canal2 {
		table[i] += 1
	}

	fmt.Println(table)

	time_end := time.Now()
	fmt.Println(time_end.Sub(time_start))
}
