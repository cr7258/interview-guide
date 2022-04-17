package main

// 正确写法，在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。
func funcMui(x, y int) (sum int, err error) {
	return x + y, nil
}
