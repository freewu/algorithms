package main

// 3510. Minimum Pair Removal to Sort Array II
// Given an array nums, you can perform the following operation any number of times:
//     1. Select the adjacent pair with the minimum sum in nums. 
//        If multiple such pairs exist, choose the leftmost one.
//     2. Replace the pair with their sum.

// Return the minimum number of operations needed to make the array non-decreasing.

// An array is said to be non-decreasing if each element is greater than or equal to its previous element (if it exists).

// Example 1:
// Input: nums = [5,2,3,1]
// Output: 2
// Explanation:
// The pair (3,1) has the minimum sum of 4. After replacement, nums = [5,2,4].
// The pair (2,4) has the minimum sum of 6. After replacement, nums = [5,6].
// The array nums became non-decreasing in two operations.

// Example 2:
// Input: nums = [1,2,2]
// Output: 0
// Explanation:
// The array nums is already sorted.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"
// import "cmp"
// import "github.com/emirpasic/gods/v2/trees/redblacktree"

// func minimumPairRemoval(nums []int) int {
//     res, n := 0, len(nums)
//     type pair struct{ s, i int }
//     // (相邻元素和，左边那个数的下标)
//     pairs := redblacktree.NewWith[pair, struct{}](func(a, b pair) int { return cmp.Or(a.s-b.s, a.i-b.i) })
//     // 剩余下标
//     idx := redblacktree.New[int, struct{}]()
//     // 递减的相邻对的个数
//     dec := 0
//     for i := range n - 1 {
//         x, y := nums[i], nums[i+1]
//         if x > y {
//             dec++
//         }
//         pairs.Put(pair{x + y, i}, struct{}{})
//     }
//     for i := range n {
//         idx.Put(i, struct{}{})
//     }
//     for dec > 0 {
//         res++
//         it := pairs.Left()
//         s := it.Key.s
//         i := it.Key.i
//         pairs.Remove(it.Key) // 删除相邻元素和最小的一对
//         // 找到 i 的位置
//         node, _ := idx.Ceiling(i + 1)
//         nxt := node.Key
//         // (当前元素，下一个数)
//         if nums[i] > nums[nxt] { // 旧数据
//             dec--
//         }
//         // (前一个数，当前元素)
//         node, _ = idx.Floor(i - 1)
//         if node != nil {
//             pre := node.Key
//             if nums[pre] > nums[i] { // 旧数据
//                 dec--
//             }
//             if nums[pre] > s { // 新数据
//                 dec++
//             }
//             pairs.Remove(pair{nums[pre] + nums[i], pre})
//             pairs.Put(pair{nums[pre] + s, pre}, struct{}{})
//         }
//         // (下一个数，下下一个数)
//         node, _ = idx.Ceiling(nxt + 1)
//         if node != nil {
//             nxt2 := node.Key
//             if nums[nxt] > nums[nxt2] { // 旧数据
//                 dec--
//             }
//             if s > nums[nxt2] { // 新数据（当前元素，下下一个数）
//                 dec++
//             }
//             pairs.Remove(pair{nums[nxt] + nums[nxt2], nxt})
//             pairs.Put(pair{s + nums[nxt2], i}, struct{}{})
//         }
//         nums[i] = s
//         idx.Remove(nxt)
//     }
//     return res
// }

import "container/heap"

type Item struct {
    val   int
    left  int
    right int
    prev  *Item
    next  *Item
    index int
    order int
}

type MinHeap []*Item
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
    if h[i].val == h[j].val {
        return h[i].order < h[j].order
    }
    return h[i].val < h[j].val
}
func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
    h[i].index, h[j].index = i, j
}

func (h *MinHeap) Push(x any) {
    item := x.(*Item)
    item.index = len(*h)
    *h = append(*h, item)
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[0 : n-1]
    return item
}

func minimumPairRemoval(nums []int) int {
    if len(nums) < 2 { return 0 }
    var hp MinHeap
    var prev *Item = nil
    var notSorted, operations int
    for i := 0; i < len(nums) - 1; i++ {
        // check whether initially sorted
        if nums[i] > nums[i+1] {
            notSorted++
        }
        cur := &Item{nums[i] + nums[i+1], nums[i], nums[i+1], prev, nil, 0, i}
        if prev != nil {
            prev.next = cur
        }
        prev = cur
        hp = append(hp, cur)
    }
    for i, item := range hp {
        item.index = i
    }
    heap.Init(&hp)
    isSorted := func(a, b int) bool {  return a <= b }
    for notSorted > 0 {
        operations++
        smallest := heap.Pop(&hp).(*Item)
        if smallest.left > smallest.right {
            notSorted--
        }
        if left := smallest.prev; left != nil {
            before := isSorted(left.left, left.right)
            left.val += smallest.right
            left.right += smallest.right
            left.next = smallest.next
            after := isSorted(left.left, left.right)
            if before && !after {
                notSorted++
            } else if !before && after {
                notSorted--
            }
            heap.Fix(&hp, left.index)
        }
        if right := smallest.next; right != nil {
            before := isSorted(right.left, right.right)
            right.val += smallest.left
            right.left += smallest.left
            right.prev = smallest.prev
            after := isSorted(right.left, right.right)
            if before && !after {
                notSorted++
            } else if !before && after {
                notSorted--
            }
            heap.Fix(&hp, right.index)
        }
    }
    return operations
}

func main() {
    // Example 1:
    // Input: nums = [5,2,3,1]
    // Output: 2
    // Explanation:
    // The pair (3,1) has the minimum sum of 4. After replacement, nums = [5,2,4].
    // The pair (2,4) has the minimum sum of 6. After replacement, nums = [5,6].
    // The array nums became non-decreasing in two operations.
    fmt.Println(minimumPairRemoval([]int{5,2,3,1})) // 2
    // Example 2:
    // Input: nums = [1,2,2]
    // Output: 0
    // Explanation:
    // The array nums is already sorted.
    fmt.Println(minimumPairRemoval([]int{1,2,2})) // 0

    fmt.Println(minimumPairRemoval([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minimumPairRemoval([]int{9,8,7,6,5,4,3,2,1})) // 7
}