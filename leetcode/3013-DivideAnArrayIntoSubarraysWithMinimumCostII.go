package main

// 3013. Divide an Array Into Subarrays With Minimum Cost II
// You are given a 0-indexed array of integers nums of length n, and two positive integers k and dist.

// The cost of an array is the value of its first element. 
// For example, the cost of [1,2,3] is 1 while the cost of [3,4,1] is 3.

// You need to divide nums into k disjoint contiguous subarrays, 
// such that the difference between the starting index of the second subarray and the starting index of the kth subarray should be less than or equal to dist. 
// In other words, if you divide nums into the subarrays nums[0..(i1 - 1)], nums[i1..(i2 - 1)], ..., nums[ik-1..(n - 1)], then ik-1 - i1 <= dist.

// Return the minimum possible sum of the cost of these subarrays.

// Example 1:
// Input: nums = [1,3,2,6,4,2], k = 3, dist = 3
// Output: 5
// Explanation: The best possible way to divide nums into 3 subarrays is: [1,3], [2,6,4], and [2]. This choice is valid because ik-1 - i1 is 5 - 2 = 3 which is equal to dist. The total cost is nums[0] + nums[2] + nums[5] which is 1 + 2 + 2 = 5.
// It can be shown that there is no possible way to divide nums into 3 subarrays at a cost lower than 5.

// Example 2:
// Input: nums = [10,1,2,2,2,1], k = 4, dist = 3
// Output: 15
// Explanation: The best possible way to divide nums into 4 subarrays is: [10], [1], [2], and [2,2,1]. This choice is valid because ik-1 - i1 is 3 - 1 = 2 which is less than dist. The total cost is nums[0] + nums[1] + nums[2] + nums[3] which is 10 + 1 + 2 + 2 = 15.
// The division [10], [1], [2,2,2], and [1] is not valid, because the difference between ik-1 and i1 is 5 - 1 = 4, which is greater than dist.
// It can be shown that there is no possible way to divide nums into 4 subarrays at a cost lower than 15.

// Example 3:
// Input: nums = [10,8,18,9], k = 3, dist = 1
// Output: 36
// Explanation: The best possible way to divide nums into 4 subarrays is: [10], [8], and [18,9]. This choice is valid because ik-1 - i1 is 2 - 1 = 1 which is equal to dist.The total cost is nums[0] + nums[1] + nums[2] which is 10 + 8 + 18 = 36.
// The division [10], [8,18], and [9] is not valid, because the difference between ik-1 and i1 is 3 - 1 = 2, which is greater than dist.
// It can be shown that there is no possible way to divide nums into 3 subarrays at a cost lower than 36.

// Constraints:
//     3 <= n <= 10^5
//     1 <= nums[i] <= 10^9
//     3 <= k <= n
//     k - 2 <= dist <= n - 2

import "fmt"
// import "github.com/emirpasic/gods/v2/maps/treemap"
// import "github.com/emirpasic/gods/trees/redblacktree"

// func minimumCost(nums []int, k int, dist int) int64 {
//     k -= 1
//     primary := treemap.NewWith[int, int](func(x, y int) int {
//         if nums[x] == nums[y] {
//             return x - y
//         }
//         return nums[x] - nums[y]
//     })
//     reserve := treemap.NewWith[int, int](func(x, y int) int {
//         if nums[x] == nums[y] {
//             return x - y
//         }
//         return nums[x] - nums[y]
//     })
//     min := func (x, y int64) int64 { if x < y { return x; }; return y; }
//     res, sum, n := int64(1 << 61), int64(0), len(nums)
//     for right := 1; right < n; right++ {
//         if right > dist {
//             left := right - dist - 1
//             if _, ok := primary.Get(left); ok {
//                 primary.Remove(left)
//                 sum -= int64(nums[left])
//             } else {
//                 reserve.Remove(left)
//             }
//         }
//         reserve.Put(right, right)
//         mn, _, _ := reserve.Min()
//         reserve.Remove(mn)
//         primary.Put(mn, mn)
//         sum += int64(nums[mn])
//         if primary.Size() > k {
//             mx, _, _ := primary.Max()
//             reserve.Put(mx, mx)
//             primary.Remove(mx)
//             sum -= int64(nums[mx])
//         }
//         if primary.Size() == k {
//             res = min(res, sum)
//         }
//     }
//     return res + int64(nums[0])
// }

