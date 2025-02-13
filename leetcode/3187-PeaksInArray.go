package main

// 3187. Peaks in Array
// A peak in an array arr is an element that is greater than its previous and next element in arr.

// You are given an integer array nums and a 2D integer array queries.

// You have to process queries of two types:
//     1. queries[i] = [1, li, ri], determine the count of peak elements in the subarray nums[li..ri].
//     2. queries[i] = [2, indexi, vali], change nums[indexi] to vali.

// Return an array answer containing the results of the queries of the first type in order.

// Notes:
//     The first and the last element of an array or a subarray cannot be a peak.

// Example 1:
// Input: nums = [3,1,4,2,5], queries = [[2,3,4],[1,0,4]]
// Output: [0]
// Explanation:
// First query: We change nums[3] to 4 and nums becomes [3,1,4,4,5].
// Second query: The number of peaks in the [3,1,4,4,5] is 0.

// Example 2:
// Input: nums = [4,1,4,2,1,5], queries = [[2,2,4],[1,0,2],[1,0,4]]
// Output: [0,1]
// Explanation:
// First query: nums[2] should become 4, but it is already set to 4.
// Second query: The number of peaks in the [4,1,4] is 0.
// Third query: The second 4 is a peak in the [4,1,4,2,1].

// Constraints:
//     3 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= queries.length <= 10^5
//     queries[i][0] == 1 or queries[i][0] == 2
//     For all i that:
//         queries[i][0] == 1: 0 <= queries[i][1] <= queries[i][2] <= nums.length - 1
//         queries[i][0] == 2: 0 <= queries[i][1] <= nums.length - 1, 1 <= queries[i][2] <= 10^5

import "fmt"

func build(idx, low, high int, arr, seg []int) {
    if low == high {
        seg[idx] = arr[low]
        return
    }
    mid := (low + high) >> 1
    build(2*idx, low, mid, arr, seg)
    build(2*idx+1, mid+1, high, arr, seg)
    seg[idx] = seg[2*idx] + seg[2*idx+1]
}

func query(idx, low, high, l, r int, seg []int) int {
    if low > r || high < l { return 0 }
    if l <= low && high <= r { return seg[idx] }
    mid := (low + high) >> 1
    return query(2*idx, low, mid, l, r, seg) + query(2*idx+1, mid+1, high, l, r, seg)
}

func update(idx, low, high, pos, val int, seg []int) {
    if pos < low || pos > high { return }
    if low == high {
        seg[idx] = val
        return
    }
    mid := (low + high) >> 1
    update(2*idx, low, mid, pos, val, seg)
    update(2*idx+1, mid+1, high, pos, val, seg)
    seg[idx] = seg[2*idx] + seg[2*idx+1]
}

// Segment Tree
func countOfPeaks(nums []int, queries [][]int) []int {
    n := len(nums)
    res, arr, seg := []int{}, make([]int, n), make([]int, n * 4)
    for i := 1; i < n-1; i++ {
        if nums[i] > nums[i+1] && nums[i] > nums[i - 1] {
            arr[i] = 1
        }
    }
    build(1, 0, n-1, arr, seg)
    for _, q := range queries {
        if q[0] == 1 {
            l, r := q[1], q[2]
            res = append(res, query(1, 0, n-1, l+1, r-1, seg))
        } else {
            pos, val := q[1], q[2]
            nums[pos] = val
            for i := pos-1; i <= pos+1; i++ {
                if i > 0 && i < n-1 {
                    if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
                        arr[i] = 1
                    } else {
                        arr[i] = 0
                    }
                    update(1, 0, n-1, i, arr[i], seg)
                }
            }
        }
    }
    return res
}

type BinaryIndexedTree struct {
    tree []int
}

func NewBinaryIndexedTree(n int) *BinaryIndexedTree {
    return &BinaryIndexedTree{make([]int, n)}
}

func (t *BinaryIndexedTree) modify(i, v int) {
    for ; i < len(t.tree); i += i & -i {
        t.tree[i] += v
    }
}

func (t *BinaryIndexedTree) query(i int) int {
    res := 0
    for ; i > 0; i -= i & -i {
        res += t.tree[i]
    }
    return res
}

func countOfPeaks1(nums []int, queries [][]int) []int {
    n := len(nums)
    bit := NewBinaryIndexedTree(n - 1)
    modify := func(i, v int) {
        if i <= 0 || i >= n - 1 { return }
        if nums[i - 1] < nums[i] && nums[i] > nums[i + 1] {
            bit.modify(i, v)
        }
    }
    for i := range nums {
        modify(i, 1)
    }
    res := []int{}
    for _, q := range queries {
        if q[0] == 1 {
            l, r := q[1], q[2] - 1
            t := 0
            if l + 1 <= r {
                t = bit.query(r) - bit.query(l)
            }
            res = append(res, t)
        } else {
            i, val := q[1], q[2]
            for k := i - 1; k <= i + 1; k++ {
                modify(k, -1)
            }
            nums[i] = val
            for k := i - 1; k <= i + 1; k++ {
                modify(k, 1)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,4,2,5], queries = [[2,3,4],[1,0,4]]
    // Output: [0]
    // Explanation:
    // First query: We change nums[3] to 4 and nums becomes [3,1,4,4,5].
    // Second query: The number of peaks in the [3,1,4,4,5] is 0.
    fmt.Println(countOfPeaks([]int{3,1,4,2,5}, [][]int{{2,3,4},{1,0,4}})) //  [0]
    // Example 2:
    // Input: nums = [4,1,4,2,1,5], queries = [[2,2,4],[1,0,2],[1,0,4]]
    // Output: [0,1]
    // Explanation:
    // First query: nums[2] should become 4, but it is already set to 4.
    // Second query: The number of peaks in the [4,1,4] is 0.
    // Third query: The second 4 is a peak in the [4,1,4,2,1].
    fmt.Println(countOfPeaks([]int{4,1,4,2,1,5}, [][]int{{2,2,4},{1,0,2},{1,0,4}})) // [0,1]

    fmt.Println(countOfPeaks1([]int{3,1,4,2,5}, [][]int{{2,3,4},{1,0,4}})) //  [0]
    fmt.Println(countOfPeaks1([]int{4,1,4,2,1,5}, [][]int{{2,2,4},{1,0,2},{1,0,4}})) // [0,1]
}