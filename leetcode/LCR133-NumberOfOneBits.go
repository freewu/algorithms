package main

// LCR 133. 位 1 的个数
// 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为 汉明重量).）。

// 提示：
//     请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
//     在 Java 中，编译器使用 二进制补码 记法来表示有符号整数。因此，在上面的 示例 3 中，输入表示有符号整数 -3。

// 示例 1：
// 输入：n = 11 (控制台输入 00000000000000000000000000001011)
// 输出：3
// 解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

// 示例 2：
// 输入：n = 128 (控制台输入 00000000000000000000000010000000)
// 输出：1
// 解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。

// 示例 3：
// 输入：n = 4294967293 (控制台输入 11111111111111111111111111111101，部分语言中 n = -3）
// 输出：31
// 解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。

// 提示：
//     输入必须是长度为 32 的 二进制串 。

import "fmt"
import "math/bits"

func hammingWeight(num uint32) int {
    res := 0
    for i := 0; i < 32; i++ {
        // fmt.Println((n & (1 << i)) >> i )
        if (int(num) & (1 << i)) >> i == 1 { // 取第 i 位判断中否为 1
            res++
        }
    }
    return res
}

// lib
func hammingWeight1(num uint32) int {
    return bits.OnesCount32(num)
}

func hammingWeight2(num uint32) int {
    res := 0
    for num != 0 {
        if num & 1 == 1 {
            res++
        }
        num >>= 1
    }
    return res
}

func hammingWeight3(num uint32) int {
    res := 0
    for num != 0 {
        res += int(num) % 2
        num = num >> 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 11
    // Output: 3
    // Explanation:
    // The input binary string 1011 has a total of three set bits.
    fmt.Println(hammingWeight(11)) // 3 (1011)
    // Example 2:
    // Input: n = 128
    // Output: 1
    // Explanation:
    // The input binary string 10000000 has a total of one set bit.
    fmt.Println(hammingWeight(128)) // 1 (10000000)
    // Example 3:
    // Input: n = 2147483645
    // Output: 30
    // Explanation:
    // The input binary string 1111111111111111111111111111101 has a total of thirty set bits.
    fmt.Println(hammingWeight(2147483645)) // 30 (1111111111111111111111111111101)

    fmt.Println(hammingWeight1(11)) // 3 (1011)
    fmt.Println(hammingWeight1(128)) // 1 (10000000)
    fmt.Println(hammingWeight1(2147483645)) // 30 (1111111111111111111111111111101)

    fmt.Println(hammingWeight2(11)) // 3 (1011)
    fmt.Println(hammingWeight2(128)) // 1 (10000000)
    fmt.Println(hammingWeight2(2147483645)) // 30 (1111111111111111111111111111101)

    fmt.Println(hammingWeight3(11)) // 3 (1011)
    fmt.Println(hammingWeight3(128)) // 1 (10000000)
    fmt.Println(hammingWeight3(2147483645)) // 30 (1111111111111111111111111111101)
}