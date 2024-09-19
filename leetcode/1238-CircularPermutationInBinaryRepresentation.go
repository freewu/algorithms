package main

// 1238. Circular Permutation in Binary Representation
// Given 2 integers n and start. Your task is return any permutation p of (0,1,2.....,2^n -1) such that :
//     p[0] = start
//     p[i] and p[i+1] differ by only one bit in their binary representation.
//     p[0] and p[2^n -1] must also differ by only one bit in their binary representation.

// Example 1:
// Input: n = 2, start = 3
// Output: [3,2,0,1]
// Explanation: The binary representation of the permutation is (11,10,00,01). 
// All the adjacent element differ by one bit. Another valid permutation is [3,1,0,2]

// Example 2:
// Input: n = 3, start = 2
// Output: [2,6,7,5,4,0,1,3]
// Explanation: The binary representation of the permutation is (010,110,111,101,100,000,001,011).

// Constraints:
//     1 <= n <= 16
//     0 <= start < 2 ^ n

import "fmt"

func circularPermutation(n int, start int) []int {
    res, count := []int{ start }, 1
    for i := 0; i < n; i++ {
        for j := count - 1; j >= 0; j-- {
            res = append(res, res[j] ^ count)
        }
        count <<= 1
    }
    return res
}

func circularPermutation1(n int, start int) []int {
    res := []int{}
    for i := 0; i < 1<<n; i++ {
        res = append(res, i ^ ( i >> 1) ^ start)
    }
    return res
}

func circularPermutation2(n int, start int) []int {
    mirror := func(nums []int, idx int) {
        for i := 0; i < idx; i++ {
            nums[idx+i] = nums[idx-i-1]
        }
    }
    flip := func (nums []int, idx int) {
        n := 1 << idx
        for i := 0; i < n; i++ {
            nums[i+n] |= 1 << idx
        }
    }
    permute := func(n int) []int {
        res := make([]int, 1<<n)
        for i := 0; i < n; i++ {
            mirror(res, 1<<i)
            flip(res, i)
        }
        return res
    }
    res := permute(n)
    index := 0
    for i := range res {
        if res[i] == start {
            index = i
            break
        }
    }
    for i := 0; i < index; i++ {
        res = append(res, res[i])
    }
    return res[index:]
}

func main() {
    // Example 1:
    // Input: n = 2, start = 3
    // Output: [3,2,0,1]
    // Explanation: The binary representation of the permutation is (11,10,00,01). 
    // All the adjacent element differ by one bit. Another valid permutation is [3,1,0,2]
    fmt.Println(circularPermutation(2,3)) // [3,2,0,1]
    // Example 2:
    // Input: n = 3, start = 2
    // Output: [2,6,7,5,4,0,1,3]
    // Explanation: The binary representation of the permutation is (010,110,111,101,100,000,001,011).
    fmt.Println(circularPermutation(3,2)) // [2,6,7,5,4,0,1,3]

    fmt.Println(circularPermutation(1,0)) // [0 1]
    //fmt.Println(circularPermutation(16,0))
    fmt.Println(circularPermutation(4,64)) // [64 65 67 66 70 71 69 68 76 77 79 78 74 75 73 72]


    fmt.Println(circularPermutation1(2,3)) // [3,2,0,1]
    fmt.Println(circularPermutation1(3,2)) // [2,6,7,5,4,0,1,3]
    fmt.Println(circularPermutation1(1,0)) // [0 1]
    //fmt.Println(circularPermutation1(16,0))
    fmt.Println(circularPermutation1(4,64)) // [64 65 67 66 70 71 69 68 76 77 79 78 74 75 73 72]

    fmt.Println(circularPermutation2(2,3)) // [3,2,0,1]
    fmt.Println(circularPermutation2(3,2)) // [2,6,7,5,4,0,1,3]
    fmt.Println(circularPermutation2(1,0)) // [0 1]
    //fmt.Println(circularPermutation2(16,0))
    fmt.Println(circularPermutation2(4,64)) // [64 65 67 66 70 71 69 68 76 77 79 78 74 75 73 72]
}
