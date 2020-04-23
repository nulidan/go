package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{} // map key为int 值为func（方法）
	//func参数为int ,返回值为int
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](3), m[3](4))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{} //set 下创建map格式， map[值的类型]bool{}
	mySet[1] = true         //添加索引存在
	n := 1
	if mySet[n] {
		t.Logf("%d is existing", n) //logf 能输出函数值
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	t.Log(mySet)
	delete(mySet, 1)
	t.Log(len(mySet))
	t.Log(mySet)
}
