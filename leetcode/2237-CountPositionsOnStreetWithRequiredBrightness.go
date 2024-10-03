package main

// 2237. Count Positions on Street With Required Brightness
// You are given an integer n. 
// A perfectly straight street is represented by a number line ranging from 0 to n - 1. 
// You are given a 2D integer array lights representing the street lamp(s) on the street. 
// Each lights[i] = [positioni, rangei] indicates that there is a street lamp at position positioni 
// that lights up the area from [max(0, positioni - rangei), min(n - 1, positioni + rangei)] (inclusive).

// The brightness of a position p is defined as the number of street lamps that light up the position p. 
// You are given a 0-indexed integer array requirement of size n where requirement[i] is the minimum brightness of the ith position on the street.

// Return the number of positions i on the street between 0 and n - 1 
// that have a brightness of at least requirement[i].

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/04/11/screenshot-2022-04-11-at-22-24-43-diagramdrawio-diagramsnet.png" />
// Input: n = 5, lights = [[0,1],[2,1],[3,2]], requirement = [0,2,1,4,1]
// Output: 4
// Explanation:
// - The first street lamp lights up the area from [max(0, 0 - 1), min(n - 1, 0 + 1)] = [0, 1] (inclusive).
// - The second street lamp lights up the area from [max(0, 2 - 1), min(n - 1, 2 + 1)] = [1, 3] (inclusive).
// - The third street lamp lights up the area from [max(0, 3 - 2), min(n - 1, 3 + 2)] = [1, 4] (inclusive).
// - Position 0 is covered by the first street lamp. It is covered by 1 street lamp which is greater than requirement[0].
// - Position 1 is covered by the first, second, and third street lamps. It is covered by 3 street lamps which is greater than requirement[1].
// - Position 2 is covered by the second and third street lamps. It is covered by 2 street lamps which is greater than requirement[2].
// - Position 3 is covered by the second and third street lamps. It is covered by 2 street lamps which is less than requirement[3].
// - Position 4 is covered by the third street lamp. It is covered by 1 street lamp which is equal to requirement[4].
// Positions 0, 1, 2, and 4 meet the requirement so we return 4.

// Example 2:
// Input: n = 1, lights = [[0,1]], requirement = [2]
// Output: 0
// Explanation:
// - The first street lamp lights up the area from [max(0, 0 - 1), min(n - 1, 0 + 1)] = [0, 0] (inclusive).
// - Position 0 is covered by the first street lamp. It is covered by 1 street lamp which is less than requirement[0].
// - We return 0 because no position meets their brightness requirement.

// Constraints:
//     1 <= n <= 10^5
//     1 <= lights.length <= 10^5
//     0 <= positioni < n
//     0 <= rangei <= 10^5
//     requirement.length == n
//     0 <= requirement[i] <= 10^5

import "fmt"

// 差分数组
func meetRequirement(n int, lights [][]int, requirement []int) int {
    res, prefixSum := 0, make([]int, n)
    for i := 0; i < len(lights); i++ {
        start:= lights[i][0] - lights[i][1]
        if start < 0 {
            start = 0
        }
        prefixSum[start]++
        end := lights[i][0] + lights[i][1]+1
        if end < n {
            prefixSum[end]--
        }
    }
    for i := 0; i < n; i++ {
        if i > 0 {
            prefixSum[i] += prefixSum[i-1]
        }
        if prefixSum[i] >= requirement[i] {
            res++
        }
    }
    return res
}

func meetRequirement1(n int, lights [][]int, requirement []int) int {
    res, diff := 0, make([]int, n + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, light := range lights {
        left, right := max(0, light[0]-light[1]), min(n-1, light[0]+light[1])
        diff[left]++
        diff[right+1]-- // 照亮的区域为 [left,right],所以这个问题中应该从right+1减小
    }
    if diff[0] >= requirement[0] {
        res++
    }
    for i := 1; i < n; i++ {
        diff[i] += diff[i-1]
        if diff[i] >= requirement[i] {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/04/11/screenshot-2022-04-11-at-22-24-43-diagramdrawio-diagramsnet.png" />
    // Input: n = 5, lights = [[0,1],[2,1],[3,2]], requirement = [0,2,1,4,1]
    // Output: 4
    // Explanation:
    // - The first street lamp lights up the area from [max(0, 0 - 1), min(n - 1, 0 + 1)] = [0, 1] (inclusive).
    // - The second street lamp lights up the area from [max(0, 2 - 1), min(n - 1, 2 + 1)] = [1, 3] (inclusive).
    // - The third street lamp lights up the area from [max(0, 3 - 2), min(n - 1, 3 + 2)] = [1, 4] (inclusive).
    // - Position 0 is covered by the first street lamp. It is covered by 1 street lamp which is greater than requirement[0].
    // - Position 1 is covered by the first, second, and third street lamps. It is covered by 3 street lamps which is greater than requirement[1].
    // - Position 2 is covered by the second and third street lamps. It is covered by 2 street lamps which is greater than requirement[2].
    // - Position 3 is covered by the second and third street lamps. It is covered by 2 street lamps which is less than requirement[3].
    // - Position 4 is covered by the third street lamp. It is covered by 1 street lamp which is equal to requirement[4].
    // Positions 0, 1, 2, and 4 meet the requirement so we return 4.
    fmt.Println(meetRequirement(5, [][]int{{0,1},{2,1},{3,2}}, []int{0,2,1,4,1})) // 4
    // Example 2:
    // Input: n = 1, lights = [[0,1]], requirement = [2]
    // Output: 0
    // Explanation:
    // - The first street lamp lights up the area from [max(0, 0 - 1), min(n - 1, 0 + 1)] = [0, 0] (inclusive).
    // - Position 0 is covered by the first street lamp. It is covered by 1 street lamp which is less than requirement[0].
    // - We return 0 because no position meets their brightness requirement.
    fmt.Println(meetRequirement(1, [][]int{{0,1}}, []int{2})) // 0

    fmt.Println(meetRequirement1(5, [][]int{{0,1},{2,1},{3,2}}, []int{0,2,1,4,1})) // 4
    fmt.Println(meetRequirement1(1, [][]int{{0,1}}, []int{2})) // 0
}