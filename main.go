package main

import "fmt"

type Message struct {
	Hello string
}

func main() {
	h := Message{Hello: "Girl"}

	fmt.Printf("%s \n", h)
}
