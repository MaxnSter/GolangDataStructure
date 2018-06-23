package queue_to_stack

/*
225. 用队列实现栈
使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
*/
type queue struct {
	data []int
}

func newQ() *queue {
	return &queue{data: make([]int, 0)}
}

func (q *queue) PushToBack(x int) {
	q.data = append(q.data,x)
}

func (q *queue) PeekFromFront() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[0]
}

func (q *queue) PopFromFront() {
	if q.IsEmpty() {
		return
	}

	q.data = q.data[1:q.Size()]
}

func (q *queue) Size() int{
	return len(q.data)
}

func (q *queue) IsEmpty() bool {
	return len(q.data) <= 0
}

type MyStack struct {
	q * queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{q:newQ()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.q.IsEmpty() {
		this.q.PushToBack(x)
		return
	}

	tmp := newQ()
	for !this.q.IsEmpty() {
		tmp.PushToBack(this.q.PeekFromFront())
		this.q.PopFromFront()
	}

	this.q.PushToBack(x)
	for !tmp.IsEmpty() {
		this.q.PushToBack(tmp.PeekFromFront())
		tmp.PopFromFront()
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.Top()
	this.q.PopFromFront()
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q.PeekFromFront()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q.IsEmpty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
