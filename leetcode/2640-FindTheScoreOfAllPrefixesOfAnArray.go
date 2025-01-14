package main

// 2640. Find the Score of All Prefixes of an Array
// We define the conversion array conver of an array arr as follows:
//     conver[i] = arr[i] + max(arr[0..i]) where max(arr[0..i]) is the maximum value of arr[j] over 0 <= j <= i.

// We also define the score of an array arr as the sum of the values of the conversion array of arr.

// Given a 0-indexed integer array nums of length n, 
// return an array ans of length n where ans[i] is the score of the prefix nums[0..i].

// Example 1:
// Input: nums = [2,3,7,5,10]
// Output: [4,10,24,36,56]
// Explanation: 
// For the prefix [2], the conversion array is [4] hence the score is 4
// For the prefix [2, 3], the conversion array is [4, 6] hence the score is 10
// For the prefix [2, 3, 7], the conversion array is [4, 6, 14] hence the score is 24
// For the prefix [2, 3, 7, 5], the conversion array is [4, 6, 14, 12] hence the score is 36
// For the prefix [2, 3, 7, 5, 10], the conversion array is [4, 6, 14, 12, 20] hence the score is 56

// Example 2:
// Input: nums = [1,1,2,4,8,16]
// Output: [2,4,8,16,32,64]
// Explanation: 
// For the prefix [1], the conversion array is [2] hence the score is 2
// For the prefix [1, 1], the conversion array is [2, 2] hence the score is 4
// For the prefix [1, 1, 2], the conversion array is [2, 2, 4] hence the score is 8
// For the prefix [1, 1, 2, 4], the conversion array is [2, 2, 4, 8] hence the score is 16
// For the prefix [1, 1, 2, 4, 8], the conversion array is [2, 2, 4, 8, 16] hence the score is 32
// For the prefix [1, 1, 2, 4, 8, 16], the conversion array is [2, 2, 4, 8, 16, 32] hence the score is 64

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func findPrefixScore(nums []int) []int64 {
    res := []int64{}
    mx, cur := -1,  0
    for _, v := range nums {
        if mx < v { mx = v }
        t := v + mx
        t += cur
        cur = t
        res = append(res, int64(t))
    }
    return res
}

func findPrefixScore1(nums []int) []int64 {
    res, mx := make([]int64, len(nums) + 1), int64(0)
    for i := 1; i < len(res); i++ {
        if int64(nums[i-1]) > mx {
            mx = int64(nums[i-1])
        }
        res[i] = res[i-1] + int64(nums[i-1]) + mx
    }
    return res[1:]
}

func main() {
    // Example 1:
    // Input: nums = [2,3,7,5,10]
    // Output: [4,10,24,36,56]
    // Explanation: 
    // For the prefix [2], the conversion array is [4] hence the score is 4
    // For the prefix [2, 3], the conversion array is [4, 6] hence the score is 10
    // For the prefix [2, 3, 7], the conversion array is [4, 6, 14] hence the score is 24
    // For the prefix [2, 3, 7, 5], the conversion array is [4, 6, 14, 12] hence the score is 36
    // For the prefix [2, 3, 7, 5, 10], the conversion array is [4, 6, 14, 12, 20] hence the score is 56
    fmt.Println(findPrefixScore([]int{2,3,7,5,10})) // [4,10,24,36,56]
    // Example 2:
    // Input: nums = [1,1,2,4,8,16]
    // Output: [2,4,8,16,32,64]
    // Explanation: 
    // For the prefix [1], the conversion array is [2] hence the score is 2
    // For the prefix [1, 1], the conversion array is [2, 2] hence the score is 4
    // For the prefix [1, 1, 2], the conversion array is [2, 2, 4] hence the score is 8
    // For the prefix [1, 1, 2, 4], the conversion array is [2, 2, 4, 8] hence the score is 16
    // For the prefix [1, 1, 2, 4, 8], the conversion array is [2, 2, 4, 8, 16] hence the score is 32
    // For the prefix [1, 1, 2, 4, 8, 16], the conversion array is [2, 2, 4, 8, 16, 32] hence the score is 64
    fmt.Println(findPrefixScore([]int{1,1,2,4,8,16})) // [2,4,8,16,32,64]

    fmt.Println(findPrefixScore([]int{1,2,3,4,5,6,7,8,9})) // [2 6 12 20 30 42 56 72 90]
    fmt.Println(findPrefixScore([]int{9,8,7,6,5,4,3,2,1})) // [18 35 51 66 80 93 105 116 126]

    fmt.Println(findPrefixScore1([]int{2,3,7,5,10})) // [4,10,24,36,56]
    fmt.Println(findPrefixScore1([]int{1,1,2,4,8,16})) // [2,4,8,16,32,64]
    fmt.Println(findPrefixScore1([]int{1,2,3,4,5,6,7,8,9})) // [2 6 12 20 30 42 56 72 90]
    fmt.Println(findPrefixScore1([]int{9,8,7,6,5,4,3,2,1})) // [18 35 51 66 80 93 105 116 126]
}