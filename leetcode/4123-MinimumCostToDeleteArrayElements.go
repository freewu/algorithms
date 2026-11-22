package main

// 4123. Minimum Cost to Delete Array
// You are given an integer array nums.

// In one operation, choose any subarray of the current array with length at least two and delete it. 
// The cost of this operation is the sum of the first and last elements of the chosen subarray.

// After each operation, the remaining elements are concatenated in their original order.

// Return the minimum total cost required to delete the entire array.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [2,-7,3,-5]
// Output: -7
// Explanation:
//     Delete [-7, 3] with cost -7 + 3 = -4. The array becomes [2, -5].
//     Delete [2, -5] with cost 2 + (-5) = -3.
//     The total cost is -4 + (-3) = -7.

// Example 2:
// Input: nums = [4,0,-5,2]
// Output: 1
// Explanation:
//     Delete [0, -5] with cost 0 + (-5) = -5. The array becomes [4, 2].
//     Delete [4, 2] with cost 4 + 2 = 6.
//     The total cost is -5 + 6 = 1.

// Example 3:
// Input: nums = [7,-8,1]
// Output: 8
// Explanation:
//     To delete the entire array, all three elements must be deleted in one operation. 
//     Deleting only two elements would leave one element, which cannot be deleted.
//     Deleting [7, -8, 1] costs 7 + 1 = 8.

// Constraints:
//     2 <= nums.length <= 10^{5}
//     -10^{9} <= nums[i] <= 10^{9}

import "fmt"
import "sort"

func minCost(nums []int) int64 {
    n := len(nums)
    val := int64(nums[0]) + int64(nums[n-1])
    if n == 2 {
        return val
    }
    arr := make([]int64, n-2)
    for i := 1; i <= n-2; i++ {
        arr[i-1] = int64(nums[i])
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    res := val
    for i := 0; i < n-3; i += 2 {
        val += arr[i] + arr[i + 1]
        if val < res {
            res = val
        }
    }
    return res
}

func minCost1(nums []int) int64 {
    a, b := 0, 1 << 61
    for i := 1; i < len(nums)-1; i++ {
        a, b = min(a, b + nums[i]), min(b, a + nums[i])
    }
    return int64(nums[0] + nums[len(nums)-1] + a)
}

func main() {
    // Example 1:
    // Input: nums = [2,-7,3,-5]
    // Output: -7
    // Explanation:
    //     Delete [-7, 3] with cost -7 + 3 = -4. The array becomes [2, -5].
    //     Delete [2, -5] with cost 2 + (-5) = -3.
    //     The total cost is -4 + (-3) = -7.
    fmt.Println(minCost([]int{2, -7, 3, -5})) // -7
    // Example 2:
    // Input: nums = [4,0,-5,2]
    // Output: 1
    // Explanation:
    //     Delete [0, -5] with cost 0 + (-5) = -5. The array becomes [4, 2].
    //     Delete [4, 2] with cost 4 + 2 = 6.
    //     The total cost is -5 + 6 = 1.
    fmt.Println(minCost([]int{4, 0, -5, 2})) // 1
    // Example 3:
    // Input: nums = [7,-8,1]
    // Output: 8
    // Explanation:
    //     To delete the entire array, all three elements must be deleted in one operation. 
    //     Deleting only two elements would leave one element, which cannot be deleted.
    //     Deleting [7, -8, 1] costs 7 + 1 = 8.
    fmt.Println(minCost([]int{7, -8, 1})) // 8

    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9})) // 10
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1})) // 10

    fmt.Println(minCost1([]int{2, -7, 3, -5})) // -7
    fmt.Println(minCost1([]int{4, 0, -5, 2})) // 1
    fmt.Println(minCost1([]int{7, -8, 1})) // 8
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9})) // 10
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1})) // 10
}