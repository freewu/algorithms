package main

// 1590. Make Sum Divisible by P
// Given an array of positive integers nums, remove the smallest subarray (possibly empty) 
// such that the sum of the remaining elements is divisible by p. 
// It is not allowed to remove the whole array.

// Return the length of the smallest subarray that you need to remove, or -1 if it's impossible.

// A subarray is defined as a contiguous block of elements in the array.

// Example 1:
// Input: nums = [3,1,4,2], p = 6
// Output: 1
// Explanation: The sum of the elements in nums is 10, which is not divisible by 6. We can remove the subarray [4], and the sum of the remaining elements is 6, which is divisible by 6.

// Example 2:
// Input: nums = [6,3,5,2], p = 9
// Output: 2
// Explanation: We cannot remove a single element to get a sum divisible by 9. The best way is to remove the subarray [5,2], leaving us with [6,3] with sum 9.

// Example 3:
// Input: nums = [1,2,3], p = 3
// Output: 0
// Explanation: Here the sum is 6. which is already divisible by 3. Thus we do not need to remove anything.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9
//     1 <= p <= 10^9

import "fmt"

func minSubarray(nums []int, p int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    remainder := sum % p // Find the remainder of the total sum divided by p
    if remainder == 0 {
        return 0 // Already divisible by p, no need to remove anything
    }
    // We want to find the smallest subarray with sum % p == remainder
    subarraySum, res, lastOccurrence := 0, len(nums), make(map[int]int)
    lastOccurrence[0] = -1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, v := range nums {
        subarraySum = (subarraySum + v) % p
        neededRemainder := (subarraySum - remainder + p) % p
        if index, ok := lastOccurrence[neededRemainder]; ok {
            res = min(res, i - index)
        }
        lastOccurrence[subarraySum] = i
    }
    if res == len(nums) {
        return -1 // No valid subarray found
    }
    return res
}

func minSubarray1(nums []int, p int) int {
    n, inf := len(nums), 1 << 31
    prefixSum := make([]int,n + 1)
    for i := 0; i < n; i++ {
        prefixSum[i+1] = prefixSum[i] + nums[i]
    }
    remainder := prefixSum[n] % p
    if remainder == 0 {
        return 0
    }
    count, res := make(map[int]int), inf
    for i := 0;i < n + 1; i++ {
        t := (prefixSum[i] - remainder) % p 
        if  index ,ok := count[t]; ok {
            res = min(res, i - index)
        }
        count[prefixSum[i] % p] = i
    }
    if res == inf || res == len(nums) {
        return -1
    }
    return res 
}

func main() {
    // Example 1:
    // Input: nums = [3,1,4,2], p = 6
    // Output: 1
    // Explanation: The sum of the elements in nums is 10, which is not divisible by 6. We can remove the subarray [4], and the sum of the remaining elements is 6, which is divisible by 6.
    fmt.Println(minSubarray([]int{3,1,4,2}, 6)) // 1
    // Example 2:
    // Input: nums = [6,3,5,2], p = 9
    // Output: 2
    // Explanation: We cannot remove a single element to get a sum divisible by 9. The best way is to remove the subarray [5,2], leaving us with [6,3] with sum 9.
    fmt.Println(minSubarray([]int{6,3,5,2}, 9)) // 2
    // Example 3:
    // Input: nums = [1,2,3], p = 3
    // Output: 0
    // Explanation: Here the sum is 6. which is already divisible by 3. Thus we do not need to remove anything.
    fmt.Println(minSubarray([]int{1,2,3}, 3)) // 0

    fmt.Println(minSubarray1([]int{3,1,4,2}, 6)) // 1
    fmt.Println(minSubarray1([]int{6,3,5,2}, 9)) // 2
    fmt.Println(minSubarray1([]int{1,2,3}, 3)) // 0
}