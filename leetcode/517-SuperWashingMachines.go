package main

// 517. Super Washing Machines
// You have n super washing machines on a line. Initially, each washing machine has some dresses or is empty.

// For each move, you could choose any m (1 <= m <= n) washing machines, 
// and pass one dress of each washing machine to one of its adjacent washing machines at the same time.

// Given an integer array machines representing the number of dresses in each washing machine from left to right on the line, 
// return the minimum number of moves to make all the washing machines have the same number of dresses. 
// If it is not possible to do it, return -1.

// Example 1:
// Input: machines = [1,0,5]
// Output: 3
// Explanation:
// 1st move:    1     0 <-- 5    =>    1     1     4
// 2nd move:    1 <-- 1 <-- 4    =>    2     1     3
// 3rd move:    2     1 <-- 3    =>    2     2     2

// Example 2:
// Input: machines = [0,3,0]
// Output: 2
// Explanation:
// 1st move:    0 <-- 3     0    =>    1     2     0
// 2nd move:    1     2 --> 0    =>    1     1     1

// Example 3:
// Input: machines = [0,2,0]
// Output: -1
// Explanation:
// It's impossible to make all three washing machines have the same number of dresses.

// Constraints:
//     n == machines.length
//     1 <= n <= 10^4
//     0 <= machines[i] <= 10^5

import "fmt"

func findMinMoves(machines []int) int {
    //since each machine can only transfer one dress at each step
    //the strategy is to find for each machine, the sum of dresses that passes left/right through this machine
    //among all the machine, the highest passes through number is also the steps to take
    sum, n := 0, len(machines)
    for _, v := range machines { // 汇总
        sum +=v
    }
    if sum % n != 0 { // 不能被平分
        return -1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, avg, leftSum := 0, sum / n, 0
    leftTarget, rightTarget := 0, sum + avg
    for _, v := range machines {
        leftSum += v  // prefix sum, i included
        rightSum := sum - leftSum + v // suffix sum, i included
        leftTarget += avg // prefix target, i included
        rightTarget -= avg  // suffix target, i included
        toRight, toLeft := 0, 0
        if leftSum > leftTarget {
            toRight = leftSum - leftTarget
        }
        if rightSum > rightTarget {
            toLeft = rightSum - rightTarget
        }
        res = max(res, toLeft + toRight)
    }
    return res
}

func findMinMoves1(machines []int) int {
    sum, n := 0, len(machines)
    for _, v := range machines { // 汇总
        sum += v
    }
    if sum % n != 0{
        return -1
    }
    res, avg := 0, sum / n
    for i := 0; i < n; i++ {
        machines[i] = machines[i] - avg
        if machines[i] > 0 && res < machines[i] {
            res = machines[i]
        }
    }
    for i := 1; i < n; i++ {
        machines[i] += machines[i-1]
        if machines[i] < 0 && res < -machines[i] {
            res = -machines[i]
        }
        if machines[i] > 0 && res < machines[i] {
            res = machines[i]
        }
    } 
    return res
}

func main() {
    // Example 1:
    // Input: machines = [1,0,5]
    // Output: 3
    // Explanation:
    // 1st move:    1     0 <-- 5    =>    1     1     4
    // 2nd move:    1 <-- 1 <-- 4    =>    2     1     3
    // 3rd move:    2     1 <-- 3    =>    2     2     2
    fmt.Println(findMinMoves([]int{1,0,5})) // 3
    // Example 2:
    // Input: machines = [0,3,0]
    // Output: 2
    // Explanation:
    // 1st move:    0 <-- 3     0    =>    1     2     0
    // 2nd move:    1     2 --> 0    =>    1     1     1
    fmt.Println(findMinMoves([]int{0,3,0})) // 2
    // Example 3:
    // Input: machines = [0,2,0]
    // Output: -1
    // Explanation:
    // It's impossible to make all three washing machines have the same number of dresses.
    fmt.Println(findMinMoves([]int{0,2,0})) // -1

    fmt.Println(findMinMoves1([]int{1,0,5})) // 3
    fmt.Println(findMinMoves1([]int{0,3,0})) // 2
    fmt.Println(findMinMoves1([]int{0,2,0})) // -1
}