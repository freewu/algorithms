package main

// LCR 089. 打家劫舍
// 一个专业的小偷，计划偷窃沿街的房屋。
// 每间房内都藏有一定的现金，影响小偷偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

// 给定一个代表每个房屋存放金额的非负整数数组 nums ，请计算 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

// 示例 1：
// 输入：nums = [1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//      偷窃到的最高金额 = 1 + 3 = 4 。

// 示例 2：
// 输入：nums = [2,7,9,3,1]
// 输出：12
// 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//      偷窃到的最高金额 = 2 + 9 + 1 = 12 。

// 提示：
//     1 <= nums.length <= 100
//     0 <= nums[i] <= 400

import "fmt"

// dp
func rob(nums []int) int {
    l := len(nums)
    if l == 0 {
        return 0
    }
    if l == 1 {
        return nums[0]
    }
    // dp[i] 代表抢 nums[0...i] 房子的最大价值
    dp := make([]int, l)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp[0], dp[1] = nums[0], max(nums[1], nums[0])
    for i := 2; i < l; i++ {
        dp[i] = max(dp[i-1], nums[i] + dp[i-2])
    }
    return dp[l-1]
}

// dp 优化辅助空间，把迭代的值保存在 2 个变量中
func rob1(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    res, pre := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(nums); i++ {
        tmp := res
        res = max(res, nums[i] + pre)
        pre = tmp
    }
    return res
}

// 模拟
func rob2(nums []int) int {
    // a 对于偶数位上的最大值的记录
    // b 对于奇数位上的最大值的记录
    a, b := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(nums); i++ {
        if i % 2 == 0 {
            a = max(a + nums[i], b)
        } else {
            b = max(a, b + nums[i])
        }
    }
    return max(a, b)
}

func main() {
    fmt.Printf("rob([]int{ 1,2,3,1 }) = %v\n",rob([]int{ 1,2,3,1 })) // 4   1 + 3
    fmt.Printf("rob([]int{ 2,7,9,3,1 }) = %v\n",rob([]int{ 2,7,9,3,1 })) // 12  2 + 9 + 1 = 12.
    fmt.Printf("rob([]int{ 2,7,9,3,1,10,3,1 }) = %v\n",rob([]int{ 2,7,9,3,1,10,3,1 })) // 22  2 + 9 + 10 + １

    fmt.Printf("rob1([]int{ 1,2,3,1 }) = %v\n",rob1([]int{ 1,2,3,1 })) // 4   1 + 3
    fmt.Printf("rob1([]int{ 2,7,9,3,1 }) = %v\n",rob1([]int{ 2,7,9,3,1 })) // 12  2 + 9 + 1 = 12.
    fmt.Printf("rob([]int{ 2,7,9,3,1,10,3,1 }) = %v\n",rob1([]int{ 2,7,9,3,1,10,3,1 })) // 22  2 + 9 + 10 + １

    fmt.Printf("rob2([]int{ 1,2,3,1 }) = %v\n",rob2([]int{ 1,2,3,1 })) // 4   1 + 3
    fmt.Printf("rob3([]int{ 2,7,9,3,1 }) = %v\n",rob2([]int{ 2,7,9,3,1 })) // 12  2 + 9 + 1 = 12.
    fmt.Printf("rob([]int{ 2,7,9,3,1,10,3,1 }) = %v\n",rob2([]int{ 2,7,9,3,1,10,3,1 })) // 22  2 + 9 + 10 + １
}
