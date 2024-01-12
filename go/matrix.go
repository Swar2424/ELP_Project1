package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func random_adjacency_matrix(n int) [][]int {
	matrix := Create_matrix(n)
	for i := 0; i < n; i++ {
		matrix[i][i] = 0
		for j := 0; j < n; j++ {
			matrix[j][i] = rand.Intn(10) + 1
			matrix[i][j] = matrix[j][i]
		}
	}
	return matrix
}

func main() {
	N := 6
	file, err := os.Create("./gorouting_dijkstra/data")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	matrix := random_adjacency_matrix(N)
	writer := bufio.NewWriter(file)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			number := strconv.Itoa(matrix[i][j])
			_, err = fmt.Fprint(writer, number+" ")
			if err != nil {
				panic(err)
			}
		}
		_, err = fmt.Fprint(writer, "\n")
		if err != nil {
			panic(err)
		}
		err = writer.Flush()
		if err != nil {
			panic(err)
		}
	}
}
