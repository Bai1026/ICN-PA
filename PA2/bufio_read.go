package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("PA2-output.txt")
	check(err)

	scanner := bufio.NewScanner(f) //
	for scanner.Scan() {           //逐行scan
		fmt.Println(scanner.Text())
	}

	f.Close()
}
