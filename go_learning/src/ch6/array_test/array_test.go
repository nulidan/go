package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int //定义数组，初始化默认为0
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 5, 6, 98, 7}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
	arr1[1] = 58
	t.Log(arr1)
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 5, 6, 98, 7}
	for idx, e := range arr2 { //for 索引，值 := range 数组名
		//range 用于for循环中迭代数组，切片等元素
		t.Log(idx, e)
	}
	for _, e := range arr2 { //当for条件前面不想写用 _ 代替
		t.Log(e)
	}
}

//二维数组

//截取 a[开始索引（包含），结束索引（不包含）]
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 6, 4, 6, 7, 3, 5}
	arr3_sec := arr3[3:]           //从第四个元素往后开始截取  6 7 3 5
	arr3_sec1 := arr3[:3]          //开始到第四个元素截取   1 6 4
	arr3_sec2 := arr3[3:5]         //从第四个元素到第六个元素之间截取  6 7
	arr3_sec3 := arr3[2:len(arr3)] //从第三个元素到数组最后截取
	t.Log(arr3_sec)
	t.Log(arr3_sec1)
	t.Log(arr3_sec2)
	t.Log(arr3_sec3)
}
