package main

// 1388. Pizza With 3n Slices
// There is a pizza with 3n slices of varying size, you and your friends will take slices of pizza as follows:
//     You will pick any pizza slice.
//     Your friend Alice will pick the next slice in the anti-clockwise direction of your pick.
//     Your friend Bob will pick the next slice in the clockwise direction of your pick.
//     Repeat until there are no more slices of pizzas.

// Given an integer array slices that represent the sizes of the pizza slices in a clockwise direction, 
// return the maximum possible sum of slice sizes that you can pick.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/02/18/sample_3_1723.png" />
// Input: slices = [1,2,3,4,5,6]
// Output: 10
// Explanation: Pick pizza slice of size 4, Alice and Bob will pick slices with size 3 and 5 respectively. Then Pick slices with size 6, finally Alice and Bob will pick slice of size 2 and 1 respectively. Total = 4 + 6.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/02/18/sample_4_1723.png" />
// Input: slices = [8,9,8,6,1,1]
// Output: 16
// Explanation: Pick pizza slice of size 8 in each turn. If you pick slice with size 9 your partners will pick slices of size 8.
 
// Constraints:
//     3 * n == slices.length
//     1 <= slices.length <= 500
//     1 <= slices[i] <= 1000

import "fmt"

// dp
func maxSizeSlices(slices []int) int {
    n := len(slices)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    helper := func(slices []int, rounds, start, end int) int {
        dp := make([][2]int, rounds + 1)
        dp[1][1] = slices[start]
        for i := start + 1; i <= end; i++ {
            for j := rounds; j > 0; j-- {
                dp[j][0] = max(dp[j][0], dp[j][1])
                dp[j][1] = dp[j-1][0] + slices[i]
            }
        }
        return max(dp[rounds][0], dp[rounds][1])
    }
    return max(helper(slices, n/3, 0, n - 2), helper(slices, n/3, 1, n - 1))
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/02/18/sample_3_1723.png" />
    // Input: slices = [1,2,3,4,5,6]
    // Output: 10
    // Explanation: Pick pizza slice of size 4, Alice and Bob will pick slices with size 3 and 5 respectively. Then Pick slices with size 6, finally Alice and Bob will pick slice of size 2 and 1 respectively. Total = 4 + 6.
    fmt.Println(maxSizeSlices([]int{1,2,3,4,5,6})) // 10
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/02/18/sample_4_1723.png" />
    // Input: slices = [8,9,8,6,1,1]
    // Output: 16
    // Explanation: Pick pizza slice of size 8 in each turn. If you pick slice with size 9 your partners will pick slices of size 8.
    fmt.Println(maxSizeSlices([]int{8,9,8,6,1,1})) // 16
}