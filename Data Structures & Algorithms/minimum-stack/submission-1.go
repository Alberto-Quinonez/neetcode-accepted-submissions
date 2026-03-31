type MinStack struct {
	stack []int
	mStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		mStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack,val)
	currMin := val
	if len(this.mStack) > 0 {
		currMin = this.mStack[len(this.mStack)-1]
	}
	this.mStack = append(this.mStack,min(currMin,val))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.mStack = this.mStack[:len(this.mStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mStack[len(this.mStack)-1]
}
