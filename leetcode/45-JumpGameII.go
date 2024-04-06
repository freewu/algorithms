package main

// 45. Jump Game II
// You are given a 0-indexed array of integers nums of length n.
// You are initially positioned at nums[0].
// Each element nums[i] represents the maximum length of a forward jump from index i. 
// In other words, if you are at nums[i], you can jump to any nums[i + j] where:
//     0 <= j <= nums[i] and
//     i + j < n

// Return the minimum number of jumps to reach nums[n - 1]. 
// The test cases are generated such that you can reach nums[n - 1].

// Example 1:
// Input: nums = [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Example 2:
// Input: nums = [2,3,0,1,4]
// Output: 2
 
// Constraints:
//     1 <= nums.length <= 10^4
//     0 <= nums[i] <= 1000
//     It's guaranteed that you can reach nums[n - 1].

// 解题思路:
//     给定一个非负整数数组，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//     你的目标是使用最少的跳跃次数到达数组的最后一个位置。

import "fmt"

func jump(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    needChoose, canReach, step := 0, 0, 0
    for i, x := range nums {
        if i + x > canReach {
            canReach = i + x
            if canReach >= len(nums) - 1 {
                return step + 1
            }
        }
        if i == needChoose {
            needChoose = canReach
            step++
        }
    }
    return step
}

// best solution
func jump1(nums []int) int {
    l, res := len(nums), 0
    if l == 1 {
        return res
    }
    for i := 0; i < l; {
        res++
        m, mj := 0, 0
        for j := i + 1; j < i + 1 + nums[i] && j < l; j++ {
            val := j - i + nums[j]
            if m <= val {
                m, mj = val, j
            }
            if j == l - 1 {
                return res
            }
        }
        i = mj
        if mj >= l - 1 {
            break
        } else if mj + nums[mj] >= l - 1{
            res++
            break
        }
    }
    return res
}

func main() {
    // Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
	fmt.Printf("jump([]int{2,3,1,1,4}) = %v\n",jump([]int{2,3,1,1,4})) // 2 
	fmt.Printf("jump([]int{2,3,0,1,4}) = %v\n",jump([]int{2,3,0,1,4})) // 2 
	fmt.Printf("jump1([]int{2,3,1,1,4}) = %v\n",jump1([]int{2,3,1,1,4})) // 2 
	fmt.Printf("jump1([]int{2,3,0,1,4}) = %v\n",jump1([]int{2,3,0,1,4})) // 2 
}