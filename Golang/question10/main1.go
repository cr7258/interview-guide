package main

import (
	"fmt"
	"sync"
)

/**
 * @description 交替打印数字和字符串: 1A2B3C4D5E6F7G8H9I10J11K12L13M14N15O16P17Q18R19S20T21U22V23W24X25Y26Z
 * @author chengzw
 * @since 2022/6/26
 * @link
 */
func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				fmt.Print(str[i : i+1])
				i++
				if i >= len(str) {
					wait.Done()
					return
				}
				number <- true
			}

		}
	}(&wait)
	// 让数字先开始打印
	number <- true
	// 等待循环结束，表示整个打印可以结束了
	wait.Wait()
	// 最后关闭 channel，防止内存泄露
	close(letter)
	close(number)
}
