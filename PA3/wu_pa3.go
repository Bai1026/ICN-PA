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
	conn, errc := net.Dial("tcp", "140.112.42.221:12000") //station IP:port
	check(errc)
	defer conn.Close()

	//2nd step
	input := ""
	fmt.Print("Input filename: ")
	fmt.Scanf("%s", &input) //read string -> input

	// 3rd step
	fi, erri := os.Open(input)
	check(erri)
	defer fi.Close()

	fii, errf := fi.Stat()
	check(errf)
	size := strconv.Itoa(int(fii.Size()))
	fmt.Println("Send the file size first:" + size)

	//4th step

	scannerf := bufio.NewScanner(fi) //read input fil e
	size_content := size + "\n"      //因為會先讀到size -> 再往下讀content
	for scannerf.Scan() {
		size_content += scannerf.Text() + "\n"
	}

	writer := bufio.NewWriter(conn)
	/*writer connect to the conn(dial to server的連線)
	server 會scan到我們放在writer.WriteString()的東西*/
	writer.WriteString(size_content)
	writer.Flush()

	// 5th, 6th step
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("Server says: %s\n", scanner.Text())
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"os"
// 	"strconv"
// )

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func main() {
// 	//1st step
// 	conn, errc := net.Dial("tcp", "140.112.42.221:12000") //station IP:port
// 	check(errc)
// 	defer conn.Close()

// 	//2nd step
// 	input := ""
// 	fmt.Print("Input filename: ")
// 	fmt.Scanf("%s", &input) //read string -> input

// 	// 3rd step
// 	fi, erri := os.Open(input)
// 	defer fi.Close()
// 	check(erri)

// 	fii, errf := fi.Stat()
// 	check(errf)
// 	t := int(fii.Size())
// 	if erri == nil {
// 		fmt.Println("Send the file size first:", t)
// 	}

// 	//4th step
// 	writer := bufio.NewWriter(conn)
// 	writer.WriteString(strconv.Itoa(t))
// 	writer.Flush()

// 	scannerf := bufio.NewScanner(fi)
// 	content := ""
// 	for scannerf.Scan() {
// 		content += scannerf.Text() + "\n"
// 	}
// 	fmt.Println("content =", content)
// 	writer.WriteString(content)
// 	writer.Flush()

// 	//5th, 6th step
// 	scanner := bufio.NewScanner(conn)
// 	if scanner.Scan() {
// 		fmt.Printf("Server says: %s\n", scanner.Text())
// 	}
// }
