package main

import (
	"os"
	"strconv"
)

// Fonction qui permet de récupérer une matrice d'un fichier
func Load_data(file string, N int) [][]int {
	b := make([]byte, 1)
	dat_numb := Create_matrix(N)

	// On doit creer la variable dat_str_list dans le cas ou on a un nombre et non un chiffre à mettre dans la matrice dat_numb
	var dat_str, dat_str_list string

	dat, _ := os.Open(file)

	for i := 0; i < N; {
		for j := 0; j < N; {
			// On lit a chaque fois un byte de la matrice et on met la variable dans b
			// On la convertit par la suite en string
			dat.Read(b)
			dat_str = string(b)

			if dat_str != " " {
				dat_str_list += dat_str
			} else {
				// A chaque fois qu'on rencontre un esp, on ajoute le nombre dans la matrice dat_numb
				// et on réinitialise dat_str_list
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
