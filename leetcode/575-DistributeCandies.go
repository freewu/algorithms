package main

// 575. Distribute Candies
// Alice has n candies, where the ith candy is of type candyType[i]. 
// Alice noticed that she started to gain weight, so she visited a doctor.

// The doctor advised Alice to only eat n / 2 of the candies she has (n is always even). 
// Alice likes her candies very much, and she wants to eat the maximum number of different types of candies while still following the doctor's advice.

// Given the integer array candyType of length n, return the maximum number of different types of candies she can eat if she only eats n / 2 of them.

// Example 1:
// Input: candyType = [1,1,2,2,3,3]
// Output: 3
// Explanation: Alice can only eat 6 / 2 = 3 candies. Since there are only 3 types, she can eat one of each type.

// Example 2:
// Input: candyType = [1,1,2,3]
// Output: 2
// Explanation: Alice can only eat 4 / 2 = 2 candies. Whether she eats types [1,2], [1,3], or [2,3], she still can only eat 2 different types.

// Example 3:
// Input: candyType = [6,6,6,6]
// Output: 1
// Explanation: Alice can only eat 4 / 2 = 2 candies. Even though she can eat 2 candies, she only has 1 type.
 
// Constraints:
//     n == candyType.length
//     2 <= n <= 10^4
//     n is even.
//     -10^5 <= candyType[i] <= 10^5

import "fmt"

func distributeCandies(candyType []int) int {
    m := make(map[int]int)
    for _, v := range candyType {
        m[v]++
    }
    mid := len(candyType) >> 1 // 只吃一半的糖
    candyTypeCount := len(m) // 糖果种类
    if mid >= candyTypeCount {
        return candyTypeCount
    }
    return mid
}

func distributeCandies1(candyType []int) int {
    mp, length := make(map[int]int), len(candyType) / 2
    for _, v := range candyType {
        mp[v]++
        if len(mp) > length {
            return length
        }
    }
    return len(mp)
}

func distributeCandies2(candyType []int) int {
    count, m := 0, make(map[int] bool, len(candyType) / 2)
    for _, v := range candyType {
        if !m[v] {
            m[v] = true
            count++
            if count == len(candyType) / 2 {
                return count
            }
        }
    }
    return count
}

func main() {
    // Example 1:
    // Input: candyType = [1,1,2,2,3,3]
    // Output: 3
    // Explanation: Alice can only eat 6 / 2 = 3 candies. Since there are only 3 types, she can eat one of each type.
    fmt.Println(distributeCandies([]int{1,1,2,2,3,3})) // 3
    // Example 2:
    // Input: candyType = [1,1,2,3]
    // Output: 2
    // Explanation: Alice can only eat 4 / 2 = 2 candies. Whether she eats types [1,2], [1,3], or [2,3], she still can only eat 2 different types.
    fmt.Println(distributeCandies([]int{1,1,2,3})) // 2
    // Example 3:
    // Input: candyType = [6,6,6,6]
    // Output: 1
    // Explanation: Alice can only eat 4 / 2 = 2 candies. Even though she can eat 2 candies, she only has 1 type.
    fmt.Println(distributeCandies([]int{6,6,6,6})) // 1

    fmt.Println(distributeCandies1([]int{1,1,2,2,3,3})) // 3
    fmt.Println(distributeCandies1([]int{1,1,2,3})) // 2
    fmt.Println(distributeCandies1([]int{6,6,6,6})) // 1

    fmt.Println(distributeCandies2([]int{1,1,2,2,3,3})) // 3
    fmt.Println(distributeCandies2([]int{1,1,2,3})) // 2
    fmt.Println(distributeCandies2([]int{6,6,6,6})) // 1
}