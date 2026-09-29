package main

// 4067. Longest Subarray With Restricted Pair Sums
// You are given an integer array nums.

// A subarray nums[l..r] is valid if there are no three distinct indices i, j, and k such that l <= i, j, k <= r and:
//     1. nums[i] + nums[j] == nums[k]
//     2. Return the maximum length of a valid subarray of nums.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [2,3,5,3,2,1]
// Output: 3
// Explanation:
// Consider the subarray [3, 5, 3]. The pairs of elements at distinct indices have the following sums:
// 3 + 5 = 8
// 3 + 3 = 6, using the two different occurrences of 3
// 5 + 3 = 8
// None of these sums is an element at the remaining index, so the subarray is valid.
// Every subarray of length 4 contains 2, 3, and 5 at distinct indices, where 2 + 3 = 5. 
// Therefore, no longer valid subarray exists, and the answer is 3.

// Example 2:
// Input: nums = [3,4,5,6]
// Output: 4
// Explanation:
// The sums obtained from every pair of elements at distinct indices are 7, 8, 9, 9, 10, and 11. 
// None of these values appears at the remaining index, so the entire array is valid.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 500

import "fmt"
import "slices"

func maxSubarray(nums []int) int{
    res, left, mx := 0, 0, slices.Max(nums)
    countS, countD := make([]int, mx * 2 + 1), make([]int, mx + 1)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    // 枚举有效子数组的右端点为 i，那么左端点 left 最小是多少？
    for i, x := range nums {
        // x 进入窗口前，先判断：
        // 如果窗口中有两数之和等于 x，或者两数之差等于 x，那么必须缩小窗口
        for countS[x] > 0 || countD[x] > 0 {
            y := nums[left]
            left++
            for _, z := range nums[left:i] {
                countS[y+z]--
                countD[abs(y-z)]--
            }
        }
        // 用子数组 [left, i] 的长度更新答案的最大值
        res = max(res, i-left+1)
        // 元素 x 进入窗口
        for _, y := range nums[left:i] {
            countS[x+y]++
            countD[abs(x-y)]++
        }
    }
    return res
}

func maxSubarray1(nums []int) int {
    const MX = 500
    count := make([]int, MX + 1)
    bad := func(v int) bool {
        for a:=1;a<=v/2; a++ {
            b := v-a
            if a == b {
                if count[a]>=2 {
                    return true
                }
            } else if count[a]>0&& count[b]>0 {
                return true
            }
        }
        for b :=1; v + b <= MX; b++ {
            if count[v+b]==0 {
                continue
            }
            if b == v {
                if count[v] >= 2 {
                    return true
                }
            } else if count[b] > 0 {
                return true
            }
        }
        return false
    }
    res, l := 0,0
    for r,v := range nums {
        count[v]++
        for bad(v){
            count[nums[l]]--
            l++
        }
        if r - l + 1 > res {
            res = r - l + 1
        }
    }
    return res  
}

func maxSubarray2(nums []int) int {
    res, count, l, n := 0, 0, 0, len(nums)
    if n == 0 {
        return 0
    }
    freq, pairSum := make([]int, 501), make([]int, 501)
    for r := 0; r < n; r++ {
        v := nums[r]
        // Add pairs of nums[r] with elements in [l, r-1]
        for j := l; j < r; j++ {
            s := v + nums[j]
            if s <= 500 {
                if freq[s] > 0 && pairSum[s] == 0 {
                    count++
                }
                pairSum[s]++
            }
        }
        // Add nums[r] to freq
        if pairSum[v] > 0 && freq[v] == 0 {
            count++
        }
        freq[v]++
        // Shrink window from left while there are violations
        for count > 0 {
            u := nums[l]
            // Remove nums[l] from freq
            if pairSum[u] > 0 && freq[u] == 1 {
                count--
            }
            freq[u]--
            // Remove pairs involving nums[l] and elements in [l+1, r]
            for j := l + 1; j <= r; j++ {
                s := u + nums[j]
                if s <= 500 {
                    pairSum[s]--
                    if freq[s] > 0 && pairSum[s] == 0 {
                        count--
                    }
                }
            }
            l++
        }
        if r-l+1 > res {
            res = r - l + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,3,5,3,2,1]
    // Output: 3
    // Explanation:
    // Consider the subarray [3, 5, 3]. The pairs of elements at distinct indices have the following sums:
    // 3 + 5 = 8
    // 3 + 3 = 6, using the two different occurrences of 3
    // 5 + 3 = 8
    // None of these sums is an element at the remaining index, so the subarray is valid.
    // Every subarray of length 4 contains 2, 3, and 5 at distinct indices, where 2 + 3 = 5. 
    // Therefore, no longer valid subarray exists, and the answer is 3.
    fmt.Println(maxSubarray([]int{2,3,5,3,2,1})) // 3
    // Example 2:
    // Input: nums = [3,4,5,6]
    // Output: 4
    // Explanation:
    // The sums obtained from every pair of elements at distinct indices are 7, 8, 9, 9, 10, and 11. 
    // None of these values appears at the remaining index, so the entire array is valid.
    fmt.Println(maxSubarray([]int{3,4,5,6})) // 4

    fmt.Println(maxSubarray([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(maxSubarray([]int{9,8,7,6,5,4,3,2,1})) // 5

    fmt.Println(maxSubarray1([]int{2,3,5,3,2,1})) // 3
    fmt.Println(maxSubarray1([]int{3,4,5,6})) // 4
    fmt.Println(maxSubarray1([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(maxSubarray1([]int{9,8,7,6,5,4,3,2,1})) // 5

    fmt.Println(maxSubarray2([]int{2,3,5,3,2,1})) // 3
    fmt.Println(maxSubarray2([]int{3,4,5,6})) // 4
    fmt.Println(maxSubarray2([]int{1,2,3,4,5,6,7,8,9})) // 5
    fmt.Println(maxSubarray2([]int{9,8,7,6,5,4,3,2,1})) // 5
}