package main

// 1172. Dinner Plate Stacks
// You have an infinite number of stacks arranged in a row and numbered (left to right) from 0, each of the stacks has the same maximum capacity.

// Implement the DinnerPlates class:
//     DinnerPlates(int capacity) 
//         Initializes the object with the maximum capacity of the stacks capacity.
//     void push(int val) 
//         Pushes the given integer val into the leftmost stack with a size less than capacity.
//     int pop() 
//         Returns the value at the top of the rightmost non-empty stack and removes it from that stack, 
//         and returns -1 if all the stacks are empty.
//     int popAtStack(int index) 
//         Returns the value at the top of the stack with the given index index and removes it from 
//         that stack or returns -1 if the stack with that given index is empty.

// Example 1:
// Input
// ["DinnerPlates", "push", "push", "push", "push", "push", "popAtStack", "push", "push", "popAtStack", "popAtStack", "pop", "pop", "pop", "pop", "pop"]
// [[2], [1], [2], [3], [4], [5], [0], [20], [21], [0], [2], [], [], [], [], []]
// Output
// [null, null, null, null, null, null, 2, null, null, 20, 21, 5, 4, 3, 1, -1]
// Explanation: 
// DinnerPlates D = DinnerPlates(2);  // Initialize with capacity = 2
// D.push(1);
// D.push(2);
// D.push(3);
// D.push(4);
// D.push(5);         // The stacks are now:  2  4
//                                            1  3  5
//                                            ﹈ ﹈ ﹈
// D.popAtStack(0);   // Returns 2.  The stacks are now:     4
//                                                        1  3  5
//                                                        ﹈ ﹈ ﹈
// D.push(20);        // The stacks are now: 20  4
//                                            1  3  5
//                                            ﹈ ﹈ ﹈
// D.push(21);        // The stacks are now: 20  4 21
//                                            1  3  5
//                                            ﹈ ﹈ ﹈
// D.popAtStack(0);   // Returns 20.  The stacks are now:     4 21
//                                                         1  3  5
//                                                         ﹈ ﹈ ﹈
// D.popAtStack(2);   // Returns 21.  The stacks are now:     4
//                                                         1  3  5
//                                                         ﹈ ﹈ ﹈ 
// D.pop()            // Returns 5.  The stacks are now:      4
//                                                         1  3 
//                                                         ﹈ ﹈  
// D.pop()            // Returns 4.  The stacks are now:   1  3 
//                                                         ﹈ ﹈   
// D.pop()            // Returns 3.  The stacks are now:   1 
//                                                         ﹈   
// D.pop()            // Returns 1.  There are no stacks.
// D.pop()            // Returns -1.  There are still no stacks.

// Constraints:
//     1 <= capacity <= 2 * 10^4
//     1 <= val <= 2 * 10^4
//     0 <= index <= 10^5
//     At most 2 * 10^5 calls will be made to push, pop, and popAtStack.

import "fmt"
import "container/heap"
import "sort"

// // Heap
// type MinHeap []int

// func (h MinHeap) Len() int { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i] < h[j]}
// func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
// func (h MinHeap) Top() interface{} { return h[0] }
// func (h *MinHeap) Push(x interface{}) {
//     // Push and Pop use pointer receivers because they modify the slice's length,
//     // not just its contents.
//     *h = append(*h, x.(int))
// }
// func (h *MinHeap) Pop() interface{} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[0 : n-1]
//     return x
// }

// type DinnerPlates struct {
//     pushHeap *MinHeap
//     popHeap  *MinHeap
//     st       map[int]*[]int
//     capacity int
// }

// func Constructor(capacity int) DinnerPlates {
//     return DinnerPlates{
//         pushHeap: &MinHeap{},
//         popHeap:  &MinHeap{},
//         st:       map[int]*[]int{},
//         capacity: capacity,
//     }
// }

// func (this *DinnerPlates) Push(val int) {
//     index := len(this.st)
//     if this.pushHeap.Len() > 0 {
//         index = heap.Pop(this.pushHeap).(int)
//     }
//     if this.st[index] == nil {
//         this.st[index] = &[]int{}
//         heap.Push(this.popHeap, -index)
//     }
//     *this.st[index] = append(*this.st[index], val)
//     if len(*this.st[index]) < this.capacity {
//         heap.Push(this.pushHeap, index)
//     }
// }

// func (this *DinnerPlates) Pop() int {
//     v := -1
//     for this.popHeap.Len() > 0 && this.st[-this.popHeap.Top().(int)] == nil {
//         heap.Pop(this.popHeap)
//     }
//     if this.popHeap.Len() > 0 {
//         index := -heap.Pop(this.popHeap).(int)
//         if len(*this.st[index]) == this.capacity {
//             heap.Push(this.pushHeap, index)
//         }
//         _st := *this.st[index]
//         v = _st[len(_st)-1]
//         *this.st[index] = _st[:len(_st)-1]

//         if len(*this.st[index]) == 0 {
//             delete(this.st, index)
//         } else {
//             heap.Push(this.popHeap, -index)
//         }
//     }
//     return v
// }


