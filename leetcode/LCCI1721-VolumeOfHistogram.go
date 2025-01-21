package main

// 面试题 17.21. Volume of Histogram LCCI

// Imagine a histogram (bar graph). 
// Design an algorithm to compute the volume of water it could hold if someone poured water across the top. 
// You can assume that each histogram bar has width 1.

// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png" />

// The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. 
// In this case, 6 units of water (blue section) are being trapped. Thanks Marcos for contributing this image!

// Example:
// Input: [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6

import "fmt"

// 动态规划
func trap(height []int) int {
    res, n := 0, len(height)
    if n == 0 { return res }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    left := make([]int, n)
    left[0] = height[0]
    for i := 1; i < n; i++ {
        left[i] = max(left[i - 1], height[i])
    }
    right := make([]int, n)
    right[n-1] = height[n-1]
    for i := n - 2; i >= 0; i-- {
        right[i] = max(right[i + 1], height[i])
    }
    for i, h := range height {
        res += min(left[i], right[i]) - h
    }
    return res
}

// 单调栈
func trap1(height []int) int {
    res, stack := 0, []int{}
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, h := range height {
        for len(stack) > 0 && h > height[stack[len(stack) - 1]] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 { break }
            left := stack[len(stack)-1]
            res += ((i - left - 1) * (min(height[left], h) - height[top]))  // width * height
        }
        stack = append(stack, i)
    }
    return res
}

// 双指针
func trap2(height []int) (ans int) {
    res, left, right, leftMax, rightMax := 0, 0, len(height) - 1, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for left < right {
        leftMax, rightMax = max(leftMax, height[left]), max(rightMax, height[right])
        if height[left] < height[right] {
            res += leftMax - height[left]
            left++
        } else {
            res += rightMax - height[right]
            right--
        }
    }
    return res
}

func main() {
    // Example:
    // Input: [0,1,0,2,1,0,1,3,2,1,2,1]
    // Output: 6
    fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})) // 6

    fmt.Println(trap1([]int{0,1,0,2,1,0,1,3,2,1,2,1})) // 6

    fmt.Println(trap2([]int{0,1,0,2,1,0,1,3,2,1,2,1})) // 6
}