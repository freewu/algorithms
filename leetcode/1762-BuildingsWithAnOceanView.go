package main

// 1762. Buildings With an Ocean View
// There are n buildings in a line. 
// You are given an integer array heights of size n that represents the heights of the buildings in the line.

// The ocean is to the right of the buildings. 
// A building has an ocean view if the building can see the ocean without obstructions. 
// Formally, a building has an ocean view if all the buildings to its right have a smaller height.

// Return a list of indices (0-indexed) of buildings that have an ocean view, sorted in increasing order.

// Example 1:
// Input: heights = [4,2,3,1]
// Output: [0,2,3]
// Explanation: Building 1 (0-indexed) does not have an ocean view because building 2 is taller.

// Example 2:
// Input: heights = [4,3,2,1]
// Output: [0,1,2,3]
// Explanation: All the buildings have an ocean view.

// Example 3:
// Input: heights = [1,3,2,4]
// Output: [3]
// Explanation: Only building 3 has an ocean view.

// Constraints:
//     1 <= heights.length <= 10^5
//     1 <= heights[i] <= 10^9

import "fmt"

// 单调栈
func findBuildings(heights []int) []int {
    stack := make([][]int,0)
    for i := 0; i < len(heights); i++ {
        for len(stack) > 0 && heights[i] >= stack[len(stack)-1][0] { // 可以挡住后面的视野，出栈
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, []int{heights[i], i}) // 高度，位置
    }
    res := make([]int, 0)
    for i := 0; i < len(stack); i++ {
        res = append(res,stack[i][1])
    }
    return res
}

func findBuildings1(heights []int) []int {
    if len(heights) == 0 {
        return nil
    }
    pos, high := []int{}, []int{} // 位置, 高度
    for i, v := range heights {
        if len(high) == 0 {
            high, pos = append(high, v), append(pos, i)
            continue
        }
        for len(high) != 0 && v >= high[len(high)-1] { // 挡住了
            high, pos = high[:len(high)-1], pos[:len(pos)-1]
        }
        high, pos = append(high, v), append(pos, i)
    }
    return pos
}

func main() {
    // Example 1:
    // Input: heights = [4,2,3,1]
    // Output: [0,2,3]
    // Explanation: Building 1 (0-indexed) does not have an ocean view because building 2 is taller.
    fmt.Println(findBuildings([]int{4,2,3,1})) // [0,2,3]
    // Example 2:
    // Input: heights = [4,3,2,1]
    // Output: [0,1,2,3]
    // Explanation: All the buildings have an ocean view.
    fmt.Println(findBuildings([]int{4,3,2,1})) // [0,1,2,3]
    // Example 3:
    // Input: heights = [1,3,2,4]
    // Output: [3]
    // Explanation: Only building 3 has an ocean view.
    fmt.Println(findBuildings([]int{1,3,2,4})) // [3]

    fmt.Println(findBuildings1([]int{4,2,3,1})) // [0,2,3]
    fmt.Println(findBuildings1([]int{4,3,2,1})) // [0,1,2,3]
    fmt.Println(findBuildings1([]int{1,3,2,4})) // [3]
}