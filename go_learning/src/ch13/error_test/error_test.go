package error_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

//区分错误类型先预制 错误类型
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThenHundredError = errors.New("n should be not laeger than 100")

//func GetFibonacci(n int) []int {  //未进行错误检测
func GetFibonacci(n int) ([]int, error) {
	// if n < 2 || n > 100 { //错误检测
	// 	return nil, errors.New("n should be in [2, 100]")
	// }
	if n < 2 { //区分错误类型
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fibList := []int{1, 1}
	for i := 2; /*短变量声明 :=  */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	//return nil, fibList //没进行错误检测的时候的返回值
	return fibList, nil
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(str); err != nil { //对err反方向执行，如果有一个错误就输出错误，并且返回
		fmt.Println("error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)

} //只有当所有都没有错误的时候才会输出最后的结果

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil { //用if判断数列值是否错误，在错误的时候输出错误提示
		if err == LessThanTwoError { //判断错误是哪一类错误
			fmt.Println("It is less")
		}
		//t.Error(err) //调用错误提示
	} else {
		t.Log(v) //未发现错误，则调用数列函数
	}
	GetFibonacci2("10") //

}
