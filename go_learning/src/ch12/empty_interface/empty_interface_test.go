package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("integer", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }
	// fmt.Println("Unlnow Type")
	switch v := p.(type) {
	case int:
		fmt.Println("integet", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(77)
	DoSomething("zhangyixing")
	DoSomething(77.3)
}

//接口最佳实践
//倾向使用小接口定义，很多接口只包含一个方法
// type Reader interface {
// 	Read(p []byte) (n int. err error)
// }
//type Writer interface {
// 	Write(p []byte) (n int. err error)
// }
// 较大接口定义，可由多个小接口定义组合而成
// type ReadWriter interface {
// 	Reader
// 	Writer
// }
// 只依赖于必要功能的最小接口
// func StoreData(reader Reader) error {
// 	...
// }
