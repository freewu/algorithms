package main

// 面试题 16.17. Contiguous Sequence LCCI
// You are given an array of integers (both positive and negative). 
// Find the contiguous sequence with the largest sum. 
// Return the sum.

// Example:
// Input:  [-2,1,-3,4,-1,2,1,-5,4]
// Output:  6
// Explanation:  [4,-1,2,1] has the largest sum 6.
// Follow Up:
// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

import "fmt"

// 动态规划
func maxSubArray(nums []int) int {
    res := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] + nums[i-1] > nums[i] {
            nums[i] += nums[i-1]
        }
        if nums[i] > res {
            res = nums[i]
        }
    }
    return res
}

// 分治
func maxSubArray1(nums []int) int {
    type Status struct { lSum, rSum, mSum, iSum int }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    pushUp := func(l, r Status) Status {
        return Status{ max(l.lSum, l.iSum + r.lSum), max(r.rSum, r.iSum + l.rSum), max(max(l.mSum, r.mSum), l.rSum + r.lSum), l.iSum + r.iSum }
    }
    var get func(nums []int, l, r int) Status
    get = func(nums []int, l, r int) Status {
        if l == r {
            return Status{ nums[l], nums[l], nums[l], nums[l] }
        }
        mid := (l + r) >> 1
        return pushUp(get(nums, l, mid), get(nums, mid + 1, r))
    }
    return get(nums, 0, len(nums) - 1).mSum
}

func main() {
    // Example:
    // Input:  [-2,1,-3,4,-1,2,1,-5,4]
    // Output:  6
    // Explanation:  [4,-1,2,1] has the largest sum 6.
    // Follow Up:
    // If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
    fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})) // 6

    fmt.Println(maxSubArray([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSubArray([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maxSubArray1([]int{-2,1,-3,4,-1,2,1,-5,4})) // 6
    fmt.Println(maxSubArray1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSubArray1([]int{9,8,7,6,5,4,3,2,1})) // 45
}