// func (this *DinnerPlates) PopAtStack(index int) int {
//     v := -1
//     st_index, f := this.st[index]
//     if index >= 0 && index <= len(this.st) && f {
//         if len(*st_index) == this.capacity {
//             heap.Push(this.pushHeap, index)
//         }
//         if len(*st_index) > 0 {
//             _st := *st_index
//             v = _st[len(_st)-1]
//             *this.st[index] = _st[:len(_st)-1]
//         }
//         if len(*this.st[index]) == 0 {
//             delete(this.st, index)
//         }
//     }
//     return v
// }

type DinnerPlates struct {
    capacity int //栈容量
    stacks [][]int //所有栈
    idx hp //最小堆: 入栈左边最小序号, 出栈右边最大序号
    //void Push(int val) - 将给出的正整数 val 推入 从左往右第一个 没有满的栈
    //int Pop() - 返回 从右往左第一个 非空栈顶部的值，并将其从栈中删除；如果所有的栈都是空的，请返回 -1
    //int PopAtStack(int index) - 返回编号 index 的栈顶部的值，并将其从栈中删除
}

func Constructor(capacity int) DinnerPlates {
    return DinnerPlates{capacity: capacity}
}

func (d *DinnerPlates) Push(val int) {
    if d.idx.Len() > 0 && d.idx.IntSlice[0] >= len(d.stacks) {
        d.idx = hp{} // 堆中都是越界下标，直接清空
    }
    if d.idx.Len() == 0 { // 所有栈都是满的
        d.stacks = append(d.stacks, []int{val}) // 添加一个新的栈
        if d.capacity > 1 { // 新的栈没有满
            heap.Push(&d.idx, len(d.stacks)-1) // 入堆
        }
    } else { // 还有未满栈
        i := d.idx.IntSlice[0]
        d.stacks[i] = append(d.stacks[i], val) // 入栈
        if len(d.stacks[i]) == d.capacity { // 栈满了
            heap.Pop(&d.idx) // 从堆中去掉
        }
    }
}

func (d *DinnerPlates) Pop() int {
    // 等价为 popAtStack 最后一个非空栈
    return d.PopAtStack(len(d.stacks) - 1)
}

func (d *DinnerPlates) PopAtStack(index int) int {
    if index < 0 || index >= len(d.stacks) || len(d.stacks[index]) == 0 {
        return -1 // 非法操作
    }
    if len(d.stacks[index]) == d.capacity { // 满栈
        heap.Push(&d.idx, index) // 元素出栈后，栈就不满了，把下标入堆
    }
    bk := len(d.stacks[index]) - 1
    val := d.stacks[index][bk]
    d.stacks[index] = d.stacks[index][:bk]
    for len(d.stacks) > 0 && len(d.stacks[len(d.stacks)-1]) == 0 {
        d.stacks = d.stacks[:len(d.stacks)-1] // 去掉末尾的空栈（懒删除，堆中下标在 push 时处理）
    }
    return val
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func main() {
    // DinnerPlates D = DinnerPlates(2);  // Initialize with capacity = 2
    obj := Constructor(2)
    fmt.Println(obj)
    // D.push(1);
    obj.Push(1)
    fmt.Println(obj)
    // D.push(2);
    obj.Push(2)
    fmt.Println(obj)
    // D.push(3);
    obj.Push(3)
    fmt.Println(obj)
    // D.push(4);
    obj.Push(4)
    fmt.Println(obj)
    // D.push(5);         // The stacks are now:  2  4
    //                                            1  3  5
    //                                            ﹈ ﹈ ﹈
    obj.Push(5)
    fmt.Println(obj)
    // D.popAtStack(0);   // Returns 2.  The stacks are now:     4
    //                                                        1  3  5
    //                                                        ﹈ ﹈ ﹈
    fmt.Println(obj.PopAtStack(0)) // 2
    fmt.Println(obj)
    // D.push(20);        // The stacks are now: 20  4
    //                                            1  3  5
    //                                            ﹈ ﹈ ﹈
    obj.Push(20)
    fmt.Println(obj)
    // D.push(21);        // The stacks are now: 20  4 21
    //                                            1  3  5
    //                                            ﹈ ﹈ ﹈
    obj.Push(21)
    fmt.Println(obj)
    // D.popAtStack(0);   // Returns 20.  The stacks are now:     4 21
    //                                                         1  3  5
    //                                                         ﹈ ﹈ ﹈
    fmt.Println(obj.PopAtStack(0)) // 20
    fmt.Println(obj)
    // D.popAtStack(2);   // Returns 21.  The stacks are now:     4
    //                                                         1  3  5
    //                                                         ﹈ ﹈ ﹈ 
    fmt.Println(obj.PopAtStack(2)) // 21
    fmt.Println(obj)
    // D.pop()            // Returns 5.  The stacks are now:      4
    //                                                         1  3 
    //                                                         ﹈ ﹈  
    fmt.Println(obj.Pop()) // 5
    fmt.Println(obj)
    // D.pop()            // Returns 4.  The stacks are now:   1  3 
    //                                                         ﹈ ﹈   
    fmt.Println(obj.Pop()) // 4
    fmt.Println(obj)
    // D.pop()            // Returns 3.  The stacks are now:   1 
    //                                                         ﹈   
    fmt.Println(obj.Pop()) // 3
    fmt.Println(obj)
    // D.pop()            // Returns 1.  There are no stacks.
    fmt.Println(obj.Pop()) // 1
    fmt.Println(obj)
    // D.pop()            // Returns -1.  There are still no stacks.
    fmt.Println(obj.Pop()) // -1
    fmt.Println(obj)
}