package main

import (
	"fmt"
	"net"
	"slices"
	"time"
)

func main() {
	// Connect to the server
	time_start := time.Now()
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send data to the server
	n := 1000
	matrice := Load_data("./data", n)
	data := Matrix{
		Rows:    n,
		Columns: n,
		Data:    matrice,
	}
	err = sendMatrix(conn, data)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la matrice :", err)
		return
	}

	// Read and process data from the server
	matrix, err := receiveMatrix(conn)
	if err != nil {
		fmt.Println("Erreur lors de la réception de la matrice :", err)
		return
	}
	time_end := time.Now()
	fmt.Println(matrix.Data)
	fmt.Println(slices.Max(matrix.Data[0]), slices.Min(matrix.Data[0]))
	fmt.Println(time_end.Sub(time_start))
}
