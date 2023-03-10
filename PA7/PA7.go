package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/* http request command:
$ curl 127.0.0.1:12004/server-test.html
*/

func handleConnection(c net.Conn) {
	// converts the socket connection conn into an I/O buffer
	reader := bufio.NewReader(c)
	// reads from the socket, process the textual input, and store
	// the information as a object
	req, err := http.ReadRequest(reader)
	check(err)

	reqUrl := strings.Split(req.RequestURI, "/")
	fi, erri := os.Open(reqUrl[1])
	defer fi.Close()

	if erri != nil {
		// handle the error
		fmt.Printf("File not found\n")
		c.Close()
		return
	}

	fii, _ := fi.Stat()
	fmt.Printf("File size = %d\n", fii.Size())
	c.Close()
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12004")
	defer ln.Close()

	//for every HTTP request
	for {
		conn, _ := ln.Accept()
		defer conn.Close()
		go handleConnection(conn)
	}
}
