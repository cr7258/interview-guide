请解释下面函数的输出结果？
- 0：make 创建切片时，长度设置为 1，第 1 个元素被设置了 int 类型的初始值 0。
- 100：在 add 函数中直接修改切片指针指向的数组的内容，生效。
- 2：第二个元素的值。
- 3：切片长度为 3，在 add 方法中虽然添加了 2 个元素，但是 append 方法实际上是在修改指针，并不是在修改指针指向的数据。因为在外面调用 add 方法时传入的 arr 和形参 arr 并不是一个值，是值的拷贝，只是他们的值都一样，都指向切片中的数据而已。
- 10：切片中的数组未发生扩容，容量等于 make 创建切片时的容量。

```go
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

# 输出
0,100,2,3,10
```

如果想要让 add 方法修改成功，可以传递切片的指针。

```go
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
```