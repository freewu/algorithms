package main

// 1052. Grumpy Bookstore Owner
// There is a bookstore owner that has a store open for n minutes. 
// Every minute, some number of customers enter the store. 
// You are given an integer array customers of length n where customers[i] is the number of the customer 
// that enters the store at the start of the ith minute and all those customers leave after the end of that minute.

// On some minutes, the bookstore owner is grumpy. 
// You are given a binary array grumpy where grumpy[i] is 1 if the bookstore owner is grumpy during the ith minute, and is 0 otherwise.

// When the bookstore owner is grumpy, the customers of that minute are not satisfied, 
// otherwise, they are satisfied.

// The bookstore owner knows a secret technique to keep themselves not grumpy for minutes consecutive minutes, but can only use it once.
// Return the maximum number of customers that can be satisfied throughout the day.

// Example 1:
// Input: customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], minutes = 3
// Output: 16
// Explanation: The bookstore owner keeps themselves not grumpy for the last 3 minutes. 
// The maximum number of customers that can be satisfied = 1 + 1 + 1 + 1 + 7 + 5 = 16.

// Example 2:
// Input: customers = [1], grumpy = [0], minutes = 1
// Output: 1
 
// Constraints:
//     n == customers.length == grumpy.length
//     1 <= minutes <= n <= 2 * 10^4
//     0 <= customers[i] <= 1000
//     grumpy[i] is either 0 or 1.

import "fmt"

func maxSatisfied(customers []int, grumpy []int, k int) int {
    n := len(customers)
    unsat, sat := make([]int, n), 0
    for i, c := range customers {
        if grumpy[i] == 1 { // 当书店老板生气时 grumpy[i] = 1 ，那一分钟的顾客就会不满意，若老板不生气则顾客是满意的
            unsat[i] += c
        } else {
            sat += c
        }
        if i > 0 {
            unsat[i] += unsat[i - 1]
        }
    }
    best, tail := 0, 0
    for i := 0; i <= n - k; i++ {
        cur := unsat[i + k - 1] - tail
        if cur > best { 
            best = cur 
        }
        tail = unsat[i]
    }
    return best + sat
}

// 双指针
func maxSatisfied1(customers []int, grumpy []int, minutes int) int {
    sum := 0 
    for i, v := range customers {
        if grumpy[i] == 0 {
            sum += v 
        }
    }
    res, tmp, left, right, n := 0, 0, 0, 0, len(customers)
    for right < n {
        if grumpy[right] == 1 {
            tmp += customers[right] * grumpy[right]
        }
        right++ 
        if right - left == minutes {
            res = max(res, sum + tmp)
            tmp -= customers[left] * grumpy[left] 
            left++ 
        }
    }
    return res 
}

func main() {
    // Explanation: The bookstore owner keeps themselves not grumpy for the last 3 minutes. 
    // The maximum number of customers that can be satisfied = 1 + 1 + 1 + 1 + 7 + 5 = 16.
    fmt.Println(maxSatisfied([]int{1,0,1,2,1,1,7,5},[]int{0,1,0,1,0,1,0,1},3)) // 16
    fmt.Println(maxSatisfied([]int{1},[]int{0},1)) // 1

    fmt.Println(maxSatisfied1([]int{1,0,1,2,1,1,7,5},[]int{0,1,0,1,0,1,0,1},3)) // 16
    fmt.Println(maxSatisfied1([]int{1},[]int{0},1)) // 1
}