package main

// 1381. Design a Stack With Increment Operation
// Design a stack that supports increment operations on its elements.
// Implement the CustomStack class:
//     CustomStack(int maxSize) 
//         Initializes the object with maxSize which is the maximum number of elements in the stack.
//     void push(int x) 
//         Adds x to the top of the stack if the stack has not reached the maxSize.
//     int pop() 
//         Pops and returns the top of the stack or -1 if the stack is empty.
//     void inc(int k, int val) 
//         Increments the bottom k elements of the stack by val. 
//         If there are less than k elements in the stack, increment all the elements in the stack.

// Example 1:
// Input
// ["CustomStack","push","push","pop","push","push","push","increment","increment","pop","pop","pop","pop"]
// [[3],[1],[2],[],[2],[3],[4],[5,100],[2,100],[],[],[],[]]
// Output
// [null,null,null,2,null,null,null,null,null,103,202,201,-1]
// Explanation
// CustomStack stk = new CustomStack(3); // Stack is Empty []
// stk.push(1);                          // stack becomes [1]
// stk.push(2);                          // stack becomes [1, 2]
// stk.pop();                            // return 2 --> Return top of the stack 2, stack becomes [1]
// stk.push(2);                          // stack becomes [1, 2]
// stk.push(3);                          // stack becomes [1, 2, 3]
// stk.push(4);                          // stack still [1, 2, 3], Do not add another elements as size is 4
// stk.increment(5, 100);                // stack becomes [101, 102, 103]
// stk.increment(2, 100);                // stack becomes [201, 202, 103]
// stk.pop();                            // return 103 --> Return top of the stack 103, stack becomes [201, 202]
// stk.pop();                            // return 202 --> Return top of the stack 202, stack becomes [201]
// stk.pop();                            // return 201 --> Return top of the stack 201, stack becomes []
// stk.pop();                            // return -1 --> Stack is empty return -1.

// Constraints:
//     1 <= maxSize, x, k <= 1000
//     0 <= val <= 100
//     At most 1000 calls will be made to each method of increment, push and pop each separately.

import "fmt"

// type CustomStack struct {
//     stack   []int
//     maxSize int
// }

// func Constructor(maxSize int) CustomStack {
//     return CustomStack{ make([]int, 0), maxSize, }
// }

// // O(1)
// func (this *CustomStack) Push(x int) {
//     if len(this.stack) < this.maxSize {
//         this.stack = append(this.stack, x)
//     }
// }

// // O(1)
// func (this *CustomStack) Pop() int {
//     res := -1
//     if 0 < len(this.stack) {
//         res = this.stack[len(this.stack) - 1]
//         this.stack = this.stack[:len(this.stack)-1]  // pop
//     }
//     return res
// }

// // O(k)
// func (this *CustomStack) Increment(k int, val int) {
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     for i := 0; i < min(len(this.stack), k); i++ {
//         this.stack[i] += val
//     }
// }

type CustomStack struct {
    Stack   []int
    Index   int
    MaxSize int
}

func Constructor(maxSize int) CustomStack {
    return CustomStack{
        Stack:   make([]int, maxSize),
        MaxSize: maxSize,
    }
}

func (this *CustomStack) Push(x int) {
    if this.Index < this.MaxSize {
        if this.Index < 0 {
            this.Index = 0
        }
        this.Stack[this.Index] = x
        this.Index++
    }
}

func (this *CustomStack) Pop() int {
    res := -1
    if this.Index > 0 {
        res = this.Stack[this.Index-1]
    }
    this.Index--
    return res
}

func (this *CustomStack) Increment(k int, val int) {
    bound := k
    if k > this.Index {
        bound = this.Index
    }
    for i := 0; i < bound; i++ {
        this.Stack[i] += val
    }
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

func main() {
    // CustomStack stk = new CustomStack(3); // Stack is Empty []
    obj := Constructor(3)
    fmt.Println(obj) // []
    // stk.push(1);                          // stack becomes [1]
    obj.Push(1)
    fmt.Println(obj) // [1]
    // stk.push(2);                          // stack becomes [1, 2]
    obj.Push(2)
    fmt.Println(obj) // [1, 2]
    // stk.pop();                            // return 2 --> Return top of the stack 2, stack becomes [1]
    fmt.Println(obj.Pop()) // 2
    fmt.Println(obj) // [1]
    // stk.push(2);                          // stack becomes [1, 2]
    obj.Push(2)
    fmt.Println(obj) // [1, 2]
    // stk.push(3);                          // stack becomes [1, 2, 3]
    obj.Push(3)
    fmt.Println(obj) // [1, 2, 3]
    // stk.push(4);                          // stack still [1, 2, 3], Do not add another elements as size is 4
    obj.Push(4)
    fmt.Println(obj) // [1, 2, 3]
    // stk.increment(5, 100);                // stack becomes [101, 102, 103]
    obj.Increment(5, 100)
    fmt.Println(obj) // [101, 102, 103]
    // stk.increment(2, 100);                // stack becomes [201, 202, 103]
    obj.Increment(2, 100)
    fmt.Println(obj) // [201, 202, 103]
    // stk.pop();                            // return 103 --> Return top of the stack 103, stack becomes [201, 202]
    fmt.Println(obj.Pop()) // 103
    fmt.Println(obj) // [201, 202]
    // stk.pop();                            // return 202 --> Return top of the stack 202, stack becomes [201]
    fmt.Println(obj.Pop()) // 202
    fmt.Println(obj) // [201]
    // stk.pop();                            // return 201 --> Return top of the stack 201, stack becomes []
    fmt.Println(obj.Pop()) // 201
    fmt.Println(obj) // []
    // stk.pop();                            // return -1 --> Stack is empty return -1.
    fmt.Println(obj.Pop()) // -1
    fmt.Println(obj) // []
}