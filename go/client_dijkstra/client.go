/*
connect
conn, err := net.Dial(tcp, portString) //portString: “IP:Port”, eg “127.0.0.1:80”

disconnect
conn.Close()
defer conn.Close()

Get yourself a reader on the connection, read some characters
reader := bufio.NewReader(conn)
message:= reader.ReadString(‘\n’)

Write content on the connection
io.WriteString(conn, fmt.Sprintf(Coucou %d\n, i))

*/

package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
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
	/*
		print("panik! ")
		tab := tab_to_str(data)
		print("panik! ")
		_, err = io.WriteString(conn, fmt.Sprintf(tab))
		print("panik! ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	*/

	// Read and process data from the server
	matrix, err := receiveMatrix(conn)
	if err != nil {
		fmt.Println("Erreur lors de la réception de la matrice :", err)
		return
	}
	fmt.Println(matrix.Data)

	/*
		tableau := make([]int, n)
		reader := bufio.NewReader(conn)
		message, _ := reader.ReadString('\r')
		k := strings.Split(message, " ")
		for j := 0; j < n; j++ {
			tableau[j], _ = strconv.Atoi(k[j])
		}
		message, _ = reader.ReadString('\n')

		fmt.Println(tableau)
	*/
}
