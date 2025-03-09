package main

// 3208. Alternating Groups II
// There is a circle of red and blue tiles. 
// You are given an array of integers colors and an integer k. 
// The color of tile i is represented by colors[i]:
//     colors[i] == 0 means that tile i is red.
//     colors[i] == 1 means that tile i is blue.

// An alternating group is every k contiguous tiles in the circle with alternating colors 
// (each tile in the group except the first and last one has a different color from its left and right tiles).

// Return the number of alternating groups.
// Note that since colors represents a circle, the first and the last tiles are considered to be next to each other.

// Example 1:
// Input: colors = [0,1,0,1,0], k = 3
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/19/screenshot-2024-05-28-183519.png" />
// Alternating groups:
// <img src="https://assets.leetcode.com/uploads/2024/05/28/screenshot-2024-05-28-182844.png" />

// Example 2:
// Input: colors = [0,1,0,0,1,0,1], k = 6
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/19/screenshot-2024-05-28-183907.png" />
// Alternating groups:
// <img src="https://assets.leetcode.com/uploads/2024/06/19/screenshot-2024-05-28-184240.png" />

// Example 3:
// Input: colors = [1,1,0,1], k = 4
// Output: 0
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/19/screenshot-2024-05-28-184516.png" />

// Constraints:
//     3 <= colors.length <= 10^5
//     0 <= colors[i] <= 1
//     3 <= k <= colors.length

import "fmt"

func numberOfAlternatingGroups(colors []int, k int) int {
    colors = append(colors, colors[:k-1]...)
    count, l, res := 1, 1, 0
    for l < len(colors) {
        if colors[l] != colors[l-1] {
            count++
        } else {
            count = 1
        }
        if count >= k {
            res++
        }
        l++
    }
    return res
}

// 双指针
func numberOfAlternatingGroups1(colors []int, k int) int {
    res, l, r, n := 0, 0, 0, len(colors)
    for l < n {
        r++
        if colors[r % n] == colors[(r - 1) % n] {
            l = r
        } else if r - l + 1 == k {
            res++
            l++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: colors = [0,1,0,1,0], k = 3
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/19/screenshot-2024-05-28-183519.png" />
    // Alternating groups:
    // <img src="https://assets.leetcode.com/uploads/2024/05/28/screenshot-2024-05-28-182844.png" />
    fmt.Println(numberOfAlternatingGroups([]int{0,1,0,1,0}, 3)) // 3
    // Example 2:
    // Input: colors = [0,1,0,0,1,0,1], k = 6
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/19/screenshot-2024-05-28-183907.png" />
    // Alternating groups:
    // <img src="https://assets.leetcode.com/uploads/2024/06/19/screenshot-2024-05-28-184240.png" />
    fmt.Println(numberOfAlternatingGroups([]int{0,1,0,0,1,0,1}, 6)) // 2
    // Example 3:
    // Input: colors = [1,1,0,1], k = 4
    // Output: 0
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/19/screenshot-2024-05-28-184516.png" />
    fmt.Println(numberOfAlternatingGroups([]int{1,1,0,1}, 4)) // 0

    fmt.Println(numberOfAlternatingGroups([]int{1,1,1,1,1,1,1,1,1,1}, 4)) // 0
    fmt.Println(numberOfAlternatingGroups([]int{0,0,0,0,0,0,0,0,0,0}, 4)) // 0
    fmt.Println(numberOfAlternatingGroups([]int{1,1,1,1,1,0,0,0,0,0}, 4)) // 0
    fmt.Println(numberOfAlternatingGroups([]int{0,0,0,0,0,1,1,1,1,1}, 4)) // 0
    fmt.Println(numberOfAlternatingGroups([]int{0,1,0,1,0,1,0,1,0,1}, 4)) // 10
    fmt.Println(numberOfAlternatingGroups([]int{1,0,1,0,1,0,1,0,1,0}, 4)) // 10

    fmt.Println(numberOfAlternatingGroups1([]int{0,1,0,1,0}, 3)) // 3
    fmt.Println(numberOfAlternatingGroups1([]int{0,1,0,0,1,0,1}, 6)) // 2
    fmt.Println(numberOfAlternatingGroups1([]int{1,1,0,1}, 4)) // 0
    fmt.Println(numberOfAlternatingGroups1([]int{1,1,1,1,1,1,1,1,1,1}, 4)) // 0
    fmt.Println(numberOfAlternatingGroups1([]int{0,0,0,0,0,0,0,0,0,0}, 4)) // 0
    fmt.Println(numberOfAlternatingGroups1([]int{1,1,1,1,1,0,0,0,0,0}, 4)) // 0
    fmt.Println(numberOfAlternatingGroups1([]int{0,0,0,0,0,1,1,1,1,1}, 4)) // 0
    fmt.Println(numberOfAlternatingGroups1([]int{0,1,0,1,0,1,0,1,0,1}, 4)) // 10
    fmt.Println(numberOfAlternatingGroups1([]int{1,0,1,0,1,0,1,0,1,0}, 4)) // 10
}