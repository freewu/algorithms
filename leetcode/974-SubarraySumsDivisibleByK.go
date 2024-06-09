package main

// 974. Subarray Sums Divisible by K
// Given an integer array nums and an integer k, return the number of non-empty subarrays that have a sum divisible by k.
// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [4,5,0,-2,-3,1], k = 5
// Output: 7
// Explanation: There are 7 subarrays with a sum divisible by k = 5:
// [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

// Example 2:
// Input: nums = [5], k = 9
// Output: 0
 
// Constraints:
//     1 <= nums.length <= 3 * 10^4
//     -10^4 <= nums[i] <= 10^4
//     2 <= k <= 10^4

import "fmt"

func subarraysDivByK(nums []int, k int) int {
    prefixSum, res, remainderCount := 0, 0, make(map[int]int) // 记录每个余数出现的次数
    remainderCount[0] = 1
    for _, num := range nums {
        prefixSum = (prefixSum + num) % k
        if prefixSum < 0 {
            prefixSum += k  // 处理负数余数
        }
        if val, ok := remainderCount[prefixSum]; ok { // 如果存在相关余数的,说明可以凑成子数组
            res += val
            remainderCount[prefixSum]++
        } else {
            remainderCount[prefixSum] = 1
        }
    }
    return res
}

func subarraysDivByK1(nums []int, k int) int {
    res, n, mp := 0, len(nums), map[int]int{0 : 1} 
    for i := 1; i < n; i++ { // [1,n)
        nums[i] += nums[i-1]
    }
    for i := 0; i < n; i++ {
        mod := nums[i] % k 
        if mod < 0 {
            mod += k 
        }
        if v, ok := mp[mod]; ok {
            res += v
            mp[mod]++
        } else {
            mp[mod] = 1 
            //若pre[i]<=>nums[i], 对于  mod==0, 需要走到这里初始化一下, 所以定义mp时需要初始化一下mp[0]=1
            //若pre[i]<=>nums[i-1], 对于 mod==0, 因默认会有冗余一个pre[0], 遍历pre设置mp时, 多余的这个pre[0]会设置mp[0]=1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,5,0,-2,-3,1], k = 5
    // Output: 7
    // Explanation: There are 7 subarrays with a sum divisible by k = 5:
    // [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
    fmt.Println(subarraysDivByK([]int{4,5,0,-2,-3,1}, 5)) // 7
    // Example 2:
    // Input: nums = [5], k = 9
    // Output: 0
    fmt.Println(subarraysDivByK([]int{5}, 9)) // 0

    fmt.Println(subarraysDivByK1([]int{4,5,0,-2,-3,1}, 5)) // 7
    fmt.Println(subarraysDivByK1([]int{5}, 9)) // 0
}