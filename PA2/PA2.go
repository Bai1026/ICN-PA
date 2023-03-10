package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, output := "", ""
	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &input)
	fmt.Printf("Output filename: ")
	fmt.Scanf("%s", &output)

	fi, erri := os.Open(input)
	check(erri)
	defer fi.Close() //final, close

	fo, erro := os.Create(output)
	check(erro)
	defer fo.Close() //final, close

	scanner := bufio.NewScanner(fi)
	writer := bufio.NewWriter(fo)

	pre := 1 //算行數的
	for scanner.Scan() {
		writer.WriteString(strconv.Itoa(pre) + " " + scanner.Text() + "\n") //int 行數轉成string 然後加上我們讀到的東東
		writer.Flush()                                                      //記得清掉
		pre++
	}

}
