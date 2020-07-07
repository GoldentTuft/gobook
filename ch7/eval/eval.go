package eval

import (
	"fmt"
	"math"
)

// Expr は算術式
type Expr interface {
	// Eval は環境env内でのこのExprの値を返します。
	Eval(env Env) float64
}

// Var は変数を特定します。例: x
type Var string

// literal は数値定数。例: 3.141
type literal float64

// unary は単項演算子を表します。例: -x
type unary struct {
	op rune // '+'か'-'のどちらか
	x  Expr
}

// binary は二項演算子を表します。例: x+y
type binary struct {
	op   rune // '+', '-', '*', '/'のどれか
	x, y Expr
}

// call は関数呼び出し式を表します。例: sin(x)
type call struct {
	fn   string // "pow", "sin", "sqrt"のどれか
	args []Expr
}

// Env は変数名を値へと対応付ける環境です。
type Env map[Var]float64

// Eval は
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// Evalは
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
