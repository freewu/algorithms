package main

// 面试题 08.13. Pile Box LCCI
// You have a stack of n boxes, with widths wi, depths di, and heights hi. 
// The boxes cannot be rotated and can only be stacked on top of one another if each box in the stack is strictly larger than the box above it in width, height, and depth. 
// Implement a method to compute the height of the tallest possible stack. 
// The height of a stack is the sum of the heights of each box.

// The input use [wi, di, hi] to represents each box.

// Example1:
// Input: box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
// Output: 6

// Example2:
// Input: box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
// Output: 10

// Note:
//     box.length <= 3000

import "fmt"
import "sort"

func pileBox(box [][]int) int {
    sort.Slice(box, func(i, j int) bool {
        if box[i][2] != box[j][2] {
            return box[i][2] < box[j][2]
        } else if box[i][1] != box[j][1] {
            return box[i][1] < box[j][1]
        } else {
            return box[i][0] < box[j][0]
        }
    })
    n := len(box)
    dp := make([]int, n)
    // 以每一个箱子为第的最大高度
    for i := 0; i < n; i++ {
        dp[i] = box[i][2]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            // 判断是否满足叠箱子的条件
            if box[i][0] > box[j][0] && box[i][1] > box[j][1] && box[i][2] > box[j][2] {
                dp[i] = max(dp[i] , dp[j] + box[i][2])
            }
        }
    }
    res:= -1 << 31
    for _, v :=range dp {
        res = max(res, v)
    }
    return res
}

func main() {
    // Example1:
    // Input: box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
    // Output: 6
    fmt.Println(pileBox([][]int{{1,1,1},{2,2,2},{3,3,3}})) // 6
    // Example2:
    // Input: box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
    // Output: 10
    fmt.Println(pileBox([][]int{{1,1,1},{2,3,4},{2,6,7},{3,4,5}})) // 10
}