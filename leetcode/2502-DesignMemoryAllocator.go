package main

// 2502. Design Memory Allocator
// You are given an integer n representing the size of a 0-indexed memory array. 
// All memory units are initially free.

// You have a memory allocator with the following functionalities:
//     1. Allocate a block of size consecutive free memory units and assign it the id mID.
//     2. Free all memory units with the given id mID.

// Note that:
//     1. Multiple blocks can be allocated to the same mID.
//     2. You should free all the memory units with mID, even if they were allocated in different blocks.

// Implement the Allocator class:
//     Allocator(int n) 
//         Initializes an Allocator object with a memory array of size n.
//     int allocate(int size, int mID) 
//         Find the leftmost block of size consecutive free memory units 
//         and allocate it with the id mID. Return the block's first index. 
//         If such a block does not exist, return -1.
//     int free(int mID) 
//         Free all memory units with the id mID. Return the number of memory units you have freed.

// Example 1:
// Input
// ["Allocator", "allocate", "allocate", "allocate", "free", "allocate", "allocate", "allocate", "free", "allocate", "free"]
// [[10], [1, 1], [1, 2], [1, 3], [2], [3, 4], [1, 1], [1, 1], [1], [10, 2], [7]]
// Output
// [null, 0, 1, 2, 1, 3, 1, 6, 3, -1, 0]
// Explanation
// Allocator loc = new Allocator(10); // Initialize a memory array of size 10. All memory units are initially free.
// loc.allocate(1, 1); // The leftmost block's first index is 0. The memory array becomes [1,_,_,_,_,_,_,_,_,_]. We return 0.
// loc.allocate(1, 2); // The leftmost block's first index is 1. The memory array becomes [1,2,_,_,_,_,_,_,_,_]. We return 1.
// loc.allocate(1, 3); // The leftmost block's first index is 2. The memory array becomes [1,2,3,_,_,_,_,_,_,_]. We return 2.
// loc.free(2); // Free all memory units with mID 2. The memory array becomes [1,_, 3,_,_,_,_,_,_,_]. We return 1 since there is only 1 unit with mID 2.
// loc.allocate(3, 4); // The leftmost block's first index is 3. The memory array becomes [1,_,3,4,4,4,_,_,_,_]. We return 3.
// loc.allocate(1, 1); // The leftmost block's first index is 1. The memory array becomes [1,1,3,4,4,4,_,_,_,_]. We return 1.
// loc.allocate(1, 1); // The leftmost block's first index is 6. The memory array becomes [1,1,3,4,4,4,1,_,_,_]. We return 6.
// loc.free(1); // Free all memory units with mID 1. The memory array becomes [_,_,3,4,4,4,_,_,_,_]. We return 3 since there are 3 units with mID 1.
// loc.allocate(10, 2); // We can not find any free block with 10 consecutive free memory units, so we return -1.
// loc.free(7); // Free all memory units with mID 7. The memory array remains the same since there is no memory unit with mID 7. We return 0.

// Constraints:
//     1 <= n, size, mID <= 1000
//     At most 1000 calls will be made to allocate and free.

import "fmt"

// type AllocNode struct {
//     prev  *AllocNode
//     next  *AllocNode
//     start int
//     size  int
// }

// type Allocator struct {
//     allocMap map[int][]*AllocNode
//     alloc    *AllocNode
//     size     int
//     free     int
// }

// func Constructor(n int) Allocator {
//     return Allocator{
//         allocMap: make(map[int][]*AllocNode),
//         size:     n,
//         free:     n,
//         alloc:    &AllocNode{},
//     }
// }

// func (al *Allocator) AddAllocNode(prev *AllocNode, size int, mID int) int {
//     node := &AllocNode{
//         size:  size,
//         prev:  prev,
//         next:  prev.next,
//         start: prev.start + prev.size,
//     }
//     if prev.next != nil {
//         prev.next.prev = node
//     }
//     prev.next = node
//     if v, ok := al.allocMap[mID]; ok {
//         al.allocMap[mID] = append(v, node)
//     } else {
//         al.allocMap[mID] = []*AllocNode{node}
//     }
//     al.free = al.free - size
//     return node.start
// }

