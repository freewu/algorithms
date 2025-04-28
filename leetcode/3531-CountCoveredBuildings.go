package main

// 3531. Count Covered Buildings
// You are given a positive integer n, representing an n x n city. 
// You are also given a 2D grid buildings, where buildings[i] = [x, y] denotes a unique building located at coordinates [x, y].

// A building is covered if there is at least one building in all four directions: left, right, above, and below.

// Return the number of covered buildings.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2025/03/04/telegram-cloud-photo-size-5-6212982906394101085-m.jpg" />
// Input: n = 3, buildings = [[1,2],[2,2],[3,2],[2,1],[2,3]]
// Output: 1
// Explanation:
// Only building [2,2] is covered as it has at least one building:
// above ([1,2])
// below ([3,2])
// left ([2,1])
// right ([2,3])
// Thus, the count of covered buildings is 1.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2025/03/04/telegram-cloud-photo-size-5-6212982906394101086-m.jpg" />
// Input: n = 3, buildings = [[1,1],[1,2],[2,1],[2,2]]
// Output: 0
// Explanation:
// No building has at least one building in all four directions.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2025/03/16/telegram-cloud-photo-size-5-6248862251436067566-x.jpg" />
// Input: n = 5, buildings = [[1,3],[3,2],[3,3],[3,5],[5,3]]
// Output: 1
// Explanation:
// Only building [3,3] is covered as it has at least one building:
// above ([1,3])
// below ([5,3])
// left ([3,2])
// right ([3,5])
// Thus, the count of covered buildings is 1.

// Constraints:
//     2 <= n <= 10^5
//     1 <= buildings.length <= 10^5 
//     buildings[i] = [x, y]
//     1 <= x, y <= n
//     All coordinates of buildings are unique.

import "fmt"
import "sort"

func countCoveredBuildings(n int, buildings [][]int) int {
    res, row, col := 0, make(map[int][]int), make(map[int][]int)
    for _, v := range buildings {
        x, y := v[0], v[1]
        row[x], col[y] = append(row[x], y), append(col[y], x)
    }
    for _, v := range row {
        sort.Ints(v)
    }
    for _, v := range col {
        sort.Ints(v)
    }
    check := func(arr []int, i int) bool {
        return arr[0] < i && i < arr[len(arr) - 1]
    }
    for _, b := range buildings {
        x, y := b[0], b[1]
        if check(row[x], y) && check(col[y], x) {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2025/03/04/telegram-cloud-photo-size-5-6212982906394101085-m.jpg" />
    // Input: n = 3, buildings = [[1,2],[2,2],[3,2],[2,1],[2,3]]
    // Output: 1
    // Explanation:
    // Only building [2,2] is covered as it has at least one building:
    // above ([1,2])
    // below ([3,2])
    // left ([2,1])
    // right ([2,3])
    // Thus, the count of covered buildings is 1.
    fmt.Println(countCoveredBuildings(3, [][]int{{1,2},{2,2},{3,2},{2,1},{2,3}})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2025/03/04/telegram-cloud-photo-size-5-6212982906394101086-m.jpg" />
    // Input: n = 3, buildings = [[1,1],[1,2],[2,1],[2,2]]
    // Output: 0
    // Explanation:
    // No building has at least one building in all four directions.
    fmt.Println(countCoveredBuildings(3, [][]int{{1,1},{1,2},{2,1},{2,2}})) // 0
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2025/03/16/telegram-cloud-photo-size-5-6248862251436067566-x.jpg" />
    // Input: n = 5, buildings = [[1,3],[3,2],[3,3],[3,5],[5,3]]
    // Output: 1
    // Explanation:
    // Only building [3,3] is covered as it has at least one building:
    // above ([1,3])
    // below ([5,3])
    // left ([3,2])
    // right ([3,5])
    // Thus, the count of covered buildings is 1.
    fmt.Println(countCoveredBuildings(5, [][]int{{1,3},{3,2},{3,3},{3,5},{5,3}})) // 1
}