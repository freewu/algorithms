package main

// 2511. Maximum Enemy Forts That Can Be Captured
// You are given a 0-indexed integer array forts of length n representing the positions of several forts. 
// forts[i] can be -1, 0, or 1 where:
//     1. -1 represents there is no fort at the ith position.
//     2. 0 indicates there is an enemy fort at the ith position.
//     3. 1 indicates the fort at the ith the position is under your command.

// Now you have decided to move your army from one of your forts at position i to an empty position j such that:
//     1. 0 <= i, j <= n - 1
//     2. The army travels over enemy forts only. 
//        Formally, for all k where min(i,j) < k < max(i,j), forts[k] == 0.

// While moving the army, all the enemy forts that come in the way are captured.

// Return the maximum number of enemy forts that can be captured. 
// In case it is impossible to move your army, or you do not have any fort under your command, return 0.

// Example 1:
// Input: forts = [1,0,0,-1,0,0,0,0,1]
// Output: 4
// Explanation:
// - Moving the army from position 0 to position 3 captures 2 enemy forts, at 1 and 2.
// - Moving the army from position 8 to position 3 captures 4 enemy forts.
// Since 4 is the maximum number of enemy forts that can be captured, we return 4.

// Example 2:
// Input: forts = [0,0,1,-1]
// Output: 0
// Explanation: Since no enemy fort can be captured, 0 is returned.

// Constraints:
//     1 <= forts.length <= 1000
//     -1 <= forts[i] <= 1

import "fmt"

func captureForts(forts []int) int {
    res, count, prev := 0, 0, -2
    for i := 0; i < len(forts); i++ {
        if forts[i] == -1 || forts[i] == 1 {
            if forts[i] == prev {
                count = 0
                continue
            }
            prev = forts[i]
            if count > res {
                res = count
            }
            count = 0
            continue
        }
        if forts[i] == 0 && (prev == -1 || prev == 1) {
            count++
        }
    }
    return res
}

func captureForts1(forts []int) int {
    res, start, local := 0, 0, 0
    for _, v := range forts {
        if v == 0 {
            local++
            continue
        }
        if v == start * -1 && local > res { 
            res = local 
        }
        local, start = 0, v
    }
    return res
}

func captureForts2(forts []int) (ans int) {
    res, p := 0, -1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range forts {
        if v != 0 {
            if p >= 0 && forts[i] != forts[p] {
                res = max(res, i - p - 1)
            }
            p = i
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: forts = [1,0,0,-1,0,0,0,0,1]
    // Output: 4
    // Explanation:
    // - Moving the army from position 0 to position 3 captures 2 enemy forts, at 1 and 2.
    // - Moving the army from position 8 to position 3 captures 4 enemy forts.
    // Since 4 is the maximum number of enemy forts that can be captured, we return 4.
    fmt.Println(captureForts([]int{1,0,0,-1,0,0,0,0,1})) // 4
    // Example 2:
    // Input: forts = [0,0,1,-1]
    // Output: 0
    // Explanation: Since no enemy fort can be captured, 0 is returned.
    fmt.Println(captureForts([]int{0,0,1,-1})) // 0

    fmt.Println(captureForts1([]int{1,0,0,-1,0,0,0,0,1})) // 4
    fmt.Println(captureForts1([]int{0,0,1,-1})) // 0

    fmt.Println(captureForts2([]int{1,0,0,-1,0,0,0,0,1})) // 4
    fmt.Println(captureForts2([]int{0,0,1,-1})) // 0
}