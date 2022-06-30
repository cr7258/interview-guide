package main

import "fmt"

/**
 * @description
 * @author chengzw
 * @since 2022/6/30
 * @link
 */

func add(arr []int) {
	// 修改指针，并不是在修改指针指向的数据
	arr = append(arr, 3)
	arr = append(arr, 4)
	arr[1] = 100 // 直接修改指针指向的数据的内容，生效
}

func main() {
	// 第一个元素初始值为 0
	arr := make([]int, 1, 10)
	arr = append(arr, 1)
	arr = append(arr, 2)
	add(arr)
	// 输出 0,1,2,3,10
	// 在 add 函数中添加的 3，4 并没有加到原有的 arr 中，长度还是 3
	fmt.Printf("%d,%d,%d,%d,%d", arr[0], arr[1], arr[2], len(arr), cap(arr))
}
