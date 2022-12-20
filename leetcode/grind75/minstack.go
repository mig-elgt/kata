package grind75

// Min Stack

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

// Implement the MinStack class:

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.

// Example 1:

// Input
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]

// Output
// [null,null,null,null,-3,null,0,-2]

// Explanation
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin(); // return -3
// minStack.pop();
// minStack.top();    // return 0
// minStack.getMin(); // return -

type valMin struct {
	value, min int
}

type MinStack struct {
	data []valMin
}

func NewMinStack() MinStack {
	return MinStack{
		data: []valMin{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.data) == 0 {
		this.data = append(this.data, valMin{val, val})
		return
	}
	min := this.data[len(this.data)-1].min
	if val < min {
		min = val
	}
	this.data = append(this.data, valMin{
		value: val,
		min:   min,
	})
}

func (this *MinStack) Pop() {
	this.data = this.data[0 : len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1].value
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1].min
}
