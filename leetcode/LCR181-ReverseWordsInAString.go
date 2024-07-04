package main

// LCR 181. 字符串中的单词反转
// 你在与一位习惯从右往左阅读的朋友发消息，他发出的文字顺序都与正常相反但单词内容正确，
// 为了和他顺利交流你决定写一个转换程序，把他所发的消息 message 转换为正常语序。

// 注意：输入字符串 message 中可能会存在前导空格、尾随空格或者单词间的多个空格。
// 返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

// 示例 1：
// 输入: message = "the sky is blue"
// 输出: "blue is sky the"

// 示例 2：
// 输入: message = "  hello world!  "
// 输出: "world! hello"
// 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

// 示例 3：
// 输入: message = "a good   example"
// 输出: "example good a"
// 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

// 提示：
//     1 <= message.length <= 10^4
//     message 中包含英文大小写字母、空格和数字
//     message 中至少有一个单词

import "fmt"
import "strings"

func reverseMessage(message string) string {
    // 先把字符串按照空格分隔成每个小单词 返回数组
    words := strings.Fields(message)
    reverse := func (m *[]string, i int, j int) {
        for i <= j {
            // 单词前后翻转
            (*m)[i], (*m)[j] = (*m)[j], (*m)[i]
            i++
            j--
        }
    }
    reverse(&words, 0, len(words) - 1)
    return strings.Join(words, " ") // 重新组合成字符串
}


// 双指针
func reverseMessage1(message string) string{
    words := strings.Fields(message)
    first, last := 0, len(words)-1
    for first < last {
        words[first], words[last] =  words[last], words[first]
        first++
        last--
    }
    return strings.Join(words, " ")
}

func main() {
    fmt.Println(reverseMessage("the sky is blue")) // blue is sky the
    fmt.Println(reverseMessage("  hello world  ")) // world hello
    fmt.Println(reverseMessage("a good   example")) // example good a

    fmt.Println(reverseMessage1("the sky is blue")) // blue is sky the
    fmt.Println(reverseMessage1("  hello world  ")) // world hello
    fmt.Println(reverseMessage1("a good   example")) // example good a
}