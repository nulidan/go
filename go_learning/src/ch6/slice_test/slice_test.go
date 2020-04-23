package slice_test

import (
	"math/rand"
	"testing"
	"time"
)

func TestSliceInit(t *testing.T) {
	var s0 []int            //定义切片
	t.Log(len(s0), cap(s0)) //len长度，cap容量  0 0
	s0 = append(s0, 1)      //添加元素
	t.Log(len(s0), cap(s0)) // 1 1

	s1 := []int{1, 5, 3, 7, 9} //创建有内容的切片
	t.Log(len(s1), cap(s1))    //5 5

	s2 := make([]int, 3, 5)              //创建类型为int,长度为3，容量为5的切片
	t.Log(len(s2), cap(s2))              //3 5
	t.Log(s2[0], s2[1], s2[2] /*s3[3]*/) // 默认0值，len为可访问长度 0 0 0
	s2 = append(s2, 7)
	t.Log(s2[0], s2[1], s2[2], s2[3]) // 0 0 0 7
	t.Log(len(s2), cap(s2))           // 4 5  添加一位，len增加，容量够因此不变
}

func TestSliceFor(t *testing.T) {
	s := []int{}
	oldCap := cap(s)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			t.Log(oldCap, newCap)
			oldCap = newCap
		}
		//t.Log(len(s), cap(s)) //1 1  2 2  3 4...8 8  9 16  10 16
		//容量存储空间着长度增加而乘2的扩展

	}
}

func TestSliceShareMemory(t *testing.T) {
	men := []string{"zhang", "yi", "xing", "jin", "zai", "zhong", "shen",
		"chang", "min", "zai", "yi", "qi"} //cap=12
	q1 := men[3:6]
	t.Log(q1, len(q1), cap(q1)) // jin zai zhong  3 9  len=6-3 cap=12-3
	//切片的连续共享存储结构导致容量一直到最后
	summer := men[6:9]                      //len = 9-6  cap=12-6
	t.Log(summer, len(summer), cap(summer)) // shen chang min  3 6
	q1[0] = "love"                          //修改元素，整个数据都被改变，因为且切片是连续共享存储结构
	t.Log(summer)
	t.Log(q1)
	t.Log(men)
}

func TestCopy(t *testing.T) {
	srcSlice := []int{1, 2, 3, 4, 56, 7}
	dstSlice := []int{13, 42, 34, 5, 223, 54, 564, 23}
	copy(dstSlice, srcSlice)
	t.Log(dstSlice)
}
func InitData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}

func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
func TestSlicefunc(t *testing.T) {
	n := 10
	s := make([]int, n) //创建一个切片，len为n
	InitData(s)
	t.Log(s)
	BubbleSort(s) //冒泡排序
	t.Log(s)
}
