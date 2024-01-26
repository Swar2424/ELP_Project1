package main

import (
	"fmt"
	"net"
	"runtime"
	"slices"
)

// Fonction pour lancer dijkstra avec les goroutines
func dijkstra_goroutine(dat_numb [][]int) []int {
	n := len(dat_numb)
	go_num := runtime.NumCPU() / 2
	Jobs := make(chan int, n)
	Results_dij := make(chan Table_int, n)
	Results_count := make(chan (int))
	Count_go := make(chan (int))
	table := make([]int, n)

	for i := 0; i < go_num; i++ {
		go Worker_dijkstra(dat_numb, Jobs, Results_dij)
		go Worker_counter(Results_dij, Results_count, Count_go)
	}

	go Launcher(n, Jobs, Results_dij, Results_count, Count_go)

	for i := range Results_count {
		table[i] += 1
	}

	fmt.Println(slices.Max(table), slices.Min(table))
	return table
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	n := 1000
	// On utilise la fonction receiveMatrix pour récupérer la matrice envoyée par les clients
	matrix, err := receiveMatrix(conn)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	Slice := dijkstra_goroutine(matrix.Data)
	//On récupère dans une ligne de matrice la matrice renvoyée par Dijkstra
	Donnees := make([][]int, 1)
	Donnees[0] = Slice
	data := Matrix{
		Rows:    1,
		Columns: n,
		Data:    Donnees,
	}
	// On envoie la matrice optenue avec Dijkstra au client
	err = sendMatrix(conn, data)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la matrice :", err)
		return
	}
}

func main() {

	//On ouvre le canal TCP
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8000")

	// On accepte plusieurs clients en parallèles grâce aux goroutines
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleClient(conn)
	}
}
