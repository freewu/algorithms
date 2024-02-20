package main

// LCR 002. 二进制求和
// 给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
// 输入为 非空 字符串且只包含数字 1 和 0。

// 示例 1:
// 输入: a = "11", b = "10"
// 输出: "101"

// 示例 2:
// 输入: a = "1010", b = "1011"
// 输出: "10101"

// 提示：
// 		每个字符串仅由字符 '0' 或 '1' 组成。
// 		1 <= a.length, b.length <= 10^4
// 		字符串如果不是 "0" ，就都不含前导零。

import "fmt"
import "strconv"
import "strings"

func addBinary(a string, b string) string {
	// 长的放前面
	if len(b) > len(a) {
		a, b = b, a
	}
	res := make([]string, len(a) + 1) // 声明一个数组来保存结果 长度是 最长的 + 1
	i, j, k, c := len(a) - 1, len(b) - 1, len(a), 0
	// 先把对齐的长度进行相加处理
	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))
		res[k] = strconv.Itoa((ai + bj + c) % 2)
		c = (ai + bj + c) / 2 // 判断是否进位
		i--
		j--
		k--
	}
	// 全合超过长度的部分
	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		res[k] = strconv.Itoa((ai + c) % 2)
		c = (ai + c) / 2
		i--
		k--
	}
	// 有进位处理
	if c > 0 {
		res[k] = strconv.Itoa(1)
	}
	return strings.Join(res, "")
}

func main() {
	fmt.Println(addBinary("11", "1"))  // "100"
	fmt.Println(addBinary("11", "11")) // "110"
	fmt.Println(addBinary("11", "10"))  // "101"
	fmt.Println(addBinary("1010", "1011"))  // "10101"
}