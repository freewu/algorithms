package main

// 面试题 03.03. Stack of Plates LCCI
// Imagine a (literal) stack of plates. If the stack gets too high, it might topple. 
// Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold. 
// Implement a data structure SetOfStacks that mimics this. 
// SetOfStacks should be composed of several stacks and should create a new stack once the previous one exceeds capacity. 
// SetOfStacks.push() and SetOfStacks.pop() should behave identically to a single stack (that is, pop() should return the same values as it would if there were just a single stack). 
// Follow Up: Implement a function popAt(int index) which performs a pop operation on a specific sub-stack.

// You should delete the sub-stack when it becomes empty. pop, popAt should return -1 when there's no element to pop.

// Example1:
// Input: 
// ["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
// [[1], [1], [2], [1], [], []]
// Output: 
// [null, null, null, 2, 1, -1]

// Example2:
// Input: 
// ["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
// [[2], [1], [2], [3], [0], [0], [0]]
// Output: 
// [null, null, null, null, 2, 1, 3]

import "fmt"

type StackOfPlates struct {
    stacks [][]int
    cap    int
}

func Constructor(cap int) StackOfPlates {
    return StackOfPlates{ stacks: make([][]int, 0), cap: cap }
}

// Push 将一个值压入栈
func (this *StackOfPlates) Push(val int) {
    if this.cap <= 0 { return } // cap <= 0 时无法存储元素
    // 如果最后一个子栈满了或者不存在，创建一个新子栈
    if len(this.stacks) == 0 || len(this.stacks[len(this.stacks) - 1]) == this.cap {
        this.stacks = append(this.stacks, []int{})
    }
    // 将值加入最后一个子栈
    this.stacks[len(this.stacks) - 1] = append(this.stacks[len(this.stacks) - 1], val)
}

// Pop 弹出最后一个子栈的栈顶元素
func (this *StackOfPlates) Pop() int {
    if len(this.stacks) == 0 { return -1 } // 栈为空
    lastStack := this.stacks[len(this.stacks) - 1]
    val := lastStack[len(lastStack) - 1] // 获取最后一个子栈
    this.stacks[len(this.stacks) - 1] = lastStack[:len(lastStack) - 1]     // 删除栈顶元素
    if len(this.stacks[len(this.stacks) - 1]) == 0 { // 如果最后一个子栈为空，则移除该子栈
        this.stacks = this.stacks[:len(this.stacks) - 1]
    }
    return val
}

// PopAt 弹出指定子栈 index 的栈顶元素
func (this *StackOfPlates) PopAt(index int) int {
    if index < 0 || index >= len(this.stacks) { return -1 } // 索引无效
    stack := this.stacks[index] // 获取指定子栈
    val := stack[len(stack) - 1]
    this.stacks[index] = stack[:len(stack) - 1] // 删除栈顶元素
    if len(this.stacks[index]) == 0 { // 如果该子栈为空，移除该子栈
        this.stacks = append(this.stacks[:index], this.stacks[index + 1:]...)
    }
    return val
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */

func main() {
    // Example1:
    // Input: 
    // ["StackOfPlates", "push", "push", "popAt", "pop", "pop"]
    // [[1], [1], [2], [1], [], []]
    // Output: 
    // [null, null, null, 2, 1, -1]
    obj1 := Constructor(1)
    fmt.Println(obj1)
    obj1.Push(1)
    fmt.Println(obj1)
    obj1.Push(1)
    fmt.Println(obj1)
    fmt.Println(obj1.PopAt(1)) // 2
    fmt.Println(obj1.Pop()) // 1
    fmt.Println(obj1.Pop()) // -1
    // Example2:
    // Input: 
    // ["StackOfPlates", "push", "push", "push", "popAt", "popAt", "popAt"]
    // [[2], [1], [2], [3], [0], [0], [0]]
    // Output: 
    // [null, null, null, null, 2, 1, 3]
    obj2 := Constructor(2)
    fmt.Println(obj2)
    obj2.Push(1)
    fmt.Println(obj2)
    obj2.Push(2)
    fmt.Println(obj2)
    obj2.Push(3)
    fmt.Println(obj2)
    fmt.Println(obj2.PopAt(0)) // 2
    fmt.Println(obj2.PopAt(0)) // 1
    fmt.Println(obj2.PopAt(0)) // 3
}