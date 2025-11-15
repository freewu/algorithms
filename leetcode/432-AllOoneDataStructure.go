package main

// 432. All O`one Data Structure
// Design a data structure to store the strings' count with the ability to return the strings with minimum and maximum counts.
// Implement the AllOne class:
//     AllOne() 
//         Initializes the object of the data structure.
//     inc(String key) 
//         Increments the count of the string key by 1. 
//         If key does not exist in the data structure, insert it with count 1.
//     dec(String key) 
//         Decrements the count of the string key by 1. 
//         If the count of key is 0 after the decrement, remove it from the data structure. 
//         It is guaranteed that key exists in the data structure before the decrement.
//     getMaxKey() 
//         Returns one of the keys with the maximal count. 
//         If no element exists, return an empty string "".
//     getMinKey() 
//         Returns one of the keys with the minimum count. 
//         If no element exists, return an empty string "".

// Note that each function must run in O(1) average time complexity.

// Example 1:
// Input
// ["AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey", "getMinKey"]
// [[], ["hello"], ["hello"], [], [], ["leet"], [], []]
// Output
// [null, null, null, "hello", "hello", null, "hello", "leet"]
// Explanation
// AllOne allOne = new AllOne();
// allOne.inc("hello");
// allOne.inc("hello");
// allOne.getMaxKey(); // return "hello"
// allOne.getMinKey(); // return "hello"
// allOne.inc("leet");
// allOne.getMaxKey(); // return "hello"
// allOne.getMinKey(); // return "leet"
 
// Constraints:
//     1 <= key.length <= 10
//     key consists of lowercase English letters.
//     It is guaranteed that for each call to dec, key is existing in the data structure.
//     At most 5 * 10^4 calls will be made to inc, dec, getMaxKey, and getMinKey.

import "fmt"

// // Wrong Answer 24 / 25 testcases passed
// type Node struct{
//     count int
//     str   string            // key, count键值对的双向链表结点
//     next *Node
//     last *Node 
// }

// type AllOne struct {
//     mp   map[string]*Node   // (key, 链表结点)的哈希表
//     head *Node              // 存储双链表头结点和尾结点
//     tail *Node              // 维护一个从头到尾count值单调减小的有序双链表
// }

// func Constructor() AllOne {
//     root := &Node{
//         count : 0x7fffffff, // 哑头结点 减少奇异点的判断
//                             // 判断链表为空 tail == head or head.next == nil
//     } 
//     return AllOne{
//         mp   : make(map[string]*Node),
//         head : root, 
//         tail : root,
//     }
// }

// func (this *AllOne) Inc(key string)  {
//     p, flag := this.mp[key]
//     if !flag {
//         p = &Node{ str : key, count : 1 }
//         // 双向链表尾部插入结点  
//         p.last = this.tail 
//         this.tail.next = p
//         this.tail = p
//         this.mp[key] = p
//     } else {
//         p.count ++
//         if p.last == this.head || p.last.count > p.count { // 若结点位于链表首部或链表仍然有序 直接inc退出
//             return
//         }
//         // 找到比结点值大的前一个结点
//         q := p
//         for q.count <= p.count {
//             q = q.last
//         }
//         if p == this.tail { // 在链表中删除结点标准操作
//             p.last.next = nil
//             this.tail = p.last
//         } else {
//             p.next.last = p.last
//             p.last.next = p.next
//         } 
//         // 将结点插入比结点值大的前一个结点之后 维持有序
//         q.next.last = p
//         p.next = q.next
//         p.last = q
//         q.next = p
//     } 
// }

// func (this *AllOne) Dec(key string) {
//     p, _ := this.mp[key]
//     p.count--
//     if p.count == 0 {
//         // 在链表中删除结点标准操作
//         delete(this.mp, key)
//         if p == this.tail {
//             p.last.next = nil
//             this.tail = p.last
//         } else {
//             p.next.last = p.last
//             p.last.next = p.next
//         }
//     } else {

//         if p == this.tail || p.next.count < p.count { // 若结点位于链表尾部或链表仍然有序 直接dec退出
//             return
//         }
//         q := p
//         for q != nil && q.count >= p.count { // 找到比结点值小的后一个结点
//             q = q.next
//         }
//         if q == nil { // 不存在这样的结点 说明结点位于链表尾部 直接退出
//             return 
//         }
//         // 将结点插入比结点值小的后一个结点之前 维持有序
//         p.next.last = p.last
//         p.last.next = p.next
//         q.last.next = p
//         p.last = q.last
//         p.next = q
//         q.last = p
//     }
// }

// // 维持有序双向链表后 max与min显然位于链表头与链表尾
// func (this *AllOne) GetMaxKey() string {
//     if this.head.next == nil {
//         return ""
//     }
//     return this.head.next.str
// }

// func (this *AllOne) GetMinKey() string {
//     if this.tail == this.head {
//         return ""
//     }
//     return this.tail.str
// }

import "container/list"

type Node struct {
    freq    int
    strSet  map[string]struct{}
    element *list.Element
}

type AllOne struct {
    dll          *list.List
    strToNodeMap map[string]*Node
}

func Constructor() AllOne {
    return AllOne{
        dll:          list.New(),
        strToNodeMap: make(map[string]*Node),
    }
}

