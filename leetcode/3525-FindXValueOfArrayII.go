package main

// 3525. Find X Value of Array II
// You are given an array of positive integers nums and a positive integer k. 
// You are also given a 2D array queries, where queries[i] = [indexi, valuei, starti, xi].

// You are allowed to perform an operation once on nums, where you can remove any suffix from nums such that nums remains non-empty.

// The x-value of nums for a given x is defined as the number of ways to perform this operation so that the product of the remaining elements leaves a remainder of x modulo k.

// For each query in queries you need to determine the x-value of nums for xi after performing the following actions:
//     1. Update nums[indexi] to valuei. Only this step persists for the rest of the queries.
//     2. Remove the prefix nums[0..(starti - 1)] (where nums[0..(-1)] will be used to represent the empty prefix).

// Return an array result of size queries.length where result[i] is the answer for the ith query.

// A prefix of an array is a subarray that starts from the beginning of the array and extends to any point within it.

// A suffix of an array is a subarray that starts at any point within the array and extends to the end of the array.

// Note that the prefix and suffix to be chosen for the operation can be empty.

// Note that x-value has a different definition in this version.

// Example 1:
// Input: nums = [1,2,3,4,5], k = 3, queries = [[2,2,0,2],[3,3,3,0],[0,1,0,1]]
// Output: [2,2,2]
// Explanation:
// For query 0, nums becomes [1, 2, 2, 4, 5], and the empty prefix must be removed. The possible operations are:
// Remove the suffix [2, 4, 5]. nums becomes [1, 2].
// Remove the empty suffix. nums becomes [1, 2, 2, 4, 5] with a product 80, which gives remainder 2 when divided by 3.
// For query 1, nums becomes [1, 2, 2, 3, 5], and the prefix [1, 2, 2] must be removed. The possible operations are:
// Remove the empty suffix. nums becomes [3, 5].
// Remove the suffix [5]. nums becomes [3].
// For query 2, nums becomes [1, 2, 2, 3, 5], and the empty prefix must be removed. The possible operations are:
// Remove the suffix [2, 2, 3, 5]. nums becomes [1].
// Remove the suffix [3, 5]. nums becomes [1, 2, 2].

// Example 2:
// Input: nums = [1,2,4,8,16,32], k = 4, queries = [[0,2,0,2],[0,2,0,1]]
// Output: [1,0]
// Explanation:
// For query 0, nums becomes [2, 2, 4, 8, 16, 32]. The only possible operation is:
// Remove the suffix [2, 4, 8, 16, 32].
// For query 1, nums becomes [2, 2, 4, 8, 16, 32]. There is no possible way to perform the operation.

// Example 3:
// Input: nums = [1,1,2,1,1], k = 2, queries = [[2,1,0,1]]
// Output: [5]

// Constraints:
//     1 <= nums[i] <= 10^9
//     1 <= nums.length <= 10^5
//     1 <= k <= 5
//     1 <= queries.length <= 2 * 10^4
//     queries[i] == [indexi, valuei, starti, xi]
//     0 <= indexi <= nums.length - 1
//     1 <= valuei <= 10^9
//     0 <= starti <= nums.length - 1
//     0 <= xi <= k - 1

import "fmt"
import "math/bits"

var K int

type data struct {
    mul int
    cnt [5]int // 比 []int 快
}

type SegmentTree []data

func mergeData(a, b data) data {
    cnt := a.cnt
    for rx, c := range b.cnt {
        cnt[a.mul * rx % K] += c
    }
    return data{a.mul * b.mul % K, cnt}
}

func newData(val int) data {
    mul := val % K
    cnt := [5]int{}
    cnt[mul] = 1
    return data{mul, cnt}
}

func (t SegmentTree) maintain(o int) {
    t[o] = mergeData(t[o<<1], t[o<<1|1])
}

func (t SegmentTree) build(a []int, o, l, r int) {
    if l == r {
        t[o] = newData(a[l])
        return
    }
    m := (l + r) >> 1
    t.build(a, o<<1, l, m)
    t.build(a, o<<1|1, m+1, r)
    t.maintain(o)
}

func (t SegmentTree) update(o, l, r, i, val int) {
    if l == r {
        t[o] = newData(val)
        return
    }
    m := (l + r) >> 1
    if i <= m {
        t.update(o<<1, l, m, i, val)
    } else {
        t.update(o<<1|1, m+1, r, i, val)
    }
    t.maintain(o)
}

func (t SegmentTree) query(o, l, r, ql, qr int) data {
    if ql <= l && r <= qr {
        return t[o]
    }
    m := (l + r) / 2
    if qr <= m {
        return t.query(o*2, l, m, ql, qr)
    }
    if ql > m {
        return t.query(o*2+1, m+1, r, ql, qr)
    }
    lRes := t.query(o*2, l, m, ql, qr)
    rRes := t.query(o*2+1, m+1, r, ql, qr)
    return mergeData(lRes, rRes)
}

