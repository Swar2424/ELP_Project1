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
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
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
	n := 20
	data := Load_data("./data", n)
	tab := tab_to_str(data)
	_, err = io.WriteString(conn, fmt.Sprintf(tab))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read and process data from the server
	tableau := make([]int, n)
	reader := bufio.NewReader(conn)
	message, _ := reader.ReadString('\r')
	k := strings.Split(message, " ")
	for j := 0; j < n; j++ {
		tableau[j], _ = strconv.Atoi(k[j])
	}
	message, _ = reader.ReadString('\n')

	fmt.Println(tableau)

}
