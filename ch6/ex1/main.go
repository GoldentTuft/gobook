package main

import (
	"bytes"
	"fmt"
)

// IntSet は負でない小さな整数のセットです。
// そのゼロ値は空セットを表しています。
type IntSet struct {
	words []uint64
}

// pc[i] はiのポピュレーションカウントです。
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popCount はxのポピュレーションカウント(1が設定されているビット数)を返します。
func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Has は負ではない値xをセットが含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add はセットに負ではない値xを追加します。
// レシーバパラメータがポインタ型でないと、コピーされたものが変更されるだけで、
// 呼び出し元のものは変更されないみたい。
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith は、sとtの和集合をsに設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String は"{1 2 3}"の形式の文字列としてセットを返します。
// レシーバパラメータが*IntSetとポインタ型なのは一貫性のためらしい。Hasも。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')

				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len は要素数を返します
func (s *IntSet) Len() int {
	sum := 0
	for _, word := range s.words {
		sum += popCount(word)
	}
	return sum
}

// Remove は値xを削除します
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] = s.words[word] &^ (1 << bit)
	}
}

// Clear はセットからすべての要素を取り除きます
func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

// Copy はセットのコピーを返します
func (s *IntSet) Copy() *IntSet {
	var res IntSet
	res.words = make([]uint64, len(s.words))
	copy(res.words, s.words)
	return &res
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	fmt.Println(x.Len())

	x.Remove(400)
	fmt.Println(x.String())
	x.Add(400)
	fmt.Println(x.String())
	x.Remove(400)
	fmt.Println(x.String())

	x.Clear()
	fmt.Println(x.String())
	x.Add(123456789)
	x.Add(300)
	fmt.Println(x.String())

	fmt.Println()
	y := x.Copy()
	y.Add(80)
	x.Add(50)
	fmt.Println(y.String())
	fmt.Println(x.String())

}
