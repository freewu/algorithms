package main

// 2591. Distribute Money to Maximum Children
// You are given an integer money denoting the amount of money (in dollars) 
// that you have and another integer children denoting the number of children that you must distribute the money to.

// You have to distribute the money according to the following rules:
//     1. All money must be distributed.
//     2. Everyone must receive at least 1 dollar.
//     3. Nobody receives 4 dollars.

// Return the maximum number of children who may receive exactly 8 dollars if you distribute the money according to the aforementioned rules. 
// If there is no way to distribute the money, return -1.

// Example 1:
// Input: money = 20, children = 3
// Output: 1
// Explanation: 
// The maximum number of children with 8 dollars will be 1. One of the ways to distribute the money is:
// - 8 dollars to the first child.
// - 9 dollars to the second child. 
// - 3 dollars to the third child.
// It can be proven that no distribution exists such that number of children getting 8 dollars is greater than 1.

// Example 2:
// Input: money = 16, children = 2
// Output: 2
// Explanation: Each child can be given 8 dollars.

// Constraints:
//     1 <= money <= 200
//     2 <= children <= 30

import "fmt"

func distMoney(money int, children int) int {
    switch {
        case children == 0 || money < children:
            return -1
        case children == 1 && money == 4:
            return -1
        case children == 1 && money == 8:
            return 1
        case children == 1:
            return 0
    }
    dpPrevious := distMoney(money - 8, children - 1)
    switch {
        case dpPrevious == -1 && children >= 2:
            return 0
        case dpPrevious == -1:
            return -1
        default:
            return 1 + dpPrevious
    }
}

func distMoney1(money int, children int) int {
    if money < children { return -1 }
    tmp := money - children
    if tmp < 0 { return -1 }
    tmp1 := tmp % 7
    if tmp / 7 > children || ( tmp / 7 == children && tmp1 != 0) { return children - 1 }
    if tmp1 == 3 && tmp / 7 == children - 1 { return children - 2 }
    return tmp / 7
}

func distMoney2(money int, children int) int {
    money -= children
    if money < 0 { return -1 }
    count, rem := money / 7, money % 7
    if count == children && rem == 0 { return children }
    if count == children - 1 && rem == 3 { return children - 2 }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return min(children - 1, count)
}

func main() {
    // Example 1:
    // Input: money = 20, children = 3
    // Output: 1
    // Explanation: 
    // The maximum number of children with 8 dollars will be 1. One of the ways to distribute the money is:
    // - 8 dollars to the first child.
    // - 9 dollars to the second child. 
    // - 3 dollars to the third child.
    // It can be proven that no distribution exists such that number of children getting 8 dollars is greater than 1.
    fmt.Println(distMoney(20, 3)) // 1
    // Example 2:
    // Input: money = 16, children = 2
    // Output: 2
    // Explanation: Each child can be given 8 dollars.
    fmt.Println(distMoney(16, 2)) // 2

    fmt.Println(distMoney(1, 2)) // -1
    fmt.Println(distMoney(200, 30)) // 24
    fmt.Println(distMoney(1, 30)) // -1
    fmt.Println(distMoney(200, 2)) // 1

    fmt.Println(distMoney1(20, 3)) // 1
    fmt.Println(distMoney1(16, 2)) // 2
    fmt.Println(distMoney1(1, 2)) // -1
    fmt.Println(distMoney1(200, 30)) // 24
    fmt.Println(distMoney1(1, 30)) // -1
    fmt.Println(distMoney1(200, 2)) // 1

    fmt.Println(distMoney2(20, 3)) // 1
    fmt.Println(distMoney2(16, 2)) // 2
    fmt.Println(distMoney2(1, 2)) // -1
    fmt.Println(distMoney2(200, 30)) // 24
    fmt.Println(distMoney2(1, 30)) // -1
    fmt.Println(distMoney2(200, 2)) // 1
}