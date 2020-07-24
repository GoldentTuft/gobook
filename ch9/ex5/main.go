package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func pingpong(out, in, res chan int, timeout <-chan time.Time) {
	count := 0
	for {
		select {
		case <-timeout:
			res <- count
			fmt.Println("*** return pingpong")
			return
		case v := <-in:
			count = v + 1
			select {
			case <-timeout:
				res <- count
				fmt.Println("*** return pingpong")
				return
			case out <- count:
			}
		}
	}
}

func main() {
	log.Printf("running %d goroutines", runtime.NumGoroutine())
	out := make(chan int)
	in := make(chan int)
	res := make(chan int, 2)
	go pingpong(out, in, res, time.After(10*time.Second))
	go pingpong(in, out, res, time.After(10*time.Second))
	log.Printf("running %d goroutines", runtime.NumGoroutine())
	start := time.Now()
	in <- 0
	resValue := <-res
	// close(res)
	elapsed := time.Now().Sub(start)
	fmt.Printf("%d count\n", resValue)
	fmt.Printf("%v elapsed time\n", elapsed)
	fmt.Printf("%f count per second\n", float64(resValue)/elapsed.Seconds())
	time.Sleep(3 * time.Second)
	log.Printf("running %d goroutines", runtime.NumGoroutine())
}
