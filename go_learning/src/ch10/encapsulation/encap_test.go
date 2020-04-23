package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct { //数据封装  关键词struct
	Id   string
	Name string
	Age  int
}

type Em struct {
	Employee
	sex  byte
	addr string
}

func TestCreatEmployeeValue(t *testing.T) {

	s1 := Em{Employee{"123", "jinzaizhong", 34}, 'm', "最喜欢昌珉了"}
	s1.Name = "nier"
	s1.sex = 'm'
	s1.Age = 18
	s1.Id = "666"
	s1.addr = "加油啊，不要死"
	t.Log(s1.Name, s1.sex, s1.Age, s1.Id, s1.addr)

}
func TestCreatEmployeeEm(t *testing.T) {
	//顺序初始化
	var s1 Em = Em{Employee{"123", "jinzaizhong", 34}, 'm', "最喜欢昌珉了"}
	t.Log(s1)
	//自动推到类型
	s2 := Em{Employee{"666", "jinzaizhong", 34}, 'm', "最喜欢昌珉了"}
	t.Log(s2)
	//指定成员初始化，没有初始化的常量自动赋值为0
	s3 := Em{sex: 'm'}
	t.Log(s3)
	s4 := Em{Employee: Employee{Name: "沈昌珉"}, addr: "最喜欢金在中了"}
	t.Log(s4)
}

//实例的创建及初始化
func TestCreatEmployeeObj(t *testing.T) {
	e := Employee{"0", "thoth", 20}
	e.Id = "jinzaizhong"
	e1 := Employee{Name: "zhangyixing", Age: 27}
	e2 := new(Employee)
	e2.Id = "49"
	e2.Name = "thhranduil"
	e2.Age = 20
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Log(e2.Age)
	t.Logf("e is %T", &e) //%T是类型，输出实例的类型
	t.Logf("e is %T", e)
	t.Logf("e1 is %T", &e1) //输出employee的指针类型
	t.Logf("e1 is %T", e1)  //输出employee的类型
	t.Logf("e2 is %T", &e2)
	t.Logf("e2 is %T", e2)
}

//行为（方法）的定义
//func (e Employee) String() string {
//fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
//有对象复制产生，结构体实例的数据被复制，整个结构被复制，有空间上复制的开销
//return fmt.Sprintf("ID:%s--Name:%s--Age:%d", e.Id, e.Name, e.Age)
//}

func (e *Employee) String() string { //指向行为（方法）的指针
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	//没有对象复制产生，实例的数据存放在同一个位置，推荐
	return fmt.Sprintf("ID:%s--Name:%s--Age:%d", e.Id, e.Name, e.Age)

}

func TestStructOperation(t *testing.T) {
	//e := &Employee{"shenchangmin", "jinzaizhong", 20}
	e := Employee{"shenchangmin", "jinzaizhong", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String()) //调用String方法
	t.Log(e)
}
