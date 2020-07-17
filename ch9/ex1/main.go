// https://github.com/ray-g/gopl/blob/master/ch09/ex9.01/bank.go
// 写経
package main

import (
	"fmt"
)

type withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int) // 入金額を送信する
var balances = make(chan int) // 残高を受信する
var withdrawals = make(chan withdrawal)

// Deposit は預金
func Deposit(amount int) { deposits <- amount }

// Balance は残高
func Balance() int { return <-balances }

// Withdraw は引き出し
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- withdrawal{amount, ch}
	return <-ch
}

func teller() {
	var balance int // balanceはtellerゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawals:
			if w.amount > balance {
				w.success <- false
				continue
			}
			balance -= w.amount
			w.success <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // モニターゴルーチンを開始する
}

func main() {

	fmt.Println(Withdraw(100))
	Deposit(100)
	fmt.Println(Withdraw(100))
	fmt.Println(Balance())

}
