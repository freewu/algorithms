package main

// LCR 031. LRU 缓存
// 运用所掌握的数据结构，设计和实现一个  LRU (Least Recently Used，最近最少使用) 缓存机制 。
// 实现 LRUCache 类：
//     LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
//     int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//     void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

// 示例：
// 输入
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
// 解释
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存是 {1=1}
// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1
// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1 (未找到)
// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1 (未找到)
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4

// 提示：
//     1 <= capacity <= 3000
//     0 <= key <= 10000
//     0 <= value <= 10^5
//     最多调用 2 * 10^5 次 get 和 put

// 进阶：是否可以在 O(1) 时间复杂度内完成这两种操作？

import "fmt"
import "container/list"

type LRUCache struct {
    cache map[int]*list.Element
    linklist *list.List
    capacity int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        cache: make(map[int]*list.Element, capacity),
        linklist: list.New(),
        capacity: capacity,
    }
}

// 0 is key and 1 is value
func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; !ok {
        return -1
    }
    elem := this.cache[key]
    this.linklist.MoveToFront(elem)
    return elem.Value.([]int)[1]
}

func (this *LRUCache) Put(key int, value int)  {
    // if capacity reached
    if elem, ok := this.cache[key]; ok {
        this.linklist.Remove(elem)
        newelem := this.linklist.PushFront([]int{key, value})
        this.cache[key] = newelem
        return
    }
    if len(this.cache) == this.capacity {
        elem := this.linklist.Back()
        v := this.linklist.Remove(elem)
        delete(this.cache, v.([]int)[0])
    }
    newelem := this.linklist.PushFront([]int{key, value})
    this.cache[key] = newelem
}

// 添加头尾节点指针，并且在初始化的时候把它们连接起来，可以大大简化判断，代码更简洁
// 因为访问时涉及到挪动操作，要在 o(1) 的时间复杂度实现，只能用链表
type LRUCache1 struct {
    m map[int]*Node
    capacity int
    count int
    // 头尾指针不存数据，单向链接到头尾节点
    head *Node
    tail *Node
}

type Node struct {
    key int
    val int
    prev *Node
    next *Node
}

func Constructor1(capacity int) LRUCache1 {
    l := LRUCache1{
        m: make(map[int]*Node, capacity),
        capacity: capacity,
        head:&Node{},
        tail:&Node{},
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}


func (this *LRUCache1) Get(key int) int {
    n, ok := this.m[key]
    if !ok {
        return -1
    }
    this.RemoveNode(n)
    this.AddToHead(n)
    return n.val
}


func (this *LRUCache1) Put(key int, value int)  {
    n, ok := this.m[key]
    if ok {
        n.val = value
        this.RemoveNode(n)
        this.AddToHead(n)
        return
    }

    if this.count == this.capacity {
        tail := this.RemoveTail()
        delete(this.m, tail.key)
        this.count--
    }
    this.count++
    n = &Node{key:key, val:value}
    this.AddToHead(n)
    this.m[key] = n
}

func (this *LRUCache1) RemoveNode(n *Node) {
    n.prev.next = n.next
    n.next.prev = n.prev
    n.next = nil
    n.prev = nil
}

func (this *LRUCache1) RemoveTail() *Node {
    tail := this.tail.prev
    this.RemoveNode(tail)
    return tail
}

func (this *LRUCache1) AddToHead(n *Node) {
    // 改变原先头结点的前驱节点为新的节点
    n.next = this.head.next
    this.head.next.prev = n
    // 改变新头结点的前驱节点为 head 指针
    this.head.next = n
    n.prev = this.head
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
    // LRUCache lRUCache = new LRUCache(2);
    obj := Constructor(2)
    fmt.Println(obj)
    // lRUCache.put(1, 1); // cache is {1=1}
    obj.Put(1,1)
    fmt.Println(obj)
    // lRUCache.put(2, 2); // cache is {1=1, 2=2}
    obj.Put(2,2)
    fmt.Println(obj)
    // lRUCache.get(1);    // return 1
    fmt.Println(obj.Get(1)) // 1
    // lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
    obj.Put(3,3)
    fmt.Println(obj)
    // lRUCache.get(2);    // returns -1 (not found)
    fmt.Println(obj.Get(2)) // -1
    // lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
    obj.Put(4,4)
    fmt.Println(obj)
    // lRUCache.get(1);    // return -1 (not found)
    fmt.Println(obj.Get(1)) // -1
    // lRUCache.get(3);    // return 3
    fmt.Println(obj.Get(3)) // 3
    // lRUCache.get(4);    // return 4
    fmt.Println(obj.Get(4)) // 4

    // LRUCache lRUCache = new LRUCache(2);
    obj1 := Constructor(2)
    fmt.Println(obj1)
    // lRUCache.put(1, 1); // cache is {1=1}
    obj1.Put(1,1)
    fmt.Println(obj1)
    // lRUCache.put(2, 2); // cache is {1=1, 2=2}
    obj1.Put(2,2)
    fmt.Println(obj1)
    // lRUCache.get(1);    // return 1
    fmt.Println(obj1.Get(1)) // 1
    // lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
    obj1.Put(3,3)
    fmt.Println(obj1)
    // lRUCache.get(2);    // returns -1 (not found)
    fmt.Println(obj1.Get(2)) // -1
    // lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
    obj1.Put(4,4)
    fmt.Println(obj1)
    // lRUCache.get(1);    // return -1 (not found)
    fmt.Println(obj1.Get(1)) // -1
    // lRUCache.get(3);    // return 3
    fmt.Println(obj1.Get(3)) // 3
    // lRUCache.get(4);    // return 4
    fmt.Println(obj1.Get(4)) // 4
}