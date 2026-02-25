package main

// 1356. Sort Integers by The Number of 1 Bits
// You are given an integer array arr. 
// Sort the integers in the array in ascending order by the number of 1's in their binary representation 
// and in case of two or more integers have the same number of 1's you have to sort them in ascending order.

// Return the array after sorting it.

// Example 1:
// Input: arr = [0,1,2,3,4,5,6,7,8]
// Output: [0,1,2,4,8,3,5,6,7]
// Explantion: [0] is the only integer with 0 bits.
// [1,2,4,8] all have 1 bit.
// [3,5,6] have 2 bits.
// [7] has 3 bits.
// The sorted array by bits is [0,1,2,4,8,3,5,6,7]

// Example 2:
// Input: arr = [1024,512,256,128,64,32,16,8,4,2,1]
// Output: [1,2,4,8,16,32,64,128,256,512,1024]
// Explantion: All integers have 1 bit in the binary representation, you should just sort them in ascending order.

// Constraints:
//     1 <= arr.length <= 500
//     0 <= arr[i] <= 10^4

import "fmt"
import "sort"
import "math/bits"

func sortByBits(arr []int) []int {
    getBitOneCount := func(val int) int { // 得到 1 的个数
        res := 0
        for val != 0 {
            res += val & 0x1
            val >>= 1
        }
        return res
    }
    sort.Slice(arr, func(i, j int) bool {
        ic, jc := getBitOneCount(arr[i]), getBitOneCount(arr[j])
        if ic == jc { // 1 一样 值从小到大
            return arr[i] < arr[j]
        }
        return ic < jc
    })
    return arr
}

func sortByBits1(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        ic, jc := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
        if ic == jc { // 1 一样 值从小到大
            return arr[i] < arr[j]
        }
        return ic < jc
    })
    return arr
}

// bits.OnesCount(uint(b[i])) < bits.OnesCount(uint(b[j]))

func main() {
    // Example 1:
    // Input: arr = [0,1,2,3,4,5,6,7,8]
    // Output: [0,1,2,4,8,3,5,6,7]
    // Explantion: [0] is the only integer with 0 bits.
    // [1,2,4,8] all have 1 bit.
    // [3,5,6] have 2 bits.
    // [7] has 3 bits.
    // The sorted array by bits is [0,1,2,4,8,3,5,6,7]
    fmt.Println(sortByBits([]int{0,1,2,3,4,5,6,7,8})) // [0,1,2,4,8,3,5,6,7]
    // Example 2:
    // Input: arr = [1024,512,256,128,64,32,16,8,4,2,1]
    // Output: [1,2,4,8,16,32,64,128,256,512,1024]
    // Explantion: All integers have 1 bit in the binary representation, you should just sort them in ascending order.
    fmt.Println(sortByBits([]int{1024,512,256,128,64,32,16,8,4,2,1})) // [1,2,4,8,16,32,64,128,256,512,1024]

    fmt.Println(sortByBits([]int{1,2,3,4,5,6,7,8,9})) // [1 2 4 8 3 5 6 9 7]
    fmt.Println(sortByBits([]int{9,8,7,6,5,4,3,2,1})) // [1 2 4 8 3 5 6 9 7]

    fmt.Println(sortByBits1([]int{0,1,2,3,4,5,6,7,8})) // [0,1,2,4,8,3,5,6,7]
    fmt.Println(sortByBits1([]int{1024,512,256,128,64,32,16,8,4,2,1})) // [1,2,4,8,16,32,64,128,256,512,1024]
    fmt.Println(sortByBits1([]int{1,2,3,4,5,6,7,8,9})) // [1 2 4 8 3 5 6 9 7]
    fmt.Println(sortByBits1([]int{9,8,7,6,5,4,3,2,1})) // [1 2 4 8 3 5 6 9 7]
}