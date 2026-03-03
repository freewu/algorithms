package main

// 3860. Unique Email Groups
// You are given an array of strings emails, where each string is a valid email address.

// Two email addresses belong to the same group if both their normalized local names and normalized domain names are identical.

// The normalization rules are as follows:
// 1 The local name is the part before the '@' symbol.
//     1.1 Ignore all dots '.'.
//     1.2 Ignore everything after the first '+', if present.
//     1.3 Convert to lowercase.
// 2 The domain name is the part after the '@' symbol.
//     2.1 Convert to lowercase.

// Return an integer denoting the number of unique email groups after normalization.

// Example 1:
// Input: emails = ["test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"]
// Output: 2
// Explanation:
// Email                               | Local                 | Normalized Local       | Domain            | Normalized Domain | Final Email  
// test.email+alex@leetcode.com	    | test.email+alex	    | testemail	             | leetcode.com      | leetcode.com	     | testemail@leetcode.com
// test.e.mail+bob.cathy@leetcode.com	| test.e.mail+bob.cathy	| testemail	             | leetcode.com      | leetcode.com	     | testemail@leetcode.com  
// testemail+david@lee.tcode.com	    | testemail+david	    | testemail	             | lee.tcode.com	 | lee.tcode.com	 | testemail@lee.tcode.com
// Unique emails are ["testemail@leetcode.com", "testemail@lee.tcode.com"]. Thus, the answer is 2.

// Example 2:
// Input: emails = ["A@B.com", "a@b.com", "ab+xy@b.com", "a.b@b.com"]
// Output: 2
// Explanation:
// Email                   | Local             | Normalized Local  | Domain            | Normalized Domain | Final Email  
// A@B.com	                | A         	    | a	                | B.com	            | b.com	            | a@b.com
// a@b.com	                | a	                | a	                | b.com	            | b.com	            | a@b.com
// ab+xy@b.com	            | ab+xy	            | ab	            | b.com	            | b.com	            | ab@b.com
// a.b@b.com	            | a.b	            | ab	            | b.com	            | b.com	            | ab@b.com
// Unique emails are ["a@b.com", "ab@b.com"]. Thus, the answer is 2.

// Example 3:
// Input: emails = ["a.b+c.d+e@DoMain.com", "ab+xyz@domain.com", "ab@domain.com"]
// Output: 1
// Explanation:
// Email                   | Local             | Normalized Local  | Domain            | Normalized Domain | Final Email   
// a.b+c.d+e@DoMain.com	| a.b+c.d+e	        | ab	            | DoMain.com	    | domain.com	    | ab@domain.com
// ab+xyz@domain.com	    | ab+xyz	        | ab	            | domain.com	    | domain.com	    | ab@domain.com
// ab@domain.com           | ab                | ab	            | domain.com	    | domain.com	    | ab@domain.com
// All emails normalize to "ab@domain.com". Thus, the answer is 1.

// Constraints:
//     1 <= emails.length <= 1000
//     1 <= emails[i].length <= 100
//     emails[i] consists of lowercase and uppercase English letters, digits, and the characters '.', '+', and '@'.
//     Each emails[i] contains exactly one '@' character.
//     All local and domain names are non-empty; local names do not start with '+'.
//     Domain names end with the ".com" suffix and contain at least one character before ".com".

import "fmt"

func uniqueEmailGroups(emails []string) int {
    mp  := make(map[string]bool) // 用map存储标准化后的唯一邮箱地址
    toLowerRune := func(c rune) rune { // 将单个字符转为小写
        if c >= 'A' && c <= 'Z' { return c + 32 }
        return c
    }
    toLower := func(s string) string { // 将字符串转为小写
        builder := []rune{}
        for _, c := range s {
            builder = append(builder, toLowerRune(c))
        }
        return string(builder)
    }
    normalizeLocalName := func(local string) string { // 标准化本地名：移除点、截断+后的内容、转小写
        var builder []rune
        for _, c := range local {
            if c == '+' { break } // 遇到第一个+就停止处理
            if c == '.' { continue  }  // 忽略点号
            builder = append(builder, toLowerRune(c)) // 转换为小写并添加到builder
        }
        return string(builder)
    }
    for _, email := range emails {
        index := -1 // 拆分本地名和域名 
        for i, c := range email { // 找到 @ 的位置
            if c == '@' {
                index = i
                break
            }
        }
        local, domain := email[:index], email[index:]
        normalizedLocal := normalizeLocalName(local) // 标准化本地名
        normalizedDomain := toLower(domain)  // 标准化域名（转小写）
        mp[normalizedLocal + normalizedDomain] = true  // 拼接标准化后的完整邮箱,存入mp
    }
    // map的长度即为唯一邮箱组的数量
    return len(mp)
}

func main() {
    // Example 1:
    // Input: emails = ["test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"]
    // Output: 2
    // Explanation:
    // Email                               | Local                 | Normalized Local       | Domain            | Normalized Domain | Final Email  
    // test.email+alex@leetcode.com	    | test.email+alex	    | testemail	             | leetcode.com      | leetcode.com	     | testemail@leetcode.com
    // test.e.mail+bob.cathy@leetcode.com	| test.e.mail+bob.cathy	| testemail	             | leetcode.com      | leetcode.com	     | testemail@leetcode.com  
    // testemail+david@lee.tcode.com	    | testemail+david	    | testemail	             | lee.tcode.com	 | lee.tcode.com	 | testemail@lee.tcode.com
    // Unique emails are ["testemail@leetcode.com", "testemail@lee.tcode.com"]. Thus, the answer is 2.
    fmt.Println(uniqueEmailGroups([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})) // 2
    // Example 2:
    // Input: emails = ["A@B.com", "a@b.com", "ab+xy@b.com", "a.b@b.com"]
    // Output: 2
    // Explanation:
    // Email                   | Local             | Normalized Local  | Domain            | Normalized Domain | Final Email  
    // A@B.com	                | A         	    | a	                | B.com	            | b.com	            | a@b.com
    // a@b.com	                | a	                | a	                | b.com	            | b.com	            | a@b.com
    // ab+xy@b.com	            | ab+xy	            | ab	            | b.com	            | b.com	            | ab@b.com
    // a.b@b.com	            | a.b	            | ab	            | b.com	            | b.com	            | ab@b.com
    // Unique emails are ["a@b.com", "ab@b.com"]. Thus, the answer is 2.
    fmt.Println(uniqueEmailGroups([]string{"A@B.com", "a@b.com", "ab+xy@b.com", "a.b@b.com"})) // 2
    // Example 3:
    // Input: emails = ["a.b+c.d+e@DoMain.com", "ab+xyz@domain.com", "ab@domain.com"]
    // Output: 1
    // Explanation:
    // Email                   | Local             | Normalized Local  | Domain            | Normalized Domain | Final Email   
    // a.b+c.d+e@DoMain.com	| a.b+c.d+e	        | ab	            | DoMain.com	    | domain.com	    | ab@domain.com
    // ab+xyz@domain.com	    | ab+xyz	        | ab	            | domain.com	    | domain.com	    | ab@domain.com
    // ab@domain.com           | ab                | ab	            | domain.com	    | domain.com	    | ab@domain.com
    // All emails normalize to "ab@domain.com". Thus, the answer is 1. 
    fmt.Println(uniqueEmailGroups([]string{"a.b+c.d+e@DoMain.com", "ab+xyz@domain.com", "ab@domain.com"})) // 1

    fmt.Println(uniqueEmailGroups([]string{"bluefrog@leetcode.com", "admin@leetcode.com", "freewu@lee.tcode.com"})) // 3
}