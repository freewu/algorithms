package main

// 3431. Minimum Unlocked Indices to Sort Nums
// You are given an array nums consisting of integers between 1 and 3, and a binary array locked of the same size.

// We consider nums sortable if it can be sorted using adjacent swaps, 
// where a swap between two indices i and i + 1 is allowed if nums[i] - nums[i + 1] == 1 and locked[i] == 0.

// In one operation, you can unlock any index i by setting locked[i] to 0.

// Return the minimum number of operations needed to make nums sortable. 
// If it is not possible to make nums sortable, return -1.

// Example 1:
// Input: nums = [1,2,1,2,3,2], locked = [1,0,1,1,0,1]
// Output: 0
// Explanation:
// We can sort nums using the following swaps:
//     swap indices 1 with 2
//     swap indices 4 with 5
// So, there is no need to unlock any index.

// Example 2:
// Input: nums = [1,2,1,1,3,2,2], locked = [1,0,1,1,0,1,0]
// Output: 2
// Explanation:
// If we unlock indices 2 and 5, we can sort nums using the following swaps:
//     swap indices 1 with 2
//     swap indices 2 with 3
//     swap indices 4 with 5
//     swap indices 5 with 6

// Example 3:
// Input: nums = [1,2,1,2,3,2,1], locked = [0,0,0,0,0,0,0]
// Output: -1
// Explanation:
// Even if all indices are unlocked, it can be shown that nums is not sortable.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 3
//     locked.length == nums.length
//     0 <= locked[i] <= 1

import "fmt"

// func minUnlockedIndices(nums []int, locked []int) int {
//     res, n, isSorted := 0, len(nums), true // 检查是否已经是升序排列
//     for i := 0; i < n - 1; i++ {
//         if nums[i] > nums[i + 1] {
//             isSorted = false
//             break
//         }
//     }
//     if isSorted { return 0 }
//     for i := 0; i < n -1 ; i++ { // 尝试通过交换排序
//         if nums[i] > nums[i+1] {
//             if nums[i]-nums[i+1] == 1 && locked[i] == 0 { // 检查是否可以交换
//                 nums[i], nums[i + 1] = nums[i + 1], nums[i] // 交换 nums[i] 和 nums[i+1]
//             } else {
//                 if locked[i] == 1 { // 需要解锁 locked[i]
//                     locked[i] = 0
//                     res++
//                 }
//                 if nums[i]-nums[i+1] == 1 && locked[i] == 0 { // 再次尝试交换
//                     nums[i], nums[i + 1] = nums[i + 1], nums[i]
//                 } else { // 无法修复，返回 -1
//                     return -1
//                 }
//             }
//         }
//     }
//     for i := 0; i < n-1; i++ { // 检查是否排序成功
//         if nums[i] > nums[i + 1] {
//             return -1
//         }
//     }
//     return res
// }

func minUnlockedIndices(nums []int, locked []int) int {
    res, n := 0, len(nums)
    left3, left2, right1, right2 := n, n, -1, -1
    min := func(x, y int) int { if x < y { return x; }; return y; }
    sum := func(arr []int) int {
        res := 0
        for _, v := range arr {
            res += v
        }
        return res
    }
    for i, v := range nums {
        if v == 3 {
            left3 = min(left3, i)
        } else if v == 1 {
            right1 = i
        } else if v == 2 {
            left2, right2 = min(left2, i), i
        }
    }
    if right1 > left3 { return -1 }
    if left3 < right2 { res += sum(locked[left3:right2]) }
    if left2 < right1 { res += sum(locked[left2:right1]) }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,2,3,2], locked = [1,0,1,1,0,1]
    // Output: 0
    // Explanation:
    // We can sort nums using the following swaps:
    //     swap indices 1 with 2
    //     swap indices 4 with 5
    // So, there is no need to unlock any index.
    fmt.Println(minUnlockedIndices([]int{1,2,1,2,3,2}, []int{1,0,1,1,0,1})) // 0
    // Example 2:
    // Input: nums = [1,2,1,1,3,2,2], locked = [1,0,1,1,0,1,0]
    // Output: 2
    // Explanation:
    // If we unlock indices 2 and 5, we can sort nums using the following swaps:
    //     swap indices 1 with 2
    //     swap indices 2 with 3
    //     swap indices 4 with 5
    //     swap indices 5 with 6
    fmt.Println(minUnlockedIndices([]int{1,2,1,1,3,2,2}, []int{1,0,1,1,0,1,0})) // 2
    // Example 3:
    // Input: nums = [1,2,1,2,3,2,1], locked = [0,0,0,0,0,0,0]
    // Output: -1
    // Explanation:
    // Even if all indices are unlocked, it can be shown that nums is not sortable.
    fmt.Println(minUnlockedIndices([]int{1,2,1,2,3,2,1}, []int{0,0,0,0,0,0,0})) // -1

    fmt.Println(minUnlockedIndices([]int{2,2,1}, []int{0,0,0})) // 0
}