package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//1st step

	// conn, errc := net.Dial("tcp", "140.112.42.221:11999") //station IP:port -> polly's address

	// conn, errc := net.Dial("tcp", "140.112.42.221:12004") //station IP:port

	conn, errc := net.Dial("tcp", "127.0.0.1:12004") //station IP:port -> we use in PA6.go
	check(errc)
	defer conn.Close()

	//2nd step
	input := ""
	fmt.Print("Input filename: ")
	fmt.Scanf("%s", &input) //read string -> input

	// 3rd step
	fi,erri:=os.Open(input)
	defer fi.Close()
	check(erri)
	fii,errf:=fi.Stat()
	check(errf)
	t:=strconv.Itoa(int(fii.Size()))
	fmt.Println("Send the file size first:"+t)

	//4th step
	writer:=bufio.NewWriter(conn)
	scannerf := bufio.NewScanner(fi)
	content:=t+"\n"
	for scannerf.Scan(){
		content+=scannerf.Text()+"\n"
	}
	writer.WriteString(content)
	writer.Flush()

	// 5th, 6th step
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("Server says: %s\n", scanner.Text())
	}
}
