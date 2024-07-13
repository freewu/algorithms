package main

// LCR 064. 实现一个魔法字典
// 设计一个使用单词列表进行初始化的数据结构，单词列表中的单词 互不相同 。 
// 如果给出一个单词，请判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单词存在于已构建的神奇字典中。

// 实现 MagicDictionary 类：
//     MagicDictionary() 初始化对象
//     void buildDict(String[] dictionary) 使用字符串数组 dictionary 设定该数据结构，dictionary 中的字符串互不相同
//     bool search(String searchWord) 给定一个字符串 searchWord ，判定能否只将字符串中 一个 字母换成另一个字母，使得所形成的新字符串能够与字典中的任一字符串匹配。如果可以，返回 true ；否则，返回 false 。

// 示例：
// 输入
// inputs = ["MagicDictionary", "buildDict", "search", "search", "search", "search"]
// inputs = [[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
// 输出
// [null, null, false, true, false, false]
// 解释
// MagicDictionary magicDictionary = new MagicDictionary();
// magicDictionary.buildDict(["hello", "leetcode"]);
// magicDictionary.search("hello"); // 返回 False
// magicDictionary.search("hhllo"); // 将第二个 'h' 替换为 'e' 可以匹配 "hello" ，所以返回 True
// magicDictionary.search("hell"); // 返回 False
// magicDictionary.search("leetcoded"); // 返回 False
 
// 提示：
//     1 <= dictionary.length <= 100
//     1 <= dictionary[i].length <= 100
//     dictionary[i] 仅由小写英文字母组成
//     dictionary 中的所有字符串 互不相同
//     1 <= searchWord.length <= 100
//     searchWord 仅由小写英文字母组成
//     buildDict 仅在 search 之前调用一次
//     最多调用 100 次 search

import "fmt"

type MagicDictionary struct {
    children [26]*MagicDictionary
    isEnd    bool
}

func Constructor() MagicDictionary {
    return MagicDictionary{}
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
    for _, word := range dictionary {
        m.insert(word)
    }
}

func (m *MagicDictionary) insert(word string) {
    cur := m
    for _, ch := range word {
        if cur.children[ch-'a'] == nil {
            cur.children[ch-'a'] = &MagicDictionary{}
        }
        cur = cur.children[ch-'a']
    }
    cur.isEnd = true
}

func (m *MagicDictionary) Search(searchWord string) bool {
	return dfs(m, searchWord, 0, 1)
}

func dfs(r *MagicDictionary, w string, i int, limit int) bool {
    // base case
    if limit < 0 {
        return false
    }
    if i == len(w) {
        return r.isEnd && limit == 0
    }
    ch := w[i] - 'a' 
    for c, t := range r.children { // iterate current node's all children
        if t == nil {
            continue
        }
        if c == int(ch) && dfs(t, w, i+1, limit) { // c == ch, represent don't need change.
            return true
        }
        if c != int(ch) && dfs(t, w, i+1, limit-1) { // c != ch, represent consume one chance for change
            return true
        }
    }
    return false
}

type MagicDictionary1 struct {
    data []string
}

func Constructor1() MagicDictionary1 {
    return MagicDictionary1{}
}

func (this *MagicDictionary1) BuildDict(dictionary []string)  {
    this.data = dictionary
}

func (this *MagicDictionary1) Search(searchWord string) bool {
    for _, d := range this.data {
        c := 0
        if len(d) != len(searchWord) { continue }
        for i := 0; i < len(searchWord); i++ {
            if searchWord[i] != d[i] { c++  }
        }
        if c == 1 {  return true }
    }
    return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

func main() {
    // MagicDictionary magicDictionary = new MagicDictionary();
    obj := Constructor()
    fmt.Println(obj)
    // magicDictionary.buildDict(["hello", "leetcode"]);
    obj.BuildDict([]string{"hello", "leetcode"})
    fmt.Println(obj)
    // magicDictionary.search("hello"); // return False
    fmt.Println(obj.Search("hello")) // false
    // magicDictionary.search("hhllo"); // We can change the second 'h' to 'e' to match "hello" so we return True
    fmt.Println(obj.Search("hhllo")) // true
    // magicDictionary.search("hell"); // return False
    fmt.Println(obj.Search("hell")) // false
    // magicDictionary.search("leetcoded"); // return False
    fmt.Println(obj.Search("leetcoded")) // false

    obj1 := Constructor1()
    fmt.Println(obj1)
    // magicDictionary.buildDict(["hello", "leetcode"]);
    obj1.BuildDict([]string{"hello", "leetcode"})
    fmt.Println(obj1)
    // magicDictionary.search("hello"); // return False
    fmt.Println(obj1.Search("hello")) // false
    // magicDictionary.search("hhllo"); // We can change the second 'h' to 'e' to match "hello" so we return True
    fmt.Println(obj1.Search("hhllo")) // true
    // magicDictionary.search("hell"); // return False
    fmt.Println(obj1.Search("hell")) // false
    // magicDictionary.search("leetcoded"); // return False
    fmt.Println(obj1.Search("leetcoded")) // false
}