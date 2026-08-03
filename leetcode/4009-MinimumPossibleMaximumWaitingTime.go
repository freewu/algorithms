package main

// 4009. Minimum Possible Maximum Waiting Time
// You are given an integer array demand, where demand[i] is the amount of fuel required by the ith car.

// You are also given an integer array fuel of length 2. 
// There are exactly two fuel dispensers, numbered 0 and 1, where fuel[j] is the initial amount of fuel available in dispenser j.

// Cars are allowed to start refueling in increasing index order. 
// Car 0 becomes allowed at time 0, and for each i > 0, car i becomes allowed exactly when car i - 1 starts refueling.

// The refueling process follows these rules:
//     1. Each dispenser can serve at most one car at a time.
//     2. A car may start refueling at any time at or after it becomes allowed.
//     3. A car can start on a dispenser only if the dispenser is free and has at least demand[i] fuel remaining.
//     4. If multiple free dispensers can serve the current car, you may choose any of them.
//     5. Refueling a car takes demand[i] seconds and reduces the remaining fuel in that dispenser by demand[i].
//     6. Once started, refueling cannot be interrupted.
//     7. When both dispensers are free, if neither has at least demand[i] fuel remaining, the process terminates and no further cars can be served.

// The waiting time of a car is the time between when it becomes allowed to start refueling and when it actually starts.

// Return the minimum possible value of the maximum waiting time among all served cars over all assignments that maximize the number of served cars. 
// If no car can be served, return -1.

// Example 1:
// Input: demand = [6,8,4,6,5], fuel = [16,13]
// Output: 6
// Explanation:
// Car | allowed at | Starts refueling at | Dispenser used | Remaining fuel before start (dispenser 0, dispenser 1) | Waiting time
// 0   | 0          | 0                   | 1              | (16, 13)                                               | 0
// 1   | 0          | 0                   | 1              | (16, 7)                                                | 0
// 2   | 0          | 0                   | 1              | (8, 7)                                                 | 6
// 3   | 0          | 6                   | 1              | (8, 3)                                                 | 2
// Car 4 becomes allowed at time 8, but when both dispensers are free, their remaining fuel is (2, 3), which is less than demand[4] = 5.
// Therefore, the process terminates. The maximum waiting time among served cars is 6.

// Example 2:
// Input: demand = [10,15], fuel = [12,17]
// Output: 0
// Explanation:
// At time 0, Car 0 becomes allowed and starts refuelling using dispenser 0.
// Car 1 becomes allowed at time 0 (when Car 0 starts) and immediately starts refuelling using dispenser 1.
// Both cars start without waiting, so the maximum waiting time is 0.

// Example 3:
// Input: demand = [10,5], fuel = [8,8]
// Output: -1
// Explanation:
// At time 0, Car 0 becomes allowed. However, neither dispenser has enough fuel to serve it, so the process terminates immediately.
// No car is served, so the answer is -1.

// Constraints:
//     1 <= demand.length <= 50
//     1 <= demand[i] <= 20
//     fuel.length == 2
//     1 <= fuel[i] <= 50

import "fmt"

