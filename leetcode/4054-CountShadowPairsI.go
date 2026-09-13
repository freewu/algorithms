package main

// 4054. Count Shadow Pairs I
// You are given an integer array nums of length n.

// A pair of indices (i, j) is called a shadow pair if all of the following conditions are satisfied:
//     1. 0 <= i < j < n
//     2. nums[i] < nums[j]
//     3. There does not exist an index k such that i < k < j and nums[k] < nums[i] < nums[j].

// Return the total number of shadow pairs.

// Example 1:
// Input: nums = [3,1,4,1,5]
// Output: 3
// Explanation:
// (i, j) | nums[i] | nums[j] | Shadow Pair
// (1, 2) | 1       | 4       | No index k exists such that 1 < k < 2
// (1, 4) | 1       | 5       | nums[2] = 4 and nums[3] = 1 are not smaller than 1
// (3, 4) | 3       | 5       | No index k exists such that 3 < k < 4
// Thus, the answer is 3.

// Example 2:
// Input: nums = [6,7,6,6,7]
// Output: 4
// Explanation:
// (i, j) | nums[i] | nums[j] | Shadow Pair
// (0, 1) | 6       | 7       | No index k exists such that 0 < k < 1
// (0, 4) | 6       | 7       | nums[1] = 7, nums[2] = 6, and nums[3] = 6 are not smaller than 6
// (2, 4) | 6       | 7       | nums[3] = 6 is not smaller than 6
// (3, 4) | 6       | 7       | No index k exists such that 3 < k < 4
// Thus, the answer is 4.

// Example 3:
// Input: nums = [1,2,3,4]
// Output: 6
// Explanation:
// (i, j) | nums[i] | nums[j] | Shadow Pair
// (0, 1) | 1       | 2       | No index k exists such that 0 < k < 1
// (0, 2) | 1       | 3       | nums[1] = 2 is not smaller than 1
// (0, 3) | 1       | 4       | nums[1] = 2 and nums[2] = 3 are not smaller than 1 
// (1, 2) | 2       | 3       | No index k exists such that 1 < k < 2
// (1, 3) | 2       | 4       | nums[2] = 3 is not smaller than 2
// (2, 3) | 3       | 4       | No index k exists such that 2 < k < 3
// Thus, the answer is 6.

// Constraints:
//     3 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

// 单调栈
func shadowPairs(nums []int) int64 {
    type Pair struct{ val, count int }
    stack := []Pair{{}} // 栈底哨兵
    res, n := int64(0), 0 // 栈的大小（count 之和） 
    for _, v := range nums { // 遍历的元素
        for stack[len(stack)-1].val > v {
            n -= stack[len(stack)-1].count
            stack = stack[:len(stack)-1] // 栈顶永远无法构成影子对
        }
        res += int64(n)
        if stack[len(stack)-1].val == v {
            // 刚好等于 val 的 nums[i] 不能构成影子对，要减掉
            res -= int64(stack[len(stack)-1].count)
            stack[len(stack)-1].count++
        } else {
            stack = append(stack, Pair{v, 1})
        }
        n++
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,1,4,1,5]
    // Output: 3
    // Explanation:
    // (i, j) | nums[i] | nums[j] | Shadow Pair
    // (1, 2) | 1       | 4       | No index k exists such that 1 < k < 2
    // (1, 4) | 1       | 5       | nums[2] = 4 and nums[3] = 1 are not smaller than 1
    // (3, 4) | 3       | 5       | No index k exists such that 3 < k < 4
    // Thus, the answer is 3.
    fmt.Println(shadowPairs([]int{3,1,4,1,5})) // 3
    // Example 2:
    // Input: nums = [6,7,6,6,7]
    // Output: 4
    // Explanation:
    // (i, j) | nums[i] | nums[j] | Shadow Pair
    // (0, 1) | 6       | 7       | No index k exists such that 0 < k < 1
    // (0, 4) | 6       | 7       | nums[1] = 7, nums[2] = 6, and nums[3] = 6 are not smaller than 6
    // (2, 4) | 6       | 7       | nums[3] = 6 is not smaller than 6
    // (3, 4) | 6       | 7       | No index k exists such that 3 < k < 4
    // Thus, the answer is 4.
    fmt.Println(shadowPairs([]int{6,7,6,6,7})) // 4
    // Example 3:
    // Input: nums = [1,2,3,4]
    // Output: 6
    // Explanation:
    // (i, j) | nums[i] | nums[j] | Shadow Pair
    // (0, 1) | 1       | 2       | No index k exists such that 0 < k < 1
    // (0, 2) | 1       | 3       | nums[1] = 2 is not smaller than 1
    // (0, 3) | 1       | 4       | nums[1] = 2 and nums[2] = 3 are not smaller than 1 
    // (1, 2) | 2       | 3       | No index k exists such that 1 < k < 2
    // (1, 3) | 2       | 4       | nums[2] = 3 is not smaller than 2
    // (2, 3) | 3       | 4       | No index k exists such that 2 < k < 3
    // Thus, the answer is 6.
    fmt.Println(shadowPairs([]int{1,2,3,4})) // 6

    fmt.Println(shadowPairs([]int{1,2,3,4,5,6,7,8,9})) // 36
    fmt.Println(shadowPairs([]int{9,8,7,6,5,4,3,2,1})) // 0
}