## 问题描述
使⽤两个 goroutine 交替打印序列，⼀个 goroutine 打印数字， 另外⼀个 goroutine 打印字⺟，最终效果如下：
```bash
1A2B3C4D5E6F7G8H9I10J11K12L13M14N15O16P17Q18R19S20T21U22V23W24X25Y26Z
```
## 解题思路
使用 channel 来控制打印的进度。使用两个 channel，来分别控制数字和字母的打印进度，数字打印完通过 channel 通知数字打印，
数字打印完通过 channel 通知字母打印。如此周而复始，直到终止条件。



## 参考资料
- [面试题:交替打印数字和字符串](https://blog.hujm.net/post/%E9%9D%A2%E8%AF%95%E9%A2%98%E4%BA%A4%E6%9B%BF%E6%89%93%E5%8D%B0%E6%95%B0%E5%AD%97%E5%92%8C%E5%AD%97%E7%AC%A6%E4%B8%B2/)