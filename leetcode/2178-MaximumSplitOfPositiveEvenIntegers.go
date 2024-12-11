package main

// 2178. Maximum Split of Positive Even Integers
// You are given an integer finalSum. Split it into a sum of a maximum number of unique positive even integers.
//     For example, given finalSum = 12, the following splits are valid (unique positive even integers summing up to finalSum):(12), (2 + 10), (2 + 4 + 6), and (4 + 8). 
//     Among them, (2 + 4 + 6) contains the maximum number of integers. 
//     Note that finalSum cannot be split into (2 + 2 + 4 + 4) as all the numbers should be unique.

// Return a list of integers that represent a valid split containing a maximum number of integers. 
// If no valid split exists for finalSum, return an empty list. 
// You may return the integers in any order.

// Example 1:
// Input: finalSum = 12
// Output: [2,4,6]
// Explanation: The following are valid splits: (12), (2 + 10), (2 + 4 + 6), and (4 + 8).
// (2 + 4 + 6) has the maximum number of integers, which is 3. Thus, we return [2,4,6].
// Note that [2,6,4], [6,2,4], etc. are also accepted.

// Example 2:
// Input: finalSum = 7
// Output: []
// Explanation: There are no valid splits for the given finalSum.
// Thus, we return an empty array.

// Example 3:
// Input: finalSum = 28
// Output: [6,8,2,12]
// Explanation: The following are valid splits: (2 + 26), (6 + 8 + 2 + 12), and (4 + 24). 
// (6 + 8 + 2 + 12) has the maximum number of integers, which is 4. Thus, we return [6,8,2,12].
// Note that [10,2,4,12], [6,2,4,16], etc. are also accepted.

// Constraints:
//     1 <= finalSum <= 10^10

import "fmt"

func maximumEvenSplit(finalSum int64) []int64 {
    if finalSum % 2 != 0 { return []int64{} }
    res, v := make([]int64, 0), int64(2)
    for finalSum > 0 {
        finalSum -= v
        res = append(res, v)
        v += 2
    }
    if finalSum == 0 { return res }
    finalSum *= -1 // make finalSum positive by multiply by -1
    toDelete := int(finalSum / 2 - 1) // calculate index of element to delete.  e.g. [2,4,6,8] index of 6 will be 6/2-1=2
    return append(res[:toDelete], res[toDelete + 1:]...) // remove element with index toDelete from result
}

func maximumEvenSplit1(finalSum int64) []int64 {
    if finalSum % 2 == 1 { return nil  }
    res, i := []int64{}, int64(2)
    for finalSum > 2*i {
        finalSum -= i
        res = append(res, i)
        i += 2
    }
    res = append(res, finalSum)
    return res
}

func main() {
    // Example 1:
    // Input: finalSum = 12
    // Output: [2,4,6]
    // Explanation: The following are valid splits: (12), (2 + 10), (2 + 4 + 6), and (4 + 8).
    // (2 + 4 + 6) has the maximum number of integers, which is 3. Thus, we return [2,4,6].
    // Note that [2,6,4], [6,2,4], etc. are also accepted.
    fmt.Println(maximumEvenSplit(12)) // [2,4,6]
    // Example 2:
    // Input: finalSum = 7
    // Output: []
    // Explanation: There are no valid splits for the given finalSum.
    // Thus, we return an empty array.
    fmt.Println(maximumEvenSplit(7)) // []
    // Example 3:
    // Input: finalSum = 28
    // Output: [6,8,2,12]
    // Explanation: The following are valid splits: (2 + 26), (6 + 8 + 2 + 12), and (4 + 24). 
    // (6 + 8 + 2 + 12) has the maximum number of integers, which is 4. Thus, we return [6,8,2,12].
    // Note that [10,2,4,12], [6,2,4,16], etc. are also accepted.
    fmt.Println(maximumEvenSplit(28)) // [4 6 8 10]

    fmt.Println(maximumEvenSplit(1)) // []
    fmt.Println(maximumEvenSplit(1024)) // [2 4 6 8 10 12 14 16 18 20 22 24 26 28 30 34 36 38 40 42 44 46 48 50 52 54 56 58 60 62 64]
    //fmt.Println(maximumEvenSplit(1e11)) // 

    fmt.Println(maximumEvenSplit1(12)) // [2,4,6]
    fmt.Println(maximumEvenSplit1(7)) // []
    fmt.Println(maximumEvenSplit1(28)) // [2 4 6 16]
    fmt.Println(maximumEvenSplit1(1)) // []
    fmt.Println(maximumEvenSplit1(1024)) // [2 4 6 8 10 12 14 16 18 20 22 24 26 28 30 32 34 36 38 40 42 44 46 48 50 52 54 56 58 60 94]
}