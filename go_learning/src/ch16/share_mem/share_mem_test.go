package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) { //非线程安全
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) { //线程安全
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second) //时间等待，等待所有协程运行完毕
	t.Logf("counter = %d", counter)
}

//同步各个线程 WaitGroup

func TestCounterWaitGroup(t *testing.T) { //线程安全
	var mut sync.Mutex
	var wa sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wa.Add(1) //增加等待的量
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wa.Done() //等待的任务结束时候告诉WaitGroup任务结束
		}()
	}
	wa.Wait() //所有的要等待的任务完成才能继续向下运行
	t.Logf("counter = %d", counter)
}
