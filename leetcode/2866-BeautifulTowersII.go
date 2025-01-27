package main

// 2866. Beautiful Towers II
// You are given a 0-indexed array maxHeights of n integers.

// You are tasked with building n towers in the coordinate line. 
// The ith tower is built at coordinate i and has a height of heights[i].

// A configuration of towers is beautiful if the following conditions hold:
//     1. 1 <= heights[i] <= maxHeights[i]
//     2. heights is a mountain array.

// Array heights is a mountain if there exists an index i such that:
//     1. For all 0 < j <= i, heights[j - 1] <= heights[j]
//     2. For all i <= k < n - 1, heights[k + 1] <= heights[k]

// Return the maximum possible sum of heights of a beautiful configuration of towers.

// Example 1:
// Input: maxHeights = [5,3,4,1,1]
// Output: 13
// Explanation: One beautiful configuration with a maximum sum is heights = [5,3,3,1,1]. This configuration is beautiful since:
// - 1 <= heights[i] <= maxHeights[i]  
// - heights is a mountain of peak i = 0.
// It can be shown that there exists no other beautiful configuration with a sum of heights greater than 13.

// Example 2:
// Input: maxHeights = [6,5,3,9,2,7]
// Output: 22
// Explanation: One beautiful configuration with a maximum sum is heights = [3,3,3,9,2,2]. This configuration is beautiful since:
// - 1 <= heights[i] <= maxHeights[i]
// - heights is a mountain of peak i = 3.
// It can be shown that there exists no other beautiful configuration with a sum of heights greater than 22.

// Example 3:
// Input: maxHeights = [3,2,5,5,2,3]
// Output: 18
// Explanation: One beautiful configuration with a maximum sum is heights = [2,2,5,5,2,2]. This configuration is beautiful since:
// - 1 <= heights[i] <= maxHeights[i]
// - heights is a mountain of peak i = 2. 
// Note that, for this configuration, i = 3 can also be considered a peak.
// It can be shown that there exists no other beautiful configuration with a sum of heights greater than 18.

// Constraints:
//     1 <= n == maxHeights.length <= 10^5
//     1 <= maxHeights[i] <= 10^9

import "fmt"

func maximumSumOfHeights(heights []int) int64 {
    res, n := int64(0), len(heights)
    left, right := make([]int, n), make([]int, n)
    f, g := make([]int64, n), make([]int64, n)
    // left stack 
    stack1 := []int{}
    for i, height := range heights {
        for len(stack1) > 0 && heights[stack1[len(stack1) - 1]] > height {
            stack1 = stack1[:len(stack1) - 1]
        }
        if len(stack1) > 0 {
            left[i] = stack1[len(stack1) - 1]
        } else {
            left[i] = -1
        }
        stack1 = append(stack1, i)
    }
    // right stack
    stack2 := []int{}
    for i := n - 1; i >= 0; i-- {
        height := heights[i]
        for len(stack2) > 0 && heights[stack2[len(stack2) - 1]] >= height {
            stack2 = stack2[:len(stack2) - 1]
        }
        if len(stack2) > 0 {
            right[i] = stack2[len(stack2) - 1]
        } else {
            right[i] = n
        }
        stack2 = append(stack2, i)
    }
    // f array
    for i, height := range heights {
        if i > 0 && height >= heights[i - 1] {
            f[i] = f[i - 1] + int64(height)
        } else {
            j := left[i]
            if j != -1 {
                f[i] = int64(height) * int64(i - j) + f[j]
            } else {
                f[i] = int64(height) * int64(i + 1)
            }
        }
    }
    // g array
    for i := n - 1; i >= 0; i-- {
        height := heights[i]
        if i < n - 1 && height >= heights[i + 1] {
            g[i] = g[i + 1] + int64(height)
        } else {
            j := right[i]
            if j != n {
                g[i] = int64(height) * int64(j - i) + g[j]
            } else {
                g[i] = int64(height) * int64(n - i)
            }
        }
    }
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    for i, height := range heights {
        res = max(res, f[i] + g[i] - int64(height))
    }
    return res
}

func maximumSumOfHeights1(heights []int) int64 {
    res, s, n := 0, 0, len(heights)
    stack, suffix := []int{ n }, make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        h := heights[i]
        for len(stack) > 1 && heights[stack[len(stack) - 1]] >= h {
            stack = stack[:len(stack) - 1]
        }
        j := stack[len(stack) - 1]
        suffix[i] = suffix[j] + h * (j - i)
        stack = append(stack, i)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    stack = []int{ -1 }
    for j, h := range heights {
        res = max(res, s + suffix[j])
        for len(stack) > 1 && heights[stack[len(stack) - 1]] >= h {
            i := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]
            s -= (i - stack[len(stack) - 1]) * heights[i]
        }
        i := stack[len(stack) - 1]
        s += h * (j-i)
        stack = append(stack, j)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: maxHeights = [5,3,4,1,1]
    // Output: 13
    // Explanation: One beautiful configuration with a maximum sum is heights = [5,3,3,1,1]. This configuration is beautiful since:
    // - 1 <= heights[i] <= maxHeights[i]  
    // - heights is a mountain of peak i = 0.
    // It can be shown that there exists no other beautiful configuration with a sum of heights greater than 13.
    fmt.Println(maximumSumOfHeights([]int{5,3,4,1,1})) // 13
    // Example 2:
    // Input: maxHeights = [6,5,3,9,2,7]
    // Output: 22
    // Explanation: One beautiful configuration with a maximum sum is heights = [3,3,3,9,2,2]. This configuration is beautiful since:
    // - 1 <= heights[i] <= maxHeights[i]
    // - heights is a mountain of peak i = 3.
    // It can be shown that there exists no other beautiful configuration with a sum of heights greater than 22.
    fmt.Println(maximumSumOfHeights([]int{6,5,3,9,2,7})) // 22
    // Example 3:
    // Input: maxHeights = [3,2,5,5,2,3]
    // Output: 18
    // Explanation: One beautiful configuration with a maximum sum is heights = [2,2,5,5,2,2]. This configuration is beautiful since:
    // - 1 <= heights[i] <= maxHeights[i]
    // - heights is a mountain of peak i = 2. 
    // Note that, for this configuration, i = 3 can also be considered a peak.
    // It can be shown that there exists no other beautiful configuration with a sum of heights greater than 18.
    fmt.Println(maximumSumOfHeights([]int{3,2,5,5,2,3})) // 18

    fmt.Println(maximumSumOfHeights([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maximumSumOfHeights([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maximumSumOfHeights1([]int{5,3,4,1,1})) // 13
    fmt.Println(maximumSumOfHeights1([]int{6,5,3,9,2,7})) // 22
    fmt.Println(maximumSumOfHeights1([]int{3,2,5,5,2,3})) // 18
    fmt.Println(maximumSumOfHeights1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maximumSumOfHeights1([]int{9,8,7,6,5,4,3,2,1})) // 45
}