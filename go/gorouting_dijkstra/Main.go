package main

import (
	"fmt"
	"time"
)

func main() {
	print("panik \n")
	time_start := time.Now()
	print("panik \n")
	Canal := make(chan ([]int))
	n := 6

	dat_numb := Load_data("./data", n)
	print("panik \n")

	for i := 0; i < n; i++ {
		fmt.Print("zebi\n")
		go Dijkstra(dat_numb, i, Canal)
		print("panik \n")
	}
	for i := 0; i < n; i++ {
		fmt.Print("zbi\n")
		fmt.Println(<-Canal)
	}

	time_end := time.Now()
	fmt.Println(time_end.Sub(time_start))
}
