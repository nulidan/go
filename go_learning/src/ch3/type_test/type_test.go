package type_test

import "testing"

type MyInt int8

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64 = 10002325
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

//不支持隐式转换
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

//支持指针，但不支持指针运算

func TestString(t *testing.T) {
	var s string
	t.Log("&" + s + "&")
	t.Log(len(s))
	if s == "" {
		t.Log("hello,world")
	}
}

//String 在go语言中初始化时是空字符串
