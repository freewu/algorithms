package main

import "fmt"

// 365. Water and Jug Problem
// You are given two jugs with capacities jug1Capacity and jug2Capacity liters. 
// There is an infinite amount of water supply available. 
// Determine whether it is possible to measure exactly targetCapacity liters using these two jugs.

// If targetCapacity liters of water are measurable, 
// you must have targetCapacity liters of water contained within one or both buckets by the end.

// Operations allowed:

// 		Fill any of the jugs with water.
// 		Empty any of the jugs.
// 		Pour water from one jug into another till the other jug is completely full, or the first jug itself is empty.
 
// Example 1:
// Input: jug1Capacity = 3, jug2Capacity = 5, targetCapacity = 4
// Output: true
// Explanation: The famous Die Hard example 

// Example 2:
// Input: jug1Capacity = 2, jug2Capacity = 6, targetCapacity = 5
// Output: false

// Example 3:
// Input: jug1Capacity = 1, jug2Capacity = 2, targetCapacity = 3
// Output: true
 

// Constraints:

// 		1 <= jug1Capacity, jug2Capacity, targetCapacity <= 10^6

// 解题思路：

// 		求最大公约数GCD（Greatest Common Divisor)。
// 		如果x与y互质（最大公约数为1），则容量范围[1, x + y]之内的任意整数体积均可以通过适当的操作得到。
// 		否则，记x与y的最大公约数为gcd，则可以获得的容量z只能为gcd的整数倍，且z <= x + y

// 简单的证明：

// 		假设最终体积z = m * x + n * y（m与n为整数，可以为0或者为负）
// 		令x = p * gcd, y = q * gcd，可知p与q互质。
// 		则z = (m * p + n * q) * gcd
// 		可以证明一定存在m, n，使得m * p + n * q = 1（p与q互质的性质，参见：裴蜀定理）
// 		此可知z一定是gcd的整数倍

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {

        if jug1Capacity > jug2Capacity {
			jug1Capacity, jug2Capacity = jug2Capacity, jug1Capacity
		}
		// 求最大公约数GCD（Greatest Common Divisor)
		// 如果x与y互质（最大公约数为1），则容量范围[1, x + y]之内的任意整数体积均可以通过适当的操作得到
        gcd := gcd(jug1Capacity, jug2Capacity)
        if gcd == 0 {
			return targetCapacity == 0
		}
		// 否则，记x与y的最大公约数为gcd，则可以获得的容量z只能为gcd的整数倍，且z <= x + y
        return (targetCapacity % gcd == 0 ) && (targetCapacity <= jug1Capacity + jug2Capacity)
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b % a, a)
}

func main() {
	fmt.Println(canMeasureWater(3,5,4)) // ture
	fmt.Println(canMeasureWater(2,6,5)) // false
	fmt.Println(canMeasureWater(1,2,3)) // true
}