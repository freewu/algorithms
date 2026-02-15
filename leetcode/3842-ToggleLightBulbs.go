package main

// 3842. Toggle Light Bulbs
// You are given an array bulbs of integers between 1 and 100.

// There are 100 light bulbs numbered from 1 to 100. All of them are switched off initially.

// For each element bulbs[i] in the array bulbs:
//     1. If the bulbs[i]th light bulb is currently off, switch it on.
//     2. Otherwise, switch it off.

// Return the list of integers denoting the light bulbs that are on in the end, sorted in ascending order. 
// If no bulb is on, return an empty list.

// Example 1:
// Input: bulbs = [10,30,20,10]
// Output: [20,30]
// Explanation:
// The bulbs[0] = 10th light bulb is currently off. We switch it on.
// The bulbs[1] = 30th light bulb is currently off. We switch it on.
// The bulbs[2] = 20th light bulb is currently off. We switch it on.
// The bulbs[3] = 10th light bulb is currently on. We switch it off.
// In the end, the 20th and the 30th light bulbs are on.

// Example 2:
// Input: bulbs = [100,100]
// Output: []
// Explanation:
// The bulbs[0] = 100th light bulb is currently off. We switch it on.
// The bulbs[1] = 100th light bulb is currently on. We switch it off.
// In the end, no light bulb is on.

// Constraints:
//     1 <= bulbs.length <= 100
//     1 <= bulbs[i] <= 100

import "fmt"

func toggleLightBulbs(bulbs []int) []int {
    res, mp := make([]int, 0), make([]bool, 101)
    for _, v := range bulbs {
        mp[v] = !mp[v]
    }
    for i := 1; i <= 100; i++ {
        if mp[i] {
            res = append(res, i)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: bulbs = [10,30,20,10]
    // Output: [20,30]
    // Explanation:
    // The bulbs[0] = 10th light bulb is currently off. We switch it on.
    // The bulbs[1] = 30th light bulb is currently off. We switch it on.
    // The bulbs[2] = 20th light bulb is currently off. We switch it on.
    // The bulbs[3] = 10th light bulb is currently on. We switch it off.
    // In the end, the 20th and the 30th light bulbs are on.
    fmt.Println(toggleLightBulbs([]int{10,30,20,10})) // [20 30]
    // Example 2:
    // Input: bulbs = [100,100]
    // Output: []
    // Explanation:
    // The bulbs[0] = 100th light bulb is currently off. We switch it on.
    // The bulbs[1] = 100th light bulb is currently on. We switch it off.
    // In the end, no light bulb is on.
    fmt.Println(toggleLightBulbs([]int{100,100})) // []

    fmt.Println(toggleLightBulbs([]int{1,2,3,4,5,6,7,8,9})) // [1 2 3 4 5 6 7 8 9]
    fmt.Println(toggleLightBulbs([]int{9,8,7,6,5,4,3,2,1})) // [1 2 3 4 5 6 7 8 9]
}