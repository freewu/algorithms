package main

// LCR 114. 火星词典
// 现有一种使用英语字母的外星文语言，这门语言的字母顺序与英语顺序不同。

// 给定一个字符串列表 words ，作为这门语言的词典，words 中的字符串已经 按这门新语言的字母顺序进行了排序 。

// 请你根据该词典还原出此语言中已知的字母顺序，并 按字母递增顺序 排列。若不存在合法字母顺序，返回 "" 。
// 若存在多种可能的合法字母顺序，返回其中 任意一种 顺序即可。

// 字符串 s 字典顺序小于 字符串 t 有两种情况：
//     在第一个不同字母处，如果 s 中的字母在这门外星语言的字母顺序中位于 t 中字母之前，那么 s 的字典顺序小于 t 。
//     如果前面 min(s.length, t.length) 字母都相同，那么 s.length < t.length 时，s 的字典顺序也小于 t 。

// 示例 1：
// 输入：words = ["wrt","wrf","er","ett","rftt"]
// 输出："wertf"

// 示例 2：
// 输入：words = ["z","x"]
// 输出："zx"

// 示例 3：
// 输入：words = ["z","x","z"]
// 输出：""
// 解释：不存在合法字母顺序，因此返回 "" 。

// 提示：
//     1 <= words.length <= 100
//     1 <= words[i].length <= 100
//     words[i] 仅由小写英文字母组成

import "fmt"
import "bytes"

// 拓扑排序
func alienOrder(words []string) string {
    mm := make(map[byte][]byte)
    addEdge := func(a, b string, k int) {
        if _, ok := mm[a[k]]; !ok {
            mm[a[k]] = make([]byte, 0)
        }
        mm[a[k]] = append(mm[a[k]], b[k])
    }
    for i := range words {
        for j := i + 1; j < len(words); j++ {
            if words[i] == words[j] {
                continue
            }
            a, b := words[i], words[j]
            k := 0
            for k < len(a) && k < len(b) && a[k] == b[k] {
                k++
            }
            if k < len(a) && k < len(b) {
                addEdge(a, b, k)
            } else if k == len(b) {
                return ""
            }
        }
    }
    alpha := make(map[byte]struct{})
    for i := range words {
        for j := range words[i] {
            alpha[words[i][j]] = struct{}{}
        }
    }
    w := &bytes.Buffer{}
    for len(mm) > 0 {
        e := make(map[byte]int)
        for x, _ := range alpha {
            e[x] = 0
        }
        for _, v := range mm {
            for _, x := range v {
                e[x]++
            }
        }
        var cur byte = 0xff
        for x, cnt := range e {
            _, ok := mm[x]
            if ok && cnt == 0 {
                cur = x
                break
            }
        }
        if cur == 0xff {
            return ""
        }
        w.WriteByte(cur)
        delete(mm, cur)
        delete(alpha, cur)
    }
    for k, _ := range alpha {
        w.WriteByte(k)
    }
    return w.String()
}

// dfs
func alienOrder1(words []string) string {
    n := len(words)
    g, vis := make([][]int, 26), make([]bool, 26)
    for _, c := range words[0] {
        vis[c-'a'] = true
    }
    for i := 1; i < n; i++ {
        w1, w2 := words[i-1], words[i]
        j := 0
        for ; j < min(len(w1), len(w2)); j++ {
            vis[w2[j]-'a'] = true
            if w1[j] == w2[j] {
                continue
            }
            g[w1[j]-'a'] = append(g[w1[j]-'a'], int(w2[j]-'a'))
            break
        }
        if j == len(w2) && j < len(w1) {
            return ""
        }
        for ; j < len(w2); j++ {
            vis[w2[j]-'a'] = true
        }
    }
    cnt := 0
    for _, v := range vis {
        if v {
            cnt++
        }
    }
    res, times := []byte{}, make([]int, 26)
    var dfs func(u int) bool
    dfs = func(u int) bool {
        if times[u] == 2 {
            return true
        } else if times[u] == 1 {
            return false
        }
        times[u] = 1
        for _, v := range g[u] {
            if !dfs(v) {
                return false
            }
        }
        res = append(res, byte(u+'a'))
        times[u] = 2
        return true
    }
    for i, v := range vis {
        if v && !dfs(i) {
            return ""
        }
    }
    for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
        res[l], res[r] = res[r], res[l]
    }
    return string(res)
}

func main() {
    fmt.Println(alienOrder([]string{"wrt","wrf","er","ett","rftt"})) // wertf
    fmt.Println(alienOrder([]string{"z","x"})) // zf
    fmt.Println(alienOrder([]string{"z","x","z"})) // ""

    fmt.Println(alienOrder1([]string{"wrt","wrf","er","ett","rftt"})) // wertf
    fmt.Println(alienOrder1([]string{"z","x"})) // zf
    fmt.Println(alienOrder1([]string{"z","x","z"})) // ""
}