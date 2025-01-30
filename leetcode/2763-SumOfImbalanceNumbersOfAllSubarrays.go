package main

// 2763. Sum of Imbalance Numbers of All Subarrays
// The imbalance number of a 0-indexed integer array arr of length n is defined as the number of indices in sarr = sorted(arr) such that:
//     1. 0 <= i < n - 1, and
//     2. sarr[i+1] - sarr[i] > 1

// Here, sorted(arr) is the function that returns the sorted version of arr.

// Given a 0-indexed integer array nums, return the sum of imbalance numbers of all its subarrays.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [2,3,1,4]
// Output: 3
// Explanation: There are 3 subarrays with non-zero imbalance numbers:
// - Subarray [3, 1] with an imbalance number of 1.
// - Subarray [3, 1, 4] with an imbalance number of 1.
// - Subarray [1, 4] with an imbalance number of 1.
// The imbalance number of all other subarrays is 0. Hence, the sum of imbalance numbers of all the subarrays of nums is 3. 

// Example 2:
// Input: nums = [1,3,3,3,5]
// Output: 8
// Explanation: There are 7 subarrays with non-zero imbalance numbers:
// - Subarray [1, 3] with an imbalance number of 1.
// - Subarray [1, 3, 3] with an imbalance number of 1.
// - Subarray [1, 3, 3, 3] with an imbalance number of 1.
// - Subarray [1, 3, 3, 3, 5] with an imbalance number of 2. 
// - Subarray [3, 3, 3, 5] with an imbalance number of 1. 
// - Subarray [3, 3, 5] with an imbalance number of 1.
// - Subarray [3, 5] with an imbalance number of 1.
// The imbalance number of all other subarrays is 0. Hence, the sum of imbalance numbers of all the subarrays of nums is 8. 

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= nums.length

import "fmt"

func sumImbalanceNumbers(nums []int) int {
    res, n := 0, len(nums)
    searchInsert := func(nums []int, target int) int {
        l, r := 0, len(nums) - 1
        for l <= r {
            mid := l + (r - l) / 2
            if nums[mid] < target {
                l = mid + 1
            } else if nums[mid] > target {
                r = mid - 1
            } else {
                return mid
            }
        }
        return l
    }
    for i := 0; i < n; i++ {
        subarr, imbalance := []int{}, 0
        for j := i; j < n; j++ {
            pos := searchInsert(subarr, nums[j])
            // We have two options
            if pos == len(subarr) {
                subarr = append(subarr, nums[j])
                if len(subarr) > 1 && subarr[len(subarr) - 1] - subarr[len(subarr) - 2] > 1 {
                    imbalance++
                }
            } else {
                // In the middle; affects the imbalance
                if pos == 0 {
                    if len(subarr) > 0 && subarr[0] - nums[j] > 1 { imbalance++ }
                } else {
                    if subarr[pos] - subarr[pos-1] > 1 { imbalance-- }
                    if nums[j] - subarr[pos-1] > 1     { imbalance++ }
                    if subarr[pos] - nums[j] > 1       { imbalance++ }
                }
                subarr = append(subarr, 0)
                copy(subarr[pos+1:], subarr[pos:])
                subarr[pos] = nums[j]
            }
            res += imbalance
        }
    }
    return res
}

func sumImbalanceNumbers1(nums []int) int {
    res, n := 0, len(nums)
    right, index := make([]int, n), make([]int, n + 1)
    for i := range index {
        index[i] = n
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        v := nums[i]
        right[i] = min(index[v], index[v - 1]) // right[i] 表示 nums[i] 右侧的 v 和 v - 1 的最近下标（不存在时为 n）
        index[v] = i
    }
    for i := range index {
        index[i] = -1
    }
    for i, v := range nums {
        // 统计 v 能产生多少贡献
        res += (i - index[v-1]) * (right[i] - i) // 子数组左端点个数 * 子数组右端点个数
        index[v] = i
    }
    // 上面计算的时候，每个子数组的最小值必然可以作为贡献，而这是不合法的
    // 所以每个子数组都多算了 1 个不合法的贡献
    return res - n * (n + 1) / 2
}

func main() {
    // Example 1:
    // Input: nums = [2,3,1,4]
    // Output: 3
    // Explanation: There are 3 subarrays with non-zero imbalance numbers:
    // - Subarray [3, 1] with an imbalance number of 1.
    // - Subarray [3, 1, 4] with an imbalance number of 1.
    // - Subarray [1, 4] with an imbalance number of 1.
    // The imbalance number of all other subarrays is 0. Hence, the sum of imbalance numbers of all the subarrays of nums is 3. 
    fmt.Println(sumImbalanceNumbers([]int{2,3,1,4})) // 3
    // Example 2:
    // Input: nums = [1,3,3,3,5]
    // Output: 8
    // Explanation: There are 7 subarrays with non-zero imbalance numbers:
    // - Subarray [1, 3] with an imbalance number of 1.
    // - Subarray [1, 3, 3] with an imbalance number of 1.
    // - Subarray [1, 3, 3, 3] with an imbalance number of 1.
    // - Subarray [1, 3, 3, 3, 5] with an imbalance number of 2. 
    // - Subarray [3, 3, 3, 5] with an imbalance number of 1. 
    // - Subarray [3, 3, 5] with an imbalance number of 1.
    // - Subarray [3, 5] with an imbalance number of 1.
    // The imbalance number of all other subarrays is 0. Hence, the sum of imbalance numbers of all the subarrays of nums is 8. 
    fmt.Println(sumImbalanceNumbers([]int{1,3,3,3,5})) // 8

    fmt.Println(sumImbalanceNumbers([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(sumImbalanceNumbers([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(sumImbalanceNumbers1([]int{2,3,1,4})) // 3
    fmt.Println(sumImbalanceNumbers1([]int{1,3,3,3,5})) // 8
    fmt.Println(sumImbalanceNumbers1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(sumImbalanceNumbers1([]int{9,8,7,6,5,4,3,2,1})) // 0
}