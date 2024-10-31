package main

// 3259. Maximum Energy Boost From Two Drinks
// You are given two integer arrays energyDrinkA and energyDrinkB of the same length n by a futuristic sports scientist. 
// These arrays represent the energy boosts per hour provided by two different energy drinks, A and B, respectively.

// You want to maximize your total energy boost by drinking one energy drink per hour. 
// However, if you want to switch from consuming one energy drink to the other, 
// you need to wait for one hour to cleanse your system (meaning you won't get any energy boost in that hour).

// Return the maximum total energy boost you can gain in the next n hours.

// Note that you can start consuming either of the two energy drinks.

// Example 1:
// Input: energyDrinkA = [1,3,1], energyDrinkB = [3,1,1]
// Output: 5
// Explanation:
// To gain an energy boost of 5, drink only the energy drink A (or only B).

// Example 2:
// Input: energyDrinkA = [4,1,1], energyDrinkB = [1,1,3]
// Output: 7
// Explanation:
// To gain an energy boost of 7:
//     Drink the energy drink A for the first hour.
//     Switch to the energy drink B and we lose the energy boost of the second hour.
//     Gain the energy boost of the drink B in the third hour.

// Constraints:
//     n == energyDrinkA.length == energyDrinkB.length
//     3 <= n <= 10^5
//     1 <= energyDrinkA[i], energyDrinkB[i] <= 10^

import "fmt"

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
    n := len(energyDrinkA)
    if n == 0 { return 0 }
    dp1, dp2 := make([]int, n), make([]int, n)
    dp1[0], dp2[0] = energyDrinkA[0], energyDrinkB[0]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        // if we don't swith the iteration
        dp1[i], dp2[i] = (dp1[i-1] + energyDrinkA[i]), (dp2[i-1] + energyDrinkB[i])
        if i > 1 { // if we want to switch to the another drink 
            dp1[i], dp2[i] = max(dp1[i], dp2[i-2] + energyDrinkA[i]), max(dp2[i], dp1[i-2] + energyDrinkB[i])
        }
    }
    return int64(max(dp1[n-1], dp2[n-1]))
}

func main() {
    // Example 1:
    // Input: energyDrinkA = [1,3,1], energyDrinkB = [3,1,1]
    // Output: 5
    // Explanation:
    // To gain an energy boost of 5, drink only the energy drink A (or only B).
    fmt.Println(maxEnergyBoost([]int{1,3,1}, []int{3,1,1})) // 5
    // Example 2:
    // Input: energyDrinkA = [4,1,1], energyDrinkB = [1,1,3]
    // Output: 7
    // Explanation:
    // To gain an energy boost of 7:
    //     Drink the energy drink A for the first hour.
    //     Switch to the energy drink B and we lose the energy boost of the second hour.
    //     Gain the energy boost of the drink B in the third hour.
    fmt.Println(maxEnergyBoost([]int{4,1,1}, []int{1,1,3})) // 7
}