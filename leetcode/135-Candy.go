package main

// 135. Candy
// There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.
// You are giving candies to these children subjected to the following requirements:
//     Each child must have at least one candy.
//     Children with a higher rating get more candies than their neighbors.

// Return the minimum number of candies you need to have to distribute the candies to the children.

 
// Example 1:
// Input: ratings = [1,0,2]
// Output: 5
// Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

// Example 2:
// Input: ratings = [1,2,2]
// Output: 4
// Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
// The third child gets 1 candy because it satisfies the above two conditions.
 
// Constraints:
//     n == ratings.length
//     1 <= n <= 2 * 10^4
//     0 <= ratings[i] <= 2 * 10^4

import "fmt"

func candy(ratings []int) int {
    candies := make([]int, len(ratings)) // 声明一个数组来保存发糖数
    for i := 1; i < len(ratings); i++ {
        if ratings[i] > ratings[i-1] { // 如果当前孩子评分 大于 前面一个孩子评分  当前评分 + 1
            candies[i] += candies[i-1] + 1 // 评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果
        }
    }
    // fmt.Printf("after round 1 %v\n",candies)
    for i := len(ratings) - 2; i >= 0; i-- { // 从后向前推进
        if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
            candies[i] = candies[i+1] + 1 // 评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果
        }
    }
    // fmt.Printf("after round 2 %v\n",candies)
    res := 0
    for _, candy := range candies {
        res += candy + 1 // 每个孩子至少分配到 1 个糖果
    }
    return res
}

// dp
func candy1(ratings []int) int {
    res, dp := 0, make([]int, len(ratings))
    for i := 0; i < len(ratings); i++ {
        dp[i] = 1
    }
    for i := 0; i < len(ratings)-1; i++ {
        if ratings[i+1] > ratings[i] {
            dp[i+1] = dp[i] + 1
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j := len(ratings) - 1; j > 0; j-- {
        if ratings[j-1] > ratings[j] {
            dp[j-1] = max(dp[j] + 1, dp[j-1])
        }
    }
    for _, v := range dp {
        res += v
    }
    return res
}

func main() {
    // Example 1:
    // Input: ratings = [1,0,2]
    // Output: 5
    // Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
    fmt.Printf("candy([]int{ 1,0,2 }) = %v\n",candy([]int{ 1,0,2 })) // 5
    // Example 2:
    // Input: ratings = [1,2,2]
    // Output: 4
    // Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
    // The third child gets 1 candy because it satisfies the above two conditions.
    fmt.Printf("candy([]int{ 1,2,2 }) = %v\n",candy([]int{ 1,2,2 })) // 4
    fmt.Printf("candy([]int{ 1,2,3,4,5,6,7,8,9 }) = %v\n",candy([]int{ 1,2,3,4,5,6,7,8,9 })) // 45
    fmt.Printf("candy([]int{ 9,8,7,6,5,4,3,2,1 }) = %v\n",candy([]int{ 9,8,7,6,5,4,3,2,1 })) // 45

    fmt.Printf("candy1([]int{ 1,0,2 }) = %v\n",candy1([]int{ 1,0,2 })) // 5
    fmt.Printf("candy1([]int{ 1,2,2 }) = %v\n",candy1([]int{ 1,2,2 })) // 4
    fmt.Printf("candy1([]int{ 1,2,3,4,5,6,7,8,9 }) = %v\n",candy([]int{ 1,2,3,4,5,6,7,8,9 })) // 45
    fmt.Printf("candy1([]int{ 9,8,7,6,5,4,3,2,1 }) = %v\n",candy([]int{ 9,8,7,6,5,4,3,2,1 })) // 45
}