// func minimumCost1(nums []int, k, dist int) int64 {
//     k--
//     L := redblacktree.NewWithIntComparator()
//     R := redblacktree.NewWithIntComparator()
//     add := func(t *redblacktree.Tree, x int) {
//         if v, ok := t.Get(x); ok {
//             t.Put(x, v.(int) + 1)
//         } else {
//             t.Put(x, 1)
//         }
//     }
//     del := func(t *redblacktree.Tree, x int) {
//         c, _ := t.Get(x)
//         if c.(int) > 1 {
//             t.Put(x, c.(int)-1)
//         } else {
//             t.Remove(x)
//         }
//     }
//     sumL := nums[0]
//     for _, x := range nums[1 : dist+2] {
//         sumL += x
//         add(L, x)
//     }
//     sizeL := dist + 1
//     l2r := func() {
//         x := L.Right().Key.(int)
//         sumL -= x
//         sizeL--
//         del(L, x)
//         add(R, x)
//     }
//     r2l := func() {
//         x := R.Left().Key.(int)
//         sumL += x
//         sizeL++
//         del(R, x)
//         add(L, x)
//     }
//     for sizeL > k {
//         l2r()
//     }
//     res := sumL
//     for i := dist + 2; i < len(nums); i++ {
//         // 移除 out
//         out := nums[i-dist-1]
//         if _, ok := L.Get(out); ok {
//             sumL -= out
//             sizeL--
//             del(L, out)
//         } else {
//             del(R, out)
//         }
//         // 添加 in
//         in := nums[i]
//         if in < L.Right().Key.(int) {
//             sumL += in
//             sizeL++
//             add(L, in)
//         } else {
//             add(R, in)
//         }
//         // 维护大小
//         if sizeL == k-1 {
//             r2l()
//         } else if sizeL == k+1 {
//             l2r()
//         }
//         res = min(res, sumL)
//     }
//     return int64(res)
// }


// type MultiSet struct {
//     tree    *redblacktree.Tree
//     counter map[int]int
//     size    int
// }

// func NewMultiSet() *MultiSet {
//     return &MultiSet{
//         tree:    redblacktree.NewWithIntComparator(),
//         counter: make(map[int]int),
//         size:    0,
//     }
// }

// func (ms *MultiSet) Add(x int) {
//     if count, exists := ms.counter[x]; exists {
//         ms.counter[x] = count + 1
//     } else {
//         ms.counter[x] = 1
//         ms.tree.Put(x, struct{}{})
//     }
//     ms.size++
// }

// func (ms *MultiSet) Remove(x int) bool {
//     if count, exists := ms.counter[x]; exists {
//         if count == 1 {
//             delete(ms.counter, x)
//             ms.tree.Remove(x)
//         } else {
//             ms.counter[x] = count - 1
//         }
//         ms.size--
//         return true
//     }
//     return false
// }

// func (ms *MultiSet) Size() int {
//     return ms.size
// }

// func (ms *MultiSet) IsEmpty() bool {
//     return ms.size == 0
// }

// func (ms *MultiSet) First() (int, bool) {
//     if ms.tree.Empty() {
//         return 0, false
//     }
//     return ms.tree.Left().Key.(int), true
// }

// func (ms *MultiSet) Last() (int, bool) {
//     if ms.tree.Empty() {
//         return 0, false
//     }
//     return ms.tree.Right().Key.(int), true
// }

// func (ms *MultiSet) Contains(x int) bool {
//     _, exists := ms.counter[x]
//     return exists
// }

// type Container struct {
//     k   int
//     st1 *MultiSet
//     st2 *MultiSet
//     sm  int64
// }

// func NewContainer(k int) *Container {
//     return &Container{
//         k:   k,
//         st1: NewMultiSet(),
//         st2: NewMultiSet(),
//         sm:  0,
//     }
// }

// func (m *Container) adjust() {
//     for m.st1.Size() < m.k && !m.st2.IsEmpty() {
//         if x, ok := m.st2.First(); ok {
//             m.st2.Remove(x)
//             m.st1.Add(x)
//             m.sm += int64(x)
//         }
//     }
//     for m.st1.Size() > m.k {
//         if x, ok := m.st1.Last(); ok {
//             m.st1.Remove(x)
//             m.st2.Add(x)
//             m.sm -= int64(x)
//         }
//     }
// }

// // 插入元素 x
// func (m *Container) add(x int) {
//     if !m.st2.IsEmpty() {
//         if first, ok := m.st2.First(); ok && x >= first {
//             m.st2.Add(x)
//         } else {
//             m.st1.Add(x)
//             m.sm += int64(x)
//         }
//     } else {
//         m.st1.Add(x)
//         m.sm += int64(x)
//     }
//     m.adjust()
// }

// // 删除元素 x
// func (m *Container) erase(x int) {
//     if m.st1.Contains(x) {
//         m.st1.Remove(x)
//         m.sm -= int64(x)
//     } else if m.st2.Contains(x) {
//         m.st2.Remove(x)
//     }
//     m.adjust()
// }

// // 前 k 小元素的和
// func (m *Container) sum() int64 {
//     return m.sm
// }

