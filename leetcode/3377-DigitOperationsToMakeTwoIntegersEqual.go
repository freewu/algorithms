package main

// 3377. Digit Operations to Make Two Integers Equal
// You are given two integers n and m that consist of the same number of digits.

// You can perform the following operations any number of times:
//     1. Choose any digit from n that is not 9 and increase it by 1.
//     2. Choose any digit from n that is not 0 and decrease it by 1.

// The integer n must not be a prime number at any point, including its original value and after each operation.

// The cost of a transformation is the sum of all values that n takes throughout the operations performed.

// Return the minimum cost to transform n into m. If it is impossible, return -1.

// Example 1:
// Input: n = 10, m = 12
// Output: 85
// Explanation:
// We perform the following operations:
// Increase the first digit, now n = 20.
// Increase the second digit, now n = 21.
// Increase the second digit, now n = 22.
// Decrease the first digit, now n = 12.

// Example 2:
// Input: n = 4, m = 8
// Output: -1
// Explanation:
// It is impossible to make n equal to m.

// Example 3:
// Input: n = 6, m = 2
// Output: -1
// Explanation: 
// Since 2 is already a prime, we can't make n equal to m.

// Constraints:
//     1 <= n, m < 10^4
//     n and m consist of the same number of digits.

import "fmt"
import "container/heap"
import "strconv"

type MinHeap [][]int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MinHeap) Pop() interface{}   {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

var sieve []bool

func runSieve() {
    sieve = make([]bool, 100000)
    for i := range sieve {
        sieve[i] = true
    }
    sieve[0], sieve[1] = false, false
    for i := 2; i < 100000; i++ {
        if sieve[i] {
            for j := 2 * i; j < 100000; j += i {
                sieve[j] = false
            }
        }
    }
}

func minOperations(n int, m int) int {
    runSieve()
    if sieve[n] || sieve[m] { return -1 }
    solve := func(n int, m int) int {
        pq := &MinHeap{}
        heap.Init(pq)
        heap.Push(pq, []int{n, n})
        visited := make(map[int]bool)
        for pq.Len() > 0 {
            top := heap.Pop(pq).([]int)
            sum, cur := top[0], top[1]
            if visited[cur] { continue }
            visited[cur] = true
            if cur == m { return sum }
            s := []rune( strconv.Itoa(cur))
            for i := 0; i < len(s); i++ {
                c := s[i]
                if s[i] < '9' {
                    s[i]++
                    next, _ := strconv.Atoi(string(s))
                    if !sieve[next] && !visited[next] {
                        heap.Push(pq, []int{sum + next, next})
                    }
                    s[i] = c
                }
                if s[i] > '0' && !(i == 0 && s[i] == '1') {
                    s[i]--
                    next, _ := strconv.Atoi(string(s))
                    if !sieve[next] && !visited[next] {
                        heap.Push(pq, []int{sum + next, next})
                    }
                    s[i] = c
                }
            }
        }
        return -1
    }
    return solve(n, m)
}

func main() {
    // Example 1:
    // Input: n = 10, m = 12
    // Output: 85
    // Explanation:
    // We perform the following operations:
    // Increase the first digit, now n = 20.
    // Increase the second digit, now n = 21.
    // Increase the second digit, now n = 22.
    // Decrease the first digit, now n = 12.
    fmt.Println(minOperations(10, 12)) // 85
    // Example 2:
    // Input: n = 4, m = 8
    // Output: -1
    // Explanation:
    // It is impossible to make n equal to m.
    fmt.Println(minOperations(4, 8)) // -1
    // Example 3:
    // Input: n = 6, m = 2
    // Output: -1
    // Explanation: 
    // Since 2 is already a prime, we can't make n equal to m.
    fmt.Println(minOperations(6, 2)) // -1

    fmt.Println(minOperations(1, 1)) // 1
    fmt.Println(minOperations(10000, 10000)) // 10000
    fmt.Println(minOperations(1, 10000)) // -1
    fmt.Println(minOperations(10000, 1)) // -1
}