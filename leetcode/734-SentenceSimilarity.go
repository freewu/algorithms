package main

// 734. Sentence Similarity
// We can represent a sentence as an array of words, 
// for example, the sentence "I am happy with leetcode" can be represented as arr = ["I","am",happy","with","leetcode"].

// Given two sentences sentence1 and sentence2 each represented as a string array 
// and given an array of string pairs similarPairs where similarPairs[i] = [xi, yi] indicates 
// that the two words xi and yi are similar.

// Return true if sentence1 and sentence2 are similar, or false if they are not similar.
// Two sentences are similar if:
//     They have the same length (i.e., the same number of words)
//     sentence1[i] and sentence2[i] are similar.

// Notice that a word is always similar to itself, also notice that the similarity relation is not transitive. 
// For example, if the words a and b are similar, and the words b and c are similar, a and c are not necessarily similar.
 
// Example 1:
// Input: sentence1 = ["great","acting","skills"], sentence2 = ["fine","drama","talent"], similarPairs = [["great","fine"],["drama","acting"],["skills","talent"]]
// Output: true
// Explanation: The two sentences have the same length and each word i of sentence1 is also similar to the corresponding word in sentence2.

// Example 2:
// Input: sentence1 = ["great"], sentence2 = ["great"], similarPairs = []
// Output: true
// Explanation: A word is similar to itself.

// Example 3:
// Input: sentence1 = ["great"], sentence2 = ["doubleplus","good"], similarPairs = [["great","doubleplus"]]
// Output: false
// Explanation: As they don't have the same length, we return false.
 
// Constraints:
//     1 <= sentence1.length, sentence2.length <= 1000
//     1 <= sentence1[i].length, sentence2[i].length <= 20
//     sentence1[i] and sentence2[i] consist of English letters.
//     0 <= similarPairs.length <= 1000
//     similarPairs[i].length == 2
//     1 <= xi.length, yi.length <= 20
//     xi and yi consist of lower-case and upper-case English letters.
//     All the pairs (xi, yi) are distinct.

import "fmt"

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
    s1l, s2l := len(sentence1), len(sentence2)
    if s1l != s2l {
        return false
    }
    for i := 0; i < s1l; i++ {
        flag := false
        if sentence1[i] == sentence2[i] {  continue; }
        for _, v := range similarPairs {
            if v[0] == sentence1[i] && v[1] == sentence2[i] || v[1] == sentence1[i] && v[0] == sentence2[i]  {
                flag = true
                break
            }
        }
        if flag == false { return false; }
    }
    return true
}

func main() {
    // Example 1:
    // Input: sentence1 = ["great","acting","skills"], sentence2 = ["fine","drama","talent"], similarPairs = [["great","fine"],["drama","acting"],["skills","talent"]]
    // Output: true
    // Explanation: The two sentences have the same length and each word i of sentence1 is also similar to the corresponding word in sentence2.
    fmt.Println(areSentencesSimilar([]string{"great","acting","skills"},[]string{"fine","drama","talent"},[][]string{{"great","fine"},{"drama","acting"},{"skills","talent"}})) // true
    // Example 2:
    // Input: sentence1 = ["great"], sentence2 = ["great"], similarPairs = []
    // Output: true
    // Explanation: A word is similar to itself.
    fmt.Println(areSentencesSimilar([]string{"great"},[]string{"great"},[][]string{})) // true
    // Example 3:
    // Input: sentence1 = ["great"], sentence2 = ["doubleplus","good"], similarPairs = [["great","doubleplus"]]
    // Output: false
    // Explanation: As they don't have the same length, we return false.
    fmt.Println(areSentencesSimilar([]string{"great"},[]string{"doubleplus","good"},[][]string{{"great","doubleplus"}})) // false
}