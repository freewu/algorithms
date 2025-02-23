package main

// 3092. Most Frequent IDs
// The problem involves tracking the frequency of IDs in a collection that changes over time. 
// You have two integer arrays, nums and freq, of equal length n. 
// Each element in nums represents an ID, and the corresponding element in freq indicates how many times that ID should be added to or removed from the collection at each step.
//     1. Addition of IDs: If freq[i] is positive, it means freq[i] IDs with the value nums[i] are added to the collection at step i.
//     2. Removal of IDs: If freq[i] is negative, it means -freq[i] IDs with the value nums[i] are removed from the collection at step i.

// Return an array ans of length n, where ans[i] represents the count of the most frequent ID in the collection after the ith step. 
// If the collection is empty at any step, ans[i] should be 0 for that step.

// Example 1:
// Input: nums = [2,3,2,1], freq = [3,2,-3,1]
// Output: [3,3,2,2]
// Explanation:
// After step 0, we have 3 IDs with the value of 2. So ans[0] = 3.
// After step 1, we have 3 IDs with the value of 2 and 2 IDs with the value of 3. So ans[1] = 3.
// After step 2, we have 2 IDs with the value of 3. So ans[2] = 2.
// After step 3, we have 2 IDs with the value of 3 and 1 ID with the value of 1. So ans[3] = 2.

// Example 2:
// Input: nums = [5,5,3], freq = [2,-2,1]
// Output: [2,0,1]
// Explanation:
// After step 0, we have 2 IDs with the value of 5. So ans[0] = 2.
// After step 1, there are no IDs. So ans[1] = 0.
// After step 2, we have 1 ID with the value of 3. So ans[2] = 1.

// Constraints:
//     1 <= nums.length == freq.length <= 10^5
//     1 <= nums[i] <= 10^5
//     -10^5 <= freq[i] <= 10^5
//     freq[i] != 0
//     The input is generated such that the occurrences of an ID will not be negative in any step.

import "fmt"
import "container/heap"

type Pair struct {
    Count int64
    Val int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].Count > h[j].Count }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MaxHeap) Pop() interface{}   {
    tmp := *h
    n := len(tmp)
    res := tmp[n - 1]
    *h = tmp[0 : n - 1]
    return res
}

func mostFrequentIDs(nums []int, freq []int) []int64 {
    n := len(nums)
    pq := make(MaxHeap, 0)
    heap.Init(&pq)
    res, mp := make([]int64, n), make(map[int]int64)
    for i := 0; i < n; i++ {
        v := nums[i]
        mp[v] += int64(freq[i])
        heap.Push(&pq, Pair{ mp[v], v })
        for len(pq) > 0 {
            count, val := pq[0].Count, pq[0].Val
            if count == mp[val] { break }
            heap.Pop(&pq)
        }
        if len(pq) == 0 {
            res[i] = 0
        } else {
            res[i] = pq[0].Count
        }
    }
    return res
}

func mostFrequentIDs1(nums []int, freq []int) []int64 {
    mx := -1
    for _, v := range nums {
        mx = max(mx, v)
    }
    res, tr := make([]int64, len(freq)), make([]int64, mx * 4 + 100)
    var update func(l, r, i int, j, x int)
    update = func(l, r, i int, j, x int) {
        if l == j && j == r {
            tr[i] += int64(x)
            return
        } 
        mid := (l + r) >> 1
        if j <= mid {
            update(l, mid, i * 2 + 1, j, x)
        } else {
            update(mid + 1, r, i * 2 + 2, j, x)
        }
        tr[i] = max(tr[i * 2 + 1], tr[i * 2 + 2])
    }
    for i := 0; i < len(nums); i++ {
        id, cv:= nums[i], freq[i]
        update(1, mx, 0, id, cv)
        res[i] = tr[0]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,3,2,1], freq = [3,2,-3,1]
    // Output: [3,3,2,2]
    // Explanation:
    // After step 0, we have 3 IDs with the value of 2. So ans[0] = 3.
    // After step 1, we have 3 IDs with the value of 2 and 2 IDs with the value of 3. So ans[1] = 3.
    // After step 2, we have 2 IDs with the value of 3. So ans[2] = 2.
    // After step 3, we have 2 IDs with the value of 3 and 1 ID with the value of 1. So ans[3] = 2.
    fmt.Println(mostFrequentIDs([]int{2,3,2,1}, []int{3,2,-3,1})) // [3,3,2,2]
    // Example 2:
    // Input: nums = [5,5,3], freq = [2,-2,1]
    // Output: [2,0,1]
    // Explanation:
    // After step 0, we have 2 IDs with the value of 5. So ans[0] = 2.
    // After step 1, there are no IDs. So ans[1] = 0.
    // After step 2, we have 1 ID with the value of 3. So ans[2] = 1.
    fmt.Println(mostFrequentIDs([]int{5,5,3}, []int{2,-2,1})) // [2,0,1]

    fmt.Println(mostFrequentIDs1([]int{2,3,2,1}, []int{3,2,-3,1})) // [3,3,2,2]
    fmt.Println(mostFrequentIDs1([]int{5,5,3}, []int{2,-2,1})) // [2,0,1]
}