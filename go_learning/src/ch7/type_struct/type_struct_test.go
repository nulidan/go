package type_struct

import (
	"fmt"
	"testing"
)

type Student struct { //定义一个结构体
	id   int
	name string
	sex  string
	age  int
	addr string
}

func TestTypeStruct(t *testing.T) {
	//顺序初始化，每个成员 必须初始化
	var s Student = Student{1, "金在中", "men", 34, "最喜欢昌珉了"}
	t.Log(s)

	//指定成员初始化，没有初始化的成员，自动赋值为0
	s1 := Student{name: "沈昌珉", addr: "最喜欢在中了"}
	t.Log(s1)

	s2 := &Student{name: "魏大勋", addr: "加油"}
	t.Log(s2)
	var s3 *Student = &Student{1, "金在中", "men", 34, "最喜欢昌珉了"}
	t.Log(s3)
	fmt.Printf("type is %T %T %T %T\n", s, s1, s2, s3)
}

func TestUseTypeStruct(t *testing.T) {
	//定义结构体普通变量
	var s4 Student
	//操作成员，用.
	s4.id = 666
	s4.name = "魏大勋"
	s4.age = 18
	s4.sex = "men"
	s4.addr = "善良"
	t.Log(s4)
}

func TestPointerTypeStruct(t *testing.T) {
	//定义一个普通结构体变量
	var s5 Student
	//定义一个指针变量
	var p1 *Student
	p1 = &s5

	//通过指针操作成员，p1.name = (*p1).name
	p1.id = 666
	p1.name = "魏大勋"
	p1.age = 18
	p1.sex = "men"
	p1.addr = "善良"
	t.Log(p1)

	//通过new申请一个结构体
	p2 := new(Student)
	p2.id = 666
	p2.name = "魏大勋"
	p2.age = 18
	p2.sex = "men"
	p2.addr = "善良"
	t.Log(p2)

}

func TestCompareType(t *testing.T) {

	//相同类型比较
	s := Student{1, "金在中", "men", 34, "最喜欢昌珉了"}
	s2 := Student{name: "魏大勋", addr: "加油"}
	if s == s2 {
		t.Log("yes")
	} else {
		t.Log("no")
	}
	var s6 Student
	s6 = s
	if s6 == s {
		t.Log("yes")
	} else {
		t.Log("no")
	}

}

func test01(s Student) {
	s.id = 777
	fmt.Println(s)
}
func test02(s *Student) {
	s.id = 888
	fmt.Println(s)
}

func TestTransmitType(t *testing.T) {
	s := Student{1, "金在中", "men", 34, "最喜欢昌珉了"}
	test01(s) //值传递。形参无法改变实参
	t.Log(s)
	test02(&s) //地址传递（引用传递），形参可以改变实参
	t.Log(s)
}

func (tmp Student) PrintfInfo() {
	fmt.Println("tmp = ", tmp)
}

func PrintfInfo(t *testing.T) {
	p := Student{1, "金在中", "men", 34, "最喜欢昌珉了"}
	t.Log(p.PrintfInfo())
}
