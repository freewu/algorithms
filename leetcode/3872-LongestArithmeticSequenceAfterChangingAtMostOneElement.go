package main

// 3872. Longest Arithmetic Sequence After Changing At Most One Element
// You are given an integer array nums.

// A subarray is arithmetic if the difference between consecutive elements in the subarray is constant.

// You can replace at most one element in nums with any integer. Then, you select an arithmetic subarray from nums.

// Return an integer denoting the maximum length of the arithmetic subarray you can select.

// Example 1:
// Input: nums = [9,7,5,10,1]
// Output: 5
// Explanation:
// Replace nums[3] = 10 with 3. The array becomes [9, 7, 5, 3, 1].
// Select the subarray [9, 7, 5, 3, 1], which is arithmetic because consecutive elements have a common difference of -2.

// Example 2:
// Input: nums = [1,2,6,7]
// Output: 3
// Explanation:
// Replace nums[0] = 1 with -2. The array becomes [-2, 2, 6, 7].
// Select the subarray [-2, 2, 6, 7], which is arithmetic because consecutive elements have a common difference of 4.

// Constraints:
//     4 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"
import "slices"

func longestArithmetic(nums []int) int {
    calc := func(nums []int) []int {
        n := len(nums)
        prefix := make([]int, n)
        prefix[0], prefix[1] = 1, 2
        for i := 2; i < n; i++ {
            if nums[i-2]+nums[i] == nums[i-1] * 2 { // 三个数等差
                prefix[i] = prefix[i-1] + 1
            } else {
                prefix[i] = 2
            }
        }
        return prefix
    }
    res, n := 0, len(nums)
    prefix := calc(nums)
    res = slices.Max(prefix) + 1
    if res >= n { // 整个数组是等差的，或者修改端点元素后是等差的
        return n
    }
    slices.Reverse(nums)
    suf := calc(nums)
    slices.Reverse(suf)
    slices.Reverse(nums)
    // 注意 max(pre) == max(suf)，无需重复计算
    for i := 1; i < n-1; i++ {
        // 把 nums[i] 改成 d2/2
        d2 := nums[i+1] - nums[i-1]
        if d2 % 2 != 0 { continue } // d2/2 必须是整数
        okLeft := i > 1 && nums[i-1]-nums[i-2] == d2/2
        okRight := i+2 < n && nums[i+2]-nums[i+1] == d2/2
        if okLeft && okRight {
            res = max(res, prefix[i-1]+1+suf[i+1])
        } else if okLeft {
            res = max(res, prefix[i-1]+2)
        } else if okRight {
            res = max(res, suf[i+1]+2)
        }
    }
    return res
}

func longestArithmetic1(nums []int) int {
    calc := func(nums []int, change bool) int {
        i, j, res := 0, 0, 0
        for i+1 < len(nums) {
            for j+1 < len(nums) && nums[i]-nums[i+1] == nums[j]-nums[j+1] {
                j++
            }
            if j+1 == len(nums) || change {
                res = max(res, j-i+1)
                break
            }
            changeIndex := j + 1
            changeNums := nums[j+1]
            nums[changeIndex] = nums[j] - (nums[i] - nums[i+1])
            k := j
            for k+1 < len(nums) && nums[i]-nums[i+1] == nums[k]-nums[k+1] {
                k++
            }
            res = max(res, k-i+1)
            nums[changeIndex] = changeNums
            i = j
        }
        return res
    }
    revCalc := func(nums []int, change bool) int {
        i, j, res := len(nums) - 1, len(nums) - 1, 0
        for i-1 >= 0 {
            for j-1 >= 0 && nums[i]-nums[i-1] == nums[j]-nums[j-1] {
                j--
            }
            if j-1 < 0 || change {
                res = max(res, i-j+1)
                break
            }
            changeIndex := j - 1
            changeNums := nums[j-1]
            nums[changeIndex] = nums[j] - (nums[i] - nums[i-1])
            k := j
            for k-1 >= 0 && nums[i]-nums[i-1] == nums[k]-nums[k-1] {
                k--
            }
            res = max(res, i-k+1)
            nums[changeIndex] = changeNums
            i = j
        }
        return res
    }
    res := calc(nums, false)
    res = max(res, revCalc(nums, false))
    return res
}

func main() {
    // Example 1:
    // Input: nums = [9,7,5,10,1]
    // Output: 5
    // Explanation:
    // Replace nums[3] = 10 with 3. The array becomes [9, 7, 5, 3, 1].
    // Select the subarray [9, 7, 5, 3, 1], which is arithmetic because consecutive elements have a common difference of -2.
    fmt.Println(longestArithmetic([]int{9,7,5,10,1})) // 5
    // Example 2:
    // Input: nums = [1,2,6,7]
    // Output: 3
    // Explanation:
    // Replace nums[0] = 1 with -2. The array becomes [-2, 2, 6, 7].
    // Select the subarray [-2, 2, 6, 7], which is arithmetic because consecutive elements have a common difference of 4.
    fmt.Println(longestArithmetic([]int{1,2,6,7})) // 3

    fmt.Println(longestArithmetic([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(longestArithmetic([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(longestArithmetic1([]int{9,7,5,10,1})) // 5
    fmt.Println(longestArithmetic1([]int{1,2,6,7})) // 3
    fmt.Println(longestArithmetic1([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(longestArithmetic1([]int{9,8,7,6,5,4,3,2,1})) // 9
}