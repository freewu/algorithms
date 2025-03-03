package main

// 面试题 03.05. Sort of Stacks LCCI
// Write a program to sort a stack such that the smallest items are on the top. 
// You can use an additional temporary stack, but you may not copy the elements into any other data structure (such as an array). 
// The stack supports the following operations: push, pop, peek, and isEmpty. 
// When the stack is empty, peek should return -1.

// Example1:
// Input: 
// ["SortedStack", "push", "push", "peek", "pop", "peek"]
// [[], [1], [2], [], [], []]
// Output: 
// [null,null,null,1,null,2]

// Example2:
// Input:  
// ["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
// [[], [], [], [1], [], []]
// Output: 
// [null,null,null,null,null,true]

// Note:
//     The total number of elements in the stack is within the range [0, 5000].

import "fmt"

type SortedStack struct {
    stack1 []int
    stack2 []int
}

func Constructor() SortedStack {
    return SortedStack{ stack1: []int{}, stack2: []int{} }
}

func (this *SortedStack) Push(val int)  {
    for len(this.stack2) > 0 && this.stack2[len(this.stack2) - 1] > val {
        this.stack1 = append(this.stack1, this.stack2[len(this.stack2) - 1])
        this.stack2 = this.stack2[:len(this.stack2) - 1]
    }
    for len(this.stack1) > 0 && this.stack1[len(this.stack1) - 1] < val {
        this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
        this.stack1 = this.stack1[:len(this.stack1) - 1]
    }
    this.stack1 = append(this.stack1, val)
}

func (this *SortedStack) Pop() {
    if this.IsEmpty() { return }
    for len(this.stack2) > 0{
        this.stack1 = append(this.stack1, this.stack2[len(this.stack2) - 1])
        this.stack2 = this.stack2[:len(this.stack2) - 1]
    } 
    this.stack1 = this.stack1[:len(this.stack1) - 1]
}

func (this *SortedStack) Peek() int {
    if this.IsEmpty() { return -1 }
    for len(this.stack2) > 0{
        this.stack1 = append(this.stack1, this.stack2[len(this.stack2) - 1])
        this.stack2 = this.stack2[:len(this.stack2) - 1]
    }
    return this.stack1[len(this.stack1) - 1]
}

func (this *SortedStack) IsEmpty() bool {
    return len(this.stack1) == 0 && len(this.stack2) == 0
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */

func main() {
    // Example1:
    // Input: 
    // ["SortedStack", "push", "push", "peek", "pop", "peek"]
    // [[], [1], [2], [], [], []]
    // Output: 
    // [null,null,null,1,null,2]
    obj1 := Constructor()
    fmt.Println(obj1)
    obj1.Push(1)
    fmt.Println(obj1)
    obj1.Push(2)
    fmt.Println(obj1)
    fmt.Println(obj1.Peek()) // 1
    obj1.Pop()
    fmt.Println(obj1)
    fmt.Println(obj1.Peek()) // 2
    // Example2:
    // Input:  
    // ["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
    // [[], [], [], [1], [], []]
    // Output: 
    // [null,null,null,null,null,true]
    obj2 := Constructor()
    fmt.Println(obj2)
    obj2.Pop()
    fmt.Println(obj2)
    obj2.Pop()
    fmt.Println(obj2)
    obj2.Push(1)
    fmt.Println(obj2)
    obj2.Pop()
    fmt.Println(obj2)
    fmt.Println(obj2.IsEmpty()) // true
}