package main

import (
	"fmt"
)

func main() {
    another:=func(done <-chan interface{}, input <-chan int) <-chan int {
  	completed := make(chan int)
  	go func() {
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
  input:=make (chan int)
  www:=another(done,input)
  produce:=func(input chan <- int ) {
  	for i:=0;i<5;i++ {
      input <- i
	}
	done<-nil
  }
  go produce(input)
  <-www



}
