package main

// 707. Design Linked List
// Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
// A node in a singly linked list should have two attributes: val and next. 
// val is the value of the current node, and next is a pointer/reference to the next node.
// If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. 
// Assume all nodes in the linked list are 0-indexed.

// Implement the MyLinkedList class:
//     MyLinkedList() Initializes the MyLinkedList object.
//     int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
//     void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
//     void addAtTail(int val) Append a node of value val as the last element of the linked list.
//     void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
//     void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.
    

// Example 1:
// Input
// ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
// [[], [1], [3], [1, 2], [1], [1], [1]]
// Output
// [null, null, null, null, 2, null, 3]
// Explanation
// MyLinkedList myLinkedList = new MyLinkedList();
// myLinkedList.addAtHead(1);
// myLinkedList.addAtTail(3);
// myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
// myLinkedList.get(1);              // return 2
// myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
// myLinkedList.get(1);              // return 3

// Constraints:
//     0 <= index, val <= 1000
//     Please do not use the built-in LinkedList library.
//     At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and deleteAtIndex.

import "fmt"

// 单链表
type node struct {
    val int
    next *node
}

type MyLinkedList struct {
    head *node
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
    cur := this.head
    for i := 0; cur != nil ; i++ {
        if i == index { return cur.val }
        cur = cur.next
    }
    return -1
}

func (this *MyLinkedList) AddAtHead(val int)  {
    if this.head == nil {
        this.head = &node{val: val}
        return
    }
    this.head = &node{
        val: val,
        next: this.head,
    }
}

func (this *MyLinkedList) AddAtTail(val int)  {
    if this.head == nil { 
        this.AddAtHead(val)
        return
    }
    // 到链表尾部
    cur := this.head
    for cur.next != nil { cur = cur.next }
    cur.next = &node{val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index <= 0 {
        this.AddAtHead(val)
        return
    }
    for i, cur := 0, this.head; cur != nil; i, cur = i+1, cur.next {
        // 找到指定位置 插入
        if i == index - 1 {
            cur.next = &node{
                val:  val,
                next: cur.next,
            }
            break
        }
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index == 0 && this.head != nil {
        this.head = this.head.next
        return
    }
    for i, cur := 0, this.head; cur != nil; i, cur = i+1, cur.next {
        // 找到指定位置删除节点
        if i == index-1 && cur.next != nil {
            cur.next = cur.next.next
            break
        }
    }
}

// 双链表
type Node struct {
    val        int
    next, prev *Node
}

type MyLinkedList1 struct {
    head, tail *Node
    size       int
}

func Constructor1() MyLinkedList1 {
    head := &Node{}
    tail := &Node{}
    head.next = tail
    tail.prev = head
    return MyLinkedList1{head, tail, 0}
}


func (this *MyLinkedList1) Get(index int) int {
    if index < 0 || index > this.size - 1 {
        return -1
    }
    var node *Node
    node = this.head
    for i := 0; i <= index; i++ {
        node = node.next;
    }
    return node.val
}


func (this *MyLinkedList1) AddAtHead(val int)  {
    node := &Node{val, this.head.next, this.head}
    this.head.next.prev = node
    this.head.next = node
    this.size += 1
}


func (this *MyLinkedList1) AddAtTail(val int)  {
    node := &Node{val, this.tail, this.tail.prev}
    this.tail.prev.next = node
    this.tail.prev = node
    this.size += 1
}


func (this *MyLinkedList1) AddAtIndex(index int, val int)  {
    if index > this.size {
        return
    }
    if index == this.size {
        this.AddAtTail(val)
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    if index < 0 {
        index = 0
    }
    var p *Node
    p = this.head
    for i := 0; i <= index; i++ {
        p = p.next
    }
    node := &Node{val, p, p.prev}
    p.prev.next = node
    p.prev = node
    this.size += 1
}


func (this *MyLinkedList1) DeleteAtIndex(index int)  {
    if index < 0 || index > this.size - 1 {
        return
    }
    var p *Node
    p = this.head
    for i := 0; i <= index; i++ {
        p = p.next
    }
    p.next.prev = p.prev
    p.prev.next = p.next
    p.prev = nil
    p.next = nil
    this.size -= 1
}

// 打印链表
func printListNode(l *node) {
    if nil == l {
        return
    }
    for {
        if nil == l.next {
            fmt.Print(l.val)
            break
        } else {
            fmt.Print(l.val, " -> ")
        }
        l = l.next
    }
    fmt.Println()
}

// 打印链表
func printListNode1(l *Node) {
    if nil == l {
        return
    }
    for {
        if nil == l.next {
            fmt.Print(l.val)
            break
        } else {
            fmt.Print(l.val, " -> ")
        }
        l = l.next
    }
    fmt.Println()
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
    // MyLinkedList myLinkedList = new MyLinkedList();
    obj := Constructor()
    // myLinkedList.addAtHead(1);
    obj.AddAtHead(1)
    printListNode(obj.head) // 1
    // myLinkedList.addAtTail(3);
    obj.AddAtTail(3)
    printListNode(obj.head) // 1 -> 3
    // myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
    obj.AddAtIndex(1,2)
    printListNode(obj.head) // 1 -> 2 -> 3
    // myLinkedList.get(1);              // return 2
    fmt.Println(obj.Get(1)) // 2
    // myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
    obj.DeleteAtIndex(1)
    printListNode(obj.head) // 1 -> 3
    // myLinkedList.get(1);              // return 3
    fmt.Println(obj.Get(1)) // 3

    obj1 := Constructor1()
    // myLinkedList.addAtHead(1);
    obj1.AddAtHead(1)
    printListNode1(obj1.head) // 1
    // myLinkedList.addAtTail(3);
    obj1.AddAtTail(3)
    printListNode1(obj1.head) // 1 -> 3
    // myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
    obj1.AddAtIndex(1,2)
    printListNode1(obj1.head) // 1 -> 2 -> 3
    // myLinkedList.get(1);              // return 2
    fmt.Println(obj1.Get(1)) // 2
    // myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
    obj1.DeleteAtIndex(1)
    printListNode1(obj1.head) // 1 -> 3
    // myLinkedList.get(1);              // return 3
    fmt.Println(obj1.Get(1)) // 3
}