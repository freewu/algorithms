package main 

// 256. Paint House
// There is a row of n houses, where each house can be painted one of three colors: red, blue, or green. 
// The cost of painting each house with a certain color is different. 
// You have to paint all the houses such that no two adjacent houses have the same color.
// The cost of painting each house with a certain color is represented by an n x 3 cost matrix costs.
//     For example, costs[0][0] is the cost of painting house 0 with the color red; 
//     costs[1][2] is the cost of painting house 1 with color green, and so on...

// Return the minimum cost to paint all houses.

// Example 1:
// Input: costs = [[17,2,17],[16,16,5],[14,3,19]]
// Output: 10
// Explanation: Paint house 0 into blue, paint house 1 into green, paint house 2 into blue.
// Minimum cost: 2 + 5 + 3 = 10.

// Example 2:
// Input: costs = [[7,6,2]]
// Output: 2
 
// Constraints:
//     costs.length == n
//     costs[i].length == 3
//     1 <= n <= 100
//     1 <= costs[i][j] <= 20

import "fmt"

// func minCost(costs [][]int) int {
//     res := 0
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     for _,v := range costs {
//         res += min(v[0],min(v[1],v[2])) // 找每行最小的数值累加即可
//     }
//     return res
// }

func minCost(costs [][]int) int {
    r, b, g := 0, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < len(costs) ; i++ {
        // 需要粉刷所有的房子并且使其相邻的两个房子颜色不能相同 所以 r 只能取 b, g
        r, b, g = costs[i][0] + min(b, g), costs[i][1] + min(r, g), costs[i][2] + min(r, b)
    }
    return min(r, min(b, g))
}

func main() {
    // Explanation: Paint house 0 into blue, paint house 1 into green, paint house 2 into blue.
    // Minimum cost: 2 + 5 + 3 = 10.
    fmt.Println(minCost([][]int{{17,2,17},{16,16,5},{14,3,19}})) // 10
    fmt.Println(minCost([][]int{{7,6,2}})) // 2
}