package main

import (
	"fmt"
)

/**
 * @description 值接收者和指针接收者的区别
 * @author chengzw
 * @since 2022/6/26
 * @link https://golang.design/go-questions/interface/receiver/
 */

/**
如果实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的方法。
*/

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	var c coder = &Gopher{"Go"}
	c.code()
	c.debug()
}

/**
但是如果我们把 main 函数的第一条语句换一下：
func main() {
	var c coder = Gopher{"Go"}
	c.code()
	c.debug()
}
src/main.go:23:6: cannot use Gopher literal (type Gopher) as type coder in assignment:
	Gopher does not implement coder (debug method has pointer receiver)

第二次报错是说，Gopher 没有实现 coder。很明显了吧，因为 Gopher 类型并没有实现 debug 方法；表面上看，
*Gopher 类型也没有实现 code 方法，但是因为 Gopher 类型实现了 code 方法，所以让 *Gopher 类型自动拥有了 code 方法。
*/
