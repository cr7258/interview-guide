package main

import (
	"fmt"
	"sync"
)

/**
 * @description 两两交替打印数字和字符串: 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
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
				fmt.Printf("%d%d", i, i+1)
				i += 2
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
				// 返回字母；str[i] 返回的是字母对应的 ASCII 码
				fmt.Print(str[i : i+2])
				i += 2
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
