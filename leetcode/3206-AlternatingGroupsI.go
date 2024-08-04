package main

// 3206. Alternating Groups I
// There is a circle of red and blue tiles. You are given an array of integers colors. 
// The color of tile i is represented by colors[i]:
//     colors[i] == 0 means that tile i is red.
//     colors[i] == 1 means that tile i is blue.

// Every 3 contiguous tiles in the circle with alternating colors (the middle tile has a different color from its left and right tiles) is called an alternating group.
// Return the number of alternating groups.
// Note that since colors represents a circle, the first and the last tiles are considered to be next to each other.

// Example 1:
// Input: colors = [1,1,1]
// Output: 0
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/16/image_2024-05-16_23-53-171.png" />

// Example 2:
// Input: colors = [0,1,0,0,1]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/16/image_2024-05-16_23-47-491.png" />
// Alternating groups:
// <img src="https://assets.leetcode.com/uploads/2024/05/16/image_2024-05-16_23-48-211.png" />

// Constraints:
//     3 <= colors.length <= 100
//     0 <= colors[i] <= 1

import "fmt"

func numberOfAlternatingGroups(colors []int) int {
    colors = append(colors, colors[0])
    colors = append(colors, colors[1])
    res := 0
    for i := 0; i<=len(colors) - 3; i++ {
        if colors[i] == 0 && colors[i+1] == 1 && colors[i+2] == 0 {
            res++ 
        }
        if colors[i] == 1 && colors[i+1] == 0 && colors[i+2] == 1 {
            res++ 
        }
    }
    return res
}

func numberOfAlternatingGroups1(colors []int) int {
    res, count, last, n := 0, 0, -1 ,len(colors)
    for i := 0; i <= n + 1; i++ {
        k := i % n
        if colors[k] == last {
            count = 1
            last = colors[k]
            continue
        }
        last = colors[k]
        count++
        if count >= 3 {
            res++
        }
    } 
    return res
}

func main() {
    // Example 1:
    // Input: colors = [1,1,1]
    // Output: 0
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/16/image_2024-05-16_23-53-171.png" />
    fmt.Println(numberOfAlternatingGroups([]int{1,1,1})) // 0
    // Example 2:
    // Input: colors = [0,1,0,0,1]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/16/image_2024-05-16_23-47-491.png" />
    // Alternating groups:
    // <img src="https://assets.leetcode.com/uploads/2024/05/16/image_2024-05-16_23-48-211.png" />
    fmt.Println(numberOfAlternatingGroups([]int{0,1,0,0,1})) // 3

    fmt.Println(numberOfAlternatingGroups1([]int{1,1,1})) // 0
    fmt.Println(numberOfAlternatingGroups1([]int{0,1,0,0,1})) // 3
}