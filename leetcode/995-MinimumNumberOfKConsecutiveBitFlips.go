package main

// 995. Minimum Number of K Consecutive Bit Flips
// You are given a binary array nums and an integer k.
// A k-bit flip is choosing a subarray of length k from nums and simultaneously changing every 0 in the subarray to 1, and every 1 in the subarray to 0.

// Return the minimum number of k-bit flips required so that there is no 0 in the array. If it is not possible, return -1.
// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [0,1,0], k = 1
// Output: 2
// Explanation: Flip nums[0], then flip nums[2].

// Example 2:
// Input: nums = [1,1,0], k = 2
// Output: -1
// Explanation: No matter how we flip subarrays of size 2, we cannot make the array become [1,1,1].

// Example 3:
// Input: nums = [0,0,0,1,0,1,1,0], k = 3
// Output: 3
// Explanation: 
// Flip nums[0],nums[1],nums[2]: nums becomes [1,1,1,1,0,1,1,0]
// Flip nums[4],nums[5],nums[6]: nums becomes [1,1,1,1,1,0,0,0]
// Flip nums[5],nums[6],nums[7]: nums becomes [1,1,1,1,1,1,1,1]
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= k <= nums.length

import "fmt"

func minKBitFlips(nums []int, k int) int {
    n, totalCost, lastFlipTime := len(nums), 0, 0
    hasFlip := make([]int, n)
    for i, v := range nums {
        if i >= k { // 首先计算当前节点被翻转过的次数
            lastFlipTime = lastFlipTime ^ hasFlip[i-k]
        }
        // 然后判断当前节点是否需要翻转
        if (v == 0 && lastFlipTime == 0) || (v == 1 && lastFlipTime == 1) {
            if n - i < k { // 特殊情况: 最后剩下的元素数量已经不及K 个，但是依然要翻转，直接返回不可能
                return -1
            }
            totalCost++
            lastFlipTime ^= 1
            hasFlip[i] = 1
        }
    }
    return totalCost
}

func minKBitFlips1(nums []int, k int) int {
    n, count, res := len(nums), 0, 0
    diff := make([]int, n+1)
    for i := 0; i < n; i++ {
        count += diff[i]
        // 首先我们明白0要变成1那么它的翻转次数一定是奇数次，1要变成0那么它的翻转次数一定是偶数次
        // 从前往后遍历，每次至少有一个元素翻转成功（k个元素中的首元素）
        // count 表示的是截止目前 nums[i] 元素的翻转次数，也是 nums[i−1] 翻转成功需要的次数
        // 所以当nums[i] + count 是偶数时，一定是未能成功翻转的状态，因此需要进行一次即可成功
        if (count + nums[i]) % 2 == 0 {
            if i + k > n {
                return -1
            }
            res++
            diff[i+1]++
            diff[i+k]--
        }
    }
    return res
}

func minKBitFlips2(nums []int, k int) int {
    n, res, b := len(nums), 0, 0
    for i := 0; i < n; i++ {
        if i - k >= 0 && nums[i-k] > 1 {
            b ^= 1
        }
        if nums[i] == b {
            if i + k > n {
                return -1
            }
            res++
            b ^= 1
            nums[i] += 2
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [0,1,0], k = 1
    // Output: 2
    // Explanation: Flip nums[0], then flip nums[2].
    fmt.Println(minKBitFlips([]int{0,1,0}, 1)) // 2
    // Example 2:
    // Input: nums = [1,1,0], k = 2
    // Output: -1
    // Explanation: No matter how we flip subarrays of size 2, we cannot make the array become [1,1,1].
    fmt.Println(minKBitFlips([]int{1,1,0}, 2)) // -1
    // Example 3:
    // Input: nums = [0,0,0,1,0,1,1,0], k = 3
    // Output: 3
    // Explanation: 
    // Flip nums[0],nums[1],nums[2]: nums becomes [1,1,1,1,0,1,1,0]
    // Flip nums[4],nums[5],nums[6]: nums becomes [1,1,1,1,1,0,0,0]
    // Flip nums[5],nums[6],nums[7]: nums becomes [1,1,1,1,1,1,1,1]
    fmt.Println(minKBitFlips([]int{0,0,0,1,0,1,1,0}, 3)) // 3

    fmt.Println(minKBitFlips1([]int{0,1,0}, 1)) // 2
    fmt.Println(minKBitFlips1([]int{1,1,0}, 2)) // -1
    fmt.Println(minKBitFlips1([]int{0,0,0,1,0,1,1,0}, 3)) // 3

    fmt.Println(minKBitFlips2([]int{0,1,0}, 1)) // 2
    fmt.Println(minKBitFlips2([]int{1,1,0}, 2)) // -1
    fmt.Println(minKBitFlips2([]int{0,0,0,1,0,1,1,0}, 3)) // 3
}