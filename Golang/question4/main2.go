package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	// append() 的第二个参数不能直接使用 slice，需使用 … 操作符，
	// 将一个切片追加到另一个切片上：append(s1,s2…)
	s1 = append(s1, s2)
	fmt.Println(s1)
}
