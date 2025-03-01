package main

// 面试题 03.02. Min Stack LCCI
// How would you design a stack which, in addition to push and pop, has a function min which returns the minimum element? 
// Push, pop and min should all operate in 0(1) time.

// Example:
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> return -3.
// minStack.pop();
// minStack.top();      --> return 0.
// minStack.getMin();   --> return -2.

import "fmt"

// type MinStack struct {
//     diff []int // 保存与最小值的差值
//     mn int
// }

// /** initialize your data structure here. */
// func Constructor() MinStack {
//     return MinStack{ mn: 1 << 31 }
// }

// func (this *MinStack) Push(x int) {
//     this.diff = append(this.diff, x - this.mn)
//     if x < this.mn {
//         this.mn = x
//     }
// }

// func (this *MinStack) Pop() {
//     top := this.diff[len(this.diff) - 1]
//     this.diff = this.diff[:len(this.diff) - 1]
//     if top < 0 {
//         this.mn -= top
//     }
// }

// func (this *MinStack) Top() int {
//     top := this.diff[len(this.diff) - 1]
//     if top < 0 {
//         return this.mn
//     } else {
//         return this.mn + top
//     }
// }

// func (this *MinStack) GetMin() int {
//     return this.mn
// }

type MinStack struct {
    data []int
    mn []int
}

func Constructor() MinStack {
    return MinStack{ data: []int{}, mn: []int{ 1 << 31 } }
}

func (this *MinStack) Push(x int)  {
    this.data = append(this.data, x)
    top := this.mn[len(this.mn) - 1]
    this.mn = append(this.mn, min(x, top))
}

func (this *MinStack) Pop()  {
    this.data = this.data[:len(this.data) - 1]
    this.mn = this.mn[:len(this.mn) - 1]
}

func (this *MinStack) Top() int {
    return this.data[len(this.data) - 1]
}

func (this *MinStack) GetMin() int {
    return this.mn[len(this.mn) - 1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
    // Example:
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
    // minStack.getMin();   --> return -3.
    fmt.Println(obj.GetMin()) // -3
    // minStack.pop();
    obj.Pop()
    fmt.Println(obj)
    // minStack.top();      --> return 0.
    fmt.Println(obj.Top()) // 0
    // minStack.getMin();   --> return -2.
    fmt.Println(obj.GetMin()) // -2
}