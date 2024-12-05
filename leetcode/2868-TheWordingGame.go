package main

// 2868. The Wording Game
// Alice and Bob each have a lexicographically sorted array of strings named a and b respectively.

// They are playing a wording game with the following rules:
//     1. On each turn, the current player should play a word from their list 
//        such that the new word is closely greater than the last played word; then it's the other player's turn.
//     2. If a player can't play a word on their turn, they lose.

// Alice starts the game by playing her lexicographically smallest word.

// Given a and b, return true if Alice can win knowing that both players play their best, and false otherwise.

// A word w is closely greater than a word z if the following conditions are met:
//     1. w is lexicographically greater than z.
//     2. If w1 is the first letter of w and z1 is the first letter of z, 
//        w1 should either be equal to z1 or be the letter after z1 in the alphabet.
//     3. or example, the word "care" is closely greater than "book" and "car", 
//        but is not closely greater than "ant" or "cook".

// A string s is lexicographically greater than a string t if in the first position where s and t differ, 
// string s has a letter that appears later in the alphabet than the corresponding letter in t. 
// If the first min(s.length, t.length) characters do not differ, 
// then the longer string is the lexicographically greater one.

// Example 1:
// Input: a = ["avokado","dabar"], b = ["brazil"]
// Output: false
// Explanation: Alice must start the game by playing the word "avokado" since it's her smallest word, then Bob plays his only word, "brazil", which he can play because its first letter, 'b', is the letter after Alice's word's first letter, 'a'.
// Alice can't play a word since the first letter of the only word left is not equal to 'b' or the letter after 'b', 'c'.
// So, Alice loses, and the game ends.

// Example 2:
// Input: a = ["ananas","atlas","banana"], b = ["albatros","cikla","nogomet"]
// Output: true
// Explanation: Alice must start the game by playing the word "ananas".
// Bob can't play a word since the only word he has that starts with the letter 'a' or 'b' is "albatros", which is smaller than Alice's word.
// So Alice wins, and the game ends.

// Example 3:
// Input: a = ["hrvatska","zastava"], b = ["bijeli","galeb"]
// Output: true
// Explanation: Alice must start the game by playing the word "hrvatska".
// Bob can't play a word since the first letter of both of his words are smaller than the first letter of Alice's word, 'h'.
// So Alice wins, and the game ends.

// Constraints:
//     1 <= a.length, b.length <= 10^5
//     a[i] and b[i] consist only of lowercase English letters.
//     a and b are lexicographically sorted.
//     All the words in a and b combined are distinct.
//     The sum of the lengths of all the words in a and b combined does not exceed 10^6.

import "fmt"

func canAliceWin(a []string, b []string) bool {
    i, j, k, w := 1, 0, 1, a[0]
    for {
        if k&1 == 1 {
            if j == len(b) { return true }
            if (b[j][0] == w[0] && w < b[j]) || b[j][0]-w[0] == 1 {
                w = b[j]
                k ^= 1
            }
            j++
        } else {
            if i == len(a) { return false }
            if (a[i][0] == w[0] && w < a[i]) || a[i][0]-w[0] == 1 {
                w = a[i]
                k ^= 1
            }
            i++
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: a = ["avokado","dabar"], b = ["brazil"]
    // Output: false
    // Explanation: Alice must start the game by playing the word "avokado" since it's her smallest word, then Bob plays his only word, "brazil", which he can play because its first letter, 'b', is the letter after Alice's word's first letter, 'a'.
    // Alice can't play a word since the first letter of the only word left is not equal to 'b' or the letter after 'b', 'c'.
    // So, Alice loses, and the game ends.
    fmt.Println(canAliceWin([]string{"avokado","dabar"}, []string{"brazil"})) // false
    // Example 2:
    // Input: a = ["ananas","atlas","banana"], b = ["albatros","cikla","nogomet"]
    // Output: true
    // Explanation: Alice must start the game by playing the word "ananas".
    // Bob can't play a word since the only word he has that starts with the letter 'a' or 'b' is "albatros", which is smaller than Alice's word.
    // So Alice wins, and the game ends.
    fmt.Println(canAliceWin([]string{"ananas","atlas","banana"}, []string{"albatros","cikla","nogomet"})) // true
    // Example 3:
    // Input: a = ["hrvatska","zastava"], b = ["bijeli","galeb"]
    // Output: true
    // Explanation: Alice must start the game by playing the word "hrvatska".
    // Bob can't play a word since the first letter of both of his words are smaller than the first letter of Alice's word, 'h'.
    // So Alice wins, and the game ends.
    fmt.Println(canAliceWin([]string{"hrvatska","zastava"}, []string{"bijeli","galeb"})) // true
}