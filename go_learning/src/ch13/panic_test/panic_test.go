package panic_test

import (
	"errors"
	"fmt"
	"testing"
)

//panic 用于不可恢复的错误
//panic 退出前会执行defer指定的内容

//os.Exit 退出前不会调用defer指定的函数
//os.Exit 退出前不输出当前调用栈信息

func TestPanicVxEcit(t *testing.T) {
	defer func() { //危险的错误修复方法
		if err := recover(); err != nil { //可能会导致health check失效
			fmt.Println("recovered panic", err) //恢复不确定错误的最佳方法是 Crash 使之重启
		}
	}()
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	fmt.Println("Start")
	panic(errors.New("this is bed"))
	//os.Exit(-1)  
}
