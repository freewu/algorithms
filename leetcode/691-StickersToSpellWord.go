package main

// 691. Stickers to Spell Word
// We are given n different types of stickers. Each sticker has a lowercase English word on it.

// You would like to spell out the given string target by cutting individual letters from your collection of stickers and rearranging them. 
// You can use each sticker more than once if you want, and you have infinite quantities of each sticker.

// Return the minimum number of stickers that you need to spell out target. 
// If the task is impossible, return -1.

// Note: In all test cases, all words were chosen randomly from the 1000 most common US English words, 
// and target was chosen as a concatenation of two random words.

// Example 1:
// Input: stickers = ["with","example","science"], target = "thehat"
// Output: 3
// Explanation:
// We can use 2 "with" stickers, and 1 "example" sticker.
// After cutting and rearrange the letters of those stickers, we can form the target "thehat".
// Also, this is the minimum number of stickers necessary to form the target string.

// Example 2:
// Input: stickers = ["notice","possible"], target = "basicbasic"
// Output: -1
// Explanation:
// We cannot form the target "basicbasic" from cutting letters from the given stickers.

// Constraints:
//     n == stickers.length
//     1 <= n <= 50
//     1 <= stickers[i].length <= 10
//     1 <= target.length <= 15
//     stickers[i] and target consist of lowercase English letters.

import "fmt"

func minStickers(stickers []string, target string) int {
    stickCount, inf := map[string]map[rune]int{}, 1 << 32 - 1
    for _, v := range stickers { // 统计每个词每个字符的频次
        stickCount[v] = map[rune]int{}
        for _, c := range v {
            stickCount[v][c]++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    charInString := func(s string, r rune) bool { // 判读字符 r 是否存在 字符 s 中
        for _, c := range s {
            if c == r {
                return true
            }
        }
        return false
    }
    dp := map[string]int{}
    var dfs func(t string, stick map[rune]int) int
    dfs = func(t string, stick map[rune]int) int {
        if v, f := dp[t]; f {
            return v
        }
        res := 0
        if stick != nil {
            res = 1
        }
        // string remains after use a stick
        remainT := []rune{}
        // substract char in stick and target string
        for _, c := range t {
            // found c in stick
            if cnt, f := stick[c]; f && cnt > 0 {
                stick[c]--
            } else {
                remainT = append(remainT, c)
            }
        }
        if len(remainT) > 0 {
            usedSticks := inf
            for st, stCount := range stickCount {
                if !charInString(st, remainT[0]) {
                    continue
                }
                // copy to new Stick
                newStick := map[rune]int{}
                for key, value := range stCount {
                    newStick[key] = value
                }
                usedSticks = min(usedSticks, dfs(string(remainT), newStick))
            }
            dp[string(remainT)] = usedSticks
            res += usedSticks
        }
        return res
    }
    if res := dfs(target, nil); res != inf {
        return res
    }
    return -1
}

func minStickers1(stickers []string, target string) int {
    // 使用图的宽度优先遍历 或者 dp
    // 图的优先遍历, 生成每个字母的对应的贴纸
    // 图上的节点为target-cur剩下的部分,算作下一层, 优化!! 2个剪枝策略,无需走每个sticker
    var charM [26][]int // 含有小写字母x的贴纸有哪些, k:char-'a' v:stickerIdx
    transform := func(str string, addM bool, idx int) [26]int8 { //统计str每个字符的数量, idxId作为这个字符串的id
        var cnt [26]int8
        for _, c := range []byte(str) {
            i := c - 'a'
            cnt[i]++
            if addM && cnt[i] == 1 { //一个字符串含有重复字符,防止被添加多次
                charM[i] = append(charM[i], idx)
            }
        }
        return cnt
    }
    cntArr := make([][26]int8, len(stickers))
    for i, sticker := range stickers {
        cnt := transform(sticker, true, i)
        cntArr[i] = cnt
    }
    next := func(t, s [26]int8) ([26]int8, int) { // t remove掉s有的字符, 并且返回t还剩余的字符数量
        leftCnt := 0
        for i := 0; i < 26; i++ {
            t[i] = max(0, t[i]-s[i])
            leftCnt += int(t[i])
        }
        return t, leftCnt
    }
    visited := make(map[[26]int8]bool) // map不支持slice做key,但支持array做key,如果不能使用array,只能将图中的节点转为string存储了.那样就麻烦了
    MAXN := 1000
    queue := make([][26]int8, MAXN)
    ql, qr := 0, 0
    tg := transform(target, false, -1)
    queue[qr] = tg
    qr++
    visited[tg] = true
    level := 0
    for ql < qr {
        level++
        size := qr - ql
        for i := 0; i < size; i++ { // 剪枝1,下一层无需走所有的sticker,只需要走含有的target剩余字符的节点就好,因为这样会使得target减少,最终求得最短路径
            cur := queue[ql]
            ql++
            for j := 0; j < 26; j++ { // 注意!! 并不一定是删除存留首字符数量越多的就越好(贪心思路不对, 比如left="aabb",s1="aa" s2="ab", 贪心思路使用1个s1,但不如使用2个s2更好
                if cur[j] > 0 { // 剪枝2: 字符串有多个字符,想要全部删除,可以从前往后删除,层数是一样的, 比如target含有3个a,不管如何尝试,最优路径都要把3个a删掉,不如先尝试删掉a.只走少量路径
                    for _, sId := range charM[j] {
                        nx, leftCnt := next(cur, cntArr[sId])
                        if leftCnt == 0 {
                            return level // 因为在生成下层节点时就判断了(而不是弹出下层节点),next层为level+1,一共转换了 level+1-1(第一层为target)=level次
                        }
                        if !visited[nx] {
                            visited[nx] = true // 注意!! 对于宽度优先遍历来说,入队列即算visited了,不然同层节点可能产生多个相同的后续节点
                            queue[qr] = nx
                            qr++
                        }
                    }
                    break
                }
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: stickers = ["with","example","science"], target = "thehat"
    // Output: 3
    // Explanation:
    // We can use 2 "with" stickers, and 1 "example" sticker.
    // After cutting and rearrange the letters of those stickers, we can form the target "thehat".
    // Also, this is the minimum number of stickers necessary to form the target string.
    fmt.Println(minStickers([]string{"with","example","science"}, "thehat")) // 3
    // Example 2:
    // Input: stickers = ["notice","possible"], target = "basicbasic"
    // Output: -1
    // Explanation:
    // We cannot form the target "basicbasic" from cutting letters from the given stickers.
    fmt.Println(minStickers([]string{"notice","possible"}, "basicbasic")) // -1

    fmt.Println(minStickers1([]string{"with","example","science"}, "thehat")) // 3
    fmt.Println(minStickers1([]string{"notice","possible"}, "basicbasic")) // -1
}