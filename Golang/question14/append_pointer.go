package main

/**
 * @description
 * @author chengzw
 * @since 2022/6/30
 * @link
 */
import "fmt"

func add(arr *[]int) {
	*arr = append(*arr, 3)
	*arr = append(*arr, 4)
}

func main() {
	// 第一个元素初始值为 0
	arr := make([]int, 1, 10)
	arr = append(arr, 1)
	arr = append(arr, 2)
	add(&arr) // 传递切片的指针
	// 输出 0,1,2,5,10
	// 在 add 函数中添加的 3，4 加到了原有的 arr 中，长度变成 5
	fmt.Printf("%d,%d,%d,%d,%d", arr[0], arr[1], arr[2], len(arr), cap(arr))
}
