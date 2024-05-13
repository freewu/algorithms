package main

// 716. Max Stack
// Design a max stack data structure that supports the stack operations and supports finding the stack's maximum element.
// Implement the MaxStack class:
//     MaxStack() Initializes the stack object.
//     void push(int x) Pushes element x onto the stack.
//     int pop() Removes the element on top of the stack and returns it.
//     int top() Gets the element on the top of the stack without removing it.
//     int peekMax() Retrieves the maximum element in the stack without removing it.
//     int popMax() Retrieves the maximum element in the stack and removes it. If there is more than one maximum element, only remove the top-most one.

// You must come up with a solution that supports O(1) for each top call and O(logn) for each other call.
 
// Example 1:
// Input
// ["MaxStack", "push", "push", "push", "top", "popMax", "top", "peekMax", "pop", "top"]
// [[], [5], [1], [5], [], [], [], [], [], []]
// Output
// [null, null, null, null, 5, 5, 1, 5, 1, 5]
// Explanation
// MaxStack stk = new MaxStack();
// stk.push(5);   // [5] the top of the stack and the maximum number is 5.
// stk.push(1);   // [5, 1] the top of the stack is 1, but the maximum is 5.
// stk.push(5);   // [5, 1, 5] the top of the stack is 5, which is also the maximum, because it is the top most one.
// stk.top();     // return 5, [5, 1, 5] the stack did not change.
// stk.popMax();  // return 5, [5, 1] the stack is changed now, and the top is different from the max.
// stk.top();     // return 1, [5, 1] the stack did not change.
// stk.peekMax(); // return 5, [5, 1] the stack did not change.
// stk.pop();     // return 1, [5] the top of the stack and the max element is now 5.
// stk.top();     // return 5, [5] the stack did not change.

// Constraints:
//     -10^7 <= x <= 10^7
//     At most 10^5 calls will be made to push, pop, top, peekMax, and popMax.
//     There will be at least one element in the stack when pop, top, peekMax, or popMax is called.

import "fmt"
import "container/heap"

type MaxStack struct {
    elems []*node // 栈存储元素
    maxHeap MaxHeap // 最大堆
}

type node struct {
    val   int
    alive bool
}

func Constructor() MaxStack {
    return MaxStack{nil, MaxHeap{}}
}

// 直接推入一个新节点
func (this *MaxStack) Push(x int)  {
    p := &node{x, true}
    this.elems = append(this.elems, p)
    heap.Push(&this.maxHeap, p)
}

// 找到离栈顶最近的第一个有效元素，取出，并将其置为无效
func (this *MaxStack) Pop() int {
    this.removDead()
    p := this.elems[len(this.elems)-1]
    this.elems = this.elems[:len(this.elems)-1]
    p.alive = false
    return p.val
}

// 找到离栈顶最近的一个有效元素，返回其值
func (this *MaxStack) Top() int {
    this.removDead()
    return this.elems[len(this.elems)-1].val
}

// 去除栈中的无效元素
func (this *MaxStack) removDead() {
    // 更新elems
    elems := this.elems
    for ; !elems[len(elems)-1].alive; elems = elems[:len(elems)-1] {}
    this.elems = elems
    // 更新堆
    p := this.maxHeap[0]
    for !p.alive {
        heap.Pop(&this.maxHeap)
        p = this.maxHeap[0]
    }
}

func (this *MaxStack) PeekMax() int {
    return this.peekMax().val
}

func (this *MaxStack) peekMax() *node {
    this.removDead()
    return this.maxHeap[0]
}

func (this *MaxStack) PopMax() int {
    // assert stack not empty
    p := this.peekMax()
    p.alive = false
    return p.val
}

// 最大堆
type MaxHeap []*node
func (h MaxHeap) Len() int { return len(h); }
func (h MaxHeap) Less(i, j int) bool { return h[i].val >= h[j].val; }
func (h *MaxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i]; }
func (h *MaxHeap) Push(x interface{})  { *h = append(*h, x.(*node)); }
func (h *MaxHeap) Pop() interface{} {
    res := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    return res
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */

func main() {
    // MaxStack stk = new MaxStack();
    obj := Constructor()
    fmt.Println(obj)
    // stk.push(5);   // [5] the top of the stack and the maximum number is 5.
    obj.Push(5)
    fmt.Println(obj)
    // stk.push(1);   // [5, 1] the top of the stack is 1, but the maximum is 5.
    obj.Push(1)
    fmt.Println(obj)
    // stk.push(5);   // [5, 1, 5] the top of the stack is 5, which is also the maximum, because it is the top most one.
    obj.Push(5)
    fmt.Println(obj)
    // stk.top();     // return 5, [5, 1, 5] the stack did not change.
    fmt.Println(obj.Top()) // 5
    // stk.popMax();  // return 5, [5, 1] the stack is changed now, and the top is different from the max.
    fmt.Println(obj.PopMax()) // 5
    fmt.Println(obj)
    // stk.top();     // return 1, [5, 1] the stack did not change.
    fmt.Println(obj.Top()) // 1
    // stk.peekMax(); // return 5, [5, 1] the stack did not change.
    fmt.Println(obj.PopMax()) // 5
    fmt.Println(obj)
    // stk.pop();     // return 1, [5] the top of the stack and the max element is now 5.
    fmt.Println(obj.Pop()) // 1
    fmt.Println(obj)
    // stk.top();     // return 5, [5] the stack did not change.
    fmt.Println(obj.Top()) // 5
    fmt.Println(obj)
}