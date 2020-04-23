package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	//p.Speak()             //调用内部中Speak 的方法
	fmt.Println("", host) //host指向测试程序中的 zhangyixing
}

type Dog struct {
	p *Pet //Dog数据类型内同样为p指向pet数据类型
}

func (d *Dog) Speak() {
	d.p.Speak() //调用内部成员变量方法

}

func (d *Dog) SpeakTo(host string) {
	//d.Speak()
	d.p.SpeakTo(host) //在Dog的数据类型中的 SpeakTo 方法中调用 Pet 数据类型中的 SpeakTo 方法中的host
	fmt.Println("jinzaizhong ", host)
	fmt.Println("", host)
}

type cat struct {
	Pet
}

func TestDog(t *testing.T) {
	dog := new(Dog)            //创建一个实例为Dog的dog对象
	dog.SpeakTo("")            //dog调用Dog中speakTo的方法,因为没有host值 所以输出为jinzaizhong
	dog.Speak()                //调用Dog中Speak的方法，输出为...
	dog.SpeakTo("zhangyixing") // 调用Dog中SpeakTo 的方法 host值为zhangyixing,
	//所以输出为zhangyixing  jinzaizhong  zhangyixing  zhangyixing
	cat := new(cat)
	cat.Speak()
}
