package main

import (
	"bufio"
	"fmt"

	// "io"
	"net"
	// "os"
	// "strconv"
	// "strings"
)
func check(e error) {
 if e != nil {
 panic(e)
 }
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":11999")
	conn, _ := ln.Accept()
	defer ln.Close()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, errr := reader.ReadString('\n')
	check(errr)
	fmt.Printf("%s", message)
	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received\n", len(message))
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()
}