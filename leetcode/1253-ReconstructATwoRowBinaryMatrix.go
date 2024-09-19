package main

// 1253. Reconstruct a 2-Row Binary Matrix
// Given the following details of a matrix with n columns and 2 rows :
//     The matrix is a binary matrix, which means each element in the matrix can be 0 or 1.
//     The sum of elements of the 0-th(upper) row is given as upper.
//     The sum of elements of the 1-st(lower) row is given as lower.
//     The sum of elements in the i-th column(0-indexed) is colsum[i], where colsum is given as an integer array with length n.

// Your task is to reconstruct the matrix with upper, lower and colsum.

// Return it as a 2-D integer array.

// If there are more than one valid solution, any of them will be accepted.

// If no valid solution exists, return an empty 2-D array.

// Example 1:
// Input: upper = 2, lower = 1, colsum = [1,1,1]
// Output: [[1,1,0],[0,0,1]]
// Explanation: [[1,0,1],[0,1,0]], and [[0,1,1],[1,0,0]] are also correct answers.

// Example 2:
// Input: upper = 2, lower = 3, colsum = [2,2,1,1]
// Output: []

// Example 3:
// Input: upper = 5, lower = 5, colsum = [2,1,2,0,1,0,1,2,0,1]
// Output: [[1,1,1,0,1,0,0,1,0,0],[1,0,1,0,0,0,1,1,0,1]]

// Constraints:
//     1 <= colsum.length <= 10^5
//     0 <= upper, lower <= colsum.length
//     0 <= colsum[i] <= 2

import "fmt"

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
    matrixRowUpper, matrixRowLower := make([]int, len(colsum)), make([]int, len(colsum))
    for i, v := range colsum {
        if v == 2 {
            matrixRowUpper[i] = 1
            matrixRowLower[i] = 1
            upper--
            lower--
        }
        if upper < 0 || lower < 0 {
            return [][]int{}
        }
    }
    for i, v := range colsum {
        if v == 1 {
            if upper > 0 {
                matrixRowUpper[i] = 1
                upper--
            } else if lower > 0 {
                matrixRowLower[i] = 1
                lower--
            } else {
                return [][]int{}
            }
        }
    }
    // upper & lower MUST BE 0 here
    if upper != 0 || lower != 0 {
        return [][]int{}
    }
    return [][]int{ matrixRowUpper, matrixRowLower }
}

func reconstructMatrix1(upper int, lower int, colsum []int) [][]int {
    matrix := make([][]int,2)
    matrix[0], matrix[1] = make([] int, len(colsum)), make([] int, len(colsum))
    for i, sum := range colsum {
        if sum == 0 {
            // pass
        } else if sum == 2 {
            matrix[0][i], matrix[1][i] = 1, 1
            upper--
            lower-- 
        } else if upper > lower{
            matrix[0][i] = 1
            upper--
        } else {
            matrix[1][i] = 1
            lower--
        }
    }
    if lower != 0 || upper != 0 {
       return nil
    }
    return matrix    
}

func main() {
    // Example 1:
    // Input: upper = 2, lower = 1, colsum = [1,1,1]
    // Output: [[1,1,0],[0,0,1]]
    // Explanation: [[1,0,1],[0,1,0]], and [[0,1,1],[1,0,0]] are also correct answers.
    fmt.Println(reconstructMatrix(2, 1, []int{1,1,1})) // [[1,1,0],[0,0,1]]
    // Example 2:
    // Input: upper = 2, lower = 3, colsum = [2,2,1,1]
    // Output: []
    fmt.Println(reconstructMatrix(2, 3, []int{2,2,1,1})) // []
    // Example 3:
    // Input: upper = 5, lower = 5, colsum = [2,1,2,0,1,0,1,2,0,1]
    // Output: [[1,1,1,0,1,0,0,1,0,0],[1,0,1,0,0,0,1,1,0,1]]
    fmt.Println(reconstructMatrix(5, 5, []int{2,1,2,0,1,0,1,2,0,1})) // [[1,1,1,0,1,0,0,1,0,0],[1,0,1,0,0,0,1,1,0,1]]

    fmt.Println(reconstructMatrix1(2, 1, []int{1,1,1})) // [[1,1,0],[0,0,1]]
    fmt.Println(reconstructMatrix1(2, 3, []int{2,2,1,1})) // []
    fmt.Println(reconstructMatrix1(5, 5, []int{2,1,2,0,1,0,1,2,0,1})) // [[1,1,1,0,1,0,0,1,0,0],[1,0,1,0,0,0,1,1,0,1]]
}