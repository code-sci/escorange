package main

import "fmt"

// go语言的继承
// 创建人: huangjinmu
// 创建时间: 2021/5/23 上午7:45
func main() {
	t := Teacher{"张三", 23}
	c := Course{t, "python", 199}

	println(c.Name, c.Teacher.Name)
	c.ToString()
	c.Teacher.ToString()
}

type Course struct {
	Teacher
	Name  string
	Price int
}
type Teacher struct {
	Name string
	Age  int
}

func (t Teacher) ToString() {
	fmt.Printf("名字:%s, 年龄:%d", t.Name, t.Age)
	println()
}

func (c Course) ToString() {
	fmt.Printf("名字:%s, 价格:%d", c.Name, c.Price)
	println()
}
