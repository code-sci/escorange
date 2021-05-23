package main

import (
	"escorange/base/ch10"
	"fmt"
	"unsafe"
)

// struct的学习
// 创建人: huangjinmu
// 创建时间: 2021/5/22 上午9:01
func main() {
	// 第一种实例化对象
	var c1 = Course{
		name:  "golang",
		price: 12,
		url:   "http://www.baidu.com",
	}
	fmt.Println(c1)

	// 第二种实例化对象
	c2 := Course{"python", 11, "555"}
	fmt.Println(c2)

	// 对象访问权限, 不在同一个包下, 首字母小写的无法访问
	var p = ch10.Person{
		Name: "zhangsan",
	}
	fmt.Println(p)

	// 指向结构体的指针
	c3 := &c2
	//fmt.Printf("%T", c3)
	// 从指针对象获取数据
	fmt.Println((*c3).name, (*c3).url, (*c3).price)
	// go语言的语法糖, go语言会自动转换,c3就优化成了(*c3)
	fmt.Println(c3.name, c3.url, c3.price)

	// 零值
	c4 := Course{}
	var c5 Course
	var c6 *Course = new(Course)
	// 指针如果只声明,而不赋值,默认值是nil, 可以采用 &Course{}来初始化
	var c7 *Course
	var c8 *Course = &Course{}

	fmt.Println(c4)
	fmt.Println(c5)
	fmt.Println(c6)
	fmt.Println(c7)
	fmt.Println(c8)

	// 结构体是值传递
	c9 := Course{"python", 11, "555"}
	c10 := c9

	c10.price = 22
	fmt.Println(c9)
	fmt.Println(c10)
	// 查看占用内存
	fmt.Println(unsafe.Sizeof(c10))
	// string 本质是结构体
	fmt.Println(unsafe.Sizeof("射流风机啊乐山大佛就"))
	p.Print()
	p.SetAge(777)
	p.Print()
}

type Course struct {
	name  string
	price int
	url   string
}
