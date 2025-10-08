package main

// 2300. Successful Pairs of Spells and Potions
// You are given two positive integer arrays spells and potions, of length n and m respectively, 
// where spells[i] represents the strength of the ith spell and potions[j] represents the strength of the jth potion.

// You are also given an integer success. 
// A spell and potion pair is considered successful if the product of their strengths is at least success.

// Return an integer array pairs of length n where pairs[i] is the number of potions that will form a successful pair with the ith spell.

// Example 1:
// Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7
// Output: [4,0,3]
// Explanation:
// - 0th spell: 5 * [1,2,3,4,5] = [5,10,15,20,25]. 4 pairs are successful.
// - 1st spell: 1 * [1,2,3,4,5] = [1,2,3,4,5]. 0 pairs are successful.
// - 2nd spell: 3 * [1,2,3,4,5] = [3,6,9,12,15]. 3 pairs are successful.

// Thus, [4,0,3] is returned.
// Example 2:
// Input: spells = [3,1,2], potions = [8,5,8], success = 16
// Output: [2,0,2]
// Explanation:
// - 0th spell: 3 * [8,5,8] = [24,15,24]. 2 pairs are successful.
// - 1st spell: 1 * [8,5,8] = [8,5,8]. 0 pairs are successful. 
// - 2nd spell: 2 * [8,5,8] = [16,10,16]. 2 pairs are successful. 
// Thus, [2,0,2] is returned.

// Constraints:
//     n == spells.length
//     m == potions.length
//     1 <= n, m <= 10^5
//     1 <= spells[i], potions[i] <= 10^5
//     1 <= success <= 10^10

import "fmt"
import "sort"

// 暴力破解
func successfulPairs(spells []int, potions []int, success int64) []int {
    res := []int{}
    for i := 0; i < len(spells); i++ {
        count := 0
        for j := 0; j < len(potions); j++ {
            if int64(spells[i] * potions[j]) >= success {
                count++
            } 
        }
        res = append(res,count)
    }
    return res
}

// 先排序，判断到小的直接出本次循环 O(n^2)
func successfulPairs1(spells []int, potions []int, success int64) []int {
    res := []int{}
    sort.Ints(potions)
    for i := 0; i < len(spells); i++ {
        count := 0
        for j := len(potions) - 1; j >= 0 ; j-- {
            if int64(spells[i] * potions[j]) < success {
                break
            } 
            count++
        }
        res = append(res,count)
    }
    return res
}

// 二分法
func successfulPairs2(spells []int, potions []int, success int64) []int {
    res := make([]int, len(spells))
    sort.Ints(potions)
    for k, v := range spells {
        left, right := 0, len(potions) - 1 
        for left <= right {
            mid := (left+right) / 2
            if int64(potions[mid] * v) >= success {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
        res[k] = len(potions) - left
    }
    return res
}

func main() {
    // Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7
    // Output: [4,0,3]
    // Explanation:
    // - 0th spell: 5 * [1,2,3,4,5] = [5,10,15,20,25]. 4 pairs are successful.
    // - 1st spell: 1 * [1,2,3,4,5] = [1,2,3,4,5]. 0 pairs are successful.
    // - 2nd spell: 3 * [1,2,3,4,5] = [3,6,9,12,15]. 3 pairs are successful.
    fmt.Println(successfulPairs([]int{5,1,3}, []int{1,2,3,4,5}, 7)) // [4,0,3]
    // - 0th spell: 3 * [8,5,8] = [24,15,24]. 2 pairs are successful.
    // - 1st spell: 1 * [8,5,8] = [8,5,8]. 0 pairs are successful. 
    // - 2nd spell: 2 * [8,5,8] = [16,10,16]. 2 pairs are successful. 
    // Thus, [2,0,2] is returned.
    fmt.Println(successfulPairs([]int{3,1,2}, []int{8,5,8}, 17)) // [2,0,2]

    fmt.Println(successfulPairs([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 17)) // [0 1 4 5 6 7 7 7 8]
    fmt.Println(successfulPairs([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 17)) // [0 1 4 5 6 7 7 7 8]
    fmt.Println(successfulPairs([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 17)) // [8 7 7 7 6 5 4 1 0]
    fmt.Println(successfulPairs([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 17)) // [8 7 7 7 6 5 4 1 0]

    fmt.Println(successfulPairs1([]int{5,1,3}, []int{1,2,3,4,5}, 7)) // [4,0,3]
    fmt.Println(successfulPairs1([]int{3,1,2}, []int{8,5,8}, 17)) // [2,0,2]
    fmt.Println(successfulPairs1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 17)) // [0 1 4 5 6 7 7 7 8]
    fmt.Println(successfulPairs1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 17)) // [0 1 4 5 6 7 7 7 8]
    fmt.Println(successfulPairs1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 17)) // [8 7 7 7 6 5 4 1 0]
    fmt.Println(successfulPairs1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 17)) // [8 7 7 7 6 5 4 1 0]

    fmt.Println(successfulPairs2([]int{5,1,3}, []int{1,2,3,4,5}, 7)) // [4,0,3]
    fmt.Println(successfulPairs2([]int{3,1,2}, []int{8,5,8}, 17)) // [2,0,2]
    fmt.Println(successfulPairs2([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 17)) // [0 1 4 5 6 7 7 7 8]
    fmt.Println(successfulPairs2([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 17)) // [0 1 4 5 6 7 7 7 8]
    fmt.Println(successfulPairs2([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 17)) // [8 7 7 7 6 5 4 1 0]
    fmt.Println(successfulPairs2([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 17)) // [8 7 7 7 6 5 4 1 0]
}