package main

import (
	"encoding/gob"
	"io"
)

func Create_matrix(N int) [][]int {
	a := make([][]int, N)
	for i := range a {
		a[i] = make([]int, N)
	}
	return a
}

type Matrix struct {
	Rows    int
	Columns int
	Data    [][]int
}

func sendMatrix(writer io.Writer, matrix Matrix) error {
	encoder := gob.NewEncoder(writer)
	err := encoder.Encode(matrix)
	if err != nil {
		return err
	}
	return nil
}

func receiveMatrix(reader io.Reader) (Matrix, error) {
	var matrix Matrix

	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&matrix)
	if err != nil {
		return matrix, err
	}

	return matrix, nil
}
