package main

import "fmt"

// 结构体的标签
// 创建人: huangjinmu
// 创建时间: 2021/5/23 下午1:56
func main() {
	b := Book{"", 11}
	fmt.Println(b)
}

type Book struct {
	Name  string `json:"name, omitempty"` //omitempty空值不显示
	Price int
}
