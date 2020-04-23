package oper_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{4, 5.3, 6, 3} 只有相同长度维数的数组才能比较
	d := [...]int{1, 2, 3, 5}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
	t.Log(b == d)

}

//按位清零  对于左右两边，只要右边的二进制数为1，左边清零，为1的话，保持不变
//1 &^ 0 --1
//1 &^ 1 --0
//0 &^ 1 --0
//0 &^ 0 --0
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
