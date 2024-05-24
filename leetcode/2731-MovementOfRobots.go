package main

// 2731. Movement of Robots
// Some robots are standing on an infinite number line with their initial coordinates given by a 0-indexed integer array nums 
// and will start moving once given the command to move. 
// The robots will move a unit distance each second.

// You are given a string s denoting the direction in which robots will move on command. 
// 'L' means the robot will move towards the left side or negative side of the number line, 
// whereas 'R' means the robot will move towards the right side or positive side of the number line.

// If two robots collide, they will start moving in opposite directions.
// Return the sum of distances between all the pairs of robots d seconds after the command. Since the sum can be very large, return it modulo 10^9 + 7.

// Note:
//     For two robots at the index i and j, pair (i,j) and pair (j,i) are considered the same pair.
//     When robots collide, they instantly change their directions without wasting any time.
//     Collision happens when two robots share the same place in a moment.
//         For example, 
//             if a robot is positioned in 0 going to the right and another is positioned in 2 going to the left, 
//             the next second they'll be both in 1 and they will change direction and the next second the first one will be in 0, 
//             heading left, and another will be in 2, heading right.
//         For example, 
//             if a robot is positioned in 0 going to the right and another is positioned in 1 going to the left, 
//             the next second the first one will be in 0, heading left, and another will be in 1, heading right.

// Example 1:
// Input: nums = [-2,0,2], s = "RLL", d = 3
// Output: 8
// Explanation: 
// After 1 second, the positions are [-1,-1,1]. Now, the robot at index 0 will move left, and the robot at index 1 will move right.
// After 2 seconds, the positions are [-2,0,0]. Now, the robot at index 1 will move left, and the robot at index 2 will move right.
// After 3 seconds, the positions are [-3,-1,1].
// The distance between the robot at index 0 and 1 is abs(-3 - (-1)) = 2.
// The distance between the robot at index 0 and 2 is abs(-3 - 1) = 4.
// The distance between the robot at index 1 and 2 is abs(-1 - 1) = 2.
// The sum of the pairs of all distances = 2 + 4 + 2 = 8.

// Example 2:
// Input: nums = [1,0], s = "RL", d = 2
// Output: 5
// Explanation: 
// After 1 second, the positions are [2,-1].
// After 2 seconds, the positions are [3,-2].
// The distance between the two robots is abs(-2 - 3) = 5.
 
// Constraints:
//     2 <= nums.length <= 10^5
//     -2 * 10^9 <= nums[i] <= 2 * 10^9
//     0 <= d <= 10^9
//     nums.length == s.length 
//     s consists of 'L' and 'R' only
//     nums[i] will be unique.

import "fmt"
import "sort"

func sumDistance(nums []int, s string, d int) int {
    res, n, mod := 0, len(nums), int(1e9) + 7
    for i := 0; i < n; i++ {
        if s[i] == 'R' {
            nums[i] += d
        } else {
            nums[i] -= d
        }
    }
    sort.Ints(nums)
    prefix := nums[0]
    for i := 1; i < n; i++ {
        res += (i * nums[i] - prefix)
        res %= mod
        prefix += nums[i]
    }
    return res
}

func sumDistance1(nums []int, s string, d int) int {
    res, n, mod := 0, len(nums), int(1e9) + 7
    for i := 0; i < n; i++ {
        if s[i] == 'R' {
            nums[i] += d
        } else {
            nums[i] -= d
        }
    }
    sort.Ints(nums)
    sum := nums[0]
    for i := 1; i < n; i++ {
        res = (res + nums[i] * i - sum) % mod
        sum = (sum + nums[i]) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [-2,0,2], s = "RLL", d = 3
    // Output: 8
    // Explanation: 
    // After 1 second, the positions are [-1,-1,1]. Now, the robot at index 0 will move left, and the robot at index 1 will move right.
    // After 2 seconds, the positions are [-2,0,0]. Now, the robot at index 1 will move left, and the robot at index 2 will move right.
    // After 3 seconds, the positions are [-3,-1,1].
    // The distance between the robot at index 0 and 1 is abs(-3 - (-1)) = 2.
    // The distance between the robot at index 0 and 2 is abs(-3 - 1) = 4.
    // The distance between the robot at index 1 and 2 is abs(-1 - 1) = 2.
    // The sum of the pairs of all distances = 2 + 4 + 2 = 8.
    fmt.Println(sumDistance([]int{-2,0,2},"RLL", 3)) // 8
    // Example 2:
    // Input: nums = [1,0], s = "RL", d = 2
    // Output: 5
    // Explanation: 
    // After 1 second, the positions are [2,-1].
    // After 2 seconds, the positions are [3,-2].
    // The distance between the two robots is abs(-2 - 3) = 5.
    fmt.Println(sumDistance([]int{1,0},"RL", 2)) // 5

    fmt.Println(sumDistance1([]int{-2,0,2},"RLL", 3)) // 8
    fmt.Println(sumDistance1([]int{1,0},"RL", 2)) // 5
}