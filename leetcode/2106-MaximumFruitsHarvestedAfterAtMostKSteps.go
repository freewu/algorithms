package main

// 2106. Maximum Fruits Harvested After at Most K Steps
// Fruits are available at some positions on an infinite x-axis. 
// You are given a 2D integer array fruits where fruits[i] = [positioni, amounti] depicts amounti fruits at the position positioni. 
//  is already sorted by positioni in ascending order, and each positioni is unique.

// You are also given an integer startPos and an integer k. Initially, you are at the position startPos. 
// From any position, you can either walk to the left or right. 
// It takes one step to move one unit on the x-axis, and you can walk at most k steps in total. 
// For every position you reach, you harvest all the fruits at that position, and the fruits will disappear from that position.

// Return the maximum total number of fruits you can harvest.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/21/1.png" />
// Input: fruits = [[2,8],[6,3],[8,6]], startPos = 5, k = 4
// Output: 9
// Explanation: 
// The optimal way is to:
// - Move right to position 6 and harvest 3 fruits
// - Move right to position 8 and harvest 6 fruits
// You moved 3 steps and harvested 3 + 6 = 9 fruits in total.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/21/2.png" />
// Input: fruits = [[0,9],[4,1],[5,7],[6,2],[7,4],[10,9]], startPos = 5, k = 4
// Output: 14
// Explanation: 
// You can move at most k = 4 steps, so you cannot reach position 0 nor 10.
// The optimal way is to:
// - Harvest the 7 fruits at the starting position 5
// - Move left to position 4 and harvest 1 fruit
// - Move right to position 6 and harvest 2 fruits
// - Move right to position 7 and harvest 4 fruits
// You moved 1 + 3 = 4 steps and harvested 7 + 1 + 2 + 4 = 14 fruits in total.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/11/21/3.png" />
// Input: fruits = [[0,3],[6,4],[8,5]], startPos = 3, k = 2
// Output: 0
// Explanation:
// You can move at most k = 2 steps and cannot reach any position with fruits.
 
// Constraints:
//     1 <= fruits.length <= 10^5
//     fruits[i].length == 2
//     0 <= startPos, positioni <= 2 * 10^5
//     positioni-1 < positioni for any i > 0 (0-indexed)
//     1 <= amounti <= 10^4
//     0 <= k <= 2 * 10^5

import "fmt"
import "sort"

// func maxTotalFruits(fruits [][]int, startPos int, k int) int {
//     j, sum, res, n := 0, 0, 0, len(fruits)
//     for j < n && fruits[j][0] < startPos - k {
//         j++
//     }
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     for i := j; i < n && fruits[i][0] <= startPos + k; i++ {
//         sum += fruits[i][1]
//         for min(startPos - 2 * fruits[j][0] + fruits[i][0], 2 * fruits[i][0] - fruits[j][0] - startPos) > k {
//             j++
//             sum -= fruits[j][1]
//         }
//         res = max(res, sum)
//     }
//     return res
// }

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
    res, left, right, sum, n  := 0, 0, 0, 0, len(fruits)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    step := func(left int, right int) int {
        if fruits[right][0] <= startPos {
            return startPos - fruits[left][0]
        } else if fruits[left][0] >= startPos {
            return fruits[right][0] - startPos
        } else {
            return min(abs(startPos-fruits[right][0]), abs(startPos-fruits[left][0])) + fruits[right][0] - fruits[left][0]
        }
    }
    for i, fruit := range fruits{
        if startPos - fruit[0] <= k{
            left, right = i, i
            break
        }
    }
    // 每次固定住窗口右边界
    for right < n {
        sum += fruits[right][1]
        // 移动左边界
        for left <= right && step(left, right) > k {
            sum -= fruits[left][1]
            left++
        }
        res = max(res, sum)
        right++
    }
    return res
}

func maxTotalFruits1(fruits [][]int, startPos int, k int) int {
    n := len(fruits)
    left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos - k })
    right, sum := left, 0
    for ; right < n && fruits[right][0] <= startPos; right++ {
        sum += fruits[right][1]
    }
    res := sum
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for ; right < n && fruits[right][0] <= startPos + k; right++ {
        sum += fruits[right][1]
        for fruits[right][0] * 2 - fruits[left][0] - startPos > k &&
            fruits[right][0] - fruits[left][0] * 2 + startPos > k {
            sum -= fruits[left][1] // fruits[left][0] 无法到达
            left++
        }
        res = max(res, sum) // 更新答案最大值
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/11/21/1.png" />
    // Input: fruits = [[2,8],[6,3],[8,6]], startPos = 5, k = 4
    // Output: 9
    // Explanation: 
    // The optimal way is to:
    // - Move right to position 6 and harvest 3 fruits
    // - Move right to position 8 and harvest 6 fruits
    // You moved 3 steps and harvested 3 + 6 = 9 fruits in total.
    fmt.Println(maxTotalFruits([][]int{{2,8},{6,3},{8,6}}, 5, 4)) // 9
    // Example 2: 
    // <img src="https://assets.leetcode.com/uploads/2021/11/21/2.png" />
    // Input: fruits = [[0,9],[4,1],[5,7],[6,2],[7,4],[10,9]], startPos = 5, k = 4
    // Output: 14
    // Explanation: 
    // You can move at most k = 4 steps, so you cannot reach position 0 nor 10.
    // The optimal way is to:
    // - Harvest the 7 fruits at the starting position 5
    // - Move left to position 4 and harvest 1 fruit
    // - Move right to position 6 and harvest 2 fruits
    // - Move right to position 7 and harvest 4 fruits
    // You moved 1 + 3 = 4 steps and harvested 7 + 1 + 2 + 4 = 14 fruits in total.
    fmt.Println(maxTotalFruits([][]int{{0,9},{4,1},{5,7},{6,2},{7,4},{10,9}}, 5, 4)) // 14
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/11/21/3.png" />
    // Input: fruits = [[0,3],[6,4],[8,5]], startPos = 3, k = 2
    // Output: 0
    // Explanation:
    // You can move at most k = 2 steps and cannot reach any position with fruits.
    fmt.Println(maxTotalFruits([][]int{{0,3},{6,4},{8,5}}, 3, 2)) // 0

    fmt.Println(maxTotalFruits1([][]int{{2,8},{6,3},{8,6}}, 5, 4)) // 9
    fmt.Println(maxTotalFruits1([][]int{{0,9},{4,1},{5,7},{6,2},{7,4},{10,9}}, 5, 4)) // 14
    fmt.Println(maxTotalFruits1([][]int{{0,3},{6,4},{8,5}}, 3, 2)) // 0
}