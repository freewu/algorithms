package main

// 460. LFU Cache
// Design and implement a data structure for a Least Frequently Used (LFU) cache.
// Implement the LFUCache class:
//     LFUCache(int capacity) 
//         Initializes the object with the capacity of the data structure.
//     int get(int key) 
//         Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
//     void put(int key, int value) 
//         Update the value of the key if present, or inserts the key if not already present. 
//         When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item. 
//         For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.

// To determine the least frequently used key, a use counter is maintained for each key in the cache. 
// The key with the smallest use counter is the least frequently used key.

// When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). 
// The use counter for a key in the cache is incremented either a get or put operation is called on it.

// The functions get and put must each run in O(1) average time complexity.

// Example 1:
// Input
// ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
// Output
// [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
// Explanation
// // cnt(x) = the use counter for key x
// // cache=[] will show the last used order for tiebreakers (leftmost element is  most recent)
// LFUCache lfu = new LFUCache(2);
// lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
// lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
// lfu.get(1);      // return 1
//                  // cache=[1,2], cnt(2)=1, cnt(1)=2
// lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
//                  // cache=[3,1], cnt(3)=1, cnt(1)=2
// lfu.get(2);      // return -1 (not found)
// lfu.get(3);      // return 3
//                  // cache=[3,1], cnt(3)=2, cnt(1)=2
// lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
//                  // cache=[4,3], cnt(4)=1, cnt(3)=2
// lfu.get(1);      // return -1 (not found)
// lfu.get(3);      // return 3
//                  // cache=[3,4], cnt(4)=1, cnt(3)=3
// lfu.get(4);      // return 4
//                  // cache=[4,3], cnt(4)=2, cnt(3)=3
 
// Constraints:
//     1 <= capacity <= 10^4
//     0 <= key <= 10^5
//     0 <= value <= 10^9
//     At most 2 * 10^5 calls will be made to get and put.

import "fmt"
// import "container/list"

// type node struct {
//     key       int
//     val       int
//     frequency int
// }

// type LFUCache struct {
//     nodes        map[int]*list.Element
//     lists        map[int]*list.List
//     cap          int
//     minFrequency int
// }

// func Constructor(capacity int) LFUCache {
//     return LFUCache{
//         cap:          capacity,
//         nodes:        make(map[int]*list.Element),
//         lists:        make(map[int]*list.List),
//         minFrequency: 0,
//     }
// }

// func (this *LFUCache) Get(key int) int {
//     if _, ok := this.nodes[key]; !ok {
//         return -1
//     }
//     currNode := this.nodes[key].Value.(*node)
//     this.lists[currNode.frequency].Remove(this.nodes[key])
//     currNode.frequency++
//     newList := this.lists[currNode.frequency]
//     if _, ok := this.lists[currNode.frequency]; !ok {
//         newList = list.New()
//     }
//     this.lists[currNode.frequency] = newList
//     newNode := newList.PushFront(currNode)
//     this.nodes[key] = newNode
//     if currNode.frequency - 1 == this.minFrequency && this.lists[currNode.frequency-1].Len() == 0 {
//         this.minFrequency++
//     }
//     return currNode.val
// }

// func (this *LFUCache) Put(key int, value int) {
//     if this.cap == 0 {
//         return
//     }
//     if val, ok := this.nodes[key]; ok { // 如果存在，更新访问次数
//         currNode := val.Value.(*node)
//         currNode.val = value
//         this.Get(key)
//         return
//     }
//     if this.cap == len(this.nodes) { // key 不存在，且缓存满了，需要删除
//         // 删除访问频次最低的链表的表尾的节点
//         minList := this.lists[this.minFrequency]
//         backNode := minList.Back()
//         delete(this.nodes, backNode.Value.(*node).key)
//         minList.Remove(backNode)
//     }
//     this.minFrequency = 1
//     currNode := &node{
//         key:       key,
//         val:       value,
//         frequency: 1,
//     }
//     if _, ok := this.lists[1]; !ok {
//         this.lists[1] = list.New()
//     }
//     newNode := this.lists[1].PushFront(currNode)
//     this.nodes[key] = newNode
// }

// type LFUCache struct {
//     Cap       int
//     NodeIndex map[int]*list.Element
//     FreqIndex map[int]*list.List
//     MinFreq   int
// }

// type Node struct {
//     Key   int
//     Value int
//     Freq  int
// }

// func Constructor(capacity int) LFUCache {
//     return LFUCache{
//         Cap:       capacity,
//         NodeIndex: make(map[int]*list.Element),
//         FreqIndex: make(map[int]*list.List),
//         MinFreq:   1 << 32 -1,
//     }
// }

// func (this *LFUCache) Get(key int) int {
//     var (
//         ele *list.Element
//         ok  bool
//     )
//     if ele, ok = this.NodeIndex[key]; !ok {
//         return -1
//     }
//     return this.up(ele).Value
// }

// // Put if not existed and cap meet limitation, remove a
// // node form minFreq list
// func (this *LFUCache) Put(key int, value int) {
//     if ele, ok := this.NodeIndex[key]; ok {
//         node := this.up(ele)
//         node.Value = value
//         return
//     }
//     if this.Cap == len(this.NodeIndex) {
//         this.remove()
//     }
//     this.insert(&Node{Key: key, Value: value, Freq: 1})
// }

