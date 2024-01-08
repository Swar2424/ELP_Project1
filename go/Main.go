package main

import "fmt"

func main() {
	Canal := make(chan ([]int))
	dat_numb := Load_data("./textfile_test.txt", 4)
	fmt.Println(dat_numb)

	for i := 1; i < 5; i++ {
		go Dijkstra(dat_numb, i, Canal)
	}
	for i := 1; i < 5; i++ {
		fmt.Println(<-Canal)
	}
}
