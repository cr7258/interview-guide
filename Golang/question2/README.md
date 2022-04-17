
下面这段代码输出什么，说明原因。
```go
package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
```
输出结果：
```bash
0 -> 3
1 -> 3
2 -> 3
3 -> 3
```

参考解析：在 for range 语句中， val 变量用于保存 slice 切片所取得的值，但是 val 变量只被声明了一次，此后是将迭代 slice 切片的值赋值给 val, val 变量的内存地址始终不变。所以 m[key] = &val 取的都是变量 val 的地址，val 在最后一次迭代时的值是 3，因此所有输出都是 3。

正确的写法：
```go
package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
        // 引入一个中间变量，每次迭代都重新声明一个变量
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
```

## 参考资料
- [踩坑记录-go中使用forRange指针取值问题](https://maaaask.com/2021/12/27/%E8%B8%A9%E5%9D%91%E8%AE%B0%E5%BD%95-go%E4%B8%AD%E4%BD%BF%E7%94%A8forRange%E6%8C%87%E9%92%88%E5%8F%96%E5%80%BC%E9%97%AE%E9%A2%98/)