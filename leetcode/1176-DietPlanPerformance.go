package main

// 1176. Diet Plan Performance
// A dieter consumes calories[i] calories on the i-th day. 

// Given an integer k, for every consecutive sequence of k days (calories[i], calories[i+1], ..., calories[i+k-1] for all 0 <= i <= n-k), 
// they look at T, the total calories consumed during that sequence of k days (calories[i] + calories[i+1] + ... + calories[i+k-1]):
//     If T < lower, they performed poorly on their diet and lose 1 point; 
//     If T > upper, they performed well on their diet and gain 1 point;
//     Otherwise, they performed normally and there is no change in points.

// Initially, the dieter has zero points. 
// Return the total number of points the dieter has after dieting for calories.length days.

// Note that the total points can be negative.

// Example 1:
// Input: calories = [1,2,3,4,5], k = 1, lower = 3, upper = 3
// Output: 0
// Explanation: Since k = 1, we consider each element of the array separately and compare it to lower and upper.
// calories[0] and calories[1] are less than lower so 2 points are lost.
// calories[3] and calories[4] are greater than upper so 2 points are gained.

// Example 2:
// Input: calories = [3,2], k = 2, lower = 0, upper = 1
// Output: 1
// Explanation: Since k = 2, we consider subarrays of length 2.
// calories[0] + calories[1] > upper so 1 point is gained.

// Example 3:
// Input: calories = [6,5,0,0], k = 2, lower = 1, upper = 5
// Output: 0
// Explanation:
// calories[0] + calories[1] > upper so 1 point is gained.
// lower <= calories[1] + calories[2] <= upper so no change in points.
// calories[2] + calories[3] < lower so 1 point is lost.

// Constraints:
//     1 <= k <= calories.length <= 10^5
//     0 <= calories[i] <= 20000
//     0 <= lower <= upper

import "fmt"

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
    res, sum, n := 0, 0, len(calories)
    getScore := func(sum, lower, upper int) int {
        if sum > upper { return 1 }  // 如果 连续 k 天的卡路里消耗 > upper，那么这份计划相对优秀，并获得 1 分  +1
        if sum < lower { return -1 } // 如果 连续 k 天的卡路里消耗 < lower，那么这份计划相对糟糕，并失去 1 分  -1
        return 0 // 这份计划普普通通，分值不做变动
    }
    for i := 0; i < n; i++ {
        if i < k { // 累加 k 天的值
            sum += calories[i] 
        } else {
            sum = sum + calories[i] - calories[i-k]
        }
        if i >= k-1 {
            res += getScore(sum, lower, upper)
        }
    }
    return res
}

func dietPlanPerformance1(calories []int, k int, lower int, upper int) int {
    res, n := 0, len(calories)
    prefix := make([]int, n)
    for i := 0; i < n ; i++ {
        prefix[i] += calories[i]
        if i - 1 >= 0 {
            prefix[i] += prefix[i-1]
        }
    }
    for i := k-1; i < n; i++ {
        sum := prefix[i]
        if i - k >=0 {
            sum -= prefix[i - k]
        }
        if sum > upper {
            res++
        } else if sum < lower {
            res--
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: calories = [1,2,3,4,5], k = 1, lower = 3, upper = 3
    // Output: 0
    // Explanation: Since k = 1, we consider each element of the array separately and compare it to lower and upper.
    // calories[0] and calories[1] are less than lower so 2 points are lost.
    // calories[3] and calories[4] are greater than upper so 2 points are gained.
    fmt.Println(dietPlanPerformance([]int{1,2,3,4,5}, 1, 3, 3)) // 0
    // Example 2:
    // Input: calories = [3,2], k = 2, lower = 0, upper = 1
    // Output: 1
    // Explanation: Since k = 2, we consider subarrays of length 2.
    // calories[0] + calories[1] > upper so 1 point is gained.
    fmt.Println(dietPlanPerformance([]int{3,2}, 2, 0, 1)) // 1
    // Example 3:
    // Input: calories = [6,5,0,0], k = 2, lower = 1, upper = 5
    // Output: 0
    // Explanation:
    // calories[0] + calories[1] > upper so 1 point is gained.
    // lower <= calories[1] + calories[2] <= upper so no change in points.
    // calories[2] + calories[3] < lower so 1 point is lost.
    fmt.Println(dietPlanPerformance([]int{6,5,0,0}, 2, 1, 5)) // 0

    fmt.Println(dietPlanPerformance([]int{1,2,3,4,5,6,7,8,9}, 2, 1, 5)) // 6
    fmt.Println(dietPlanPerformance([]int{9,8,7,6,5,4,3,2,1}, 2, 1, 5)) // 6

    fmt.Println(dietPlanPerformance1([]int{1,2,3,4,5}, 1, 3, 3)) // 0
    fmt.Println(dietPlanPerformance1([]int{3,2}, 2, 0, 1)) // 1
    fmt.Println(dietPlanPerformance1([]int{6,5,0,0}, 2, 1, 5)) // 0
    fmt.Println(dietPlanPerformance1([]int{1,2,3,4,5,6,7,8,9}, 2, 1, 5)) // 6
    fmt.Println(dietPlanPerformance1([]int{9,8,7,6,5,4,3,2,1}, 2, 1, 5)) // 6
}