// func (al *Allocator) Allocate(size int, mID int) int {
//     if size > al.free {
//         return -1
//     }
//     node := al.alloc
//     for node != nil {
//         if (node.next == nil) && ((al.size - node.start - node.size) >= size) {
//             return al.AddAllocNode(node, size, mID)
//         }
//         if (node.next != nil) && ((node.next.start - node.start - node.size) >= size) {
//             return al.AddAllocNode(node, size, mID)
//         }
//         node = node.next
//     }
//     return -1
// }

// func (al *Allocator) Free(mID int) int {
//     res := 0
//     for _, node := range al.allocMap[mID] {
//         node.prev.next = node.next
//         if node.next != nil {
//             node.next.prev = node.prev
//         }
//         res = res + node.size
//     }
//     delete(al.allocMap, mID)
//     al.free = al.free + res
//     return res
// }

type Pair struct {
    Start, End int
}

type Allocator struct {
    data []bool
    idmp [][]Pair
}

func Constructor(n int) Allocator {
    return Allocator{make([]bool, n), make([][]Pair, 1001)}
}

func (this *Allocator) Allocate(size int, mID int) int {
    count := 0
    for i := 0; i < len(this.data); i++ {
        if !this.data[i] { count++ }
        if i >= size && !this.data[i - size] { count-- }
        if count == size {
            this.idmp[mID] = append(this.idmp[mID], Pair{i - size + 1, i})
            for j := i - size + 1; j <= i; j++ {
                this.data[j] = true
            }
            return i - size + 1
        }
    }
    return -1
}

func (this *Allocator) Free(mID int) int {
    sum := 0
    for _, seg := range this.idmp[mID] {
        sum += (seg.End - seg.Start + 1)
        for i := seg.Start; i <= seg.End; i++ {
            this.data[i] = false
        }
    }
    this.idmp[mID]= []Pair{}
    return sum
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */

func main() {
    // Allocator loc = new Allocator(10); // Initialize a memory array of size 10. All memory units are initially free.
    obj := Constructor(10)
    fmt.Println(obj)
    // loc.allocate(1, 1); // The leftmost block's first index is 0. The memory array becomes [1,_,_,_,_,_,_,_,_,_]. We return 0.
    fmt.Println(obj.Allocate(1, 1)) // 0
    fmt.Println(obj)
    // loc.allocate(1, 2); // The leftmost block's first index is 1. The memory array becomes [1,2,_,_,_,_,_,_,_,_]. We return 1.
    fmt.Println(obj.Allocate(1, 2)) // 1
    fmt.Println(obj)
    // loc.allocate(1, 3); // The leftmost block's first index is 2. The memory array becomes [1,2,3,_,_,_,_,_,_,_]. We return 2.
    fmt.Println(obj.Allocate(1, 3)) // 2
    fmt.Println(obj)
    // loc.free(2); // Free all memory units with mID 2. The memory array becomes [1,_, 3,_,_,_,_,_,_,_]. We return 1 since there is only 1 unit with mID 2.
    fmt.Println(obj.Free(2)) // 1
    fmt.Println(obj)
    // loc.allocate(3, 4); // The leftmost block's first index is 3. The memory array becomes [1,_,3,4,4,4,_,_,_,_]. We return 3.
    fmt.Println(obj.Allocate(3, 4)) // 3
    fmt.Println(obj)
    // loc.allocate(1, 1); // The leftmost block's first index is 1. The memory array becomes [1,1,3,4,4,4,_,_,_,_]. We return 1.
    fmt.Println(obj.Allocate(1, 1)) // 1
    fmt.Println(obj)
    // loc.allocate(1, 1); // The leftmost block's first index is 6. The memory array becomes [1,1,3,4,4,4,1,_,_,_]. We return 6.
    fmt.Println(obj.Allocate(1, 1)) // 6
    fmt.Println(obj)
    // loc.free(1); // Free all memory units with mID 1. The memory array becomes [_,_,3,4,4,4,_,_,_,_]. We return 3 since there are 3 units with mID 1.
    fmt.Println(obj.Free(2)) // 1
    fmt.Println(obj)
    // loc.allocate(10, 2); // We can not find any free block with 10 consecutive free memory units, so we return -1.
    fmt.Println(obj.Allocate(10, 2)) // -1
    fmt.Println(obj)
    // loc.free(7); // Free all memory units with mID 7. The memory array remains the same since there is no memory unit with mID 7. We return 0.
    fmt.Println(obj.Free(7)) // 0
    fmt.Println(obj)
}