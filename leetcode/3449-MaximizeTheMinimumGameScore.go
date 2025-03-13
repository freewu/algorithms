package main

// 3449. Maximize the Minimum Game Score
// You are given an array points of size n and an integer m. 
// There is another array gameScore of size n, where gameScore[i] represents the score achieved at the ith game. 
// Initially, gameScore[i] == 0 for all i.

// You start at index -1, which is outside the array (before the first position at index 0). 
// You can make at most m moves. In each move, you can either:
//     1. Increase the index by 1 and add points[i] to gameScore[i].
//     2. Decrease the index by 1 and add points[i] to gameScore[i].

// Note that the index must always remain within the bounds of the array after the first move.

// Return the maximum possible minimum value in gameScore after at most m moves.

// Example 1:
// Input: points = [2,4], m = 3
// Output: 4
// Explanation:
// Initially, index i = -1 and gameScore = [0, 0].
// Move	Index	gameScore
// Increase i	0	[2, 0]
// Increase i	1	[2, 4]
// Decrease i	0	[4, 4]
// The minimum value in gameScore is 4, and this is the maximum possible minimum among all configurations. Hence, 4 is the output.

// Example 2:
// Input: points = [1,2,3], m = 5
// Output: 2
// Explanation:
// Initially, index i = -1 and gameScore = [0, 0, 0].
// Move	Index	gameScore
// Increase i	0	[1, 0, 0]
// Increase i	1	[1, 2, 0]
// Decrease i	0	[2, 2, 0]
// Increase i	1	[2, 4, 0]
// Increase i	2	[2, 4, 3]
// The minimum value in gameScore is 2, and this is the maximum possible minimum among all configurations. Hence, 2 is the output.

// Constraints:
//     2 <= n == points.length <= 5 * 10^4
//     1 <= points[i] <= 10^6
//     1 <= m <= 10^9

import "fmt"
import "sort"
import "slices"

func maxScore(points []int, m int) int64 {
    check := func(bound int) bool {
        a, b, moves, n := 0, 0, 0, len(points)
        for i := 0; i < n - 1; i++ {
            if a >= bound {
                moves++
                a, b = b, 0
                if i + 1 == n - 1 {
                    if a < bound {
                        k := (bound - a + points[i + 1] - 1) / points[i+1]
                        moves += 2*k - 1
                        if moves > m {
                            return false
                        }
                    }
                    break
                } else {
                    continue
                }
            }
            k := (bound - a + points[i] - 1) / points[i]
            a, b = b + points[i+1]*(k-1), 0
            moves += 2 * k - 1
            if i + 1 == n - 1 {
                if a < bound {
                    k = (bound - a + points[i+1] - 1) / points[i+1]
                    moves += 2*k - 1
                }
            }
            if moves > m {
                return false
            }
        }
        return true
    }
    l, r := 0, 500000000000001
    for l < r {
        t := (l + r) / 2
        if check(t) {
            l = t + 1
        } else {
            r = t
        }
    }
    return int64(l - 1)
}

func maxScore1(points []int, m int) int64 {
    right := (m + 1) / 2 * slices.Min(points)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := sort.Search(right, func(low int) bool {
        // 二分最小的不满足要求的 low + 1，即可得到最大的满足要求的 low
        low++
        left, pre := m, 0
        for i, p := range points {
            k := (low - 1)/p + 1 - pre // 还需要操作的次数
            if i == len(points) - 1 && k <= 0 { // 最后一个数已经满足要求
                break
            }
            k = max(k, 1) // 至少要走 1 步
            left -= k * 2 - 1 // 左右横跳
            if left < 0 {
                return true
            }
            pre = k - 1 // 右边那个数顺带操作了 k-1 次
        }
        return false
    })
    return int64(res)
}

func main() {
    // Example 1:
    // Input: points = [2,4], m = 3
    // Output: 4
    // Explanation:
    // Initially, index i = -1 and gameScore = [0, 0].
    // Move	Index	gameScore
    // Increase i	0	[2, 0]
    // Increase i	1	[2, 4]
    // Decrease i	0	[4, 4]
    // The minimum value in gameScore is 4, and this is the maximum possible minimum among all configurations. Hence, 4 is the output.
    fmt.Println(maxScore([]int{2,4}, 3)) // 4
    // Example 2:
    // Input: points = [1,2,3], m = 5
    // Output: 2
    // Explanation:
    // Initially, index i = -1 and gameScore = [0, 0, 0].
    // Move	Index	gameScore
    // Increase i	0	[1, 0, 0]
    // Increase i	1	[1, 2, 0]
    // Decrease i	0	[2, 2, 0]
    // Increase i	1	[2, 4, 0]
    // Increase i	2	[2, 4, 3]
    // The minimum value in gameScore is 2, and this is the maximum possible minimum among all configurations. Hence, 2 is the output.
    fmt.Println(maxScore([]int{1,2,3}, 5)) // 2

    fmt.Println(maxScore([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0
    fmt.Println(maxScore([]int{9,8,7,6,5,4,3,2,1}, 2)) // 0

    fmt.Println(maxScore1([]int{2,4}, 3)) // 4
    fmt.Println(maxScore1([]int{1,2,3}, 5)) // 2
    fmt.Println(maxScore1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0
    fmt.Println(maxScore1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 0
}