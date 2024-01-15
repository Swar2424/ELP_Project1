package main

import (
	"fmt"
	"runtime"
	"slices"
	"time"
)

func main() {
	go_num := runtime.NumCPU() / 2
	time_start := time.Now()
	n := 2000
	Jobs := make(chan int, n)
	Results_dij := make(chan Table_int, n)
	Results_count := make(chan (int))
	Count_go := make(chan (int))
	dat_numb := Load_data("./data", n)
	table := make([]int, n)

	for i := 0; i < go_num; i++ {
		go Worker_dijkstra(dat_numb, Jobs, Results_dij)
		go Worker_counter(Results_dij, Results_count, Count_go)
	}

	go Launcher(n, Jobs, Results_dij, Results_count, Count_go)

	for i := range Results_count {
		table[i] += 1
	}

	time_end := time.Now()
	fmt.Println(slices.Max(table), slices.Min(table))
	fmt.Println(time_end.Sub(time_start))
}
