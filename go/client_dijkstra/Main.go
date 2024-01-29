package main

import (
	"fmt"
	"net"
	"slices"
	"time"
)

func main() {
	// On se connecte au serveur
	time_start := time.Now()
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	n := 1000
	matrice := Load_data("./data", n)
	// data : la matrice que l'on souhaite envoyer au serveur
	data := Matrix{
		Rows:    n,
		Columns: n,
		Data:    matrice,
	}
	// On envoie la matrice au serveur grâce à la connexion TCP
	err = sendMatrix(conn, data)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la matrice :", err)
		return
	}

	// On récupère la matrice que le serveur nous envoie
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
