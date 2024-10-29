package main

// 1681. Minimum Incompatibility
// You are given an integer array nums​​​ and an integer k. 
// You are asked to distribute this array into k subsets of equal size such that there are no two equal elements in the same subset.

// A subset's incompatibility is the difference between the maximum and minimum elements in that array.

// Return the minimum possible sum of incompatibilities of the k subsets after distributing the array optimally, 
// or return -1 if it is not possible.

// A subset is a group integers that appear in the array with no particular order.

// Example 1:
// Input: nums = [1,2,1,4], k = 2
// Output: 4
// Explanation: The optimal distribution of subsets is [1,2] and [1,4].
// The incompatibility is (2-1) + (4-1) = 4.
// Note that [1,1] and [2,4] would result in a smaller sum, but the first subset contains 2 equal elements.

// Example 2:
// Input: nums = [6,3,8,1,3,1,2,2], k = 4
// Output: 6
// Explanation: The optimal distribution of subsets is [1,2], [2,3], [6,8], and [1,3].
// The incompatibility is (2-1) + (3-2) + (8-6) + (3-1) = 6.

// Example 3:
// Input: nums = [5,3,3,6,3,3], k = 3
// Output: -1
// Explanation: It is impossible to distribute nums into 3 subsets where no two elements are equal in the same subset.

// Constraints:
//     1 <= k <= nums.length <= 16
//     nums.length is divisible by k
//     1 <= nums[i] <= nums.length

import "fmt"
import "sort"

func minimumIncompatibility(nums []int, k int) int {
    res, n := 1 << 31, len(nums)
    sort.Ints(nums)
    eachSize := n / k
    count, orders := make([]int, n + 1), []int{}
    for _, v := range nums {
        count[v]++
        if count[v] > k {
            return -1
        }
    }
    for i := range count {
        orders = append(orders, i)
    }
    sort.Ints(orders)
    var dfs func(start, sum int, group []int)
    dfs = func(start, sum int, group []int) {
        if len(group) > 0 && len(group) % eachSize == 0 {
            sum += group[len(group)-1] - group[len(group) - eachSize]
            start = 0
        }
        if sum >= res { return }
        if len(group) == len(nums) {
            res = min(res, sum)
            return
        }
        for i := start; i < len(count); i++ {
            if count[orders[i]] == 0  { continue }
            count[orders[i]]--
            group = append(group, orders[i])
            dfs(i + 1, sum, group)
            group = group[:len(group)-1]
            count[orders[i]]++
            if start == 0 { break } // important
        }
    }
    dfs(0, 0, []int{})
    if res == 1 << 31 { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,4], k = 2
    // Output: 4
    // Explanation: The optimal distribution of subsets is [1,2] and [1,4].
    // The incompatibility is (2-1) + (4-1) = 4.
    // Note that [1,1] and [2,4] would result in a smaller sum, but the first subset contains 2 equal elements.
    fmt.Println(minimumIncompatibility([]int{1,2,1,4}, 2)) // 4
    // Example 2:
    // Input: nums = [6,3,8,1,3,1,2,2], k = 4
    // Output: 6
    // Explanation: The optimal distribution of subsets is [1,2], [2,3], [6,8], and [1,3].
    // The incompatibility is (2-1) + (3-2) + (8-6) + (3-1) = 6.
    fmt.Println(minimumIncompatibility([]int{6,3,8,1,3,1,2,2}, 4)) // 6
    // Example 3:
    // Input: nums = [5,3,3,6,3,3], k = 3
    // Output: -1
    // Explanation: It is impossible to distribute nums into 3 subsets where no two elements are equal in the same subset.
    fmt.Println(minimumIncompatibility([]int{5,3,3,6,3,3}, 3)) // -1
}