package main

// 225. Implement Stack using Queues
// Implement a last-in-first-out (LIFO) stack using only two queues. 
// The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).
// Implement the MyStack class:
//         void push(int x) Pushes element x to the top of the stack.
//         int pop() Removes the element on the top of the stack and returns it.
//         int top() Returns the element on the top of the stack.
//         boolean empty() Returns true if the stack is empty, false otherwise.

// Notes:
//     You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
//     Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.
 
// Example 1:
// Input
// ["MyStack", "push", "push", "top", "pop", "empty"]
// [[], [1], [2], [], [], []]
// Output
// [null, null, null, 2, 2, false]
// Explanation
// MyStack myStack = new MyStack();
// myStack.push(1);
// myStack.push(2);
// myStack.top(); // return 2
// myStack.pop(); // return 2
// myStack.empty(); // return False

// Constraints:
//         1 <= x <= 9
//         At most 100 calls will be made to push, pop, top, and empty.
//         All the calls to pop and top are valid.

import "fmt"

type MyStack struct {
    enque []int
	deque []int
}


func Constructor() MyStack {
    return MyStack{[]int{}, []int{}}
}


func (this *MyStack) Push(x int)  {
    this.enque = append(this.enque, x)
}


func (this *MyStack) Pop() int {
	length := len(this.enque)
	for i := 0; i < length-1; i++ {
		this.deque = append(this.deque, this.enque[0])
		this.enque = this.enque[1:]
	}
	topEle := this.enque[0]
	this.enque = this.deque
	this.deque = nil
	return topEle
}

func (this *MyStack) Top() int {
	topEle := this.Pop()
	this.enque = append(this.enque, topEle)
	return topEle
}

func (this *MyStack) Empty() bool {
	return len(this.enque) == 0
}


type MyStack1 struct {
	data []int
}

func Constructor1() MyStack1 {
	return MyStack1{ []int{} }
}

func (this *MyStack1) Push(x int) {
	// 压栈是加到 slice 前面
	this.data = append([]int{x}, this.data...)
}

func (this *MyStack1) Pop() int {
	res := this.Top()
	// pop 是从前面出来
	this.data = this.data[1:len(this.data)]
	return res
}

func (this *MyStack1) Top() int {
	return this.data[0]
}

func (this *MyStack1) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

 func main() {
    obj := Constructor();
    obj.Push(1);
	obj.Push(2);
    fmt.Println("obj.Top(): ", obj.Top());
    fmt.Println("obj.Empty(): ", obj.Empty());
	fmt.Println("obj.Pop(): ", obj.Pop());
	fmt.Println("obj.Pop(): ", obj.Pop());
	fmt.Println("obj.Empty(): ", obj.Empty());

	obj1 := Constructor1();
    obj1.Push(1);
	obj1.Push(2);
    fmt.Println("obj1.Top(): ", obj1.Top());
    fmt.Println("obj1.Empty(): ", obj1.Empty());
	fmt.Println("obj1.Pop(): ", obj1.Pop());
	fmt.Println("obj1.Pop(): ", obj1.Pop());
	fmt.Println("obj1.Empty(): ", obj1.Empty());
 }