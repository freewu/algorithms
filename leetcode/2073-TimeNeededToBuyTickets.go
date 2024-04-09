package main

// 2073. Time Needed to Buy Tickets
// There are n people in a line queuing to buy tickets, where the 0th person is at the front of the line and the (n - 1)th person is at the back of the line.
// You are given a 0-indexed integer array tickets of length n where the number of tickets that the ith person would like to buy is tickets[i].

// Each person takes exactly 1 second to buy a ticket. 
// A person can only buy 1 ticket at a time and has to go back to the end of the line (which happens instantaneously) in order to buy more tickets. If a person does not have any tickets left to buy, the person will leave the line.

// Return the time taken for the person at position k (0-indexed) to finish buying tickets.

// Example 1:
// Input: tickets = [2,3,2], k = 2
// Output: 6
// Explanation: 
// - In the first pass, everyone in the line buys a ticket and the line becomes [1, 2, 1].
// - In the second pass, everyone in the line buys a ticket and the line becomes [0, 1, 0].
// The person at position 2 has successfully bought 2 tickets and it took 3 + 3 = 6 seconds.

// Example 2:
// Input: tickets = [5,1,1,1], k = 0
// Output: 8
// Explanation:
// - In the first pass, everyone in the line buys a ticket and the line becomes [4, 0, 0, 0].
// - In the next 4 passes, only the person in position 0 is buying tickets.
// The person at position 0 has successfully bought 5 tickets and it took 4 + 1 + 1 + 1 + 1 = 8 seconds.

// Constraints:
//     n == tickets.length
//     1 <= n <= 100
//     1 <= tickets[i] <= 100
//     0 <= k < n

import "fmt"

// 暴力解法
func timeRequiredToBuy(tickets []int, k int) int {
    res := 0
    for {
        // 循环 tickets[k] 轮
        for i := 0; i < len(tickets); i++ {
            if tickets[k] == 0 { // 如果买完了直接返回
                return res
            }
            if tickets[i] == 0 { // 如果为 0 了,就片算作离队不需要继续下面的处理
                continue
            }
            tickets[i]--
            res++
        }
    }
    return res
}

// O(n)
func timeRequiredToBuy1(tickets []int, k int) int {
    res := 0
    c := tickets[k]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, n := range tickets {
        if i == k {
            res += c
        } else if i < k { // k 之前的
            res += min(c, n)
        } else { // k 之后的
            res += min(c - 1, n)
        }
    }
    return res
}

func main() {
    // - In the first pass, everyone in the line buys a ticket and the line becomes [1, 2, 1].
    // - In the second pass, everyone in the line buys a ticket and the line becomes [0, 1, 0].
    // The person at position 2 has successfully bought 2 tickets and it took 3 + 3 = 6 seconds.
    fmt.Println(timeRequiredToBuy([]int{2,3,2},2)) // 6

    // - In the first pass, everyone in the line buys a ticket and the line becomes [4, 0, 0, 0].
    // - In the next 4 passes, only the person in position 0 is buying tickets.
    // The person at position 0 has successfully bought 5 tickets and it took 4 + 1 + 1 + 1 + 1 = 8 seconds.
    fmt.Println(timeRequiredToBuy([]int{5,1,1,1},0)) // 8

    fmt.Println(timeRequiredToBuy1([]int{2,3,2},2)) // 6
    fmt.Println(timeRequiredToBuy1([]int{5,1,1,1},0)) // 8
}