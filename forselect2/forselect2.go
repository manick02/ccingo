package main

import (
	"fmt"
	"time"
)

func main() {

	square:=func(done <-chan int, input <-chan int) <-chan int {
		output:=make(chan int)
		//defer close(output)
		go func() {
       // defer close(output)
		for {
			select {
			case <-done:
				fmt.Println("done is done")
				return
			case b:=<-input:
				fmt.Println("received input to square")
				result:=b*b
				output <- result
			}
		}
		}()
		return output
	}

	consumeOutput:= func(mainDone <-chan int,read <-chan int ) {
		for {
			select {
			case <-mainDone:
				fmt.Println("mainDone is done")
				return
			case out:=<-read:
				fmt.Println("Received square to print")
				fmt.Println(out)
				continue
			}
		}
	}
	done:=make(chan int)
	input:=make(chan int)
	mainDone:=make(chan int)
	//defer close(done)
	//defer close(mainDone)

	output:=square(done,input)

	go consumeOutput(mainDone,output)
	for _,i:= range [] int{1,2,3,4,5} {
		fmt.Println("Sending input:",i)
		input <-i
	}
	time.Sleep(3*time.Second)
	close(done)
    close(mainDone)


}