func minMaxWaitingTime(demand []int, fuel []int) int {
    inf, mx := 1 << 61, 51 * 21 * 21
    encode := func(f1, t1, t2 int) int {
        return (f1 * 21 + t1) * 21 + t2
    }
    decode := func(state, fuelSum, usedFuel int) (int, int, int, int) {
        t2 := state % 21
        state /= 21
        t1 := state % 21
        f1 := state / 21
        f2 := fuelSum - usedFuel - f1
        return f1, f2, t1, t2
    }
    if demand[0] > max(fuel[0], fuel[1]) {
        return -1
    }
    dp := make([]int, mx)
    dpp := make([]int, mx)
    for i := 0; i < mx; i++ {
        dp[i] = inf
        dpp[i] = inf
    }
    curr := make([]int, 0)
    next := make([]int, 0)
    fuelSum := fuel[0] + fuel[1]
    usedFuel := 0
    initialState := encode(fuel[0], 0, 0)
    dp[initialState] = 0
    curr = append(curr, initialState)
    for _, d := range demand {
        for i := 0; i < mx; i++ {
            dpp[i] = inf
        }
        next = next[:0]
        for _, state := range curr {
            f1, f2, t1, t2 := decode(state, fuelSum, usedFuel)
            if f1 >= d {
                remaining := t2 - t1
                if remaining < 0 {
                    remaining = 0
                }
                nextState := encode(f1-d, d, remaining)
                nextWait := max(dp[state], t1)
                if dpp[nextState] == inf {
                    next = append(next, nextState)
                }
                if nextWait < dpp[nextState] {
                    dpp[nextState] = nextWait
                }
            }
            if f2 >= d {
                remaining := t1 - t2
                if remaining < 0 {
                    remaining = 0
                }
                nextState := encode(f1, remaining, d)
                nextWait := max(dp[state], t2)
                if dpp[nextState] == inf {
                    next = append(next, nextState)
                }
                if nextWait < dpp[nextState] {
                    dpp[nextState] = nextWait
                }
            }
        }
        if len(next) == 0 {
            break
        }
        usedFuel += d
        dp, dpp = dpp, dp
        curr, next = next, curr
    }
    res := inf
    for _, state := range curr {
        if dp[state] < res {
            res = dp[state]
        }
    }
    return res
}

