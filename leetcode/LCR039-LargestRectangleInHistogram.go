package main

// LCR 039. 柱状图中最大的矩形
// 给定非负整数数组 heights ，数组中的数字用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

// 示例 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/04/histogram.jpg" />
// 输入：heights = [2,1,5,6,2,3]
// 输出：10
// 解释：最大的矩形为图中红色区域，面积为 10

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2021/01/04/histogram-1.jpg" />
// 输入： heights = [2,4]
// 输出： 4

// 提示：
//     1 <= heights.length <=10^5
//     0 <= heights[i] <= 10^4

import "fmt"

// stack
func largestRectangleArea(heights []int) int {
    res, n := 0, len(heights) + 2
    // Add a sentry at the beginning and the end
    getHeight := func(i int) int { if i == 0 || n-1 == i { return 0; }; return heights[i-1]; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    stack := make([]int, 0, n/2)
    for i := 0; i < n; i++ {
        for len(stack) > 0 && getHeight(stack[len(stack) - 1]) > getHeight(i) {
            // pop stack
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res = max(res, getHeight(idx)*(i- stack[len(stack) -1 ] - 1))
        }
        // push stack
        stack = append(stack, i)
    }
    return res
}

// dp
func largestRectangleArea1(heights []int) int {
    l := len(heights)
    if l == 0 {
        return 0
    }
    left, right:= make([]int, l), make([]int, l)
    for i := 1 ; i < l; i++ {
        left[i] = i
        j := i
        for left[j] - 1 >= 0 && heights[left[j] - 1] >= heights[i] {
            j = left[j] - 1
        }
        left[i] = left[j]
    }
    right[l - 1] = l - 1
    for i := l - 2; i >= 0; i-- {
        right[i] = i
        j := i
        for right[j] + 1 <= l - 1 && heights[right[j] + 1] >= heights[i] {
            j = right[j] + 1
        }
        right[i] = right[j]
    }
    res := 0
    for i := 0; i < l; i++  {
        if heights[i] * (right[i] - left[i] + 1) > res {
            res = heights[i] * (right[i] - left[i] + 1)
        }
    }
    return res
}

func main() {
    // Explanation: The above is a histogram where width of each bar is 1.
    // The largest rectangle is shown in the red area, which has an area = 10 units.
    fmt.Printf("largestRectangleArea([]int{2,1,5,6,2,3}) = %v\n",largestRectangleArea([]int{2,1,5,6,2,3})) // 10
    fmt.Printf("largestRectangleArea([]int{2,4}) = %v\n",largestRectangleArea([]int{2,4})) // 4
    fmt.Printf("largestRectangleArea([]int{2,3}) = %v\n",largestRectangleArea([]int{2,3})) // 4

    fmt.Printf("largestRectangleArea1([]int{2,1,5,6,2,3}) = %v\n",largestRectangleArea1([]int{2,1,5,6,2,3})) // 10
    fmt.Printf("largestRectangleArea1([]int{2,4}) = %v\n",largestRectangleArea1([]int{2,4})) // 4
    fmt.Printf("largestRectangleArea1([]int{2,3}) = %v\n",largestRectangleArea1([]int{2,3})) // 4
}
