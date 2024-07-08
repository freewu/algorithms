package main

// 1286. Iterator for Combination
// Design the CombinationIterator class:
//     CombinationIterator(string characters, int combinationLength) Initializes the object with a string characters of sorted distinct lowercase English letters and a number combinationLength as arguments.
//     next() Returns the next combination of length combinationLength in lexicographical order.
//     hasNext() Returns true if and only if there exists a next combination.

// Example 1:
// Input
// ["CombinationIterator", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
// [["abc", 2], [], [], [], [], [], []]
// Output
// [null, "ab", true, "ac", true, "bc", false]
// Explanation
// CombinationIterator itr = new CombinationIterator("abc", 2);
// itr.next();    // return "ab"
// itr.hasNext(); // return True
// itr.next();    // return "ac"
// itr.hasNext(); // return True
// itr.next();    // return "bc"
// itr.hasNext(); // return False

// Constraints:
//     1 <= combinationLength <= characters.length <= 15
//     All the characters of characters are unique.
//     At most 10^4 calls will be made to next and hasNext.
//     It is guaranteed that all calls of the function next are valid.

import "fmt"

type CombinationIterator struct {
    index int
    combination []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
    combination := []string{}
    var dfs func (start, length int, s, chars string) 
    dfs = func(start, length int, s, chars string) {
        if length == 0 { 
            combination = append(combination, s)
            return 
        }
        for i := start; i < len(chars); i++ { 
            dfs(i + 1, length - 1, s + string(chars[i]), chars) 
        }
    }
    dfs(0, combinationLength, "", characters)
    return CombinationIterator{ 0, combination }
}

func (this *CombinationIterator) Next() string {
    v := this.combination[this.index]
    this.index++
    return v
}

func (this *CombinationIterator) HasNext() bool {
    return this.index < len(this.combination)
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
    // CombinationIterator itr = new CombinationIterator("abc", 2);
    obj := Constructor("abc", 2)
    fmt.Println(obj)
    // itr.next();    // return "ab"
    fmt.Println(obj.Next()) // "ab"
    // itr.hasNext(); // return True
    fmt.Println(obj.HasNext()) // true
    // itr.next();    // return "ac"
    fmt.Println(obj.Next()) // "ac"
    // itr.hasNext(); // return True
    fmt.Println(obj.HasNext()) // true
    // itr.next();    // return "bc"
    fmt.Println(obj.Next()) // "bc"
    // itr.hasNext(); // return False
    fmt.Println(obj.HasNext()) // false
}