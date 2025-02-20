package main

// 面试题 16.25. LRU Cache LCCI
// Design and build a "least recently used" cache, which evicts the least recently used item. 
// The cache should map from keys to values (allowing you to insert and retrieve a value associ­ated with a particular key) and be initialized with a max size. 
// When it is full, it should evict the least recently used item.

// You should implement following operations:  get and put.
//     Get a value by key: 
//         get(key) - If key is in the cache, return the value, otherwise return -1.
//     Write a key-value pair to the cache: 
//         put(key, value) - If the key is not in the cache, then write its value to the cache. 
//                           Evict the least recently used item before writing if necessary.

// Example:
// LRUCache cache = new LRUCache( 2 /* capacity */ );
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4

import "fmt"

// type LRUCache struct {
//     ch map[int]int
//     stack []int
//     capacity int
// }

// func Constructor(capacity int) LRUCache {
//     return LRUCache {
//         capacity : capacity,
//         ch : map[int]int{},
//         stack : []int{},
//     }
// }

// func (this *LRUCache) Get(key int) int {
//     if _, ok := this.ch[key]; !ok { return -1 }
//     res := this.ch[key]
//     for k, v := range this.stack {
//         if v == key {
//           stack1 := this.stack[:k]
//           stack2 := this.stack[k+1:]
//           this.stack = append(stack1, stack2...)
//           this.stack = append(this.stack, key)  
//         }
//     }
//     return res
// }

// func (this *LRUCache) Put(key int, value int)  {
//     if _, ok := this.ch[key]; ok {
//             for k, v := range this.stack {
//                 if v == key {
//                     stack1 := this.stack[:k]
//                     stack2 := this.stack[k+1:]
//                     this.stack = append(stack1, stack2...)
//                }
//             }
//     } else {
//         if len(this.ch) == this.capacity {
//             tmp := this.stack[0]
//             this.stack = this.stack[1:]
//             delete(this.ch, tmp)
//         }
//     }
//     this.ch[key] = value
//     this.stack = append(this.stack, key)
// }

type LRUCache struct {
    Head *Node
    Tail *Node
    Cap int
    Length int
    NodeKeyMap map[int]*Node
}

type Node struct{
    Key int
    Value int
    Next *Node
    Pre *Node
}

func Constructor(capacity int) LRUCache {
    head := &Node{}
    tail := &Node{}
    head.Next = tail
    tail.Pre = head
    return LRUCache{ Head: head, Tail: tail, Cap: capacity,  NodeKeyMap: make(map[int]*Node) }
}

func (this *LRUCache) Get(key int) int {
    node,ok := this.NodeKeyMap[key]
    if !ok { return -1 }
    res := node.Value
    if this.Head.Next == node { return res }
    current := this.Head.Next
    node.Pre.Next = node.Next
    node.Next.Pre = node.Pre
    this.Head.Next = node
    node.Pre = this.Head
    node.Next = current
    current.Pre = node
    return res
}

func (this *LRUCache) Put(key int, value int)  {
    node,ok := this.NodeKeyMap[key]
    if !ok {
        node = &Node{ Key:key, Value: value }
        this.NodeKeyMap[key] = node
        current := this.Head.Next
        this.Head.Next = node
        node.Pre = this.Head
        node.Next = current
        current.Pre = node
        if this.Length < this.Cap {
            this.Length++
            return 
        }
        deleted := this.Tail.Pre
        deleted.Pre.Next = this.Tail
        this.Tail.Pre = deleted.Pre
        delete(this.NodeKeyMap, deleted.Key)
    } else {
        node.Value = value
        if this.Head.Next == node { return }
        current := this.Head.Next
        node.Pre.Next = node.Next
        node.Next.Pre = node.Pre
        this.Head.Next = node
        node.Pre = this.Head
        node.Next = current
        current.Pre = node
    }
    return 
}

func main() {
    // LRUCache cache = new LRUCache( 2 /* capacity */ );
    obj := Constructor(2)
    fmt.Println(obj)
    // cache.put(1, 1);
    obj.Put(1, 1)
    fmt.Println(obj)
    // cache.put(2, 2);
    obj.Put(2, 2)
    fmt.Println(obj)
    // cache.get(1);       // returns 1
    fmt.Println(obj.Get(1)) // 1
    // cache.put(3, 3);    // evicts key 2
    obj.Put(3, 3)
    fmt.Println(obj)
    // cache.get(2);       // returns -1 (not found)
    fmt.Println(obj.Get(2)) // -1
    // cache.put(4, 4);    // evicts key 1
    obj.Put(4, 4)
    fmt.Println(obj)
    // cache.get(1);       // returns -1 (not found)
    fmt.Println(obj.Get(1)) // -1
    // cache.get(3);       // returns 3
    fmt.Println(obj.Get(3)) // 3
    // cache.get(4);       // returns 4
    fmt.Println(obj.Get(4)) // 4
}