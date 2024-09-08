package main

// 3066. Minimum Operations to Exceed Threshold Value II
// You are given a 0-indexed integer array nums, and an integer k.

// In one operation, you will:
//     Take the two smallest integers x and y in nums.
//     Remove x and y from nums.
//     Add min(x, y) * 2 + max(x, y) anywhere in the array.

// Note that you can only apply the described operation if nums contains at least two elements.

// Return the minimum number of operations needed so that all elements of the array are greater than or equal to k.

// Example 1:
// Input: nums = [2,11,10,1,3], k = 10
// Output: 2
// Explanation: In the first operation, we remove elements 1 and 2, then add 1 * 2 + 2 to nums. nums becomes equal to [4, 11, 10, 3].
// In the second operation, we remove elements 3 and 4, then add 3 * 2 + 4 to nums. nums becomes equal to [10, 11, 10].
// At this stage, all the elements of nums are greater than or equal to 10 so we can stop.
// It can be shown that 2 is the minimum number of operations needed so that all elements of the array are greater than or equal to 10.

// Example 2:
// Input: nums = [1,1,2,4,9], k = 20
// Output: 4
// Explanation: After one operation, nums becomes equal to [2, 4, 9, 3].
// After two operations, nums becomes equal to [7, 4, 9].
// After three operations, nums becomes equal to [15, 9].
// After four operations, nums becomes equal to [33].
// At this stage, all the elements of nums are greater than 20 so we can stop.
// It can be shown that 4 is the minimum number of operations needed so that all elements of the array are greater than or equal to 20.

// Constraints:
//     2 <= nums.length <= 2 * 10^5
//     1 <= nums[i] <= 10^9
//     1 <= k <= 10^9
//     The input is generated such that an answer always exists. That is, there exists some sequence of operations after which all elements of the array are greater than or equal to k.

import "fmt"
import "container/heap"
import "sort"

// Priority Queue 
type PriorityQueue []int

func (this PriorityQueue) isEmpty() bool { return this.Len() == 0 }
func (this PriorityQueue) Len() int { return len(this) }
func (this PriorityQueue) Less(i, j int) bool { return this[i] < this[j] }// min-heap
func (this PriorityQueue) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
func (this PriorityQueue) Peek() int { return this[0] }
func (this *PriorityQueue) Push(num interface{}) { *this = append(*this, num.(int))}
func (this *PriorityQueue) Pop() interface{} {
    n := len(*this)
    num := (*this)[n - 1]
    *this = (*this)[:n - 1]
    return num
}

func minOperations(nums []int, k int) int {
    res, pq := 0, new(PriorityQueue)
    heap.Init(pq)
    for _, v := range nums {
        heap.Push(pq, v)
    }
    for pq.Peek() < k {
        if pq.Peek() < k {
            one := heap.Pop(pq).(int)
            two := heap.Pop(pq).(int)
            heap.Push(pq, one * 2 + two)
            res++
        }
    }
    return res
}

func minOperations1(nums []int, k int) int {
    n, size, count, d, r, minNum := len(nums), 0, 0, k / 3, k % 3, k
    arr := make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range nums {
        if v - r < d {
            arr[size] = v
            size++
        } else if v < k {
            count++
            minNum = min(minNum, v)
        }
    }
    sort.Ints(arr[:size])
    queue := make([]int, size)
    write, read, index, mn, mx, steps := 0, 0, 0, 0, 0, 0
    for index < size || read < write {
        if index < size && (read == write || arr[index] < queue[read]) {
            mn = arr[index]
            index++
        } else {
            mn = queue[read]
            read++
        }
        steps++
        if index == size && read == write {
            x := mn * 2 + minNum
            if x >= k { count-- }
            break
        }
        if index < size && (read == write || arr[index] < queue[read]) {
            mx = arr[index]
            index++
        } else {
            mx = queue[read]
            read++
        }
        x := mn * 2 + mx
        if x - r < d {
            queue[write] = x
            write++
        } else if x < k {
            count++
            minNum = min(minNum, x)
        }
    }
    return steps + (count + 1) / 2
}

func main() {
    // Example 1:
    // Input: nums = [2,11,10,1,3], k = 10
    // Output: 2
    // Explanation: In the first operation, we remove elements 1 and 2, then add 1 * 2 + 2 to nums. nums becomes equal to [4, 11, 10, 3].
    // In the second operation, we remove elements 3 and 4, then add 3 * 2 + 4 to nums. nums becomes equal to [10, 11, 10].
    // At this stage, all the elements of nums are greater than or equal to 10 so we can stop.
    // It can be shown that 2 is the minimum number of operations needed so that all elements of the array are greater than or equal to 10.
    fmt.Println(minOperations([]int{2,11,10,1,3}, 10)) // 2
    // Example 2:
    // Input: nums = [1,1,2,4,9], k = 20
    // Output: 4
    // Explanation: After one operation, nums becomes equal to [2, 4, 9, 3].
    // After two operations, nums becomes equal to [7, 4, 9].
    // After three operations, nums becomes equal to [15, 9].
    // After four operations, nums becomes equal to [33].
    // At this stage, all the elements of nums are greater than 20 so we can stop.
    // It can be shown that 4 is the minimum number of operations needed so that all elements of the array are greater than or equal to 20.
    fmt.Println(minOperations([]int{1,1,2,4,9}, 20)) // 4

    fmt.Println(minOperations1([]int{2,11,10,1,3}, 10)) // 2
    fmt.Println(minOperations1([]int{1,1,2,4,9}, 20)) // 4
}