func (this *AllOne) Inc(key string) {
    if node, exists := this.strToNodeMap[key]; exists {
        freq := node.freq
        delete(node.strSet, key)

        var nextNode *Node
        if node.element.Next() != nil {
            nextNode = node.element.Next().Value.(*Node)
        }

        if nextNode == nil || nextNode.freq != freq+1 {
            newNode := &Node{
                freq:    freq + 1,
                strSet:  make(map[string]struct{}),
                element: this.dll.InsertAfter(nil, node.element),
            }
            newNode.strSet[key] = struct{}{}
            newNode.element.Value = newNode
            this.strToNodeMap[key] = newNode
        } else {
            nextNode.strSet[key] = struct{}{}
            this.strToNodeMap[key] = nextNode
        }

        if len(node.strSet) == 0 {
            this.dll.Remove(node.element)
        }
    } else {
        var firstNode *Node
        if this.dll.Len() > 0 {
            firstNode = this.dll.Front().Value.(*Node)
        }

        if firstNode == nil || firstNode.freq > 1 {
            newNode := &Node{
                freq:    1,
                strSet:  make(map[string]struct{}),
                element: this.dll.PushFront(nil),
            }
            newNode.strSet[key] = struct{}{}
            newNode.element.Value = newNode
            this.strToNodeMap[key] = newNode
        } else {
            firstNode.strSet[key] = struct{}{}
            this.strToNodeMap[key] = firstNode
        }
    }
}

func (this *AllOne) Dec(key string) {
    if node, exists := this.strToNodeMap[key]; exists {
        delete(node.strSet, key)
        freq := node.freq

        if freq == 1 {
            delete(this.strToNodeMap, key)
        } else {
            var prevNode *Node
            if node.element.Prev() != nil {
                prevNode = node.element.Prev().Value.(*Node)
            }

            if prevNode == nil || prevNode.freq != freq-1 {
                newNode := &Node{
                    freq:    freq - 1,
                    strSet:  make(map[string]struct{}),
                    element: this.dll.InsertBefore(nil, node.element),
                }
                newNode.strSet[key] = struct{}{}
                newNode.element.Value = newNode
                this.strToNodeMap[key] = newNode
            } else {
                prevNode.strSet[key] = struct{}{}
                this.strToNodeMap[key] = prevNode
            }
        }

        if len(node.strSet) == 0 {
            this.dll.Remove(node.element)
        }
    }
}

func (this *AllOne) GetMaxKey() string {
    if this.dll.Len() == 0 {
        return ""
    }
    lastNode := this.dll.Back().Value.(*Node)
    for k := range lastNode.strSet {
        return k
    }
    return ""
}

func (this *AllOne) GetMinKey() string {
    if this.dll.Len() == 0 {
        return ""
    }
    firstNode := this.dll.Front().Value.(*Node)
    for k := range firstNode.strSet {
        return k
    }
    return ""
}

// 超出时间限制
type AllOne1 struct {
    data map[string]int
}

func Constructor1() AllOne1 {
    return AllOne1{ data: make(map[string]int, 0) }
}

func (this *AllOne1) Inc(key string)  {
    if _, ok := this.data[key]; ok {
        this.data[key]++
    }else{
        this.data[key] = 1
    }
}

func (this *AllOne1) Dec(key string)  {
    if _, ok := this.data[key]; ok {
        this.data[key]--
        if this.data[key] < 1 {
            delete(this.data,key)
        }
    }
    
}

func (this *AllOne1) GetMaxKey() string {
    mx, key := 0, ""
    for i,v := range this.data {
        if v > mx {
            key = i
            mx = v
        }
    }
    return key
}

func (this *AllOne1) GetMinKey() string {
    mn, key := 1 << 32 - 1, ""
    for i, v := range this.data {
        if v < mn {
            key = i
            mn = v
        }
    }
    return key
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

func main() {
    // AllOne allOne = new AllOne();
    obj := Constructor()
    fmt.Println(obj)
    // allOne.inc("hello");
    obj.Inc("hello")
    fmt.Println(obj)
    // allOne.inc("hello");
    obj.Inc("hello")
    fmt.Println(obj)
    // allOne.getMaxKey(); // return "hello"
    fmt.Println(obj.GetMaxKey()) // hello
    // allOne.getMinKey(); // return "hello"
    fmt.Println(obj.GetMinKey()) // hello
    // allOne.inc("leet");
    obj.Inc("leet")
    fmt.Println(obj)
    // allOne.getMaxKey(); // return "hello"
    fmt.Println(obj.GetMaxKey()) // hello
    // allOne.getMinKey(); // return "leet"
    fmt.Println(obj.GetMinKey()) // leet

    // AllOne allOne = new AllOne();
    obj1 := Constructor1()
    fmt.Println(obj1)
    // allOne.inc("hello");
    obj1.Inc("hello")
    fmt.Println(obj1)
    // allOne.inc("hello");
    obj1.Inc("hello")
    fmt.Println(obj1)
    // allOne.getMaxKey(); // return "hello"
    fmt.Println(obj1.GetMaxKey()) // hello
    // allOne.getMinKey(); // return "hello"
    fmt.Println(obj1.GetMinKey()) // hello
    // allOne.inc("leet");
    obj1.Inc("leet")
    fmt.Println(obj1)
    // allOne.getMaxKey(); // return "hello"
    fmt.Println(obj1.GetMaxKey()) // hello
    // allOne.getMinKey(); // return "leet"
    fmt.Println(obj1.GetMinKey()) // leet
}