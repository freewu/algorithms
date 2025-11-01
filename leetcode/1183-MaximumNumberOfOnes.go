package main

// 1183. Maximum Number of Ones
// Consider a matrix M with dimensions width * height, such that every cell has value 0 or 1, 
// and any square sub-matrix of M of size sideLength * sideLength has at most maxOnes ones.
// Return the maximum possible number of ones that the matrix M can have.

// Example 1:
// Input: width = 3, height = 3, sideLength = 2, maxOnes = 1
// Output: 4
// Explanation:
// In a 3*3 matrix, no 2*2 sub-matrix can have more than 1 one.
// The best solution that has 4 ones is:
// [1,0,1]
// [0,0,0]
// [1,0,1]

// Example 2:
// Input: width = 3, height = 3, sideLength = 2, maxOnes = 2
// Output: 6
// Explanation:
// [1,0,1]
// [1,0,1]
// [1,0,1]

// Constraints:
//     1 <= width, height <= 100
//     1 <= sideLength <= width, height
//     0 <= maxOnes <= sideLength * sideLength

import "fmt"
import "sort"

func maximumNumberOfOnes(width int, height int, sideLength int, maxOnes int) int {
    sum, arr := 0, []int{}
    wq, wr, hq, hr := width / sideLength, width % sideLength, height / sideLength, height % sideLength
    for i := 0; i < sideLength; i++ {
        for j := 0; j < sideLength; j++ {
            ww,hh := wq, hq
            if i < wr {
                ww++
            }
            if j < hr{
                hh++
            }
            arr = append(arr, ww * hh)
        }
    }
    sort.Ints(arr)
    for i := 0; i < maxOnes && i < len(arr); i++ {
        sum += arr[len(arr) - 1 - i]
    }
    return sum
}


func main() {
    // Example 1:
    // Input: width = 3, height = 3, sideLength = 2, maxOnes = 1
    // Output: 4
    // Explanation:
    // In a 3*3 matrix, no 2*2 sub-matrix can have more than 1 one.
    // The best solution that has 4 ones is:
    // [1,0,1]
    // [0,0,0]
    // [1,0,1]
    fmt.Println(maximumNumberOfOnes(3,3,2,1)) // 4
    // Example 2:
    // Input: width = 3, height = 3, sideLength = 2, maxOnes = 2
    // Output: 6
    // Explanation:
    // [1,0,1]
    // [1,0,1]
    // [1,0,1]
    fmt.Println(maximumNumberOfOnes(3,3,2,2)) // 6

    fmt.Println(maximumNumberOfOnes(1,2,3,4)) // 2
    fmt.Println(maximumNumberOfOnes(9,8,7,6)) // 16
    fmt.Println(maximumNumberOfOnes(5,5,5,5)) // 5
    fmt.Println(maximumNumberOfOnes(1,1,1,0)) // 0
    fmt.Println(maximumNumberOfOnes(1,1,1,10000)) // 1
    fmt.Println(maximumNumberOfOnes(100,1,100,10000)) // 100
    fmt.Println(maximumNumberOfOnes(100,100,1,10000)) // 10000
    fmt.Println(maximumNumberOfOnes(100,100,100,0)) // 0
    fmt.Println(maximumNumberOfOnes(1,1,100,10000)) // 1
    fmt.Println(maximumNumberOfOnes(100,1,1,10000)) // 100
    fmt.Println(maximumNumberOfOnes(100,100,1,0)) // 0
    fmt.Println(maximumNumberOfOnes(100,100,100,10000)) // 10000
}