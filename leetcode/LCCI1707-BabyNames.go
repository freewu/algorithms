package main

// 面试题 17.07. Baby Names LCCI
// Each year, the government releases a list of the 10000 most common baby names and their frequencies (the number of babies with that name). 
// The only problem with this is that some names have multiple spellings. 
// For example,"John" and ''Jon" are essentially the same name but would be listed separately in the list. 
// Given two lists, one of names/frequencies and the other of pairs of equivalent names, write an algorithm to print a new list of the true frequency of each name. 
// Note that if John and Jon are synonyms, and Jon and Johnny are synonyms, then John and Johnny are synonyms. 
// (It is both transitive and symmetric.) In the final list, choose the name that are lexicographically smallest as the "real" name.

// Example:
// Input: names = ["John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"], synonyms = ["(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"]
// Output: ["John(27)","Chris(36)"]

// Note:
//     names.length <= 100000

import "fmt"
import "strconv"
import "strings"
import "cmp"

// // 解答错误 35 / 36 
// func trulyMostPopular(names []string, synonyms []string) []string {
//     freq, set := make(map[string]int), make(map[string]string)
//     find := func(name string) string {
//         if _, exit := set[name]; !exit {
//             return ""
//         }
//         for set[name] != name {
//             name = set[name]
//         }
//         return name
//     }
//     union := func(name1, name2 string) {
//         set1, set2 := find(name1), find(name2)
//         if set1 != "" && set2 != "" && set1 != set2 {
//             if set1 < set2 {
//                 set[set2] = set1
//                 freq[set1] += freq[set2]
//                 freq[set2] = 0
//             } else {
//                 set[set1] = set2
//                 freq[set2] += freq[set1]
//                 freq[set1] = 0
//             }
//         }
//     }
//     for _, v := range names {
//         end := 0
//         for v[end] != '(' {
//             end++
//         }
//         name := v[:end]
//         set[name] = v[:end]
//         freq[name], _ = strconv.Atoi(v[end + 1: len(v) - 1])
//     }
//     for _, v := range synonyms {
//         end := 0
//         for v[end] != ',' {
//             end++
//         }
//         union(v[1: end], v[end + 1: len(v) - 1])
//     }
//     res := []string{}
//     for name := range freq {
//         if freq[name] != 0 {
//             res = append(res, name + "(" + strconv.Itoa(freq[name]) + ")")
//         }
//     }
//     return res
// }

func trulyMostPopular(names []string, synonyms []string) []string {
    p := make(map[string]string)
    var find func(x string) string
    find = func(x string) string {
        v, ok := p[x]
        if !ok {
            p[x] = x
        }
        if v != x {
            p[x] = find(p[x])
        }
        return p[x]
    }
    union := func(x, y string) {
        px, py := find(x), find(y)
        if px == py { return }
        // 字典序最小的名字作为父节点
        if px < py {
            p[py] = px
        } else {
            p[px] = py
        }
    }
    for _, s := range synonyms {
        dot := strings.Index(s, ",")
        union(s[1:dot], s[dot+1:len(s) - 1])
    }
    count := make(map[string]int)
    for _, name := range names {
        left := strings.Index(name, "(")
        pa := find(name[:left])
        c, _ := strconv.Atoi(name[left+1 : len(name)-1])
        count[pa] += c
    }
    res := make([]string, 0, len(count))
    for k, v := range count {
        res = append(res, fmt.Sprintf("%s(%d)", k, v))
    }
    return res
}

func trulyMostPopular1(names []string, synonyms []string) (ans []string) {
    type Item struct {
        name string
        count int
    }
    n := len(names)
    arr, seq := make([]Item, n), make(map[string]int)
    for i, v := range names {
        name, str, _ := strings.Cut(v, "(")
        count, _ := strconv.Atoi(strings.TrimRight(str, ")"))
        arr[i] = Item{ name: name, count: count }
        seq[name] = i
    }
    p := make([]int, n)
    for i := range p {
        p[i] = i
    }
    var find func(i int) int 
    find = func(i int) int {
        if i == p[i] { return i }
        return find(p[i])
    }
    union := func(a, b int) {
        pa, pb := find(a), find(b)
        if pa == pb { return }
        namea, nameb := arr[pa].name, arr[pb].name
        if cmp.Compare(namea, nameb) < 0 {
            p[pb] = pa
            return
        }
        p[pa] = pb
        
    }
    for _, v := range synonyms {
        namea, nameb, _ := strings.Cut(v[1:len(v)-1], ",")
        a, b := seq[namea], seq[nameb]
        union(a, b)
    }
    for i, v := range arr {
        pi := find(i)
        if pi == i {
            continue
        }
        arr[pi].count += v.count
        
    }
    res := []string{}
    for i, v := range arr {
        pi := find(i)
        if pi == i {
            res = append(res, fmt.Sprintf("%s(%d)", v.name, v.count))
        }
    }
    return res
}

func main() {
    // Example:
    // Input: names = ["John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"], synonyms = ["(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"]
    // Output: ["John(27)","Chris(36)"]
    fmt.Println(trulyMostPopular([]string{"John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"}, []string{"(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"})) // ["John(27)","Chris(36)"]

    fmt.Println(trulyMostPopular([]string{"a(10)","c(13)"}, []string{"(a,b)","(c,d)","(b,c)"})) // ["a(23)"]

    fmt.Println(trulyMostPopular1([]string{"John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"}, []string{"(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"})) // ["John(27)","Chris(36)"]
    fmt.Println(trulyMostPopular1([]string{"a(10)","c(13)"}, []string{"(a,b)","(c,d)","(b,c)"})) // ["a(23)"]
}