package main

// 1477. Find Two Non-overlapping Sub-arrays Each With Target Sum
// You are given an array of integers arr and an integer target.

// You have to find two non-overlapping sub-arrays of arr each with a sum equal target. 
// There can be multiple answers so you have to find an answer where the sum of the lengths of the two sub-arrays is minimum.

// Return the minimum sum of the lengths of the two required sub-arrays, 
// or return -1 if you cannot find such two sub-arrays.

// Example 1:
// Input: arr = [3,2,2,4,3], target = 3
// Output: 2
// Explanation: Only two sub-arrays have sum = 3 ([3] and [3]). The sum of their lengths is 2.

// Example 2:
// Input: arr = [7,3,4,7], target = 7
// Output: 2
// Explanation: Although we have three non-overlapping sub-arrays of sum = 7 ([7], [3,4] and [7]), but we will choose the first and third sub-arrays as the sum of their lengths is 2.

// Example 3:
// Input: arr = [4,3,2,6,2,3,4], target = 6
// Output: -1
// Explanation: We have only one sub-array of sum = 6.

// Constraints:
//     1 <= arr.length <= 10^5
//     1 <= arr[i] <= 1000
//     1 <= target <= 10^8

import "fmt"

// 滑动窗口
func minSumOfLengths(arr []int, target int) int {
    dp :=[100001]int{}
    i, j, sum, mn, res  := 0, 0, 0, 1 << 31, 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for j < len(arr) {
        sum += arr[j]
        for sum > target {
            sum -= arr[i]
            i++
        }
        if sum == target {
            l := j - i + 1
            mn = min(mn, l)
            if i > 0 {
                res = min(res, dp[i-1] + l)
            }
        }
        dp[j] = mn
        j++
    }
    if res == 1 << 31 { return -1 }
    return res
}

func minSumOfLengths1(arr []int, target int) int {
    n, sum := len(arr), 0
    suf := make([]int, n + 1)
    suf[n] = n
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, j := n - 1, n - 1; i >= 0; i -- {
        sum += arr[i]
        for sum > target {
            sum -= arr[j]
            j -- 
        }
        suf[i] = suf[i + 1]
        if sum == target {
            suf[i] = min(suf[i], j - i + 1)
        }
    }
    sum = 0
    res, pre := -1, n
    for i, j := 0, 0; j < n; j ++ {
        sum += arr[j]
        for sum > target {
            sum -= arr[i]
            i ++ 
        }
        if sum == target {
            pre = min(pre, j - i + 1)
        }
        if pre > 0 && pre < n && suf[j + 1] > 0 && suf[j + 1] < n {
            if res == -1 {
                res = pre + suf[j + 1]
            } else {
                res = min(res, pre + suf[j + 1])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: arr = [3,2,2,4,3], target = 3
    // Output: 2
    // Explanation: Only two sub-arrays have sum = 3 ([3] and [3]). The sum of their lengths is 2.
    fmt.Println(minSumOfLengths([]int{3,2,2,4,3}, 3)) // 2
    // Example 2:
    // Input: arr = [7,3,4,7], target = 7
    // Output: 2
    // Explanation: Although we have three non-overlapping sub-arrays of sum = 7 ([7], [3,4] and [7]), but we will choose the first and third sub-arrays as the sum of their lengths is 2.
    fmt.Println(minSumOfLengths([]int{7,3,4,7}, 7)) // 2
    // Example 3:
    // Input: arr = [4,3,2,6,2,3,4], target = 6
    // Output: -1
    // Explanation: We have only one sub-array of sum = 6.
    fmt.Println(minSumOfLengths([]int{4,3,2,6,2,3,4}, 6)) // -1

    fmt.Println(minSumOfLengths1([]int{3,2,2,4,3}, 3)) // 2
    fmt.Println(minSumOfLengths1([]int{7,3,4,7}, 7)) // 2
    fmt.Println(minSumOfLengths1([]int{4,3,2,6,2,3,4}, 6)) // -1
}