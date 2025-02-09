package main

// 3186. Maximum Total Damage With Spell Casting
// A magician has various spells.

// You are given an array power, where each element represents the damage of a spell. 
// Multiple spells can have the same damage value.

// It is a known fact that if a magician decides to cast a spell with a damage of power[i], 
// they cannot cast any spell with a damage of power[i] - 2, power[i] - 1, power[i] + 1, or power[i] + 2.

// Each spell can be cast only once.

// Return the maximum possible total damage that a magician can cast.

// Example 1:
// Input: power = [1,1,3,4]
// Output: 6
// Explanation:
// The maximum possible damage of 6 is produced by casting spells 0, 1, 3 with damage 1, 1, 4.

// Example 2:
// Input: power = [7,1,6,6]
// Output: 13
// Explanation:
// The maximum possible damage of 13 is produced by casting spells 1, 2, 3 with damage 1, 6, 6.

// Constraints:
//     1 <= power.length <= 10^5
//     1 <= power[i] <= 10^9

import "fmt"
import "sort"

func maximumTotalDamage(power []int) int64 {
    sort.Ints(power)
    dp := make([]int, len(power))
    j, k := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, p := range power {
        d := p
        for power[k] + 2 < p {
            k++
        }
        if k > 0 {
            d += dp[k - 1]
        }
        if power[j] != p {
            j = i
        }
        d += ((i - j) * p)
        if i > 0 {
            dp[i] = max(dp[i - 1], d)
        } else {
            dp[i] = d
        }
    }
    return int64(dp[len(dp) - 1])
}

func maximumTotalDamage1(power []int) int64 {
    sort.Ints(power)
    p2, p3, p1v, p2v, p3v, count := -3, -3, 0, 0, 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, p := range power {
        count++
        if i == len(power)-1 || power[i] != power[i+1] {
            y := 0
            if p != p3+1 && p != p3+2 {
                y = p3v
            } else if p != p2+2 {
                y = p2v
            } else {
                y = p1v
            }
            y = max(y + p * count, p3v)
            p2, p3, p1v, p2v, p3v = p3, p, p2v, p3v, y
            count = 0
        }
    }
    return int64(p3v)
}

func main() {
    // Example 1:
    // Input: power = [1,1,3,4]
    // Output: 6
    // Explanation:
    // The maximum possible damage of 6 is produced by casting spells 0, 1, 3 with damage 1, 1, 4.
    fmt.Println(maximumTotalDamage([]int{1,1,3,4})) // 6
    // Example 2:
    // Input: power = [7,1,6,6]
    // Output: 13
    // Explanation:
    // The maximum possible damage of 13 is produced by casting spells 1, 2, 3 with damage 1, 6, 6.
    fmt.Println(maximumTotalDamage([]int{7,1,6,6})) // 13

    fmt.Println(maximumTotalDamage([]int{1,2,3,4,5,6,7,8,9})) // 18
    fmt.Println(maximumTotalDamage([]int{9,8,7,6,5,4,3,2,1})) // 18

    fmt.Println(maximumTotalDamage1([]int{1,1,3,4})) // 6
    fmt.Println(maximumTotalDamage1([]int{7,1,6,6})) // 13
    fmt.Println(maximumTotalDamage1([]int{1,2,3,4,5,6,7,8,9})) // 18
    fmt.Println(maximumTotalDamage1([]int{9,8,7,6,5,4,3,2,1})) // 18
}