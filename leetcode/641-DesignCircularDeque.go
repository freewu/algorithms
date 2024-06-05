package main

// 641. Design Circular Deque
// Design your implementation of the circular double-ended queue (deque).
// Implement the MyCircularDeque class:
//     MyCircularDeque(int k) Initializes the deque with a maximum size of k.
//     boolean insertFront() Adds an item at the front of Deque. Returns true if the operation is successful, or false otherwise.
//     boolean insertLast() Adds an item at the rear of Deque. Returns true if the operation is successful, or false otherwise.
//     boolean deleteFront() Deletes an item from the front of Deque. Returns true if the operation is successful, or false otherwise.
//     boolean deleteLast() Deletes an item from the rear of Deque. Returns true if the operation is successful, or false otherwise.
//     int getFront() Returns the front item from the Deque. Returns -1 if the deque is empty.
//     int getRear() Returns the last item from Deque. Returns -1 if the deque is empty.
//     boolean isEmpty() Returns true if the deque is empty, or false otherwise.
//     boolean isFull() Returns true if the deque is full, or false otherwise.
 
// Example 1:
// Input
// ["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"]
// [[3], [1], [2], [3], [4], [], [], [], [4], []]
// Output
// [null, true, true, true, false, 2, true, true, true, 4]
// Explanation
// MyCircularDeque myCircularDeque = new MyCircularDeque(3);
// myCircularDeque.insertLast(1);  // return True
// myCircularDeque.insertLast(2);  // return True
// myCircularDeque.insertFront(3); // return True
// myCircularDeque.insertFront(4); // return False, the queue is full.
// myCircularDeque.getRear();      // return 2
// myCircularDeque.isFull();       // return True
// myCircularDeque.deleteLast();   // return True
// myCircularDeque.insertFront(4); // return True
// myCircularDeque.getFront();     // return 4
 
// Constraints:
//     1 <= k <= 1000
//     0 <= value <= 1000
//     At most 2000 calls will be made to insertFront, insertLast, deleteFront, deleteLast, getFront, getRear, isEmpty, isFull.

import "fmt"

type MyCircularDeque struct {
    queue []int // the Deque
    cap int // the lenght of the Deque
}

func Constructor(len int) MyCircularDeque {
    return MyCircularDeque{[]int{}, len}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
    // If the lenght of the deque is not equal to k we can add to the front of the array and return true
    if this.IsFull() { 
        return false 
    }
    this.queue = append(this.queue[:0], append([]int{ value }, this.queue[0:]...)...)
    return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
    // If the lenght of the deque is not equal to k we can add to the back of the array and return true
    if this.IsFull() { 
        return false 
    } 
    this.queue = append(this.queue, value)
    return true
}

func (this *MyCircularDeque) DeleteFront() bool {
    // If the lenght of the deque is not equal to 0 we can remove from the front of the array and return true
    if this.IsEmpty() { 
        return false 
    }
    this.queue = this.queue[1:]
    return true
}

func (this *MyCircularDeque) DeleteLast() bool {
    // If the lenght of the deque is not equal to 0 we can remove from the end of the array and return true
    if this.IsEmpty() {
        return false
    }
    this.queue = this.queue[:len(this.queue) - 1]
    return true
}

func (this *MyCircularDeque) GetFront() int {
    // If the lenght of the deque is not equal to 0 we can return the first element else return -1
    if this.IsEmpty() { 
        return -1 
    }
    return this.queue[0]
}

func (this *MyCircularDeque) GetRear() int {
    // If the lenght of the deque is not equal to 0 we can return the last element else return -1
    if this.IsEmpty() { 
        return -1
    }
    return this.queue[len(this.queue) - 1]
}

func (this *MyCircularDeque) IsEmpty() bool {
    return len(this.queue) == 0
}

func (this *MyCircularDeque) IsFull() bool {
    return len(this.queue) == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func main() {
    // MyCircularDeque myCircularDeque = new MyCircularDeque(3);
    obj := Constructor(3)
    fmt.Println(obj) // 
    fmt.Println(obj.IsEmpty()) // true
    // myCircularDeque.insertLast(1);  // return True
    fmt.Println(obj.InsertLast(1)) // true
    fmt.Println(obj) // 
    // myCircularDeque.insertLast(2);  // return True
    fmt.Println(obj.InsertLast(2)) // true
    fmt.Println(obj) // 
    // myCircularDeque.insertFront(3); // return True
    fmt.Println(obj.InsertLast(3)) // true
    fmt.Println(obj) // 
    // myCircularDeque.insertFront(4); // return False, the queue is full.
    fmt.Println(obj.InsertLast(4)) // false
    fmt.Println(obj) // 
    // myCircularDeque.getRear();      // return 2
    fmt.Println(obj.GetRear()) // 2
    // myCircularDeque.isFull();       // return True
    fmt.Println(obj.IsFull()) // True
    // myCircularDeque.deleteLast();   // return True
    fmt.Println(obj.DeleteLast()) // True
    fmt.Println(obj) // 
    // myCircularDeque.insertFront(4); // return True
    fmt.Println(obj.InsertFront(4)) // true
    fmt.Println(obj) // 
    // myCircularDeque.getFront();     // return 4
    fmt.Println(obj.GetFront()) // 4
}