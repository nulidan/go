package main

import "fmt"

type Student struct { //定义一个结构体
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main() {
	//顺序初始化，每个成员 必须初始化
	var s Student = Student{1, "金在中", "men", 34, "最喜欢昌珉了"}

	//指定成员初始化，没有初始化的成员，自动赋值为0
	s1 := Student{name: "沈昌珉", addr: "最喜欢在中了"}

	s2 := &Student{name: "魏大勋", addr: "加油"}

	var s3 *Student = &Student{1, "金在中", "men", 34, "最喜欢昌珉了"}

	fmt.Printf("type is %T %T %T %T\n", s, s1, s2, s3)

	p := Student{1, "金在中", "men", 34, "最喜欢昌珉了"}
	p.PrintfInfo()

	var p2 Student
	(&p2).SetInfo(1, "金在中", "men", 4, "最喜欢昌珉了")
	p2.PrintfInfo()
}

//通过一个函数，给成员赋值
func (p *Student) SetInfo(n int, m string, b string, v int, c string) {
	p.id = n
	p.name = m
	p.sex = b
	p.age = v
	p.addr = c
}

//带有接受者的函数叫做方法
//只要接收者类型不一样，就是不同方法
//func (接收者名字 接收者类型) 方法名（）
func (tmp Student) PrintfInfo() {
	fmt.Println("tmp = ", tmp)
}

type char byte //定义一个结构体，当作方法接收者
//type xxx 结构体类型（不能是指针）

func (tmp char) PrintfInfo() {
	fmt.Println("tmp = ", tmp)
}
