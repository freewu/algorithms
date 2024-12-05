package main

// 2729. Check if The Number is Fascinating
// You are given an integer n that consists of exactly 3 digits.

// We call the number n fascinating if, 
// after the following modification, 
// the resulting number contains all the digits from 1 to 9 exactly once and does not contain any 0's:
//     Concatenate n with the numbers 2 * n and 3 * n.

// Return true if n is fascinating, or false otherwise.

// Concatenating two numbers means joining them together. 
// For example, the concatenation of 121 and 371 is 121371.

// Example 1:
// Input: n = 192
// Output: true
// Explanation: We concatenate the numbers n = 192 and 2 * n = 384 and 3 * n = 576. The resulting number is 192384576. This number contains all the digits from 1 to 9 exactly once.

// Example 2:
// Input: n = 100
// Output: false
// Explanation: We concatenate the numbers n = 100 and 2 * n = 200 and 3 * n = 300. The resulting number is 100200300. This number does not satisfy any of the conditions.

// Constraints:
//     100 <= n <= 999

import "fmt"
import "strconv"
import "sort"

func isFascinating(n int) bool {
    a, b := n * 2, n * 3
    strA, strB, strN := strconv.Itoa(a), strconv.Itoa(b), strconv.Itoa(n)
    conArr, concatVal := []int{}, strN + strA + strB
    for _, v := range concatVal {
        conArr = append(conArr, int(v - '0'))
    }
    sort.Ints(conArr)
    for i := 0; i < len(conArr); i++ {
        if i + 1 == conArr[i] { continue }
        return false
    }
    return true
}

func isFascinating1(n int) bool {
    count, t := [10]int{}, n
    for ; t > 0; t /= 10 {
        count[t % 10]++
    }
    t = 2 * n
    for ; t > 0; t /= 10 {
        count[t % 10]++
    }
    t = 3 * n
    for ; t > 0; t /= 10 {
        count[t % 10]++
    }
    for i := 1; i < 10; i++ {
        if count[i] != 1 { return false }
    }
    return true
}

func main() {
    // Example 1:
    // Input: n = 192
    // Output: true
    // Explanation: We concatenate the numbers n = 192 and 2 * n = 384 and 3 * n = 576. The resulting number is 192384576. This number contains all the digits from 1 to 9 exactly once.
    fmt.Println(isFascinating(192)) // true
    // Example 2:
    // Input: n = 100
    // Output: false
    // Explanation: We concatenate the numbers n = 100 and 2 * n = 200 and 3 * n = 300. The resulting number is 100200300. This number does not satisfy any of the conditions.
    fmt.Println(isFascinating(100)) // false

    fmt.Println(isFascinating(101)) // false
    fmt.Println(isFascinating(666)) // false
    fmt.Println(isFascinating(888)) // false
    fmt.Println(isFascinating(998)) // false
    fmt.Println(isFascinating(999)) // false

    fmt.Println(isFascinating1(192)) // true
    fmt.Println(isFascinating1(100)) // false
    fmt.Println(isFascinating1(101)) // false
    fmt.Println(isFascinating1(666)) // false
    fmt.Println(isFascinating1(888)) // false
    fmt.Println(isFascinating1(998)) // false
    fmt.Println(isFascinating1(999)) // false
}