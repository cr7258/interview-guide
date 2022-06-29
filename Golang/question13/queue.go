package question13

import (
	"fmt"
	"sync"
)

/**
 * @description 实现一个线程安全的队列
 * @author chengzw
 * @since 2022/6/29
 */

type Queue struct {
	sync.Mutex
	data []interface{}
}

func (q *Queue) Push(data int) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, data)
}

func (q *Queue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.data) > 0 {
		data := q.data[0]
		q.data = q.data[1:]
		return data, true
	}
	return nil, false
}

func main() {
	q := &Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
