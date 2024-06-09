package main

// 1167. Minimum Cost to Connect Sticks
// You have some number of sticks with positive integer lengths. 
// These lengths are given as an array sticks, where sticks[i] is the length of the ith stick.

// You can connect any two sticks of lengths x and y into one stick by paying a cost of x + y. 
// You must connect all the sticks until there is only one stick remaining.

// Return the minimum cost of connecting all the given sticks into one stick in this way.

// Example 1:
// Input: sticks = [2,4,3]
// Output: 14
// Explanation: You start with sticks = [2,4,3].
// 1. Combine sticks 2 and 3 for a cost of 2 + 3 = 5. Now you have sticks = [5,4].
// 2. Combine sticks 5 and 4 for a cost of 5 + 4 = 9. Now you have sticks = [9].
// There is only one stick left, so you are done. The total cost is 5 + 9 = 14.

// Example 2:
// Input: sticks = [1,8,3,5]
// Output: 30
// Explanation: You start with sticks = [1,8,3,5].
// 1. Combine sticks 1 and 3 for a cost of 1 + 3 = 4. Now you have sticks = [4,8,5].
// 2. Combine sticks 4 and 5 for a cost of 4 + 5 = 9. Now you have sticks = [9,8].
// 3. Combine sticks 9 and 8 for a cost of 9 + 8 = 17. Now you have sticks = [17].
// There is only one stick left, so you are done. The total cost is 4 + 9 + 17 = 30.

// Example 3:
// Input: sticks = [5]
// Output: 0
// Explanation: There is only one stick, so you don't need to do anything. The total cost is 0.

// Constraints:
//     1 <= sticks.length <= 10^4
//     1 <= sticks[i] <= 10^4

import "fmt"
import "sort"

// 贪心
func connectSticks(sticks []int) int {
    if len(sticks) == 1 {
        return 0
    }
    sort.Ints(sticks)
    res := 0
    instert := func (sticks *[]int, lenght int) {
        *sticks = (*sticks)[2:]
        i := 0
        for i = 0; i < len(*sticks); i++ {
            if lenght < (*sticks)[i] {
                break
            }
        }
        if i == 0 {
            *sticks = append([]int{lenght}, *sticks...)
        } else if i == len(*sticks) {
            *sticks = append(*sticks, lenght)
        } else {
            rear := append([]int{},(*sticks)[i:]... )
            tmp := append((*sticks)[0:i],lenght )
            *sticks = append(tmp, rear...)
        }
    }
    for len(sticks) > 1 {
        res = res + sticks[0] + sticks[1]
        instert(&sticks, sticks[0] + sticks[1])
    }
    return res
}

// 小顶堆
func connectSticks1(sticks []int) int {
    if len(sticks) == 0 {
        return 0
    }
    h := NewMinHeap() // 初始化小顶堆
    for _, stick := range sticks { // 将木棍长度数组转换为小顶堆
        h.Insert(stick)
    }
    totalCost := 0 // 计算总成本
    for h.size > 1 {
        min1, min2 := h.Pop(), h.Pop() // 弹出两个最小长度的木棍
        cost := min1 + min2 // 计算连接后的总成本
        totalCost += cost
        h.Insert(cost) // 将连接后的木棍长度重新插入小顶堆中
    }
    return totalCost
}

// 建立小顶堆结构体
type MinHeap struct {
    array []int
    size  int
}

// 初始化小顶堆
func NewMinHeap() *MinHeap {
    return &MinHeap{
        array: make([]int, 0),
        size:  0,
    }
}

// 向小顶堆中插入元素
func (h *MinHeap) Insert(val int) {
    h.array = append(h.array, val)
    h.size++
    h.shiftUp(h.size - 1)
}

// 上移操作，使得小顶堆恢复性质
func (h *MinHeap) shiftUp(idx int) {
    for idx > 0 {
        parentIdx := (idx - 1) / 2
        if h.array[idx] >= h.array[parentIdx] {
            break
        }
        h.array[idx], h.array[parentIdx] = h.array[parentIdx], h.array[idx]
        idx = parentIdx
    }
}

// 弹出小顶堆中的最小元素
func (h *MinHeap) Pop() int {
    if h.size == 0 {
        return -1
    }
    minVal := h.array[0]
    h.size--
    h.array[0] = h.array[h.size]
    h.array = h.array[:h.size]
    h.shiftDown(0)
    return minVal
}

// 下移操作，使得小顶堆恢复性质
func (h *MinHeap) shiftDown(idx int) {
    for idx < h.size {
        leftChildIdx := idx*2 + 1
        rightChildIdx := idx*2 + 2
        smallestIdx := idx
        if leftChildIdx < h.size && h.array[leftChildIdx] < h.array[smallestIdx] {
            smallestIdx = leftChildIdx
        }
        if rightChildIdx < h.size && h.array[rightChildIdx] < h.array[smallestIdx] {
            smallestIdx = rightChildIdx
        }
        if smallestIdx == idx {
            break
        }
        h.array[idx], h.array[smallestIdx] = h.array[smallestIdx], h.array[idx]
        idx = smallestIdx
    }
}

func main() {
    // Example 1:
    // Input: sticks = [2,4,3]
    // Output: 14
    // Explanation: You start with sticks = [2,4,3].
    // 1. Combine sticks 2 and 3 for a cost of 2 + 3 = 5. Now you have sticks = [5,4].
    // 2. Combine sticks 5 and 4 for a cost of 5 + 4 = 9. Now you have sticks = [9].
    // There is only one stick left, so you are done. The total cost is 5 + 9 = 14.
    fmt.Println(connectSticks([]int{2,4,3})) // 14
    // Example 2:
    // Input: sticks = [1,8,3,5]
    // Output: 30
    // Explanation: You start with sticks = [1,8,3,5].
    // 1. Combine sticks 1 and 3 for a cost of 1 + 3 = 4. Now you have sticks = [4,8,5].
    // 2. Combine sticks 4 and 5 for a cost of 4 + 5 = 9. Now you have sticks = [9,8].
    // 3. Combine sticks 9 and 8 for a cost of 9 + 8 = 17. Now you have sticks = [17].
    // There is only one stick left, so you are done. The total cost is 4 + 9 + 17 = 30.
    fmt.Println(connectSticks([]int{1,8,3,5})) // 30
    // Example 3:
    // Input: sticks = [5]
    // Output: 0
    // Explanation: There is only one stick, so you don't need to do anything. The total cost is 0.
    fmt.Println(connectSticks([]int{5})) // 5

    fmt.Println(connectSticks1([]int{2,4,3})) // 14
    fmt.Println(connectSticks1([]int{1,8,3,5})) // 30
    fmt.Println(connectSticks1([]int{5})) // 5
}