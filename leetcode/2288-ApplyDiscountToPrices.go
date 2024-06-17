package main

// 2288. Apply Discount to Prices
// A sentence is a string of single-space separated words where each word can contain digits, lowercase letters, and the dollar sign '$'. 
// A word represents a price if it is a sequence of digits preceded by a dollar sign.
//     For example, "$100", "$23", and "$6" represent prices while "100", "$", and "$1e5" do not.

// You are given a string sentence representing a sentence and an integer discount. 
// For each word representing a price, apply a discount of discount% on the price and update the word in the sentence. 
// All updated prices should be represented with exactly two decimal places.

// Return a string representing the modified sentence.
// Note that all prices will contain at most 10 digits.

// Example 1:
// Input: sentence = "there are $1 $2 and 5$ candies in the shop", discount = 50
// Output: "there are $0.50 $1.00 and 5$ candies in the shop"
// Explanation: 
// The words which represent prices are "$1" and "$2". 
// - A 50% discount on "$1" yields "$0.50", so "$1" is replaced by "$0.50".
// - A 50% discount on "$2" yields "$1". Since we need to have exactly 2 decimal places after a price, we replace "$2" with "$1.00".

// Example 2:
// Input: sentence = "1 2 $3 4 $5 $6 7 8$ $9 $10$", discount = 100
// Output: "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"
// Explanation: 
// Applying a 100% discount on any price will result in 0.
// The words representing prices are "$3", "$5", "$6", and "$9".
// Each of them is replaced by "$0.00".

// Constraints:
//     1 <= sentence.length <= 10^5
//     sentence consists of lowercase English letters, digits, ' ', and '$'.
//     sentence does not have leading or trailing spaces.
//     All words in sentence are separated by a single space.
//     All prices will be positive numbers without leading zeros.
//     All prices will have at most 10 digits.
//     0 <= discount <= 100

import "fmt"
import "strings"
import "strconv"
import "regexp"

func discountPrices(sentence string, discount int) string {
    words := strings.Split(sentence, " ")
    for index, word := range words {
        if word[0] != '$' { // 不以 $ 开头
            continue
        }
        num, err := strconv.Atoi(word[1:len(word)]) // 取 $后面的转数字
        if err != nil {
            continue
        }
        discountPrice := fmt.Sprintf("%.2f", float64(num) * (100 - float64(discount)) / 100)
        words[index] = "$" + string(discountPrice)
    }
    return strings.Join(words, " ")
}

// 正则
func discountPrices1(sentence string, discount int) string {
    words := strings.Fields(sentence)
    r, _ := regexp.Compile(`\A\$\d+\z`)
    for i, w := range words {
        if r.MatchString(w) {
            dollari := strings.Index(w, "$")
            floatVal, _ := strconv.ParseFloat(w[dollari+1:], 64)
            np := fmt.Sprintf("$%.2f", floatVal * float64(100-discount)/100.0)
            words[i] = np
        }
    }
    return strings.Join(words, " ")
}

func discountPrices2(sentence string, discount int) string {
    words := strings.Split(sentence, " ")
    for i, word := range words {
        if word[0] == '$' {
            if v, err := strconv.Atoi(word[1:]); err == nil {
                words[i] = fmt.Sprintf("$%.2f", float64(v * (100 - discount)) / 100)
            }
        }
    }
    return strings.Join(words, " ")
}

func main() {
    // Example 1:
    // Input: sentence = "there are $1 $2 and 5$ candies in the shop", discount = 50
    // Output: "there are $0.50 $1.00 and 5$ candies in the shop"
    // Explanation: 
    // The words which represent prices are "$1" and "$2". 
    // - A 50% discount on "$1" yields "$0.50", so "$1" is replaced by "$0.50".
    // - A 50% discount on "$2" yields "$1". Since we need to have exactly 2 decimal places after a price, we replace "$2" with "$1.00".
    fmt.Println(discountPrices("there are $1 $2 and 5$ candies in the shop", 50)) // "there are $0.50 $1.00 and 5$ candies in the shop"
    // Example 2:
    // Input: sentence = "1 2 $3 4 $5 $6 7 8$ $9 $10$", discount = 100
    // Output: "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"
    // Explanation: 
    // Applying a 100% discount on any price will result in 0.
    // The words representing prices are "$3", "$5", "$6", and "$9".
    // Each of them is replaced by "$0.00".
    fmt.Println(discountPrices("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100)) // "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"

    fmt.Println(discountPrices1("there are $1 $2 and 5$ candies in the shop", 50)) // "there are $0.50 $1.00 and 5$ candies in the shop"
    fmt.Println(discountPrices1("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100)) // "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"

    fmt.Println(discountPrices2("there are $1 $2 and 5$ candies in the shop", 50)) // "there are $0.50 $1.00 and 5$ candies in the shop"
    fmt.Println(discountPrices2("1 2 $3 4 $5 $6 7 8$ $9 $10$", 100)) // "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"
}