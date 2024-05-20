package main

// 319. Bulb Switcher
// There are n bulbs that are initially off. 
// You first turn on all the bulbs, then you turn off every second bulb.

// On the third round, you toggle every third bulb (turning on if it's off or turning off if it's on). 
// For the ith round, you toggle every i bulb. For the nth round, you only toggle the last bulb.

// Return the number of bulbs that are on after n rounds.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/bulb.jpg" />
// Input: n = 3
// Output: 1
// Explanation: At first, the three bulbs are [off, off, off].
// After the first round, the three bulbs are [on, on, on].
// After the second round, the three bulbs are [on, off, on].
// After the third round, the three bulbs are [on, off, off]. 
// So you should return 1 because there is only one bulb is on.

// Example 2:
// Input: n = 0
// Output: 0

// Example 3:
// Input: n = 1
// Output: 1
 
// Constraints:
//     0 <= n <= 10^9

import "fmt"
import "math"

// 1 到 n 中的某个数 x 有奇数个约数,也即 x 是完全平方数
// 计算 1 到 n 中完全平方数的个数 sqrt(n)
func bulbSwitch(n int) int {
    return int(math.Sqrt(float64(n)))
}

// 能够被拆解为平方根的数字说明都会被开关奇数次，
// 那么就会是被点亮的；所以只要找到小于n的可以平方根就可以了
func bulbSwitch1(n int) int {
    res := 0
    for i := 1; i*i <= n; i++ {
        res++
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/11/05/bulb.jpg" />
    // Input: n = 3
    // Output: 1
    // Explanation: At first, the three bulbs are [off, off, off].
    // After the first round, the three bulbs are [on, on, on].
    // After the second round, the three bulbs are [on, off, on].
    // After the third round, the three bulbs are [on, off, off]. 
    // So you should return 1 because there is only one bulb is on.
    fmt.Println(bulbSwitch(3)) // 1
    // Example 2:
    // Input: n = 0
    // Output: 0
    fmt.Println(bulbSwitch(0)) // 0
    // Example 3:
    // Input: n = 1
    // Output: 1
    fmt.Println(bulbSwitch(1)) // 1

    fmt.Println(bulbSwitch(4)) // 2
    fmt.Println(bulbSwitch(9)) // 3
    fmt.Println(bulbSwitch(16)) // 4

    fmt.Println(bulbSwitch1(3)) // 1
    fmt.Println(bulbSwitch1(0)) // 0
    fmt.Println(bulbSwitch1(1)) // 1
    fmt.Println(bulbSwitch1(4)) // 2
    fmt.Println(bulbSwitch1(9)) // 3
    fmt.Println(bulbSwitch1(16)) // 4
}