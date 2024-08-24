package main

// 3044. Most Frequent Prime
// You are given a m x n 0-indexed 2D matrix mat. From every cell, you can create numbers in the following way:
//     There could be at most 8 paths from the cells namely: 
//         east, south-east, south, south-west, west, north-west, north, and north-east.
//     Select a path from them and append digits in this path to the number being formed by traveling in this direction.
//     Note that numbers are generated at every step, 
//         for example, if the digits along the path are 1, 9, 1, 
//         then there will be three numbers generated along the way: 1, 19, 191.

// Return the most frequent prime number greater than 10 out of all the numbers created by traversing the matrix or -1 if no such prime number exists. 
// If there are multiple prime numbers with the highest frequency, then return the largest among them.

// Note: It is invalid to change the direction during the move.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/02/15/south" />
// Input: mat = [[1,1],[9,9],[1,1]]
// Output: 19
// Explanation: 
// From cell (0,0) there are 3 possible directions and the numbers greater than 10 which can be created in those directions are:
// East: [11], South-East: [19], South: [19,191].
// Numbers greater than 10 created from the cell (0,1) in all possible directions are: [19,191,19,11].
// Numbers greater than 10 created from the cell (1,0) in all possible directions are: [99,91,91,91,91].
// Numbers greater than 10 created from the cell (1,1) in all possible directions are: [91,91,99,91,91].
// Numbers greater than 10 created from the cell (2,0) in all possible directions are: [11,19,191,19].
// Numbers greater than 10 created from the cell (2,1) in all possible directions are: [11,19,19,191].
// The most frequent prime number among all the created numbers is 19.

// Example 2:
// Input: mat = [[7]]
// Output: -1
// Explanation: The only number which can be formed is 7. It is a prime number however it is not greater than 10, so return -1.

// Example 3:
// Input: mat = [[9,7,8],[4,6,5],[2,8,6]]
// Output: 97
// Explanation: 
// Numbers greater than 10 created from the cell (0,0) in all possible directions are: [97,978,96,966,94,942].
// Numbers greater than 10 created from the cell (0,1) in all possible directions are: [78,75,76,768,74,79].
// Numbers greater than 10 created from the cell (0,2) in all possible directions are: [85,856,86,862,87,879].
// Numbers greater than 10 created from the cell (1,0) in all possible directions are: [46,465,48,42,49,47].
// Numbers greater than 10 created from the cell (1,1) in all possible directions are: [65,66,68,62,64,69,67,68].
// Numbers greater than 10 created from the cell (1,2) in all possible directions are: [56,58,56,564,57,58].
// Numbers greater than 10 created from the cell (2,0) in all possible directions are: [28,286,24,249,26,268].
// Numbers greater than 10 created from the cell (2,1) in all possible directions are: [86,82,84,86,867,85].
// Numbers greater than 10 created from the cell (2,2) in all possible directions are: [68,682,66,669,65,658].
// The most frequent prime number among all the created numbers is 97.

// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 6
//     1 <= mat[i][j] <= 9

import "fmt"

