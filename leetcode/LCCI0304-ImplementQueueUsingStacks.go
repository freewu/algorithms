package main

// 面试题 03.04. Implement Queue using Stacks LCCI
// Implement a MyQueue class which implements a queue using two stacks.

// Example:
// MyQueue queue = new MyQueue();
// queue.push(1);
// queue.push(2);
// queue.peek();  // return 1
// queue.pop();   // return 1
// queue.empty(); // return false

// Notes:
//     You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.
//     Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.
//     You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).

import "fmt"

type MyQueue struct {
    stack1 []int
    stack2  []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    // 入栈1
    this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    // 如果stack2不为空，直接返回栈顶元素
    if len(this.stack2) != 0 {
        x := this.stack2[len(this.stack2) - 1]
        this.stack2 = this.stack2[:len(this.stack2) - 1]
        return x
    }
    for len(this.stack1) != 0 {
        // 将所有栈1的元素都入2栈
        this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
        this.stack1 = this.stack1[:len(this.stack1) - 1]
    }
    return this.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
    if len(this.stack2) != 0 {
        return this.stack2[len(this.stack2) - 1]
    }
    for len(this.stack1) != 0 {
        // 将所有栈1的元素都入2栈
        this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
        this.stack1 = this.stack1[:len(this.stack1) - 1]
    }
    return this.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    if len(this.stack1) == 0 && len(this.stack2) == 0 {
        return true
    }
    return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
    // Example:
    // MyQueue queue = new MyQueue();
    obj := Constructor()
    fmt.Println(obj)
    // queue.push(1);
    obj.Push(1)
    fmt.Println(obj)
    // queue.push(2);
    obj.Push(2)
    fmt.Println(obj)
    // queue.peek();  // return 1
    fmt.Println(obj.Peek()) // 1
    // queue.pop();   // return 1
    fmt.Println(obj.Pop()) // 1
    fmt.Println(obj)
    // queue.empty(); // return false
    fmt.Println(obj.Empty()) // false
}