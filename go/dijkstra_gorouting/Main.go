package main

import (
	"fmt"
	"time"
)

func main() {
	time_start := time.Now()
	Canal := make(chan ([]int))
	n := 500
	dat_numb := Load_data("./data", n)

	for i := 0; i < n; i++ {
		go Dijkstra(dat_numb, i, Canal)
	}
	for i := 0; i < n; i++ {
		fmt.Println(<-Canal)
	}

	time_end := time.Now()
	fmt.Println(time_end.Sub(time_start))
}
