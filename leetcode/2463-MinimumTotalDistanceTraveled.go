package main

// 2463. Minimum Total Distance Traveled
// There are some robots and factories on the X-axis. 
// You are given an integer array robot where robot[i] is the position of the ith robot. 
// You are also given a 2D integer array factory where factory[j] = [positionj, limitj] indicates 
// that positionj is the position of the jth factory and that the jth factory can repair at most limitj robots.

// The positions of each robot are unique. The positions of each factory are also unique. 
// Note that a robot can be in the same position as a factory initially.

// All the robots are initially broken; they keep moving in one direction. 
// The direction could be the negative or the positive direction of the X-axis. 
// When a robot reaches a factory that did not reach its limit, the factory repairs the robot, and it stops moving.

// At any moment, you can set the initial direction of moving for some robot. 
// Your target is to minimize the total distance traveled by all the robots.

// Return the minimum total distance traveled by all the robots. 
// The test cases are generated such that all the robots can be repaired.

// Note that
//     1. All robots move at the same speed.
//     2. If two robots move in the same direction, they will never collide.
//     3. If two robots move in opposite directions and they meet at some point, they do not collide. 
//        They cross each other.
//     4. If a robot passes by a factory that reached its limits, it crosses it as if it does not exist.
//     5. If the robot moved from a position x to a position y, the distance it moved is |y - x|.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/09/15/example1.jpg" />
// Input: robot = [0,4,6], factory = [[2,2],[6,2]]
// Output: 4
// Explanation: As shown in the figure:
// - The first robot at position 0 moves in the positive direction. It will be repaired at the first factory.
// - The second robot at position 4 moves in the negative direction. It will be repaired at the first factory.
// - The third robot at position 6 will be repaired at the second factory. It does not need to move.
// The limit of the first factory is 2, and it fixed 2 robots.
// The limit of the second factory is 2, and it fixed 1 robot.
// The total distance is |2 - 0| + |2 - 4| + |6 - 6| = 4. It can be shown that we cannot achieve a better total distance than 4.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/09/15/example-2.jpg" />
// Input: robot = [1,-1], factory = [[-2,1],[2,1]]
// Output: 2
// Explanation: As shown in the figure:
// - The first robot at position 1 moves in the positive direction. It will be repaired at the second factory.
// - The second robot at position -1 moves in the negative direction. It will be repaired at the first factory.
// The limit of the first factory is 1, and it fixed 1 robot.
// The limit of the second factory is 1, and it fixed 1 robot.
// The total distance is |2 - 1| + |(-2) - (-1)| = 2. It can be shown that we cannot achieve a better total distance than 2.

// Constraints:
//     1 <= robot.length, factory.length <= 100
//     factory[j].length == 2
//     -10^9 <= robot[i], positionj <= 10^9
//     0 <= limitj <= robot.length
//     The input will be generated such that it is always possible to repair every robot.

import "fmt"
import "sort"
import "slices"
import "math"

func minimumTotalDistance(robot []int, factory [][]int) int64 {
    memo := make(map[[3]int]int64, 0)
    sort.Ints(robot)
    slices.SortFunc(factory, func(a, b []int) int {
        return a[0] - b[0]
    })
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    dist := func (a, b int) int64 {
        if (a >= 0 && b >= 0) || (a < 0 && b < 0) {
            if a >= b { return int64(a - b) }
            return int64(b - a)
        } else if a >= 0 && b < 0 {
            return int64(a - b)
        } else if b >= 0 && a < 0 {
            return int64(b - a)
        }
        return 0
    }
    var dp func(r []int, f [][]int) int64
    dp = func(r []int, f [][]int) int64 {
        key := [3]int{len(r), len(f), f[0][1]}
        if v, ok := memo[key]; ok { return v }
        if len(r) == 0 { return 0 }
        if f[0][1] == 0 {
            if len(f) > 1 {
                return dp(r, f[1:])
            } else {
                return math.MaxInt64
            }
        }
        f[0][1]--
        cur := dp(r[1:], f)
        if cur < math.MaxInt64 {
            cur += dist(r[0], f[0][0])
        }
        f[0][1]++
        if len(f) > 1 {
            cur = min(cur, dp(r, f[1:]))
        }
        memo[[3]int{len(r), len(f), f[0][1]}] = cur
        return cur
    }
    return dp(robot, factory)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/09/15/example1.jpg" />
    // Input: robot = [0,4,6], factory = [[2,2],[6,2]]
    // Output: 4
    // Explanation: As shown in the figure:
    // - The first robot at position 0 moves in the positive direction. It will be repaired at the first factory.
    // - The second robot at position 4 moves in the negative direction. It will be repaired at the first factory.
    // - The third robot at position 6 will be repaired at the second factory. It does not need to move.
    // The limit of the first factory is 2, and it fixed 2 robots.
    // The limit of the second factory is 2, and it fixed 1 robot.
    // The total distance is |2 - 0| + |2 - 4| + |6 - 6| = 4. It can be shown that we cannot achieve a better total distance than 4.
    fmt.Println(minimumTotalDistance([]int{0,4,6}, [][]int{{2,2},{6,2}})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/09/15/example-2.jpg" />
    // Input: robot = [1,-1], factory = [[-2,1],[2,1]]
    // Output: 2
    // Explanation: As shown in the figure:
    // - The first robot at position 1 moves in the positive direction. It will be repaired at the second factory.
    // - The second robot at position -1 moves in the negative direction. It will be repaired at the first factory.
    // The limit of the first factory is 1, and it fixed 1 robot.
    // The limit of the second factory is 1, and it fixed 1 robot.
    // The total distance is |2 - 1| + |(-2) - (-1)| = 2. It can be shown that we cannot achieve a better total distance than 2.
    fmt.Println(minimumTotalDistance([]int{1,-1}, [][]int{{-2,1},{2,1}})) // 2
}