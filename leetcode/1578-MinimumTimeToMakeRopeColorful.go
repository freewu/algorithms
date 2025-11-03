package main

// 1578. Minimum Time to Make Rope Colorful
// Alice has n balloons arranged on a rope. 
// You are given a 0-indexed string colors where colors[i] is the color of the ith balloon.

// Alice wants the rope to be colorful. 
// She does not want two consecutive balloons to be of the same color, so she asks Bob for help. 
// Bob can remove some balloons from the rope to make it colorful. 
// You are given a 0-indexed integer array neededTime where neededTime[i] is the time (in seconds) that Bob needs to remove the ith balloon from the rope.

// Return the minimum time Bob needs to make the rope colorful.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/13/ballon1.jpg" />
// Input: colors = "abaac", neededTime = [1,2,3,4,5]
// Output: 3
// Explanation: In the above image, 'a' is blue, 'b' is red, and 'c' is green.
// Bob can remove the blue balloon at index 2. This takes 3 seconds.
// There are no longer two consecutive balloons of the same color. Total time = 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/13/balloon2.jpg" />
// Input: colors = "abc", neededTime = [1,2,3]
// Output: 0
// Explanation: The rope is already colorful. Bob does not need to remove any balloons from the rope.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/12/13/balloon3.jpg" />
// Input: colors = "aabaa", neededTime = [1,2,3,4,1]
// Output: 2
// Explanation: Bob will remove the balloons at indices 0 and 4. Each balloons takes 1 second to remove.
// There are no longer two consecutive balloons of the same color. Total time = 1 + 1 = 2.

// Constraints:
//     n == colors.length == neededTime.length
//     1 <= n <= 10^5
//     1 <= neededTime[i] <= 10^4
//     colors contains only lowercase English letters.

import "fmt"

func minCost(colors string, neededTime []int) int {
    res := 0
    for i := 0; i < len(colors); i++ {
        sum, mx, it := neededTime[i], neededTime[i], 0
        for j := i + 1; j < len(colors); j++ {
            if colors[i] != colors[j] {
                break
            } else {
                it++
                sum += neededTime[j]
                if neededTime[j] > mx {
                    mx = neededTime[j]
                }
            }
        }
        if it != 0 {
            res += sum - mx
            i = i + it
        }
    }
    return res
}

func minCost1(colors string, neededTime []int) int {
    // 分组循环 + 贪心
    // 贪心: (移除的气球越少越好=>连续相同的气球保留一个)
    // 让每一组连续相同颜色的变为1个
    // 每组花费时间最少=> 要保留花费时间最大的那个  即: sum(group)-max(group[i])
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, n := 0, len(colors); i < n; {
        cur, mx, sum := colors[i], 0, 0
        for i < n && colors[i] == cur { // 循环是包含start自身的,所以mx,sum不用在外面处理  [start,i)
            mx = max(mx, neededTime[i])
            sum += neededTime[i]
            i++
        }
        res += sum - mx // 留下花费时间最长的那个,其余全移除
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/13/ballon1.jpg" />
    // Input: colors = "abaac", neededTime = [1,2,3,4,5]
    // Output: 3
    // Explanation: In the above image, 'a' is blue, 'b' is red, and 'c' is green.
    // Bob can remove the blue balloon at index 2. This takes 3 seconds.
    // There are no longer two consecutive balloons of the same color. Total time = 3.
    fmt.Println(minCost("abaac", []int{1,2,3,4,5})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/13/balloon2.jpg" />
    // Input: colors = "abc", neededTime = [1,2,3]
    // Output: 0
    // Explanation: The rope is already colorful. Bob does not need to remove any balloons from the rope.
    fmt.Println(minCost("abc", []int{1,2,3})) // 0
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/12/13/balloon3.jpg" />
    // Input: colors = "aabaa", neededTime = [1,2,3,4,1]
    // Output: 2
    // Explanation: Bob will remove the balloons at indices 0 and 4. Each balloons takes 1 second to remove.
    // There are no longer two consecutive balloons of the same color. Total time = 1 + 1 = 2.
    fmt.Println(minCost("aabaa", []int{1,2,3,4,1})) // 2

    fmt.Println(minCost("leetcode", []int{1,2,3,4,5,6,7,8})) // 2
    fmt.Println(minCost("bluefrog", []int{1,2,3,4,5,6,7,8})) // 0

    fmt.Println(minCost1("abaac", []int{1,2,3,4,5})) // 3
    fmt.Println(minCost1("abc", []int{1,2,3})) // 0
    fmt.Println(minCost1("aabaa", []int{1,2,3,4,1})) // 2
    fmt.Println(minCost1("leetcode", []int{1,2,3,4,5,6,7,8})) // 2
    fmt.Println(minCost1("bluefrog", []int{1,2,3,4,5,6,7,8})) // 0
}