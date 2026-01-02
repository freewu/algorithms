package main

// 961. N-Repeated Element in Size 2N Array
// You are given an integer array nums with the following properties:
//     nums.length == 2 * n.
//     nums contains n + 1 unique elements.
//     Exactly one element of nums is repeated n times.

// Return the element that is repeated n times.

// Example 1:
// Input: nums = [1,2,3,3]
// Output: 3

// Example 2:
// Input: nums = [2,1,2,5,3,2]
// Output: 2

// Example 3:
// Input: nums = [5,1,5,2,5,3,5,4]
// Output: 5

// Constraints:
//     2 <= n <= 5000
//     nums.length == 2 * n
//     0 <= nums[i] <= 10^4
//     nums contains n + 1 unique elements and one of them is repeated exactly n times.

import "fmt"

func repeatedNTimes(nums []int) int {
    n, mp := len(nums), make(map[int]int)
    for _, v := range nums {
        mp[v]++
    }
    for k, v := range mp {
        if v == n / 2 {
            return k
        }
    }
    return -1
}

func repeatedNTimes1(nums []int) int {
    arr := make([]int, 10000)
    for i := 0; i < len(nums); i++ {
        arr[nums[i]]++  
        if arr[nums[i]] > 1 { // nums 中恰有一个元素重复 n 次 || nums 包含 n + 1 个 不同的 元素
            return nums[i]
        }
    }
    return 0
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,3]
    // Output: 3
    fmt.Println(repeatedNTimes([]int{1,2,3,3})) // 3
    // Example 2: 
    // Input: nums = [2,1,2,5,3,2]
    // Output: 2
    fmt.Println(repeatedNTimes([]int{2,1,2,5,3,2})) // 2
    // Example 3:
    // Input: nums = [5,1,5,2,5,3,5,4]
    // Output: 5
    fmt.Println(repeatedNTimes([]int{5,1,5,2,5,3,5,4})) // 5

    fmt.Println(repeatedNTimes([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(repeatedNTimes([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(repeatedNTimes1([]int{1,2,3,3})) // 3
    fmt.Println(repeatedNTimes1([]int{2,1,2,5,3,2})) // 2
    fmt.Println(repeatedNTimes1([]int{5,1,5,2,5,3,5,4})) // 5
    fmt.Println(repeatedNTimes1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(repeatedNTimes1([]int{9,8,7,6,5,4,3,2,1})) // 0
}