package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{}

	sn2 := struct {
		age  int
		name string
	}{}
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	// 无法通过编译
	// 结构体中的 map 是无法进行比较的
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
