package main

// 2859. Sum of Values at Indices With K Set Bits 
// You are given a 0-indexed integer array nums and an integer k.
// Return an integer that denotes the sum of elements in nums whose corresponding indices have exactly k set bits in their binary representation.
// The set bits in an integer are the 1's present when it is written in binary.
// For example, the binary representation of 21 is 10101, which has 3 set bits.

// Example 1:
// Input: nums = [5,10,1,5,2], k = 1
// Output: 13
// Explanation: The binary representation of the indices are: 
// 0 = 000
// 1 = 001
// 2 = 010
// 3 = 011
// 4 = 100
// Indices 1, 2, and 4 have k = 1 set bits in their binary representation.
// Hence, the answer is nums[1] + nums[2] + nums[4] = 13.

// Example 2:
// Input: nums = [4,3,2,1], k = 2
// Output: 1
// Explanation: The binary representation of the indices are:
// 0 = 00
// 1 = 01
// 2 = 10
// 3 = 11
// Only index 3 has k = 2 set bits in its binary representation.
// Hence, the answer is nums[3] = 1.
 
// Constraints:
//         1 <= nums.length <= 1000
//         1 <= nums[i] <= 10^5
//         0 <= k <= 10

import "fmt"

func sumIndicesWithKSetBits(nums []int, k int) int {
    if k == 0 {
        return nums[0]
    }
    nums[0] = 0
    for i := 1 << k - 1; i < len(nums); {
        nums[0] += nums[i]
        b := i & -i
        c := i + b
        i = (c^i) >> 2 / b | c
    }
    return nums[0]
}

func sumIndicesWithKSetBits1(nums []int, k int) int {
    var sum int
    for i := 0; i < len(nums); i++ {
       num := i
       loc := 0
       // 计算 1 的个数
       for num != 0 {
          loc += num % 2
          num = num >> 1
       }
       // 符合标准累加起来
       if loc == k {
           sum += nums[i]
       }
    }
    return sum
}

// use bits
// func sumIndicesWithKSetBits2(nums []int, k int) int {
//     sum := 0
//     for i, v := range nums {
//         if bits.OnesCount(uint(i)) == k {
//             sum += v
//         }
//     }
//     return sum
// }

func main() {
    fmt.Println(sumIndicesWithKSetBits([]int{5,10,1,5,2}, 1)) // 13
    fmt.Println(sumIndicesWithKSetBits([]int{4,3,2,1}, 2)) // 1

    fmt.Println(sumIndicesWithKSetBits1([]int{5,10,1,5,2}, 1)) // 13
    fmt.Println(sumIndicesWithKSetBits1([]int{4,3,2,1}, 2)) // 1
}