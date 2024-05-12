package main

// 288. Unique Word Abbreviation
// The abbreviation of a word is a concatenation of its first letter, the number of characters between the first and last letter, and its last letter. 
// If a word has only two characters, then it is an abbreviation of itself.
// For example:
//     dog --> d1g because there is one letter between the first letter 'd' and the last letter 'g'.
//     internationalization --> i18n because there are 18 letters between the first letter 'i' and the last letter 'n'.
//     it --> it because any word with only two characters is an abbreviation of itself.

// Implement the ValidWordAbbr class:
//     ValidWordAbbr(String[] dictionary) Initializes the object with a dictionary of words.
//     boolean isUnique(string word) Returns true if either of the following conditions are met (otherwise returns false):
//         There is no word in dictionary whose abbreviation is equal to word's abbreviation.
//         For any word in dictionary whose abbreviation is equal to word's abbreviation, that word and word are the same.

// Example 1:
// Input
// ["ValidWordAbbr", "isUnique", "isUnique", "isUnique", "isUnique", "isUnique"]
// [[["deer", "door", "cake", "card"]], ["dear"], ["cart"], ["cane"], ["make"], ["cake"]]
// Output
// [null, false, true, false, true, true]
// Explanation
// ValidWordAbbr validWordAbbr = new ValidWordAbbr(["deer", "door", "cake", "card"]);
// validWordAbbr.isUnique("dear"); // return false, dictionary word "deer" and word "dear" have the same abbreviation "d2r" but are not the same.
// validWordAbbr.isUnique("cart"); // return true, no words in the dictionary have the abbreviation "c2t".
// validWordAbbr.isUnique("cane"); // return false, dictionary word "cake" and word "cane" have the same abbreviation  "c2e" but are not the same.
// validWordAbbr.isUnique("make"); // return true, no words in the dictionary have the abbreviation "m2e".
// validWordAbbr.isUnique("cake"); // return true, because "cake" is already in the dictionary and no other word in the dictionary has "c2e" abbreviation.
 
// Constraints:
//     1 <= dictionary.length <= 3 * 10^4
//     1 <= dictionary[i].length <= 20
//     dictionary[i] consists of lowercase English letters.
//     1 <= word.length <= 20
//     word consists of lowercase English letters.
//     At most 5000 calls will be made to isUnique.

import "fmt"
import "strconv"

type ValidWordAbbr struct {
    dict1 map[string]bool
    dict2 map[string]string
}

func Constructor(dictionary []string) ValidWordAbbr {
    dict1, dict2 := make(map[string]bool), make(map[string]string)
    for _, word := range dictionary {
        abbr := GetAbbr(word) 
        v, ok := dict1[abbr]
        if !ok { // abbr 已注册 → 判断 word == dict2[abbr] → 不等于时进入终态
            dict1[abbr] = true
            dict2[abbr] = word // abbr未注册 → 注册状态 = dict1 true，dict2 存word
        } else if v && word != dict2[abbr] { 
            dict1[abbr] = false
            delete(dict2, abbr) // 终态= dict1 false, dict2 删除abbr
        }
    }
    return ValidWordAbbr{dict1, dict2}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
    abbr := GetAbbr(word)
    v, ok := this.dict1[abbr]
    if !ok { // abbr未注册和abbr已注册dict1[abbr]==true
        return true
    }
    if v && word == this.dict2[abbr] { // word==dict2[abbr]对应True状态
        return true
    }
    return false
}

func GetAbbr(word string) string {
    if len(word) <= 2 {
        return word
    }
    return string(word[0]) + strconv.Itoa(len(word)-2) + string(word[len(word)-1])
}

/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */

func main() {
    // ValidWordAbbr validWordAbbr = new ValidWordAbbr(["deer", "door", "cake", "card"]);
    obj := Constructor([]string{"deer", "door", "cake", "card"})
    // validWordAbbr.isUnique("dear"); // return false, dictionary word "deer" and word "dear" have the same abbreviation "d2r" but are not the same.
    fmt.Println(obj.IsUnique("dear")) // false
    // validWordAbbr.isUnique("cart"); // return true, no words in the dictionary have the abbreviation "c2t".
    fmt.Println(obj.IsUnique("cart")) // true
    // validWordAbbr.isUnique("cane"); // return false, dictionary word "cake" and word "cane" have the same abbreviation  "c2e" but are not the same.
    fmt.Println(obj.IsUnique("cane")) // false
    // validWordAbbr.isUnique("make"); // return true, no words in the dictionary have the abbreviation "m2e".
    fmt.Println(obj.IsUnique("make")) // true
    // validWordAbbr.isUnique("cake"); // return true, because "cake" is already in the dictionary and no other word in the dictionary has "c2e" abbreviation.
    fmt.Println(obj.IsUnique("cake")) // true
}