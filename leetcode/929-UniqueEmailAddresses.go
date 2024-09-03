package main

// 929. Unique Email Addresses
// Every valid email consists of a local name and a domain name, separated by the '@' sign. 
// Besides lowercase letters, the email may contain one or more '.' or '+'.
//     For example, in "alice@leetcode.com", "alice" is the local name, and "leetcode.com" is the domain name.

// If you add periods '.' between some characters in the local name part of an email address, 
// mail sent there will be forwarded to the same address without dots in the local name.
// Note that this rule does not apply to domain names.
//     For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.

// If you add a plus '+' in the local name, everything after the first plus sign will be ignored. 
// This allows certain emails to be filtered. Note that this rule does not apply to domain names.
//     For example, "m.y+name@email.com" will be forwarded to "my@email.com".

// It is possible to use both of these rules at the same time.

// Given an array of strings emails where we send one email to each emails[i], 
// return the number of different addresses that actually receive mails.

// Example 1:
// Input: emails = ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
// Output: 2
// Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails.

// Example 2:
// Input: emails = ["a@leetcode.com","b@leetcode.com","c@leetcode.com"]
// Output: 3

// Constraints:
//     1 <= emails.length <= 100
//     1 <= emails[i].length <= 100
//     emails[i] consist of lowercase English letters, '+', '.' and '@'.
//     Each emails[i] contains exactly one '@' character.
//     All local and domain names are non-empty.
//     Local names do not start with a '+' character.
//     Domain names end with the ".com" suffix.

import "fmt"
import "strings"

func numUniqueEmails(emails []string) int {
    mp := map[string]int{}
    for _, email := range emails {
        t1, t2 := "", ""
        for i, v := range email {
            if v != '.' && v != '@' {
                t1 = t1 + string(v)
            } else if v == '@' {
                t2 = email[i:]
                break
            }
        }
        for i, v := range t1 {
            if v == '+' {
                t1 = t1[:i]
                break
            }
        }
        mp[t1+t2]++
    }
    return len(mp)
}

func numUniqueEmails1(emails []string) int {
    mp, str := make(map[string]struct{}), []byte{}
    for _, email := range emails {
        str = str[:0]
        index := strings.IndexByte(email, '@') // 找到 @ 的位置
        username, host := email[index:], email[:index]
        rep := strings.ReplaceAll(host, ".", "")
        before, _, _ := strings.Cut(rep, "+")
        str = append(str, []byte(before)...)
        str = append(str, username...)
        mp[string(str)] = struct{}{}
    }
    return len(mp)
}

func main() {
    // Example 1:
    // Input: emails = ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
    // Output: 2
    // Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails.
    fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"})) // 2
    // Example 2:
    // Input: emails = ["a@leetcode.com","b@leetcode.com","c@leetcode.com"]
    // Output: 3
    fmt.Println(numUniqueEmails([]string{"a@leetcode.com","b@leetcode.com","c@leetcode.com"})) // 3

    fmt.Println(numUniqueEmails1([]string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"})) // 2
    fmt.Println(numUniqueEmails1([]string{"a@leetcode.com","b@leetcode.com","c@leetcode.com"})) // 3
}