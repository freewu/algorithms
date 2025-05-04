package main

// 3538. Merge Operations for Minimum Travel Time
// You are given a straight road of length l km, an integer n, an integer k, and two integer arrays, position and time, each of length n.

// The array position lists the positions (in km) of signs in strictly increasing order (with position[0] = 0 and position[n - 1] = l).

// Each time[i] represents the time (in minutes) required to travel 1 km between position[i] and position[i + 1].

// You must perform exactly k merge operations. 
// In one merge, you can choose any two adjacent signs at indices i and i + 1 (with i > 0 and i + 1 < n) and:
//     1. Update the sign at index i + 1 so that its time becomes time[i] + time[i + 1].
//     2. Remove the sign at index i.

// Return the minimum total travel time (in minutes) to travel from 0 to l after exactly k merges.

// Example 1:
// Input: l = 10, n = 4, k = 1, position = [0,3,8,10], time = [5,8,3,6]
// Output: 62
// Explanation:
// Merge the signs at indices 1 and 2. Remove the sign at index 1, and change the time at index 2 to 8 + 3 = 11.
// After the merge:
// position array: [0, 8, 10]
// time array: [5, 11, 6]
// Segment	Distance (km)	Time per km (min)	Segment Travel Time (min)
// 0 → 8	8	5	8 × 5 = 40
// 8 → 10	2	11	2 × 11 = 22
// Total Travel Time: 40 + 22 = 62, which is the minimum possible time after exactly 1 merge.

// Example 2:
// Input: l = 5, n = 5, k = 1, position = [0,1,2,3,5], time = [8,3,9,3,3]
// Output: 34
// Explanation:
// Merge the signs at indices 1 and 2. Remove the sign at index 1, and change the time at index 2 to 3 + 9 = 12.
// After the merge:
// position array: [0, 2, 3, 5]
// time array: [8, 12, 3, 3]
// Segment	Distance (km)	Time per km (min)	Segment Travel Time (min)
// 0 → 2	2	8	2 × 8 = 16
// 2 → 3	1	12	1 × 12 = 12
// 3 → 5	2	3	2 × 3 = 6
// Total Travel Time: 16 + 12 + 6 = 34, which is the minimum possible time after exactly 1 merge.
 
// Constraints:
//     1 <= l <= 10^5
//     2 <= n <= min(l + 1, 50)
//     0 <= k <= min(n - 2, 10)
//     position.length == n
//     position[0] = 0 and position[n - 1] = l
//     position is sorted in strictly increasing order.
//     time.length == n
//     1 <= time[i] <= 100​
//     1 <= sum(time) <= 100​​​​​​

import "fmt"

func minTravelTime(l int, n int, k int, position []int, time []int) int {
    if k == 0 {
        total := 0
        for i := 0; i < n-1; i++ {
            total += (position[i+1] - position[i]) * time[i]
        }
        return total
    }
    sumTime := 0
    for _, t := range time {
        sumTime += t
    }
    memo := make([][][]int, n)
    for i := range memo {
        memo[i] = make([][]int, k+1)
        for j := range memo[i] {
            memo[i][j] = make([]int, sumTime+1)
            for l := range memo[i][j] {
                memo[i][j][l] = -1
            }
        }
    }
    var dfs func(pos, currK, currTime int) int
    dfs = func(pos, currK, currTime int) int {
        if pos == n-1 {
            if currK == 0 {
                return 0
            }
            return 1 << 32
        }
        if memo[pos][currK][currTime] != -1 {
            return memo[pos][currK][currTime]
        }
        minCost := 1 << 32
        nextPos := pos + 1
        if nextPos < n {
            res := dfs(nextPos, currK, time[nextPos])
            if res !=  1 << 32 {
                cost := (position[nextPos] - position[pos]) * currTime
                if cost+res < minCost {
                    minCost = cost + res
                }
            }
        }
        if currK > 0 {
            maxNext := pos + currK + 1
            if maxNext >= n {
                maxNext = n - 1
            }
            timeSum := time[pos+1]
            mergesUsed := 0
            for nextIdx := pos + 2; nextIdx <= maxNext; nextIdx++ {
                mergesUsed++
                if mergesUsed > currK {
                    break
                }
                timeSum += time[nextIdx]
                res := dfs(nextIdx, currK-mergesUsed, timeSum)
                if res != 1 << 32 {
                    cost := (position[nextIdx] - position[pos]) * currTime
                    totalCost := cost + res
                    if totalCost < minCost {
                        minCost = totalCost
                    }
                }
            }
        }
        memo[pos][currK][currTime] = minCost
        return minCost
    }
    return dfs(0, k, time[0])
}

func main() {
    // Example 1:
    // Input: l = 10, n = 4, k = 1, position = [0,3,8,10], time = [5,8,3,6]
    // Output: 62
    // Explanation:
    // Merge the signs at indices 1 and 2. Remove the sign at index 1, and change the time at index 2 to 8 + 3 = 11.
    // After the merge:
    // position array: [0, 8, 10]
    // time array: [5, 11, 6]
    // Segment	Distance (km)	Time per km (min)	Segment Travel Time (min)
    // 0 → 8	8	5	8 × 5 = 40
    // 8 → 10	2	11	2 × 11 = 22
    // Total Travel Time: 40 + 22 = 62, which is the minimum possible time after exactly 1 merge.
    fmt.Println(minTravelTime(10, 4, 1, []int{0,3,8,10}, []int{5,8,3,6})) // 62
    // Example 2:
    // Input: l = 5, n = 5, k = 1, position = [0,1,2,3,5], time = [8,3,9,3,3]
    // Output: 34
    // Explanation:
    // Merge the signs at indices 1 and 2. Remove the sign at index 1, and change the time at index 2 to 3 + 9 = 12.
    // After the merge:
    // position array: [0, 2, 3, 5]
    // time array: [8, 12, 3, 3]
    // Segment	Distance (km)	Time per km (min)	Segment Travel Time (min)
    // 0 → 2	2	8	2 × 8 = 16
    // 2 → 3	1	12	1 × 12 = 12
    // 3 → 5	2	3	2 × 3 = 6
    // Total Travel Time: 16 + 12 + 6 = 34, which is the minimum possible time after exactly 1 merge.
    fmt.Println(minTravelTime(5, 5, 1, []int{0,1,2,3,5}, []int{8,3,9,3,3})) // 34
}