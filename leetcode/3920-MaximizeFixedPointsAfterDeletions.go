package main

// 3920. Maximize Fixed Points After Deletions
// You are given an integer array nums.

// A position i is called a fixed point if nums[i] == i.

// You are allowed to delete any number of elements (including zero) from the array. 
// After each deletion, the remaining elements shift left, and indices are reassigned starting from 0.

// Return an integer denoting the maximum number of fixed points that can be achieved after performing any number of deletions.

// Example 1:
// Input: nums = [0,2,1]
// Output: 2
// Explanation:
// Delete nums[1] = 2. The array becomes [0, 1].
// Now, nums[0] = 0 and nums[1] = 1, so both indices are fixed points.
// Thus, the answer is 2.

// Example 2:
// Input: nums = [3,1,2]
// Output: 2
// Explanation:
// Do not delete any elements. The array remains [3, 1, 2].
// Here, nums[1] = 1 and nums[2] = 2, so these indices are fixed points.
// Thus, the answer is 2.

// Example 3:
// Input: nums = [1,0,1,2]
// Output: 3
// Explanation:
// Delete nums[0] = 1. The array becomes [0, 1, 2].
// Now, nums[0] = 0, nums[1] = 1, and nums[2] = 2, so all indices are fixed points.
// Thus, the answer is 3.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"
import "sort"
import "slices"
import "cmp"

func maxFixedPoints(nums []int) int {
    arr := [][2]int{}
    for i, v := range nums {
        if i >= v {
            arr = append(arr, [2]int{v, i - v})
        }
    }
    maxEnvelopes := func(envelopes [][2]int) int {
        slices.SortFunc(envelopes, func(a, b [2]int) int {
            return cmp.Or(a[0] - b[0], b[1] - a[1])
        })
        g := []int{}
        for _, e := range envelopes {
            h := e[1]
            j := sort.SearchInts(g, h+1) // 允许 LIS 相邻元素相等
            if j < len(g) {
                g[j] = h
            } else {
                g = append(g, h)
            }
        }
        return len(g)
    }
    return maxEnvelopes(arr)
}

func maxFixedPoints1(nums []int) int {
    res, good, n := 0, 0, len(nums)
    for i, v := range nums {
        if i >= v {
            nums[good] = v<<20 | (100000 - i + v)
            good++
        }
    }
    slices.Sort(nums[:good])
    bit := make([]int, n + 1)
    update := func(pos, val int) {
        for pos++; pos <= n; pos += pos & -pos {
            bit[pos] = max(bit[pos], val)
        }
    }
    query := func(pos int) int {
        res := 0
        for pos++; pos > 0; pos -= pos & -pos {
            res = max(res, bit[pos])
        }
        return res
    }
    for _, p := range nums[:good] {
        x := 100000 - (p & (1<<20 - 1))
        best := query(x) + 1
        update(x, best)
        res = max(res, best)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,2,1]
    // Output: 2
    // Explanation:
    // Delete nums[1] = 2. The array becomes [0, 1].
    // Now, nums[0] = 0 and nums[1] = 1, so both indices are fixed points.
    // Thus, the answer is 2.
    fmt.Println(maxFixedPoints([]int{0,2,1})) // 2
    // Example 2:
    // Input: nums = [3,1,2]
    // Output: 2
    // Explanation:
    // Do not delete any elements. The array remains [3, 1, 2].
    // Here, nums[1] = 1 and nums[2] = 2, so these indices are fixed points.
    // Thus, the answer is 2.
    fmt.Println(maxFixedPoints([]int{3,1,2})) // 2
    // Example 3:
    // Input: nums = [1,0,1,2]
    // Output: 3
    // Explanation:
    // Delete nums[0] = 1. The array becomes [0, 1, 2].
    // Now, nums[0] = 0, nums[1] = 1, and nums[2] = 2, so all indices are fixed points.
    // Thus, the answer is 3. 
    fmt.Println(maxFixedPoints([]int{1,0,1,2})) // 3

    fmt.Println(maxFixedPoints([]int{2,0,0})) // 1
    fmt.Println(maxFixedPoints([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maxFixedPoints([]int{9,8,7,6,5,4,3,2,1})) // 4

    fmt.Println(maxFixedPoints1([]int{0,2,1})) // 2
    fmt.Println(maxFixedPoints1([]int{3,1,2})) // 2
    fmt.Println(maxFixedPoints1([]int{1,0,1,2})) // 3
    fmt.Println(maxFixedPoints1([]int{2,0,0})) // 1
    fmt.Println(maxFixedPoints1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(maxFixedPoints1([]int{9,8,7,6,5,4,3,2,1})) // 4
}