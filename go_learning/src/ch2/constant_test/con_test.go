package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1 //0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

//变量定义 a:=1
//连续常量定义 const(
// 	Readable = 1 << iota
// 	Writable
// 	Executable
// )     多行注释按CTRL+k，c
