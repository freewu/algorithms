package main

// 2563. Count the Number of Fair Pairs
// Given a 0-indexed integer array nums of size n and two integers lower and upper, return the number of fair pairs.

// A pair (i, j) is fair if:
//     0 <= i < j < n, and
//     lower <= nums[i] + nums[j] <= upper

// Example 1:
// Input: nums = [0,1,7,4,4,5], lower = 3, upper = 6
// Output: 6
// Explanation: There are 6 fair pairs: (0,3), (0,4), (0,5), (1,3), (1,4), and (1,5).

// Example 2:
// Input: nums = [1,7,9,2,5], lower = 11, upper = 11
// Output: 1
// Explanation: There is a single fair pair: (2,3).

// Constraints:
//     1 <= nums.length <= 10^5
//     nums.length == n
//     -10^9 <= nums[i] <= 10^9
//     -10^9 <= lower <= upper <= 10^9

import "fmt"
import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    res := 0
    for i, v := range nums { // 锁定第一个点
        // 第一次 二分
        l, r := 0, i // 判断从0 到 i 之间满足的值
        for l < r {  // 二分查找大于lower的值
            mid := l + (r - l) / 2 // 防止溢出
            if nums[mid] + v >= lower {
                r = mid // 左闭右开区间
            } else {
                l = mid + 1
            }
        }
        i1 := l // 存储起始位置
        // 第二次 二分
        l, r = i1, i // 判断从0 到 i 之间满足的值
        for l < r {
            mid := l + (r-l)/2
            if nums[mid] + v > upper { // l 的值 截至在等于时候 或者可以写 nums[mid]+v > upper 但下面要i2 = r
                r = mid
            } else {
                l = mid + 1
            }
        }
        i2 := r // 存储终止位置
        res += (i2 - i1) // 结果累加
    }
    return int64(res)
}

func countFairPairs1(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    res, n := 0, len(nums)
    r1, r2 := n - 1, n - 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for r1 > i && nums[r1]+nums[i] >= lower { r1-- }
        for r2 > i && nums[r2]+nums[i] > upper  { r2-- }
        if r2 <= i { break }
        res += (r2 - max(i, r1))
    }
    return int64(res)
}

func countFairPairs2(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    lower_bound := func(value int) int64 { // Calculate the number of pairs with sum less than `value`.
        res, left, right := int64(0), 0, len(nums) - 1
        for left < right {
            sum := nums[left] + nums[right]
            if sum < value { // If sum is less than value, add the size of window to result and move to the next index
                res += int64(right - left)
                left++
            } else { // Otherwise, shift the right pointer backwards, until we get a valid window.
                right--
            }
        }
        return res
    }
    return lower_bound(upper + 1) - lower_bound(lower)
}

func main() {
    // Example 1:
    // Input: nums = [0,1,7,4,4,5], lower = 3, upper = 6
    // Output: 6
    // Explanation: There are 6 fair pairs: (0,3), (0,4), (0,5), (1,3), (1,4), and (1,5).
    fmt.Println(countFairPairs([]int{0,1,7,4,4,5}, 3, 6)) // 6
    // Example 2:
    // Input: nums = [1,7,9,2,5], lower = 11, upper = 11
    // Output: 1
    // Explanation: There is a single fair pair: (2,3).
    fmt.Println(countFairPairs([]int{1,7,9,2,5}, 11, 11)) // 1

    fmt.Println(countFairPairs1([]int{0,1,7,4,4,5}, 3, 6)) // 6
    fmt.Println(countFairPairs1([]int{1,7,9,2,5}, 11, 11)) // 1

    fmt.Println(countFairPairs2([]int{0,1,7,4,4,5}, 3, 6)) // 6
    fmt.Println(countFairPairs2([]int{1,7,9,2,5}, 11, 11)) // 1
}