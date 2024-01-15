package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func random_adjacency_matrix(n int, max int) [][]int {
	matrix := Create_matrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][i] = rand.Intn(max) + 1
			matrix[i][j] = matrix[j][i]
		}
		matrix[i][i] = 0
	}
	return matrix
}

func write(file *os.File, matrix [][]int, N int) {
	writer := bufio.NewWriter(file)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			number := strconv.Itoa(matrix[i][j]) // convertir en string un int
			_, err := fmt.Fprint(writer, number+" ")
			if err != nil {
				panic(err)
			}
		}
		_, err := fmt.Fprint(writer, "\r\n")
		if err != nil {
			panic(err)
		}
		err = writer.Flush()
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	N := 20
	max := 100

	file, err := os.Create("./goroutine_dijkstra/data")
	if err != nil {
		panic(err)
	}
	file2, err2 := os.Create("./normal_dijkstra/data")
	if err2 != nil {
		panic(err2)
	}
	file3, err3 := os.Create("./client_dijkstra/data")
	if err3 != nil {
		panic(err3)
	}

	defer file.Close()
	defer file2.Close()
	defer file3.Close()

	matrix := random_adjacency_matrix(N, max)

	write(file, matrix, N)
	write(file2, matrix, N)
	write(file3, matrix, N)
}
