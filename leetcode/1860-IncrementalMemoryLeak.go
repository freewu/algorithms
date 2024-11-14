package main

// 1860. Incremental Memory Leak
// You are given two integers memory1 and memory2 representing the available memory in bits on two memory sticks. 
// There is currently a faulty program running that consumes an increasing amount of memory every second.

// At the ith second (starting from 1), i bits of memory are allocated to the stick with more available memory 
// (or from the first memory stick if both have the same available memory). 
// If neither stick has at least i bits of available memory, the program crashes.

// Return an array containing [crashTime, memory1crash, memory2crash], 
// where crashTime is the time (in seconds) 
// when the program crashed and memory1crash and memory2crash are the available bits of memory in the first and second sticks respectively.

// Example 1:
// Input: memory1 = 2, memory2 = 2
// Output: [3,1,0]
// Explanation: The memory is allocated as follows:
// - At the 1st second, 1 bit of memory is allocated to stick 1. The first stick now has 1 bit of available memory.
// - At the 2nd second, 2 bits of memory are allocated to stick 2. The second stick now has 0 bits of available memory.
// - At the 3rd second, the program crashes. The sticks have 1 and 0 bits available respectively.

// Example 2:
// Input: memory1 = 8, memory2 = 11
// Output: [6,0,4]
// Explanation: The memory is allocated as follows:
// - At the 1st second, 1 bit of memory is allocated to stick 2. The second stick now has 10 bit of available memory.
// - At the 2nd second, 2 bits of memory are allocated to stick 2. The second stick now has 8 bits of available memory.
// - At the 3rd second, 3 bits of memory are allocated to stick 1. The first stick now has 5 bits of available memory.
// - At the 4th second, 4 bits of memory are allocated to stick 2. The second stick now has 4 bits of available memory.
// - At the 5th second, 5 bits of memory are allocated to stick 1. The first stick now has 0 bits of available memory.
// - At the 6th second, the program crashes. The sticks have 0 and 4 bits available respectively.

// Constraints:
//     0 <= memory1, memory2 <= 2^31 - 1

import "fmt"

func memLeak(memory1 int, memory2 int) []int {
    i := 1
    for {
        if memory1 < memory2 { // pick memory2
            if i > memory2 { break }
            memory2 -= i
        } else { // pick memory1
            if i > memory1 { break } 
            memory1 -= i
        }
        i++
    } 
    return []int{ i, memory1, memory2 }
}

func memLeak1(memory1 int, memory2 int) []int {
    i := 1
    for ; i <= memory1 || i <= memory2; i++ {
        if memory1 >= memory2 {
            memory1 -= i
        } else {
            memory2 -= i
        }
    }
    return []int{ i, memory1, memory2 }
}

func main() {
    // Example 1:
    // Input: memory1 = 2, memory2 = 2
    // Output: [3,1,0]
    // Explanation: The memory is allocated as follows:
    // - At the 1st second, 1 bit of memory is allocated to stick 1. The first stick now has 1 bit of available memory.
    // - At the 2nd second, 2 bits of memory are allocated to stick 2. The second stick now has 0 bits of available memory.
    // - At the 3rd second, the program crashes. The sticks have 1 and 0 bits available respectively.
    fmt.Println(memLeak(2,2)) // [3,1,0]
    // Example 2:
    // Input: memory1 = 8, memory2 = 11
    // Output: [6,0,4]
    // Explanation: The memory is allocated as follows:
    // - At the 1st second, 1 bit of memory is allocated to stick 2. The second stick now has 10 bit of available memory.
    // - At the 2nd second, 2 bits of memory are allocated to stick 2. The second stick now has 8 bits of available memory.
    // - At the 3rd second, 3 bits of memory are allocated to stick 1. The first stick now has 5 bits of available memory.
    // - At the 4th second, 4 bits of memory are allocated to stick 2. The second stick now has 4 bits of available memory.
    // - At the 5th second, 5 bits of memory are allocated to stick 1. The first stick now has 0 bits of available memory.
    // - At the 6th second, the program crashes. The sticks have 0 and 4 bits available respectively.
    fmt.Println(memLeak(8,11)) // [6,0,4]

    fmt.Println(memLeak(0,0)) // [1 0 0]
    fmt.Println(memLeak(1024,1024)) //[64 0 32]
    fmt.Println(memLeak(0,1024)) // [45 0 34]
    fmt.Println(memLeak(1024,0)) // [45 34 0]
    fmt.Println(memLeak(1 << 31, 1 << 31)) // [92681 88048 41708]
    fmt.Println(memLeak(1 << 31, 0)) // [65536 32768 0]
    fmt.Println(memLeak(0, 1 << 31)) // [65536 0 32768]

    fmt.Println(memLeak1(2,2)) // [3,1,0]
    fmt.Println(memLeak1(8,11)) // [6,0,4]
    fmt.Println(memLeak1(0,0)) // [1 0 0]
    fmt.Println(memLeak1(1024,1024)) //[64 0 32]
    fmt.Println(memLeak1(0,1024)) // [45 0 34]
    fmt.Println(memLeak1(1024,0)) // [45 34 0]
    fmt.Println(memLeak1(1 << 31, 1 << 31)) // [92681 88048 41708]
    fmt.Println(memLeak1(1 << 31, 0)) // [65536 32768 0]
    fmt.Println(memLeak1(0, 1 << 31)) // [65536 0 32768]
}