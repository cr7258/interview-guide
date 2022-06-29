package question13

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := &Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)

	num1, _ := q.Pop()
	num2, _ := q.Pop()
	num3, _ := q.Pop()
	num4, ok4 := q.Pop()

	if !(num1 == 1 && num2 == 2 && num3 == 3 && num4 == nil && !ok4) {
		t.Errorf("测试失败")
	}
}
