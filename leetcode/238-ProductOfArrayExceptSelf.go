package main

// 238. Product of Array Except Self
// Given an integer array nums, 
// return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Example 2:
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:
//     2 <= nums.length <= 10^5
//     -30 <= nums[i] <= 30
//     The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
 
// Follow up: Can you solve the problem in O(1) extra space complexity? 
// (The output array does not count as extra space for space complexity analysis.)

import "fmt"

// 暴力法 遍历就对了 Time Limit Exceeded
func productExceptSelf(nums []int) []int {
    res := make([]int,len(nums))
    for i := 0; i < len(nums); i++ {
        sum := 1
        for j := 0; j < len(nums); j++ {
            // 不等于自己都相乘
            if j != i {
                sum = sum * nums[j]
            }
        }
        res[i] = sum
    }
    return res
}

// 计算整体乘积再除自己
func productExceptSelf1(nums []int) []int {
    res := make([]int,len(nums))
    // 求出所有的乘积
    t := 1
    zero := 0
    zero_index := 0
    for i := 0; i < len(nums); i++ {
        // 排除 0 的
        if nums[i] == 0 {
            zero++
            zero_index = i
            // 出现两 0 直接返回 全 0 的结果
            if zero >= 2 {
                return res
            }
            continue
        }
        t *= nums[i]
    }
    // 有一个 0 存在,把 0 的位置替换成 t 结束
    if zero == 1 {
        res[zero_index] = t 
        return res
    }
    // 没有 0 出现在的情况 
    for i := 0; i < len(nums); i++ {
        res[i] = t / nums[i]
    }
    return res
}

// 不用除法
func productExceptSelf2(nums []int) []int {
    left := make([]int,len(nums))
    // 从左到右累乘 0 --> len(nums) - 1
    for i := range nums {
        if i == 0 {
            left[i] = int(1)
            continue
        }
        left[i] = left[i-1] * nums[i-1]
    }
    // 从右到左累乘 0 --> len(nums) - 1
    right := make([]int,len(nums))
    for i := range nums {
        k := len(nums) - 1 - i
        if i == 0 {
            right[k] = int(1)
            continue
        }
        right[k] = right[k+1] * nums[k+1]
    }
    answer :=  make([]int,len(nums))
    // fmt.Println("left: ",left," right: ",right)
    // 左右结果相乘得出结果
    for i := range nums{
         answer[i] = left[i] * right[i]
    }
    return answer
}

// 不用除法 容易理解的
func productExceptSelf3(nums []int) []int {
    l := len(nums)
    left := make([]int,l)
    // 从左到右累乘 0 --> len(nums) - 1
    for i := 0; i < l; i++ {
        if i == 0 {
            left[i] = int(1)
            continue
        }
        left[i] = left[i - 1] * nums[i - 1]
    }
    // 从右到左累乘 0 --> len(nums) - 1
    right := make([]int,len(nums))
    for i := l - 1 ; i >= 0; i-- {
        if i == l - 1 {
            right[i] = int(1)
            continue
        }
        right[i] = right[i + 1] * nums[i + 1]
    }
    res :=  make([]int,len(nums))
    //fmt.Println("left: ",left," right: ",right)
    // 左右结果相乘得出结果
    for i := range nums{
        res[i] = left[i] * right[i]
    }
    return res
}

func main() {
    fmt.Println(productExceptSelf([]int{ 1,2,3,4 })) // [24,12,8,6]
    fmt.Println(productExceptSelf([]int{ -1,1,0,-3,3 })) // [0,0,9,0,0]

    fmt.Println(productExceptSelf1([]int{ 1,2,3,4 })) // [24,12,8,6]
    fmt.Println(productExceptSelf1([]int{ -1,1,0,-3,3 })) // [0,0,9,0,0]

    fmt.Println(productExceptSelf2([]int{ 1,2,3,4 })) // [24,12,8,6]
    fmt.Println(productExceptSelf2([]int{ -1,1,0,-3,3 })) // [0,0,9,0,0]

    fmt.Println(productExceptSelf3([]int{ 1,2,3,4 })) // [24,12,8,6]
    fmt.Println(productExceptSelf3([]int{ -1,1,0,-3,3 })) // [0,0,9,0,0]
}