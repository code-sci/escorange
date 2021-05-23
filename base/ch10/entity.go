package ch10

import "fmt"

//
// 创建人: huangjinmu
// 创建时间: 2021/5/22 上午9:45

type Person struct {
	Name string
	age  int
}

// 通过首字母大小写控制访问权限
type student struct {
	Name string
	age  int
}

// 结构体绑定方法
func (p Person) Print() {
	fmt.Printf("名字:%s, 年龄 %d", p.Name, p.age)
	fmt.Println()
}

func (p *Person) SetAge(age int) {
	p.age = age
}
