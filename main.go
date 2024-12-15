package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("What is your name? >")
	file, _ := os.OpenFile("file.txt", os.O_RDONLY, 0666)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {

		readFile, err := reader.ReadString('\n')
		fmt.Printf("-> %s", readFile)
		if err != nil {
			return

		}
	}

}
