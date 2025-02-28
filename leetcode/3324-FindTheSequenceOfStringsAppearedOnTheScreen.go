package main

// 3324. Find the Sequence of Strings Appeared on the Screen
// You are given a string target.

// Alice is going to type target on her computer using a special keyboard that has only two keys:
//     1. Key 1 appends the character "a" to the string on the screen.
//     2. Key 2 changes the last character of the string on the screen to its next character in the English alphabet. 
//        For example, "c" changes to "d" and "z" changes to "a".

// Note that initially there is an empty string "" on the screen, so she can only press key 1.

// Return a list of all strings that appear on the screen as Alice types target, in the order they appear, using the minimum key presses.

// Example 1:
// Input: target = "abc"
// Output: ["a","aa","ab","aba","abb","abc"]
// Explanation:
// The sequence of key presses done by Alice are:
// Press key 1, and the string on the screen becomes "a".
// Press key 1, and the string on the screen becomes "aa".
// Press key 2, and the string on the screen becomes "ab".
// Press key 1, and the string on the screen becomes "aba".
// Press key 2, and the string on the screen becomes "abb".
// Press key 2, and the string on the screen becomes "abc".

// Example 2:
// Input: target = "he"
// Output: ["a","b","c","d","e","f","g","h","ha","hb","hc","hd","he"]

// Constraints:
//     1 <= target.length <= 400
//     target consists only of lowercase English letters.

import "fmt"

func stringSequence(target string) []string {
    n, sum := len(target), 0
    for i := 0; i < n; i++ { // Counting the total number of rows in the result
        sum += int(target[i] - 'a' + 1)
    }
    res, subseq, index := make([]string, sum), make([]byte, 0, n), 0 // res fill pointer
    for i := 0; i < n; i++ {
        v, start := target[i], len(subseq) // current subseq len
        for j := uint8('a'); j <= v; j++ { // Add characters from 'a' to v
            subseq = append(subseq[:start], j)
            // Convert to string only before writing to result
            res[index] = string(subseq)
            index++
        }
    }
    return res
}

func stringSequence1(target string) []string {
    res, cur := make([]string, 0), ""
    for _, i := range target {
        s := 'a'
        for {
            res = append(res, cur + string(s))
            if s == i { break }
            s++
        }
        cur = cur + string(s)
    }
    return res
}

func main() {
    // Example 1:
    // Input: target = "abc"
    // Output: ["a","aa","ab","aba","abb","abc"]
    // Explanation:
    // The sequence of key presses done by Alice are:
    // Press key 1, and the string on the screen becomes "a".
    // Press key 1, and the string on the screen becomes "aa".
    // Press key 2, and the string on the screen becomes "ab".
    // Press key 1, and the string on the screen becomes "aba".
    // Press key 2, and the string on the screen becomes "abb".
    // Press key 2, and the string on the screen becomes "abc".
    fmt.Println(stringSequence("abc")) // ["a","aa","ab","aba","abb","abc"]
    // Example 2:
    // Input: target = "he"
    // Output: ["a","b","c","d","e","f","g","h","ha","hb","hc","hd","he"]
    fmt.Println(stringSequence("he")) // ["a","b","c","d","e","f","g","h","ha","hb","hc","hd","he"]

    fmt.Println(stringSequence("bluefrog")) // [a b ba bb bc bd be bf bg bh bi bj bk bl bla blb blc bld ble blf blg blh bli blj blk bll blm bln blo blp blq blr bls blt blu blua blub bluc blud blue bluea blueb bluec blued bluee bluef bluefa bluefb bluefc bluefd bluefe blueff bluefg bluefh bluefi bluefj bluefk bluefl bluefm bluefn bluefo bluefp bluefq bluefr bluefra bluefrb bluefrc bluefrd bluefre bluefrf bluefrg bluefrh bluefri bluefrj bluefrk bluefrl bluefrm bluefrn bluefro bluefroa bluefrob bluefroc bluefrod bluefroe bluefrof bluefrog]
    fmt.Println(stringSequence("leetcode")) // [a b c d e f g h i j k l la lb lc ld le lea leb lec led lee leea leeb leec leed leee leef leeg leeh leei leej leek leel leem leen leeo leep leeq leer lees leet leeta leetb leetc leetca leetcb leetcc leetcd leetce leetcf leetcg leetch leetci leetcj leetck leetcl leetcm leetcn leetco leetcoa leetcob leetcoc leetcod leetcoda leetcodb leetcodc leetcodd leetcode]

    fmt.Println(stringSequence1("abc")) // ["a","aa","ab","aba","abb","abc"]
    fmt.Println(stringSequence1("he")) // ["a","b","c","d","e","f","g","h","ha","hb","hc","hd","he"]
    fmt.Println(stringSequence1("bluefrog")) // [a b ba bb bc bd be bf bg bh bi bj bk bl bla blb blc bld ble blf blg blh bli blj blk bll blm bln blo blp blq blr bls blt blu blua blub bluc blud blue bluea blueb bluec blued bluee bluef bluefa bluefb bluefc bluefd bluefe blueff bluefg bluefh bluefi bluefj bluefk bluefl bluefm bluefn bluefo bluefp bluefq bluefr bluefra bluefrb bluefrc bluefrd bluefre bluefrf bluefrg bluefrh bluefri bluefrj bluefrk bluefrl bluefrm bluefrn bluefro bluefroa bluefrob bluefroc bluefrod bluefroe bluefrof bluefrog]
    fmt.Println(stringSequence1("leetcode")) // [a b c d e f g h i j k l la lb lc ld le lea leb lec led lee leea leeb leec leed leee leef leeg leeh leei leej leek leel leem leen leeo leep leeq leer lees leet leeta leetb leetc leetca leetcb leetcc leetcd leetce leetcf leetcg leetch leetci leetcj leetck leetcl leetcm leetcn leetco leetcoa leetcob leetcoc leetcod leetcoda leetcodb leetcodc leetcodd leetcode]
}
