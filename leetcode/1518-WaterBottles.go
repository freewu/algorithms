package main

// 1518. Water Bottles
// There are numBottles water bottles that are initially full of water. 
// You can exchange numExchange empty water bottles from the market with one full water bottle.

// The operation of drinking a full water bottle turns it into an empty bottle.

// Given the two integers numBottles and numExchange, return the maximum number of water bottles you can drink.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/07/01/sample_1_1875.png" />
// Input: numBottles = 9, numExchange = 3
// Output: 13
// Explanation: You can exchange 3 empty bottles to get 1 full water bottle.
// Number of water bottles you can drink: 9 + 3 + 1 = 13.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/07/01/sample_2_1875.png" />
// Input: numBottles = 15, numExchange = 4
// Output: 19
// Explanation: You can exchange 4 empty bottles to get 1 full water bottle. 
// Number of water bottles you can drink: 15 + 3 + 1 = 19.

// Constraints:
//     1 <= numBottles <= 100
//     2 <= numExchange <= 100

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
    res, emptyBottles := numBottles, numBottles
    for emptyBottles >= numExchange {
        t := emptyBottles / numExchange // 空瓶可以换到的数量
        res += t
        emptyBottles = t + (emptyBottles % numExchange)  // 加上不兑换的空瓶
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/07/01/sample_1_1875.png" />
    // Input: numBottles = 9, numExchange = 3
    // Output: 13
    // Explanation: You can exchange 3 empty bottles to get 1 full water bottle.
    // Number of water bottles you can drink: 9 + 3 + 1 = 13.
    fmt.Println(numWaterBottles(9, 3)) // 13
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/07/01/sample_2_1875.png" />
    // Input: numBottles = 15, numExchange = 4
    // Output: 19
    // Explanation: You can exchange 4 empty bottles to get 1 full water bottle. 
    // Number of water bottles you can drink: 15 + 3 + 1 = 19.
    fmt.Println(numWaterBottles(15, 4)) // 19

    fmt.Println(numWaterBottles(1, 2)) // 1
    fmt.Println(numWaterBottles(100, 100)) // 101
    fmt.Println(numWaterBottles(1, 100)) // 1
    fmt.Println(numWaterBottles(100, 2)) // 199
}