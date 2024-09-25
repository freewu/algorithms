package main

// 1670. Design Front Middle Back Queue
// Design a queue that supports push and pop operations in the front, middle, and back.

// Implement the FrontMiddleBack class:
//     FrontMiddleBack() Initializes the queue.
//     void pushFront(int val) Adds val to the front of the queue.
//     void pushMiddle(int val) Adds val to the middle of the queue.
//     void pushBack(int val) Adds val to the back of the queue.
//     int popFront() Removes the front element of the queue and returns it. If the queue is empty, return -1.
//     int popMiddle() Removes the middle element of the queue and returns it. If the queue is empty, return -1.
//     int popBack() Removes the back element of the queue and returns it. If the queue is empty, return -1.

// Notice that when there are two middle position choices, the operation is performed on the frontmost middle position choice. 
// For example:
//     Pushing 6 into the middle of [1, 2, 3, 4, 5] results in [1, 2, 6, 3, 4, 5].
//     Popping the middle from [1, 2, 3, 4, 5, 6] returns 3 and results in [1, 2, 4, 5, 6].

// Example 1:
// Input:
// ["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
// [[], [1], [2], [3], [4], [], [], [], [], []]
// Output:
// [null, null, null, null, null, 1, 3, 4, 2, -1]
// Explanation:
// FrontMiddleBackQueue q = new FrontMiddleBackQueue();
// q.pushFront(1);   // [1]
// q.pushBack(2);    // [1, 2]
// q.pushMiddle(3);  // [1, 3, 2]
// q.pushMiddle(4);  // [1, 4, 3, 2]
// q.popFront();     // return 1 -> [4, 3, 2]
// q.popMiddle();    // return 3 -> [4, 2]
// q.popMiddle();    // return 4 -> [2]
// q.popBack();      // return 2 -> []
// q.popFront();     // return -1 -> [] (The queue is empty)

// Constraints:
//     1 <= val <= 10^9
//     At most 1000 calls will be made to pushFront, pushMiddle, pushBack, popFront, popMiddle, and popBack.

import "fmt"
import "math"
import "container/list"

type FrontMiddleBackQueue struct {
    data []int
}

func Constructor() FrontMiddleBackQueue {
    return FrontMiddleBackQueue{ []int{} }
}

