package main 

// 3948. Lexicographically Maximum MEX Array
// You are given an integer array nums.

// You want to construct an array result by repeatedly performing the following operation until nums becomes empty:
//     1. Choose an integer k such that 1 <= k <= len(nums).
//     2. Compute the MEX of the first k elements of nums.
//     3. Append this MEX to result.
//     4. Remove the first k elements from nums.

// Return the lexicographically maximum array result that can be obtained after performing the operations.

// The MEX of an array is the smallest non-negative integer not present in the array.

// An array a is lexicographically greater than an array b if in the first position where a and b differ, 
// array a has an element that is greater than the corresponding element in b. 
// If the first min(a.length, b.length) elements do not differ, then the longer array is the lexicographically greater one.

// Example 1:
// Input: nums = [0,1,0]
// Output: [2,1]
// Explanation:
// Take the first k = 2 elements [0, 1] which has MEX = 2. Current result = [2].
// Remaining array [0] has MEX = 1. Thus, the final result = [2, 1].

// Example 2:
// Input: nums = [1,0,2]
// Output: [3]
// Explanation:
// Take the first k = 3 elements [1, 0, 2] which has MEX = 3.
// nums is now empty. Thus, the final result = [3].

// Example 3:
// Input: nums = [3,1]
// Output: [0,0]
// Explanation:​​​​​​​
// Take k = 1, first element [3] has MEX = 0. Current result = [0].
// Remaining array [1] has MEX = 0. Thus, the final result = [0, 0].

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^5

import "fmt"

func maximumMEX(nums []int) []int {
    res, n := []int{}, len(nums)
    // mex 最大是 n，>= n 的数无需考虑
    pos := make([][]int, n+1) // n 作为哨兵
    for i, x := range nums {
        if x < n {
            pos[x] = append(pos[x], i)
        }
    }
    for i := 0; i < n; i++ {
        start := i // 这一段的起点
        // 枚举这一段的 mex，越大越好（字典序越大）
        mex := 0
        for ; ; mex++ {
            // 清理在 start 之前的下标
            for len(pos[mex]) > 0 && pos[mex][0] < start {
                pos[mex] = pos[mex][1:]
            }
            if len(pos[mex]) == 0 {
                break
            }
            // 这一段包含剩余元素中的最左边的 mex
            i = max(i, pos[mex][0])
        }
        res = append(res, mex)
    }
    return res
}

func maximumMEX1(nums []int) []int {
    n, m := len(nums), 0
    res, arr := []int{}, make([]int, n + 2)
    visited := make([]bool, n + 2)
    for _, v := range nums {
        if v <= n{
            arr[v]++
        }
    }
    for arr[m] > 0 {
        m++
    }
    for i := 0; i < n; {
        if m == 0{
            res = append(res, 0)
            i++
            continue
        }
        t, c := m, 0
        for ; i < n; i++ {
            v := nums[i]
            if v < t && !visited[v] {
                visited[v] = true
                c++
            }
            if v <= n {
                arr[v]--
                if arr[v] == 0 && v < m {
                    m = v
                }
            }
            if c == t {
                i++
                break
            }
        }
        for j := range t {
            visited[j] = false
        }
        res = append(res, t)
    }
    return res
}

func maximumMEX2(nums []int) []int {
    mex, n := 0, len(nums)
    count := make([]int, n + 1)
    for _, v := range nums {
        if v <= n {
            count[v]++
        }
    }
    for mex <= n && count[mex] > 0 {
        mex++
    }
    res, seen := make([]int, 0, n), make([]int, n + 1)
    i, segID := 0, 1
    for i < n {
        m := mex
        if m == 0 {
            res = append(res, 0)
            v := nums[i]
            if v <= n {
                count[v]--
                if count[v] == 0 && v < mex {
                    mex = v
                }
            }
            i++
            continue
        }
        need := m
        segID++
        for need > 0 {
            v := nums[i]
            if v <= n {
                count[v]--
                if count[v] == 0 && v < mex {
                    mex = v
                }
            }
            if v < m && seen[v] != segID {
                seen[v] = segID
                need--
            }
            i++
        }
        res = append(res, m)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,0]
    // Output: [2,1]
    // Explanation:
    // Take the first k = 2 elements [0, 1] which has MEX = 2. Current result = [2].
    // Remaining array [0] has MEX = 1. Thus, the final result = [2, 1].
    fmt.Println(maximumMEX([]int{0,1,0})) // [2,1]
    // Example 2:
    // Input: nums = [1,0,2]
    // Output: [3]
    // Explanation:
    // Take the first k = 3 elements [1, 0, 2] which has MEX = 3.
    fmt.Println(maximumMEX([]int{1,0,2})) // [3]
    // Example 3:
    // Input: nums = [3,1]
    // Output: [0,0]
    // Explanation:​​​​​​​
    // Take k = 1, first element [3] has MEX = 0. Current result = [0].
    // Remaining array [1] has MEX = 0. Thus, the final result = [0, 0].
    fmt.Println(maximumMEX([]int{3,1})) // [0,0]

    fmt.Println(maximumMEX([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0]
    fmt.Println(maximumMEX([]int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 0 0 0 0 0]

    fmt.Println(maximumMEX1([]int{0,1,0})) // [2,1]
    fmt.Println(maximumMEX1([]int{1,0,2})) // [3]
    fmt.Println(maximumMEX1([]int{3,1})) // [0,0]
    fmt.Println(maximumMEX1([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0]
    fmt.Println(maximumMEX1([]int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 0 0 0 0 0]

    fmt.Println(maximumMEX2([]int{0,1,0})) // [2,1]
    fmt.Println(maximumMEX2([]int{1,0,2})) // [3]
    fmt.Println(maximumMEX2([]int{3,1})) // [0,0]
    fmt.Println(maximumMEX2([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0]
    fmt.Println(maximumMEX2([]int{9,8,7,6,5,4,3,2,1})) // [0 0 0 0 0 0 0 0 0]
}