package main

// 2468. Split Message Based on Limit
// You are given a string, message, and a positive integer, limit.

// You must split message into one or more parts based on limit. 
// Each resulting part should have the suffix "<a/b>", where "b" is to be replaced with the total number of parts and "a" is to be replaced with the index of the part, starting from 1 and going up to b. 
// Additionally, the length of each resulting part (including its suffix) should be equal to limit, except for the last part whose length can be at most limit.

// The resulting parts should be formed such that when their suffixes are removed and they are all concatenated in order, they should be equal to message. 
// Also, the result should contain as few parts as possible.

// Return the parts message would be split into as an array of strings. 
// If it is impossible to split message as required, return an empty array.

// Example 1:
// Input: message = "this is really a very awesome message", limit = 9
// Output: ["thi<1/14>","s i<2/14>","s r<3/14>","eal<4/14>","ly <5/14>","a v<6/14>","ery<7/14>"," aw<8/14>","eso<9/14>","me<10/14>"," m<11/14>","es<12/14>","sa<13/14>","ge<14/14>"]
// Explanation:
// The first 9 parts take 3 characters each from the beginning of message.
// The next 5 parts take 2 characters each to finish splitting message. 
// In this example, each part, including the last, has length 9. 
// It can be shown it is not possible to split message into less than 14 parts.

// Example 2:
// Input: message = "short message", limit = 15
// Output: ["short mess<1/2>","age<2/2>"]
// Explanation:
// Under the given constraints, the string can be split into two parts: 
// - The first part comprises of the first 10 characters, and has a length 15.
// - The next part comprises of the last 3 characters, and has a length 8.

// Constraints:
//     1 <= message.length <= 10^4
//     message consists only of lowercase English letters and ' '.
//     1 <= limit <= 10^4

import "fmt"
import "math"
import "strconv"

func splitMessage(message string, limit int) []string {
    res := []string{}

    // Start by assuming the denominator (total number of splits) is
    // one digit long, and work up from there.
    for denomDigits := 1; denomDigits <= 4; denomDigits++ {
        numParts, m  := 0, message
        splits := []string{}
        for len(m) > 0 {
            numParts++
            numDigits := int(math.Log10(float64(numParts))) + 1
            if numDigits > denomDigits {
                splits = nil
                break
            }
            toGrab := limit - numDigits - denomDigits - 3
            if toGrab <= 0 {
                splits = nil
                break
            }
            if toGrab > len(m) {
                toGrab = len(m)
            }
            // Add our split to the list and slice off the message.
            splits = append(splits, m[0:toGrab])
            m = m[toGrab:]
        }
        if splits != nil {
            d := len(splits)
            for i, s := range splits {
                res = append(res, fmt.Sprintf("%s<%d/%d>", s, i + 1, d))
            }
            break
        }
    }
    return res
}

func splitMessage1(message string, limit int) []string {
    res := []string{}
    getPages := func(limit, msgLen int) int{
        left, initParts, extraSpace := limit - 5, 9, 0
        for msgLen > left * initParts + extraSpace {
            left -= 2
            extraSpace += initParts
            initParts = initParts * 10 + 9
            if left <= 0 {
                return 0
            }
        }
        pages := (msgLen - extraSpace) / left
        if  (msgLen - extraSpace) % left != 0 {
            pages++
        }
        return pages
    }
    pages := getPages(limit, len(message))
    pos := 0
    for i := 0; i < pages; i++ {
        tag := "<" + strconv.Itoa(i + 1) + "/" + strconv.Itoa(pages) + ">"
        wordLen := limit - len(tag)
        end := min(len(message), pos + wordLen)
        res = append(res, message[pos:end] + tag)
        pos += wordLen

    }
    return res
}

func main() {
    // Example 1:
    // Input: message = "this is really a very awesome message", limit = 9
    // Output: ["thi<1/14>","s i<2/14>","s r<3/14>","eal<4/14>","ly <5/14>","a v<6/14>","ery<7/14>"," aw<8/14>","eso<9/14>","me<10/14>"," m<11/14>","es<12/14>","sa<13/14>","ge<14/14>"]
    // Explanation:
    // The first 9 parts take 3 characters each from the beginning of message.
    // The next 5 parts take 2 characters each to finish splitting message. 
    // In this example, each part, including the last, has length 9. 
    // It can be shown it is not possible to split message into less than 14 parts.
    fmt.Println(splitMessage("this is really a very awesome message", 9)) // ["thi<1/14>","s i<2/14>","s r<3/14>","eal<4/14>","ly <5/14>","a v<6/14>","ery<7/14>"," aw<8/14>","eso<9/14>","me<10/14>"," m<11/14>","es<12/14>","sa<13/14>","ge<14/14>"]
    // Example 2:
    // Input: message = "short message", limit = 15
    // Output: ["short mess<1/2>","age<2/2>"]
    // Explanation:
    // Under the given constraints, the string can be split into two parts: 
    // - The first part comprises of the first 10 characters, and has a length 15.
    // - The next part comprises of the last 3 characters, and has a length 8.
    fmt.Println(splitMessage("short message", 15)) // ["short mess<1/2>","age<2/2>"]

    fmt.Println(splitMessage1("this is really a very awesome message", 9)) // ["thi<1/14>","s i<2/14>","s r<3/14>","eal<4/14>","ly <5/14>","a v<6/14>","ery<7/14>"," aw<8/14>","eso<9/14>","me<10/14>"," m<11/14>","es<12/14>","sa<13/14>","ge<14/14>"]
    fmt.Println(splitMessage1("short message", 15)) // ["short mess<1/2>","age<2/2>"]
}