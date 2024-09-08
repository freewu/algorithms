package main

// 3023. Find Pattern in Infinite Stream I
// You are given a binary array pattern and an object stream of class InfiniteStream representing a 0-indexed infinite stream of bits.

// The class InfiniteStream contains the following function:
//     int next(): Reads a single bit (which is either 0 or 1) from the stream and returns it.

// Return the first starting index where the pattern matches the bits read from the stream. 
// For example, if the pattern is [1, 0], the first match is the highlighted part in the stream [0, 1, 0, 1, ...].

// Example 1:
// Input: stream = [1,1,1,0,1,1,1,...], pattern = [0,1]
// Output: 3
// Explanation: The first occurrence of the pattern [0,1] is highlighted in the stream [1,1,1,0,1,...], which starts at index 3.

// Example 2:
// Input: stream = [0,0,0,0,...], pattern = [0]
// Output: 0
// Explanation: The first occurrence of the pattern [0] is highlighted in the stream [0,...], which starts at index 0.

// Example 3:
// Input: stream = [1,0,1,1,0,1,1,0,1,...], pattern = [1,1,0,1]
// Output: 2
// Explanation: The first occurrence of the pattern [1,1,0,1] is highlighted in the stream [1,0,1,1,0,1,...], which starts at index 2.

// Constraints:
//     1 <= pattern.length <= 100
//     pattern consists only of 0 and 1.
//     stream consists only of 0 and 1.
//     The input is generated such that the pattern's start index exists in the first 105 bits of the stream.

import "fmt"
import "strconv"

type InfiniteStream struct {
    data []int
    index int
}

func (this *InfiniteStream) Next() int {
    this.index++
    if this.index == len(this.data) {
        this.index = 0
    } 
    return this.data[this.index]
}

func Constructor(data []int) InfiniteStream {
    return InfiniteStream{ data, 0 }
}

/**
 * Definition for an infinite stream.
 * type InfiniteStream interface {
 *     Next() int
 * }
 */
func findPattern(stream InfiniteStream, pattern []int) int {
    pStr, nowStr, index := "", "", 0
    for i := 0; i < len(pattern); i++ {
        pStr += strconv.Itoa(pattern[i])
    }
    for {
        s := strconv.Itoa(stream.Next())
        nowStr += s
        if len(nowStr) == len(pStr) {
            if pStr == nowStr {
                return index - len(pattern) + 1
            }
            nowStr = nowStr[1:]
        }
        index++
    }
    return -1
}

func main() {
    // Example 1:
    // Input: stream = [1,1,1,0,1,1,1,...], pattern = [0,1]
    // Output: 3
    // Explanation: The first occurrence of the pattern [0,1] is highlighted in the stream [1,1,1,0,1,...], which starts at index 3.
    stream1 := Constructor([]int{1,1,1,0,1,1,1,0})
    fmt.Println(findPattern(stream1, []int{0,1})) // 3
    // Example 2:
    // Input: stream = [0,0,0,0,...], pattern = [0]
    // Output: 0
    // Explanation: The first occurrence of the pattern [0] is highlighted in the stream [0,...], which starts at index 0.
    stream2 := Constructor([]int{0,0,0,0})
    fmt.Println(findPattern(stream2, []int{0})) // 0
    // Example 3:
    // Input: stream = [1,0,1,1,0,1,1,0,1,...], pattern = [1,1,0,1]
    // Output: 2
    // Explanation: The first occurrence of the pattern [1,1,0,1] is highlighted in the stream [1,0,1,1,0,1,...], which starts at index 2.
    stream3 := Constructor([]int{1,0,1,1,0,1,1,0,1})
    fmt.Println(findPattern(stream3, []int{1,1,0,1})) // 2
}