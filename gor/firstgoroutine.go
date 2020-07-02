package main
import ("fmt"
"sync")
func main() {
	var wg sync.WaitGroup
	x:= func() {
		defer wg.Done()
		fmt.Println("I would be used as go routine")
	}
    wg.Add(2)
	go x()
	go amAnotherGoRoutine(&wg)
	fmt.Println("hello")
	wg.Wait()
}

func amAnotherGoRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("I am another go routine")
}