package main

// 3100. Water Bottles II
// You are given two integers numBottles and numExchange.

// numBottles represents the number of full water bottles that you initially have. 
// In one operation, you can perform one of the following operations:
//     1. Drink any number of full water bottles turning them into empty bottles.
//     2. Exchange numExchange empty bottles with one full water bottle. 
//        Then, increase numExchange by one.

// Note that you cannot exchange multiple batches of empty bottles for the same value of numExchange. 
// For example, if numBottles == 3 and numExchange == 1, you cannot exchange 3 empty water bottles for 3 full bottles.

// Return the maximum number of water bottles you can drink.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/01/28/exampleone1.png" />
// Input: numBottles = 13, numExchange = 6
// Output: 15
// Explanation: The table above shows the number of full water bottles, empty water bottles, the value of numExchange, and the number of bottles drunk.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/01/28/example231.png" />
// Input: numBottles = 10, numExchange = 3
// Output: 13
// Explanation: The table above shows the number of full water bottles, empty water bottles, the value of numExchange, and the number of bottles drunk.

// Constraints:
//     1 <= numBottles <= 100
//     1 <= numExchange <= 100

import "fmt"

func maxBottlesDrunk(numBottles int, numExchange int) int {
    res, sum := numBottles, numBottles
    for numExchange <= sum {
        sum -= numExchange
        sum++;
        numExchange++
        res++
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/01/28/exampleone1.png" />
    // Input: numBottles = 13, numExchange = 6
    // Output: 15
    // Explanation: The table above shows the number of full water bottles, empty water bottles, the value of numExchange, and the number of bottles drunk.
    fmt.Println(maxBottlesDrunk(13,6)) // 15
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/01/28/example231.png" />
    // Input: numBottles = 10, numExchange = 3
    // Output: 13
    // Explanation: The table above shows the number of full water bottles, empty water bottles, the value of numExchange, and the number of bottles drunk.
    fmt.Println(maxBottlesDrunk(10,3)) // 13

    fmt.Println(maxBottlesDrunk(1,1)) // 2
    fmt.Println(maxBottlesDrunk(100,100)) // 101
    fmt.Println(maxBottlesDrunk(1,100)) // 1
    fmt.Println(maxBottlesDrunk(100,1)) // 114
}