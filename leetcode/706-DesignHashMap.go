package main

// 706. Design HashMap
// Design a HashMap without using any built-in hash table libraries.
// Implement the MyHashMap class:
//     MyHashMap() initializes the object with an empty map.
//     void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
//     int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
//     void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
    
// Example 1:
// Input
// ["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
// [[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
// Output
// [null, null, null, 1, -1, null, 1, null, -1]
// Explanation
// MyHashMap myHashMap = new MyHashMap();
// myHashMap.put(1, 1); // The map is now [[1,1]]
// myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
// myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
// myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
// myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
// myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
// myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
// myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]
 
// Constraints:
//     0 <= key, value <= 10^6
//     At most 10^4 calls will be made to put, get, and remove.

import "fmt"
// import "fnv"
// import "binary"

type MyHashMap struct {
    data map[int]int
}

func Constructor() MyHashMap {
    return MyHashMap{
        data: make(map[int]int),
    }
}

func (this *MyHashMap) Put(key int, value int) {
    this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
    if val, ok := this.data[key]; ok {
        return val
    } else {
        return -1
    }
}

func (this *MyHashMap) Remove(key int) {
    delete(this.data, key)
}

// 使用数组来存放
type Node struct {
    key int
    value int
    next *Node
}

type MyHashMap1 struct {
    data []*Node
}

const MyHashMapSize = 1000

func Constructor1() MyHashMap1 {
    return MyHashMap1{
        data: make([]*Node, MyHashMapSize),
    }
}


func (this *MyHashMap1) Put(key int, value int)  {
    keyHash := key % MyHashMapSize
    node := this.data[keyHash]
    for nil != node {
        if key == node.key {
            node.value = value
            return
        }
        node = node.next
    }
    this.data[keyHash] = &Node{
        key: key,
        value: value,
        next: this.data[keyHash],
    }
}


func (this *MyHashMap1) Get(key int) int {
    keyHash := key % MyHashMapSize
    node := this.data[keyHash]
    for nil != node {
        if key == node.key {
            return node.value
        }
        node = node.next
    }
    return -1
}


func (this *MyHashMap1) Remove(key int)  {
    keyHash := key % MyHashMapSize
    dummy := &Node{
        next: this.data[keyHash],
    }
    node := dummy
    for nil != node.next {
        if key == node.next.key {
            node.next = node.next.next
            break 
        }
        node = node.next
    }
    this.data[keyHash] = dummy.next
    return
}

// type MyHashMap2 struct {
//     data [][]*pair
//     size int
// }

// type pair struct {
//     key int
//     val int
// }

// func Constructor2() MyHashMap2 {
//     return MyHashMap2 {
//         data: make([][]*pair, 1000),
//         size: 1000,
//     }
// }

// func hash(n, max int) int {
// 	h := fnv.New32a()
// 	bs := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(bs, uint32(n))
// 	h.Write(bs)
// 	v := h.Sum32()
// 	return int(v % uint32(max))
// }

// func (this *MyHashMap2) Put(key int, value int)  {
//     ind := hash(key, this.size)
//     for i, val := range this.data[ind] {
//         if val.key == key {
//             this.data[ind][i].val = value
//             return
//         }
//     }
//     this.data[ind] = append(this.data[ind], &pair{key: key, val: value})
// }


// func (this *MyHashMap2) Get(key int) int {
//     ind := hash(key, this.size)
//     for _, val := range this.data[ind] {
//         if val.key == key {
//             return val.val
//         }
//     }
//     return -1
// }


// func (this *MyHashMap2) Remove(key int)  {
//     ind := hash(key, this.size)
//     for i, val := range this.data[ind] {
//         if val.key == key {
//             l := len(this.data[ind])
//             this.data[ind][i] = this.data[ind][l-1]
//             this.data[ind] = this.data[ind][:l-1]
//             return
//         }
//     }
// }


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main() {
    // MyHashMap myHashMap = new MyHashMap();
    obj := Constructor()
    // myHashMap.put(1, 1); // The map is now [[1,1]]
    obj.Put(1,1)
    fmt.Println(obj) // [[1,1]]
    // myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
    obj.Put(2,2)
    fmt.Println(obj) // [[1,1], [2,2]]
    // myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
    fmt.Println(obj.Get(1)) // 1
    fmt.Println(obj) // [[1,1], [2,2]]
    // myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
    fmt.Println(obj.Get(3)) // -1
    fmt.Println(obj) // [[1,1], [2,2]]
    // myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
    obj.Put(2,1)
    fmt.Println(obj) // [[1,1], [2,1]]
    // myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
    fmt.Println(obj.Get(2)) // 1
    fmt.Println(obj) // [[1,1], [2,1]]
    // myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
    obj.Remove(2)
    fmt.Println(obj) // [[1,1]]
    // myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]
    fmt.Println(obj.Get(2)) // -1
    fmt.Println(obj) // [[1,1]]

    // MyHashMap myHashMap = new MyHashMap();
    obj1 := Constructor1()
    // myHashMap.put(1, 1); // The map is now [[1,1]]
    obj1.Put(1,1)
    fmt.Println(obj1) // [[1,1]]
    // myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
    obj1.Put(2,2)
    fmt.Println(obj1) // [[1,1], [2,2]]
    // myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
    fmt.Println(obj1.Get(1)) // 1
    fmt.Println(obj1) // [[1,1], [2,2]]
    // myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
    fmt.Println(obj1.Get(3)) // -1
    fmt.Println(obj1) // [[1,1], [2,2]]
    // myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
    obj1.Put(2,1)
    fmt.Println(obj1) // [[1,1], [2,1]]
    // myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
    fmt.Println(obj1.Get(2)) // 1
    fmt.Println(obj1) // [[1,1], [2,1]]
    // myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
    obj1.Remove(2)
    fmt.Println(obj1) // [[1,1]]
    // myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]
    fmt.Println(obj1.Get(2)) // -1
    fmt.Println(obj1) // [[1,1]]


    // // MyHashMap myHashMap = new MyHashMap();
    // obj2 := Constructor2()
    // // myHashMap.put(1, 1); // The map is now [[1,1]]
    // obj2.Put(1,1)
    // fmt.Println(obj2) // [[1,1]]
    // // myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
    // obj2.Put(2,2)
    // fmt.Println(obj2) // [[1,1], [2,2]]
    // // myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
    // fmt.Println(obj2.Get(1)) // 1
    // fmt.Println(obj2) // [[1,1], [2,2]]
    // // myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
    // fmt.Println(obj2.Get(3)) // -1
    // fmt.Println(obj2) // [[1,1], [2,2]]
    // // myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
    // obj2.Put(2,1)
    // fmt.Println(obj2) // [[1,1], [2,1]]
    // // myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
    // fmt.Println(obj2.Get(2)) // 1
    // fmt.Println(obj2) // [[1,1], [2,1]]
    // // myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
    // obj2.Remove(2)
    // fmt.Println(obj2) // [[1,1]]
    // // myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]
    // fmt.Println(obj2.Get(2)) // -1
    // fmt.Println(obj2) // [[1,1]]
}