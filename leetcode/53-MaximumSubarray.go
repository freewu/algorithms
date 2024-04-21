package main

// 53. Maximum Subarray
// Given an integer array nums, find the subarray with the largest sum, and return its sum.

// Example 1:
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum 6.

// Example 2:
// Input: nums = [1]
// Output: 1
// Explanation: The subarray [1] has the largest sum 1.

// Example 3:
// Input: nums = [5,4,-1,7,8]
// Output: 23
// Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
 
// Constraints:
//     1 <= nums.length <= 10^5
//     -10^4 <= nums[i] <= 10^4
 
// Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

import "fmt"

func maxSubArray(nums []int) int {
    mx, sum := -1 >> 32 - 1, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range nums {
        sum = max(sum + nums[i], nums[i]) // 如果累加值都小于当前值，从当前值开始
        mx = max(mx, sum)
    }
    return mx
}

func maxSubArray1(nums []int) int {
    var l = len(nums)
    var max = 0
    array := make([][]int, l)

    for i := 0; i < l; i++ {
        subArray := make([]int, l)
        subArray[0] = nums[i]
        for j := i + 1; j < l; j++ {
            subArray[i] += nums[j]
            if subArray[i] > max {
                max = subArray[i]
            }
        }
        array[i] = subArray
    }

    // // 输出
    // for i := range array {
    //     for j := range array[i] {
    //         fmt.Printf("%v ", array[i][j])
    //     }
    //     fmt.Println()
    // }
    return max
}

// DP
func maxSubArray2(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    max := func (a int, b int) int { if a > b { return a; }; return b; }
    dp, res := make([]int, len(nums)), nums[0]
    dp[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        if dp[i-1] > 0 { // 如果值大于 0 则累加
            dp[i] = nums[i] + dp[i-1]
        } else {
            dp[i] = nums[i]
        }
        res = max(res, dp[i])
    }
    return res
}

// 模拟
func maxSubArray3(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    maxSum, res, p := nums[0], 0, 0
    for p < len(nums) {
        res += nums[p]
        if res > maxSum {
            maxSum = res
        }
        if res < 0 {
            res = 0
        }
        p++
    }
    return maxSum
}

// best solution 快慢指针
func maxSubArray4(nums []int) int {
    length := len(nums)
    fast, slow := 0,0
    result, sum :=  -1 >> 32 - 1, 0
    for fast < length && slow < length {
        sum = sum + nums[fast]
        if sum > result {
            result = sum
        }
        for sum< 0 {
            sum = sum - nums[slow]
            slow++
        }
        fast++
    }
    return result
}

func main() {
    fmt.Printf("maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSubArray([]int{1}) = %v\n",maxSubArray([]int{1})) // 1
    fmt.Printf("maxSubArray([]int{5,4,-1,7,8}) = %v\n",maxSubArray([]int{5,4,-1,7,8})) // 23

    fmt.Printf("maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSubArray1([]int{1}) = %v\n",maxSubArray1([]int{1})) // 1
    fmt.Printf("maxSubArray1([]int{5,4,-1,7,8}) = %v\n",maxSubArray1([]int{5,4,-1,7,8})) // 23

    fmt.Printf("maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSubArray2([]int{1}) = %v\n",maxSubArray2([]int{1})) // 1
    fmt.Printf("maxSubArray2([]int{5,4,-1,7,8}) = %v\n",maxSubArray2([]int{5,4,-1,7,8})) // 23

    fmt.Printf("maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSubArray3([]int{1}) = %v\n",maxSubArray3([]int{1})) // 1
    fmt.Printf("maxSubArray3([]int{5,4,-1,7,8}) = %v\n",maxSubArray3([]int{5,4,-1,7,8})) // 23

    fmt.Printf("maxSubArray4([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) = %v\n",maxSubArray4([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
    fmt.Printf("maxSubArray4([]int{1}) = %v\n",maxSubArray4([]int{1})) // 1
    fmt.Printf("maxSubArray4([]int{5,4,-1,7,8}) = %v\n",maxSubArray4([]int{5,4,-1,7,8})) // 23
}
