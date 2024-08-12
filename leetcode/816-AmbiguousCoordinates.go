package main

// 816. Ambiguous Coordinates
// We had some 2-dimensional coordinates, like "(1, 3)" or "(2, 0.5)". 
// Then, we removed all commas, decimal points, and spaces and ended up with the string s.
//     For example, "(1, 3)" becomes s = "(13)" and "(2, 0.5)" becomes s = "(205)".

// Return a list of strings representing all possibilities for what our original coordinates could have been.

// Our original representation never had extraneous zeroes, 
// so we never started with numbers like "00", "0.0", "0.00", "1.0", "001", "00.01", 
// or any other number that can be represented with fewer digits. 
// Also, a decimal point within a number never occurs without at least one digit occurring before it, 
// so we never started with numbers like ".1".

// The final answer list can be returned in any order. 
// All coordinates in the final answer have exactly one space between them (occurring after the comma.)

// Example 1:
// Input: s = "(123)"
// Output: ["(1, 2.3)","(1, 23)","(1.2, 3)","(12, 3)"]

// Example 2:
// Input: s = "(0123)"
// Output: ["(0, 1.23)","(0, 12.3)","(0, 123)","(0.1, 2.3)","(0.1, 23)","(0.12, 3)"]
// Explanation: 0.0, 00, 0001 or 00.01 are not allowed.

// Example 3:
// Input: s = "(00011)"
// Output: ["(0, 0.011)","(0.001, 1)"]

// Constraints:
//     4 <= s.length <= 12
//     s[0] == '(' and s[s.length - 1] == ')'.
//     The rest of s are digits.

import "fmt"

func ambiguousCoordinates(s string) []string {
    res := []string{}
    s = s[1:len(s) - 1] // remove parenthesis
    gen := func (s string, start, end int) []string {
        res := []string{} 
        for p := 1; p <= end - start; p++ { // partitioning for whole number and decimal part
            // whole number part
            w := s[start:start + p]
            // decimal number part
            d := s[start + p:end]
            // either whole number part is a zero or it cannot start with zero
            // and either decimal part doesn't exist or it cannot end with a zero
            if (w == "0" || w[0] != '0') && (len(d) == 0 || d[len(d) - 1] != '0') {
                if len(d) == 0 { // decimal part does not exist
                    res = append(res, w)
                } else { // decimal exists
                    res = append(res, w + "." + d)
                }
            }
        }
        return res
    }
    for i := 1; i < len(s); i++ {
        lefts, rights := gen(s, 0, i), gen(s, i, len(s))
        for _, left := range lefts { // for each possible pair of  coordinates
            for _, right := range rights {
                val := "(" + left + ", " + right + ")"
                res = append(res, val)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "(123)"
    // Output: ["(1, 2.3)","(1, 23)","(1.2, 3)","(12, 3)"]
    fmt.Println(ambiguousCoordinates("(123)")) // ["(1, 2.3)","(1, 23)","(1.2, 3)","(12, 3)"]
    // Example 2:
    // Input: s = "(0123)"
    // Output: ["(0, 1.23)","(0, 12.3)","(0, 123)","(0.1, 2.3)","(0.1, 23)","(0.12, 3)"]
    // Explanation: 0.0, 00, 0001 or 00.01 are not allowed.
    fmt.Println(ambiguousCoordinates("(0123)")) // ["(0, 1.23)","(0, 12.3)","(0, 123)","(0.1, 2.3)","(0.1, 23)","(0.12, 3)"]
    // Example 3:
    // Input: s = "(00011)"
    // Output: ["(0, 0.011)","(0.001, 1)"]
    fmt.Println(ambiguousCoordinates("(00011)")) // ["(0, 0.011)","(0.001, 1)"]
}