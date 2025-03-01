package main

// 面试题 03.01. Three in One LCCI
// Describe how you could use a single array to implement three stacks.

// You should implement push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum) methods. 
// stackNum is the index of the stack. value is the value that pushed to the stack.

// The constructor requires a stackSize parameter, which represents the size of each stack.

// Example1:
// Input: 
// ["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
// [[1], [0, 1], [0, 2], [0], [0], [0], [0]]
//  Output: 
// [null, null, null, 1, -1, -1, true]
// Explanation: When the stack is empty, `pop, peek` return -1. When the stack is full, `push` does nothing.

// Example2:
// Input: 
// ["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
// [[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
//  Output: 
// [null, null, null, null, 2, 1, -1, -1]

import "fmt"

// type TripleInOne struct {
//     data []int
// }

// func Constructor(stackSize int) TripleInOne {
//     return TripleInOne{ append(make([]int, stackSize * 3), stackSize * 3, 2, 1, 0)}
// }

// func (this *TripleInOne) Push(stackNum int, value int) {
//     if this.data[len(this.data) + ^stackNum] < this.data[len(this.data) + ^3] {
//         this.data[this.data[len(this.data) + ^stackNum]] = value
//         this.data[len(this.data) + ^stackNum] += 3
//     }
// }

// func (this *TripleInOne) Pop(stackNum int) int {
//     if this.data[len(this.data) + ^stackNum] >= 3 {
//         this.data[len(this.data) + ^stackNum] -= 3
//         return this.data[this.data[len(this.data) + ^stackNum]]
//     }
//     return -1
// }

// func (this *TripleInOne) Peek(stackNum int) int {
//     if this.data[len(this.data) + ^stackNum] >= 3 {
//         return this.data[this.data[len(this.data) + ^stackNum] - 3]
//     }
//     return -1
// }

// func (this *TripleInOne) IsEmpty(stackNum int) bool {
//     return this.data[len(this.data) + ^stackNum] < 3
// }

type TripleInOne struct {
    stack     []int // 用于存储所有栈的数据
    stackSize int   // 每个栈的大小
    tops      []int // 栈顶索引数组，每个栈有一个对应的栈顶
}

func Constructor(stackSize int) TripleInOne {
    // 初始化三个栈的栈顶索引
    return TripleInOne{
        stack:     make([]int, stackSize * 3), // 提前分配空间
        stackSize: stackSize,
        tops:      []int{-1, stackSize - 1, stackSize * 2 - 1},
    }
}

// Push 将值压入指定栈
func (this *TripleInOne) Push(stackNum int, value int) {
    if this.tops[stackNum] <= (stackNum+1)*this.stackSize - 2 { // 判断是否栈满
        this.tops[stackNum]++
        this.stack[this.tops[stackNum]] = value
    }
}

// Pop 从指定栈弹出值
func (this *TripleInOne) Pop(stackNum int) int {
    if this.IsEmpty(stackNum) { // 判断是否为空栈
        return -1
    }
    value := this.stack[this.tops[stackNum]]
    this.tops[stackNum]-- // 弹出元素
    return value
}

func (this *TripleInOne) Peek(stackNum int) int {
    if this.IsEmpty(stackNum) { // 判断是否为空栈
        return -1
    }
    return this.stack[this.tops[stackNum]]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
    return this.tops[stackNum] < stackNum * this.stackSize
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */

func main() {
    // Example1:
    // Input: 
    // ["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
    // [[1], [0, 1], [0, 2], [0], [0], [0], [0]]
    //  Output: 
    // [null, null, null, 1, -1, -1, true]
    // Explanation: When the stack is empty, `pop, peek` return -1. When the stack is full, `push` does nothing.
    obj1 := Constructor(1)
    fmt.Println(obj1)
    obj1.Push(0, 1)
    fmt.Println(obj1)
    obj1.Push(0, 2)
    fmt.Println(obj1)
    fmt.Println(obj1.Pop(0)) // 1
    fmt.Println(obj1.Pop(0)) // -1
    fmt.Println(obj1.Pop(0)) // -1
    fmt.Println(obj1.IsEmpty(0)) // true
    // Example2:
    // Input: 
    // ["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
    // [[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
    //  Output: 
    // [null, null, null, null, 2, 1, -1, -1]
    obj2 := Constructor(2)
    fmt.Println(obj2)
    obj2.Push(0, 1)
    fmt.Println(obj2)
    obj2.Push(0, 2)
    fmt.Println(obj2)
    obj2.Push(0, 3)
    fmt.Println(obj2)
    fmt.Println(obj2.Pop(0)) // 2
    fmt.Println(obj2.Pop(0)) // 1
    fmt.Println(obj2.Pop(0)) // -1
    fmt.Println(obj2.Peek(0)) // -1
}