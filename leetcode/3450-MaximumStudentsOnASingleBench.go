package main 

// 3450. Maximum Students on a Single Bench
// You are given a 2D integer array of student data students,
// where students[i] = [student_id, bench_id] represents that student student_id is sitting on the bench bench_id.

// Return the maximum number of unique students sitting on any single bench. 
// If no students are present, return 0.

// Note: A student can appear multiple times on the same bench in the input, but they should be counted only once per bench.

// Example 1:
// Input: students = [[1,2],[2,2],[3,3],[1,3],[2,3]]
// Output: 3
// Explanation:
// Bench 2 has two unique students: [1, 2].
// Bench 3 has three unique students: [1, 2, 3].
// The maximum number of unique students on a single bench is 3.

// Example 2:
// Input: students = [[1,1],[2,1],[3,1],[4,2],[5,2]]
// Output: 3
// Explanation:
// Bench 1 has three unique students: [1, 2, 3].
// Bench 2 has two unique students: [4, 5].
// The maximum number of unique students on a single bench is 3.

// Example 3:
// Input: students = [[1,1],[1,1]]
// Output: 1
// Explanation:
// The maximum number of unique students on a single bench is 1.

// Example 4:
// Input: students = []
// Output: 0
// Explanation:
// Since no students are present, the output is 0.

// Constraints:
//     0 <= students.length <= 100
//     students[i] = [student_id, bench_id]
//     1 <= student_id <= 100
//     1 <= bench_id <= 100

import "fmt"

func maxStudentsOnBench(students [][]int) int {
    mp := make(map[int]map[int]int)
    for _, v:= range students {
        a, b := v[0], v[1]
        if _, ok := mp[b]; !ok {
            mp[b] = make(map[int]int)
        }
        mp[b][a]++
    }
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _,v := range mp {
        res = max(res, len(v))
    }
    return res
}

func main() {
    // Example 1:
    // Input: students = [[1,2],[2,2],[3,3],[1,3],[2,3]]
    // Output: 3
    // Explanation:
    // Bench 2 has two unique students: [1, 2].
    // Bench 3 has three unique students: [1, 2, 3].
    // The maximum number of unique students on a single bench is 3.
    fmt.Println(maxStudentsOnBench([][]int{{1,2},{2,2},{3,3},{1,3},{2,3}})) // 3
    // Example 2:
    // Input: students = [[1,1],[2,1],[3,1],[4,2],[5,2]]
    // Output: 3
    // Explanation:
    // Bench 1 has three unique students: [1, 2, 3].
    // Bench 2 has two unique students: [4, 5].
    // The maximum number of unique students on a single bench is 3.
    fmt.Println(maxStudentsOnBench([][]int{{1,1},{2,1},{3,1},{4,2},{5,2}})) // 3
    // Example 3:
    // Input: students = [[1,1],[1,1]]
    // Output: 1
    // Explanation:
    // The maximum number of unique students on a single bench is 1.
    fmt.Println(maxStudentsOnBench([][]int{{1,1},{1,1}})) // 1
    // Example 4:
    // Input: students = []
    // Output: 0
    // Explanation:
    // Since no students are present, the output is 0.
    fmt.Println(maxStudentsOnBench([][]int{})) // 0
}