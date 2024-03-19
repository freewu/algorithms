package main

// 269. Alien Dictionary
// There is a new alien language that uses the English alphabet. However, the order of the letters is unknown to you.
// You are given a list of strings words from the alien language's dictionary. 
// Now it is claimed that the strings in words are sorted lexicographically by the rules of this new language.
// If this claim is incorrect, and the given arrangement of string in words cannot correspond to any order of letters, return "".
// Otherwise, return a string of the unique letters in the new alien language sorted in lexicographically increasing order by the new language's rules. 
// If there are multiple solutions, return any of them.

// Example 1:
// Input: words = ["wrt","wrf","er","ett","rftt"]
// Output: "wertf"

// Example 2:
// Input: words = ["z","x"]
// Output: "zx"
// Example 3:

// Input: words = ["z","x","z"]
// Output: ""
// Explanation: The order is invalid, so return "".
 
// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 100
//     words[i] consists of only lowercase English letters.

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
	const CHARSET int = 26
	n := len(words)
	g := make([][]int, CHARSET)
	vis := make([]bool, CHARSET)
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
	ans := []byte{}
	times := make([]int, CHARSET)
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
		ans = append(ans, byte(u+'a'))
		times[u] = 2
		return true
	}
	for i, v := range vis {
		if v && !dfs(i) {
			return ""
		}
	}
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return string(ans)
}

func main() {
    fmt.Println(alienOrder([]string{"wrt","wrf","er","ett","rftt"})) // wertf
    fmt.Println(alienOrder([]string{"z","x"})) // zf
    fmt.Println(alienOrder([]string{"z","x","z"})) // ""

    fmt.Println(alienOrder1([]string{"wrt","wrf","er","ett","rftt"})) // wertf
    fmt.Println(alienOrder1([]string{"z","x"})) // zf
    fmt.Println(alienOrder1([]string{"z","x","z"})) // ""
}