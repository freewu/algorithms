package main 

// 1654. Minimum Jumps to Reach Home
// A certain bug's home is on the x-axis at position x. Help them get there from position 0.

// The bug jumps according to the following rules:
//     It can jump exactly a positions forward (to the right).
//     It can jump exactly b positions backward (to the left).
//     It cannot jump backward twice in a row.
//     It cannot jump to any forbidden positions.

// The bug may jump forward beyond its home, but it cannot jump to positions numbered with negative integers.

// Given an array of integers forbidden, where forbidden[i] means that the bug cannot jump to the position forbidden[i], 
// and integers a, b, and x, return the minimum number of jumps needed for the bug to reach its home.
// If there is no possible sequence of jumps that lands the bug on position x, return -1.

// Example 1:
// Input: forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
// Output: 3
// Explanation: 3 jumps forward (0 -> 3 -> 6 -> 9) will get the bug home.

// Example 2:
// Input: forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
// Output: -1

// Example 3:
// Input: forbidden = [1,6,2,14,5,17,4], a = 16, b = 9, x = 7
// Output: 2
// Explanation: One jump forward (0 -> 16) then one jump backward (16 -> 7) will get the bug home.

// Constraints:
//     1 <= forbidden.length <= 1000
//     1 <= a, b, forbidden[i] <= 2000
//     0 <= x <= 2000
//     All the elements in forbidden are distinct.
//     Position x is not forbidden.

import "fmt"

// // bfs 解答错误 94 / 96
// func minimumJumps(forbidden []int, a int, b int, x int) int {
//     if x == 0 { return 0 }
//     type Node struct {
//         Pos int
//         Back bool
//     }
//     count, queue, visited := 0, make([]Node, 0), make(map[int]bool)
//     queue = append(queue, Node{ 0, true })
//     for i := range forbidden {
//         visited[forbidden[i]] = true
//     }
//     for len(queue) > 0 {
//         n := len(queue)
//         for i := 0; i < n; i++ {
//             v := queue[i]
//             next1 := v.Pos + a
//             if next1 <= 4000 && !visited[next1] {
//                 if next1 == x { return count + 1 }
//                 queue = append(queue, Node{ next1, false })
//                 visited[next1] = true
//             }
//             next2 := v.Pos - b
//             if next2 >= 0 && !visited[next2] && !v.Back {
//                 if next2 == x { return count + 1 }
//                 queue = append(queue, Node{ next2, true })
//             }
//         }
//         queue = queue[n:] // pop
//         count++
//     }
//     return -1
// }

func minimumJumps(forbidden []int, a int, b int, x int) int {
    queue, set, steps := []int{0}, make(map[int]bool), make([]int, 10000)
    for _, v := range forbidden { set[v] = true } // set forbidden
    for i := 1; i < len(steps); i++ { steps[i] = 1 << 31 } // fill steps
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++  {
            if queue[i] + a < len(steps) && !set[queue[i] + a] {
                if steps[queue[i]] + 1 < steps[queue[i] + a] {
                    queue = append(queue, queue[i] + a)
                    steps[queue[i] + a] = steps[queue[i]] + 1
                }
            }
            if queue[i] - b > 0 && !set[queue[i] - b] {
                if steps[queue[i]] + 1 < steps[queue[i] - b ] {
                    steps[queue[i] - b ] = steps[queue[i]] + 1
                }
                if queue[i] - b + a < len(steps) && !set[queue[i] - b + a] {
                    if steps[queue[i]] + 2 < steps[queue[i] - b + a] {
                        queue = append(queue, queue[i] - b + a)
                        steps[queue[i] - b + a] = steps[queue[i]] + 2
                    }
                }
            }
        }
        queue = queue[n:] // pop
    }
    if steps[x] == 1 << 31 { return -1 }
    return steps[x]
}

func minimumJumps1(forbidden []int, a int, b int, x int) int {
    visited := make([]bool, 6000)
    for i := range forbidden { visited[forbidden[i]] = true }
    queue, res := [][2]int{{0, 0}}, -1
    for len(queue) > 0 {
        n := len(queue)
        res++
        for i := 0; i < n; i++ {
            curr, isBack := queue[i][0], queue[i][1]
            if curr == x { return res }
            if isBack == 0 && curr - b > 0 && !visited[curr - b] {
                visited[curr - b] = true
                queue = append(queue, [2]int{curr - b, 1})
            }
            if curr + a < len(visited) && !visited[curr + a] {
                visited[curr + a] = true
                queue = append(queue, [2]int{curr + a, 0})
            }
        }
        queue = queue[n:] // pop
    }
    return -1
}

func main() {
    // Example 1:
    // Input: forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
    // Output: 3
    // Explanation: 3 jumps forward (0 -> 3 -> 6 -> 9) will get the bug home.
    fmt.Println(minimumJumps([]int{14,4,18,1,15}, 3, 15, 9)) // 3
    // Example 2:
    // Input: forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
    // Output: -1
    fmt.Println(minimumJumps([]int{8,3,16,6,12,20}, 15, 13, 11)) // -1
    // Example 3:
    // Input: forbidden = [1,6,2,14,5,17,4], a = 16, b = 9, x = 7
    // Output: 2
    // Explanation: One jump forward (0 -> 16) then one jump backward (16 -> 7) will get the bug home.
    fmt.Println(minimumJumps([]int{1,6,2,14,5,17,4}, 16, 9, 7)) // 2

    fmt.Println(minimumJumps([]int{1998}, 1999, 2000, 2000)) // 3998

    fmt.Println(minimumJumps1([]int{14,4,18,1,15}, 3, 15, 9)) // 3
    fmt.Println(minimumJumps1([]int{8,3,16,6,12,20}, 15, 13, 11)) // -1
    fmt.Println(minimumJumps1([]int{1,6,2,14,5,17,4}, 16, 9, 7)) // 2
    fmt.Println(minimumJumps1([]int{1998}, 1999, 2000, 2000)) // 3998
}