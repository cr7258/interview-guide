## 1 Go 语言中 make 和 new 有什么区别？
-   `new`函数主要是为类型申请一片内存空间，**返回执行内存的指针**。
-   `make`函数能够分配并初始化类型所需的内存空间和结构，**返回复合类型的本身**。`make` 函数仅支持 `channel`、`map`、`slice` 三种类型，**因为这三种类型就是引用类型，所以就没必要返回他们的指针了**。
-   `new`函数在日常开发中使用是比较少的，可以被替代。
-   `make`函数初始化`slice`会初始化零值，日常开发要注意这个问题。

参考资料：
- [Go语言中make和new有什么区别？](https://asong.cloud/make%E5%92%8Cnew%E6%9C%89%E4%BB%80%E4%B9%88%E5%8C%BA%E5%88%AB/)
- [5.5 make 和 new](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-make-and-new/)

## 2 select 是随机的还是顺序的？
- select 语句不使用 default 分支时，处于**阻塞状态**，直到其中一个 channel 的收/发操作准备就绪（或者 channel 关闭或者缓冲区有值），如果同时有多个 channel 的收/发操作准备就绪（或者 channe l关闭）则**随机**选择其中一个。
- select 语句使用 default 分支时，处于**非阻塞状态**，从所有准备就绪（或者 channel 关闭或者缓冲区有值）的 channel 中**随机**选择其中一个，如果没有则执行 default 分支。

参考资料
- [Go面试：select是随机的还是顺序的？](https://blog.csdn.net/pengpengzhou/article/details/107036700)
- [go : select 的执行顺序](https://segmentfault.com/a/1190000022520711)