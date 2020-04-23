package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValuses() (int, int) { //定义方法名字returnMultiValues()  返回值为(int, int )
	//多返回值函数
	return rand.Intn(10), rand.Intn(20) //Intn 随机数
}

func timeSpent(inner func(op int) int) func(op int) int { //计时函数
	//func timeSpent(参数传入一个其他的函数inner func(输入参数为一个整形）返回参数为一个整形) 返回值为一个函数func (输入参数为一个整形）返回参数为一个整形)
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
	a, b := returnMultiValuses() //或者只想输出一个 则可以写 a，_ :=
	//t.Log(a)
	t.Log(a, b)
	tsSF := timeSpent(slowFun) //调用timeSpent函数来作为参数
	t.Log(tsSF(15))
}

func Sum(ops ...int) int { //可变参数
	ret := 0
	for _, op := range ops { //数组求和
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(2, 4, 5, 8, 7))
}

//defer 函数   延迟函数

func Clear() {
	fmt.Println("这都是什么鬼啊，懵逼树下只有我")
}

func TestDefer(t *testing.T) {
	defer Clear() //延迟输出
	fmt.Println("一个刺激的开始")
	//panic("err") //panic 退出环节  但defer仍然执行，只是延迟，在panic之前
}
