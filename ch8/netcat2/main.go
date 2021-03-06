package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go mustCopy(os.Stdout, conn)

	// 入力?サーバー側?が終了してしまうと、ここが失敗して、
	// メインが死ぬことによって上のゴルーチンも終了してしまう?
	mustCopy(conn, os.Stdin)
}
