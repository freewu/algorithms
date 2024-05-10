package main

// 1732. Find the Highest Altitude
// There is a biker going on a road trip. 
// The road trip consists of n + 1 points at different altitudes. 
// The biker starts his trip on point 0 with altitude equal 0.

// You are given an integer array gain of length n where gain[i] is the net gain in altitude between points i​​​​​​ and i + 1 for all (0 <= i < n). 
// Return the highest altitude of a point.

// Example 1:
// Input: gain = [-5,1,5,0,-7]
// Output: 1
// Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.

// Example 2:
// Input: gain = [-4,-3,-2,-1,4,3,2]
// Output: 0
// Explanation: The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.

// Constraints:
//     n == gain.length
//     1 <= n <= 100
//     -100 <= gain[i] <= 100

import "fmt"

func largestAltitude(gain []int) int {
    res, altitude := 0, 0
    // 自行车手从海拔为 0 的点 0 开始骑行
    // 给的是净海拔高度差的数组,依次重放即可
    for i := 0 ; i < len(gain); i++ {
        altitude += gain[i]
        if altitude > res {
            res = altitude
        }
    }
    return res
}

func largestAltitude1(gain []int) int {
    res, altitude := 0, 0
    for _, v := range gain {
        altitude += v
        if altitude > res {
            res = altitude
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: gain = [-5,1,5,0,-7]
    // Output: 1
    // Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
    fmt.Println(largestAltitude([]int{-5,1,5,0,-7})) // 1
    // Example 2:
    // Input: gain = [-4,-3,-2,-1,4,3,2]
    // Output: 0
    // Explanation: The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.
    fmt.Println(largestAltitude([]int{-4,-3,-2,-1,4,3,2})) // 0

    fmt.Println(largestAltitude1([]int{-5,1,5,0,-7})) // 1
    fmt.Println(largestAltitude1([]int{-4,-3,-2,-1,4,3,2})) // 0
}