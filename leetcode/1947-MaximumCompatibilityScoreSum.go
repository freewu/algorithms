package main

// 1947. Maximum Compatibility Score Sum
// There is a survey that consists of n questions where each question's answer is either 0 (no) or 1 (yes).

// The survey was given to m students numbered from 0 to m - 1 and m mentors numbered from 0 to m - 1. 
// The answers of the students are represented by a 2D integer array students where students[i] is an integer array 
// that contains the answers of the ith student (0-indexed). 
// The answers of the mentors are represented by a 2D integer array mentors where mentors[j] is an integer array 
// that contains the answers of the jth mentor (0-indexed).

// Each student will be assigned to one mentor, and each mentor will have one student assigned to them. 
// The compatibility score of a student-mentor pair is the number of answers that are the same for both the student and the mentor.
//     For example, if the student's answers were [1, 0, 1] and the mentor's answers were [0, 0, 1], 
//     then their compatibility score is 2 because only the second and the third answers are the same.

// You are tasked with finding the optimal student-mentor pairings to maximize the sum of the compatibility scores.

// Given students and mentors, return the maximum compatibility score sum that can be achieved.

// Example 1:
// Input: students = [[1,1,0],[1,0,1],[0,0,1]], mentors = [[1,0,0],[0,0,1],[1,1,0]]
// Output: 8
// Explanation: We assign students to mentors in the following way:
// - student 0 to mentor 2 with a compatibility score of 3.
// - student 1 to mentor 0 with a compatibility score of 2.
// - student 2 to mentor 1 with a compatibility score of 3.
// The compatibility score sum is 3 + 2 + 3 = 8.

// Example 2:
// Input: students = [[0,0],[0,0],[0,0]], mentors = [[1,1],[1,1],[1,1]]
// Output: 0
// Explanation: The compatibility score of any student-mentor pair is 0.

// Constraints:
//     m == students.length == mentors.length
//     n == students[i].length == mentors[j].length
//     1 <= m, n <= 8
//     students[i][k] is either 0 or 1.
//     mentors[j][k] is either 0 or 1.

import "fmt"

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
    res, n := 0, len(students)
    visited := make([]bool, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    getCompatibility := func(arr1, arr2 []int) int {
        res := 0
        for i := 0; i < len(arr1); i++ {
            if arr1[i] == arr2[i] { res++ }
        }
        return res
    }
    var backtrack func(index int, currentCompatibility int)
    backtrack = func(index int, currentCompatibility int) {
        if index == n {
            res = max(res, currentCompatibility)
        } else {
            for i := 0; i < n; i++ {
                if !visited[i]{
                    visited[i] = true
                    backtrack(index + 1, currentCompatibility + getCompatibility(students[index], mentors[i])) 
                    visited[i] = false
                }  
            }
        }
    }
    backtrack(0, 0)
    return res
}

func main() {
    // Example 1:
    // Input: students = [[1,1,0],[1,0,1],[0,0,1]], mentors = [[1,0,0],[0,0,1],[1,1,0]]
    // Output: 8
    // Explanation: We assign students to mentors in the following way:
    // - student 0 to mentor 2 with a compatibility score of 3.
    // - student 1 to mentor 0 with a compatibility score of 2.
    // - student 2 to mentor 1 with a compatibility score of 3.
    // The compatibility score sum is 3 + 2 + 3 = 8.
    fmt.Println(maxCompatibilitySum([][]int{{1,1,0},{1,0,1},{0,0,1}}, [][]int{{1,0,0},{0,0,1},{1,1,0}})) // 8
    // Example 2:
    // Input: students = [[0,0],[0,0],[0,0]], mentors = [[1,1],[1,1],[1,1]]
    // Output: 0
    // Explanation: The compatibility score of any student-mentor pair is 0.
    fmt.Println(maxCompatibilitySum([][]int{{0,0},{0,0},{0,0}}, [][]int{{1,1},{1,1},{1,1}})) // 0
}