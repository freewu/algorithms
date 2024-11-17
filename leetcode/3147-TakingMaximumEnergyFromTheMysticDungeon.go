package main

// 3147. Taking Maximum Energy From the Mystic Dungeon
// In a mystic dungeon, n magicians are standing in a line. 
// Each magician has an attribute that gives you energy. 
// Some magicians can give you negative energy, which means taking energy from you.

// You have been cursed in such a way that after absorbing energy from magician i, you will be instantly transported to magician (i + k). 
// This process will be repeated until you reach the magician where (i + k) does not exist.

// In other words, you will choose a starting point and then teleport with k jumps until you reach the end of the magicians' sequence, absorbing all the energy during the journey.

// You are given an array energy and an integer k. Return the maximum possible energy you can gain.

// Example 1:
// Input:  energy = [5,2,-10,-5,1], k = 3
// Output: 3
// Explanation: We can gain a total energy of 3 by starting from magician 1 absorbing 2 + 1 = 3.

// Example 2:
// Input: energy = [-2,-3,-1], k = 2
// Output: -1
// Explanation: We can gain a total energy of -1 by starting from magician 2.

// Constraints:
//     1 <= energy.length <= 10^5
//     -1000 <= energy[i] <= 1000
//     1 <= k <= energy.length - 1

import "fmt"

func maximumEnergy(energy []int, k int) int {
    res, count, n := -1 << 31, 0, len(energy)
    arr := make([]int, n)
    for i := n - 1; i >= 0; i-- {
        if count >= k {
            arr[i] = energy[i] + arr[i + k]
        } else {
            arr[i] = energy[i]
        }
        count++
    }
    for _, v := range arr {
        if v > res {
            res = v
        }
    }
    return res
}

func maximumEnergy1(energy []int, k int) int {
    n := len(energy)
    res := energy[n - 1]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - k; i < n; i++ {
        sum := 0
        for j := i; j >= 0; j -= k {
            sum += energy[j]
            res = max(res, sum)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input:  energy = [5,2,-10,-5,1], k = 3
    // Output: 3
    // Explanation: We can gain a total energy of 3 by starting from magician 1 absorbing 2 + 1 = 3.
    fmt.Println(maximumEnergy([]int{5,2,-10,-5,1}, 3)) // 3
    // Example 2:
    // Input: energy = [-2,-3,-1], k = 2
    // Output: -1
    // Explanation: We can gain a total energy of -1 by starting from magician 2.
    fmt.Println(maximumEnergy([]int{-2,-3,-1}, 2)) // -1

    fmt.Println(maximumEnergy1([]int{5,2,-10,-5,1}, 3)) // 3
    fmt.Println(maximumEnergy1([]int{-2,-3,-1}, 2)) // -1
}