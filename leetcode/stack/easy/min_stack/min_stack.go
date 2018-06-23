package min_stack

/*
155. 最小栈
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
*/
//思路,同时维护一个栈,记录着当前最小值
type MinStack struct {
	stack     []int
	smallList []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:     make([]int, 0),
		smallList: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.smallList) == 0 {
		this.smallList = append(this.smallList, x)
		return
	}

	curMin := this.smallList[len(this.smallList)-1]
	if x < curMin {
		curMin = x
	}
	this.smallList = append(this.smallList, curMin)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	this.stack = this.stack[:len(this.stack)-1]
	this.smallList = this.smallList[:len(this.smallList)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return -1
	}

	return this.smallList[len(this.smallList)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
