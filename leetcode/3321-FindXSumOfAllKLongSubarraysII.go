package main

// 3321. Find X-Sum of All K-Long Subarrays II
// You are given an array nums of n integers and two integers k and x.

// The x-sum of an array is calculated by the following procedure:
//     1. Count the occurrences of all elements in the array.
//     2. Keep only the occurrences of the top x most frequent elements. 
//        If two elements have the same number of occurrences, the element with the bigger value is considered more frequent.
//     3. Calculate the sum of the resulting array.

// Note that if an array has less than x distinct elements, its x-sum is the sum of the array.

// Return an integer array answer of length n - k + 1 where answer[i] is the x-sum of the subarray nums[i..i + k - 1].

// Example 1:
// Input: nums = [1,1,2,2,3,4,2,3], k = 6, x = 2
// Output: [6,10,12]
// Explanation:
// For subarray [1, 1, 2, 2, 3, 4], only elements 1 and 2 will be kept in the resulting array. Hence, answer[0] = 1 + 1 + 2 + 2.
// For subarray [1, 2, 2, 3, 4, 2], only elements 2 and 4 will be kept in the resulting array. Hence, answer[1] = 2 + 2 + 2 + 4. Note that 4 is kept in the array since it is bigger than 3 and 1 which occur the same number of times.
// For subarray [2, 2, 3, 4, 2, 3], only elements 2 and 3 are kept in the resulting array. Hence, answer[2] = 2 + 2 + 2 + 3 + 3.

// Example 2:
// Input: nums = [3,8,7,8,7,5], k = 2, x = 2
// Output: [11,15,15,15,12]
// Explanation:
// Since k == x, answer[i] is equal to the sum of the subarray nums[i..i + k - 1].

// Constraints:
//     nums.length == n
//     1 <= n <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= x <= k <= nums.length

import "fmt"
import "container/heap"

type Item struct{ Count, Value, Version int } // 出现次数，元素值, 版本号

type MaxHeap struct {
    Data  []Item
    Order int
}

func (h *MaxHeap) Len() int { return len(h.Data) }
func (h *MaxHeap) Swap(i, j int) { h.Data[i], h.Data[j] = h.Data[j], h.Data[i]}
func (h *MaxHeap) Less(i, j int) bool {
    less := func (a, b Item) bool { return a.Count < b.Count || a.Count == b.Count && a.Value < b.Value }
    b := less(h.Data[i], h.Data[j])
    if h.Order > 0 { return !b }
    return b
}
func (h *MaxHeap) Push(v any) { h.Data = append(h.Data, v.(Item)) }
func (h *MaxHeap) Pop() any {
    n := h.Len()
    res := h.Data[n-1]
    h.Data = h.Data[:n-1]
    return res
}

func findXSum(nums []int, k, x int) []int64 {
    res, mp := []int64{},make(map[int]Item)
    l, r := &MaxHeap{}, &MaxHeap{ Order: 1 }
    sum, lsize := 0, 0
    clear := func(h *MaxHeap) {
        for h.Len() > 0 {
            top := h.Data[0]
            t := mp[top.Value]
            if h.Data[0].Count > 0 && h.Data[0].Version == t.Version { break }
            heap.Pop(h)
        }
    }
    less := func (a, b Item) bool { return a.Count < b.Count || a.Count == b.Count && a.Value < b.Value }
    rb := func() {
        clear(r)
        for lsize < x && r.Len() > 0 {
            t := heap.Pop(r).(Item)
            clear(r)
            heap.Push(l, t)
            lsize++
            sum += t.Count * t.Value
        }
        clear(l)
        clear(r)
        for l.Len() > 0 && r.Len() > 0 && less(l.Data[0], r.Data[0]) {
            t := heap.Pop(l).(Item)
            clear(l)
            heap.Push(r, t)
            sum -= t.Count * t.Value
            t = heap.Pop(r).(Item)
            clear(r)
            heap.Push(l, t)
            sum += t.Count * t.Value
        }
    }
    for i, v := range nums {
        if i >= k {
            v := nums[i-k]
            t := mp[v]
            if l.Len() > 0 && !less(t, l.Data[0]) {
                lsize--
                sum -= t.Count * t.Value
            }
            t.Count--
            t.Version++
            mp[v] = t
            heap.Push(r, t)
            rb()
        }
        t, ok := mp[v]
        if !ok {
            t = Item{ Value: v }
        }
        clear(l)
        if l.Len() > 0 && !less(t, l.Data[0]) {
            lsize--
            sum -= t.Count * t.Value
        }
        t.Count++
        t.Version++
        mp[v] = t
        heap.Push(r, t)
        rb()
        if i >= k-1 {
            res = append(res, int64(sum))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,1,2,2,3,4,2,3], k = 6, x = 2
    // Output: [6,10,12]
    // Explanation:
    // For subarray [1, 1, 2, 2, 3, 4], only elements 1 and 2 will be kept in the resulting array. Hence, answer[0] = 1 + 1 + 2 + 2.
    // For subarray [1, 2, 2, 3, 4, 2], only elements 2 and 4 will be kept in the resulting array. Hence, answer[1] = 2 + 2 + 2 + 4. Note that 4 is kept in the array since it is bigger than 3 and 1 which occur the same number of times.
    // For subarray [2, 2, 3, 4, 2, 3], only elements 2 and 3 are kept in the resulting array. Hence, answer[2] = 2 + 2 + 2 + 3 + 3.
    fmt.Println(findXSum([]int{1,1,2,2,3,4,2,3}, 6, 2)) // [6,10,12]
    // Example 2:
    // Input: nums = [3,8,7,8,7,5], k = 2, x = 2
    // Output: [11,15,15,15,12]
    // Explanation:
    // Since k == x, answer[i] is equal to the sum of the subarray nums[i..i + k - 1].
    fmt.Println(findXSum([]int{3,8,7,8,7,5}, 2, 2)) // [11,15,15,15,12]

    fmt.Println(findXSum([]int{1,2,3,4,5,6,7,8,9}, 2, 2)) // [3 5 7 9 11 13 15 17]
    fmt.Println(findXSum([]int{9,8,7,6,5,4,3,2,1}, 2, 2)) // [17 15 13 11 9 7 5 3]
}