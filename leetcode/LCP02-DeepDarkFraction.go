package main

// LCP 02. 分式化简
// 有一个同学在学习分式。他需要将一个连分数化成最简分数，你能帮助他吗？
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/09/09/fraction_example_1.jpg" / >
// 连分数是形如上图的分式。在本题中，所有系数都是大于等于0的整数。
// 输入的cont代表连分数的系数（cont[0]代表上图的a0，以此类推）。
// 返回一个长度为2的数组[n, m]，使得连分数的值等于n / m，且n, m最大公约数为1。

// 示例 1：
// 输入：cont = [3, 2, 0, 2]
// 输出：[13, 4]
// 解释：原连分数等价于3 + (1 / (2 + (1 / (0 + 1 / 2))))。注意[26, 8], [-13, -4]都不是正确答案。

// 示例 2：
// 输入：cont = [0, 0, 3]
// 输出：[3, 1]
// 解释：如果答案是整数，令分母为1即可。
 
// 限制：
//         cont[i] >= 0
//         1 <= cont的长度 <= 10
//         cont最后一个元素不等于0
//         答案的n, m的取值都能被32位int整型存下（即不超过2 ^ 31 - 1）。

import "fmt"

func fraction(cont []int) []int {
    res := make([]int,2)
    n := len(cont)
    // 将最后两个作为第一个，逆推回去
    res[0], res[1] = cont[n-1],1
    for i := n-2; i >= 0; i-- {
        res[1], res[0] = res[0], cont[i] * res[0] + res[1]
    }
    return res
}

// def fraction(self, cont: List[int]) -> List[int]:
//     n,m = 0, 1
//     for a in cont[::-1]:
//         n,m = m, (m * a + n)
//     return [m, n]

func main() {
    // 示例 1：
    // 输入：cont = [3, 2, 0, 2]
    // 输出：[13, 4]
    // 解释：原连分数等价于3 + (1 / (2 + (1 / (0 + 1 / 2))))。注意[26, 8], [-13, -4]都不是正确答案。
    fmt.Println(fraction([]int{ 3, 2, 0, 2 })) // [13, 4]
    // 示例 2：
    // 输入：cont = [0, 0, 3]
    // 输出：[3, 1]
    // 解释：如果答案是整数，令分母为1即可。
    fmt.Println(fraction([]int{ 0, 0, 3 })) // [3, 1]
}