package main

// LCR 125. 图书整理 II
// 读者来到图书馆排队借还书，图书管理员使用两个书车来完成整理借还书的任务。
// 书车中的书从下往上叠加存放，图书管理员每次只能拿取书车顶部的书。排队的读者会有两种操作：
//     push(bookID)：把借阅的书籍还到图书馆。
//     pop()：从图书馆中借出书籍。

// 为了保持图书的顺序，图书管理员每次取出供读者借阅的书籍是 最早 归还到图书馆的书籍。你需要返回 每次读者借出书的值 。
// 如果没有归还的书可以取出，返回 -1 。

// 示例 1：
// 输入：
// ["BookQueue", "push", "push", "pop"]
// [[], [1], [2], []]
// 输出：[null,null,null,1]
// 解释：
// MyQueue myQueue = new MyQueue();
// myQueue.push(1); // queue is: [1]
// myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
// myQueue.pop(); // return 1, queue is [2]
 
// 提示：
//     1 <= bookID <= 10000
//     最多会对 push、pop 进行 10000 次调用

import "fmt"

// type CQueue struct {
//     data []int
// }

// func Constructor() CQueue {
//     return CQueue{data: []int{} }
// }

// func (this *CQueue) AppendTail(value int)  {
//     this.data = append(this.data, value)
// }

// func (this *CQueue) DeleteHead() int {
//     if len(this.data) == 0 {
//         return -1
//     }
//     v := this.data[0]
//     this.data = this.data[1:]
//     return v
// }

type CQueue struct {
    stack1, stack2 []int
}

func Constructor() CQueue {
    return CQueue{[]int{},[]int{}}
}

func (this *CQueue) AppendTail(value int)  {
    this.stack1 = append(this.stack1,value)
}

func (this *CQueue) DeleteHead() int {
    if len(this.stack1) == 0 && len(this.stack2) == 0 {
        return -1
    }
    if len(this.stack2) > 0 {
        v := this.stack2[len(this.stack2) - 1]
        this.stack2 = this.stack2[:len(this.stack2)-1]
        return v
    }
    for len(this.stack1) > 0 {
        t := this.stack1[len(this.stack1)-1]
        this.stack1 = this.stack1[:len(this.stack1) - 1]
        this.stack2 = append(this.stack2, t )
    }
    v := this.stack2[len(this.stack2)-1]
    this.stack2 = this.stack2[:len(this.stack2)-1]
    return v
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
    // MyQueue myQueue = new MyQueue();
    obj := Constructor()
    fmt.Println(obj)
    // myQueue.push(1); // queue is: [1]
    obj.AppendTail(1)
    fmt.Println(obj)
    // myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
    obj.AppendTail(2)
    fmt.Println(obj)
    // myQueue.pop(); // return 1, queue is [2]
    fmt.Println(obj.DeleteHead()) // 1
    fmt.Println(obj)
}