package main

// 888. Fair Candy Swap
// Alice and Bob have a different total number of candies. 
// You are given two integer arrays aliceSizes and bobSizes where aliceSizes[i] is the number of candies of the ith box of candy 
// that Alice has and bobSizes[j] is the number of candies of the jth box of candy that Bob has.

// Since they are friends, they would like to exchange one candy box each so that after the exchange, 
// they both have the same total amount of candy. 
// The total amount of candy a person has is the sum of the number of candies in each box they have.

// Return an integer array answer where answer[0] is the number of candies in the box that Alice must exchange, 
// and answer[1] is the number of candies in the box that Bob must exchange. 
// If there are multiple answers, you may return any one of them. It is guaranteed that at least one answer exists.

// Example 1:
// Input: aliceSizes = [1,1], bobSizes = [2,2]
// Output: [1,2]

// Example 2:
// Input: aliceSizes = [1,2], bobSizes = [2,3]
// Output: [1,2]

// Example 3:
// Input: aliceSizes = [2], bobSizes = [1,3]
// Output: [2,3]

// Constraints:
//     1 <= aliceSizes.length, bobSizes.length <= 10^4
//     1 <= aliceSizes[i], bobSizes[j] <= 10^5
//     Alice and Bob have a different total number of candies.
//     There will be at least one valid answer for the given input.

import "fmt"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    sumArray := func(arr []int) int {
        res := 0
        for _, v := range arr {
            res += v
        }
        return res
    }
    sum1 := sumArray(aliceSizes)
    sum := sum1 + sumArray(bobSizes)
    mp, res := make(map[int]int), []int{}
    for i, v := range bobSizes {
        mp[v] = i
    }
    for _, v := range aliceSizes {
        need := sum / 2 - (sum1 - v)
        if _, ok := mp[need]; ok {
            res = append(res, v, need)
            break
        }
    }
    return res
}

func fairCandySwap1(aliceSizes []int, bobSizes []int) []int {
    asum, bsum, mp := 0, 0, make(map[int]bool, len(bobSizes))
    for _, v := range aliceSizes {
        asum += v
    }
    for _, v := range bobSizes {
        bsum += v
        mp[v] = true
    }
    k := (asum + bsum) / 2
    for _, v := range aliceSizes {
        t := k + v - asum
        if mp[t] {
            return []int{v, t}
        }
    }
    return nil
}

func main() {
    // Example 1:
    // Input: aliceSizes = [1,1], bobSizes = [2,2]
    // Output: [1,2]
    fmt.Println(fairCandySwap([]int{1,1}, []int{2,2})) // [1,2]
    // Example 2:
    // Input: aliceSizes = [1,2], bobSizes = [2,3]
    // Output: [1,2]
    fmt.Println(fairCandySwap([]int{1,2}, []int{2,3})) // [1,2]
    // Example 3:
    // Input: aliceSizes = [2], bobSizes = [1,3]
    // Output: [2,3]
    fmt.Println(fairCandySwap([]int{2}, []int{1,3})) // [2,3]

    fmt.Println(fairCandySwap1([]int{1,1}, []int{2,2})) // [1,2]
    fmt.Println(fairCandySwap1([]int{1,2}, []int{2,3})) // [1,2]
    fmt.Println(fairCandySwap1([]int{2}, []int{1,3})) // [2,3]
}