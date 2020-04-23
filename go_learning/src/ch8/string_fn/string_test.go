package string_fn

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",") //字符串分割，以逗号,为分隔符
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-")) //字符串连接，以 - 为连接符
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10) //int 转string 用 Itoa
	// 10 转换为字符串 "10"
	t.Log("str " + s)
	if i, err := strconv.Atoi("10"); err == nil {
		// string 转 int 用 Atoi ,再判断是否为空
		//"10" 转换为 int 10 赋值给 i
		t.Log(10 + i)
	}
}
