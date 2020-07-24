package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func pipeline(start chan int) chan int {
	var n, p chan int

	for i := 0; i < 1000000; i++ {
		if i == 0 {
			p = start
		} else {
			p = n
		}
		n = make(chan int)
		go func(next, prev chan int) {
			v := <-prev
			next <- v
			// for v := range prev {
			// 	next <- v
			// }
			close(next)
		}(n, p)
	}
	return n
}

func pipeTest(sendVal int) {
	start := make(chan int)
	end := pipeline(start)
	log.Printf("running %d goroutines", runtime.NumGoroutine())
	fmt.Println("press key to start sending")
	os.Stdin.Read(make([]byte, 10))
	startTime := time.Now()
	start <- sendVal
	res := <-end
	fmt.Println(time.Since(startTime))
	log.Printf("running %d goroutines", runtime.NumGoroutine())
	fmt.Println(res)
	close(start)
}

func main() {
	fmt.Println("キーを押すと1回目スタート")
	os.Stdin.Read(make([]byte, 10))
	pipeTest(666)
	fmt.Println("1回目終了")

	fmt.Println("キーを押すと2回目スタート")
	os.Stdin.Read(make([]byte, 10))
	pipeTest(777)
	fmt.Println("2回目終了")

	time.Sleep(120 * time.Second)

}