func mostFrequentPrime(mat [][]int) int {
    dirs := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
    m, n := len(mat), len(mat[0])
    cnt := map[int]int{}
    isPrime := func (n int) bool {
        for i := 2; i*i <= n; i++ {
            if n % i == 0 {
                return false
            }
        }
        return true
    }
    for i, row := range mat {
        for j, v := range row {
            for _, d := range dirs {
                x, y, v := i+d.x, j+d.y, v
                for 0 <= x && x < m && 0 <= y && y < n {
                    v = v * 10 + mat[x][y]
                    if cnt[v] > 0 || isPrime(v) { // 如果 v 在 cnt 中，那么 v 一定是质数
                        cnt[v]++
                    }
                    x += d.x
                    y += d.y
                }
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, mx := -1, 0
    for k, v := range cnt {
        if v > mx {
            res, mx = k, v
        } else if v == mx {
            res = max(res, k)
        }
    }
    return res
}

func mostFrequentPrime1(mat [][]int) int {
    primes := map[int]int{-1: 0}
    isPrime := func(x int) bool {
        if primes[x] > 0 {
            return true
        }
        for i := 2; i*i <= x; i++ {
            if x%i == 0 {
                return false
            }
        }
        return true
    }
    dirs := [][]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}
    res, m, n := -1, len(mat), len(mat[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            for _, dir := range dirs {
                p, q := i, j
                cur := mat[p][q]
                for {
                    np, nq := p+dir[0], q+dir[1]
                    if np >= 0 && np < m && nq >= 0 && nq < n {
                        p, q = np, nq
                    } else {
                        break
                    }
                    cur = cur*10 + mat[p][q]
                    if isPrime(cur) {
                        primes[cur]++
                        if primes[cur] > primes[res] || (primes[cur] == primes[res] && cur > res) {
                            res = cur
                        }
                    }
                }
            }
        }
    }
    return res
}

func mostFrequentPrime2(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
    freq := make(map[int]int)
    isPrime := func(num int) bool {
        if num < 2 {
            return false
        }
        for i := 2; i*i <= num; i++ {
            if num%i == 0 {
                return false
            }
        }
        return true
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            for _, dir := range dirs {
                x, y := i, j
                num := 0
                for x >= 0 && x < m && y >= 0 && y < n {
                    num = num*10 + mat[x][y]
                    if num > 10 && isPrime(num) {
                        freq[num]++
                    }
                    x += dir[0]
                    y += dir[1]
               }
            }
        }
    }
    maxFreq, mostFreqPrime := 0, -1
    for prime, v := range freq {
        if v > maxFreq || (v == maxFreq && prime > mostFreqPrime) {
            maxFreq, mostFreqPrime = v, prime
        }
    }
    return mostFreqPrime
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/02/15/south" />
    // Input: mat = [[1,1],[9,9],[1,1]]
    // Output: 19
    // Explanation: 
    // From cell (0,0) there are 3 possible directions and the numbers greater than 10 which can be created in those directions are:
    // East: [11], South-East: [19], South: [19,191].
    // Numbers greater than 10 created from the cell (0,1) in all possible directions are: [19,191,19,11].
    // Numbers greater than 10 created from the cell (1,0) in all possible directions are: [99,91,91,91,91].
    // Numbers greater than 10 created from the cell (1,1) in all possible directions are: [91,91,99,91,91].
    // Numbers greater than 10 created from the cell (2,0) in all possible directions are: [11,19,191,19].
    // Numbers greater than 10 created from the cell (2,1) in all possible directions are: [11,19,19,191].
    // The most frequent prime number among all the created numbers is 19.
    fmt.Println(mostFrequentPrime([][]int{{1,1},{9,9},{1,1}})) // 19
    // Example 2:
    // Input: mat = [[7]]
    // Output: -1
    // Explanation: The only number which can be formed is 7. It is a prime number however it is not greater than 10, so return -1.
    fmt.Println(mostFrequentPrime([][]int{{7}})) // -1
    // Example 3:
    // Input: mat = [[9,7,8],[4,6,5],[2,8,6]]
    // Output: 97
    // Explanation: 
    // Numbers greater than 10 created from the cell (0,0) in all possible directions are: [97,978,96,966,94,942].
    // Numbers greater than 10 created from the cell (0,1) in all possible directions are: [78,75,76,768,74,79].
    // Numbers greater than 10 created from the cell (0,2) in all possible directions are: [85,856,86,862,87,879].
    // Numbers greater than 10 created from the cell (1,0) in all possible directions are: [46,465,48,42,49,47].
    // Numbers greater than 10 created from the cell (1,1) in all possible directions are: [65,66,68,62,64,69,67,68].
    // Numbers greater than 10 created from the cell (1,2) in all possible directions are: [56,58,56,564,57,58].
    // Numbers greater than 10 created from the cell (2,0) in all possible directions are: [28,286,24,249,26,268].
    // Numbers greater than 10 created from the cell (2,1) in all possible directions are: [86,82,84,86,867,85].
    // Numbers greater than 10 created from the cell (2,2) in all possible directions are: [68,682,66,669,65,658].
    // The most frequent prime number among all the created numbers is 97.
    fmt.Println(mostFrequentPrime([][]int{{9,7,8},{4,6,5},{2,8,6}})) // 97

    fmt.Println(mostFrequentPrime1([][]int{{1,1},{9,9},{1,1}})) // 19
    fmt.Println(mostFrequentPrime1([][]int{{7}})) // -1
    fmt.Println(mostFrequentPrime1([][]int{{9,7,8},{4,6,5},{2,8,6}})) // 97

    fmt.Println(mostFrequentPrime2([][]int{{1,1},{9,9},{1,1}})) // 19
    fmt.Println(mostFrequentPrime2([][]int{{7}})) // -1
    fmt.Println(mostFrequentPrime2([][]int{{9,7,8},{4,6,5},{2,8,6}})) // 97
}