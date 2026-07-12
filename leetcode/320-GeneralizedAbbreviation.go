package main

// 320. Generalized Abbreviation
// A word's generalized abbreviation can be constructed by taking any number of non-overlapping and non-adjacent 
// substrings and replacing them with their respective lengths.

// For example, "abcde" can be abbreviated into:
//     "a3e" ("bcd" turned into "3")
//     "1bcd1" ("a" and "e" both turned into "1")
//     "5" ("abcde" turned into "5")
//     "abcde" (no substrings replaced)
// However, these abbreviations are invalid:
//     "23" ("ab" turned into "2" and "cde" turned into "3") is invalid as the substrings chosen are adjacent.
//     "22de" ("ab" turned into "2" and "bc" turned into "2") is invalid as the substring chosen overlap.

// Given a string word, return a list of all the possible generalized abbreviations of word. 
// Return the answer in any order.

// Example 1:
// Input: word = "word"
// Output: ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]

// Example 2:
// Input: word = "a"
// Output: ["1","a"]
 
// Constraints:
//     1 <= word.length <= 15
//     word consists of only lowercase English letters.

import "fmt"
import "strconv"

func generateAbbreviations(word string) []string {
    n, res := len(word), []string{}
    var dfs func(path []byte, p int, c bool)
    dfs = func(path []byte, p int, c bool){
        if p == n {
            res = append(res, string(path))
            return
        }
        if c { // 当前可以缩写
            for i := 1; i <= n - p; i++ { // 分别缩写多个长度
                dfs(append(path, []byte(strconv.Itoa(i))...), p + i, false)
            }
        } else { // 不可以缩写
            for i := 1; i <= n - p; i++ { // 分别不缩写多个长度
                dfs(append(path,word[p:p+i]...), p+i, true)
            }
        }
    }
    dfs([]byte{}, 0, true)
    dfs([]byte{}, 0, false)
    return res
}

