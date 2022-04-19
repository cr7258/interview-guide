package main

import "fmt"

func main() {
	list := new([]int)
	// new([]int) 之后的 list 是一个 *[]int 类型的指针，
	// 不能对指针执行 append 操作。可以使用 make() 初始化之后再用。
	list = append(list, 1)
	fmt.Println(list)
}
