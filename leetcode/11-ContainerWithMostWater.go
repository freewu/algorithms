package main

// 11. Container With Most Water  Medium
// You are given an integer array height of length n. 
// There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

// Example 1:
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
// In this case, the max area of water (blue section) the container can contain is 49.

// Example 2:
// Input: height = [1,1]
// Output: 1

// Constraints:
//     n == height.length
//     2 <= n <= 10^5
//     0 <= height[i] <= 10^4

// # 解题思路
// 	对撞指针的思路。首尾分别 2 个指针，每次移动以后都分别判断长宽的乘积是否最大

import "fmt"

func maxArea(height []int) int {
    max, start, end := 0, 0, len(height)-1
    // 从头尾向中间靠拢  [ start -> ... <- end ]
    for start < end {
        width := end - start // 宽度
        high := 0
        if height[start] < height[end] {// 如果 开头的高度 < 结尾的高度  从头部向 中间走一步
            high = height[start]
            start++
        } else { // 如果 开头的高度 >= 结尾的高度  从尾部向 中间走一步
            high = height[end]
            end--
        }
        temp := width * high // 计算出面积
        if temp > max { // 保存最大的面试,如果需要得到 起始位置 加两个变量保存位置即可
            max = temp
        }
    }
    return max
}

// best solution
func maxArea1(height []int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    area, res := 0, 0
    low, high := 0, len(height) - 1

    for low < high { // 从头尾向中间靠拢  [ start -> ... <- end ]
        area = (high - low) * min(height[high], height[low]) // 计算面积
        if area > res {
            res = area
        }
        if height[high] > height[low] { // 如果 结尾的高度 > 开头的高度 从头部向 中间走一步
            low++
        } else {  // 如果 结尾的高度 <= 开头的高度 从头部向 中间走一步
            high--
        }
    }
    return res
}

func maxArea2(height []int) int {
    l, r, res := 0, len(height) - 1, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for l < r {
        area :=  (r - l) * min(height[l], height[r])
        res = max(res, area)
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: height = [1,8,6,2,5,4,8,3,7]
    // Output: 49
    // Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
    // In this case, the max area of water (blue section) the container can contain is 49.
    fmt.Printf("maxArea([]int{1,8,6,2,5,4,8,3,7} = %v \n",maxArea([]int{1,8,6,2,5,4,8,3,7})); // 49
    // Example 2:
    // Input: height = [1,1]
    // Output: 1
    fmt.Printf("maxArea([]int{1,1} = %v \n",maxArea([]int{1,1})); // 1

    fmt.Printf("maxArea([]int{1,2,3,4,5,6,7,8,9} = %v \n",maxArea([]int{1,2,3,4,5,6,7,8,9})); // 20
    fmt.Printf("maxArea([]int{9,8,7,6,5,4,3,2,1} = %v \n",maxArea([]int{9,8,7,6,5,4,3,2,1})); // 20

    fmt.Printf("maxArea1([]int{1,8,6,2,5,4,8,3,7} = %v \n",maxArea1([]int{1,8,6,2,5,4,8,3,7})); // 49
    fmt.Printf("maxArea1([]int{1,1} = %v \n",maxArea1([]int{1,1})); // 1
    fmt.Printf("maxArea1([]int{1,2,3,4,5,6,7,8,9} = %v \n",maxArea1([]int{1,2,3,4,5,6,7,8,9})); // 20
    fmt.Printf("maxArea1([]int{9,8,7,6,5,4,3,2,1} = %v \n",maxArea1([]int{9,8,7,6,5,4,3,2,1})); // 20

    fmt.Printf("maxArea2([]int{1,8,6,2,5,4,8,3,7} = %v \n",maxArea2([]int{1,8,6,2,5,4,8,3,7})); // 49
    fmt.Printf("maxArea2([]int{1,1} = %v \n",maxArea2([]int{1,1})); // 1
    fmt.Printf("maxArea2([]int{1,2,3,4,5,6,7,8,9} = %v \n",maxArea2([]int{1,2,3,4,5,6,7,8,9})); // 20
    fmt.Printf("maxArea2([]int{9,8,7,6,5,4,3,2,1} = %v \n",maxArea2([]int{9,8,7,6,5,4,3,2,1})); // 20
}