func (this *FrontMiddleBackQueue) PushFront(val int)  {
    this.data = append([]int{val}, this.data...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
    middle := len(this.data) / 2
    if len(this.data) == middle {
        this.data = append(this.data, val)
        return
    }
    this.data = append(this.data[:middle+1], this.data[middle:]...)
    this.data[middle] = val
}

func (this *FrontMiddleBackQueue) PushBack(val int)  {
    this.data = append(this.data, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
    if len(this.data) == 0 { return -1 }
    res := this.data[0]
    this.data = this.data[1:]
    return res
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
    if len(this.data) == 0 { return -1 }
    middle := math.Ceil(float64(len(this.data)) / 2) - 1
    res := this.data[int(middle)]
    this.data = append(this.data[:int(middle)], this.data[int(middle)+1:]...)
    return res
}

func (this *FrontMiddleBackQueue) PopBack() int {
    if len(this.data) == 0 { return -1 }
    res := this.data[len(this.data)-1]
    this.data = this.data[:len(this.data)-1]
    return res
}


type FrontMiddleBackQueue1 struct {
    left  *list.List
    right *list.List
}

// 用两个列表表示队列的左右两部分，一遍从中间操作元素
// 如果是奇数个元素，维护左边少右边多，所以：
// 1、如果有偶数个元素时，pushMiddle 优先向右边添加
// 2、如果有奇数个元素时，popMiddle 优先从右边删除
// 3、如果只有 1 个元素，popFront 的时候，要去右边删除
// 要把以上三个特点写到代码里，才能保证细节不出错

// 维护左边少右边多的状态，每次增删元素之后都要执行一次
func (q *FrontMiddleBackQueue1) balance() {
    // 右边最多比左边多一个元素
    if q.right.Len() > q.left.Len()+1 {
        // 右边多，匀一个给左边
        q.left.PushBack(q.right.Remove(q.right.Front()))
    }
    if q.left.Len() > q.right.Len() {
        // 左边多，匀一个给右边
        q.right.PushFront(q.left.Remove(q.left.Back()))
    }
}

func Constructor1() FrontMiddleBackQueue1 {
    return FrontMiddleBackQueue1{
        left:  list.New(),
        right: list.New(),
    }
}

func (q *FrontMiddleBackQueue1) PushFront(val int) {
    q.left.PushFront(val)
    q.balance()
}

func (q *FrontMiddleBackQueue1) PushMiddle(val int) {
    if (q.left.Len()+q.right.Len())%2 == 0 {
        // 如果有偶数个元素时，pushMiddle 优先向右边添加
        q.right.PushFront(val)
    } else {
        q.left.PushBack(val)
    }
    q.balance()
}

func (q *FrontMiddleBackQueue1) PushBack(val int) {
    q.right.PushBack(val)
    q.balance()
}

func (q *FrontMiddleBackQueue1) PopFront() int {
    if q.left.Len()+q.right.Len() == 0 {
        return -1
    }
    if q.left.Len()+q.right.Len() == 1 {
        // 如果只有 1 个元素，popFront 的时候，要去右边删除
        return q.right.Remove(q.right.Front()).(int)
    }
    e := q.left.Remove(q.left.Front()).(int)
    q.balance()
    return e
}

func (q *FrontMiddleBackQueue1) PopMiddle() int {
    if q.left.Len()+q.right.Len() == 0 {
        return -1
    }
    var e int
    if (q.left.Len()+q.right.Len())%2 == 0 {
        e = q.left.Remove(q.left.Back()).(int)
    } else {
        // 如果有奇数个元素时，popMiddle 优先从右边删除
        e = q.right.Remove(q.right.Front()).(int)
    }
    q.balance()
    return e
}

func (q *FrontMiddleBackQueue1) PopBack() int {
    if q.left.Len()+q.right.Len() == 0 {
        return -1
    }
    e := q.right.Remove(q.right.Back()).(int)
    q.balance()
    return e
}

func (q *FrontMiddleBackQueue1) Size() int {
    return q.left.Len() + q.right.Len()
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */

func main() {
    // FrontMiddleBackQueue q = new FrontMiddleBackQueue();
    obj := Constructor()
    fmt.Println(obj)
    // q.pushFront(1);   // [1]
    obj.PushFront(1)
    fmt.Println(obj) // [1]
    // q.pushBack(2);    // [1, 2]
    obj.PushBack(2)
    fmt.Println(obj) // [1, 2]
    // q.pushMiddle(3);  // [1, 3, 2]
    obj.PushMiddle(2)
    fmt.Println(obj) // [1, 3, 2]
    // q.pushMiddle(4);  // [1, 4, 3, 2]
    obj.PushMiddle(4)
    fmt.Println(obj) // [1, 4, 3, 2]
    // q.popFront();     // return 1 -> [4, 3, 2]
    fmt.Println(obj.PopFront()) // 1
    fmt.Println(obj) // [1, 4, 3, 2]
    // q.popMiddle();    // return 3 -> [4, 2]
    fmt.Println(obj.PopMiddle()) // 3
    fmt.Println(obj) // [4, 2]
    // q.popMiddle();    // return 4 -> [2]
    fmt.Println(obj.PopMiddle()) // 4
    fmt.Println(obj) // []
    // q.popBack();      // return 2 -> []
    fmt.Println(obj.PopBack()) // 2
    fmt.Println(obj) // []
    // q.popFront();     // return -1 -> [] (The queue is empty)
    fmt.Println(obj.PopFront()) // -1
    fmt.Println(obj) // []

    // FrontMiddleBackQueue q = new FrontMiddleBackQueue();
    obj1 := Constructor1()
    fmt.Println(obj1)
    // q.pushFront(1);   // [1]
    obj1.PushFront(1)
    fmt.Println(obj1) // [1]
    // q.pushBack(2);    // [1, 2]
    obj1.PushBack(2)
    fmt.Println(obj1) // [1, 2]
    // q.pushMiddle(3);  // [1, 3, 2]
    obj1.PushMiddle(2)
    fmt.Println(obj1) // [1, 3, 2]
    // q.pushMiddle(4);  // [1, 4, 3, 2]
    obj1.PushMiddle(4)
    fmt.Println(obj1) // [1, 4, 3, 2]
    // q.popFront();     // return 1 -> [4, 3, 2]
    fmt.Println(obj1.PopFront()) // 1
    fmt.Println(obj1) // [1, 4, 3, 2]
    // q.popMiddle();    // return 3 -> [4, 2]
    fmt.Println(obj1.PopMiddle()) // 3
    fmt.Println(obj1) // [4, 2]
    // q.popMiddle();    // return 4 -> [2]
    fmt.Println(obj1.PopMiddle()) // 4
    fmt.Println(obj1) // []
    // q.popBack();      // return 2 -> []
    fmt.Println(obj1.PopBack()) // 2
    fmt.Println(obj1) // []
    // q.popFront();     // return -1 -> [] (The queue is empty)
    fmt.Println(obj1.PopFront()) // -1
    fmt.Println(obj1) // []
}