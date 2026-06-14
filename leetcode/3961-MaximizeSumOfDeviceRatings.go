package main

// 3961. Maximize Sum of Device Ratings
// You are given a 2D integer array units of size m × n where units[i][j] represents the capacity of the jth unit in the ith device. 
// Each device contains exactly n units.

// The rating of a device is the minimum capacity among all its units.

// You may perform the following operation any number of times (including zero):
//     1. Choose a device i that has not been used as a source before.
//     2. Remove exactly one unit from device i and add it to any different device.
//     3. Then mark device i as used, so it cannot be chosen again as a source.

// Return the maximum possible sum of the ratings of all devices after any number of such operations.

// Note:
//     1. Devices can receive units from multiple devices, regardless of whether they have been selected.
//     2. The rating of an empty device is 0.

// Example 1:
// Input: units = [[1,3],[2,2]]
// Output: 4
// Explanation:
// ​​​​​​​​​​​​​​Select device i = 0 and transfer units[0][0] = 1 to device i = 1.
// After the transfer, the ratings are:
// Device 0 = [3]: rating[0] = 3
// Device 1 = [2, 2, 1]: rating[1] = 1
// Thus, the sum of ratings is 3 + 1 = 4.

// Example 2:
// Input: units = [[1,2,3],[4,5,6]]
// Output: 6
// Explanation:
// Select device i = 1 and transfer units[1][0] = 4 to device i = 0.
// After the transfer, the ratings are:
// Device 0 = [1, 2, 3, 4]: rating[0] = 1
// Device 1 = [5, 6]: rating[1] = 5
// Thus, the sum of ratings is 1 + 5 = 6.

// Example 3:
// Input: units = [[5,5,5],[1,1,1]]
// Output: 6
// Explanation:
// No transfers increase the sum of ratings. Thus, the sum of ratings is 5 + 1 = 6.

// Constraints:
//     1 <= m == units.length <= 10^5
//     1 <= n == units[i].length <= 10^5
//     m * n <= 2 * 10^5
//     1 <= units[i][j] <= 10^5

import "fmt"

func maxRatings(units [][]int) int64 {
    if len(units[0]) == 1 { // single-column case
        sum := 0
        for _, v := range units {
            sum += v[0]
        }
        return int64(sum)
    }
    sum, gmn1, gmn2 := 0, 1 << 61, 1 << 61
    for _, unit := range units {
        mn1, mn2 := 1 << 61, 1 << 61
        for _, val := range unit {
            if val < mn1 {
                mn2 = mn1
                mn1 = val
            } else if val < mn2 {
                mn2 = val
            }
        }
        sum += mn2 // Accumulate the sum of all 2nd minimums
        if mn1 < gmn1 { // Track the absolute smallest 1st and 2nd minimums across all rows
            gmn1 = mn1
        }
        if mn2 < gmn2 {
            gmn2 = mn2
        }
    }
    return int64(sum - gmn2 + gmn1)
}

func main() {
    // Example 1:
    // Input: units = [[1,3],[2,2]]
    // Output: 4
    // Explanation:
    // ​​​​​​​​​​​​​​Select device i = 0 and transfer units[0][0] = 1 to device i = 1.
    // After the transfer, the ratings are:
    // Device 0 = [3]: rating[0] = 3
    // Device 1 = [2, 2, 1]: rating[1] = 1
    // Thus, the sum of ratings is 3 + 1 = 4.
    fmt.Println(maxRatings([][]int{{1,3},{2,2}})) // 4
    // Example 2:
    // Input: units = [[1,2,3],[4,5,6]]
    // Output: 6
    // Explanation:
    // Select device i = 1 and transfer units[1][0] = 4 to device i = 0.
    // After the transfer, the ratings are:
    // Device 0 = [1, 2, 3, 4]: rating[0] = 1
    // Device 1 = [5, 6]: rating[1] = 5
    // Thus, the sum of ratings is 1 + 5 = 6.
    fmt.Println(maxRatings([][]int{{1,2,3},{4,5,6}})) // 6
    // Example 3:
    // Input: units = [[5,5,5],[1,1,1]]
    // Output: 6
    // Explanation:
    // No transfers increase the sum of ratings. Thus, the sum of ratings is 5 + 1 = 6. 
    fmt.Println(maxRatings([][]int{{5,5,5},{1,1,1}})) // 6
}