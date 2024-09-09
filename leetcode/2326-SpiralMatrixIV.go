package main

// 2326. Spiral Matrix IV
// You are given two integers m and n, which represent the dimensions of a matrix.

// You are also given the head of a linked list of integers.

// Generate an m x n matrix that contains the integers in the linked list presented in spiral order (clockwise), starting from the top-left of the matrix. 
// If there are remaining empty spaces, fill them with -1.

// Return the generated matrix.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/05/09/ex1new.jpg" />
// Input: m = 3, n = 5, head = [3,0,2,6,8,1,7,9,4,2,5,5,0]
// Output: [[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
// Explanation: The diagram above shows how the values are printed in the matrix.
// Note that the remaining spaces in the matrix are filled with -1.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/05/11/ex2.jpg" />
// Input: m = 1, n = 4, head = [0,1,2]
// Output: [[0,1,2,-1]]
// Explanation: The diagram above shows how the values are printed from left to right in the matrix.
// The last space in the matrix is set to -1.

// Constraints:
//     1 <= m, n <= 10^5
//     1 <= m * n <= 10^5
//     The number of nodes in the list is in the range [1, m * n].
//     0 <= Node.val <= 1000

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

// 打印链表
func printListNode(l *ListNode) {
    if nil == l {
        return
    }
    for {
        if nil == l.Next {
            fmt.Print(l.Val)
            break
        } else {
            fmt.Print(l.Val, " -> ")
        }
        l = l.Next
    }
    fmt.Println()
}

// 数组创建链表
func makeListNode(arr []int) *ListNode {
    if (len(arr) == 0) {
        return nil
    }
    l := len(arr) - 1
    head := &ListNode{arr[l], nil}
    for i := l - 1; i >= 0; i-- {
        n := &ListNode{arr[i], head}
        head = n
    }
    return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    res := make([][]int, m)
    for i := 0; i < m; i++ { res[i] = make([]int, n) }
    for i := 0; i < m; i++ { // 生成一个全是 -1的  m * n 的矩阵 
        for j := 0; j < n; j++ { 
            res[i][j] = -1 
        }
    }
    y, x := 0, 0
    dirs := [][]int{ {0,0}, {0,1}, {1,0}, {0,-1}, {-1,0}, }
    d, shrink, lim := 1, 0, 0
    for head != nil { // 遍历链表
        res[y][x] = head.Val
        head = head.Next
        i, j := y + dirs[d][0], x + dirs[d][1]
        if i >= shrink && i < m - lim && j >= lim && j < n - lim {
            y, x = i, j
        } else {
            d++
            if d == 4 { 
                shrink++
            } else if d > 4 { 
                d = 1
                lim++ 
            }
            y, x = y + dirs[d][0], x + dirs[d][1]
        }
    }
    return res
}

func spiralMatrix1(m int, n int, head *ListNode) [][]int {
    res := make([][]int, m)
    for i := range res {
        res[i] = make([]int, n)
        for j := range res[i] {
            res[i][j] = -1
        }
    }
    up, down, left, right := 0, m - 1, 0, n - 1
    cur := head
    for cur != nil {
        for i := left; i <= right; i++ {
            if cur == nil {
                return res
            }
            res[up][i] = cur.Val
            cur = cur.Next
        }
        up++
        for i := up; i <= down; i++ {
            if cur == nil {
                return res
            }
            res[i][right] = cur.Val
            cur = cur.Next
        }
        right--
        for i := right; i >= left; i-- {
            if cur == nil {
                return res
            }
            res[down][i] = cur.Val
            cur = cur.Next
        }
        down--
        for i := down; i >= up; i-- {
            if cur == nil {
                return res
            }
            res[i][left] = cur.Val
            cur = cur.Next
        }
        left++
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/05/09/ex1new.jpg" />
    // Input: m = 3, n = 5, head = [3,0,2,6,8,1,7,9,4,2,5,5,0]
    // Output: [[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
    // Explanation: The diagram above shows how the values are printed in the matrix.
    // Note that the remaining spaces in the matrix are filled with -1.
    list1 := makeListNode([]int{3,0,2,6,8,1,7,9,4,2,5,5,0})
    printListNode(list1) // 3 -> 0 -> 2 -> 6 -> 8 -> 1 -> 7 -> 9 -> 4 -> 2 -> 5 -> 5 -> 0
    fmt.Println(spiralMatrix(3, 5, list1)) // [[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/05/11/ex2.jpg" />
    // Input: m = 1, n = 4, head = [0,1,2]
    // Output: [[0,1,2,-1]]
    // Explanation: The diagram above shows how the values are printed from left to right in the matrix.
    // The last space in the matrix is set to -1.
    list2 := makeListNode([]int{0,1,2}) // 0 -> 1 -> 2
    printListNode(list2)
    fmt.Println(spiralMatrix(1, 4, list2)) // [[0,1,2,-1]]

    list11 := makeListNode([]int{3,0,2,6,8,1,7,9,4,2,5,5,0})
    printListNode(list11) // 3 -> 0 -> 2 -> 6 -> 8 -> 1 -> 7 -> 9 -> 4 -> 2 -> 5 -> 5 -> 0
    fmt.Println(spiralMatrix1(3, 5, list11)) // [[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
    list12 := makeListNode([]int{0,1,2}) // 0 -> 1 -> 2
    printListNode(list12)
    fmt.Println(spiralMatrix1(1, 4, list12)) // [[0,1,2,-1]]
}