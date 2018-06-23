package stack_to_queue

/*
232. 用栈实现队列
使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。
示例:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false
说明:

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
*/

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{data: make([]int, 0)}
}

func (s *stack) push(x int) {
	s.data = append(s.data, x)
}

func (s *stack) pop() {
	if s.isEmpty() {
		return
	}
	s.data = s.data[:s.size()-1]
}

func (s *stack) top() int {
	if s.isEmpty() {
		return -1
	}

	return s.data[s.size()-1]
}

func (s *stack) size() int {
	return len(s.data)
}

func (s *stack) isEmpty() bool {
	return len(s.data) <= 0
}

type MyQueue struct {
	sIn  *stack
	sOut *stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{sIn: newStack(), sOut: newStack()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.sIn.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.sIn.isEmpty() && this.sOut.isEmpty() {
		return -1
	} else if this.sOut.isEmpty() {
		this.inToOut()
	}

	x := this.sOut.top()
	this.sOut.pop()
	return x
}

func (this *MyQueue) inToOut() {
	for !this.sIn.isEmpty() {
		this.sOut.push(this.sIn.top())
		this.sIn.pop()
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.sIn.isEmpty() && this.sOut.isEmpty() {
		return -1
	} else if this.sOut.isEmpty() {
		this.inToOut()
	}

	return this.sOut.top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.sIn.isEmpty() && this.sOut.isEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */