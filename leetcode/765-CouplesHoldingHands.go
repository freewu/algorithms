package main

// 765. Couples Holding Hands
// There are n couples sitting in 2n seats arranged in a row and want to hold hands.

// The people and seats are represented by an integer array row where row[i] is the ID of the person sitting in the ith seat. 
// The couples are numbered in order, the first couple being (0, 1), the second couple being (2, 3), and so on with the last couple being (2n - 2, 2n - 1).

// Return the minimum number of swaps so that every couple is sitting side by side. 
// A swap consists of choosing any two people, then they stand up and switch seats.

// Example 1:
// Input: row = [0,2,1,3]
// Output: 1
// Explanation: We only need to swap the second (row[1]) and third (row[2]) person.

// Example 2:
// Input: row = [3,2,0,1]
// Output: 0
// Explanation: All couples are already seated side by side.

// Constraints:
//     2n == row.length
//     2 <= n <= 30
//     n is even.
//     0 <= row[i] < 2n
//     All the elements of row are unique.

import "fmt"

// greedy
func minSwapsCouples(row []int) int {
    res := 0
    matchNumber := func (n int) int { if n % 2 == 0 { return n + 1; }; return n - 1; }
    match := func (a, b int) bool { return matchNumber(a) == b; }
    swap := func (row []int, i, j int) { row[i], row[j] = row[j], row[i]; }
    for i := 0; i < len(row); i += 2 {
        if !match(row[i], row[i+1]) {
            for j := i + 2; j < len(row); j++ {
                if match(row[i], row[j]) {
                    // check if left side of pair seat can fix two couples
                    if j % 2 == 0 && match(row[i+1], row[j+1]) {
                        swap(row, i + 1, j)
                        res++
                    }
                    if j % 2 == 1 && match(row[i+1], row[j-1]) {
                        swap(row, i + 1, j)
                        res++
                    }
                }
            }
            for j := i + 2; j < len(row); j++ { // otherwize, swap right side seat
                if match(row[i+1], row[j]) {
                    swap(row, i, j)
                    res++
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: row = [0,2,1,3]
    // Output: 1
    // Explanation: We only need to swap the second (row[1]) and third (row[2]) person.
    fmt.Println(minSwapsCouples([]int{0,2,1,3})) // 1
    // Example 2:
    // Input: row = [3,2,0,1]
    // Output: 0
    // Explanation: All couples are already seated side by side.
    fmt.Println(minSwapsCouples([]int{3,2,0,1})) // 0
}