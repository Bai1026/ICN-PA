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
	f, err := os.Create("PA2-output2.txt")
	check(err)
	defer f.Close() //記得關掉

	writerr := bufio.NewWriter(f) //把東西寫進f裏面
	len, _ := writerr.WriteString("This is aaa test! \nyes\nmother")
	fmt.Println(len)
	writerr.Flush()
}
