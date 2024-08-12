package main 

// 818. Race Car
// Your car starts at position 0 and speed +1 on an infinite number line. 
// Your car can go into negative positions. 
// Your car drives automatically according to a sequence of instructions 'A' (accelerate) and 'R' (reverse):
//     When you get an instruction 'A', your car does the following:
//         position += speed
//         speed *= 2
//     When you get an instruction 'R', your car does the following:
//         If your speed is positive then speed = -1
//         otherwise speed = 1
//     Your position stays the same.

// For example, after commands "AAR", your car goes to positions 0 --> 1 --> 3 --> 3, and your speed goes to 1 --> 2 --> 4 --> -1.

// Given a target position target, return the length of the shortest sequence of instructions to get there.

// Example 1:
// Input: target = 3
// Output: 2
// Explanation: 
// The shortest instruction sequence is "AA".
// Your position goes from 0 --> 1 --> 3.

// Example 2:
// Input: target = 6
// Output: 5
// Explanation: 
// The shortest instruction sequence is "AAARA".
// Your position goes from 0 --> 1 --> 3 --> 7 --> 7 --> 6.

// Constraints:
//     1 <= target <= 10^4

import "fmt"
import "math"

func racecar(target int) int {
    dp := make([]int, target + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(target int) int
    dfs = func(target int) int {
        if dp[target] > 0 { // if we processed target, then dp[target] != 0
            return dp[target]
        }
        // quick way to find the shortest straight route longer than target
        n := int(math.Log2(float64(target))) + 1 
        noReverse := 1 << n - 1
        if noReverse == target {
            dp[target] = n
        } else {
            // cost to get here n accelerates + 1 reverse
            dp[target] = dfs(noReverse - target) + n + 1
            // get here faster by getting as close to the target (n-1) accelerates
            // then try reverse, and accelerate again. Reverse up to n-1 which would get back to 0
            maxNoReverse := 1 << (n-1)
            for m := 0; m < n-1; m++ {
                reverses := 1 << m
                // n - 1 accelerates, m reverses + instruction to reverse and accelerate
                moves := n - 1 + m + 2
                nextTarget := target - maxNoReverse + reverses
                dp[target] = min(dp[target], dfs(nextTarget) + moves)
            }
        }
        return dp[target]
    }
    return dfs(target)
}

func main() {
    // Example 1:
    // Input: target = 3
    // Output: 2
    // Explanation: 
    // The shortest instruction sequence is "AA".
    // Your position goes from 0 --> 1 --> 3.
    fmt.Println(racecar(3)) // 2
    // Example 2:
    // Input: target = 6
    // Output: 5
    // Explanation: 
    // The shortest instruction sequence is "AAARA".
    // Your position goes from 0 --> 1 --> 3 --> 7 --> 7 --> 6.
    fmt.Println(racecar(6)) // 5
}