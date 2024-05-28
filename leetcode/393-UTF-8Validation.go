package main

// 393. UTF-8 Validation
// Given an integer array data representing the data, 
// return whether it is a valid UTF-8 encoding (i.e. it translates to a sequence of valid UTF-8 encoded characters).

// A character in UTF8 can be from 1 to 4 bytes long, subjected to the following rules:
//     For a 1-byte character, the first bit is a 0, followed by its Unicode code.
//     For an n-bytes character, the first n bits are all one's, the n + 1 bit is 0, followed by n - 1 bytes with the most significant 2 bits being 10.

// This is how the UTF-8 encoding would work:

//      Number of Bytes   |        UTF-8 Octet Sequence
//                        |              (binary)
//    --------------------+-----------------------------------------
//             1          |   0xxxxxxx
//             2          |   110xxxxx 10xxxxxx
//             3          |   1110xxxx 10xxxxxx 10xxxxxx
//             4          |   11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
// x denotes a bit in the binary form of a byte that may be either 0 or 1.

// Note: The input is an array of integers. Only the least significant 8 bits of each integer is used to store the data. This means each integer represents only 1 byte of data.

// Example 1:
// Input: data = [197,130,1]
// Output: true
// Explanation: data represents the octet sequence: 11000101 10000010 00000001.
// It is a valid utf-8 encoding for a 2-bytes character followed by a 1-byte character.

// Example 2:
// Input: data = [235,140,4]
// Output: false
// Explanation: data represented the octet sequence: 11101011 10001100 00000100.
// The first 3 bits are all one's and the 4th bit is 0 means it is a 3-bytes character.
// The next byte is a continuation byte which starts with 10 and that's correct.
// But the second continuation byte does not start with 10, so it is invalid.

// Constraints:
//     1 <= data.length <= 2 * 10^4
//     0 <= data[i] <= 255

import "fmt"
import "errors"

func validUtf8(data []int) bool {
    count := 0
    for _, num := range data {
        if count == 0 {
            if num >> 3 == 0b11110 {
                count = 3
            } else if num >> 4 == 0b1110 {
                count = 2
            } else if num >> 5 == 0b110 {
                count = 1
            } else if num >> 7 != 0b0 { // 对于 1 字节 的字符，字节的第一位设为 0 ，后面 7 位为这个符号的 unicode 码
                return false
            }
        } else {
            if num >> 6 != 0b10 {
                return false
            }
            count--
        }
    }
    return count == 0
}

func validUtf81(data []int) bool {
    head := [4][3]int{
        {0b10000000, 0b00000000, 1},
        {0b11100000, 0b11000000, 2},
        {0b11110000, 0b11100000, 3},
        {0b11111000, 0b11110000, 4},
    }
    getLength := func(data int) int {
        for _, h := range head {
            if data & h[0] == h[1] { return h[2] }
        }
        return 0
    }
    for i := 0; i < len(data); {
        length := getLength(data[i])
        if length == 0 || i+length > len(data) {
            return false
        }
        for _, ch := range data[i+1 : i+length] {
            if ch & 0b11000000 != 0b10000000 {
                return false
            }
        }
        i += length
    }
    return true
}

func validUtf82(data []int) bool {
    howManyBytes := func (n int) (int, error) {
        switch {
            case n & 0b10000000 == 0b00000000: { return 1, nil } 
            case n & 0b11100000 == 0b11000000: { return 2, nil }
            case n & 0b11110000 == 0b11100000: { return 3, nil }
            case n & 0b11111000 == 0b11110000: { return 4, nil }
            default: { return 0, errors.New("invalid leading header") }
        }
    }
    subsequentBytesHeader := func(n int) error {
        if n & 0b11000000 == 0b10000000 { return nil }
        return errors.New("invalid subsequent header")
    }
    for i := 0; i < len(data); {
        next, err := howManyBytes(data[i])
        if err != nil {
            return false
        }
        for y := 1; y < next; y++ {
            if i+y >= len(data) { return false }
            if err := subsequentBytesHeader(data[i+y]); err != nil { return false }
        }
        i += next
    }
    return true
}

func main() {
    // Example 1:
    // Input: data = [197,130,1]
    // Output: true
    // Explanation: data represents the octet sequence: 11000101 10000010 00000001.
    // It is a valid utf-8 encoding for a 2-bytes character followed by a 1-byte character.
    fmt.Println(validUtf8([]int{197,130,1})) // true
    // Example 2:
    // Input: data = [235,140,4]
    // Output: false
    // Explanation: data represented the octet sequence: 11101011 10001100 00000100.
    // The first 3 bits are all one's and the 4th bit is 0 means it is a 3-bytes character.
    // The next byte is a continuation byte which starts with 10 and that's correct.
    // But the second continuation byte does not start with 10, so it is invalid.
    fmt.Println(validUtf8([]int{235,140,4})) // false

    fmt.Println(validUtf81([]int{197,130,1})) // true
    fmt.Println(validUtf81([]int{235,140,4})) // false

    fmt.Println(validUtf82([]int{197,130,1})) // true
    fmt.Println(validUtf82([]int{235,140,4})) // false
}