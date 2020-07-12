package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup // 活動中のゴルーチンの数
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go func() {
			go echo(c, input.Text(), 1*time.Second)
			wg.Done()
		}()
	}
	wg.Wait()
	// 注意: input.Err()からの潜在的エラーを無視している
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
