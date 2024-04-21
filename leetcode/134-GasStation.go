package main

// 134. Gas Station
// There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].
// You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. 
// You begin the journey with an empty tank at one of the gas stations.

// Given two integer arrays gas and cost,
// return the starting gas station's index if you can travel around the circuit once in the clockwise direction, 
// otherwise return -1. 
// If there exists a solution, it is guaranteed to be unique

// Example 1:
// Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// Output: 3
// Explanation:
// Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
// Travel to station 4. Your tank = 4 - 1 + 5 = 8
// Travel to station 0. Your tank = 8 - 2 + 1 = 7
// Travel to station 1. Your tank = 7 - 3 + 2 = 6
// Travel to station 2. Your tank = 6 - 4 + 3 = 5
// Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
// Therefore, return 3 as the starting index.

// Example 2:
// Input: gas = [2,3,4], cost = [3,4,3]
// Output: -1
// Explanation:
// You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
// Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
// Travel to station 0. Your tank = 4 - 3 + 2 = 3
// Travel to station 1. Your tank = 3 - 3 + 3 = 3
// You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
// Therefore, you can't travel around the circuit once no matter where you start.
 
// Constraints:
//     n == gas.length == cost.length
//     1 <= n <= 10^5
//     0 <= gas[i], cost[i] <= 10^4

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
    tank, total, start := 0, 0, 0
    for i := 0; i < len(gas); i++ {
        tank += gas[i] - cost[i]
        total += gas[i] - cost[i] // 计算总体的油耗
        if tank < 0 {
            start = i + 1
            tank = 0
        }        
    }
    if total >= 0 { // 如果总油耗够说明可以从 start 为起点开始一圈
        return start
    }      
    return -1   
}

func canCompleteCircuit1(gas []int, cost []int) int {
    res, gasSum, costSum, suffixSum := -1, 0, 0, 0 
    for i := range gas {
        gasSum += gas[i]
        costSum += cost[i]
        suffixSum += gas[i] - cost[i]
        if suffixSum >= 0 {
            if res == -1 {
                res = i
            }
        } else {
            suffixSum = 0
            res = -1
        }
    }
    if gasSum >= costSum {
        return res
    }
    return -1
}

func canCompleteCircuit2(gas []int, cost []int) int {
    sliceSum := func(s []int) (res int) { for _, v := range s { res += v; }; return res; }
    if sliceSum(gas) < sliceSum(cost) {
        return -1
    }
    start, sum := 0, 0
    for i := 0; i < len(gas); i++ {
        sum += gas[i] - cost[i]
        if sum < 0 {
            start = i + 1
            sum = 0
        }
    }
    return start
}


func main() {
    // Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
    // Travel to station 4. Your tank = 4 - 1 + 5 = 8
    // Travel to station 0. Your tank = 8 - 2 + 1 = 7
    // Travel to station 1. Your tank = 7 - 3 + 2 = 6
    // Travel to station 2. Your tank = 6 - 4 + 3 = 5
    // Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
    // Therefore, return 3 as the starting index.
    fmt.Println(canCompleteCircuit([]int{1,2,3,4,5},[]int{3,4,5,1,2})) // 3
    // Example 2:
    // Input: gas = [2,3,4], cost = [3,4,3]
    // Output: -1
    // Explanation:
    // You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
    // Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
    // Travel to station 0. Your tank = 4 - 3 + 2 = 3
    // Travel to station 1. Your tank = 3 - 3 + 3 = 3
    // You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
    // Therefore, you can't travel around the circuit once no matter where you start.
    fmt.Println(canCompleteCircuit([]int{2,3,4},[]int{3,4,3})) // -1

    fmt.Println(canCompleteCircuit1([]int{1,2,3,4,5},[]int{3,4,5,1,2})) // 3
    fmt.Println(canCompleteCircuit1([]int{2,3,4},[]int{3,4,3})) // -1

    fmt.Println(canCompleteCircuit2([]int{1,2,3,4,5},[]int{3,4,5,1,2})) // 3
    fmt.Println(canCompleteCircuit2([]int{2,3,4},[]int{3,4,3})) // -1
}
