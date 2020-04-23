package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 78: 2, 3: 3}
	m2 := map[string]int{"zhang": 1, "yi": 2, "xing": 3}
	m3 := map[string]string{"zhang": "jin", "yi": "zai", "xing": "zhong"}
	m4 := map[int]int{}
	m5 := make(map[string]int, 7)
	// map[key的类型]数值的类型{ key : 数值}
	t.Log(m1)
	t.Log(m1[78])
	t.Logf("len m1 = %d", len(m1))
	t.Log(m2)
	t.Log(m2["yi"])
	t.Logf("len m2 = %d", len(m2))
	t.Log(m3)
	t.Log(m3["yi"])
	t.Logf("len m3 = %d", len(m3))
	m4[6] = 25
	t.Log(m4)
	t.Logf("len m4 = %d", len(m4))
	t.Log(m5)
	t.Logf("len m5 = %d", len(m5))

}

func TestAccessNotExistingKey(t *testing.T) {
	//m1 := map[string]int{}
	//key,value
	// t.Log(m1["shenchangmin"])
	// m1["jinzaizhong"] = 520
	// t.Log(m1["jinzaizhong"])
	// m1["zhangyixing"] = 666
	// t.Log(m1["zhangyixing"])
	// if v, ok := m1["zhangyixing"]; ok { // 如果ok为true，则证明mi["zhangyixing"]存在
	// 	//则返回key所对应的v
	// 	t.Log("key is value is ", v)
	// } else {
	// 	t.Log("key is not existing.")
	// }
	m1 := map[int]string{1: "shenchangmin", 2: "jinzaizhong", 3: "zhangyixing"}
	v, ok := m1[2]
	if ok { // 如果ok为true，则证明mi[0]存在
		//则返回key所对应的v
		t.Log("key is value is ", v)
	} else {
		t.Log("key is not existing.")
	}
	m2 := map[int]int{}
	m2[3] = 0
	if v, ok := m2[3]; ok {
		t.Log("this is ", v)
	} else {
		t.Log("this is existing.")
	} // map 元素的访问，当访问的key 不存在的时候，任返回 0 ，所以需要判断是否里面的元素是否为0
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 3, 3: 4}
	for k, v := range m1 { //切片的for k 为key，而不是数组的index
		t.Log(k, v)
	}
}

func test(m2 map[int]string) {
	delete(m2, 666)

}
func TestDeleteMap(t *testing.T) {
	m2 := map[int]string{1: "jinzaizhong", 520: "shenchangmin", 666: "lost"}
	delete(m2, 666) //删除key为1的内容
	t.Log(m2)

	test(m2)
	t.Log(m2)
}
