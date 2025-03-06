package main

// 面试题 16.02. Words Frequency LCCI
// Design a method to find the frequency of occurrences of any given word in a book. 
// What if we were running this algorithm multiple times?

// You should implement following methods:
//     WordsFrequency(book) constructor, parameter is a array of strings, representing the book.
//     get(word) get the frequency of word in the book. 

// Example:
// WordsFrequency wordsFrequency = new WordsFrequency({"i", "have", "an", "apple", "he", "have", "a", "pen"});
// wordsFrequency.get("you"); //returns 0，"you" is not in the book
// wordsFrequency.get("have"); //returns 2，"have" occurs twice in the book
// wordsFrequency.get("an"); //returns 1
// wordsFrequency.get("apple"); //returns 1
// wordsFrequency.get("pen"); //returns 1

// Note:
//     There are only lowercase letters in book[i].
//     1 <= book.length <= 100000
//     1 <= book[i].length <= 10
//     get function will not be called more than 100000 times.

import "fmt"

type WordsFrequency struct {
    data map[string]int
}

func Constructor(book []string) WordsFrequency {
    mp := make(map[string]int)
    for _, v := range book {
        mp[v]++
    }
    return WordsFrequency{ data: mp }
}

func (this *WordsFrequency) Get(word string) int {
    return this.data[word]
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */

func main() {
    // WordsFrequency wordsFrequency = new WordsFrequency({"i", "have", "an", "apple", "he", "have", "a", "pen"});
    obj := Constructor([]string{"i", "have", "an", "apple", "he", "have", "a", "pen"})
    // wordsFrequency.get("you"); //returns 0，"you" is not in the book
    fmt.Println(obj.Get("you")) // 0
    // wordsFrequency.get("have"); //returns 2，"have" occurs twice in the book
    fmt.Println(obj.Get("have")) // 2
    // wordsFrequency.get("an"); //returns 1
    fmt.Println(obj.Get("an")) // 1
    // wordsFrequency.get("apple"); //returns 1
    fmt.Println(obj.Get("apple")) // 1
    // wordsFrequency.get("pen"); //returns 1
    fmt.Println(obj.Get("pen")) // 1
}