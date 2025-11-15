package main

// 232. Implement Queue using Stacks
// Implement a first in first out (FIFO) queue using only two stacks. 
// The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

// Implement the MyQueue class:
//     void push(int x) Pushes element x to the back of the queue.
//     int pop() Removes the element from the front of the queue and returns it.
//     int peek() Returns the element at the front of the queue.
//     boolean empty() Returns true if the queue is empty, false otherwise.

// Notes:
//     You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
//     Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.

// Example 1:
// Input
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
// [[], [1], [2], [], [], []]
// Output
// [null, null, null, 1, 1, false]
// Explanation
// MyQueue myQueue = new MyQueue();
// myQueue.push(1); // queue is: [1]
// myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
// myQueue.peek(); // return 1
// myQueue.pop(); // return 1, queue is [2]
// myQueue.empty(); // return false
 
// Constraints:
//     1 <= x <= 9
//     At most 100 calls will be made to push, pop, peek, and empty.
//     All the calls to pop and peek are valid.

// Follow-up: Can you implement the queue such that each operation is amortized O(1) time complexity? 
// In other words, performing n operations will take overall O(n) time even if one of those operations may take longer.

import "fmt"

type MyQueue struct {
    // 将一个栈当作输入栈，用于压入 push 传入的数据；另一个栈当作输出栈，用于 pop 和 peek 操作。
    inStack, outStack []int
}

func Constructor() MyQueue {
    return MyQueue{}
}

func (q *MyQueue) Push(x int) {
    q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
    for len(q.inStack) > 0 {
        q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
        q.inStack = q.inStack[:len(q.inStack)-1]
    }
}

func (q *MyQueue) Pop() int {
	// 	每次 pop 或 peek 时，若输出栈为空则将输入栈的全部数据依次弹出并压入输出栈，这样输出栈从栈顶往栈底的顺序就是队列从队首往队尾的顺序
    if len(q.outStack) == 0 {
        q.in2out()
    }
    x := q.outStack[len(q.outStack)-1]
    q.outStack = q.outStack[:len(q.outStack)-1]
    return x
}

func (q *MyQueue) Peek() int {
	// 每次 pop 或 peek 时，若输出栈为空则将输入栈的全部数据依次弹出并压入输出栈
    if len(q.outStack) == 0 {
        q.in2out()
    }
    return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
    return len(q.inStack) == 0 && len(q.outStack) == 0
}

func main() {
	myQueue := Constructor();
	myQueue.Push(1) // queue is: [1]
	fmt.Println(myQueue)
	myQueue.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Println(myQueue)
	fmt.Println("myQueue.Peek()",myQueue.Peek()) // return 1
	fmt.Println(myQueue)
	fmt.Println("myQueue.Pop()",myQueue.Pop()) // return 1, queue is [2]
	fmt.Println(myQueue)
	fmt.Println(myQueue.Empty()) // return false
	fmt.Println(myQueue)

    // MyQueue myQueue = new MyQueue();
    q1 := Constructor();
    // myQueue.push(1); // queue is: [1]
    q1.Push(1)
    fmt.Println(q1) // [1]
    // myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
    q1.Push(2)
    fmt.Println(q1) // [1 2]
    // myQueue.peek(); // return 1
    fmt.Println(q1.Peek()) // 1
    // myQueue.pop(); // return 1, queue is [2]
    fmt.Println(q1.Pop()) // 1
    fmt.Println(q1) // [2]
    // myQueue.empty(); // return false
    fmt.Println(q1.Empty()) // false
}