func NewSegmentTree(arr []int) SegmentTree {
    n := len(arr)
    t := make(SegmentTree, 2 << bits.Len(uint(n-1)))
    t.build(arr, 1, 0, n-1)
    return t
}

func resultArray(nums []int, k int, queries [][]int) []int {
    K = k
    t := NewSegmentTree(nums)
    n := len(nums)
    res := make([]int, len(queries))
    for qi, q := range queries {
        t.update(1, 0, n-1, q[0], q[1])
        v := t.query(1, 0, n-1, q[2], n-1)
        res[qi] = v.cnt[q[3]]
    }
    return res
}

func resultArray1(nums []int, k int, queries [][]int) []int {
    mod := make([]int, len(nums))
    for i := range nums {
        mod[i] = nums[i] % k
    }
    res := make([]int, len(queries))
    for qidx, q := range queries {
        index, value, start, xi := q[0], q[1], q[2], q[3]
        nums[index], mod[index] = value, value % k
        m := len(nums) - start
        if m <= 0 {
            res[qidx] = 0
            continue
        }
        if k == 1 {
            if xi == 0 {
                res[qidx] = m
            } else {
                res[qidx] = 0
            }
            continue
        }
        currentMod, count := 1, 0
        for i := 0; i < m; i++ {
            currentMod = (currentMod * mod[start+i]) % k
            if currentMod == xi {
                count++
            }
            if currentMod == 0 {
                if xi == 0 {
                    count += m - i - 1
                }
                break
            }
        }
        res[qidx] = count
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,5], k = 3, queries = [[2,2,0,2],[3,3,3,0],[3,3,3,0]]
    // Output: [2,2,2]
    // Explanation:
    // For query 0, nums becomes [1, 2, 2, 4, 5], and the empty prefix must be removed. The possible operations are:
    // Remove the suffix [2, 4, 5]. nums becomes [1, 2].
    // Remove the empty suffix. nums becomes [1, 2, 2, 4, 5] with a product 80, which gives remainder 2 when divided by 3.
    // For query 1, nums becomes [1, 2, 2, 3, 5], and the prefix [1, 2, 2] must be removed. The possible operations are:
    // Remove the empty suffix. nums becomes [3, 5].
    // Remove the suffix [5]. nums becomes [3].
    // For query 2, nums becomes [1, 2, 2, 3, 5], and the empty prefix must be removed. The possible operations are:
    // Remove the suffix [2, 2, 3, 5]. nums becomes [1].
    // Remove the suffix [3, 5]. nums becomes [1, 2, 2].
    fmt.Println(resultArray([]int{1,2,3,4,5}, 3, [][]int{{2,2,0,2},{3,3,3,0},{3,3,3,0}})) // [2,2,2]
    // Example 2:
    // Input: nums = [1,2,4,8,16,32], k = 4, queries = [[0,2,0,2],[0,2,0,1]]
    // Output: [1,0]
    // Explanation:
    // For query 0, nums becomes [2, 2, 4, 8, 16, 32]. The only possible operation is:
    // Remove the suffix [2, 4, 8, 16, 32].
    // For query 1, nums becomes [2, 2, 4, 8, 16, 32]. There is no possible way to perform the operation.
    fmt.Println(resultArray([]int{1,2,4,8,16,32}, 4, [][]int{{0,2,0,2},{0,2,0,1}})) // [1,0]
    // Example 3:
    // Input: nums = [1,1,2,1,1], k = 2, queries = [[2,1,0,1]]
    // Output: [5]
    fmt.Println(resultArray([]int{1,1,2,1,1}, 2, [][]int{{2,1,0,1}})) // [5]

    fmt.Println(resultArray([]int{1,2,3,4,5,6,7,8,9}, 2, [][]int{{2,1,0,1}})) // [1]
    fmt.Println(resultArray([]int{9,8,7,6,5,4,3,2,1}, 2, [][]int{{2,1,0,1}})) // [1]

    fmt.Println(resultArray1([]int{1,2,3,4,5}, 3, [][]int{{2,2,0,2},{3,3,3,0},{3,3,3,0}})) // [2,2,2]
    fmt.Println(resultArray1([]int{1,2,4,8,16,32}, 4, [][]int{{0,2,0,2},{0,2,0,1}})) // [1,0]
    fmt.Println(resultArray1([]int{1,1,2,1,1}, 2, [][]int{{2,1,0,1}})) // [5]
    fmt.Println(resultArray1([]int{1,2,3,4,5,6,7,8,9}, 2, [][]int{{2,1,0,1}})) // [1]
    fmt.Println(resultArray1([]int{9,8,7,6,5,4,3,2,1}, 2, [][]int{{2,1,0,1}})) // [1]
}