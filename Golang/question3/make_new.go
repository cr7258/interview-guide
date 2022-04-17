package main

import "fmt"

type struct1 struct {
	name string
	age  int
}

func main() {
	// 使用 make 创建 map
	s1 := make(map[string]int)
	fmt.Println(s1) // map[]
	s1["dog"] = 1
	s1["cat"] = 2
	fmt.Println(s1) // map[cat:2 dog:1]

	// 使用 make 创建一个长度为 3，容量为 6 的切片
	s2 := make([]int, 3, 6)
	fmt.Println(s2) // []
	s2[0] = 1
	s2[1] = 2
	s2[2] = 3
	fmt.Println(s2) // [1 2 3]

	// 使用 new 创建一个对象
	s3 := new(struct1) // 返回的是指针 *Type
	fmt.Println(s3)

	// 使用 new 创建一个 map
	s4 := new(map[string]int)
	fmt.Println(s4)        // 返回的是指针 *Type，&map[]
	*s4 = map[string]int{} // 初始化 map
	(*s4)["dog"] = 1
	(*s4)["cat"] = 2
	fmt.Println(s4) // &map[cat:2 dog:1]
}
