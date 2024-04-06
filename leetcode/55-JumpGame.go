package main

// 55. Jump Game
// You are given an integer array nums. 
// You are initially positioned at the array's first index, 
// and each element in the array represents your maximum jump length at that position.

// Return true if you can reach the last index, or false otherwise.

// Example 1:
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Example 2:
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: 
//     You will always arrive at index 3 no matter what. 
//     Its maximum jump length is 0, which makes it impossible to reach the last index.

// Constraints:
//     1 <= nums.length <= 10^4
//     0 <= nums[i] <= 10^5

// 解题思路:
// 	给出一个非负数组，要求判断从数组 0 下标开始，能否到达数组最后一个位置。
//     如果某一个作为 起跳点 的格子可以跳跃的距离是 n，那么表示后面 n 个格子都可以作为 起跳点。
//     可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离maxJump 不断更新。
//     如果可以一直跳到最后，就成功了。如果中间有一个点比 maxJump 还要大，说明在这个点和 maxJump 中间连不上了，有些点不能到达最后一个位置。

import "fmt"

func canJump(nums []int) bool {
    if len(nums) == 0 {
        return false
    }
    if len(nums) == 1 {
        return true
    }
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range nums {
        // 从 0 开始，所以第一次不会进入   如果可跳的最大值 < 当前位置  说明跳不到本位置 退出 返回 false
        if i > res { 
            return false
        }
        res = max(res, i + v) // i + v =   当前点 + 在当前点可以跳的最大值
    }
    return true
}

// best solution
func canJump1(nums []int) bool {
    steps := 0 // 所需步数
    // 从最后一个位置倒推
    for i := len(nums) - 2; i >= 0; i-- { // len(nums) - 2 最后一个位置
        steps++
        // 因为是 倒推 当前位置可跳步数  >= 需要的步数,说明到了这个位置 肯定能到最后 把 当前所需步数据 设置为 0
        if nums[i] >= steps { 
            steps = 0
        }
    }
    return steps == 0
}

func main() {
    fmt.Printf("canJump([]int{2,3,1,1,4} = %v\n",canJump([]int{2,3,1,1,4})) // true
    fmt.Printf("canJump([]int{3,2,1,0,4} = %v\n",canJump([]int{3,2,1,0,4})) // false

    fmt.Printf("canJump1([]int{2,3,1,1,4} = %v\n",canJump1([]int{2,3,1,1,4})) // true
    fmt.Printf("canJump1([]int{3,2,1,0,4} = %v\n",canJump1([]int{3,2,1,0,4})) // false
}
