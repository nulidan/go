package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	fmt.Println("hello,world")
// 	var a int = 10
// 	fmt.Println("a=", a)
// 	var b, c int
// 	fmt.Println("b c=", b, c)
// 	d := 30 //:= 自动推导类型 只能用一次，用于初次化
// 	fmt.Println("d=", d)
// 	d = 10
// 	fmt.Println("d=", d)
// 	//多重赋值
// 	s, n, k := 10, 20, 30
// 	fmt.Println("s,n,k=", s, n, k)
// 	//交换俩个变量值
// 	i, j := 10, 20

// 	i, j = j, i
// 	fmt.Printf("i=%d,j=%d\n", i, j)

// 	//匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用。才有优势
// 	i, _ = j, i
// 	fmt.Printf("i=%d,j=%d\n", i, j)

// 	//常量const
// 	//变量var
// 	const dw = 10
// 	fmt.Println("dw=", dw)

// 	//不同类型变量定义/声明
// 	var (
// 		q int     = 1
// 		e float32 = 11.2
// 	)
// 	q, e = 10, 3.14
// 	fmt.Println("q=", q)
// 	fmt.Print("e=", e)

// 	//枚举 iota
// 	//常量自动生成器，每个一行，自动累加1
// 	//iota给常量赋值使用
// 	const (
// 		r = iota
// 		t = iota
// 		y = iota
// 	)
// 	fmt.Println("r=,t=,y=", r, t, y)

// 	//iota遇到const，重置为0
// 	const u = iota
// 	fmt.Println(u)
// 	//可以只写一个iota
// 	const (
// 		o = iota
// 		l
// 		k1
// 	)
// 	fmt.Printf("o=%d,p=%d,f=%d", o, l, k1)
// 	//如果同一行，值为相同,
// 	const (
// 		q1         = iota
// 		l1, l2, l3 = iota, iota, iota
// 		k2         = iota
// 	)
// 	fmt.Println(q1, l1, l2, l3, k2)

// }
// func TestFn(t *testing.T) {
// 	fmt.Println("hello,world")
// }

func GetNum(p *int) { //产生随机数
	var num int
	rand.Seed(time.Now().UnixNano()) //根据时间来随机数
	for {
		num = rand.Intn(10000) //规定4位数
		if num >= 1000 {       //控制条件为4位数
			break
		}
	}

	*p = num //将取到的随机数赋值给p指针 p指针将数据返回到GetNum
}
func CreatNum(s []int, num int) { //保存每位数的函数 randSlice随机数切片 randNum 随机数
	s[0] = num / 1000       //取商 取千位
	s[1] = num % 1000 / 100 //取百位
	s[2] = num % 100 / 10   //取十位
	s[3] = num % 10         //取个位
}
func oneGame(randSlice []int) {
	var num int
	keySlice := make([]int, 4)
	for {
		for {
			fmt.Println("请输入一个4位数")
			fmt.Scan(&num)
			if num < 10000 && num > 999 {
				break
			}

			fmt.Println("输入数字不符合要求")

		}
		//fmt.Println("num:", num)
		CreatNum(keySlice, num) //将输入的数的数据存入keySlice中
		//fmt.Println("keySlice:", keySlice)
		n := 0
		for i := 0; i < 4; i++ { //0-4之间的位数大小
			if keySlice[i] > randSlice[i] { //循环对比输入的数和随机产生的数之间的大小
				fmt.Printf("第%d位大了一点\n", i+1) //大了，进行下一位的比较

			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了一点\n", i+1)
			} else {
				fmt.Printf("第%d猜对了\n", i+1)
				n++ //在猜对位数之后，n+1，用于记录猜对的位数
			}
		}
		if n == 4 { //在比较数字大小的循环结束之后，判断位数全部猜对，跳出循环
			fmt.Println("全部猜对了")
			break
		}
	}
}
func main() {
	var randNum int
	//产生一个4位随机数
	GetNum(&randNum)
	//fmt.Println("num:", randNum)
	randSlice := make([]int, 4) //创建一个randSlice切片

	CreatNum(randSlice, randNum) //保存这个4位数的每一位
	//fmt.Println("randSlice:", randSlice)

	oneGame(randSlice)
}
