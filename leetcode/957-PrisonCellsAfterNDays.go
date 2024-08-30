package main

// 957. Prison Cells After N Days
// There are 8 prison cells in a row and each cell is either occupied or vacant.

// Each day, whether the cell is occupied or vacant changes according to the following rules:
//     If a cell has two adjacent neighbors that are both occupied or both vacant, then the cell becomes occupied.
//     Otherwise, it becomes vacant.

// Note that because the prison is a row, the first and the last cells in the row can't have two adjacent neighbors.

// You are given an integer array cells where cells[i] == 1 if the ith cell is occupied and cells[i] == 0 if the ith cell is vacant, and you are given an integer n.

// Return the state of the prison after n days (i.e., n such changes described above).

// Example 1:
// Input: cells = [0,1,0,1,1,0,0,1], n = 7
// Output: [0,0,1,1,0,0,0,0]
// Explanation: The following table summarizes the state of the prison on each day:
// Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
// Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
// Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
// Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
// Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
// Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
// Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
// Day 7: [0, 0, 1, 1, 0, 0, 0, 0]

// Example 2:
// Input: cells = [1,0,0,1,0,0,1,0], n = 1000000000
// Output: [0,0,1,1,1,1,1,0]

// Constraints:
//     cells.length == 8
//     cells[i] is either 0 or 1.
//     1 <= n <= 10^9

import "fmt"

func prisonAfterNDays(cells []int, n int) []int {
    arr := [8]int{}
    for i, v := range cells { 
        arr[i] = v 
    }
    nextDay := func(cells [8]int) [8]int {
        t := [8]int{}
        for i := 1; i < 7; i++ {
            if cells[i - 1] == cells[i + 1] { // 如果一间牢房的两个相邻的房间都被占用或都是空的，那么该牢房就会被占用
                t[i] = 1
            } else { // 否则，它就会被空置
                t[i] = 0
            }
        }
        return t
    }
    t, count, hasCycle := make(map[[8]int]bool), 0, false
    for i := 0; i < n; i++ {
        next := nextDay(arr)
        if _, ok := t[next]; !ok {
            t[next] = true
            count++
            arr = next
        } else {
            hasCycle = true
            break
        }
    }
    if hasCycle { 
        for i := 0; i < n % count; i++ { 
            arr = nextDay(arr) 
        } 
    }
    return arr[:]
}

func main() {
    // Example 1:
    // Input: cells = [0,1,0,1,1,0,0,1], n = 7
    // Output: [0,0,1,1,0,0,0,0]
    // Explanation: The following table summarizes the state of the prison on each day:
    // Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
    // Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
    // Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
    // Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
    // Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
    // Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
    // Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
    // Day 7: [0, 0, 1, 1, 0, 0, 0, 0]
    fmt.Println(prisonAfterNDays([]int{0,1,0,1,1,0,0,1}, 7)) // [0,0,1,1,0,0,0,0]
    // Example 2:
    // Input: cells = [1,0,0,1,0,0,1,0], n = 1000000000
    // Output: [0,0,1,1,1,1,1,0]
    fmt.Println(prisonAfterNDays([]int{1,0,0,1,0,0,1,0}, 1000000000)) // [0,0,1,1,1,1,1,0]
}