package series

import "fmt"

func init() {
	fmt.Println("this is init1")
}
func init() {
	fmt.Println("this is init2")
}

func Squars(n int) int { //开头大写字母的函数能被外部访问，开头小写不行
	return n * n
}
func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
