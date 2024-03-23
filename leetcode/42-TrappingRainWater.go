package main

// 42. Trapping Rain Water
// Given n non-negative integers representing an elevation map where the width of each bar is 1, 
// compute how much water it can trap after raining.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png" />
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: 
//     The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. 
//     In this case, 6 units of rain water (blue section) are being trapped.

// Example 2:
// Input: height = [4,2,0,3,2,5]
// Output: 9

// Constraints:
//     n == height.length
//     1 <= n <= 2 * 10^4
//     0 <= height[i] <= 10^5

// 解题思路:
//     从 x 轴开始，给出一个数组，数组里面的数字代表从 (0,0) 点开始，宽度为 1 个单位，高度为数组元素的值。
//     如果下雨了，问这样一个容器能装多少单位的水

import "fmt"

func trap(height []int) int {
    res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
    for left <= right { // 从开头 0 和结尾 len(height) - 1 向中间靠拢
        if height[left] <= height[right] { // 如果 左边(开头) <= 右边(结尾) 开头向里走一步
            if height[left] > maxLeft { // 如果 左边 最大
                maxLeft = height[left]
            } else {
                res += maxLeft - height[left] // 这里是关键
            }
            left++
        } else {
            if height[right] >= maxRight {
                maxRight = height[right]
            } else {
                res += maxRight - height[right] //
            }
            right--
        }
    }
    return res
}

// best solution
func trap1(heights []int) int {
    computeLeftBoundary := func (heights []int) []int {
        res := make([]int, len(heights))
        for i := 1; i < len(heights); i++ {
            res[i] = max(res[i-1], heights[i-1])
        }
        return res
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }

    leftBoundaries := computeLeftBoundary(heights)
    right := heights[len(heights)-1]
    res := 0
    for i := len(heights) - 2; i >= 1; i-- {
        ht := min(leftBoundaries[i], right)
        if ht > heights[i] {
            res += ht - heights[i]
        }
        right = max(right, heights[i])
    }
    return res
}

func trap2(height []int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }

    n := len(height)
    lmax := make([]int, n)
    lmax[0] = height[0]
    for i := 1; i < n; i++ {
        lmax[i] = max(lmax[i-1], height[i])
    }
    rmax := make([]int, n)
    rmax[n-1] = height[n-1]
    for i := n-2; i >= 0; i-- {
        rmax[i] = max(rmax[i+1], height[i])
    }
    res := 0
    for i := 0; i < n; i++ {
        res += (min(lmax[i], rmax[i]) - height[i])
    }
    return res
}


func main() {
    fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})) // 6
    fmt.Println(trap([]int{4,2,0,3,2,5})) // 9

    fmt.Println(trap1([]int{0,1,0,2,1,0,1,3,2,1,2,1})) // 6
    fmt.Println(trap1([]int{4,2,0,3,2,5})) // 9
    
    fmt.Println(trap2([]int{0,1,0,2,1,0,1,3,2,1,2,1})) // 6
    fmt.Println(trap2([]int{4,2,0,3,2,5})) // 9
}
