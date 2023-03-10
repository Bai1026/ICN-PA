package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
 if e != nil {
 panic(e)
 }
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12004")
	conn, _ := ln.Accept()
	fo,_:=os.Create("whatever.txt")
	defer ln.Close()
	defer conn.Close()
	defer fo.Close()

	// read in file size
	reader := bufio.NewReader(conn)
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

		// message+=strconv.Itoa(pre) + line + "\n"	
		//write to ouput.txt
		newline := strconv.Itoa(pre)  + " " + line
		newfile_writer.WriteString(newline)
		newfile_writer.Flush()

		//store new me
		message += newline //to calculate the message size
		pre++ // to calculate the line number
		count += len((line)) //to calculate the current message size
	}
	outputSize := len(message) 
	fmt.Printf("Output file size: %d\n" , outputSize)

	//send back to client
	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received, %d bytes file generated",filesize ,outputSize)
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()
}
