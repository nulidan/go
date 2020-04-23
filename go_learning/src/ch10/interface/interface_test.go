package interface_test

import (
	"fmt"
	"testing"
	"time"
)

//duck type 形式
type Programmer interface { //定义接口
	//只有声明，没有实现，由其他类型实现
	WriteHelloWorld() string //接口签名
}

type GoProgrammer struct { //定义接口的实现

}

func (g *GoProgrammer) WriteHelloWorld() string { //方法中与接口签名一样

	return "fmt.Println(\"hello,world\")" + "测试一下"
}

func TestClient(t *testing.T) {
	var p Programmer //接口变量  接口类型是Programmer
	//只要实现了此接口方法的类型，这个类型的变量（接收者类型）就可以给p赋值
	p = new(GoProgrammer) //实现Programmer的实例是GoProgrammer
	t.Log(p.WriteHelloWorld())
}

//Go接口
//接口定义可以包含在接口使用者包内
//接口为非侵略性，实现不依赖接口定义

//自定义类型定义
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv { //计时函数
	//func timeSpent(参数传入一个其他的函数inner func(输入参数为一个整形）返回参数为一个整形)
	//返回值为一个函数func (输入参数为一个整形）返回参数为一个整形)
	return func(n int) int {
		start := time.Now()
		ret := inner(n)                                         //调用内部函数
		fmt.Println("time spent:", time.Since(start).Seconds()) //输出内部函数运行时间
		return ret
	}
}

func slowFun(op int) int { //测试
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun) //调用timeSpent函数来作为参数
	t.Log(tsSF(15))

}

func TestSayHi(t *testing.T) {
	x := make([]Programmer, 3)
	x[0] = 1
	x[1] = 2
	x[2] = 3

	for _, i := range x {
		t.Log(x.WriteHelloWorld())
	}
}
