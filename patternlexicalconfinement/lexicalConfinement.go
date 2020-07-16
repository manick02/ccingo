package main

import "fmt"

func main() {

	chanOwner:=func() <-chan int {
		channel:=make(chan int,5) // scoping production into the channel only in this go routine
		go func() {
			defer close(channel) // results
			for i:=0;i < 5;i++ {
				fmt.Println("Produced ",string(i))
             channel<-i
			}
		}()
		return channel
	}
	output:=chanOwner()
	consumer:=func(output <-chan int) { //receive only read only copy of channel
      for result:=range output{
      	fmt.Println(result)
	  }
	}
	consumer(output) //
}
