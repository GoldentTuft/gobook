package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	go hoge()
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
	}
	log.Printf("running %d goroutines", runtime.NumGoroutine())
	panic("")
}

func hoge() {
	go piyo()
	for i := 0; i < 10; i++ {
		fmt.Println("hoge")
		time.Sleep(500 * time.Millisecond)
	}
}

func piyo() {
	for {
		fmt.Println("piyo")
		time.Sleep(500 * time.Millisecond)
	}
}
