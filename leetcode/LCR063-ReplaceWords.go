package main

// LCR 063. 单词替换
// 在英语中，有一个叫做 词根(root) 的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。
// 例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。

// 现在，给定一个由许多词根组成的词典和一个句子，需要将句子中的所有继承词用词根替换掉。
// 如果继承词有许多可以形成它的词根，则用最短的词根替换它。

// 需要输出替换之后的句子。

// 示例 1：
// 输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// 输出："the cat was rat by the bat"

// 示例 2：
// 输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
// 输出："a a b c"

// 示例 3：
// 输入：dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
// 输出："a a a a a a a a bbb baba a"

// 示例 4：
// 输入：dictionary = ["catt","cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// 输出："the cat was rat by the bat"

// 示例 5：
// 输入：dictionary = ["ac","ab"], sentence = "it is abnormal that this solution is accepted"
// 输出："it is ab that this solution is ac"

// 提示：
//     1 <= dictionary.length <= 1000
//     1 <= dictionary[i].length <= 100
//     dictionary[i] 仅由小写字母组成。
//     1 <= sentence.length <= 10^6
//     sentence 仅由小写字母和空格组成。
//     sentence 中单词的总量在范围 [1, 1000] 内。
//     sentence 中每个单词的长度在范围 [1, 1000] 内。
//     sentence 中单词之间由一个空格隔开。
//     sentence 没有前导或尾随空格。

import "fmt"
import "strings"

type TrieNode struct {
    childrens [26]*TrieNode
    isWordEnd bool
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{
        root: &TrieNode{},
    }
}

func (this *Trie) insert(word string) {
    n, current := len(word), this.root
    for i := 0; i < n; i++ {
        index := word[i] - 'a'
        if current.childrens[index] == nil {
            current.childrens[index] = &TrieNode{}
        }
        current = current.childrens[index]
    }
    current.isWordEnd = true
}

func (this *Trie) find(word string) (bool, int) {
    n, current := len(word), this.root
    for i := 0; i < n; i++ {
        index := word[i] - 'a'
        if current.childrens[index] == nil {
            return false, -1
        }
        current = current.childrens[index]
        if current.isWordEnd {
            return true, i
        }
    }
    return false, -1
}

func replaceWords(dictionary []string, sentence string) string {
    t := NewTrie()
    for _, word := range dictionary {
        t.insert(word)
    }
    words := strings.Split(sentence, " ")
    for i, word := range words {
        found, index := t.find(word)
        if found {
            words[i] = word[:index+1]
        }
    }
    return strings.Join(words, " ")
}

func main() { 
    // 示例 1：
    // 输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
    // 输出："the cat was rat by the bat"
    fmt.Println(replaceWords([]string{"cat","bat","rat"},"the cattle was rattled by the battery")) // "the cat was rat by the bat" 
    // 示例 2：
    // 输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
    // 输出："a a b c"
    fmt.Println(replaceWords([]string{"a","b","c"}, "aadsfasf absbs bbab cadsfafs")) // "a a b c"
    // 示例 3：
    // 输入：dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
    // 输出："a a a a a a a a bbb baba a"
    fmt.Println(replaceWords([]string{"a", "aa", "aaa", "aaaa"}, "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa")) // "a a a a a a a a bbb baba a"
    // 示例 4：
    // 输入：dictionary = ["catt","cat","bat","rat"], sentence = "the cattle was rattled by the battery"
    // 输出："the cat was rat by the bat"
    fmt.Println(replaceWords([]string{"catt","cat","bat","rat"}, "the cattle was rattled by the battery")) // "the cat was rat by the bat"
    // 示例 5：
    // 输入：dictionary = ["ac","ab"], sentence = "it is abnormal that this solution is accepted"
    // 输出："it is ab that this solution is ac"
    fmt.Println(replaceWords([]string{"ac","ab"}, "it is abnormal that this solution is accepted")) // "it is ab that this solution is ac"
}