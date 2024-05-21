package main

// 737. Sentence Similarity II
// We can represent a sentence as an array of words, for example, the sentence "I am happy with leetcode" can be represented as arr = ["I","am",happy","with","leetcode"].
// Given two sentences sentence1 and sentence2 each represented as a string array and given an array of string pairs similarPairs where similarPairs[i] = [xi, yi] indicates that the two words xi and yi are similar.
// Return true if sentence1 and sentence2 are similar, or false if they are not similar.

// Two sentences are similar if:
//     They have the same length (i.e., the same number of words)
//     sentence1[i] and sentence2[i] are similar.

// Notice that a word is always similar to itself, also notice that the similarity relation is transitive. 
// For example, if the words a and b are similar, and the words b and c are similar, then a and c are similar.

// Example 1:
// Input: sentence1 = ["great","acting","skills"], sentence2 = ["fine","drama","talent"], similarPairs = [["great","good"],["fine","good"],["drama","acting"],["skills","talent"]]
// Output: true
// Explanation: The two sentences have the same length and each word i of sentence1 is also similar to the corresponding word in sentence2.

// Example 2:
// Input: sentence1 = ["I","love","leetcode"], sentence2 = ["I","love","onepiece"], similarPairs = [["manga","onepiece"],["platform","anime"],["leetcode","platform"],["anime","manga"]]
// Output: true
// Explanation: "leetcode" --> "platform" --> "anime" --> "manga" --> "onepiece".
// Since "leetcode is similar to "onepiece" and the first two words are the same, the two sentences are similar.

// Example 3:
// Input: sentence1 = ["I","love","leetcode"], sentence2 = ["I","love","onepiece"], similarPairs = [["manga","hunterXhunter"],["platform","anime"],["leetcode","platform"],["anime","manga"]]
// Output: false
// Explanation: "leetcode" is not similar to "onepiece".
 
// Constraints:
//     1 <= sentence1.length, sentence2.length <= 1000
//     1 <= sentence1[i].length, sentence2[i].length <= 20
//     sentence1[i] and sentence2[i] consist of lower-case and upper-case English letters.
//     0 <= similarPairs.length <= 2000
//     similarPairs[i].length == 2
//     1 <= xi.length, yi.length <= 20
//     xi and yi consist of English letters.

import "fmt"

func areSentencesSimilarTwo(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
    if len(sentence1) != len(sentence2) {
        return false
    }
    dicts := map[string]string{} // 集
    var find func(x string) string 
    find = func(x string) string { // 查
        if dicts[x] == "" || dicts[x] == x {
            return x
        } else {
            dicts[x] = find(dicts[x])  // 路径压缩
        }
        return dicts[x]
    }
    var union = func(x, y string) { // 并
        dicts[find(x)] = find(y)
    }
    for _, p := range similarPairs { // 遍历关联相似词
        union(p[0], p[1])
    }
    for i := range sentence1 { //  遍历句子单词判断是否相似
        if find(sentence1[i]) != find(sentence2[i]) {
            return false
        }
    }
    return true
}

type UnionFind struct {
    arr []int
}

func NewUnionFind(n int) *UnionFind {
    arr := make([]int, n)
    for i := range arr {
        arr[i] = i
    }
    return &UnionFind{arr: arr}
}

func (this *UnionFind) Find(u int) int {
    if this.arr[u] == u {
        return u
    }
    this.arr[u] = this.Find(this.arr[u])
    return this.arr[u]
}

func (this *UnionFind) IsSame(u int, v int) bool {
    return this.Find(u) == this.Find(v)
}

func (this *UnionFind) Join(u int, v int) {
    u, v = this.Find(u), this.Find(v)
    if u == v {
        return
    }
    this.arr[v] = u
}

func areSentencesSimilarTwo1(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
    if len(sentence1) != len(sentence2) {
        return false
    }
    wordsMap := make(map[string]int)
    for _, similar := range similarPairs {
        if _, has := wordsMap[similar[0]]; !has {
            wordsMap[similar[0]] = len(wordsMap)
        }
        if _, has := wordsMap[similar[1]]; !has {
            wordsMap[similar[1]] = len(wordsMap)
        }
    }
    uf := NewUnionFind(len(wordsMap))
    for _, similar := range similarPairs {
        uf.Join(wordsMap[similar[0]], wordsMap[similar[1]])
    }
    for i := range sentence1 {
        if sentence1[i] == sentence2[i] { continue }
        if _, has := wordsMap[sentence1[i]]; !has { return false }
        if _, has := wordsMap[sentence2[i]]; !has { return false }
        if !uf.IsSame(wordsMap[sentence1[i]], wordsMap[sentence2[i]]) { return false }
    }
    return true
}

func main() {
    // Example 1:
    // Input: sentence1 = ["great","acting","skills"], sentence2 = ["fine","drama","talent"], similarPairs = [["great","good"],["fine","good"],["drama","acting"],["skills","talent"]]
    // Output: true
    // Explanation: The two sentences have the same length and each word i of sentence1 is also similar to the corresponding word in sentence2.
    fmt.Println(areSentencesSimilarTwo([]string{"great","acting","skills"},[]string{"fine","drama","talent"},[][]string{{"great","good"},{"fine","good"},{"drama","acting"},{"skills","talent"}})) // true
    // Example 2:
    // Input: sentence1 = ["I","love","leetcode"], sentence2 = ["I","love","onepiece"], similarPairs = [["manga","onepiece"],["platform","anime"],["leetcode","platform"],["anime","manga"]]
    // Output: true
    // Explanation: "leetcode" --> "platform" --> "anime" --> "manga" --> "onepiece".
    // Since "leetcode is similar to "onepiece" and the first two words are the same, the two sentences are similar.
    fmt.Println(areSentencesSimilarTwo([]string{"I","love","leetcode"},[]string{"I","love","onepiece"},[][]string{{"manga","onepiece"},{"platform","anime"},{"leetcode","platform"},{"anime","manga"}})) // true
    // Input: sentence1 = ["I","love","leetcode"], sentence2 = ["I","love","onepiece"], similarPairs = [["manga","hunterXhunter"],["platform","anime"],["leetcode","platform"],["anime","manga"]]
    // Output: false
    // Explanation: "leetcode" is not similar to "onepiece".
    fmt.Println(areSentencesSimilarTwo([]string{"I","love","leetcode"},[]string{"I","love","onepiece"},[][]string{{"manga","hunterXhunter"},{"platform","anime"},{"leetcode","platform"},{"anime","manga"}})) // false

    fmt.Println(areSentencesSimilarTwo1([]string{"great","acting","skills"},[]string{"fine","drama","talent"},[][]string{{"great","good"},{"fine","good"},{"drama","acting"},{"skills","talent"}})) // true
    fmt.Println(areSentencesSimilarTwo1([]string{"I","love","leetcode"},[]string{"I","love","onepiece"},[][]string{{"manga","onepiece"},{"platform","anime"},{"leetcode","platform"},{"anime","manga"}})) // true
    fmt.Println(areSentencesSimilarTwo1([]string{"I","love","leetcode"},[]string{"I","love","onepiece"},[][]string{{"manga","hunterXhunter"},{"platform","anime"},{"leetcode","platform"},{"anime","manga"}})) // false
}