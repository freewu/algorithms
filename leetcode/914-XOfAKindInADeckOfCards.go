package main

// 914. X of a Kind in a Deck of Cards
// You are given an integer array deck where deck[i] represents the number written on the ith card.

// Partition the cards into one or more groups such that:
//     Each group has exactly x cards where x > 1, and
//     All the cards in one group have the same integer written on them.

// Return true if such partition is possible, or false otherwise.

// Example 1:
// Input: deck = [1,2,3,4,4,3,2,1]
// Output: true
// Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].

// Example 2:
// Input: deck = [1,1,1,2,2,2,3,3]
// Output: false
// Explanation: No possible partition.

// Constraints:
//     1 <= deck.length <= 10^4
//     0 <= deck[i] < 10^4

import "fmt"

func hasGroupsSizeX(deck []int) bool {
    if len(deck) < 2 {
        return false
    }
    mapAppend := func(data []int, pos, val int) []int {
        if l := len(data);l > pos {
        } else if c := cap(data); c >= pos+1 {
            data = data[:pos+1]
        } else {
            data = append(data, make([]int, pos+1-c)...)
        }
        data[pos] += val
        return data
    }
    gcd := func (x, y int) int {  for y != 0 {  x, y = y, x % y; }; return x; }
    count := []int{}
    for _, v := range deck {
        count = mapAppend(count, v, 1)
    }
    g := 0
    for _, v := range count {
        if v == 0 {
            continue
        } else if g == 0 {
            g = v
        } else if g = gcd(g, v); g == 1 {
            return false
        }
    }
    return true
}

func hasGroupsSizeX1(deck []int) bool {
    g, nums := 0, make([]int, 10004)
    for _, v := range deck {
        nums[v]++
    }
    gcd := func (x, y int) int {  for y != 0 {  x, y = y, x % y; }; return x; }
    for _, v := range nums {
        g = gcd(g, v) // 最大公约数为1，说明不能分组成功
        if g == 1 {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: deck = [1,2,3,4,4,3,2,1]
    // Output: true
    // Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].
    fmt.Println(hasGroupsSizeX([]int{1,2,3,4,4,3,2,1})) // true
    // Example 2:
    // Input: deck = [1,1,1,2,2,2,3,3]
    // Output: false
    // Explanation: No possible partition.
    fmt.Println(hasGroupsSizeX([]int{1,1,1,2,2,2,3,3})) // false

    fmt.Println(hasGroupsSizeX([]int{1})) // false
    fmt.Println(hasGroupsSizeX([]int{1,1,1,1,2,2})) // true

    fmt.Println(hasGroupsSizeX1([]int{1,2,3,4,4,3,2,1})) // true
    fmt.Println(hasGroupsSizeX1([]int{1,1,1,2,2,2,3,3})) // false
    fmt.Println(hasGroupsSizeX1([]int{1})) // false
    fmt.Println(hasGroupsSizeX1([]int{1,1,1,1,2,2})) // true
}