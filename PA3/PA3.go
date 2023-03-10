package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//1st step
	conn, errc := net.Dial("tcp", "140.112.42.221:12000") //station IP:port
	// conn, errc := net.Dial("tcp", "140.112.42.221:11999") //station IP:port
	check(errc)
	defer conn.Close()

	//2nd step
	input := ""
	fmt.Print("Input filename: ")
	fmt.Scanf("%s", &input) //read string -> input

	// fi, erri := os.Open(input) //file_in and error_in
	// check(erri)
	// defer fi.Close()

	//3rd step
	// fi, erri := os.Stat(input) //獲取文件消息
	// check(erri)
	// if erri == nil{
	// 	fmt.Println(fi.Size())
	// }
	fi, erri := ioutil.ReadFile(input) //讀檔案
	check(erri)
	if erri == nil {
		fmt.Println(len(fi))
	}

	//4th step
	fmt.Println(string(fi))

	//5th, 6th step
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("Server replies: %s\n", scanner.Text())
	}
	// conn.Close()
}
