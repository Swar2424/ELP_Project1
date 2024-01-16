package main

/*
Take a TCP port on the machine and ask connection attempts to that port to be
redirected to your app
ln, err := net.Listen("tcp", portString) //”:8000”
● Accept a new connection on that port
conn, errconn := ln.Accept()
● Close the connection of a client
conn.Close()
● Receive/Send byte: same as for the client
*/

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

type Table_int struct {
	Table  *[][]int
	Origin int
}

func dijkstra_goroutine(dat_numb [][]int) []int {
	//Canal1 := make(chan Table_int)
	//Canal2 := make(chan (int))
	//Canal3 := make(chan (int))
	n := len(dat_numb)
	go_num := runtime.NumCPU() / 2
	//time_start := time.Now()
	//n := 2000
	Jobs := make(chan int, n)
	Results_dij := make(chan Table_int, n)
	Results_count := make(chan (int))
	Count_go := make(chan (int))
	//dat_numb := Load_data("./data", n)
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

	// Read and process data from the client
	n := 20
	tableau := Create_matrix(n)
	reader := bufio.NewReader(conn)
	for i := 0; i < n; i++ {
		message, _ := reader.ReadString('\r')
		k := strings.Split(message, " ")
		for j := 0; j < n; j++ {
			tableau[i][j], _ = strconv.Atoi(k[j])
		}
		message, _ = reader.ReadString('\n')
	}
	envoie := dijkstra_goroutine(tableau)
	// Write data back to the client
	tab := slice_to_str(envoie)

	_, err := io.WriteString(conn, fmt.Sprintf(tab))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func main() {

	//Openning TCP server connexion

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8000")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		// Handle client connection in a goroutine
		go handleClient(conn)
	}

}
