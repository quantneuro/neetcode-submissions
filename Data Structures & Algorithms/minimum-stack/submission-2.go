type MinStack struct {
	stack []int
	mstack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack=append(this.stack,val)

	if len(this.mstack) == 0{
		this.mstack=append(this.mstack,val)
	}else if val <this.mstack[len(this.mstack)-1]{
		this.mstack=append(this.mstack,val)
	}else{
		this.mstack=append(this.mstack,this.mstack[len(this.mstack)-1])
	}
}

func (this *MinStack) Pop() {
	this.stack=this.stack[:len(this.stack)-1]
	this.mstack=this.mstack[:len(this.mstack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mstack[len(this.mstack)-1]
}
