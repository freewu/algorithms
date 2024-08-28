package main

// 1258. Synonymous Sentences
// You are given a list of equivalent string pairs synonyms where synonyms[i] = [si, ti] indicates that si and ti are equivalent strings. 
// You are also given a sentence text.

// Return all possible synonymous sentences sorted lexicographically.

// Example 1:
// Input: synonyms = [["happy","joy"],["sad","sorrow"],["joy","cheerful"]], text = "I am happy today but was sad yesterday"
// Output: ["I am cheerful today but was sad yesterday","I am cheerful today but was sorrow yesterday","I am happy today but was sad yesterday","I am happy today but was sorrow yesterday","I am joy today but was sad yesterday","I am joy today but was sorrow yesterday"]

// Example 2:
// Input: synonyms = [["happy","joy"],["cheerful","glad"]], text = "I am happy today but was sad yesterday"
// Output: ["I am happy today but was sad yesterday","I am joy today but was sad yesterday"]

// Constraints:
//     0 <= synonyms.length <= 10
//     synonyms[i].length == 2
//     1 <= si.length, ti.length <= 10
//     si != ti
//     text consists of at most 10 words.
//     All the pairs of synonyms are unique.
//     The words of text are separated by single spaces.

import "fmt"
import "strings"
import "sort"

// // 并查集
// func generateSentences(synonyms [][]string, text string) []string {
//     mp, index, sz, id := map[string]int{}, 0, make([]int, 0), make([]int, 0)
//     for _, s := range synonyms {
//         if _, ok := mp[s[0]]; !ok {
//             mp[s[0]] = index
//             sz = append(sz, 1)
//             id = append(id, index)
//             index++
//         }
//         if _, ok := mp[s[1]]; !ok {
//             mp[s[1]] = index
//             sz = append(sz, 1)
//             id = append(id, index)
//             index++
//         }
//     }
//     find := func(x int, id []int) int {
//         for x != id[x] {
//             id[x] = id[id[x]]
//             x = id[x]
//         }
//         return id[x]
//     }
//     union := func(i, j int, sz, id []int) {
//         x, y := find(i, id), find(j, id)
//         if x == y { return }
//         if sz[x] < sz[y] {
//             id[x] = y
//             sz[y] += sz[x]
//         } else {
//             id[y] = x
//             sz[x] += sz[y]
//         }
//     }
//     for _, s := range synonyms {
//         union(mp[s[0]], mp[s[1]], sz, id)
//     }
//     same := map[int][]string{}
//     for k, v := range mp {
//         same[id[v]] = append(same[id[v]], k)
//     }
//     var ctrc func(texts []string, m map[string]int, id []int, same map[int][]string) []string 
//     ctrc = func(texts []string, m map[string]int, id []int, same map[int][]string) []string {
//         if len(texts) == 0 {
//             return []string{}
//         }
//         add, sub := []string{}, ctrc(texts[1:], m, id, same)
//         if v, ok := m[texts[0]]; !ok {
//             add = []string{ texts[0] }
//         } else {
//             add = same[find(v, id)]
//         }
//         if len(sub) == 0 {
//             return add
//         }
//         res := []string{}
//         for _, a := range add {
//             for _, s := range sub {
//                 res = append(res, a + " " + s)
//             }
//         }
//         return res
//     }
//     res := ctrc(strings.Split(text, " "), mp, id, same)
//     sort.Strings(res)
//     return res
// }

type Union struct {
    parent *Union
    val    string
}

func (u *Union) getParent() *Union {
    if u.parent == nil || u.parent == u {
        return u
    }
    p := u.parent.getParent()
    u.parent = p
    return p
}

func generateSentences(synonyms [][]string, text string) []string {
    res, texts,syn := []string{}, strings.Split(text, " "), map[string]*Union{}
    for _, synonym := range synonyms {
        a, b := synonym[0], synonym[1]
        if syn[a] == nil {
            syn[a] = &Union{val: a}
        }
        if syn[b] == nil {
            syn[b] = &Union{val: b}
        }
        syn[a].getParent().parent = syn[b].getParent()
    }
    ss := map[string][]string{}
    for s, u := range syn {
        ss[u.getParent().val] = append(ss[u.getParent().val], s)
    }
    for s := range ss {
        sort.Strings(ss[s])
    }
    for i := len(texts) - 1; i >= 0; i-- {
        str := texts[i]
        resource := []string{str}
        if syn[str] != nil {
            resource = ss[syn[str].getParent().val]
        }
        if len(res) == 0 {
            res = resource
            continue
        }
        n := len(res)
        for _, s := range resource {
            for j := 0; j < n; j++ {
                res = append(res, s + " " + res[j])
            }
        }
        res = res[n:]
    }
    return res
}

func main() {
    // Example 1:
    // Input: synonyms = [["happy","joy"],["sad","sorrow"],["joy","cheerful"]], text = "I am happy today but was sad yesterday"
    // Output: ["I am cheerful today but was sad yesterday","I am cheerful today but was sorrow yesterday","I am happy today but was sad yesterday","I am happy today but was sorrow yesterday","I am joy today but was sad yesterday","I am joy today but was sorrow yesterday"]
    synonyms1 := [][]string{
        {"happy","joy"},
        {"sad","sorrow"},
        {"joy","cheerful"},
    }
    fmt.Println(generateSentences(synonyms1, "I am happy today but was sad yesterday")) // ["I am cheerful today but was sad yesterday","I am cheerful today but was sorrow yesterday","I am happy today but was sad yesterday","I am happy today but was sorrow yesterday","I am joy today but was sad yesterday","I am joy today but was sorrow yesterday"]
    // Example 2:
    // Input: synonyms = [["happy","joy"],["cheerful","glad"]], text = "I am happy today but was sad yesterday"
    // Output: ["I am happy today but was sad yesterday","I am joy today but was sad yesterday"]
    synonyms2 := [][]string{
        {"happy","joy"},
        {"cheerful","glad"},
    }
    fmt.Println(generateSentences(synonyms2, "I am happy today but was sad yesterday")) // ["I am happy today but was sad yesterday","I am joy today but was sad yesterday"]

    synonyms3 := [][]string{
        {"a","b"},
        {"b","c"},
        {"d","e"},
        {"c","d"},
    }
    fmt.Println(generateSentences(synonyms3, "a b")) // ["a a","a b","a c","a d","a e","b a","b b","b c","b d","b e","c a","c b","c c","c d","c e","d a","d b","d c","d d","d e","e a","e b","e c","e d","e e"]
}