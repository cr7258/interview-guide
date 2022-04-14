
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

参考解析：
for range 循环的时候会**创建每个元素的副本，而不是元素的引用**，所以 m[key] = &val 取的都是变量 val 的地址，val 在最后一次迭代时的值是 3，因此所有输出都是 3。

正确的写法：
```go
package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
        // 在赋值前用另一个变量替换
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
```