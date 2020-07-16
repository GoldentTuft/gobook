package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	Out  chan<- string
	Name string
}

// type client chan<- string // 送信用メッセージチャネル

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // クライアントから受信するすべてのメッセージ
)

func broadcaster() {
	clients := make(map[client]bool) // すべての接続されているクライアント
	for {
		select {
		case msg := <-messages:
			// 受信するメッセージをすべてのクライアントの
			// 送信用メッセージチャネルへブロードキャストする。
			for cli := range clients {
				cli.Out <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.Out <- "*** online clients ***"
			for c := range clients {
				cli.Out <- c.Name
			}
			cli.Out <- "**********************"

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // 送信用のクライアントメッセージ
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{ch, who}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	timeout := 5 * time.Minute
	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// 注意: input.Err()からの潜在的なエラーを無視している

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意: ネットワークのエラーを無視している
	}
}
