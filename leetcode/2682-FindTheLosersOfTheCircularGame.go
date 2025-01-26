package main

// 2682. Find the Losers of the Circular Game
// There are n friends that are playing a game. 
// The friends are sitting in a circle and are numbered from 1 to n in clockwise order.
// More formally, moving clockwise from the ith friend brings you to the (i+1)th friend for 1 <= i < n, 
// and moving clockwise from the nth friend brings you to the 1st friend.

// The rules of the game are as follows:

// 1st friend receives the ball.

//     1. After that, 1st friend passes it to the friend who is k steps away from them in the clockwise direction.
//     2. After that, the friend who receives the ball should pass it to the friend who is 2 * k steps away from them in the clockwise direction.
//     3. After that, the friend who receives the ball should pass it to the friend who is 3 * k steps away from them in the clockwise direction, and so on and so forth.

// In other words, on the ith turn, the friend holding the ball should pass it to the friend who is i * k steps away from them in the clockwise direction.

// The game is finished when some friend receives the ball for the second time.

// The losers of the game are friends who did not receive the ball in the entire game.

// Given the number of friends, n, and an integer k, return the array answer, which contains the losers of the game in the ascending order.


// Example 1:
// Input: n = 5, k = 2
// Output: [4,5]
// Explanation: The game goes as follows:
// 1) Start at 1st friend and pass the ball to the friend who is 2 steps away from them - 3rd friend.
// 2) 3rd friend passes the ball to the friend who is 4 steps away from them - 2nd friend.
// 3) 2nd friend passes the ball to the friend who is 6 steps away from them  - 3rd friend.
// 4) The game ends as 3rd friend receives the ball for the second time.

// Example 2:
// Input: n = 4, k = 4
// Output: [2,3,4]
// Explanation: The game goes as follows:
// 1) Start at the 1st friend and pass the ball to the friend who is 4 steps away from them - 1st friend.
// 2) The game ends as 1st friend receives the ball for the second time.

// Constraints:
//     1 <= k <= n <= 50

import "fmt"

func circularGameLosers(n int, k int) []int {
    mp := make(map[int]bool)
    v, i := 1, 1
    for !mp[v] {
        mp[v] = true
        v = ((k * i) % n) + v
        if v > n { v -= n }
        i++
    }
    res := []int{}
    for i = 1; i <= n; i++ {
        if !mp[i] { 
            res = append(res, i) 
        }
    }
    return res
}

func circularGameLosers1(n int, k int) []int {
    visit := make(map[int]bool, n)
    j := 0
    for i := k; !visit[j]; i += k {
        visit[j] = true
        j = (j + i) % n
    }
    res := make([]int, 0, n)
    for i := 0; i < n; i++ {
        if !visit[i] {
            res = append(res, i + 1)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5, k = 2
    // Output: [4,5]
    // Explanation: The game goes as follows:
    // 1) Start at 1st friend and pass the ball to the friend who is 2 steps away from them - 3rd friend.
    // 2) 3rd friend passes the ball to the friend who is 4 steps away from them - 2nd friend.
    // 3) 2nd friend passes the ball to the friend who is 6 steps away from them  - 3rd friend.
    // 4) The game ends as 3rd friend receives the ball for the second time.
    fmt.Println(circularGameLosers(5, 2)) // [4,5]
    // Example 2:
    // Input: n = 4, k = 4
    // Output: [2,3,4]
    // Explanation: The game goes as follows:
    // 1) Start at the 1st friend and pass the ball to the friend who is 4 steps away from them - 1st friend.
    // 2) The game ends as 1st friend receives the ball for the second time.
    fmt.Println(circularGameLosers(4, 4)) // [2,3,4]

    fmt.Println(circularGameLosers(1, 1)) // []
    fmt.Println(circularGameLosers(50, 50)) // [2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50]
    fmt.Println(circularGameLosers(50, 1)) // [3 5 8 9 10 12 13 14 15 18 19 20 21 23 24 25 26 27 28 30 31 32 33 34 35 36 38 39 40 41 42 43 44 45 47 48 49 50]

    fmt.Println(circularGameLosers1(5, 2)) // [4,5]
    fmt.Println(circularGameLosers1(4, 4)) // [2,3,4]
    fmt.Println(circularGameLosers1(1, 1)) // []
    fmt.Println(circularGameLosers1(50, 50)) // [2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50]
    fmt.Println(circularGameLosers1(50, 1)) // [3 5 8 9 10 12 13 14 15 18 19 20 21 23 24 25 26 27 28 30 31 32 33 34 35 36 38 39 40 41 42 43 44 45 47 48 49 50]

}