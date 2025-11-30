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
    prefix := make([]int,n + 1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    res, rem := inf, prefix[n] % p
    if rem == 0 {
        return 0
    }
    count := make(map[int]int)
    for i := 0;i < n + 1; i++ {
        t := (prefix[i] - rem) % p 
        if  index ,ok := count[t]; ok {
            res = min(res, i - index)
        }
        count[prefix[i] % p] = i
    }
    if res == inf || res == len(nums) {
        return -1
    }
    return res 
}

func minSubarray2(nums []int, p int) int {
    res, sum, n := 1 << 31, 0, len(nums)
    // 整个数组的 sum%p = x, x=0 说明不用移除元素
    // 去掉数组中的 一部分和%p = x, 说明移除这部分%p=0
    // (后面的和+p)%p-(前面的和+p)%p == x
    // s - z = x
    for _, v := range nums {
        sum = (sum + v) % p
    }
    x, s := sum % p, 0
    if x == 0 { return 0 }
    mp := make(map[int]int) // 前缀和出现k的最近一次下标v
    mp[0] = -1              // 需要这个 前缀和为0的
    for i := 0; i < n; i++ {
        s = (s + nums[i] + p) % p
        if pos, ok := mp[(s-x+p)%p]; ok { // 去掉(pos,i]
            if i-pos != n {
                res = min(res, i-pos)
            }
        }
        mp[s] = i   
    }
    if res == 1 << 31 {
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

    fmt.Println(minSubarray([]int{1,2,3,4,5,6,7,8,9}, 3)) // 0
    fmt.Println(minSubarray([]int{9,8,7,6,5,4,3,2,1}, 3)) // 0

    fmt.Println(minSubarray1([]int{3,1,4,2}, 6)) // 1
    fmt.Println(minSubarray1([]int{6,3,5,2}, 9)) // 2
    fmt.Println(minSubarray1([]int{1,2,3}, 3)) // 0
    fmt.Println(minSubarray1([]int{1,2,3,4,5,6,7,8,9}, 3)) // 0
    fmt.Println(minSubarray1([]int{9,8,7,6,5,4,3,2,1}, 3)) // 0

    fmt.Println(minSubarray2([]int{3,1,4,2}, 6)) // 1
    fmt.Println(minSubarray2([]int{6,3,5,2}, 9)) // 2
    fmt.Println(minSubarray2([]int{1,2,3}, 3)) // 0
    fmt.Println(minSubarray2([]int{1,2,3,4,5,6,7,8,9}, 3)) // 0
    fmt.Println(minSubarray2([]int{9,8,7,6,5,4,3,2,1}, 3)) // 0
}