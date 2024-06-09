package main

// 466. Count The Repetitions
// We define str = [s, n] as the string str which consists of the string s concatenated n times.
//     For example, str == ["abc", 3] =="abcabcabc".

// We define that string s1 can be obtained from string s2 if we can remove some characters from s2 such that it becomes s1.
//     For example, s1 = "abc" can be obtained from s2 = "abdbec" based on our definition by removing the bolded underlined characters.

// You are given two strings s1 and s2 and two integers n1 and n2. 
// You have the two strings str1 = [s1, n1] and str2 = [s2, n2].
// Return the maximum integer m such that str = [str2, m] can be obtained from str1.

// Example 1:
// Input: s1 = "acb", n1 = 4, s2 = "ab", n2 = 2
// Output: 2

// Example 2:
// Input: s1 = "acb", n1 = 1, s2 = "acb", n2 = 1
// Output: 1
 
// Constraints:
//     1 <= s1.length, s2.length <= 100
//     s1 and s2 consist of lowercase English letters.
//     1 <= n1, n2 <= 10^6

import "fmt"

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
    n := len(s2)
    cnt := make([]int, n)
    for i := 0; i < n; i++ { // 如果重新给一个s1 并且s2是从第i位开始匹配 那么s2可以走多少位（走完了就从头开始走
        p1, p2 := 0, i
        for p1 < len(s1) {
            if s1[p1] == s2[p2 % n] {
                p2++
            }
            p1++
        }
        cnt[i] = p2 - i // 统计如果是从s2的第i位开始走 给一个新的s1 s2能走多少位
    }
    index := 0
    for i := 0; i <n1; i++ { // 直接模拟不断给s1 然后看s2能新走多少位
        index += cnt[index % n]
    }
    return index / n / n2
}

func getMaxRepetitions1(s1 string, n1 int, s2 string, n2 int) int {
    len1, len2 := len(s1), len(s2)
    if len1 == 0 || len2 == 0 || len1*n1 < len2 * n2 {
        return 0
    }
    // idx1 、cnt1 表示 s1 的下标和 s1 的遍历了多少次
    // idx2 、cnt2 表示 s2 的下标和 s2 的遍历了多少次
    // 用 map 记录每一个 idx2 对应的 cnt1 和 cnt2 ，如果之前出现过，则表示出现循环了
    // 那么整体情况是：
    // - 前 m[idx2][0] 个 s1 包含 m[idx2][1] 个 s2
    // - 以后每 cnt1 - m[idx2][0] 个 s2 包含 cnt2 - m[idx2][1] 个 s2
    // - idx2 这时候如果还没有到结尾，需要继续走完
    idx1, cnt1, idx2, cnt2 := 0, 0, 0, 0
    m := make(map[int][]int)
    loopS1 := func() {
        for idx1 = 0; idx1 < len1; idx1++ { // 遍历 s1
            if s1[idx1] == s2[idx2] {
                idx2++
                if idx2 == len2 {
                    idx2 = 0
                    cnt2++ // s2 遍历完一次了
                }
            }
        }
        cnt1++ // s1 遍历完一次了
    }
    for {
        loopS1()
        if cnt1 == n1 { // 已经结束了，还未出现循环
            return cnt2 / n2
        }
        if v, ok := m[idx2]; ok { // 出现循环了
            // - 前 m[idx2][0] 个 s1 包含 m[idx2][1] 个 s2
            // n1 = n1 - cnt1
            // n2 = n2 - cnt2
            // - 以后每 cnt1 - m[idx2][0] 个 s2 包含 cnt2 - m[idx2][1] 个 s2
            // 循环次数： (n1 - cnt1) / (cnt1 - v[0])
            times := (n1 - cnt1) / (cnt1 - v[0])
            //    前面的 + 循环的：次数 * 循环跨度
            cnt1 = cnt1 + times * (cnt1 - v[0])
            //    前面的 + 循环的：次数 * 循环跨度
            cnt2 = cnt2 + times * (cnt2 - v[1])
            // - idx2 这时候如果还没有到结尾，需要继续走完
            for cnt1 < n1 {
                loopS1()
            }
            return cnt2 / n2
        }
        m[idx2] = []int{cnt1, cnt2}
    }
}

func main() {
    // Example 1:
    // Input: s1 = "acb", n1 = 4, s2 = "ab", n2 = 2
    // Output: 2
    fmt.Println(getMaxRepetitions("acb",4,"ab",2)) // 2
    // Example 2:
    // Input: s1 = "acb", n1 = 1, s2 = "acb", n2 = 1
    // Output: 1
    fmt.Println(getMaxRepetitions("acb",1,"acb",1)) // 1

    fmt.Println(getMaxRepetitions1("acb",4,"ab",2)) // 2
    fmt.Println(getMaxRepetitions1("acb",1,"acb",1)) // 1
}