package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("hello.go") //..到上個資料夾的層級去找hello
	check(err)

	word1, word2 := "", ""
	fmt.Fscanln(f, &word1, &word2)
	fmt.Printf("%s %s\n", word1, word2)

	for i := 2; i <= 5; i++ {
		word1, word2 = "", ""
		fmt.Fscanln(f, &word1, &word2)
		fmt.Println(word1, word2)
	}

	f.Close()
}
