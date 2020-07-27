package main

import "fmt"

func main() {

	ch:=make(chan interface{})
	go func() {
		defer close(ch)
		fmt.Println("Hello i am good");

	}()
	fmt.Println("Hello how are you")
	<-ch
}
