package main

import (
	"os"
	"strconv"
)

func Load_data(file string, N int) [][]int {
	b := make([]byte, 1)
	dat_numb := Create_matrix(N)
	var dat_str, dat_str_list string

	dat, _ := os.Open(file)

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

	return (dat_numb)
}
