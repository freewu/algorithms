package main

// 3951. Minimum Energy to Maintain Brightness
// You are given an integer n, representing n light bulbs arranged in a line and indexed from 0 to n - 1.

// You are also given an integer brightness and a 2D integer array intervals, 
// where intervals[i] = [starti, endi] represents an inclusive time interval during which the lighting requirement must be satisfied.

// At each time unit, every bulb can independently be either on or off. 
// A bulb that is on illuminates its own position and its adjacent positions, if they exist.

// The total illumination at a time unit is the number of illuminated positions. 
// Each position is counted at most once.

// For every integer time unit covered by at least one interval in intervals, the total illumination must be at least brightness. 
// At time units not covered by any interval, all bulbs may remain off. 
// Each bulb that is on consumes 1 unit of energy for that time unit.

// Return an integer denoting the minimum total energy required.

// Example 1:
// Input: n = 5, brightness = 5, intervals = [[6,12]]
// Output: 14
// Explanation:
// Turn on the light bulbs at positions 1 and 4.
// Current state of line: 0 1 0 0 1.
// All 5 positions are illuminated, so the required brightness is reached.
// The active interval has length 12 - 6 + 1 = 7, so the total energy is 2 * 7 = 14.

// Example 2:
// Input: n = 2, brightness = 1, intervals = [[0,0],[2,2]]
// Output: 2
// Explanation:
// Turn on one light bulb during each active interval.
// Each interval has length 1, so the total active time is 1 + 1 = 2.
// The total energy is 1 * 2 = 2.

// Example 3:
// Input: n = 4, brightness = 2, intervals = [[1,3],[2,4]]
// Output: 4
// Explanation:
// Turn on one light bulb. It can illuminate at least 2 positions.
// The active intervals overlap, so the total active time is the length of [1,4], which is 4.
// The total energy is 1 * 4 = 4.
 
// Constraints:
//     1 <= n <= 10^6
//     1 <= brightness <= n
//     1 <= intervals.length <= 10^5
//     intervals[i] == [starti, endi]
//     0 <= starti <= endi <= 10^9

import "fmt"
import "slices"

func minEnergy(n int, brightness int, intervals [][]int) int64 {
    slices.SortFunc(intervals, func(p, q []int) int { 
        return p[0] - q[0] // 按照左端点从小到大排序
    })
    l, left, right := 0, 0, -1
    for _, p := range intervals {
        if p[0] <= right { // 左端点在合并区间内，可以合并
            right = max(right, p[1]) // 更新合并区间的右端点
        } else { // 不相交，无法合并
            l += right - left + 1
            left, right = p[0], p[1] // 新的合并区间
        }
    }
    l += right - left + 1
    bulbs := (brightness + 2) / 3 // 至少要开启 bulbs 个灯泡
    return int64(bulbs * l)
}

func minEnergy1(n int, brightness int, intervals [][]int) int64 {
    slices.SortFunc(intervals, func(p, q []int) int { 
        return p[0] - q[0] // 按照左端点从小到大排序
    })
    sum, x, y := 0, intervals[0][0], intervals[0][1]
    for i,itr := range intervals[1:] {
        if intervals[i][1] >= intervals[i + 1][1] {
            continue
        }
        a, b := itr[0],itr[1]
        if a > y + 1 {
            sum += y - x + 1
            x = a
            y = b
        }
        if b > y {
            y = b
        }
    }
    k := (brightness - 1) / 3 + 1
    sum += y - x + 1
    return int64(k * sum)
}

func main() {
    // Example 1:
    // Input: n = 5, brightness = 5, intervals = [[6,12]]
    // Output: 14
    // Explanation:
    // Turn on the light bulbs at positions 1 and 4.
    // Current state of line: 0 1 0 0 1.
    // All 5 positions are illuminated, so the required brightness is reached.
    // The active interval has length 12 - 6 + 1 = 7, so the total energy is 2 * 7 = 14.
    fmt.Println(minEnergy(5, 5, [][]int{{6,12}})) // 14
    // Example 2:
    // Input: n = 2, brightness = 1, intervals = [[0,0],[2,2]]
    // Output: 2
    // Explanation:
    // Turn on one light bulb during each active interval.
    // Each interval has length 1, so the total active time is 1 + 1 = 2.
    // The total energy is 1 * 2 = 2.
    fmt.Println(minEnergy(2, 1, [][]int{{0,0},{2,2}})) // 2
    // Example 3:
    // Input: n = 4, brightness = 2, intervals = [[1,3],[2,4]]
    // Output: 4
    // Explanation:
    // Turn on one light bulb. It can illuminate at least 2 positions.
    // The active intervals overlap, so the total active time is the length of [1,4], which is 4.
    // The total energy is 1 * 4 = 4.
    fmt.Println(minEnergy(4, 2, [][]int{{1,3},{2,4}})) // 4

    fmt.Println(minEnergy1(5, 5, [][]int{{6,12}})) // 14
    fmt.Println(minEnergy1(2, 1, [][]int{{0,0},{2,2}})) // 2
    fmt.Println(minEnergy1(4, 2, [][]int{{1,3},{2,4}})) // 4
}