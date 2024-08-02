package main

// LCR 090. 打家劫舍 II
// 一个专业的小偷，计划偷窃一个环形街道上沿街的房屋，每间房内都藏有一定的现金。
// 这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
// 同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

// 给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

// 示例 1：
// 输入：nums = [2,3,2]
// 输出：3
// 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

// 示例 2：
// 输入：nums = [1,2,3,1]
// 输出：4
// 解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//      偷窃到的最高金额 = 1 + 3 = 4 。

// 示例 3：
// 输入：nums = [0]
// 输出：0

// 提示：
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 1000

// # 解题思路 #
//     在一个环形的街道中，即最后一个元素和第一个元素是邻居
//     由于首尾是相邻的，所以在取了第一个房子以后就不能取第 n 个房子，
//     那么就在 [0,n - 1] 的区间内找出总价值最多的解，然后再 [1,n] 的区间内找出总价值最多的解，两者取最大值即可。

import "fmt"

func rob(nums []int) int {
    n := len(nums)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    if n == 0 { return 0 }
    if n == 1 { return nums[0] }
    if n == 2 { return max(nums[0], nums[1]) }
    helper := func (nums []int, start, end int) int {
        preMax := nums[start]
        curMax := max(preMax, nums[start+1])
        for i := start + 2; i <= end; i++ {
            tmp := curMax
            curMax = max(curMax, nums[i] + preMax)
            preMax = tmp
        }
        return curMax
    }
    return max(helper(nums, 0, n - 2), helper(nums, 1, n - 1)) // 由于首尾是相邻的，所以需要对比 [0，n-1]、[1，n] 这两个区间的最大值
}

// best solution
func rob1(nums []int) int {
    if len(nums) == 1 { return nums[0] }
    p1, p2, mx := 0, 0, 0
    for i := 0; i < len(nums) - 1; i += 1 {
        tmp := p1
        if p2 + nums[i] > p1 {
            tmp = p2 + nums[i]
        }
        p2 = p1
        p1 = tmp
    }
    mx = p1
    p1, p2 = 0, 0
    for i := 1; i < len(nums); i += 1 {
        tmp := p1
        if p2 + nums[i] > p1 {
            tmp = p2 + nums[i]
        }
        p2 = p1
        p1 = tmp
    }
    if p1 > mx {
        mx = p1
    }
    return mx
}

func main() {
    fmt.Printf("rob([]int{ 2,3,2 }) = %v\n",rob([]int{ 2,3,2 })) // 3
    fmt.Printf("rob([]int{ 1,2,3,1 }) = %v\n",rob([]int{ 1,2,3,1 })) // 4  1 + 3
    fmt.Printf("rob([]int{ 1,2,3 }) = %v\n",rob([]int{ 1,2,3 })) // 3

    fmt.Printf("rob1([]int{ 2,3,2 }) = %v\n",rob1([]int{ 2,3,2 })) // 3
    fmt.Printf("rob1([]int{ 1,2,3,1 }) = %v\n",rob1([]int{ 1,2,3,1 })) // 4
    fmt.Printf("rob1([]int{ 1,2,3 }) = %v\n",rob1([]int{ 1,2,3 })) // 3
}