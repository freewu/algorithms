package main

// 1739. Building Boxes
// You have a cubic storeroom where the width, length, and height of the room are all equal to n units. 
// You are asked to place n boxes in this room where each box is a cube of unit side length. 
// There are however some rules to placing the boxes:
//     1. You can place the boxes anywhere on the floor.
//     2. If box x is placed on top of the box y, 
//        then each side of the four vertical sides of the box y must either be adjacent to another box or to a wall.

// Given an integer n, return the minimum possible number of boxes touching the floor.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/04/3-boxes.png" />
// Input: n = 3
// Output: 3
// Explanation: The figure above is for the placement of the three boxes.
// These boxes are placed in the corner of the room, where the corner is on the left side.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/04/4-boxes.png" />
// Input: n = 4
// Output: 3
// Explanation: The figure above is for the placement of the four boxes.
// These boxes are placed in the corner of the room, where the corner is on the left side.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/01/04/10-boxes.png" />
// Input: n = 10
// Output: 6
// Explanation: The figure above is for the placement of the ten boxes.
// These boxes are placed in the corner of the room, where the corner is on the back side.
 
// Constraints:
//     1 <= n <= 10^9

import "fmt"
import "sort"

func minimumBoxes(n int) int {
    c := func(n int) int {
        if n == 1 { return 1 }
        return (n + 1) * n / 2
    }
    b := func(n int) int {
        t, s := 1, 0
        for n > 0 {
            if n >= t {
                n -= t
                s += c(t)
                t++
            } else {
                s += c(n)
                break
            }
        }
        return s
    }
    return sort.Search(n, func(i int) bool { 
        return n <= b(i) 
    })
}

func minimumBoxes1(n int) int {
    totalBoxes, baseSize, nextPyramidHeight := 0, 0, 1
    for totalBoxes + baseSize + nextPyramidHeight < n {
        baseSize += nextPyramidHeight
        nextPyramidHeight += 1
        totalBoxes += baseSize
    }
    n -= totalBoxes
    numTriangleBoxes := func(height int) int { return height * (height + 1) / 2 }
    low, high := 0, n
    for low < high {
        mid := int((low + high) / 2)
        if numTriangleBoxes(mid) < n {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return baseSize + low
}

func minimumBoxes2 (n int) int {
    cur, i, j := 1, 1, 1
    for n > cur {
        n -= cur
        i++
        cur += i
    }
    cur = 1
    for n > cur {
        n -= cur
        j++
        cur++
    }
    return (i-1) * i / 2 + j
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/01/04/3-boxes.png" />
    // Input: n = 3
    // Output: 3
    // Explanation: The figure above is for the placement of the three boxes.
    // These boxes are placed in the corner of the room, where the corner is on the left side.
    fmt.Println(minimumBoxes(3)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/01/04/4-boxes.png" />
    // Input: n = 4
    // Output: 3
    // Explanation: The figure above is for the placement of the four boxes.
    // These boxes are placed in the corner of the room, where the corner is on the left side.
    fmt.Println(minimumBoxes(4)) // 3
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/01/04/10-boxes.png" />
    // Input: n = 10
    // Output: 6
    // Explanation: The figure above is for the placement of the ten boxes.
    // These boxes are placed in the corner of the room, where the corner is on the back side.
    fmt.Println(minimumBoxes(10)) // 6

    fmt.Println(minimumBoxes(1)) // 3
    fmt.Println(minimumBoxes(2)) // 3
    fmt.Println(minimumBoxes(1024)) // 3
    fmt.Println(minimumBoxes(9999999)) // 3

    fmt.Println(minimumBoxes1(3)) // 3
    fmt.Println(minimumBoxes1(4)) // 3
    fmt.Println(minimumBoxes1(10)) // 6
    fmt.Println(minimumBoxes1(1)) // 3
    fmt.Println(minimumBoxes1(2)) // 3
    fmt.Println(minimumBoxes1(1024)) // 3
    fmt.Println(minimumBoxes1(9999999)) // 3

    fmt.Println(minimumBoxes2(3)) // 3
    fmt.Println(minimumBoxes2(4)) // 3
    fmt.Println(minimumBoxes2(10)) // 6
    fmt.Println(minimumBoxes2(1)) // 3
    fmt.Println(minimumBoxes2(2)) // 3
    fmt.Println(minimumBoxes2(1024)) // 3
    fmt.Println(minimumBoxes2(9999999)) // 3
}