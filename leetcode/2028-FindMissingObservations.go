package main

// 2028. Find Missing Observations
// You have observations of n + m 6-sided dice rolls with each face numbered from 1 to 6. 
// n of the observations went missing, and you only have the observations of m rolls. 
// Fortunately, you have also calculated the average value of the n + m rolls.

// You are given an integer array rolls of length m where rolls[i] is the value of the ith observation. 
// You are also given the two integers mean and n.

// Return an array of length n containing the missing observations such that the average value of the n + m rolls is exactly mean. 
// If there are multiple valid answers, return any of them. If no such array exists, return an empty array.

// The average value of a set of k numbers is the sum of the numbers divided by k.
// Note that mean is an integer, so the sum of the n + m rolls should be divisible by n + m.

// Example 1:
// Input: rolls = [3,2,4,3], mean = 4, n = 2
// Output: [6,6]
// Explanation: The mean of all n + m rolls is (3 + 2 + 4 + 3 + 6 + 6) / 6 = 4.

// Example 2:
// Input: rolls = [1,5,6], mean = 3, n = 4
// Output: [2,3,2,2]
// Explanation: The mean of all n + m rolls is (1 + 5 + 6 + 2 + 3 + 2 + 2) / 7 = 3.

// Example 3:
// Input: rolls = [1,2,3,4], mean = 6, n = 4
// Output: []
// Explanation: It is impossible for the mean to be 6 no matter what the 4 missing rolls are.

// Constraints:
//     m == rolls.length
//     1 <= n, m <= 10^5
//     1 <= rolls[i], mean <= 6

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
    m, length := 0, len(rolls) + n // 数据长度要加上缺失的
    total := mean * length // total is what the sum of n + m should be
    for i := 0; i < len(rolls); i++ { // get the sum of rolls aka m
        m += rolls[i]
    }
    target := total - m // target is our target sum of n
    part, sum, res :=  target / n, 0, make([]int, n)
    if part < 1 || part > 6 { // its impossible  要凑的值小于 1 或 大于 6 （骰子只有6个面 1 - 6）
        return []int{}
    }
    for i := 0; i < len(res); i++ {
        sum += part
        res[i] = part
    }
    i := 0
    for sum < target { // increment dice by 1 until we reach our target n
        for i < len(res) && res[i] == 6 { i++; }
        if i == len(res) { return []int{}; }
        res[i]++
        sum++
    }
    return res
}

func missingRolls1(rolls []int, mean int, n int) []int {
    sum, res, cnt := 0, make([]int, n), make([]int, 7)
    for _, v := range rolls { // 累加出可知到的总值
        sum += v
    }
    left := mean * (n + len(rolls)) - sum // 得到缺失的总数
    if left < n || left > 6 * n {
        return []int{}
    }
    for i := 1; i <= 6; i++ {
        if left > n {
            cnt[i-1] -= n
            cnt[i] += n
            left -= n
        } else {
            cnt[i-1] -= left
            cnt[i] += left
            break
        }
    }
    for i, index := 1, 0; i <= 6; i++ {
        for ; cnt[i] > 0; cnt[i]-- {
            res[index] = i
            index++
        }
    }
    return res
}

func missingRolls2(rolls []int, mean int, n int) []int {
    mTotal := 0
    for _, v := range rolls { // 统计出未缺失的观测数据总和
        mTotal += v
    }
    nTotal := ((mean) * (n + len(rolls))) - mTotal //  计算出缺失的观测数据总和
    if nTotal > n * 6 || nTotal < n {
        return []int{}
    }
    res, rem := make([]int, n), nTotal
    for i := range res {
        res[i] = 1
        rem -= 1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := range res {
        currAdd := min(5, rem)
        res[i] += currAdd
        rem -= currAdd
        if rem == 0 {
            break
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: rolls = [3,2,4,3], mean = 4, n = 2
    // Output: [6,6]
    // Explanation: The mean of all n + m rolls is (3 + 2 + 4 + 3 + 6 + 6) / 6 = 4.
    fmt.Println(missingRolls([]int{3,2,4,3}, 4, 2)) // [6,6]
    // Example 2:
    // Input: rolls = [1,5,6], mean = 3, n = 4
    // Output: [2,3,2,2]
    // Explanation: The mean of all n + m rolls is (1 + 5 + 6 + 2 + 3 + 2 + 2) / 7 = 3.
    fmt.Println(missingRolls([]int{1,5,6}, 3, 4)) // [2,3,2,2]
    // Example 3:
    // Input: rolls = [1,2,3,4], mean = 6, n = 4
    // Output: []
    // Explanation: It is impossible for the mean to be 6 no matter what the 4 missing rolls are.
    fmt.Println(missingRolls([]int{1,2,3,4}, 6, 4)) // []

    fmt.Println(missingRolls1([]int{3,2,4,3}, 4, 2)) // [6,6]
    fmt.Println(missingRolls1([]int{1,5,6}, 3, 4)) // [2,3,2,2]
    fmt.Println(missingRolls1([]int{1,2,3,4}, 6, 4)) // []

    fmt.Println(missingRolls2([]int{3,2,4,3}, 4, 2)) // [6,6]
    fmt.Println(missingRolls2([]int{1,5,6}, 3, 4)) // [2,3,2,2]
    fmt.Println(missingRolls2([]int{1,2,3,4}, 6, 4)) // []
}