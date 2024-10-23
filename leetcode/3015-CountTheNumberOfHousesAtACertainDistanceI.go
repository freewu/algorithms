package main

// 3015. Count the Number of Houses at a Certain Distance I
// You are given three positive integers n, x, and y.

// In a city, there exist houses numbered 1 to n connected by n streets. 
// There is a street connecting the house numbered i with the house numbered i + 1 for all 1 <= i <= n - 1 . 
// An additional street connects the house numbered x with the house numbered y.

// For each k, such that 1 <= k <= n, you need to find the number of pairs of houses (house1, house2) 
// such that the minimum number of streets that need to be traveled to reach house2 from house1 is k.

// Return a 1-indexed array result of length n where result[k] represents the total number of pairs of houses 
// such that the minimum streets required to reach one house from the other is k.

// Note that x and y can be equal.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/12/20/example2.png" />
// Input: n = 3, x = 1, y = 3
// Output: [6,0,0]
// Explanation: Let's look at each pair of houses:
// - For the pair (1, 2), we can go from house 1 to house 2 directly.
// - For the pair (2, 1), we can go from house 2 to house 1 directly.
// - For the pair (1, 3), we can go from house 1 to house 3 directly.
// - For the pair (3, 1), we can go from house 3 to house 1 directly.
// - For the pair (2, 3), we can go from house 2 to house 3 directly.
// - For the pair (3, 2), we can go from house 3 to house 2 directly.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/12/20/example3.png" />
// Input: n = 5, x = 2, y = 4
// Output: [10,8,2,0,0]
// Explanation: For each distance k the pairs are:
// - For k == 1, the pairs are (1, 2), (2, 1), (2, 3), (3, 2), (2, 4), (4, 2), (3, 4), (4, 3), (4, 5), and (5, 4).
// - For k == 2, the pairs are (1, 3), (3, 1), (1, 4), (4, 1), (2, 5), (5, 2), (3, 5), and (5, 3).
// - For k == 3, the pairs are (1, 5), and (5, 1).
// - For k == 4 and k == 5, there are no pairs.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2023/12/20/example5.png" />
// Input: n = 4, x = 1, y = 1
// Output: [6,4,2,0]
// Explanation: For each distance k the pairs are:
// - For k == 1, the pairs are (1, 2), (2, 1), (2, 3), (3, 2), (3, 4), and (4, 3).
// - For k == 2, the pairs are (1, 3), (3, 1), (2, 4), and (4, 2).
// - For k == 3, the pairs are (1, 4), and (4, 1).
// - For k == 4, there are no pairs.

// Constraints:
//     2 <= n <= 100
//     1 <= x, y <= n

import "fmt"

// bfs
func countOfPairs(n int, x int, y int) []int {
    graph := make([][]int, n + 1)
    for i := 1; i < n; i++ {
        graph[i]   = append(graph[i], i+1)
        graph[i+1] = append(graph[i+1], i)
    }
    graph[x] = append(graph[x], y)
    graph[y] = append(graph[y], x)
    res := make([]int, n)
    for i := 1; i <= n; i++ {
        dist := make([]int, n+1)
        for i := range dist { 
            dist[i] = 1 << 31
        }
        queue := []int{ i }
        dist[i] = 0
        for len(queue) > 0 {
            u := queue[0]
            du := dist[u]
            queue = queue[1:]
            for _, v := range graph[u] {
                if dist[v] > du + 1 {
                    dist[v] = du + 1
                    res[dist[v]-1]++
                    queue = append(queue, v)
                }                
            }
        }
    }
    return res
}

func countOfPairs1(n int, x int, y int) []int {
    res := make([]int, n)
    if x > y {
        x, y = y, x
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i := 1; i <= n; i++ {
        for j := i + 1; j <= n; j++ {
            v := min(j - i, 1 + abs(i - x) + abs(j - y))
            res[v - 1] += 2
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/12/20/example2.png" />
    // Input: n = 3, x = 1, y = 3
    // Output: [6,0,0]
    // Explanation: Let's look at each pair of houses:
    // - For the pair (1, 2), we can go from house 1 to house 2 directly.
    // - For the pair (2, 1), we can go from house 2 to house 1 directly.
    // - For the pair (1, 3), we can go from house 1 to house 3 directly.
    // - For the pair (3, 1), we can go from house 3 to house 1 directly.
    // - For the pair (2, 3), we can go from house 2 to house 3 directly.
    // - For the pair (3, 2), we can go from house 3 to house 2 directly.
    fmt.Println(countOfPairs(3, 1, 3))  // [6,0,0]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/12/20/example3.png" />
    // Input: n = 5, x = 2, y = 4
    // Output: [10,8,2,0,0]
    // Explanation: For each distance k the pairs are:
    // - For k == 1, the pairs are (1, 2), (2, 1), (2, 3), (3, 2), (2, 4), (4, 2), (3, 4), (4, 3), (4, 5), and (5, 4).
    // - For k == 2, the pairs are (1, 3), (3, 1), (1, 4), (4, 1), (2, 5), (5, 2), (3, 5), and (5, 3).
    // - For k == 3, the pairs are (1, 5), and (5, 1).
    // - For k == 4 and k == 5, there are no pairs.
    fmt.Println(countOfPairs(5, 2, 4)) // [10,8,2,0,0]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2023/12/20/example5.png" />
    // Input: n = 4, x = 1, y = 1
    // Output: [6,4,2,0]
    // Explanation: For each distance k the pairs are:
    // - For k == 1, the pairs are (1, 2), (2, 1), (2, 3), (3, 2), (3, 4), and (4, 3).
    // - For k == 2, the pairs are (1, 3), (3, 1), (2, 4), and (4, 2).
    // - For k == 3, the pairs are (1, 4), and (4, 1).
    // - For k == 4, there are no pairs.
    fmt.Println(countOfPairs(4, 1, 1))  // [6,4,2,0]

    fmt.Println(countOfPairs(2, 1, 1))  // [2 0]
    fmt.Println(countOfPairs(2, 2, 2))  // [2 0]
    fmt.Println(countOfPairs(100, 2, 2))  // [198 196 194, ... ]
    fmt.Println(countOfPairs(100, 1, 1))  // [198 196 194 192, ... ]
    fmt.Println(countOfPairs(100, 100, 100))  // [198 196 194 192 190, ...]

    fmt.Println(countOfPairs1(3, 1, 3))  // [6,0,0]
    fmt.Println(countOfPairs1(5, 2, 4)) // [10,8,2,0,0]
    fmt.Println(countOfPairs1(4, 1, 1))  // [6,4,2,0]
    fmt.Println(countOfPairs1(2, 1, 1))  // [2 0]
    fmt.Println(countOfPairs1(2, 2, 2))  // [2 0]
    fmt.Println(countOfPairs1(100, 2, 2))  // [198 196 194, ... ]
    fmt.Println(countOfPairs1(100, 1, 1))  // [198 196 194 192, ... ]
    fmt.Println(countOfPairs1(100, 100, 100))  // [198 196 194 192 190, ...]
}