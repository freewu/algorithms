package main

// 4109. Count Valid Suffix Sorts II
// You are given two integer arrays nums and target of lengths n and m, respectively, where n <= m.

// For each index i, consider sorting the suffix nums[i..n - 1] of the original array in non-increasing order.

// Return an integer array ans of length m - n + 1, where ans[k] is the number of indices i for which the resulting array equals target[k..k + n - 1].

// A subarray is a contiguous non-empty sequence of elements within an array.

// A suffix of an array is a subarray that begins at any index and extends to its end.

// Example 1:
// Input: nums = [2,2,3], target = [2,2,3,2,1,3,2,2]
// Output: [1,1,0,0,0,1]
// Explanation:
// Sorting the suffix starting at indices 0, 1, and 2 produces [3, 2, 2], [2, 3, 2], and [2, 2, 3], respectively.
//     k = 0: [2, 2, 3] is produced by i = 2.
//     k = 1: [2, 3, 2] is produced by i = 1.
//     k = 2, 3, 4: None of the corresponding subarrays can be produced.
//     k = 5: [3, 2, 2] is produced by i = 0.
// Therefore, the answer is [1, 1, 0, 0, 0, 1].

// Example 2:
// Input: nums = [5,6,5], target = [5,6,5,5,6,5]
// Output: [2,1,0,2]
// Explanation:
// Sorting the suffix starting at index 0 produces [6, 5, 5], while starting at either index 1 or 2 produces [5, 6, 5].
//     k = 0: [5, 6, 5] is produced by i = 1 and i = 2.
//     k = 1: [6, 5, 5] is produced by i = 0.
//     k = 2: [5, 5, 6] cannot be produced.
//     k = 3: [5, 6, 5] is produced by i = 1 and i = 2.
// Therefore, the answer is [2, 1, 0, 2].

// Constraints:
//     n == nums.length
//     m == target.length
//     1 <= n <= m <= 10^{5}
//     1 <= nums[i], target[i] <= 10^{5}

import "fmt"
import "slices"

func countSuffix(nums []int, target []int) []int {
    calc := func (s []int) []int { // 计算并返回 z 数组，其中 z[i] = |LCP(s[i:], s)|
        n := len(s)
        res := make([]int, n)
        l, r := 0, 0
        for i := 1; i < n; i++ {
            if i <= r {
                res[i] = min(res[i-l], r-i+1)
            }
            for i+res[i] < n && s[res[i]] == s[i+res[i]] {
                l, r = i, i + res[i]
                res[i]++
            }
        }
        return res
    }
    findAnagrams := func (s, p []int) []bool { // 找到字符串中所有字母异位词
        mx := max(slices.Max(s), slices.Max(p))
        diff := make([]int, mx+1) // slice 比 map 快
        // 统计 p 的每种整数的出现次数
        for _, x := range p {
            diff[x]++
        }
        res := make([]bool, len(s) - len(p) + 1)
        left := 0
        for right, x := range s {
            diff[x]-- // x 进入窗口
            for diff[x] < 0 { // x 太多了
                diff[s[left]]++ // 左端点离开窗口
                left++
            }
            if right-left+1 == len(p) { // 当前窗口与 p 的每种整数的出现次数都相同（证明见 438 题解）
                res[left] = true
            }
        }
        return res
    }
    n, m := len(nums), len(target)
    arr := findAnagrams(target, nums)
    z := calc(append(nums, target...))
    lastInc := make([]int, m)
    for i := 1; i < m; i++ {
        if target[i-1] < target[i] {
            lastInc[i] = i
        } else {
            lastInc[i] = lastInc[i-1]
        }
    }
    res := make([]int, len(arr))
    for k, valid := range arr {
        if valid {
            res[k] = max(min(z[n+k], n-1)-max(lastInc[k+n-1]-k, 0)+1, 0)
        }
    }
    return res
}

