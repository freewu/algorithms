package main 

// 146. LRU Cache
// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
// Implement the LRUCache class:
//     LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
//     int get(int key) Return the value of the key if the key exists, otherwise return -1.
//     void put(int key, int value) 
//         Update the value of the key if the key exists. 
//         Otherwise, add the key-value pair to the cache. 
//         If the number of keys exceeds the capacity from this operation, evict the least recently used key.

// The functions get and put must each run in O(1) average time complexity.

// Example 1:
// Input
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// Output
// [null, null, null, 1, null, -1, null, -1, 3, 4]
// Explanation
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // cache is {1=1}
// lRUCache.put(2, 2); // cache is {1=1, 2=2}
// lRUCache.get(1);    // return 1
// lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
// lRUCache.get(2);    // returns -1 (not found)
// lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
// lRUCache.get(1);    // return -1 (not found)
// lRUCache.get(3);    // return 3
// lRUCache.get(4);    // return 4
 
// Constraints:
//     1 <= capacity <= 3000
//     0 <= key <= 10^4
//     0 <= value <= 10^5
//     At most 2 * 10^5 calls will be made to get and put.

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