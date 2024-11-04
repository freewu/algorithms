package main

// 1742. Maximum Number of Balls in a Box
// You are working in a ball factory where you have n balls numbered from lowLimit up to highLimit inclusive (i.e., n == highLimit - lowLimit + 1), 
// and an infinite number of boxes numbered from 1 to infinity.

// Your job at this factory is to put each ball in the box with a number equal to the sum of digits of the ball's number. 
// For example, the ball number 321 will be put in the box number 3 + 2 + 1 = 6 
// and the ball number 10 will be put in the box number 1 + 0 = 1.

// Given two integers lowLimit and highLimit, return the number of balls in the box with the most balls.

// Example 1:
// Input: lowLimit = 1, highLimit = 10
// Output: 2
// Explanation:
// Box Number:  1 2 3 4 5 6 7 8 9 10 11 ...
// Ball Count:  2 1 1 1 1 1 1 1 1 0  0  ...
// Box 1 has the most number of balls with 2 balls.

// Example 2:
// Input: lowLimit = 5, highLimit = 15
// Output: 2
// Explanation:
// Box Number:  1 2 3 4 5 6 7 8 9 10 11 ...
// Ball Count:  1 1 1 1 2 2 1 1 1 0  0  ...
// Boxes 5 and 6 have the most number of balls with 2 balls in each.

// Example 3:
// Input: lowLimit = 19, highLimit = 28
// Output: 2
// Explanation:
// Box Number:  1 2 3 4 5 6 7 8 9 10 11 12 ...
// Ball Count:  0 1 1 1 1 1 1 1 1 2  0  0  ...
// Box 10 has the most number of balls with 2 balls.

// Constraints:
//     1 <= lowLimit <= highLimit <= 10^5

import "fmt"
import "strconv"

func countBalls(lowLimit int, highLimit int) int {
    res, sum, box := -1, 0, make([]int, 46)
    for i := lowLimit; i <= highLimit; i++ {
        sum = 0
        for j := 0; j < len(string(strconv.Itoa(i))); j++ {
            a, _ := strconv.Atoi(string(strconv.Itoa(i))[j : j+1])
            sum = sum + a
        }
        box[sum] = box[sum] + 1
    }
    for i := 0; i < len(box); i++ {
        if res < box[i] {
            res = box[i]
        }
    }
    return res
}

func countBalls1(lowLimit int, highLimit int) int {
    res, box := 0, make([]int, 46)
    max := func(x, y int) int { if x > y { return x; }; return y; }
    tonum := func(i int) int {
        res := 0
        for ; i > 0; i /= 10 {
            res += i % 10
        }
        return res
    }
    for i := lowLimit; i <= highLimit; i++ {
        v := tonum(i)
        box[v]++
        res = max(res, box[v])
    }
    return res
}

func main() {
    // Example 1:
    // Input: lowLimit = 1, highLimit = 10
    // Output: 2
    // Explanation:
    // Box Number:  1 2 3 4 5 6 7 8 9 10 11 ...
    // Ball Count:  2 1 1 1 1 1 1 1 1 0  0  ...
    // Box 1 has the most number of balls with 2 balls.
    fmt.Println(countBalls(1, 10)) // 2
    // Example 2:
    // Input: lowLimit = 5, highLimit = 15
    // Output: 2
    // Explanation:
    // Box Number:  1 2 3 4 5 6 7 8 9 10 11 ...
    // Ball Count:  1 1 1 1 2 2 1 1 1 0  0  ...
    // Boxes 5 and 6 have the most number of balls with 2 balls in each.
    fmt.Println(countBalls(5, 15)) // 2
    // Example 3:
    // Input: lowLimit = 19, highLimit = 28
    // Output: 2
    // Explanation:
    // Box Number:  1 2 3 4 5 6 7 8 9 10 11 12 ...
    // Ball Count:  0 1 1 1 1 1 1 1 1 2  0  0  ...
    // Box 10 has the most number of balls with 2 balls.
    fmt.Println(countBalls(19, 28)) // 2

    fmt.Println(countBalls(1, 1)) // 1
    fmt.Println(countBalls(10000, 10000)) // 1
    fmt.Println(countBalls(1, 10000)) // 670
    fmt.Println(countBalls(10000, 1)) // 0

    fmt.Println(countBalls1(1, 10)) // 2
    fmt.Println(countBalls1(5, 15)) // 2
    fmt.Println(countBalls1(19, 28)) // 2
    fmt.Println(countBalls1(1, 1)) // 1
    fmt.Println(countBalls1(10000, 10000)) // 1
    fmt.Println(countBalls1(1, 10000)) // 670
    fmt.Println(countBalls1(10000, 1)) // 0
}