// func minimumCost(nums []int, k int, dist int) int64 {
//     n := len(nums)
//     count := NewContainer(k - 2)
//     for i := 1; i < k-1; i++ {
//         count.add(nums[i])
//     }
//     res := count.sum() + int64(nums[k-1])
//     for i := k; i < n; i++ {
//         j := i - dist - 1
//         if j > 0 {
//             count.erase(nums[j])
//         }
//         count.add(nums[i-1])
//         curr := count.sum() + int64(nums[i])     
//         if curr < res {
//             res = curr
//         }
//     }
//     return res + int64(nums[0])
// }

import "sort"
import "container/heap"

func minimumCost(nums []int, k int, dist int) int64 {
    k--
    l, r := &MaxHeap{todo: map[int]int{}}, &MaxHeap{todo: map[int]int{}}
    for _, x := range nums[1 : dist+2] {
        l.push(x)
    }
    for l.sz > k {
        r.push(-l.pop())
    }
    mn := l.sum
    for i := dist + 2; i < len(nums); i++ {
        out := nums[i-dist-1]
        if out <= l.top() {
            l.del(out)
        } else {
            r.del(-out)
        }
        in := nums[i]
        if in <= l.top() {
            l.push(in)
        } else {
            r.push(-in)
        }
        if l.sz == k-1 {
            l.push(-r.pop())
        } else if l.sz == k+1 {
            r.push(-l.pop())
        }
        mn = min(mn, l.sum)
    }
    return int64(nums[0] + mn)
}

type MaxHeap struct {
    sort.IntSlice
    todo map[int]int
    sz, sum int
}

func (h *MaxHeap) Less(i, j int) bool    { return h.IntSlice[i] > h.IntSlice[j] }
func (h *MaxHeap) Push(v any)            { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MaxHeap) Pop() (v any)          { a := h.IntSlice; h.IntSlice, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *MaxHeap) del(v int)             { h.todo[v]++; h.sz--; h.sum -= v }
func (h *MaxHeap) push(v int) { 
    if h.todo[v] > 0 {
        h.todo[v]--
    } else {
        heap.Push(h, v)
    }
    h.sz++
    h.sum += v
}

func (h *MaxHeap) pop() int { h.do(); h.sz--; v := heap.Pop(h).(int); h.sum -= v; return v }
func (h *MaxHeap) top() int { h.do(); return h.IntSlice[0] }
func (h *MaxHeap) do() {
    for h.Len() > 0 && h.todo[h.IntSlice[0]] > 0 {
        h.todo[h.IntSlice[0]]--
        heap.Pop(h)
    }
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2,6,4,2], k = 3, dist = 3
    // Output: 5
    // Explanation: The best possible way to divide nums into 3 subarrays is: [1,3], [2,6,4], and [2]. This choice is valid because ik-1 - i1 is 5 - 2 = 3 which is equal to dist. The total cost is nums[0] + nums[2] + nums[5] which is 1 + 2 + 2 = 5.
    // It can be shown that there is no possible way to divide nums into 3 subarrays at a cost lower than 5.
    fmt.Println(minimumCost([]int{1,3,2,6,4,2}, 3, 3)) // 5
    // Example 2:
    // Input: nums = [10,1,2,2,2,1], k = 4, dist = 3
    // Output: 15
    // Explanation: The best possible way to divide nums into 4 subarrays is: [10], [1], [2], and [2,2,1]. This choice is valid because ik-1 - i1 is 3 - 1 = 2 which is less than dist. The total cost is nums[0] + nums[1] + nums[2] + nums[3] which is 10 + 1 + 2 + 2 = 15.
    // The division [10], [1], [2,2,2], and [1] is not valid, because the difference between ik-1 and i1 is 5 - 1 = 4, which is greater than dist.
    // It can be shown that there is no possible way to divide nums into 4 subarrays at a cost lower than 15.
    fmt.Println(minimumCost([]int{10,1,2,2,2,1}, 4, 3)) // 15
    // Example 3:
    // Input: nums = [10,8,18,9], k = 3, dist = 1
    // Output: 36
    // Explanation: The best possible way to divide nums into 4 subarrays is: [10], [8], and [18,9]. This choice is valid because ik-1 - i1 is 2 - 1 = 1 which is equal to dist.The total cost is nums[0] + nums[1] + nums[2] which is 10 + 8 + 18 = 36.
    // The division [10], [8,18], and [9] is not valid, because the difference between ik-1 and i1 is 3 - 1 = 2, which is greater than dist.
    // It can be shown that there is no possible way to divide nums into 3 subarrays at a cost lower than 36.
    fmt.Println(minimumCost([]int{10,8,18,9}, 3, 1)) // 36

    fmt.Println(minimumCost([]int{1,2,3,4,5,6,7,8,9}, 3, 1)) // 6
    fmt.Println(minimumCost([]int{9,8,7,6,5,4,3,2,1}, 3, 1)) // 12
}