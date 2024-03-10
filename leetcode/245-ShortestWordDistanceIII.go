package main

// 245. Shortest Word Distance III
// Given an array of strings wordsDict and two strings that already exist in the array word1 and word2, 
// return the shortest distance between the occurrence of these two words in the list.
// Note that word1 and word2 may be the same. 
// It is guaranteed that they represent two individual words in the list.

// Example 1:
// Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "makes", word2 = "coding"
// Output: 1

// Example 2:
// Input: wordsDict = ["practice", "makes", "perfect", "coding", "makes"], word1 = "makes", word2 = "makes"
// Output: 3
 
// Constraints:
//         1 <= wordsDict.length <= 10^5
//         1 <= wordsDict[i].length <= 10
//         wordsDict[i] consists of lowercase English letters.
//         word1 and word2 are in wordsDict.

import "fmt"
import "math"

func shortestWordDistance(wordsDict []string, word1, word2 string) int {
    //fmt.Println("wordsDict: ", wordsDict)
    //fmt.Println("word1: ", word1,"word2: ", word2)
    ans := len(wordsDict)
    index1, index2 := -1, -1
    // 从左到右遍历数组 wordsDict，
    // 当遍历到 word1时，
    //     如果已经遍历的单词中存在 word2，为了计算最短距离，应该取最后一个已经遍历到的 word2 所在的下标，计算和当前下标的距离。
    //     同理，当遍历到 word2  时，应该取最后一个已经遍历到的 word1 所在的下标，计算和当前下标的距离。
    for i, word := range wordsDict {
        if word == word1 { // 找到单词1
            index1 = i
            // 如果 2 也找到了
            if index2 >= 0 {
                ans = min(ans, index1 - index2);
            }
        } 
        if word == word2 { // 找到单词2
            index2 = i
            // 如果 1 也找到了 且不是同一位置
            if(index1 >= 0 && index1 != index2) {
                ans = min(ans, index2 - index1);
            }
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

// best solution
func shortestWordDistance1(wordsDict []string, word1 string, word2 string) int {
	var ans = math.MaxInt32
	var index1, index2 = -1, -1
	for i, word := range wordsDict {
        // word1 和 word2 都一样的情况
		if word == word1 && word == word2 {
            // 找到第二个的时候
			if index1 != -1 {
				ans = min(ans, i - index1)
			}
			index1 = i
		} else if word == word1 { // 找到 word1 的时候
			if index2 != -1 {
				ans = min(ans, i - index2)
			}
			index1 = i
		} else if word == word2 { // 找到 word2 的时候
			if index1 != -1 {
				ans = min(ans, i - index1)
			}
			index2 = i
		}
	}
	return ans
}

func main() {
    fmt.Println(shortestWordDistance(
        []string{"practice", "makes", "perfect", "coding", "makes"},
        "makes",
        "coding",
    )) // 1

    fmt.Println(shortestWordDistance(
        []string{"practice", "makes", "perfect", "coding", "makes"},
        "makes",
        "makes",
    )) // 3

    fmt.Println(shortestWordDistance1(
        []string{"practice", "makes", "perfect", "coding", "makes"},
        "makes",
        "coding",
    )) // 1

    fmt.Println(shortestWordDistance1(
        []string{"practice", "makes", "perfect", "coding", "makes"},
        "makes",
        "makes",
    )) // 3
}