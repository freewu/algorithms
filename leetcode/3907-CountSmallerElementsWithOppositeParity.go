package main

// 3907. Count Smaller Elements With Opposite Parity
// You are given an integer array nums of length n.

// The score of an index i is defined as the number of indices j such that:
//     1. i < j < n
//     2. nums[j] < nums[i]
//     3. nums[i] and nums[j] have different parity (one is even and the other is odd).

// Return an integer array answer of length n, where answer[i] is the score of index i.

// Example 1:
// Input: nums = [5,2,4,1,3]
// Output: [2,1,2,0,0]
// Explanation:
// For i = 0, the elements nums[1] = 2 and nums[2] = 4 are smaller and have different parity.
// For i = 1, the element nums[3] = 1 is smaller and has different parity.
// For i = 2, the elements nums[3] = 1 and nums[4] = 3 are smaller and have different parity.
// No valid elements exist for the remaining indices.
// Thus, the answer = [2, 1, 2, 0, 0].

// Example 2:
// Input: nums = [4,4,1]
// Output: [1,1,0]
// Explanation:​​​​​​​
// For i = 0 and i = 1, the element nums[2] = 1 is smaller and has different parity. Thus, the answer = [1, 1, 0].

// Example 3:
// Input: nums = [7]
// Output: [0]
// Explanation:
// No elements exist to the right of index 0, so its score is 0. Thus, the answer = [0].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9​​​​​​​

import "fmt"
import "sort"

// 超出时间限制 552 / 598 
func countSmallerOppositeParity(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    for i := range nums {
        for j := i + 1; j < n; j++ {
            if nums[j] < nums[i] && (nums[i]&1) != (nums[j]&1) {
                res[i]++
            }
        }
    }
    return res
}

func countSmallerOppositeParity1(nums []int) []int {
    n := len(nums)
    res, evens, odds := make([]int, n), make([]int, 0), make([]int, 0) // 维护两个有序切片：evens存右侧偶数，odds存右侧奇数
    // 逆序遍历数组
    for i := n - 1; i >= 0; i-- {
        curr,count  := nums[i], 0
        isCurEven := curr % 2 == 0
        if isCurEven {
            // 当前是偶数，找右侧【奇数】中 < current 的数量
            count = sort.Search(len(odds), func(k int) bool { 
                return odds[k] >= curr 
            })
        } else {
            // 当前是奇数，找右侧【偶数】中 < current 的数量
            count = sort.Search(len(evens), func(k int) bool { 
                return evens[k] >= curr 
            })
        }
        res[i] = count
        // 将当前值插入对应有序切片
        if isCurEven {
            // 插入偶数有序切片
            index := sort.Search(len(evens), func(k int) bool { 
                return evens[k] >= curr 
            })
            // 切片插入：扩容+复制
            evens = append(evens, 0)
            copy(evens[index+1:], evens[index:])
            evens[index] = curr
        } else {
            // 插入奇数有序切片
            index := sort.Search(len(odds), func(k int) bool { 
                return odds[k] >= curr 
            })
            odds = append(odds, 0)
            copy(odds[index+1:], odds[index:])
            odds[index] = curr
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,2,4,1,3]
    // Output: [2,1,2,0,0]
    // Explanation:
    // For i = 0, the elements nums[1] = 2 and nums[2] = 4 are smaller and have different parity.
    // For i = 1, the element nums[3] = 1 is smaller and has different parity.
    // For i = 2, the elements nums[3] = 1 and nums[4] = 3 are smaller and have different parity.
    // No valid elements exist for the remaining indices.
    // Thus, the answer = [2, 1, 2, 0, 0].
    fmt.Println(countSmallerOppositeParity([]int{5,2,4,1,3})) // [2,1,2,0,0]
    // Example 2:
    // Input: nums = [4,4,1]
    // Output: [1,1,0]
    // Explanation:​​​​​​​
    // For i = 0 and i = 1, the element nums[2] = 1 is smaller and has different parity. Thus, the answer = [1, 1, 0].
    fmt.Println(countSmallerOppositeParity([]int{4,4,1})) // [1,1,0]
    // Example 3:
    // Input: nums = [7]
    // Output: [0]
    // Explanation:
    // No elements exist to the right of index 0, so its score is 0. Thus, the answer = [0].
    fmt.Println(countSmallerOppositeParity([]int{7})) // [0]

    fmt.Println(countSmallerOppositeParity([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0]
    fmt.Println(countSmallerOppositeParity([]int{9,8,7,6,5,4,3,2,1})) // [4 4 3 3 2 2 1 1 0]

    fmt.Println(countSmallerOppositeParity1([]int{5,2,4,1,3})) // [2,1,2,0,0]
    fmt.Println(countSmallerOppositeParity1([]int{4,4,1})) // [1,1,0]
    fmt.Println(countSmallerOppositeParity1([]int{7})) // [0]
    fmt.Println(countSmallerOppositeParity1([]int{1,2,3,4,5,6,7,8,9})) // [0 0 0 0 0 0 0 0 0]
    fmt.Println(countSmallerOppositeParity1([]int{9,8,7,6,5,4,3,2,1})) // [4 4 3 3 2 2 1 1 0]
}
