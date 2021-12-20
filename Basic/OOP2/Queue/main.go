package queue

import "fmt"
//队列
type Queue []interface{}
//后进
// e.q.  q.Push(123)
func (q *Queue) Push(value interface{}) {
	*q = append(*q, value.(int))
}
//先出
func (q Queue) Pop() interface{} {
	head := (q)[0]
	(q) = (q)[1:]
	return head.(int)
}
func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
func main() {
	uuu := &Queue{1}
	uuu.Push(2)
	fmt.Println(uuu)
	uuu.Pop()
	fmt.Println(uuu)
	fmt.Println(uuu.IsEmpty())
	uuu.Pop()
	fmt.Println(uuu.IsEmpty())

}
