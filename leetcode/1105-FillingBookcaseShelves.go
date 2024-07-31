package main

// 1105. Filling Bookcase Shelves
// You are given an array books where books[i] = [thicknessi, heighti] indicates the thickness and height of the ith book.
// You are also given an integer shelfWidth.

// We want to place these books in order onto bookcase shelves that have a total width shelfWidth.

// We choose some of the books to place on this shelf such that the sum of their thickness is less than or equal to shelfWidth, 
// then build another level of the shelf of the bookcase so that the total height of the bookcase has increased by the maximum height of the books we just put down. We repeat this process until there are no more books to place.

// Note that at each step of the above process, 
// the order of the books we place is the same order as the given sequence of books.
//     For example, if we have an ordered list of 5 books, 
//     we might place the first and second book onto the first shelf, the third book on the second shelf, 
//     and the fourth and fifth book on the last shelf.

// Return the minimum possible height that the total bookshelf can be after placing shelves in this manner.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/06/24/shelves.png" />
// Input: books = [[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]], shelfWidth = 4
// Output: 6
// Explanation:
// The sum of the heights of the 3 shelves is 1 + 3 + 2 = 6.
// Notice that book number 2 does not have to be on the first shelf.

// Example 2:
// Input: books = [[1,3],[2,4],[3,2]], shelfWidth = 6
// Output: 4

// Constraints:
//     1 <= books.length <= 1000
//     1 <= thicknessi <= shelfWidth <= 1000
//     1 <= heighti <= 1000

import  "fmt"

// dfs
func minHeightShelves(books [][]int, shelfWidth int) int {
    n, inf := len(books), 1 << 32 - 1
    cache := make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i int) int
    dfs = func(i int) int {
        if i < 0 {
            return 0
        }
        if cache[i]	> 0 {
            return cache[i]
        }
        res := inf
        for j, w, v := i, 0, 0; j >= 0; j-- {
            w += books[j][0]
            if w > shelfWidth {
                break
            }
            v = max(v, books[j][1])
            res = min(res, dfs(j-1)+v)
        }
        cache[i] = res
        return res
    }
    return dfs(n - 1)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/06/24/shelves.png" />
    // Input: books = [[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]], shelfWidth = 4
    // Output: 6
    // Explanation:
    // The sum of the heights of the 3 shelves is 1 + 3 + 2 = 6.
    // Notice that book number 2 does not have to be on the first shelf.
    fmt.Println(minHeightShelves([][]int{{1,1},{2,3},{2,3},{1,1},{1,1},{1,1},{1,2}}, 4)) // 6
    // Example 2:
    // Input: books = [[1,3],[2,4],[3,2]], shelfWidth = 6
    // Output: 4
    fmt.Println(minHeightShelves([][]int{{1,3},{2,4},{3,2}}, 6)) // 4
}