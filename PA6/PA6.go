package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)
func check(e error) {
 if e != nil {
 panic(e)
 }
}

func handleConnection (c net.Conn) {
	fo,_:=os.Create("whatever.txt")
	defer fo.Close()

	reader := bufio.NewReader(c)
	filesize_str, errr := reader.ReadString('\n')
	check(errr)
	// transform filesize type

	filesize_str = strings.TrimSuffix(filesize_str, "\n")
	filesize , _ := strconv.Atoi(filesize_str)
	fmt.Printf("Upload file size: %d\n", filesize)
	

	message := ""
	pre:=1
	count:=0
	newfile_writer:=bufio.NewWriter(fo) 

	// read in bytes
	for count < filesize {
		line, err := reader.ReadString('\n')
		//handle error
		if err == io.EOF {
			fmt.Printf("n: %v, err: %v\n", line, err)
			break
		}

		message+=strconv.Itoa(pre) + line + "\n"	
		//write to ouput.txt
		newline := strconv.Itoa(pre)  + " " + line
		newfile_writer.WriteString(newline)
		newfile_writer.Flush()

		pre++
		count += len((line))
	}

	outputSize := len(message) 
	fmt.Printf("Output file size: %d\n" , outputSize)
	
	// //send back to client
	writer := bufio.NewWriter(c)
	newline := fmt.Sprintf("%d bytes received, %d bytes file generated",filesize ,outputSize)
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()

	time.Sleep(5*time.Second) //wait for checking if there are something else
	c.Close() //wait for 5s then close
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12004")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		// defer conn.Close(), since already close at the handleConnection

		fmt.Printf("connect!\n")
		go handleConnection(conn)
	}
}
