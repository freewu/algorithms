package main

// 2336. Smallest Number in Infinite Set
// You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].
// Implement the SmallestInfiniteSet class:
//     SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
//     int popSmallest() Removes and returns the smallest integer contained in the infinite set.
//     void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.
    
// Example 1:
// Input
// ["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
// [[], [2], [], [], [], [1], [], [], []]
// Output
// [null, null, 1, 2, 3, null, 1, 4, 5]
// Explanation
// SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
// smallestInfiniteSet.addBack(2);    // 2 is already in the set, so no change is made.
// smallestInfiniteSet.popSmallest(); // return 1, since 1 is the smallest number, and remove it from the set.
// smallestInfiniteSet.popSmallest(); // return 2, and remove it from the set.
// smallestInfiniteSet.popSmallest(); // return 3, and remove it from the set.
// smallestInfiniteSet.addBack(1);    // 1 is added back to the set.
// smallestInfiniteSet.popSmallest(); // return 1, since 1 was added back to the set and
//                                    // is the smallest number, and remove it from the set.
// smallestInfiniteSet.popSmallest(); // return 4, and remove it from the set.
// smallestInfiniteSet.popSmallest(); // return 5, and remove it from the set.

// Constraints:
//     1 <= num <= 1000
//     At most 1000 calls will be made in total to popSmallest and addBack.

import "fmt"
import "container/heap"

// Define the min-heap for integers
type IntHeap []int
// Define the methods required by the heap interface
func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i int, j int) bool {return h[i] < h[j]}
func (h IntHeap) Swap(i int, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push (x any) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop () any {
    n := len(*h)
    x := (*h)[n - 1]
    *h = (*h)[:n - 1]
    return x
}

type SmallestInfiniteSet struct {
    backHeap *IntHeap
    inHeap map[int]bool
    next int
}

func Constructor() SmallestInfiniteSet {
    return SmallestInfiniteSet{
        backHeap: &IntHeap{},
        inHeap: make(map[int]bool),
        next: 1,
    }
}

func (this *SmallestInfiniteSet) PopSmallest() int {
    if (*this.backHeap).Len() == 0 || this.next < (*this.backHeap)[0] {
        this.next++
        return this.next - 1
    } else {
        smallestVal := heap.Pop(this.backHeap).(int)
        this.inHeap[smallestVal] = false
        return smallestVal
    }
}

func (this *SmallestInfiniteSet) AddBack(num int)  {
    if num < this.next {
        if !this.inHeap[num] {
            this.inHeap[num] = true
            heap.Push(this.backHeap, num)
        }
    }
}


// //有序集合 treeset.Set有序 不重复 k(可为空),v(不可为空)
// type SmallestInfiniteSet struct {
//     confine int
//     s *treeset.Set
// }


// func Constructor() SmallestInfiniteSet {
//     return SmallestInfiniteSet{
//         confine: 1, //可删除的最小数集合边界 
//         s: treeset.NewWithIntComparator(),
//     }
// }


// //删除最左侧数据即s集合开头数据 如果s为空将confine返回并后移
// func (this *SmallestInfiniteSet) PopSmallest() int {
//     if this.s.Empty(){
//         res := this.confine
//         this.confine++
//         return res
//     }
//     it := this.s.Iterator()
//     it.Next()
//     res := it.Value().(int)
//     this.s.Remove(res)
//     return res
// }

// //如果小于边界的左侧 即在s集合内则加入(因为有可能已经弹出) 否则不加入(本就包含)
// func (this *SmallestInfiniteSet) AddBack(num int)  {
//     if num < this.confine{
//         this.s.Add(num)
//     }
// }


/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

func main() {
    // SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
    obj := Constructor()
    // smallestInfiniteSet.addBack(2);    // 2 is already in the set, so no change is made.
    obj.AddBack(2)
    fmt.Println(obj)
    // smallestInfiniteSet.popSmallest(); // return 1, since 1 is the smallest number, and remove it from the set.
    fmt.Println(obj.PopSmallest()) // 1
    // smallestInfiniteSet.popSmallest(); // return 2, and remove it from the set.
    fmt.Println(obj.PopSmallest()) // 2
    // smallestInfiniteSet.popSmallest(); // return 3, and remove it from the set.
    fmt.Println(obj.PopSmallest()) // 3
    // smallestInfiniteSet.addBack(1);    // 1 is added back to the set.
    obj.AddBack(1)
    // smallestInfiniteSet.popSmallest(); // return 1, since 1 was added back to the set and
    //                                    // is the smallest number, and remove it from the set.
    fmt.Println(obj.PopSmallest()) // 1
    // smallestInfiniteSet.popSmallest(); // return 4, and remove it from the set.
    fmt.Println(obj.PopSmallest()) // 4
    // smallestInfiniteSet.popSmallest(); // return 5, and remove it from the set.
    fmt.Println(obj.PopSmallest()) // 5
}