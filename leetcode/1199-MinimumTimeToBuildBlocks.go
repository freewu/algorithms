package main

// 1199. Minimum Time to Build Blocks
// You are given a list of blocks, where blocks[i] = t means that the i-th block needs t units of time to be built. 
// A block can only be built by exactly one worker.

// A worker can either split into two workers (number of workers increases by one) or build a block then go home. 
// Both decisions cost some time.

// The time cost of spliting one worker into two workers is given as an integer split. 
// Note that if two workers split at the same time, they split in parallel so the cost would be split.

// Output the minimum time needed to build all blocks.

// Initially, there is only one worker.

// Example 1:
// Input: blocks = [1], split = 1
// Output: 1
// Explanation: We use 1 worker to build 1 block in 1 time unit.

// Example 2:
// Input: blocks = [1,2], split = 5
// Output: 7
// Explanation: We split the worker into 2 workers in 5 time units then assign each of them to a block so the cost is 5 + max(1, 2) = 7.

// Example 3:
// Input: blocks = [1,2,3], split = 1
// Output: 4
// Explanation: Split 1 worker into 2, then assign the first worker to the last block and split the second worker into 2.
// Then, use the two unassigned workers to build the first two blocks.
// The cost is 1 + max(3, 1 + max(1, 2)) = 4.

// Constraints:
//     1 <= blocks.length <= 1000
//     1 <= blocks[i] <= 10^5
//     1 <= split <= 100

import "fmt"
import "container/heap"
import "sort"

type PriorityQueue []int

func (this PriorityQueue) Len() int {return len(this)}
func (this PriorityQueue) Less(a, b int) bool {  return this[a] < this[b] }
func (this PriorityQueue) Swap(a, b int) { this[a], this[b] = this[b], this[a] }
func (this *PriorityQueue) Push(x interface{}) { *this = append(*this, x.(int)) }
func (this *PriorityQueue) Pop() interface{} {
    res := (*this)[len(*this) - 1]
    (*this) = (*this)[:this.Len() - 1]
    return res
}

func minBuildTime(blocks []int, split int) int {
    queue := &PriorityQueue{}
    heap.Init(queue)
    for _, v := range blocks {
        heap.Push(queue, v)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for queue.Len() > 1 {
        cn1 := heap.Pop(queue).(int)
        cn2 := heap.Pop(queue).(int)

        heap.Push(queue, max(cn1, cn2) + split)
    }
    return heap.Pop(queue).(int)
}

func minBuildTime1(blocks []int, split int) int {
    sort.Ints(blocks)
    check := func(blocks []int, cost, limit int) bool {
        n := len(blocks)
        queue := make([]int, n * 2 - 1)
        queue[0] = 0
        l, r := 0, 1
        for i := n - 1; i >= 0; i-- {
            if l == r {
                return false
            }
            for r - l <= i && queue[l] + cost + blocks[i] <= limit {
                queue[r] = queue[l] + cost
                r++
                queue[r] = queue[l] + cost
                r++
                l++
            }
            if queue[l] + blocks[i] > limit {
                return false
            }
            l++
        }
        return true
    }
    left, right:= 1, 1000000000
    for left < right {
        m := (left + right) / 2
        if check(blocks, split, m) {
            right = m
        } else {
            left = m + 1
        }
    }
    return left
}

func main() {
    // Example 1:
    // Input: blocks = [1], split = 1
    // Output: 1
    // Explanation: We use 1 worker to build 1 block in 1 time unit.
    fmt.Println(minBuildTime([]int{1}, 1)) // 1
    // Example 2:
    // Input: blocks = [1,2], split = 5
    // Output: 7
    // Explanation: We split the worker into 2 workers in 5 time units then assign each of them to a block so the cost is 5 + max(1, 2) = 7.
    fmt.Println(minBuildTime([]int{1,2}, 5)) // 7
    // Example 3:
    // Input: blocks = [1,2,3], split = 1
    // Output: 4
    // Explanation: Split 1 worker into 2, then assign the first worker to the last block and split the second worker into 2.
    // Then, use the two unassigned workers to build the first two blocks.
    // The cost is 1 + max(3, 1 + max(1, 2)) = 4.
    fmt.Println(minBuildTime([]int{1,2,3}, 1)) // 4

    fmt.Println(minBuildTime1([]int{1}, 1)) // 1
    fmt.Println(minBuildTime1([]int{1,2}, 5)) // 7
    fmt.Println(minBuildTime1([]int{1,2,3}, 1)) // 4
}