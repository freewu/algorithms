package main 

// 622. Design Circular Queue
// Design your implementation of the circular queue. 
// The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle, 
// and the last position is connected back to the first position to make a circle. 
// It is also called "Ring Buffer".

// One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. 
// In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. 
// But using the circular queue, we can use the space to store new values.

// Implement the MyCircularQueue class:
//     MyCircularQueue(k) Initializes the object with the size of the queue to be k.
//     int Front() Gets the front item from the queue. If the queue is empty, return -1.
//     int Rear() Gets the last item from the queue. If the queue is empty, return -1.
//     boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
//     boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
//     boolean isEmpty() Checks whether the circular queue is empty or not.
//     boolean isFull() Checks whether the circular queue is full or not.

// You must solve the problem without using the built-in queue data structure in your programming language. 

// Example 1:
// Input
// ["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
// [[3], [1], [2], [3], [4], [], [], [], [4], []]
// Output
// [null, true, true, true, false, 3, true, true, true, 4]
// Explanation
// MyCircularQueue myCircularQueue = new MyCircularQueue(3);
// myCircularQueue.enQueue(1); // return True
// myCircularQueue.enQueue(2); // return True
// myCircularQueue.enQueue(3); // return True
// myCircularQueue.enQueue(4); // return False
// myCircularQueue.Rear();     // return 3
// myCircularQueue.isFull();   // return True
// myCircularQueue.deQueue();  // return True
// myCircularQueue.enQueue(4); // return True
// myCircularQueue.Rear();     // return 4
 
// Constraints:
//     1 <= k <= 1000
//     0 <= value <= 1000
//     At most 3000 calls will be made to enQueue, deQueue, Front, Rear, isEmpty, and isFull.

import "fmt"

type MyCircularQueue struct {
	data    []int // 数据存放
	front   int
    rear    int
    size    int // 初始化队列长度
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size:  k,
		data: make([]int, k),
		front: 0,
		rear:  -1,
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
    // 队列满了,不能再加入队列了
	if q.IsFull() {
		return false
	}
	q.rear++
	q.data[q.rear % q.size] = value
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front++
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.front % q.size]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[q.rear % q.size]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.rear < q.front
}

func (q *MyCircularQueue) IsFull() bool {
	return q.rear - q.front == q.size - 1
}

func main() {
    // MyCircularQueue myCircularQueue = new MyCircularQueue(3);
    obj := Constructor(3)
    fmt.Println(obj) // {[0 0 0] 0 -1 3}
    // myCircularQueue.enQueue(1); // return True
    fmt.Println(obj.EnQueue(1)) // true
    fmt.Println(obj) // {[1 0 0] 0 0 3}
    // myCircularQueue.enQueue(2); // return True
    fmt.Println(obj.EnQueue(2)) // true
    fmt.Println(obj) // {[1 2 0] 0 1 3}
    // myCircularQueue.enQueue(3); // return True
    fmt.Println(obj.EnQueue(3)) // true
    fmt.Println(obj) // {[1 2 3] 0 2 3}
    // myCircularQueue.enQueue(4); // return False
    fmt.Println(obj.EnQueue(4)) // false
    fmt.Println(obj) // {[1 2 3] 0 2 3}
    // myCircularQueue.Rear();     // return 3
    fmt.Println(obj.Rear()) // 3
    fmt.Println(obj) // {[1 2 3] 0 2 3}
    // myCircularQueue.isFull();   // return True
    fmt.Println(obj.IsFull()) // true
    // myCircularQueue.deQueue();  // return True
    fmt.Println(obj.DeQueue()) // true
    fmt.Println(obj) // {[1 2 3] 1 2 3}
    // myCircularQueue.enQueue(4); // return True
    fmt.Println(obj.EnQueue(4)) // true
    fmt.Println(obj) // {[4 2 3] 1 3 3}
    // myCircularQueue.Rear();     // return 4
    fmt.Println(obj.Rear()) // 4
    fmt.Println(obj) // {[4 2 3] 1 3 3}
}