const MX = 100_001
var zn, freq [100_001]int32
var res [100_001]int

func countSuffix1(nums, target []int) []int {
    ext := func(nums, text []int, i, lim int, l, r *int) int {
        z := 0
        if i < *r {
            z = min(*r-i, int(zn[i-*l]))
        }
        for z < lim && nums[z] == text[i+z] {
            z++
        }
        if i+z > *r {
            *l, *r = i, i+z
        }
        return z
    }
    n, m := len(nums), len(target)
    zn[0] = int32(n)
    for i, l, r := 1, 0, 0; i < n; i++ {
        zn[i] = int32(ext(nums, nums, i, n-i, &l, &r))
    }
    bad, st := 0, 0
    add := func(v int, d int32) {
        if freq[v] == 0 {
            bad++
        }
        freq[v] += d
        if freq[v] == 0 {
            bad--
        }
    }
    push := func(e int) {
        add(target[e], 1)
        if e > 0 && target[e-1] < target[e] {
            st = e
        }
    }
    for i, v := range nums {
        add(v, -1)
        push(i)
    }
    l, r := 0, 0
    for k := 0; ; k++ {
        val := 0
        if bad == 0 {
            z := ext(nums, target, k, n, &l, &r)
            val = max(0, min(z, n-1)-max(0, st-k)+1)
        }
        res[k] = val
        e := k + n
        if e == m {
            break
        }
        add(target[k], -1)
        push(e)
    }
    for _, v := range nums {
        freq[v] = 0
    }
    for _, v := range target[m-n:] {
        freq[v] = 0
    }
    return res[:m-n+1]
}

func main() {
    // Example 1:
    // Input: nums = [2,2,3], target = [2,2,3,2,1,3,2,2]
    // Output: [1,1,0,0,0,1]
    // Explanation:
    // Sorting the suffix starting at indices 0, 1, and 2 produces [3, 2, 2], [2, 3, 2], and [2, 2, 3], respectively.
    //     k = 0: [2, 2, 3] is produced by i = 2.
    //     k = 1: [2, 3, 2] is produced by i = 1.
    //     k = 2, 3, 4: None of the corresponding subarrays can be produced.
    //     k = 5: [3, 2, 2] is produced by i = 0.
    // Therefore, the answer is [1, 1, 0, 0, 0, 1].
    fmt.Println(countSuffix([]int{2,2,3}, []int{2,2,3,2,1,3,2,2})) // [1,1,0,0,0,1]
    // Example 2:
    // Input: nums = [5,6,5], target = [5,6,5,5,6,5]
    // Output: [2,1,0,2]
    // Explanation:
    // Sorting the suffix starting at index 0 produces [6, 5, 5], while starting at either index 1 or 2 produces [5, 6, 5].
    //     k = 0: [5, 6, 5] is produced by i = 1 and i = 2.
    //     k = 1: [6, 5, 5] is produced by i = 0.
    //     k = 2: [5, 5, 6] cannot be produced.
    //     k = 3: [5, 6, 5] is produced by i = 1 and i = 2.
    // Therefore, the answer is [2, 1, 0, 2].
    fmt.Println(countSuffix([]int{5,6,5}, []int{5,6,5,5,6,5})) // [2,1,0,2]

    fmt.Println(countSuffix([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [1]
    fmt.Println(countSuffix([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [1]
    fmt.Println(countSuffix([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [0]
    fmt.Println(countSuffix([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // [9]

    fmt.Println(countSuffix1([]int{2,2,3}, []int{2,2,3,2,1,3,2,2})) // [1,1,0,0,0,1]
    fmt.Println(countSuffix1([]int{5,6,5}, []int{5,6,5,5,6,5})) // [2,1,0,2]
    fmt.Println(countSuffix1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [1]
    fmt.Println(countSuffix1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [1]
    fmt.Println(countSuffix1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [0]
    fmt.Println(countSuffix1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // [9]
}