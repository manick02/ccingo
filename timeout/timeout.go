package main

import (
	"fmt"
	"time"
)

func main() {
	someCh:=make(chan int)
	someCh2:=time.After(1*time.Second)
	go func() {
		time.Sleep(1*time.Second)
		someCh<-1
	}()

	for {
		select {
		case x:=<-someCh:
			fmt.Println(x)
			return
		case y:=<-someCh2:
			fmt.Println(y)
			return
		default:
			fmt.Println("Non Activity")

		}
	}

}
