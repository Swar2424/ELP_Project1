package main

import (
	"fmt"
)

func main() {
	Canal := make(chan ([]int))
	dat_numb := Load_data("./textfile_test.txt", 4)
	fmt.Println(dat_numb)

	for i := 0; i < 4; i++ {
		go Dijkstra(dat_numb, i, Canal)
	}
	for i := 0; i < 4; i++ {
		fmt.Println(<-Canal)
	}
}
