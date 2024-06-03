package main

// LCR 012. 寻找数组的中心下标
// 给你一个整数数组 nums ，请计算数组的 中心下标 。
// 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
// 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
// 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。

// 示例 1：
// 输入：nums = [1,7,3,6,5,6]
// 输出：3
// 解释：
// 中心下标是 3 。
// 左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
// 右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。

// 示例 2：
// 输入：nums = [1, 2, 3]
// 输出：-1
// 解释：
// 数组中不存在满足此条件的中心下标。

// 示例 3：
// 输入：nums = [2, 1, -1]
// 输出：0
// 解释：
// 中心下标是 0 。
// 左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
// 右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。

// 提示：
//     1 <= nums.length <= 10^4
//     -1000 <= nums[i] <= 1000

import "fmt"

func pivotIndex(nums []int) int {
    rSum, lSum := 0, 0
    for _, v := range nums[1:] { // <--- 累加所有值(除 0 的位置)
        rSum += v
    }
    if rSum == 0 {
        return 0
    }
    for i := 1; i < len(nums); i++ {
        lSum += nums[i-1]
        rSum -= nums[i]
        if lSum == rSum {
            return i
        }
    }
    return -1
}

func pivotIndex1(nums []int) int {
    right, left := 0, 0
    for _, num := range nums {
        right += num
    }   
    for i := 0; i < len(nums); i++ {
        right -= nums[i]
        if left == right {
            return i
        }
        left += nums[i]
    }
    return -1
}

func pivotIndex2(nums []int) int {
    leftPrefixSum, rightPrefixSum := make([]int, len(nums)), make([]int, len(nums))
    leftPrefixSum[0] = nums[0]
    rightPrefixSum[len(nums) - 1] = nums[len(nums) - 1]
    for i := 1; i < len(nums); i++ {
        leftPrefixSum[i] = nums[i] + leftPrefixSum[i - 1]
        endIdx := len(nums) - i - 1 
        rightPrefixSum[endIdx] = nums[endIdx] + rightPrefixSum[endIdx + 1]
    }
    for i := 0; i < len(nums); i++ {
        if leftPrefixSum[i] - nums[i] == rightPrefixSum[i] - nums[i] {
            return i
        }
    }
    return -1
}

func pivotIndex3(nums []int) int {
    sum, cur := 0, 0
    for _, x := range nums {
        sum += x
    }
    for i, x := range nums {
        if 2 * cur == sum - x {
            return i
        }
        cur += x
    }
    return -1
}

func main() {
    // Example 1:
    // Input: nums = [1,7,3,6,5,6]
    // Output: 3
    // Explanation:
    // The pivot index is 3.
    // Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
    // Right sum = nums[4] + nums[5] = 5 + 6 = 11
    fmt.Println(pivotIndex([]int{1,7,3,6,5,6})) // 3
    // Example 2:
    // Input: nums = [1,2,3]
    // Output: -1
    // Explanation:
    // There is no index that satisfies the conditions in the problem statement.
    fmt.Println(pivotIndex([]int{1,2,3})) // -1
    // Example 3:
    // Input: nums = [2,1,-1]
    // Output: 0
    // Explanation:
    // The pivot index is 0.
    // Left sum = 0 (no elements to the left of index 0)
    // Right sum = nums[1] + nums[2] = 1 + -1 = 0
    fmt.Println(pivotIndex([]int{2,1,-1})) // 0

    fmt.Println(pivotIndex1([]int{1,7,3,6,5,6})) // 3
    fmt.Println(pivotIndex1([]int{1,2,3})) // -1
    fmt.Println(pivotIndex1([]int{2,1,-1})) // 0

    fmt.Println(pivotIndex2([]int{1,7,3,6,5,6})) // 3
    fmt.Println(pivotIndex2([]int{1,2,3})) // -1
    fmt.Println(pivotIndex2([]int{2,1,-1})) // 0

    fmt.Println(pivotIndex3([]int{1,7,3,6,5,6})) // 3
    fmt.Println(pivotIndex3([]int{1,2,3})) // -1
    fmt.Println(pivotIndex3([]int{2,1,-1})) // 0
}