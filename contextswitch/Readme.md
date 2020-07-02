In order to prove the effectiveness of goroutine its key to understand and measure how to measure the context
switch
This would help us understand and appreciate what advantage go brings to table over thread context switching
Operating system provides following boundaries of concurrency
 -> Process
   -> Thread

Context switching at a layer closer to hardware is costly

go test -bench=. -cpu=2