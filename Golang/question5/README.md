## 1 结构体的比较
```go
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
```

参考答案及解析：编译不通过  invalid operation: sm1 == sm2
这道题考的是结构体的比较，有几个需要注意的地方：
- 1.结构体只能比较是否相等，不能比较大小。
- 2.相同类型的结构体才能进行比较（字段的值和类型一致），结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1 就是不同的结构体。
- 3.如果 struct 的所有成员都可以比较，则该 struct 才可以通过 == 或 != 进行比较，比较时逐项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等。

**那么什么是可比较的呢？常见的有 bool、数值型、字符、指针、数组等，像切片、map、函数等是不能比较的。** 具体可以参考 Go 说明文档： https://golang.org/ref/spec#Comparison_operators。
