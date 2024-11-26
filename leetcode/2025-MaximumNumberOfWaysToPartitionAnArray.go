package main

// 2025. Maximum Number of Ways to Partition an Array
// You are given a 0-indexed integer array nums of length n. 
// The number of ways to partition nums is the number of pivot indices that satisfy both conditions:
//     1 <= pivot < n
//     nums[0] + nums[1] + ... + nums[pivot - 1] == nums[pivot] + nums[pivot + 1] + ... + nums[n - 1]

// You are also given an integer k. 
// You can choose to change the value of one element of nums to k, or to leave the array unchanged.

// Return the maximum possible number of ways to partition nums to satisfy both conditions after changing at most one element.

// Example 1:
// Input: nums = [2,-1,2], k = 3
// Output: 1
// Explanation: One optimal approach is to change nums[0] to k. The array becomes [3,-1,2].
// There is one way to partition the array:
// - For pivot = 2, we have the partition [3,-1 | 2]: 3 + -1 == 2.

// Example 2:
// Input: nums = [0,0,0], k = 1
// Output: 2
// Explanation: The optimal approach is to leave the array unchanged.
// There are two ways to partition the array:
// - For pivot = 1, we have the partition [0 | 0,0]: 0 == 0 + 0.
// - For pivot = 2, we have the partition [0,0 | 0]: 0 + 0 == 0.

// Example 3:
// Input: nums = [22,4,-25,-20,-15,15,-16,7,19,-10,0,-13,-14], k = -33
// Output: 4
// Explanation: One optimal approach is to change nums[2] to k. The array becomes [22,4,-33,-20,-15,15,-16,7,19,-10,0,-13,-14].
// There are four ways to partition the array.

// Constraints:
//     n == nums.length
//     2 <= n <= 10^5
//     -10^5 <= k, nums[i] <= 10^5

import "fmt"

func waysToPartition(nums []int, k int) int {
    sum, n, diffFreq := 0, len(nums), make(map[int]int)
    for _, v := range nums {
        sum += v
    }
    leftSum := 0
    for i := 0; i < n - 1; i++ {
        leftSum += nums[i]
        diffFreq[2 * leftSum - sum]++
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, leftFreq := diffFreq[0], make(map[int]int)
    leftSum = 0
    for i := 0; i < n; i++  {
        delta := k - nums[i]
        res = max(res, leftFreq[delta] + diffFreq[-delta])
        leftSum += nums[i]
        diff := 2 * leftSum - sum
        leftFreq[diff]++
        diffFreq[diff]--
    }
    return res
}

func waysToPartition1(nums []int, k int) int {
    res, n := 0, len(nums)
    s := make([]int, n)
    s[0] = nums[0]
    right, left := make(map[int]int),  make(map[int]int)
    for i := range nums[:n - 1] {
        right[s[i]]++
        s[i + 1] = s[i] + nums[i+1]
    }
    if s[n - 1] % 2 == 0 {
        res = right[s[n - 1] / 2]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums {
        diff := k - v
        if (s[n-1] + diff) % 2 == 0 {
            res = max(res, left[(s[n - 1] + diff)/2] + right[(s[n - 1] - diff) / 2])
        }
        left[s[i]]++
        right[s[i]]--
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,-1,2], k = 3
    // Output: 1
    // Explanation: One optimal approach is to change nums[0] to k. The array becomes [3,-1,2].
    // There is one way to partition the array:
    // - For pivot = 2, we have the partition [3,-1 | 2]: 3 + -1 == 2.
    fmt.Println(waysToPartition([]int{2,-1,2}, 3)) // 1
    // Example 2:
    // Input: nums = [0,0,0], k = 1
    // Output: 2
    // Explanation: The optimal approach is to leave the array unchanged.
    // There are two ways to partition the array:
    // - For pivot = 1, we have the partition [0 | 0,0]: 0 == 0 + 0.
    // - For pivot = 2, we have the partition [0,0 | 0]: 0 + 0 == 0.
    fmt.Println(waysToPartition([]int{0,0,0}, 1)) // 2
    // Example 3:
    // Input: nums = [22,4,-25,-20,-15,15,-16,7,19,-10,0,-13,-14], k = -33
    // Output: 4
    // Explanation: One optimal approach is to change nums[2] to k. The array becomes [22,4,-33,-20,-15,15,-16,7,19,-10,0,-13,-14].
    // There are four ways to partition the array.
    fmt.Println(waysToPartition([]int{22,4,-25,-20,-15,15,-16,7,19,-10,0,-13,-14}, -33)) // 4

    fmt.Println(waysToPartition1([]int{2,-1,2}, 3)) // 1
    fmt.Println(waysToPartition1([]int{0,0,0}, 1)) // 2
    fmt.Println(waysToPartition1([]int{22,4,-25,-20,-15,15,-16,7,19,-10,0,-13,-14}, -33)) // 4
}