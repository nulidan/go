package interface_test

import (
	"fmt"
	"testing"
)

type Code string //定义一个code类型，但实际还是String类型
type Programmer interface {
	WriteHelloWorld() Code //string的别名code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")" //%T为输出实例p或者说是传入类型
}

func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	//goProg := new(GoProgrammer)
	goProg := &GoProgrammer{} //interface 对应的是指针，所以这种方式需要用&取值符
	javaProg := new(JavaProgrammer)
	writeFirstProgrammer(goProg)
	writeFirstProgrammer(javaProg)
}

//空接口
