package main

import (
	"fmt"
	"time"
)

func main() {
	//done := make(chan string)
	//stringStream := make(chan string)
	//for _, s := range []string{"a", "b", "c"} {
	//	select {
	//	case <-done:
	//		return
	//	case stringStream <- s:
	//
	//	}
	//}
	//close(done)
	worker:=func(done <-chan interface{},input <-chan string) <-chan string {
		completed := make(chan string)
        go func() {
         defer fmt.Println("Work completed")
         defer close(completed)
         for {
			 select {
              case s:= <-input:
              	fmt.Println(s)
              	case <-done:
					return
			}
		 }
		}()
		return completed
	}
	done:=make(chan interface{})
	input:=make(chan string)
	defer close(input)
	workerCompleted:=worker(done,input)

	go func(){
		defer close(done)
		input<-"S"
		input<-"M"
		input<-"A"
		time.Sleep(3*time.Second)
		fmt.Println("Work is done")

	}()
    <-workerCompleted
	fmt.Println("All experiments done.")

}