func minMaxWaitingTime1(demand []int, fuel []int) int {
    const INF = 1 << 61
    n := len(demand)
    // Step 1: find maximum number of served cars (must be a prefix).
    // poss[s] = can assign prefix cars so that dispenser 0 uses exactly s fuel.
    poss := make([]bool, fuel[0]+1)
    poss[0] = true
    pref, kStar := 0, 0
    for i := 0; i < n; i++ {
        d := demand[i]
        np := make([]bool, fuel[0]+1)
        for s := 0; s <= fuel[0]; s++ {
            if poss[s] {
                np[s] = true
                if s+d <= fuel[0] {
                    np[s+d] = true
                }
            }
        }
        poss = np
        pref += d
        ok := false
        for s := 0; s <= fuel[0]; s++ {
            if poss[s] && pref-s <= fuel[1] {
                ok = true
                break
            }
        }
        if ok {
            kStar = i + 1
        } else {
            break
        }
    }
    if kStar == 0 {
        return -1
    }
    F0 := fuel[0]
    // dp[f0][last][r] = minimal possible max waiting time so far, where
    // f0 = fuel used on dispenser 0, last = dispenser of previous car,
    // r = clamped (availability time of other dispenser - start time of previous car).
    dp := make([][2][21]int, F0+1)
    for a := range dp {
        for l := 0; l < 2; l++ {
            for r := 0; r < 21; r++ {
                dp[a][l][r] = INF
            }
        }
    }
    if demand[0] <= fuel[0] {
        dp[demand[0]][0][0] = 0
    }
    if demand[0] <= fuel[1] {
        dp[0][1][0] = 0
    }
    prefUsed := demand[0]
    for i := 1; i < kStar; i++ {
        d := demand[i]
        pd := demand[i-1]
        ndp := make([][2][21]int, F0+1)
        for a := range ndp {
            for l := 0; l < 2; l++ {
                for r := 0; r < 21; r++ {
                    ndp[a][l][r] = INF
                }
            }
        }
        for f0 := 0; f0 <= F0; f0++ {
            f1 := prefUsed - f0
            for last := 0; last < 2; last++ {
                for r := 0; r < 21; r++ {
                    v := dp[f0][last][r]
                    if v >= INF {
                        continue
                    }
                    // Option A: same dispenser as previous car -> wait = pd
                    {
                        nv := v
                        if pd > nv {
                            nv = pd
                        }
                        nr := r - pd
                        if nr < 0 {
                            nr = 0
                        }
                        if last == 0 {
                            if f0+d <= fuel[0] && nv < ndp[f0+d][0][nr] {
                                ndp[f0+d][0][nr] = nv
                            }
                        } else {
                            if f1+d <= fuel[1] && nv < ndp[f0][1][nr] {
                                ndp[f0][1][nr] = nv
                            }
                        }
                    }
                    // Option B: other dispenser -> wait = r
                    {
                        nv := v
                        if r > nv {
                            nv = r
                        }
                        nr := pd - r
                        if nr < 0 {
                            nr = 0
                        }
                        if last == 0 {
                            if f1+d <= fuel[1] && nv < ndp[f0][1][nr] {
                                ndp[f0][1][nr] = nv
                            }
                        } else {
                            if f0+d <= fuel[0] && nv < ndp[f0+d][0][nr] {
                                ndp[f0+d][0][nr] = nv
                            }
                        }
                    }
                }
            }
        }
        dp = ndp
        prefUsed += d
    }
    res := INF
    for f0 := 0; f0 <= F0; f0++ {
        for l := 0; l < 2; l++ {
            for r := 0; r < 21; r++ {
                if dp[f0][l][r] < res {
                    res = dp[f0][l][r]
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: demand = [6,8,4,6,5], fuel = [16,13]
    // Output: 6
    // Explanation:
    // Car | allowed at | Starts refueling at | Dispenser used | Remaining fuel before start (dispenser 0, dispenser 1) | Waiting time
    // 0   | 0          | 0                   | 1              | (16, 13)                                               | 0
    // 1   | 0          | 0                   | 1              | (16, 7)                                                | 0
    // 2   | 0          | 0                   | 1              | (8, 7)                                                 | 6
    // 3   | 0          | 6                   | 1              | (8, 3)                                                 | 2
    // Car 4 becomes allowed at time 8, but when both dispensers are free, their remaining fuel is (2, 3), which is less than demand[4] = 5.
    // Therefore, the process terminates. The maximum waiting time among served cars is 6.
    fmt.Println(minMaxWaitingTime([]int{6,8,4,6,5}, []int{16,13})) // 6
    // Example 2:
    // Input: demand = [10,15], fuel = [12,17]
    // Output: 0
    // Explanation:
    // At time 0, Car 0 becomes allowed and starts refuelling using dispenser 0.
    // Car 1 becomes allowed at time 0 (when Car 0 starts) and immediately starts refuelling using dispenser 1.
    // Both cars start without waiting, so the maximum waiting time is 0.
    fmt.Println(minMaxWaitingTime([]int{10,15}, []int{12,17})) // 0
    // Example 3:
    // Input: demand = [10,5], fuel = [8,8]
    // Output: -1
    // Explanation:
    // At time 0, Car 0 becomes allowed. However, neither dispenser has enough fuel to serve it, so the process terminates immediately.
    // No car is served, so the answer is -1.
    fmt.Println(minMaxWaitingTime([]int{10,5}, []int{8,8})) // -1
    // Example 4:
    // Input: demand = [2,3,5], fuel = [1,8]
    // Output: 2
    fmt.Println(minMaxWaitingTime([]int{2,3,5}, []int{1,8})) // 2
    // Example 5:
    // Input: demand = [3,2,4,4], fuel = [4,5]
    // Output: 3
    fmt.Println(minMaxWaitingTime([]int{3,2,4,4}, []int{4,5})) // 3

    fmt.Println(minMaxWaitingTime([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minMaxWaitingTime([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 2
    fmt.Println(minMaxWaitingTime([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(minMaxWaitingTime([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(minMaxWaitingTime1([]int{6,8,4,6,5}, []int{16,13})) // 6
    fmt.Println(minMaxWaitingTime1([]int{10,15}, []int{12,17})) // 0
    fmt.Println(minMaxWaitingTime1([]int{10,5}, []int{8,8})) // -1
    fmt.Println(minMaxWaitingTime1([]int{2,3,5}, []int{1,8})) // 2
    fmt.Println(minMaxWaitingTime1([]int{3,2,4,4}, []int{4,5})) // 3
    fmt.Println(minMaxWaitingTime1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minMaxWaitingTime1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 2
    fmt.Println(minMaxWaitingTime1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(minMaxWaitingTime1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // 0
}