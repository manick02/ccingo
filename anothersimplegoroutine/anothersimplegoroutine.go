package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _,i:= range []int {1,3,4,5} {
		wg.Add(1)
		go func(i int) {
            defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	fmt.Println("Waiting for go routines to finish")
	wg.Wait()
	fmt.Println("Go routines are complete")
}
