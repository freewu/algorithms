package main

// 1695. Maximum Erasure Value
// You are given an array of positive integers nums and want to erase a subarray containing unique elements. 
// The score you get by erasing the subarray is equal to the sum of its elements.

// Return the maximum score you can get by erasing exactly one subarray.

// An array b is called to be a subarray of a if it forms a contiguous subsequence of a, 
// that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).

// Example 1:
// Input: nums = [4,2,4,5,6]
// Output: 17
// Explanation: The optimal subarray here is [2,4,5,6].

// Example 2:
// Input: nums = [5,2,1,2,5,2,1,2,5]
// Output: 8
// Explanation: The optimal subarray here is [5,2,1] or [1,2,5].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^4

import "fmt"
import "slices"

func maximumUniqueSubarray(nums []int) int {
    res, sum, left := 0, 0, 0
    lastPos := make(map[int]int) // record each num's last position
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums { // iterate over the array
        if pos, ok := lastPos[v]; ok { // if the number appears before
            for j := left; j <= pos; j++ { // find its last position, and remove all the numbers before that position
                sum -= nums[j]           // derease from current sliding window
                delete(lastPos, nums[j]) // remove from the map
            }
            left = pos + 1 // update the start position of the sliding window
        }
        lastPos[v] = i       // record the num's position
        sum += v             // increase the current sliding window
        res = max(res, sum) // add to the map
    }
    return res
}

func maximumUniqueSubarray1(nums []int) int {
    res, sum, left, mx := 0, 0, 0, slices.Max(nums) // res, 当前窗口和, 窗口左边界, 数组中的最大值
    seen := make([]bool, mx + 1) // 创建布尔数组记录元素是否出现过
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        for seen[v] { // 移除重复元素直到当前元素唯一
            seen[nums[left]] = false
            sum -= nums[left]
            left++
        }
        // 添加当前元素
        seen[v] = true
        sum += v
        res = max(res, sum) // 更新最大和
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,2,4,5,6]
    // Output: 17
    // Explanation: The optimal subarray here is [2,4,5,6].
    fmt.Println(maximumUniqueSubarray([]int{4,2,4,5,6})) // 17
    // Example 2:
    // Input: nums = [5,2,1,2,5,2,1,2,5]
    // Output: 8
    // Explanation: The optimal subarray here is [5,2,1] or [1,2,5].
    fmt.Println(maximumUniqueSubarray([]int{5,2,1,2,5,2,1,2,5})) // 8

    fmt.Println(maximumUniqueSubarray([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maximumUniqueSubarray([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maximumUniqueSubarray1([]int{4,2,4,5,6})) // 17
    fmt.Println(maximumUniqueSubarray1([]int{5,2,1,2,5,2,1,2,5})) // 8
    fmt.Println(maximumUniqueSubarray1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maximumUniqueSubarray1([]int{9,8,7,6,5,4,3,2,1})) // 45
}