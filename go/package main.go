package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	b := make([]byte, 1)
	N := 4
	var dat_numb [4][4]int
	var dat_str, dat_str_list string

	dat, _ := os.Open("./textfile_test.txt")

	for i := 0; i < N; {
		for j := 0; j < N; {
			dat_read, _ := dat.Read(b)
			dat_str = string(b[:dat_read])

			if dat_str != " " {
				dat_str_list += dat_str
			} else {
				dat_int, _ := strconv.Atoi(dat_str_list)
				dat_numb[i][j] = dat_int
				j++
				dat_str_list = ""
			}
		}
		i++
		dat.Read(b)
		dat.Read(b)
	}

	fmt.Println(dat_numb)
	fmt.Println(Dijkstra(dat_numb, 0))
}
