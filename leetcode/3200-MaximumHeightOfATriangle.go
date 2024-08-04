package main

// 3200. Maximum Height of a Triangle
// You are given two integers red and blue representing the count of red and blue colored balls. 
// You have to arrange these balls to form a triangle such that the 1st row will have 1 ball, the 2nd row will have 2 balls, the 3rd row will have 3 balls, and so on.

// All the balls in a particular row should be the same color, and adjacent rows should have different colors.
// Return the maximum height of the triangle that can be achieved.

// Example 1:
// Input: red = 2, blue = 4
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/16/brb.png" />
// The only possible arrangement is shown above.

// Example 2:
// Input: red = 2, blue = 1
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/16/br.png" />
// The only possible arrangement is shown above.

// Example 3:
// Input: red = 1, blue = 1
// Output: 1

// Example 4:
// Input: red = 10, blue = 1
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/16/br.png" />
// The only possible arrangement is shown above.

// Constraints:
//     1 <= red, blue <= 100

import "fmt"

func maxHeightOfTriangle(red int, blue int) int {
    check := func (a int, b int)int{
        h := 0
        for i := 1; ; i++ {
            if i % 2 != 0 {
                if a - i >= 0 {
                    a -= i
                    h++
                } else {
                    return h
                }
            } else {
                if b - i >= 0 {
                    b -= i
                    h++
                } else {
                    return h
                }
            }
        }
    }
    h1, h2 := check(red, blue), check(blue, red)
    if h1 > h2 {
        return h1
    }
    return h2
}

func main() {
    // Example 1:
    // Input: red = 2, blue = 4
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/16/brb.png" />
    // The only possible arrangement is shown above.
    fmt.Println(maxHeightOfTriangle(2,4)) // 3
    // Example 2:
    // Input: red = 2, blue = 1
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/16/br.png" />
    // The only possible arrangement is shown above.
    fmt.Println(maxHeightOfTriangle(2,1)) // 2
    // Example 3:
    // Input: red = 1, blue = 1
    // Output: 1
    fmt.Println(maxHeightOfTriangle(1,1)) // 1
    // Example 4:
    // Input: red = 10, blue = 1
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/16/br.png" />
    // The only possible arrangement is shown above.
    fmt.Println(maxHeightOfTriangle(10,1)) // 2
}