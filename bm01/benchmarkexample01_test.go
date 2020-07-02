package bm01
import ("testing"
        "fmt")

func Benchmarksample01_test(b *testing.B){
	receive := make(chan int)
	for i:=0; i < b.N; i++ {
		fmt.Println(i)
		fmt.Println(<-receive)
	}

	for j:=0; j<b.N;j++ {
		receive<-j
	}
	b.StartTimer()
}

func Benchmarksample02_test(b *testing.B){
	for i:=0; i <=b.N; i++ {
		calculate(i);
	}
}

func calculate(j int) int {
	 sum := 0;
	for i:=0; i<=j; i++ {
		sum = sum + i
	}
	return sum
}
