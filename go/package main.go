package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	b := make([]byte, 1)
	N := 4
	var dat_numb [4]int
	var dat_str, dat_str_list string

	dat, _ := os.Open("./textfile_test.txt")

	for i := 0; i < N; {
		dat_read, _ := dat.Read(b)
		dat_str = string(b[:dat_read])

		if dat_str != " " {
			dat_str_list += dat_str
		} else {
			dat_int, _ := strconv.Atoi(dat_str_list)
			dat_numb[i] = dat_int
			i++
			dat_str_list = ""
		}
	}

	fmt.Println(dat_numb)
}
