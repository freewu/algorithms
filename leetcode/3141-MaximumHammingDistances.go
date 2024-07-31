package main

// 3141. Maximum Hamming Distances
// Given an array nums and an integer m, with each element nums[i] satisfying 0 <= nums[i] < 2m, return an array answer. 
// The answer array should be of the same length as nums, where each element answer[i] represents the maximum Hamming distance between nums[i] and any other element nums[j] in the array.

// The Hamming distance between two binary integers is defined as the number of positions at which the corresponding bits differ (add leading zeroes if needed).

// Example 1:
// Input: nums = [9,12,9,11], m = 4
// Output: [2,3,2,3]
// Explanation:
// The binary representation of nums = [1001,1100,1001,1011].
// The maximum hamming distances for each index are:
// nums[0]: 1001 and 1100 have a distance of 2.
// nums[1]: 1100 and 1011 have a distance of 3.
// nums[2]: 1001 and 1100 have a distance of 2.
// nums[3]: 1011 and 1100 have a distance of 3.

// Example 2:
// Input: nums = [3,4,6,10], m = 4
// Output: [3,3,2,3]
// Explanation:
// The binary representation of nums = [0011,0100,0110,1010].
// The maximum hamming distances for each index are:
// nums[0]: 0011 and 0100 have a distance of 3.
// nums[1]: 0100 and 0011 have a distance of 3.
// nums[2]: 0110 and 1010 have a distance of 2.
// nums[3]: 1010 and 0100 have a distance of 3.

// Constraints:
//     1 <= m <= 17
//     2 <= nums.length <= 2^m
//     0 <= nums[i] < 2^m

import "fmt"

func maxHammingDistances(nums []int, m int) []int {
    maxNum, inf := (1 << m) - 1, 1 << 32 - 1
    distances := make([]int, maxNum + 1)
    for i := 0; i < len(distances); i++ {
        distances[i] = inf
    }
    queue := []int{}
    for _, v := range nums {
        if distances[v] > 0 {
            distances[v] = 0
            queue = append(queue, v)
        }
    }
    for len(queue) > 0 {
        num := queue[0] // 队列取出 Dequeue
        queue = queue[1:]
        for i := 0; i < m; i++ {
            next := num ^ (1 << i)
            if distances[num] + 1 < distances[next] {
                distances[next] = distances[num] + 1
                queue = append(queue, next) // Enqueue
            }
        }
    }
    n := len(nums)
    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = m - distances[nums[i] ^ maxNum]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [9,12,9,11], m = 4
    // Output: [2,3,2,3]
    // Explanation:
    // The binary representation of nums = [1001,1100,1001,1011].
    // The maximum hamming distances for each index are:
    // nums[0]: 1001 and 1100 have a distance of 2.
    // nums[1]: 1100 and 1011 have a distance of 3.
    // nums[2]: 1001 and 1100 have a distance of 2.
    // nums[3]: 1011 and 1100 have a distance of 3.
    fmt.Println(maxHammingDistances([]int{9,12,9,11}, 4)) // [2,3,2,3]
    // Example 2:
    // Input: nums = [3,4,6,10], m = 4
    // Output: [3,3,2,3]
    // Explanation:
    // The binary representation of nums = [0011,0100,0110,1010].
    // The maximum hamming distances for each index are:
    // nums[0]: 0011 and 0100 have a distance of 3.
    // nums[1]: 0100 and 0011 have a distance of 3.
    // nums[2]: 0110 and 1010 have a distance of 2.
    // nums[3]: 1010 and 0100 have a distance of 3.
    fmt.Println(maxHammingDistances([]int{3,4,6,10}, 4)) // [3,3,2,3]
}