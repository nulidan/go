package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a { //if 变量赋值；判断条件（表达式为布尔值）{}
		t.Log("1==1")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {

		switch i {
		case 0, 1:
			t.Log("Even")
		case 2, 3:
			t.Log("Odd")
		default:
			t.Log("it is a unknow")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {

		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("it is a unknow")
		}
	}
}

//switch 可以在case判断条件
