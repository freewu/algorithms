package main

// 1711. Count Good Meals
// A good meal is a meal that contains exactly two different food items with a sum of deliciousness equal to a power of two.

// You can pick any two different foods to make a good meal.

// Given an array of integers deliciousness where deliciousness[i] is the deliciousness of the i​​​​​​th​​​​​​​​ item of food, 
// return the number of different good meals you can make from this list modulo 10^9 + 7.

// Note that items with different indices are considered different even if they have the same deliciousness value.

// Example 1:
// Input: deliciousness = [1,3,5,7,9]
// Output: 4
// Explanation: The good meals are (1,3), (1,7), (3,5) and, (7,9).
// Their respective sums are 4, 8, 8, and 16, all of which are powers of 2.

// Example 2:
// Input: deliciousness = [1,1,1,3,3,3,7]
// Output: 15
// Explanation: The good meals are (1,1) with 3 ways, (1,3) with 9 ways, and (1,7) with 3 ways.

// Constraints:
//     1 <= deliciousness.length <= 10^5
//     0 <= deliciousness[i] <= 2^20

import "fmt"
import "sort"
import "math/bits"

// 超出时间限制 63 / 72 
func countPairs(deliciousness []int) int {
    sort.Ints(deliciousness)
    count, fin, i, j := 0, 0, len(deliciousness) - 1, len(deliciousness) - 2
    for i > 0 {
        n := deliciousness[i] + deliciousness[j]
        if fin == 0 {
            for n > 1 {
                if n % 2 == 0 {
                    n = n / 2
                } else {
                    break
                }
            }
            if n == 1 {
                fin = deliciousness[i] + deliciousness[j]
                count++
            }
        } else {
            if n != 0 && fin % n == 0 {
                count++
            }
        }
        if j == 0 {
            i--
            j = i - 1
        } else {
            j--
        }
    }
    return count
}

func countPairs1(deliciousness []int) int {
    res, mp := 0, make(map[int]int)
    for _, v := range deliciousness {
        mp[v] += 1
    }
    for k, _ := range mp {
        // no need to process the 0 key, it will be processed
        // in the next line (the power of two case)
        if k == 0 { continue }
        // If the current N is a power of two, then bitwise
        // operation of (N and (N-1)) will equal to 0
        if (k & (k - 1)) == 0 {
            if _, ok := mp[0]; ok {
                res += (mp[k] * mp[0])
            }
        }
        // mask is the next bigger power of 2 after key, so
        // we'll shift 1 as many as the bit length of key
        mask := (1 << bits.Len(uint(k)))
        pair := mask - k
        if _, ok := mp[pair]; !ok { continue }
        // if the pair is its own self use the (m * (m-1))/2
        // else use (m * n)
        if k == pair {
            res += (mp[k] * (mp[pair] - 1) / 2 )
        } else {
            res += (mp[k] * mp[pair])
        }
    }
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: deliciousness = [1,3,5,7,9]
    // Output: 4
    // Explanation: The good meals are (1,3), (1,7), (3,5) and, (7,9).
    // Their respective sums are 4, 8, 8, and 16, all of which are powers of 2.
    fmt.Println(countPairs([]int{1,3,5,7,9})) // 4
    // Example 2:
    // Input: deliciousness = [1,1,1,3,3,3,7]
    // Output: 15
    // Explanation: The good meals are (1,1) with 3 ways, (1,3) with 9 ways, and (1,7) with 3 ways.
    fmt.Println(countPairs([]int{1,1,1,3,3,3,7})) // 15

    fmt.Println(countPairs1([]int{1,3,5,7,9})) // 4
    fmt.Println(countPairs1([]int{1,1,1,3,3,3,7})) // 15
}