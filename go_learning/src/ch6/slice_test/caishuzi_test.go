package slice_test

import (
	"math/rand"
	"testing"
	"time"
)

func GetNum(p *int) {
	var num int
	rand.Seed(time.Now().UnixNano())
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}
	*p = num
}

func TestNumGuess(t *testing.T) {
	var Num int
	GetNum(&Num)
	t.Log(Num)

}
