package main

// 2355. Maximum Number of Books You Can Take
// You are given a 0-indexed integer array books of length n where books[i] denotes the number of books on the ith shelf of a bookshelf.

// You are going to take books from a contiguous section of the bookshelf spanning from l to r where 0 <= l <= r < n. 
// For each index i in the range l <= i < r, you must take strictly fewer books from shelf i than shelf i + 1.

// Return the maximum number of books you can take from the bookshelf.

// Example 1:
// Input: books = [8,5,2,7,9]
// Output: 19
// Explanation:
// - Take 1 book from shelf 1.
// - Take 2 books from shelf 2.
// - Take 7 books from shelf 3.
// - Take 9 books from shelf 4.
// You have taken 19 books, so return 19.
// It can be proven that 19 is the maximum number of books you can take.

// Example 2:
// Input: books = [7,0,3,4,5]
// Output: 12
// Explanation:
// - Take 3 books from shelf 2.
// - Take 4 books from shelf 3.
// - Take 5 books from shelf 4.
// You have taken 12 books so return 12.
// It can be proven that 12 is the maximum number of books you can take.

// Example 3:
// Input: books = [8,2,3,7,3,4,0,1,4,3]
// Output: 13
// Explanation:
// - Take 1 book from shelf 0.
// - Take 2 books from shelf 1.
// - Take 3 books from shelf 2.
// - Take 7 books from shelf 3.
// You have taken 13 books so return 13.
// It can be proven that 13 is the maximum number of books you can take.

// Constraints:
//     1 <= books.length <= 10^5
//     0 <= books[i] <= 10^5

import "fmt"

func maximumBooks(books []int) int64 {
    type Pair struct{ Index, Sum int }
    res, stack := 0, []Pair{{-1, 0}} // 单调栈里面存 (下标, dp 值)，加个哨兵
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, v := range books {
        for len(stack) > 1 && books[stack[len(stack) - 1].Index] - stack[len(stack) - 1].Index >= v - i {
            stack = stack[:len(stack)-1] // pop
        }
        n := min(i - stack[len(stack) - 1].Index, v)
        sum := (v * 2 - n + 1) * n / 2 + stack[len(stack) - 1].Sum
        res = max(res, sum)
        stack = append(stack, Pair{ i, sum }) // push
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: books = [8,5,2,7,9]
    // Output: 19
    // Explanation:
    // - Take 1 book from shelf 1.
    // - Take 2 books from shelf 2.
    // - Take 7 books from shelf 3.
    // - Take 9 books from shelf 4.
    // You have taken 19 books, so return 19.
    // It can be proven that 19 is the maximum number of books you can take.
    fmt.Println(maximumBooks([]int{8,5,2,7,9})) // 19
    // Example 2:
    // Input: books = [7,0,3,4,5]
    // Output: 12
    // Explanation:
    // - Take 3 books from shelf 2.
    // - Take 4 books from shelf 3.
    // - Take 5 books from shelf 4.
    // You have taken 12 books so return 12.
    // It can be proven that 12 is the maximum number of books you can take.
    fmt.Println(maximumBooks([]int{7,0,3,4,5})) // 12
    // Example 3:
    // Input: books = [8,2,3,7,3,4,0,1,4,3]
    // Output: 13
    // Explanation:
    // - Take 1 book from shelf 0.
    // - Take 2 books from shelf 1.
    // - Take 3 books from shelf 2.
    // - Take 7 books from shelf 3.
    // You have taken 13 books so return 13.
    // It can be proven that 13 is the maximum number of books you can take.
    fmt.Println(maximumBooks([]int{8,2,3,7,3,4,0,1,4,3})) // 13
}