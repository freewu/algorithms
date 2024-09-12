package main

// 2398. Maximum Number of Robots Within Budget
// You have n robots. 
// You are given two 0-indexed integer arrays, chargeTimes and runningCosts, both of length n. 
// The ith robot costs chargeTimes[i] units to charge and costs runningCosts[i] units to run. 
// You are also given an integer budget.

// The total cost of running k chosen robots is equal to max(chargeTimes) + k * sum(runningCosts), 
// where max(chargeTimes) is the largest charge cost among the k robots and sum(runningCosts) is the sum of running costs among the k robots.

// Return the maximum number of consecutive robots you can run such that the total cost does not exceed budget.

// Example 1:
// Input: chargeTimes = [3,6,1,3,4], runningCosts = [2,1,3,4,5], budget = 25
// Output: 3
// Explanation: 
// It is possible to run all individual and consecutive pairs of robots within budget.
// To obtain answer 3, consider the first 3 robots. The total cost will be max(3,6,1) + 3 * sum(2,1,3) = 6 + 3 * 6 = 24 which is less than 25.
// It can be shown that it is not possible to run more than 3 consecutive robots within budget, so we return 3.

// Example 2:
// Input: chargeTimes = [11,12,19], runningCosts = [10,8,7], budget = 19
// Output: 0
// Explanation: No robot can be run that does not exceed the budget, so we return 0.

// Constraints:
//     chargeTimes.length == runningCosts.length == n
//     1 <= n <= 5 * 10^4
//     1 <= chargeTimes[i], runningCosts[i] <= 10^5
//     1 <= budget <= 10^15

import "fmt"

// type Deque struct{ front, back []int }

// func (q *Deque) empty() bool    { return len(q.front) == 0 && len(q.back) == 0 }
// func (q *Deque) pushback(x int) { q.back = append(q.back, x) }
// func (q *Deque) popback() int {
//     x, i := 0, len(q.back) - 1
//     if i >= 0 {
//         x, q.back = q.back[i], q.back[:i]
//     } else {
//         x, q.front = q.front[0], q.front[1:]
//     }
//     return x
// }
// func (q *Deque) last() int {
//     x, i := 0, len(q.back) - 1
//     if i >= 0 {
//         x = q.back[i]
//     } else {
//         x = q.front[0]
//     }
//     return x
// }
// func (q *Deque) popfront() int {
//     x, i := 0, len(q.front) - 1
//     if i >= 0 {
//         x, q.front = q.front[i], q.front[:i]
//     } else {
//         x, q.back = q.back[0], q.back[1:]
//     }
//     return x
// }

// func (q *Deque) first() int {
//     x, i := 0, len(q.front) - 1
//     if i >= 0 {
//         x = q.front[i]
//     } else {
//         x = q.back[0]
//     }
//     return x
// }

// // 解答错误 80 / 82
// func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
//     i, j := 0, 0
//     queue := &Deque{}  
//     for total := int64(0); j < len(chargeTimes); j++ {
//         for !queue.empty() && queue.last() <= chargeTimes[j] {
//             queue.popback()
//         }
//         queue.pushback(chargeTimes[j])
//         total += int64(runningCosts[j])
//         if int64(queue.first()) + total * int64(j-i+1) > budget {
//             if queue.first() == chargeTimes[i] {
//                 queue.popfront()
//             }
//             total -= int64(runningCosts[i])
//             i++
//         }
//     }
//     return j - i
// }

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
    res, n, total, queue := 0, len(chargeTimes), int64(0), []int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, j := 0, 0; i < n; i++ {
        total += int64(runningCosts[i])
        for len(queue) > 0 && chargeTimes[queue[len(queue) - 1]] <= chargeTimes[i] {
            queue = queue[:len(queue) - 1] // pop
        }
        queue = append(queue, i)
        for j <= i && int64(i - j + 1) * total + int64(chargeTimes[queue[0]]) > budget {
            if len(queue) > 0 && queue[0] == j {
                queue = queue[1:] // front pop
            }
            total -= int64(runningCosts[j])
            j++
        }
        res = max(res, i - j + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: chargeTimes = [3,6,1,3,4], runningCosts = [2,1,3,4,5], budget = 25
    // Output: 3
    // Explanation: 
    // It is possible to run all individual and consecutive pairs of robots within budget.
    // To obtain answer 3, consider the first 3 robots. The total cost will be max(3,6,1) + 3 * sum(2,1,3) = 6 + 3 * 6 = 24 which is less than 25.
    // It can be shown that it is not possible to run more than 3 consecutive robots within budget, so we return 3.
    fmt.Println(maximumRobots([]int{3,6,1,3,4}, []int{2,1,3,4,5}, 25)) // 3
    // Example 2:
    // Input: chargeTimes = [11,12,19], runningCosts = [10,8,7], budget = 19
    // Output: 0
    // Explanation: No robot can be run that does not exceed the budget, so we return 0.
    fmt.Println(maximumRobots([]int{11,12,19}, []int{10,8,7}, 19)) // 0

    fmt.Println(maximumRobots([]int{4,4,1}, []int{3,1,2}, 7)) // 1
}