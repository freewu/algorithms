package main

// 843. Guess the Word
// You are given an array of unique strings words where words[i] is six letters long. 
// One word of words was chosen as a secret word.

// You are also given the helper object Master. 
// You may call Master.guess(word) where word is a six-letter-long string, and it must be from words. 
// Master.guess(word) returns:
//     -1 if word is not from words, or
//     an integer representing the number of exact matches (value and position) of your guess to the secret word.

// There is a parameter allowedGuesses for each test case where allowedGuesses is the maximum number of times you can call Master.guess(word).

// For each test case, you should call Master.guess with the secret word without exceeding the maximum number of allowed guesses. You will get:
//     1. "Either you took too many guesses, or you did not find the secret word." 
//         if you called Master.guess more than allowedGuesses times or 
//         if you did not call Master.guess with the secret word, or
//     2. "You guessed the secret word correctly." 
//         if you called Master.guess with the secret word with the number of calls to Master.guess less than or equal to allowedGuesses.

// The test cases are generated such that you can guess the secret word with a reasonable strategy (other than using the bruteforce method).

// Example 1:
// Input: secret = "acckzz", words = ["acckzz","ccbazz","eiowzz","abcczz"], allowedGuesses = 10
// Output: You guessed the secret word correctly.
// Explanation:
// master.guess("aaaaaa") returns -1, because "aaaaaa" is not in wordlist.
// master.guess("acckzz") returns 6, because "acckzz" is secret and has all 6 matches.
// master.guess("ccbazz") returns 3, because "ccbazz" has 3 matches.
// master.guess("eiowzz") returns 2, because "eiowzz" has 2 matches.
// master.guess("abcczz") returns 4, because "abcczz" has 4 matches.
// We made 5 calls to master.guess, and one of them was the secret, so we pass the test case.

// Example 2:
// Input: secret = "hamada", words = ["hamada","khaled"], allowedGuesses = 10
// Output: You guessed the secret word correctly.
// Explanation: Since there are two words, you can guess both.

// Constraints:
//     1 <= words.length <= 100
//     words[i].length == 6
//     words[i] consist of lowercase English letters.
//     All the strings of wordlist are unique.
//     secret exists in words.
//     10 <= allowedGuesses <= 30

import "fmt"

type Master struct {
}

func (this *Master) Guess(word string) int {
    return 1
}

/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */
func findSecretWord(wordlist []string, master *Master) {
    // map word to a score 
    wordMap := make(map[string]int)
    for i := range wordlist {
        wordMap[wordlist[i]] = 0
    }
    getMaxKey := func(m map[string]int) string {
        res, max := "", -1
        for k, v := range m {
            if v > max {
                max = v
                res = k
            }
        }
        delete(m, res)
        return res
    }
    pruneMap := func(m map[string]int, word string, score int) string {
        maxScore, bestGuess := -1, ""
        for k := range m {
            for i := range word {
                if k[i] == word[i] {
                    m[k] += score
                    if score == 0 {
                        delete(m, k)
                    }
                }
            }
            if m[k] > maxScore {
                maxScore = m[k]
                bestGuess = k
            }
        }
        delete(m, bestGuess)
        return bestGuess
    }
    // just use first word as first guess
    guess := wordlist[0]
    for i := 0; i < 10; i++ {
        score := master.Guess(guess)
        if score == -1 {
            guess = getMaxKey(wordMap) // if there were no matches to guess, then use the less best guess
        }
        guess = pruneMap(wordMap, guess, score) // re-update map scores and return best guess with highest score
    }
}

func findSecretWord1(words []string, master *Master) {
    mp, g := map[string]bool{}, ""
    for _, w := range words {
        mp[w] = true
    }
    check := func (a, b string) int {
        res := 0
        for i := 0; i < 6; i++ {
            if a[i] == b[i] {
            res++
            }
        }
        return res
    }
    for {
        // pick a word and remove from map
        for w, _ := range mp {
            g = w
            delete(mp,g)
            break
        }
        n := master.Guess(g)
        if n == 6 { // correct guess
            return
        }
        // remove all impossible words
        for w, _ := range mp {
            if check(g,w) != n {
                delete(mp,w)
            }
        }
    }
}

func main() {
    // Example 1:
    // Input: secret = "acckzz", words = ["acckzz","ccbazz","eiowzz","abcczz"], allowedGuesses = 10
    // Output: You guessed the secret word correctly.
    // Explanation:
    // master.guess("aaaaaa") returns -1, because "aaaaaa" is not in wordlist.
    // master.guess("acckzz") returns 6, because "acckzz" is secret and has all 6 matches.
    // master.guess("ccbazz") returns 3, because "ccbazz" has 3 matches.
    // master.guess("eiowzz") returns 2, because "eiowzz" has 2 matches.
    // master.guess("abcczz") returns 4, because "abcczz" has 4 matches.
    // We made 5 calls to master.guess, and one of them was the secret, so we pass the test case.
    fmt,Println(findSecretWord([]string{"acckzz","ccbazz","eiowzz","abcczz"}, master1)) // You guessed the secret word correctly.
    // Example 2:
    // Input: secret = "hamada", words = ["hamada","khaled"], allowedGuesses = 10
    // Output: You guessed the secret word correctly.
    // Explanation: Since there are two words, you can guess both.
    fmt,Println(findSecretWord([]string{"hamada","khaled"}, master2)) // You guessed the secret word correctly.
}