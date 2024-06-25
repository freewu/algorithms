package main

// LCR 147. 最小栈
// 请你设计一个 最小栈 。它提供 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// 实现 MinStack 类:
//     MinStack() 初始化堆栈对象。
//     void push(int val) 将元素val推入堆栈。
//     void pop() 删除堆栈顶部的元素。
//     int top() 获取堆栈顶部的元素。
//     int getMin() 获取堆栈中的最小元素。
 
// 示例 1:
// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[2],[-3],[],[],[],[]]
// 输出：
// [null,null,null,null,-3,null,2,-2]
// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(2);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 2.
// minStack.getMin();   --> 返回 -2.

// 提示：
//     -2^31 <= val <= 2^31 - 1
//     pop、top 和 getMin 操作总是在 非空栈 上调用
//     push、pop、top 和 getMin 最多被调用 3 * 10^4 次

import "fmt"

type MinStack struct {
    stack []int // 栈
    minElement int // 用来存放最小元素
}

func Constructor() MinStack {
    return MinStack{stack: []int{}}
}

func (this *MinStack) Push(val int)  {
    // 栈为空时
    if len(this.stack) == 0 {
        this.minElement = val
        this.stack = append(this.stack, val)
    } else if val < this.minElement { // 如果入栈值为最小值
        this.stack = append(this.stack, 2*val - this.minElement)
        this.minElement = val;
    } else {
        this.stack = append(this.stack,val)
    }
}

func (this *MinStack) Pop()  {
    if this.stack[len(this.stack)-1] < this.minElement {
        this.minElement = 2*this.minElement-this.stack[len(this.stack)-1]
    }
    this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
    if this.stack[len(this.stack)-1] < this.minElement {
        return this.minElement
    }
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.minElement
}

// 两个数组维护
type MinStack1 struct {
    min []int
    stack []int
}

func Constructor1() MinStack1 {
    return MinStack1{stack:[]int{}}
}

func (this *MinStack1) Push(val int)  {
    if len(this.min) == 0 || val <= this.GetMin() {
        this.min = append(this.min,val)
    }
    this.stack = append(this.stack,val)
}

func (this *MinStack1) Pop()  {
    if this.Top() == this.GetMin() {
        this.min = this.min[:len(this.min)-1]
    }
    this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack1) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack1) GetMin() int {
    return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
    // MinStack minStack = new MinStack();
    obj := Constructor()
    fmt.Println(obj)
    // minStack.push(-2);
    obj.Push(-2)
    fmt.Println(obj)
    // minStack.push(0);
    obj.Push(0)
    fmt.Println(obj)
    // minStack.push(-3);
    obj.Push(-3)
    fmt.Println(obj)
    // minStack.getMin(); // return -3
    fmt.Println(obj.GetMin()) // -3
    // minStack.pop();
    obj.Pop() 
    fmt.Println(obj)
    // minStack.top();    // return 0
    fmt.Println(obj.Top()) // 0
    // minStack.getMin(); // return -2
    fmt.Println(obj.GetMin()) // -2

    // MinStack minStack = new MinStack();
    obj1 := Constructor1()
    fmt.Println(obj1)
    // minStack.push(-2);
    obj1.Push(-2)
    fmt.Println(obj1)
    // minStack.push(0);
    obj1.Push(0)
    fmt.Println(obj1)
    // minStack.push(-3);
    obj1.Push(-3)
    fmt.Println(obj1)
    // minStack.getMin(); // return -3
    fmt.Println(obj1.GetMin()) // -3
    // minStack.pop();
    obj1.Pop() 
    fmt.Println(obj1)
    // minStack.top();    // return 0
    fmt.Println(obj1.Top()) // 0
    // minStack.getMin(); // return -2
    fmt.Println(obj1.GetMin()) // -2
}