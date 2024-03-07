package main

// 244. Shortest Word Distance II
// Design a data structure that will be initialized with a string array, 
// and then it should answer queries of the shortest distance between two different strings from the array.
// Implement the WordDistance class:
//         WordDistance(String[] wordsDict) initializes the object with the strings array wordsDict.
//         int shortest(String word1, String word2) returns the shortest distance between word1 and word2 in the array wordsDict.
 
// Example 1:
// Input
// ["WordDistance", "shortest", "shortest"]
// [[["practice", "makes", "perfect", "coding", "makes"]], ["coding", "practice"], ["makes", "coding"]]
// Output
// [null, 3, 1]
// Explanation
// WordDistance wordDistance = new WordDistance(["practice", "makes", "perfect", "coding", "makes"]);
// wordDistance.shortest("coding", "practice"); // return 3
// wordDistance.shortest("makes", "coding");    // return 1

// Constraints:
//         1 <= wordsDict.length <= 3 * 104
//         1 <= wordsDict[i].length <= 10
//         wordsDict[i] consists of lowercase English letters.
//         word1 and word2 are in wordsDict.
//         word1 != word2
//         At most 5000 calls will be made to shortest.

import "fmt"

type WordDistance struct {
    data []string
}

func Constructor(wordsDict []string) WordDistance {
    return WordDistance{wordsDict}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
    //fmt.Println("wordsDict: ", wordsDict)
    //fmt.Println("word1: ", word1,"word2: ", word2)
    ans := len(this.data)
    index1, index2 := -1, -1
    // 从左到右遍历数组 wordsDict，
    // 当遍历到 word1时，
    //     如果已经遍历的单词中存在 word2，为了计算最短距离，应该取最后一个已经遍历到的 word2 所在的下标，计算和当前下标的距离。
    //     同理，当遍历到 word2  时，应该取最后一个已经遍历到的 word1 所在的下标，计算和当前下标的距离。
    for i, word := range this.data {
        if word == word1 { // 找到单词1
            index1 = i
        } else if word == word2 { // 找到单词2
            index2 = i
        }
        //fmt.Println("i: ", i , "index1: ", index1 , "index2: ", index2)
        // 都找到了，计算距离
        if index1 >= 0 && index2 >= 0 {
            ans = min(ans, abs(index1 - index2))
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

type WordDistance1 struct {
    g map[string]([]int)
}

func Constructor1(wordsDict []string) WordDistance1 {
    this := WordDistance1{}
    g := make(map[string]([]int))
    for i, w := range wordsDict {
        g[w] = append(g[w], i)
    }
    this.g = g
    return this
}


func (this *WordDistance1) Shortest(word1 string, word2 string) (ans int) {
    g1, g2 := this.g[word1], this.g[word2]
    ans = 0x3f3f3f3f
    for i := 0; i < len(g1); i++ {
        for j := 0; j < len(g2); j++ {
            ans = min(ans, abs(g1[i] - g2[j]))
        }
    }
    return
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */

func main() {
    obj := Constructor([]string{"practice", "makes", "perfect", "coding", "makes"});
    fmt.Println(obj.Shortest("coding", "practice")) // 3
    fmt.Println(obj.Shortest("makes", "coding")) // 1


    obj1 := Constructor1([]string{"practice", "makes", "perfect", "coding", "makes"});
    fmt.Println(obj1.Shortest("coding", "practice")) // 3
    fmt.Println(obj1.Shortest("makes", "coding")) // 1
}