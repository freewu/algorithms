package main

// 831. Masking Personal Information
// You are given a personal information string s, representing either an email address or a phone number. 
// Return the masked personal information using the below rules.

// Email address:
// An email address is:
//     A name consisting of uppercase and lowercase English letters, followed by
//     The '@' symbol, followed by
//     The domain consisting of uppercase and lowercase English letters with a dot '.' somewhere in the middle (not the first or last character).

// To mask an email:
//     The uppercase letters in the name and domain must be converted to lowercase letters.
//     The middle letters of the name (i.e., all but the first and last letters) must be replaced by 5 asterisks "*****".

// Phone number:
// A phone number is formatted as follows:
//     The phone number contains 10-13 digits.
//     The last 10 digits make up the local number.
//     The remaining 0-3 digits, in the beginning, make up the country code.
//     Separation characters from the set {'+', '-', '(', ')', ' '} separate the above digits in some way.

// To mask a phone number:
//     Remove all separation characters.
//     The masked phone number should have the form:
//     "***-***-XXXX" if the country code has 0 digits.
//     "+*-***-***-XXXX" if the country code has 1 digit.
//     "+**-***-***-XXXX" if the country code has 2 digits.
//     "+***-***-***-XXXX" if the country code has 3 digits.
//     "XXXX" is the last 4 digits of the local number.

// Example 1:
// Input: s = "LeetCode@LeetCode.com"
// Output: "l*****e@leetcode.com"
// Explanation: s is an email address.
// The name and domain are converted to lowercase, and the middle of the name is replaced by 5 asterisks.

// Example 2:
// Input: s = "AB@qq.com"
// Output: "a*****b@qq.com"
// Explanation: s is an email address.
// The name and domain are converted to lowercase, and the middle of the name is replaced by 5 asterisks.
// Note that even though "ab" is 2 characters, it still must have 5 asterisks in the middle.

// Example 3:
// Input: s = "1(234)567-890"
// Output: "***-***-7890"
// Explanation: s is a phone number.
// There are 10 digits, so the local number is 10 digits and the country code is 0 digits.
// Thus, the resulting masked number is "***-***-7890".

// Constraints:
//     s is either a valid email or a phone number.
//     If s is an email:
//     8 <= s.length <= 40
//     s consists of uppercase and lowercase English letters and exactly one '@' symbol and '.' symbol.
//     If s is a phone number:
//     10 <= s.length <= 20
//     s consists of digits, spaces, and the symbols '(', ')', '-', and '+'.

import "fmt"
import "strings"
import "unicode"

func maskPII(s string) string {
    if strings.Contains(s, "@") {
        maskEmail := func(s string) string {
            parts := strings.Split(s, "@")
            name, domain := parts[0], parts[1]
            name = string(name[0]) + "*****" + string(name[len(name) - 1])
            return name + "@" + domain
        }
        return maskEmail(strings.ToLower(s))
    }
    maskPhone := func(s string) string {
        vals := ""
        for i := 0; i < len(s); i++ { // 只留 0-9 的数字
            if s[i] >= '0' && s[i] <= '9' {
                vals += string(s[i])
            }
        }
        country := []string{"", "+*-", "+**-", "+***-"}
        return country[len(vals)-10] + "***-***-" + vals[len(vals)-4:]
    }
    return maskPhone(s)
}

func maskPII1(s string) string {
    at := strings.Index(s, "@")
    if at > 0 {
        s = strings.ToLower(s)
        return string(s[0]) + "*****" + s[at-1:]
    }
    var sb strings.Builder
    for _, r := range s {
        if unicode.IsDigit(r) {
            sb.WriteRune(r)
        }
    }
    s = sb.String()
    country := []string{"", "+*-", "+**-", "+***-"}
    return country[len(s)-10] + "***-***-" + s[len(s)-4:]
}

func main() {
    // Example 1:
    // Input: s = "LeetCode@LeetCode.com"
    // Output: "l*****e@leetcode.com"
    // Explanation: s is an email address.
    // The name and domain are converted to lowercase, and the middle of the name is replaced by 5 asterisks.
    fmt.Println(maskPII("LeetCode@LeetCode.com")) // "l*****e@leetcode.com"
    // Example 2:
    // Input: s = "AB@qq.com"
    // Output: "a*****b@qq.com"
    // Explanation: s is an email address.
    // The name and domain are converted to lowercase, and the middle of the name is replaced by 5 asterisks.
    // Note that even though "ab" is 2 characters, it still must have 5 asterisks in the middle.
    fmt.Println(maskPII("AB@qq.com")) // "a*****b@qq.com"
    // Example 3:
    // Input: s = "1(234)567-890"
    // Output: "***-***-7890"
    // Explanation: s is a phone number.
    // There are 10 digits, so the local number is 10 digits and the country code is 0 digits.
    // Thus, the resulting masked number is "***-***-7890".
    fmt.Println(maskPII("1(234)567-890")) // "***-***-7890"

    fmt.Println(maskPII1("LeetCode@LeetCode.com")) // "l*****e@leetcode.com"
    fmt.Println(maskPII1("AB@qq.com")) // "a*****b@qq.com"
    fmt.Println(maskPII1("1(234)567-890")) // "***-***-7890"
}