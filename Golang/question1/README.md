下面代码的输出内容是什么？

```go
package main

import "fmt"

func main() {
	defer_call()
}

func defer_call() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

```

输出结果：

```go
打印后
打印中
打印前
panic: 触发异常
```

参考解析：

defer 的执行顺序是后进先出的，当出现 panic 语句的时候，会按照 defer 后进先出的顺序执行，最后才会执行 panic。