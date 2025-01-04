package main

// 2499. Minimum Total Cost to Make Arrays Unequal
// You are given two 0-indexed integer arrays nums1 and nums2, of equal length n.

// In one operation, you can swap the values of any two indices of nums1. 
// The cost of this operation is the sum of the indices.

// Find the minimum total cost of performing the given operation any number of times 
// such that nums1[i] != nums2[i] for all 0 <= i <= n - 1 after performing all the operations.

// Return the minimum total cost such that nums1 and nums2 satisfy the above condition. 
// In case it is not possible, return -1.

// Example 1:
// Input: nums1 = [1,2,3,4,5], nums2 = [1,2,3,4,5]
// Output: 10
// Explanation: 
// One of the ways we can perform the operations is:
// - Swap values at indices 0 and 3, incurring cost = 0 + 3 = 3. Now, nums1 = [4,2,3,1,5]
// - Swap values at indices 1 and 2, incurring cost = 1 + 2 = 3. Now, nums1 = [4,3,2,1,5].
// - Swap values at indices 0 and 4, incurring cost = 0 + 4 = 4. Now, nums1 =[5,3,2,1,4].
// We can see that for each index i, nums1[i] != nums2[i]. The cost required here is 10.
// Note that there are other ways to swap values, but it can be proven that it is not possible to obtain a cost less than 10.

// Example 2:
// Input: nums1 = [2,2,2,1,3], nums2 = [1,2,2,3,3]
// Output: 10
// Explanation: 
// One of the ways we can perform the operations is:
// - Swap values at indices 2 and 3, incurring cost = 2 + 3 = 5. Now, nums1 = [2,2,1,2,3].
// - Swap values at indices 1 and 4, incurring cost = 1 + 4 = 5. Now, nums1 = [2,3,1,2,2].
// The total cost needed here is 10, which is the minimum possible.

// Example 3:
// Input: nums1 = [1,2,2], nums2 = [1,2,2]
// Output: -1
// Explanation: 
// It can be shown that it is not possible to satisfy the given conditions irrespective of the number of operations we perform.
// Hence, we return -1.

// Constraints:
//     n == nums1.length == nums2.length
//     1 <= n <= 10^5
//     1 <= nums1[i], nums2[i] <= n

import "fmt"

func minimumTotalCost(nums1 []int, nums2 []int) int64 {
    res, n, mx, mxc, total := 0, len(nums1), 0, 0, 0
    count := make([]int, n + 1)
    for i := 0; i < n; i++ {
        if nums1[i] == nums2[i] {
            res += i
            total++
            count[nums1[i]]++
            if count[nums1[i]] > mxc {
                mxc, mx = count[nums1[i]],nums1[i]
            }
        }
    }
    swap := (mxc << 1) - total
    for i := 0; i < n && swap > 0; i++ {
        if nums1[i] != nums2[i] && nums1[i] != mx && nums2[i] != mx {
            swap--
            res += i
        }
    }
    if swap > 0 { return -1 }
    return int64(res)
}

func minimumTotalCost1(nums1 []int, nums2 []int) int64 {
    res, same, n := 0, 0, len(nums1)
    count := make([]int, n + 1)
    for i, v := range nums1 {
        if v == nums2[i] {
            same++
            res += i
            count[v]++
        }
    }
    mode, lead := 0, 0
    for i, v := range count {
        if t := v * 2 - same; t > 0 {
            mode, lead = t, i
            break
        }
    }
    for i, v := range nums1 {
        if mode > 0 && v != nums2[i] && v != lead && nums2[i] != lead {
            res += i
            mode--
        }
    }
    if mode > 0 { return -1 }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums1 = [1,2,3,4,5], nums2 = [1,2,3,4,5]
    // Output: 10
    // Explanation: 
    // One of the ways we can perform the operations is:
    // - Swap values at indices 0 and 3, incurring cost = 0 + 3 = 3. Now, nums1 = [4,2,3,1,5]
    // - Swap values at indices 1 and 2, incurring cost = 1 + 2 = 3. Now, nums1 = [4,3,2,1,5].
    // - Swap values at indices 0 and 4, incurring cost = 0 + 4 = 4. Now, nums1 =[5,3,2,1,4].
    // We can see that for each index i, nums1[i] != nums2[i]. The cost required here is 10.
    // Note that there are other ways to swap values, but it can be proven that it is not possible to obtain a cost less than 10.
    fmt.Println(minimumTotalCost([]int{1,2,3,4,5}, []int{1,2,3,4,5})) // 10
    // Example 2:
    // Input: nums1 = [2,2,2,1,3], nums2 = [1,2,2,3,3]
    // Output: 10
    // Explanation: 
    // One of the ways we can perform the operations is:
    // - Swap values at indices 2 and 3, incurring cost = 2 + 3 = 5. Now, nums1 = [2,2,1,2,3].
    // - Swap values at indices 1 and 4, incurring cost = 1 + 4 = 5. Now, nums1 = [2,3,1,2,2].
    // The total cost needed here is 10, which is the minimum possible.
    fmt.Println(minimumTotalCost([]int{2,2,2,1,3}, []int{1,2,2,3,3})) // 10
    // Example 3:
    // Input: nums1 = [1,2,2], nums2 = [1,2,2]
    // Output: -1
    // Explanation: 
    // It can be shown that it is not possible to satisfy the given conditions irrespective of the number of operations we perform.
    // Hence, we return -1.
    fmt.Println(minimumTotalCost([]int{1,2,2}, []int{1,2,2})) // -1

    fmt.Println(minimumTotalCost1([]int{1,2,3,4,5}, []int{1,2,3,4,5})) // 10
    fmt.Println(minimumTotalCost1([]int{2,2,2,1,3}, []int{1,2,2,3,3})) // 10
    fmt.Println(minimumTotalCost1([]int{1,2,2}, []int{1,2,2})) // -1
}