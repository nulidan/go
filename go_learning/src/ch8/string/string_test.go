package string_test

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值
	s = "thoth"
	t.Log(len(s))
	//s[1] = '3'  //string 是不可变的byte slice
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "张艺兴"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("张艺兴 Unicode %x", c[0]) //unicode 是字符集
	t.Logf("张艺兴 UTF8 %x", s)       //UTF8 是转花为字符序列的规则
}

func TestStringToRune(t *testing.T) {
	s := "沈昌珉爱金在中"
	for _, c := range s { //range 遍历字符串里面每一个rune的时候使用
		t.Logf("%[1]c %[1]x", c) //输出的【1】表示都是以第一个参数对应，以%c 转换，
		//以%x 转换
	}
}
