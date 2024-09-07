package main

// 2545. Sort the Students by Their Kth Score
// There is a class with m students and n exams. 
// You are given a 0-indexed m x n integer matrix score, 
// where each row represents one student and score[i][j] denotes the score the ith student got in the jth exam. 
// The matrix score contains distinct integers only.

// You are also given an integer k. 
// Sort the students (i.e., the rows of the matrix) by their scores in the kth (0-indexed) exam from the highest to the lowest.

// Return the matrix after sorting it.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/11/30/example1.png" />
// Input: score = [[10,6,9,1],[7,5,11,2],[4,8,3,15]], k = 2
// Output: [[7,5,11,2],[10,6,9,1],[4,8,3,15]]
// Explanation: In the above diagram, S denotes the student, while E denotes the exam.
// - The student with index 1 scored 11 in exam 2, which is the highest score, so they got first place.
// - The student with index 0 scored 9 in exam 2, which is the second highest score, so they got second place.
// - The student with index 2 scored 3 in exam 2, which is the lowest score, so they got third place.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/11/30/example2.png" />
// Input: score = [[3,4],[5,6]], k = 0
// Output: [[5,6],[3,4]]
// Explanation: In the above diagram, S denotes the student, while E denotes the exam.
// - The student with index 1 scored 5 in exam 0, which is the highest score, so they got first place.
// - The student with index 0 scored 3 in exam 0, which is the lowest score, so they got second place.

// Constraints:
//     m == score.length
//     n == score[i].length
//     1 <= m, n <= 250
//     1 <= score[i][j] <= 10^5
//     score consists of distinct integers.
//     0 <= k < n

import "fmt"
import "sort"

func sortTheStudents(score [][]int, k int) [][]int {
    arr := []int{}
    for _, v := range score { // 取出每行第 k 个 
        arr = append(arr, v[k])
    }
    sort.Ints(arr)
    n, i := len(arr), 0
    res := make([][]int, n)
    for i < n {
        for _, v := range score {
            if v[k] == arr[i] { // 找到第 k 行一样的
                res[n - i - 1] = v
                break
            }
        }
        i++
    }
    return res
}

func sortTheStudents1(score [][]int, k int) [][]int {
    sort.Slice(score, func(i, j int) bool {
        return score[i][k] > score[j][k]
    })
    return score
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/11/30/example1.png" />
    // Input: score = [[10,6,9,1],[7,5,11,2],[4,8,3,15]], k = 2
    // Output: [[7,5,11,2],[10,6,9,1],[4,8,3,15]]
    // Explanation: In the above diagram, S denotes the student, while E denotes the exam.
    // - The student with index 1 scored 11 in exam 2, which is the highest score, so they got first place.
    // - The student with index 0 scored 9 in exam 2, which is the second highest score, so they got second place.
    // - The student with index 2 scored 3 in exam 2, which is the lowest score, so they got third place.
    score1 := [][]int{
        {10,6,9,1},
        {7,5,11,2},
        {4,8,3,15},
    }
    fmt.Println(sortTheStudents(score1, 2)) // [[7,5,11,2],[10,6,9,1],[4,8,3,15]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/11/30/example2.png" />
    // Input: score = [[3,4],[5,6]], k = 0
    // Output: [[5,6],[3,4]]
    // Explanation: In the above diagram, S denotes the student, while E denotes the exam.
    // - The student with index 1 scored 5 in exam 0, which is the highest score, so they got first place.
    // - The student with index 0 scored 3 in exam 0, which is the lowest score, so they got second place.
    score2 := [][]int{
        {3,4},
        {5,6},
    }
    fmt.Println(sortTheStudents(score2, 0)) // [[5,6],[3,4]]

    fmt.Println(sortTheStudents1(score1, 2)) // [[7,5,11,2],[10,6,9,1],[4,8,3,15]]
    fmt.Println(sortTheStudents1(score2, 0)) // [[5,6],[3,4]]
}