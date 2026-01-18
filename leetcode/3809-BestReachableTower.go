package main

// 3809. Best Reachable Tower
// You are given a 2D integer array towers, where towers[i] = [xi, yi, qi] represents the coordinates (xi, yi) and quality factor qi of the ith tower.

// You are also given an integer array center = [cx, cy​​​​​​​] representing your location, and an integer radius.

// A tower is reachable if its Manhattan distance from center is less than or equal to radius.

// Among all reachable towers:

//     1. Return the coordinates of the tower with the maximum quality factor.
//     2. If there is a tie, return the tower with the lexicographically smallest coordinate. If no tower is reachable, return [-1, -1].

// The Manhattan Distance between two cells (xi, yi) and (xj, yj) is |xi - xj| + |yi - yj|.
// A coordinate [xi, yi] is lexicographically smaller than [xj, yj] if xi < xj, or xi == xj and yi < yj.

// |x| denotes the absolute value of x.

// Example 1:
// Input: towers = [[1,2,5], [2,1,7], [3,1,9]], center = [1,1], radius = 2
// Output: [3,1]
// Explanation:
// Tower [1, 2, 5]: Manhattan distance = |1 - 1| + |2 - 1| = 1, reachable.
// Tower [2, 1, 7]: Manhattan distance = |2 - 1| + |1 - 1| = 1, reachable.
// Tower [3, 1, 9]: Manhattan distance = |3 - 1| + |1 - 1| = 2, reachable.
// All towers are reachable. The maximum quality factor is 9, which corresponds to tower [3, 1].

// Example 2:
// Input: towers = [[1,3,4], [2,2,4], [4,4,7]], center = [0,0], radius = 5
// Output: [1,3]
// Explanation:
// Tower [1, 3, 4]: Manhattan distance = |1 - 0| + |3 - 0| = 4, reachable.
// Tower [2, 2, 4]: Manhattan distance = |2 - 0| + |2 - 0| = 4, reachable.
// Tower [4, 4, 7]: Manhattan distance = |4 - 0| + |4 - 0| = 8, not reachable.
// Among the reachable towers, the maximum quality factor is 4. Both [1, 3] and [2, 2] have the same quality, so the lexicographically smaller coordinate is [1, 3].

// Example 3:
// Input: towers = [[5,6,8], [0,3,5]], center = [1,2], radius = 1
// Output: [-1,-1]
// Explanation:
// Tower [5, 6, 8]: Manhattan distance = |5 - 1| + |6 - 2| = 8, not reachable.
// Tower [0, 3, 5]: Manhattan distance = |0 - 1| + |3 - 2| = 2, not reachable.
// No tower is reachable within the given radius, so [-1, -1] is returned.

// Constraints:
//     1 <= towers.length <= 10^5
//     towers[i] = [xi, yi, qi]
//     center = [cx, cy]
//     0 <= xi, yi, qi, cx, cy <= 10^5​​​​​​​
//     0 <= radius <= 10^5

import "fmt"

func bestTower(towers [][]int, center []int, radius int) []int {
    x, y, mx := 0, 0, -1
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 0; i < len(towers); i++ {
        if abs(center[0] - towers[i][0]) + abs(center[1] - towers[i][1]) <= radius {
            if towers[i][2] > mx {
                mx = towers[i][2]
                x = towers[i][0]
                y = towers[i][1]
            } else if towers[i][2] == mx {
                if towers[i][0] < x || (towers[i][0] == x && towers[i][1] < y) {
                    x = towers[i][0]
                    y = towers[i][1]
                }
            }
        }
    }
    if mx == -1 { return []int{-1, -1} }
    return []int{x, y}
}

func bestTower1(towers [][]int, center []int, radius int) []int {
    res := []int{ 1000000, 1000000, -1000000}
    x0, y0 := center[0],center[1]
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    check := func(p1,p2 []int)bool {
        if p1[2] > p2[2] { return true }
        if p1[2] < p2[2] { return false }
        x1,y1 := p1[0],p1[1]
        x2,y2 := p2[0],p2[2]
        if x1 == x2 { return y1 < y2 }
        return x1 < x2
    }
    for _, point := range towers {
        x1, y1 := point[0], point[1]
        tmp := abs(x1 - x0) + abs(y1 - y0)
        if tmp <= radius && check(point,res){
            res = point
        }
    }
    if res[0] > 100001 { return []int{-1,-1} }
    return res[:2]
}

func main() {
    // Example 1:
    // Input: towers = [[1,2,5], [2,1,7], [3,1,9]], center = [1,1], radius = 2
    // Output: [3,1]
    // Explanation:
    // Tower [1, 2, 5]: Manhattan distance = |1 - 1| + |2 - 1| = 1, reachable.
    // Tower [2, 1, 7]: Manhattan distance = |2 - 1| + |1 - 1| = 1, reachable.
    // Tower [3, 1, 9]: Manhattan distance = |3 - 1| + |1 - 1| = 2, reachable.
    // All towers are reachable. The maximum quality factor is 9, which corresponds to tower [3, 1].
    fmt.Println(bestTower([][]int{{1,2,5}, {2,1,7}, {3,1,9}}, []int{1,1}, 2)) // [3,1]
    // Example 2:
    // Input: towers = [[1,3,4], [2,2,4], [4,4,7]], center = [0,0], radius = 5
    // Output: [1,3]
    // Explanation:
    // Tower [1, 3, 4]: Manhattan distance = |1 - 0| + |3 - 0| = 4, reachable.
    // Tower [2, 2, 4]: Manhattan distance = |2 - 0| + |2 - 0| = 4, reachable.
    // Tower [4, 4, 7]: Manhattan distance = |4 - 0| + |4 - 0| = 8, not reachable.
    // Among the reachable towers, the maximum quality factor is 4. Both [1, 3] and [2, 2] have the same quality, so the lexicographically smaller coordinate is [1, 3].
    fmt.Println(bestTower([][]int{{1,3,4}, {2,2,4}, {4,4,7}}, []int{0,0}, 5)) // [1,3]
    // Example 3:
    // Input: towers = [[5,6,8], [0,3,5]], center = [1,2], radius = 1
    // Output: [-1,-1]
    // Explanation:
    // Tower [5, 6, 8]: Manhattan distance = |5 - 1| + |6 - 2| = 8, not reachable.
    // Tower [0, 3, 5]: Manhattan distance = |0 - 1| + |3 - 2| = 2, not reachable.
    // No tower is reachable within the given radius, so [-1, -1] is returned.
    fmt.Println(bestTower([][]int{{5,6,8}, {0,3,5}}, []int{1,2}, 1)) // [-1,-1]

    fmt.Println(bestTower1([][]int{{1,2,5}, {2,1,7}, {3,1,9}}, []int{1,1}, 2)) // [3,1]
    fmt.Println(bestTower1([][]int{{1,3,4}, {2,2,4}, {4,4,7}}, []int{0,0}, 5)) // [1,3]
    fmt.Println(bestTower1([][]int{{5,6,8}, {0,3,5}}, []int{1,2}, 1)) // [-1,-1]
}