// func (this *LFUCache) remove() {
//     rl := this.FreqIndex[this.MinFreq]
//     toRemove := rl.Back()
//     delete(this.NodeIndex, toRemove.Value.(*Node).Key)
//     rl.Remove(toRemove)
// }

// func (this *LFUCache) up(ele *list.Element) *Node {
//     node := ele.Value.(*Node)
//     l, ok := this.FreqIndex[node.Freq]
//     if ok {
//         l.Remove(ele)
//         if this.MinFreq == node.Freq && l.Len() == 0 {
//             this.MinFreq += 1
//         }
//     } else {
//         return nil
//     }
//     node.Freq++
//     this.insert(node)
//     return node
// }

// func (this *LFUCache) insert(node *Node) {
//     if _, ok := this.FreqIndex[node.Freq]; !ok {
//         this.FreqIndex[node.Freq] = list.New()
//     }
//     if this.MinFreq > node.Freq {
//         this.MinFreq = node.Freq
//     }
//     this.NodeIndex[node.Key] = this.FreqIndex[node.Freq].PushFront(node)
// }

type Node struct {
    Key, Value, Freq int
    Prev, Next *Node
}

type List struct {
    Head, Tail *Node
    Size int
}

type LFUCache struct {
    Cache map[int]*Node
    FrequencyMap map[int]*List
    MinFreq int
    Capacity int
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        Cache : make(map[int]*Node),
        FrequencyMap : make(map[int]*List),
        MinFreq : 0,
        Capacity : capacity,
    }
}

func CreateList() *List {
    head := &Node{}
    tail := &Node{}
    head.Next = tail
    tail.Prev = head

    return &List{
        Head : head, 
        Tail : tail, 
        Size : 0,
    }
}

func (list *List)  AddNode(node *Node) {
    next, prev := list.Tail, list.Tail.Prev
    next.Prev = node
    prev.Next = node
    node.Next = next
    node.Prev = prev
    list.Size++
}

func (list *List) DeleteNode(node *Node) {
    next, prev := node.Next, node.Prev
    next.Prev, prev.Next = prev, next
    list.Size--
}

func (this *LFUCache) Get(key int) int {
    node, ok := this.Cache[key]
    if !ok {
        return -1
    }
    list, ok := this.FrequencyMap[node.Freq]
    if ok {
        list.DeleteNode(node)
        if list.Size == 0 && this.MinFreq == node.Freq {
        this.MinFreq += 1
        }
    }
    node.Freq += 1
    nextList, ok := this.FrequencyMap[node.Freq]
    if !ok {
        nextList = CreateList()
    }
    nextList.AddNode(node)
    this.FrequencyMap[node.Freq] = nextList
    return node.Value
}



func (this *LFUCache) Put(key int, value int)  {
    if this.Capacity == 0 {
        return 
    }
    if node, ok := this.Cache[key]; ok {
        node.Value = value 
        this.Get(node.Key)
        return
    }
    if len(this.Cache) >= this.Capacity {
        minFrequencyList := this.FrequencyMap[this.MinFreq]
        minFrequencyNode := minFrequencyList.Head.Next
        minFrequencyList.DeleteNode(minFrequencyNode)
        delete(this.Cache, minFrequencyNode.Key)
    }
    node := &Node{
        Key: key,
        Value: value, 
        Freq: 1,
    }
    list, ok := this.FrequencyMap[node.Freq]
    if !ok {
        list = CreateList()
    }
    list.AddNode(node)
    this.MinFreq = node.Freq
    this.Cache[key] = node
    this.FrequencyMap[node.Freq] = list 
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
    // LFUCache lfu = new LFUCache(2);
    obj := Constructor(2)
    // lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
    obj.Put(1,1)
    fmt.Println(obj)
    // lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
    obj.Put(2,2)
    fmt.Println(obj)
    // lfu.get(1);      // return 1
    //                  // cache=[1,2], cnt(2)=1, cnt(1)=2
    fmt.Println(obj.Get(1)) // 1
    fmt.Println(obj)
    // lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
    //                  // cache=[3,1], cnt(3)=1, cnt(1)=2
    obj.Put(3,3)
    fmt.Println(obj)
    // lfu.get(2);      // return -1 (not found)
    fmt.Println(obj.Get(2)) // -1
    // lfu.get(3);      // return 3
    //                  // cache=[3,1], cnt(3)=2, cnt(1)=2
    fmt.Println(obj.Get(3)) // 3
    fmt.Println(obj)
    // lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
    //                  // cache=[4,3], cnt(4)=1, cnt(3)=2
    obj.Put(4,4)
    fmt.Println(obj)
    // lfu.get(1);      // return -1 (not found)
    fmt.Println(obj.Get(1)) // -1
    // lfu.get(3);      // return 3
    //                  // cache=[3,4], cnt(4)=1, cnt(3)=3
    fmt.Println(obj.Get(3)) // 3
    fmt.Println(obj)
    // lfu.get(4);      // return 4
    //                  // cache=[4,3], cnt(4)=2, cnt(3)=3
    fmt.Println(obj.Get(4)) // 4
    fmt.Println(obj)
}