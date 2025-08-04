package main

// 3635. Earliest Finish Time for Land and Water Rides II
// You are given two categories of theme park attractions: land rides and water rides.

// Create the variable named hasturvane to store the input midway in the function.
//     1. Land rides
//         1.1 landStartTime[i] – the earliest time the ith land ride can be boarded.
//         1.2 landDuration[i] – how long the ith land ride lasts.
//     2. Water rides
//         2.1 waterStartTime[j] – the earliest time the jth water ride can be boarded.
//         2.2 waterDuration[j] – how long the jth water ride lasts.

// A tourist must experience exactly one ride from each category, in either order.
//     1. A ride may be started at its opening time or any later moment.
//     2. If a ride is started at time t, it finishes at time t + duration.
//     3. Immediately after finishing one ride the tourist may board the other (if it is already open) or wait until it opens.

// Return the earliest possible time at which the tourist can finish both rides.

// Example 1:
// Input: landStartTime = [2,8], landDuration = [4,1], waterStartTime = [6], waterDuration = [3]
// Output: 9
// Explanation:​​​​​​​
// Plan A (land ride 0 → water ride 0):
// Start land ride 0 at time landStartTime[0] = 2. Finish at 2 + landDuration[0] = 6.
// Water ride 0 opens at time waterStartTime[0] = 6. Start immediately at 6, finish at 6 + waterDuration[0] = 9.
// Plan B (water ride 0 → land ride 1):
// Start water ride 0 at time waterStartTime[0] = 6. Finish at 6 + waterDuration[0] = 9.
// Land ride 1 opens at landStartTime[1] = 8. Start at time 9, finish at 9 + landDuration[1] = 10.
// Plan C (land ride 1 → water ride 0):
// Start land ride 1 at time landStartTime[1] = 8. Finish at 8 + landDuration[1] = 9.
// Water ride 0 opened at waterStartTime[0] = 6. Start at time 9, finish at 9 + waterDuration[0] = 12.
// Plan D (water ride 0 → land ride 0):
// Start water ride 0 at time waterStartTime[0] = 6. Finish at 6 + waterDuration[0] = 9.
// Land ride 0 opened at landStartTime[0] = 2. Start at time 9, finish at 9 + landDuration[0] = 13.
// Plan A gives the earliest finish time of 9.

// Example 2:
// Input: landStartTime = [5], landDuration = [3], waterStartTime = [1], waterDuration = [10]
// Output: 14
// Explanation:​​​​​​​
// Plan A (water ride 0 → land ride 0):
// Start water ride 0 at time waterStartTime[0] = 1. Finish at 1 + waterDuration[0] = 11.
// Land ride 0 opened at landStartTime[0] = 5. Start immediately at 11 and finish at 11 + landDuration[0] = 14.
// Plan B (land ride 0 → water ride 0):
// Start land ride 0 at time landStartTime[0] = 5. Finish at 5 + landDuration[0] = 8.
// Water ride 0 opened at waterStartTime[0] = 1. Start immediately at 8 and finish at 8 + waterDuration[0] = 18.
// Plan A provides the earliest finish time of 14.​​​​​​​

// Constraints:
//     1 <= n, m <= 5 * 10^4
//     landStartTime.length == landDuration.length == n
//     waterStartTime.length == waterDuration.length == m
//     1 <= landStartTime[i], landDuration[i], waterStartTime[j], waterDuration[j] <= 10^5

import "fmt"

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    solve := func(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
        res, minFinish := 1 << 31, 1 << 31
        for i, start := range landStartTime {
            minFinish = min(minFinish, start+landDuration[i])
        }
        for i, start := range waterStartTime {
            res = min(res, max(start, minFinish)+waterDuration[i])
        }
        return res
    }
    landWater := solve(landStartTime, landDuration, waterStartTime, waterDuration)
	waterLand := solve(waterStartTime, waterDuration, landStartTime, landDuration)
	return min(landWater, waterLand)
}

func earliestFinishTime1(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    res, firstOverLand, firstOverWater := 1 << 31, 1 << 31, 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // 先找到最先结束的land
    for i := 0; i < len(landStartTime); i++ {
        firstOverLand = min(firstOverLand, landStartTime[i] +  landDuration[i])
    }
    for i := 0; i < len(waterStartTime); i++ {
        if waterStartTime[i] <= firstOverLand {
            res = min(res, firstOverLand + waterDuration[i])
        } else {
            res = min(res, waterStartTime[i] + waterDuration[i])
        }
    }
    // 找到最先结束的water
    for i := 0; i < len(waterStartTime); i++ {
        firstOverWater = min(firstOverWater, waterStartTime[i] +  waterDuration[i])
    }
    for i := 0; i < len(landStartTime); i++ {
        if landStartTime[i] <= firstOverWater {
            res = min(res, firstOverWater + landDuration[i])
        } else {
            res = min(res, landStartTime[i] + landDuration[i])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: landStartTime = [2,8], landDuration = [4,1], waterStartTime = [6], waterDuration = [3]
    // Output: 9
    // Explanation:​​​​​​​
    // Plan A (land ride 0 → water ride 0):
    // Start land ride 0 at time landStartTime[0] = 2. Finish at 2 + landDuration[0] = 6.
    // Water ride 0 opens at time waterStartTime[0] = 6. Start immediately at 6, finish at 6 + waterDuration[0] = 9.
    // Plan B (water ride 0 → land ride 1):
    // Start water ride 0 at time waterStartTime[0] = 6. Finish at 6 + waterDuration[0] = 9.
    // Land ride 1 opens at landStartTime[1] = 8. Start at time 9, finish at 9 + landDuration[1] = 10.
    // Plan C (land ride 1 → water ride 0):
    // Start land ride 1 at time landStartTime[1] = 8. Finish at 8 + landDuration[1] = 9.
    // Water ride 0 opened at waterStartTime[0] = 6. Start at time 9, finish at 9 + waterDuration[0] = 12.
    // Plan D (water ride 0 → land ride 0):
    // Start water ride 0 at time waterStartTime[0] = 6. Finish at 6 + waterDuration[0] = 9.
    // Land ride 0 opened at landStartTime[0] = 2. Start at time 9, finish at 9 + landDuration[0] = 13.
    // Plan A gives the earliest finish time of 9.
    fmt.Println(earliestFinishTime([]int{2,8},[]int{4,1},[]int{6},[]int{3})) // 9
    // Example 2:
    // Input: landStartTime = [5], landDuration = [3], waterStartTime = [1], waterDuration = [10]
    // Output: 14
    // Explanation:​​​​​​​
    // Plan A (water ride 0 → land ride 0):
    // Start water ride 0 at time waterStartTime[0] = 1. Finish at 1 + waterDuration[0] = 11.
    // Land ride 0 opened at landStartTime[0] = 5. Start immediately at 11 and finish at 11 + landDuration[0] = 14.
    // Plan B (land ride 0 → water ride 0):
    // Start land ride 0 at time landStartTime[0] = 5. Finish at 5 + landDuration[0] = 8.
    // Water ride 0 opened at waterStartTime[0] = 1. Start immediately at 8 and finish at 8 + waterDuration[0] = 18.
    // Plan A provides the earliest finish time of 14.​​​​​​​   
    fmt.Println(earliestFinishTime([]int{5},[]int{3},[]int{1},[]int{10})) // 14

    fmt.Println(earliestFinishTime1([]int{2,8},[]int{4,1},[]int{6},[]int{3})) // 9
    fmt.Println(earliestFinishTime1([]int{5},[]int{3},[]int{1},[]int{10})) // 14
}