func main() {
    // Example 1:
    // Input: word = "word"
    // Output: ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]
    fmt.Println(generateAbbreviations("word")) // ["4","3d","2r1","2rd","1o2","1o1d","1or1","1ord","w3","w2d","w1r1","w1rd","wo2","wo1d","wor1","word"]
    // Example 2:
    // Input: word = "a"
    // Output: ["1","a"]
    fmt.Println(generateAbbreviations("a")) // ["1","a"]

    fmt.Println(generateAbbreviations("bluefrog")) // [1l1e1r1g 1l1e1r2 1l1e1ro1 1l1e1rog 1l1e2o1 1l1e2og 1l1e3g 1l1e4 1l1ef1o1 1l1ef1og 1l1ef2g 1l1ef3 1l1efr1g 1l1efr2 1l1efro1 1l1efrog 1l2f1o1 1l2f1og 1l2f2g 1l2f3 1l2fr1g 1l2fr2 1l2fro1 1l2frog 1l3r1g 1l3r2 1l3ro1 1l3rog 1l4o1 1l4og 1l5g 1l6 1lu1f1o1 1lu1f1og 1lu1f2g 1lu1f3 1lu1fr1g 1lu1fr2 1lu1fro1 1lu1frog 1lu2r1g 1lu2r2 1lu2ro1 1lu2rog 1lu3o1 1lu3og 1lu4g 1lu5 1lue1r1g 1lue1r2 1lue1ro1 1lue1rog 1lue2o1 1lue2og 1lue3g 1lue4 1luef1o1 1luef1og 1luef2g 1luef3 1luefr1g 1luefr2 1luefro1 1luefrog 2u1f1o1 2u1f1og 2u1f2g 2u1f3 2u1fr1g 2u1fr2 2u1fro1 2u1frog 2u2r1g 2u2r2 2u2ro1 2u2rog 2u3o1 2u3og 2u4g 2u5 2ue1r1g 2ue1r2 2ue1ro1 2ue1rog 2ue2o1 2ue2og 2ue3g 2ue4 2uef1o1 2uef1og 2uef2g 2uef3 2uefr1g 2uefr2 2uefro1 2uefrog 3e1r1g 3e1r2 3e1ro1 3e1rog 3e2o1 3e2og 3e3g 3e4 3ef1o1 3ef1og 3ef2g 3ef3 3efr1g 3efr2 3efro1 3efrog 4f1o1 4f1og 4f2g 4f3 4fr1g 4fr2 4fro1 4frog 5r1g 5r2 5ro1 5rog 6o1 6og 7g 8 b1u1f1o1 b1u1f1og b1u1f2g b1u1f3 b1u1fr1g b1u1fr2 b1u1fro1 b1u1frog b1u2r1g b1u2r2 b1u2ro1 b1u2rog b1u3o1 b1u3og b1u4g b1u5 b1ue1r1g b1ue1r2 b1ue1ro1 b1ue1rog b1ue2o1 b1ue2og b1ue3g b1ue4 b1uef1o1 b1uef1og b1uef2g b1uef3 b1uefr1g b1uefr2 b1uefro1 b1uefrog b2e1r1g b2e1r2 b2e1ro1 b2e1rog b2e2o1 b2e2og b2e3g b2e4 b2ef1o1 b2ef1og b2ef2g b2ef3 b2efr1g b2efr2 b2efro1 b2efrog b3f1o1 b3f1og b3f2g b3f3 b3fr1g b3fr2 b3fro1 b3frog b4r1g b4r2 b4ro1 b4rog b5o1 b5og b6g b7 bl1e1r1g bl1e1r2 bl1e1ro1 bl1e1rog bl1e2o1 bl1e2og bl1e3g bl1e4 bl1ef1o1 bl1ef1og bl1ef2g bl1ef3 bl1efr1g bl1efr2 bl1efro1 bl1efrog bl2f1o1 bl2f1og bl2f2g bl2f3 bl2fr1g bl2fr2 bl2fro1 bl2frog bl3r1g bl3r2 bl3ro1 bl3rog bl4o1 bl4og bl5g bl6 blu1f1o1 blu1f1og blu1f2g blu1f3 blu1fr1g blu1fr2 blu1fro1 blu1frog blu2r1g blu2r2 blu2ro1 blu2rog blu3o1 blu3og blu4g blu5 blue1r1g blue1r2 blue1ro1 blue1rog blue2o1 blue2og blue3g blue4 bluef1o1 bluef1og bluef2g bluef3 bluefr1g bluefr2 bluefro1 bluefrog]
    fmt.Println(generateAbbreviations("leetcode")) // [1e1t1o1e 1e1t1o2 1e1t1od1 1e1t1ode 1e1t2d1 1e1t2de 1e1t3e 1e1t4 1e1tc1d1 1e1tc1de 1e1tc2e 1e1tc3 1e1tco1e 1e1tco2 1e1tcod1 1e1tcode 1e2c1d1 1e2c1de 1e2c2e 1e2c3 1e2co1e 1e2co2 1e2cod1 1e2code 1e3o1e 1e3o2 1e3od1 1e3ode 1e4d1 1e4de 1e5e 1e6 1ee1c1d1 1ee1c1de 1ee1c2e 1ee1c3 1ee1co1e 1ee1co2 1ee1cod1 1ee1code 1ee2o1e 1ee2o2 1ee2od1 1ee2ode 1ee3d1 1ee3de 1ee4e 1ee5 1eet1o1e 1eet1o2 1eet1od1 1eet1ode 1eet2d1 1eet2de 1eet3e 1eet4 1eetc1d1 1eetc1de 1eetc2e 1eetc3 1eetco1e 1eetco2 1eetcod1 1eetcode 2e1c1d1 2e1c1de 2e1c2e 2e1c3 2e1co1e 2e1co2 2e1cod1 2e1code 2e2o1e 2e2o2 2e2od1 2e2ode 2e3d1 2e3de 2e4e 2e5 2et1o1e 2et1o2 2et1od1 2et1ode 2et2d1 2et2de 2et3e 2et4 2etc1d1 2etc1de 2etc2e 2etc3 2etco1e 2etco2 2etcod1 2etcode 3t1o1e 3t1o2 3t1od1 3t1ode 3t2d1 3t2de 3t3e 3t4 3tc1d1 3tc1de 3tc2e 3tc3 3tco1e 3tco2 3tcod1 3tcode 4c1d1 4c1de 4c2e 4c3 4co1e 4co2 4cod1 4code 5o1e 5o2 5od1 5ode 6d1 6de 7e 8 l1e1c1d1 l1e1c1de l1e1c2e l1e1c3 l1e1co1e l1e1co2 l1e1cod1 l1e1code l1e2o1e l1e2o2 l1e2od1 l1e2ode l1e3d1 l1e3de l1e4e l1e5 l1et1o1e l1et1o2 l1et1od1 l1et1ode l1et2d1 l1et2de l1et3e l1et4 l1etc1d1 l1etc1de l1etc2e l1etc3 l1etco1e l1etco2 l1etcod1 l1etcode l2t1o1e l2t1o2 l2t1od1 l2t1ode l2t2d1 l2t2de l2t3e l2t4 l2tc1d1 l2tc1de l2tc2e l2tc3 l2tco1e l2tco2 l2tcod1 l2tcode l3c1d1 l3c1de l3c2e l3c3 l3co1e l3co2 l3cod1 l3code l4o1e l4o2 l4od1 l4ode l5d1 l5de l6e l7 le1t1o1e le1t1o2 le1t1od1 le1t1ode le1t2d1 le1t2de le1t3e le1t4 le1tc1d1 le1tc1de le1tc2e le1tc3 le1tco1e le1tco2 le1tcod1 le1tcode le2c1d1 le2c1de le2c2e le2c3 le2co1e le2co2 le2cod1 le2code le3o1e le3o2 le3od1 le3ode le4d1 le4de le5e le6 lee1c1d1 lee1c1de lee1c2e lee1c3 lee1co1e lee1co2 lee1cod1 lee1code lee2o1e lee2o2 lee2od1 lee2ode lee3d1 lee3de lee4e lee5 leet1o1e leet1o2 leet1od1 leet1ode leet2d1 leet2de leet3e leet4 leetc1d1 leetc1de leetc2e leetc3 leetco1e leetco2 leetcod1 leetcode]
    fmt.Println(generateAbbreviations("freewu")) // [1r1e1u 1r1e2 1r1ew1 1r1ewu 1r2w1 1r2wu 1r3u 1r4 1re1w1 1re1wu 1re2u 1re3 1ree1u 1ree2 1reew1 1reewu 2e1w1 2e1wu 2e2u 2e3 2ee1u 2ee2 2eew1 2eewu 3e1u 3e2 3ew1 3ewu 4w1 4wu 5u 6 f1e1w1 f1e1wu f1e2u f1e3 f1ee1u f1ee2 f1eew1 f1eewu f2e1u f2e2 f2ew1 f2ewu f3w1 f3wu f4u f5 fr1e1u fr1e2 fr1ew1 fr1ewu fr2w1 fr2wu fr3u fr4 fre1w1 fre1wu fre2u fre3 free1u free2 freew1 freewu]
}