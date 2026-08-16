package main

// 4024. Nearest Available Drone
// You are given a 2D integer array drones, where drones[i] = [xi, yi, rangei] represents the x-coordinate, y-coordinate, and travel range of the ith drone.

// You are also given an integer array target = [tx, ty], representing the coordinates of the target.

// A drone drones[i] can reach the target if the Manhattan distance between its coordinates and the target coordinates is less than or equal to its rangei.

// Return the index of the reachable drone with the minimum Manhattan distance to the target. 
// If there is a tie, return the smallest index. 
// If no drone can reach the target, return -1.

// The Manhattan distance between two coordinates (xi, yi) and (xj, yj) is |xi - xj| + |yi - yj|.

// Example 1:
// Input: drones = [[0,0,8],[2,2,9]], target = [3,4]
// Output: 1
// Explanation:
// The distance between drones[0] and target is |0 - 3| + |0 - 4| = 7, which is within its range of 8.
// The distance between drones[1] and target is |2 - 3| + |2 - 4| = 3, which is within its range of 9.
// Since drones[1] is the nearest drone, the answer is 1.

// Example 2:
// Input: drones = [[2,1,5],[4,4,5],[6,6,8]], target = [5,5]
// Output: 1
// Explanation:
// The distance between drones[0] and target is |2 - 5| + |1 - 5| = 7, which is greater than its range of 5.
// The distance between drones[1] and target is |4 - 5| + |4 - 5| = 2, which is within its range of 5.
// The distance between drones[2] and target is |6 - 5| + |6 - 5| = 2, which is within its range of 8.
// Both drones[1] and drones[2] are the nearest drones. Since we should return the smallest index, the answer is 1.

// Example 3:
// Input: drones = [[4,4,5]], target = [8,6]
// Output: -1
// Explanation:
// The distance between drones[0] and target is |4 - 8| + |4 - 6| = 6, which is greater than its range of 5.
// No drone can reach the target, so the answer is -1.

// Constraints:
//     1 <= drones.length <= 100
//     drones[i] = [xi, yi, rangei]
//     target = [tx, ty]
//     -25 <= xi, yi, tx, ty <= 25
//     1 <= rangei <= 100

import "fmt"

func nearestDrone(drones [][]int, target []int) int {
    res, best, dx, dy := -1, 1 << 61, target[0], target[1]
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, drone := range drones {
        x, y, r := drone[0], drone[1], drone[2]
        dist := abs(x-dx) + abs(y-dy)
        if dist <= r && dist < best {
            res, best = i, dist
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: drones = [[0,0,8],[2,2,9]], target = [3,4]
    // Output: 1
    // Explanation:
    // The distance between drones[0] and target is |0 - 3| + |0 - 4| = 7, which is within its range of 8.
    // The distance between drones[1] and target is |2 - 3| + |2 - 4| = 3, which is within its range of 9.
    // Since drones[1] is the nearest drone, the answer is 1.
    fmt.Println(nearestDrone([][]int{ []int{0,0,8},[]int{2,2,9}}, []int{3,4})) // 1
    // Example 2:
    // Input: drones = [[2,1,5],[4,4,5],[6,6,8]], target = [5,5]
    // Output: 1
    // Explanation:
    // The distance between drones[0] and target is |2 - 5| + |1 - 5| = 7, which is greater than its range of 5.
    // The distance between drones[1] and target is |4 - 5| + |4 - 5| = 2, which is within its range of 5.
    // The distance between drones[2] and target is |6 - 5| + |6 - 5| = 2, which is within its range of 8.
    // Both drones[1] and drones[2] are the nearest drones. Since we should return the smallest index, the answer is 1.
    fmt.Println(nearestDrone([][]int{ []int{2,1,5},[]int{4,4,5},[]int{6,6,8}}, []int{5,5})) // 1
    // Example 3:
    // Input: drones = [[4,4,5]], target = [8,6]
    // Output: -1
    // Explanation:
    // The distance between drones[0] and target is |4 - 8| + |4 - 6| = 6, which is greater than its range of 5.
    // No drone can reach the target, so the answer is -1.
    fmt.Println(nearestDrone([][]int{ []int{4,4,5}}, []int{8,6})) // -1
}