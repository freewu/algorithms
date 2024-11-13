package main

// 2403. Minimum Time to Kill All Monsters
// You are given an integer array power where power[i] is the power of the ith monster.

// You start with 0 mana points, 
// and each day you increase your mana points by gain where gain initially is equal to 1.

// Each day, after gaining gain mana, 
// you can defeat a monster if your mana points are greater than or equal to the power of that monster. 
// When you defeat a monster:
//     your mana points will be reset to 0, and
//     the value of gain increases by 1.

// Return the minimum number of days needed to defeat all the monsters.

// Example 1:
// Input: power = [3,1,4]
// Output: 4
// Explanation: The optimal way to beat all the monsters is to:
// - Day 1: Gain 1 mana point to get a total of 1 mana point. Spend all mana points to kill the 2nd monster.
// - Day 2: Gain 2 mana points to get a total of 2 mana points.
// - Day 3: Gain 2 mana points to get a total of 4 mana points. Spend all mana points to kill the 3rd monster.
// - Day 4: Gain 3 mana points to get a total of 3 mana points. Spend all mana points to kill the 1st monster.
// It can be proven that 4 is the minimum number of days needed. 

// Example 2:
// Input: power = [1,1,4]
// Output: 4
// Explanation: The optimal way to beat all the monsters is to:
// - Day 1: Gain 1 mana point to get a total of 1 mana point. Spend all mana points to kill the 1st monster.
// - Day 2: Gain 2 mana points to get a total of 2 mana points. Spend all mana points to kill the 2nd monster.
// - Day 3: Gain 3 mana points to get a total of 3 mana points.
// - Day 4: Gain 3 mana points to get a total of 6 mana points. Spend all mana points to kill the 3rd monster.
// It can be proven that 4 is the minimum number of days needed. 

// Example 3:
// Input: power = [1,2,4,9]
// Output: 6
// Explanation: The optimal way to beat all the monsters is to:
// - Day 1: Gain 1 mana point to get a total of 1 mana point. Spend all mana points to kill the 1st monster.
// - Day 2: Gain 2 mana points to get a total of 2 mana points. Spend all mana points to kill the 2nd monster.
// - Day 3: Gain 3 mana points to get a total of 3 mana points.
// - Day 4: Gain 3 mana points to get a total of 6 mana points.
// - Day 5: Gain 3 mana points to get a total of 9 mana points. Spend all mana points to kill the 4th monster.
// - Day 6: Gain 4 mana points to get a total of 4 mana points. Spend all mana points to kill the 3rd monster.
// It can be proven that 6 is the minimum number of days needed.

// Constraints:
//     1 <= power.length <= 17
//     1 <= power[i] <= 10^9

import "fmt"

func minimumTime(power []int) int64 {
    n := len(power)
    cache := map[int]int64{ 0: 0 }
    bitCount := func(n int) int {
        count := 0
        for n > 0 {
            count++
            n &= n - 1
        }
        return count
    }
    var dp func(mask int) int64
    dp = func(mask int) int64 {
        if v, ok := cache[mask]; ok { return v }
        res, bits := int64(1 << 42), bitCount(mask)
        for i := 0; i < n; i++ {
            if (1 << i) & mask > 0 {
                t := dp(mask ^ (1 << i))
                t += int64(power[i] / bits)
                if power[i] % bits != 0 {
                    t++
                }
                if t < res {
                    res = t
                }
            }
        }
        cache[mask] = res
        return res
    }
    return dp((1 << n) - 1)
}

func main() {
    // Example 1:
    // Input: power = [3,1,4]
    // Output: 4
    // Explanation: The optimal way to beat all the monsters is to:
    // - Day 1: Gain 1 mana point to get a total of 1 mana point. Spend all mana points to kill the 2nd monster.
    // - Day 2: Gain 2 mana points to get a total of 2 mana points.
    // - Day 3: Gain 2 mana points to get a total of 4 mana points. Spend all mana points to kill the 3rd monster.
    // - Day 4: Gain 3 mana points to get a total of 3 mana points. Spend all mana points to kill the 1st monster.
    // It can be proven that 4 is the minimum number of days needed. 
    fmt.Println(minimumTime([]int{3,1,4})) // 4
    // Example 2:
    // Input: power = [1,1,4]
    // Output: 4
    // Explanation: The optimal way to beat all the monsters is to:
    // - Day 1: Gain 1 mana point to get a total of 1 mana point. Spend all mana points to kill the 1st monster.
    // - Day 2: Gain 2 mana points to get a total of 2 mana points. Spend all mana points to kill the 2nd monster.
    // - Day 3: Gain 3 mana points to get a total of 3 mana points.
    // - Day 4: Gain 3 mana points to get a total of 6 mana points. Spend all mana points to kill the 3rd monster.
    // It can be proven that 4 is the minimum number of days needed. 
    fmt.Println(minimumTime([]int{1,1,4})) // 4
    // Example 3:
    // Input: power = [1,2,4,9]
    // Output: 6
    // Explanation: The optimal way to beat all the monsters is to:
    // - Day 1: Gain 1 mana point to get a total of 1 mana point. Spend all mana points to kill the 1st monster.
    // - Day 2: Gain 2 mana points to get a total of 2 mana points. Spend all mana points to kill the 2nd monster.
    // - Day 3: Gain 3 mana points to get a total of 3 mana points.
    // - Day 4: Gain 3 mana points to get a total of 6 mana points.
    // - Day 5: Gain 3 mana points to get a total of 9 mana points. Spend all mana points to kill the 4th monster.
    // - Day 6: Gain 4 mana points to get a total of 4 mana points. Spend all mana points to kill the 3rd monster.
    // It can be proven that 6 is the minimum number of days needed.
    fmt.Println(minimumTime([]int{1,2,4,9})) // 6

    fmt.Println(minimumTime([]int{1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000,1000000000})) // 3439552527
}