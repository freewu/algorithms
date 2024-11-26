package main

// 2078. Two Furthest Houses With Different Colors
// There are n houses evenly lined up on the street, and each house is beautifully painted. 
// You are given a 0-indexed integer array colors of length n, where colors[i] represents the color of the ith house.

// Return the maximum distance between two houses with different colors.

// The distance between the ith and jth houses is abs(i - j), where abs(x) is the absolute value of x.

// Example 1:
// <img scr="https://assets.leetcode.com/uploads/2021/10/31/eg1.png" />
// Input: colors = [1,1,1,6,1,1,1]
// Output: 3
// Explanation: In the above image, color 1 is blue, and color 6 is red.
// The furthest two houses with different colors are house 0 and house 3.
// House 0 has color 1, and house 3 has color 6. The distance between them is abs(0 - 3) = 3.
// Note that houses 3 and 6 can also produce the optimal answer.

// Example 2:
// <img scr="https://assets.leetcode.com/uploads/2021/10/31/eg2.png" />
// Input: colors = [1,8,3,8,3]
// Output: 4
// Explanation: In the above image, color 1 is blue, color 8 is yellow, and color 3 is green.
// The furthest two houses with different colors are house 0 and house 4.
// House 0 has color 1, and house 4 has color 3. The distance between them is abs(0 - 4) = 4.

// Example 3:
// Input: colors = [0,1]
// Output: 1
// Explanation: The furthest two houses with different colors are house 0 and house 1.
// House 0 has color 0, and house 1 has color 1. The distance between them is abs(0 - 1) = 1.

// Constraints:
//     n == colors.length
//     2 <= n <= 100
//     0 <= colors[i] <= 100
//     Test data are generated such that at least two houses have different colors.

import "fmt"

func maxDistance(colors []int) int {
    n := len(colors)
    i, j := 0, n - 1
    for i < n{
        if colors[i] != colors[n - 1] { return n - 1 - i }
        i++
        if colors[j] != colors[0] { return j }
        j--
    }
    return -1
}

func maxDistance1(colors []int) int {
    n := len(colors)
    if colors[0] != colors[n - 1] {  return n - 1 }
    c, l, r := colors[0], 1, n - 2
    for l < n && colors[l] == c {l++}
    for r >= 0 && colors[r] == c {r--}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(r, n - 1 - l)
}

func main() {
    // Example 1:
    // <img scr="https://assets.leetcode.com/uploads/2021/10/31/eg1.png" />
    // Input: colors = [1,1,1,6,1,1,1]
    // Output: 3
    // Explanation: In the above image, color 1 is blue, and color 6 is red.
    // The furthest two houses with different colors are house 0 and house 3.
    // House 0 has color 1, and house 3 has color 6. The distance between them is abs(0 - 3) = 3.
    // Note that houses 3 and 6 can also produce the optimal answer.
    fmt.Println(maxDistance([]int{1,1,1,6,1,1,1})) // 3
    // Example 2:
    // <img scr="https://assets.leetcode.com/uploads/2021/10/31/eg2.png" />
    // Input: colors = [1,8,3,8,3]
    // Output: 4
    // Explanation: In the above image, color 1 is blue, color 8 is yellow, and color 3 is green.
    // The furthest two houses with different colors are house 0 and house 4.
    // House 0 has color 1, and house 4 has color 3. The distance between them is abs(0 - 4) = 4.
    fmt.Println(maxDistance([]int{1,8,3,8,3})) // 4
    // Example 3:
    // Input: colors = [0,1]
    // Output: 1
    // Explanation: The furthest two houses with different colors are house 0 and house 1.
    // House 0 has color 0, and house 1 has color 1. The distance between them is abs(0 - 1) = 1.
    fmt.Println(maxDistance([]int{0,1})) // 1

    fmt.Println(maxDistance([]int{1,1,1,1,1,1,1,1,1})) // -1

    fmt.Println(maxDistance1([]int{1,1,1,6,1,1,1})) // 3
    fmt.Println(maxDistance1([]int{1,8,3,8,3})) // 4
    fmt.Println(maxDistance1([]int{0,1})) // 1
    fmt.Println(maxDistance1([]int{1,1,1,1,1,1,1,1,1})) // -1
}