package main

import (
	"os"
)

func Load_data(file string, N int) string {
	b := make([]byte, 1)
	//dat_numb := Create_matrix(N)
	var dat_str, dat_str_list string

	dat, _ := os.Open(file)

	for i := 0; i < N; {

		for j := 0; j < N-1; {
			dat_read, _ := dat.Read(b)
			dat_str = string(b[:dat_read])
			dat_str_list += dat_str
			if dat_str == " " {
				j++
			}
		}
		i++
		dat.Read(b)
		dat.Read(b)
		dat_str_list += "\r\n"
	}

	return (dat_str_list)
}
