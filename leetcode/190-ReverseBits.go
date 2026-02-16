package main

// 190. Reverse Bits
// Reverse bits of a given 32 bits unsigned integer.
// Note:
//     Note that in some languages, such as Java, there is no unsigned integer type. 
//     In this case, both input and output will be given as a signed integer type. 
//     They should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
    
//     In Java, the compiler represents the signed integers using 2's complement notation. 
//     Therefore, in Example 2 above, the input represents the signed integer -3 and the output represents the signed integer -1073741825.
 
// Example 1:
// Input: n = 00000010100101000001111010011100
// Output:    964176192 (00111001011110000010100101000000)
// Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, 
// so return 964176192 which its binary representation is 00111001011110000010100101000000.

// Example 2:
// Input: n = 11111111111111111111111111111101
// Output:   3221225471 (10111111111111111111111111111111)
// Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, 
// so return 3221225471 which its binary representation is 10111111111111111111111111111111.

// Constraints:
//     The input must be a binary string of length 32

// Follow up: If this function is called many times, how would you optimize it?

import "fmt"
import "strconv"

func reverseBits(num int) int {
    res := 0
    for i := 0; i < 32; i++ {
        // 把 num 往右移动，不断的消灭右边最低位的 1，将这个 1 给 res，
        // res 不断的左移即可实现反转二进制位的目的
        res = res << 1 | num & 1
        num >>= 1
    }
    return res
}

func reverseBits1(num int) int {
    const (
        m1 = 0x55555555 // 01010101010101010101010101010101
        m2 = 0x33333333 // 00110011001100110011001100110011
        m4 = 0x0f0f0f0f // 00001111000011110000111100001111
        m8 = 0x00ff00ff // 00000000111111110000000011111111
    )
    num = num >> 1 &m1 | num & m1 << 1
    num = num >> 2 &m2 | num & m2 << 2
    num = num >> 4 &m4 | num & m4 << 4
    num = num >> 8 &m8 | num & m8 << 8
    return num >> 16 | num << 16
}

func main() {
    // Example 1:
    // Input: n = 00000010100101000001111010011100
    // Output:    964176192 (00111001011110000010100101000000)
    // Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, 
    // so return 964176192 which its binary representation is 00111001011110000010100101000000.
    decimal, _ := strconv.ParseInt("00000010100101000001111010011100", 2, 64)
    fmt.Printf("%v\n",decimal)
    // 00000010100101000001111010011100
    fmt.Println(reverseBits(int(43261596))) // 964176192 (00111001011110000010100101000000)
    // Example 2:
    // Input: n = 11111111111111111111111111111101
    // Output:   3221225471 (10111111111111111111111111111111)
    // Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, 
    // so return 3221225471 which its binary representation is 10111111111111111111111111111111.
    decimal, _ = strconv.ParseInt("11111111111111111111111111111101", 2, 64)
    fmt.Printf("%v\n",decimal)
    // 11111111111111111111111111111101
    fmt.Println(reverseBits(int(4294967293))) // 3221225471 (10111111111111111111111111111111)

    fmt.Println(reverseBits(int(1))) // 3221225471 (10111111111111111111111111111111)
    fmt.Println(reverseBits(int(1024))) // 2097152
    fmt.Println(reverseBits(int(1 << 31))) // 1

    fmt.Println(reverseBits1(int(43261596))) // 964176192 (00111001011110000010100101000000)
    fmt.Println(reverseBits1(int(4294967293))) // 3221225471 (10111111111111111111111111111111)
    fmt.Println(reverseBits1(int(1))) // 2147483648
    fmt.Println(reverseBits1(int(1024))) // 2097152
    fmt.Println(reverseBits1(int(1 << 31))) // 4294967297
}