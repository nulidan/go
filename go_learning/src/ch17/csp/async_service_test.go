package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50) //睡眠等待
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

//串行输出

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() { //启动协程
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret //把ret数据放给retCh
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) //从retCH中取出数据
	time.Sleep(time.Second * 1)
}

//异步返回方式
