package main

// 2588. Count the Number of Beautiful Subarrays
// You are given a 0-indexed integer array nums. 
// In one operation, you can:
//     1. Choose two different indices i and j such that 0 <= i, j < nums.length.
//     2. Choose a non-negative integer k such that the kth bit (0-indexed) in the binary representation of nums[i] and nums[j] is 1.
//     3. Subtract 2k from nums[i] and nums[j].

// A subarray is beautiful if it is possible to make all of its elements equal to 0 after applying the above operation any number of times.

// Return the number of beautiful subarrays in the array nums.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [4,3,1,2,4]
// Output: 2
// Explanation: There are 2 beautiful subarrays in nums: [4,3,1,2,4] and [4,3,1,2,4].
// - We can make all elements in the subarray [3,1,2] equal to 0 in the following way:
//   - Choose [3, 1, 2] and k = 1. Subtract 21 from both numbers. The subarray becomes [1, 1, 0].
//   - Choose [1, 1, 0] and k = 0. Subtract 20 from both numbers. The subarray becomes [0, 0, 0].
// - We can make all elements in the subarray [4,3,1,2,4] equal to 0 in the following way:
//   - Choose [4, 3, 1, 2, 4] and k = 2. Subtract 22 from both numbers. The subarray becomes [0, 3, 1, 2, 0].
//   - Choose [0, 3, 1, 2, 0] and k = 0. Subtract 20 from both numbers. The subarray becomes [0, 2, 0, 2, 0].
//   - Choose [0, 2, 0, 2, 0] and k = 1. Subtract 21 from both numbers. The subarray becomes [0, 0, 0, 0, 0].

// Example 2:
// Input: nums = [1,10,4]
// Output: 0
// Explanation: There are no beautiful subarrays in nums.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^6

import "fmt"

func beautifulSubarrays(nums []int) int64 {
    count := map[int]int{ 0 : 1 }
    res, xor := 0, 0
    for _, v := range nums {
        xor ^= v
        res += count[xor]
        count[xor]++
    }
    return int64(res)
}

func beautifulSubarrays1(nums []int) int64 {
    count := make(map[int]int, len(nums) + 1)
    count[0] = 1
    res, sum := 0, 0
    for _, v := range nums {
        sum ^= v
        res += count[sum]
        count[sum]++
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [4,3,1,2,4]
    // Output: 2
    // Explanation: There are 2 beautiful subarrays in nums: [4,3,1,2,4] and [4,3,1,2,4].
    // - We can make all elements in the subarray [3,1,2] equal to 0 in the following way:
    //   - Choose [3, 1, 2] and k = 1. Subtract 21 from both numbers. The subarray becomes [1, 1, 0].
    //   - Choose [1, 1, 0] and k = 0. Subtract 20 from both numbers. The subarray becomes [0, 0, 0].
    // - We can make all elements in the subarray [4,3,1,2,4] equal to 0 in the following way:
    //   - Choose [4, 3, 1, 2, 4] and k = 2. Subtract 22 from both numbers. The subarray becomes [0, 3, 1, 2, 0].
    //   - Choose [0, 3, 1, 2, 0] and k = 0. Subtract 20 from both numbers. The subarray becomes [0, 2, 0, 2, 0].
    //   - Choose [0, 2, 0, 2, 0] and k = 1. Subtract 21 from both numbers. The subarray becomes [0, 0, 0, 0, 0].
    fmt.Println(beautifulSubarrays([]int{4,3,1,2,4})) // 2
    // Example 2:
    // Input: nums = [1,10,4]
    // Output: 0
    // Explanation: There are no beautiful subarrays in nums.
    fmt.Println(beautifulSubarrays([]int{1,10,4})) // 0

    fmt.Println(beautifulSubarrays([]int{1,2,3,4,5,6,7,8,9})) // 6
    fmt.Println(beautifulSubarrays([]int{9,8,7,6,5,4,3,2,1})) // 6

    fmt.Println(beautifulSubarrays1([]int{4,3,1,2,4})) // 2
    fmt.Println(beautifulSubarrays1([]int{1,10,4})) // 0
    fmt.Println(beautifulSubarrays1([]int{1,2,3,4,5,6,7,8,9})) // 6
    fmt.Println(beautifulSubarrays1([]int{9,8,7,6,5,4,3,2,1})) // 6
}