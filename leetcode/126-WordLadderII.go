package main

import "fmt"
import "math"

// 126. Word Ladder II
// A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
// 		Every adjacent pair of words differs by a single letter.
// 		Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
// 		sk == endWord

// Given two words, beginWord and endWord, and a dictionary wordList, return all the shortest transformation sequences from beginWord to endWord, or an empty list if no such sequence exists. Each sequence should be returned as a list of the words [beginWord, s1, s2, ..., sk].
 
// Example 1:
// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
// Output: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
// Explanation: There are 2 shortest transformation sequences:
// "hit" -> "hot" -> "dot" -> "dog" -> "cog"
// "hit" -> "hot" -> "lot" -> "log" -> "cog"


// Example 2:
// Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
// Output: []
// Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
 

// Constraints:

// 		1 <= beginWord.length <= 5
// 		endWord.length == beginWord.length
// 		1 <= wordList.length <= 500
// 		wordList[i].length == beginWord.length
// 		beginWord, endWord, and wordList[i] consist of lowercase English letters.
// 		beginWord != endWord
// 		All the words in wordList are unique.
// 		The sum of all shortest transformation sequences does not exceed 10^5.

// Time Limit Exceeded 执行时间太长了
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	result, wordMap := make([][]string, 0), make(map[string]bool)
	// 把单词数组转成 map
	for _, w := range wordList {
		wordMap[w] = true
	}
	// 如果要找的词不在单词列表中 直接返回 false
	if !wordMap[endWord] {
		return result
	}
	// 生成一个多维数组( 步骤结果有多个 )
	// create a queue, track the path
	queue := make([][]string, 0)
	queue = append(queue, []string{beginWord})
	// queueLen is used to track how many slices in queue are in the same level
	// if found a result, I still need to finish checking current level cause I need to return all possible paths
	queueLen := 1
	// use to track strings that this level has visited
	// when queueLen == 0, remove levelMap keys in wordMap
	levelMap := make(map[string]bool)
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastWord := path[len(path)-1]
		for i := 0; i < len(lastWord); i++ {
			for c := 'a'; c <= 'z'; c++ {
				nextWord := lastWord[:i] + string(c) + lastWord[i+1:]
				if nextWord == endWord {
					path = append(path, endWord)
					result = append(result, path)
					continue
				}
				if wordMap[nextWord] {
					// different from word ladder, don't remove the word from wordMap immediately
					// same level could reuse the key.
					// delete from wordMap only when currently level is done.
					levelMap[nextWord] = true
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, nextWord)
					queue = append(queue, newPath)
				}
			}
		}
		queueLen--
		// if queueLen is 0, means finish traversing current level. if result is not empty, return result
		if queueLen == 0 {
			if len(result) > 0 {
				return result
			}
			for k := range levelMap {
				delete(wordMap, k)
			}
			// clear levelMap
			levelMap = make(map[string]bool)
			queueLen = len(queue)
		}
	}
	return result
}

// Memory Limit Exceeded
func findLadders1(beginWord string, endWord string, wordList []string) [][]string {
	n := len(wordList)
	ids := map[string]int{}
	// 把单词数组转成 map
	for i, s := range wordList { 
		ids[s] = i
	}
	ret := [][]string{}
	// 如果要找的词不在单词列表中 直接返回 false
	if _, ok := ids[endWord]; !ok {  
		return ret
	}
	// beginWord在不在wordList里面，不在就给加上
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = n
		n++
	}

	cost := make([]int, n)             // 生成一个cost slice，标识从beginWord通过蛋疼的只能改一个字母的变化，
									   // 变到这个单词的时候最少的变化次数
	edge := make([][]int, n)           // 哎呀重点来了，edge是一个[][]int slice，长度是n，里面每一个member，
									   // 都是一个[]int slice，表示当前wordList所在位置的word，经过一次蛋疼变换
									   // 都能变换到哪些位置的word上去，说白了整完这个edge，从一个指定word出发，不就
									   // 能建立这个word的下一层的所有节点么
	for i := range edge {              // 开始迭代wordList里面的每一个word，每次选一个word出来
		for j := i + 1; j < n; j++ {   // 然后从这个word的下一个word开始迭代，再选一个word出来
			for b := 0; b < len(beginWord); b++ {   // 对这两个word按照每个字节进行对比
				if wordList[i][b] != wordList[j][b] {   // 假设这两个word在某一个字节处，不一样了
					if wordList[i][b+1:] == wordList[j][b+1:] { // 但是剩下部分却一样，这说明啥？这说明这两个word只有这个字节不一样啊
						edge[i] = append(edge[i], j)    // 那还犹豫啥，直接把这两个word对应位置的edge slice加上彼此
						edge[j] = append(edge[j], i)    // 这样给定这两个word的任何一个word，不就知道蛋疼变换能变成啥了么
					}
					break                               // 无论剩下部分一不一样，这两个词就比对完了，break掉这个循环
				}
			}
		}
		cost[i] = math.MaxInt64         // 咋的都是迭代，不如顺手把cost也都给初始化了，假设从beginWord到所有word，步骤都是很长很长
	}
	cost[ids[beginWord]] = 0            // beginWord到beginWord，最短长度肯定是0嘛

	q := [][]int{{ids[beginWord]}}      // BFS标准套路已经厌烦了，就是把beginWord放到queue里面然后开始
									    // 但是此时此刻问题来了，到底咋样才能知道搜索到一个endWord的时候，这个搜索的path路径都
									    // 包含了啥单词？ 所以其实这里顺手把能代表path的一个[]int slice扔进queue里面，作为
										// queue的基本类型，然后每次往queue里面加新成员的时候，把新成员append到path里面，不是
										// 挺香的一个小技巧么？
	for i := 0; i < len(q); i++ {       // 废话不多说，虽然已经说了很多，立刻开始BFS标准套路之迭代循环到queue里面已经没啥可迭代为止
		cur := q[i]                     // BFS标准套路之从queue里面取当前node出来
		curLast := cur[len(cur)-1]	    // 看看当前node代表的path的最后一个访问到的word是啥
		if wordList[curLast] == endWord {   // 唉呀妈呀，这个word要是endWord，
			p := []string{}                 // 那就赶紧把这个path转换成[]string，然后放到返回值里面吧
			for _, id := range cur {
				p = append(p, wordList[id])
			}
			ret = append(ret, p)
			continue                        // 那必须是continue这个循环啊，不需要再往下进行下一层的搜索了
		}
		for _, to := range edge[curLast] {  // 行吧，这个word不是endWord，那让我们把跟这个word能通过蛋疼变换
											// 勾搭到一起的word都给找出来，然后append进path，然后放到queue里面
			if cost[curLast]+1 <= cost[to] {  // 本篇最喵喵喵的部分来了：啥样的word能被放进下一层的queue？
										      // 比如这个word和已经在path里面（BFS搜索过的）word重复了，那说明啥？
											  // 说明这个word的cost在很久之前就已经被set过了，但是如果当前搜索路径
											  // 能让这个cost变小，那不是很香么，那就把更新这个word的cost呀
											  // 当然要是这个word压根还没人摸过，那就是第一次搜索到这个word，那也顺手
											  // 把这个word的cost给设置了。
										      // 总之，要是从当前path的last word（听着瘆得慌，毕竟last word这么命名有
										      // 歧义）走到新的word，cost值超过了之前搜索的路径，那就没必要再谈了
											  // 因为这一定会导致未来就算搜索到endWord的cost，不是最短。并且也有搜索
											  // 成环的死循环风险。
				cost[to] = cost[curLast] + 1
				path := make([]int, len(cur)+1)  // 好了废话太多了，复制path吧，最后一位加上这个toWord
				copy(path, cur)
				path[len(path)-1] = to
				q = append(q, path)           // 然后BFS标准套路之我要把这个节点捅进queue里面就不管了。
			}
		}
	}
	return ret
}

//bfs+dfs(如果是双向bfs，效果会更好)
func findLadders2(beginWord string, endWord string, wordList []string) [][]string {
    //字典表（将wordList中的单词放入hash表中，方便查找）
    dict:=make(map[string]bool,0)
    for _,v:=range wordList{
        dict[v]=true
    }
    //如果endWord不在hash表中，表示不存在转换列表，返回空集合
    if !dict[endWord]{
        return [][]string{}
    }
    //将第一个单词放入hash表中，方便实现邻接表（构建图）
    dict[beginWord]=true
    //构建邻接表
    graph:=make(map[string][]string,0)
    //执行bfs搜索，找到每个点到endWord的距离
    distance:=bfs(endWord,dict,graph)
    res:=make([][]string,0)//保存结果
    //执行dfs操作
    dfs(beginWord,endWord,&res,[]string{},distance,graph)
    return res
}

//回溯实现方式一：（个人偏好这个，更符合模板）
func dfs(beginWord string,endWord string,res *[][]string,path []string,distance map[string]int,graph map[string][]string){
    //出递归条件
    if beginWord==endWord{
        path=append(path,beginWord) //加入endWord节点
        tmp:=make([]string,len(path))
        copy(tmp,path)
        (*res)=append((*res),tmp)
        path=path[:len(path)-1] //移除endWord节点
        return
    }
    //否则遍历图
    for _,v:=range graph[beginWord]{
        //遍历图时，朝着距离与终点越来越近的方向进行（当前节点的距离肯定要比下一个距离大1）
        if distance[beginWord]==distance[v]+1{
            path=append(path,beginWord)
            dfs(v,endWord,res,path,distance,graph)
            //回溯（执行完上述的所有时，将其回溯回去）
            path=path[:len(path)-1]
        } 
    }
}
//回溯实现方式二：
// func dfs(beginWord string,endWord string,res *[][]string,path []string,distance map[string]int,graph map[string][]string){
//     path=append(path,beginWord)
//     //出递归条件
//     if beginWord==endWord{
//         tmp:=make([]string,len(path))
//         copy(tmp,path)
//         (*res)=append((*res),tmp)
//         return
//     }
//     //否则遍历图
//     for _,v:=range graph[beginWord]{
//         //遍历图时，朝着距离与终点越来越近的方向进行（当前节点的距离肯定要比下一个距离大1）
//         if distance[beginWord]==distance[v]+1{
//             dfs(v,endWord,res,path,distance,graph)
//         } 
//     }
//     //回溯（执行完上述的所有时，将其回溯回去）
//     path=path[:len(path)-1]
// }


//从终点出发，进行bfs，计算每一个点到达终点的距离
func bfs(endWord string,dict map[string]bool,graph map[string][]string)map[string]int{
    distance:=make(map[string]int,0) //每个点到终点的距离
    queue:=make([]string,0)
    queue=append(queue,endWord)
    distance[endWord]=0 //初始值
    for len(queue)!=0{
        cursize:=len(queue)
        for i:=0;i<cursize;i++{
            word:=queue[0]
            queue=queue[1:]
            //找到和word有一位不同的单词列表
            expansion:=expand(word,dict)
            for _,v:=range expansion{
                //构造邻接表
                //我们是从beginWord到endWord构造邻接表，而bfs是从endWord开始，所以构造时，反过来构造
                //即graph[v]=append(graph[v],word)而不是graph[word]=append(graph[word],v)
                graph[v]=append(graph[v],word)
                //表示没访问过
                if _,ok:=distance[v];!ok{
                    distance[v]=distance[word]+1 //距离加一
                    queue=append(queue,v) //入队列
                }
            }
        }
    }
    return distance
}

//获得邻接点
func expand(word string,dict map[string]bool)[]string{
    expansion:=make([]string,0) //保存word的邻接点
    //从word的每一位开始
    chs:=[]rune(word)
    for i:=0;i<len(word);i++{
        tmp:=chs[i] //保存当前位，方便后序进行复位
        for c:='a';c<='z';c++{
            //如果一样则直接跳过，之所以用tmp，是因为chs[i]在变
            if tmp==c{ 
                continue
            }
            chs[i]=c
            newstr:=string(chs)
            //新单词在dict中不存在，则跳过
            if dict[newstr]{
                expansion=append(expansion,newstr)
            }
        }
        chs[i]=tmp //单词位复位
    }
    return expansion
}

// best solution
func findLadders3(beginWord string, endWord string, wordList []string) [][]string {
	distancemp := make([]int, len(wordList))
	q := make([]int, 0)

	dis := 1
	for pos, word := range wordList {
		if dif(beginWord, word) == 1 {
			distancemp[pos] = dis
			q = append(q, pos)
            if word == endWord {
                return [][]string{[]string{beginWord, endWord}}
            }
		}
	}
	q1 := make([]int, 0)
	found := false

	for len(q) > 0 && !found {
		dis++
		for _, j := range q {
			for i := range wordList {
				if distancemp[i] > 0 {
					continue
				}
				if dif(wordList[j], wordList[i]) == 1 {
					distancemp[i] = dis
					q1 = append(q1, i)
					if wordList[i] == endWord {
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
		q, q1 = q1, q[:0]
	}
    if !found {
        return nil
    }

	ans := [][]string{{endWord}}
	tmp := [][]string{}

	for i := dis - 1; i >= 1; i-- {
		for j := 0; j < len(distancemp); j++ {
			if distancemp[j] == i {
				for _, wordarr := range ans {
					if dif(wordarr[len(wordarr)-1], wordList[j]) == 1 {
						res := make([]string, 0, len(wordarr)+1)
						res = append(res, wordarr...)
						res = append(res, wordList[j])
						tmp = append(tmp, res)
					}
				}
			}
		}
		ans, tmp = tmp, [][]string{}
	}

	for i := 0; i < len(ans); i++ {
		ans[i] = revert(append(ans[i], beginWord))
	}

	return ans
}

func revert(wordarr []string) []string {
	for i, j := 0, len(wordarr)-1; i < j; {
		wordarr[i], wordarr[j] = wordarr[j], wordarr[i]
		i++
		j--
	}
	return wordarr
}

func dif(s1, s2 string) int {
	k := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			k++
			if k > 1 {
				return k
			}
		}
	}
	return k
}

func main() {
	// [[hit hot dot dog cog] [hit hot lot log cog]]
	fmt.Println(findLadders("hit","cog",[]string{"hot","dot","dog","lot","log","cog"}))
	// []
	fmt.Println(findLadders("hit","cog",[]string{"hot","dot","dog","lot","log"}))

	// []
	//fmt.Println(findLadders("aaaaa","ggggg",[]string{"aaaaa","caaaa","cbaaa","daaaa","dbaaa","eaaaa","ebaaa","faaaa","fbaaa","gaaaa","gbaaa","haaaa","hbaaa","iaaaa","ibaaa","jaaaa","jbaaa","kaaaa","kbaaa","laaaa","lbaaa","maaaa","mbaaa","naaaa","nbaaa","oaaaa","obaaa","paaaa","pbaaa","bbaaa","bbcaa","bbcba","bbdaa","bbdba","bbeaa","bbeba","bbfaa","bbfba","bbgaa","bbgba","bbhaa","bbhba","bbiaa","bbiba","bbjaa","bbjba","bbkaa","bbkba","bblaa","bblba","bbmaa","bbmba","bbnaa","bbnba","bboaa","bboba","bbpaa","bbpba","bbbba","abbba","acbba","dbbba","dcbba","ebbba","ecbba","fbbba","fcbba","gbbba","gcbba","hbbba","hcbba","ibbba","icbba","jbbba","jcbba","kbbba","kcbba","lbbba","lcbba","mbbba","mcbba","nbbba","ncbba","obbba","ocbba","pbbba","pcbba","ccbba","ccaba","ccaca","ccdba","ccdca","cceba","cceca","ccfba","ccfca","ccgba","ccgca","cchba","cchca","cciba","ccica","ccjba","ccjca","cckba","cckca","cclba","cclca","ccmba","ccmca","ccnba","ccnca","ccoba","ccoca","ccpba","ccpca","cccca","accca","adcca","bccca","bdcca","eccca","edcca","fccca","fdcca","gccca","gdcca","hccca","hdcca","iccca","idcca","jccca","jdcca","kccca","kdcca","lccca","ldcca","mccca","mdcca","nccca","ndcca","occca","odcca","pccca","pdcca","ddcca","ddaca","ddada","ddbca","ddbda","ddeca","ddeda","ddfca","ddfda","ddgca","ddgda","ddhca","ddhda","ddica","ddida","ddjca","ddjda","ddkca","ddkda","ddlca","ddlda","ddmca","ddmda","ddnca","ddnda","ddoca","ddoda","ddpca","ddpda","dddda","addda","aedda","bddda","bedda","cddda","cedda","fddda","fedda","gddda","gedda","hddda","hedda","iddda","iedda","jddda","jedda","kddda","kedda","lddda","ledda","mddda","medda","nddda","nedda","oddda","oedda","pddda","pedda","eedda","eeada","eeaea","eebda","eebea","eecda","eecea","eefda","eefea","eegda","eegea","eehda","eehea","eeida","eeiea","eejda","eejea","eekda","eekea","eelda","eelea","eemda","eemea","eenda","eenea","eeoda","eeoea","eepda","eepea","eeeea","ggggg","agggg","ahggg","bgggg","bhggg","cgggg","chggg","dgggg","dhggg","egggg","ehggg","fgggg","fhggg","igggg","ihggg","jgggg","jhggg","kgggg","khggg","lgggg","lhggg","mgggg","mhggg","ngggg","nhggg","ogggg","ohggg","pgggg","phggg","hhggg","hhagg","hhahg","hhbgg","hhbhg","hhcgg","hhchg","hhdgg","hhdhg","hhegg","hhehg","hhfgg","hhfhg","hhigg","hhihg","hhjgg","hhjhg","hhkgg","hhkhg","hhlgg","hhlhg","hhmgg","hhmhg","hhngg","hhnhg","hhogg","hhohg","hhpgg","hhphg","hhhhg","ahhhg","aihhg","bhhhg","bihhg","chhhg","cihhg","dhhhg","dihhg","ehhhg","eihhg","fhhhg","fihhg","ghhhg","gihhg","jhhhg","jihhg","khhhg","kihhg","lhhhg","lihhg","mhhhg","mihhg","nhhhg","nihhg","ohhhg","oihhg","phhhg","pihhg","iihhg","iiahg","iiaig","iibhg","iibig","iichg","iicig","iidhg","iidig","iiehg","iieig","iifhg","iifig","iighg","iigig","iijhg","iijig","iikhg","iikig","iilhg","iilig","iimhg","iimig","iinhg","iinig","iiohg","iioig","iiphg","iipig","iiiig","aiiig","ajiig","biiig","bjiig","ciiig","cjiig","diiig","djiig","eiiig","ejiig","fiiig","fjiig","giiig","gjiig","hiiig","hjiig","kiiig","kjiig","liiig","ljiig","miiig","mjiig","niiig","njiig","oiiig","ojiig","piiig","pjiig","jjiig","jjaig","jjajg","jjbig","jjbjg","jjcig","jjcjg","jjdig","jjdjg","jjeig","jjejg","jjfig","jjfjg","jjgig","jjgjg","jjhig","jjhjg","jjkig","jjkjg","jjlig","jjljg","jjmig","jjmjg","jjnig","jjnjg","jjoig","jjojg","jjpig","jjpjg","jjjjg","ajjjg","akjjg","bjjjg","bkjjg","cjjjg","ckjjg","djjjg","dkjjg","ejjjg","ekjjg","fjjjg","fkjjg","gjjjg","gkjjg","hjjjg","hkjjg","ijjjg","ikjjg","ljjjg","lkjjg","mjjjg","mkjjg","njjjg","nkjjg","ojjjg","okjjg","pjjjg","pkjjg","kkjjg","kkajg","kkakg","kkbjg","kkbkg","kkcjg","kkckg","kkdjg","kkdkg","kkejg","kkekg","kkfjg","kkfkg","kkgjg","kkgkg","kkhjg","kkhkg","kkijg","kkikg","kkljg","kklkg","kkmjg","kkmkg","kknjg","kknkg","kkojg","kkokg","kkpjg","kkpkg","kkkkg","ggggx","gggxx","ggxxx","gxxxx","xxxxx","xxxxy","xxxyy","xxyyy","xyyyy","yyyyy","yyyyw","yyyww","yywww","ywwww","wwwww","wwvww","wvvww","vvvww","vvvwz","avvwz","aavwz","aaawz","aaaaz"}))

	// [[hit hot dot dog cog] [hit hot lot log cog]]
	fmt.Println(findLadders1("hit","cog",[]string{"hot","dot","dog","lot","log","cog"}))
	// []
	fmt.Println(findLadders1("hit","cog",[]string{"hot","dot","dog","lot","log"}))

	// []
	//fmt.Println(findLadders1("aaaaa","ggggg",[]string{"aaaaa","caaaa","cbaaa","daaaa","dbaaa","eaaaa","ebaaa","faaaa","fbaaa","gaaaa","gbaaa","haaaa","hbaaa","iaaaa","ibaaa","jaaaa","jbaaa","kaaaa","kbaaa","laaaa","lbaaa","maaaa","mbaaa","naaaa","nbaaa","oaaaa","obaaa","paaaa","pbaaa","bbaaa","bbcaa","bbcba","bbdaa","bbdba","bbeaa","bbeba","bbfaa","bbfba","bbgaa","bbgba","bbhaa","bbhba","bbiaa","bbiba","bbjaa","bbjba","bbkaa","bbkba","bblaa","bblba","bbmaa","bbmba","bbnaa","bbnba","bboaa","bboba","bbpaa","bbpba","bbbba","abbba","acbba","dbbba","dcbba","ebbba","ecbba","fbbba","fcbba","gbbba","gcbba","hbbba","hcbba","ibbba","icbba","jbbba","jcbba","kbbba","kcbba","lbbba","lcbba","mbbba","mcbba","nbbba","ncbba","obbba","ocbba","pbbba","pcbba","ccbba","ccaba","ccaca","ccdba","ccdca","cceba","cceca","ccfba","ccfca","ccgba","ccgca","cchba","cchca","cciba","ccica","ccjba","ccjca","cckba","cckca","cclba","cclca","ccmba","ccmca","ccnba","ccnca","ccoba","ccoca","ccpba","ccpca","cccca","accca","adcca","bccca","bdcca","eccca","edcca","fccca","fdcca","gccca","gdcca","hccca","hdcca","iccca","idcca","jccca","jdcca","kccca","kdcca","lccca","ldcca","mccca","mdcca","nccca","ndcca","occca","odcca","pccca","pdcca","ddcca","ddaca","ddada","ddbca","ddbda","ddeca","ddeda","ddfca","ddfda","ddgca","ddgda","ddhca","ddhda","ddica","ddida","ddjca","ddjda","ddkca","ddkda","ddlca","ddlda","ddmca","ddmda","ddnca","ddnda","ddoca","ddoda","ddpca","ddpda","dddda","addda","aedda","bddda","bedda","cddda","cedda","fddda","fedda","gddda","gedda","hddda","hedda","iddda","iedda","jddda","jedda","kddda","kedda","lddda","ledda","mddda","medda","nddda","nedda","oddda","oedda","pddda","pedda","eedda","eeada","eeaea","eebda","eebea","eecda","eecea","eefda","eefea","eegda","eegea","eehda","eehea","eeida","eeiea","eejda","eejea","eekda","eekea","eelda","eelea","eemda","eemea","eenda","eenea","eeoda","eeoea","eepda","eepea","eeeea","ggggg","agggg","ahggg","bgggg","bhggg","cgggg","chggg","dgggg","dhggg","egggg","ehggg","fgggg","fhggg","igggg","ihggg","jgggg","jhggg","kgggg","khggg","lgggg","lhggg","mgggg","mhggg","ngggg","nhggg","ogggg","ohggg","pgggg","phggg","hhggg","hhagg","hhahg","hhbgg","hhbhg","hhcgg","hhchg","hhdgg","hhdhg","hhegg","hhehg","hhfgg","hhfhg","hhigg","hhihg","hhjgg","hhjhg","hhkgg","hhkhg","hhlgg","hhlhg","hhmgg","hhmhg","hhngg","hhnhg","hhogg","hhohg","hhpgg","hhphg","hhhhg","ahhhg","aihhg","bhhhg","bihhg","chhhg","cihhg","dhhhg","dihhg","ehhhg","eihhg","fhhhg","fihhg","ghhhg","gihhg","jhhhg","jihhg","khhhg","kihhg","lhhhg","lihhg","mhhhg","mihhg","nhhhg","nihhg","ohhhg","oihhg","phhhg","pihhg","iihhg","iiahg","iiaig","iibhg","iibig","iichg","iicig","iidhg","iidig","iiehg","iieig","iifhg","iifig","iighg","iigig","iijhg","iijig","iikhg","iikig","iilhg","iilig","iimhg","iimig","iinhg","iinig","iiohg","iioig","iiphg","iipig","iiiig","aiiig","ajiig","biiig","bjiig","ciiig","cjiig","diiig","djiig","eiiig","ejiig","fiiig","fjiig","giiig","gjiig","hiiig","hjiig","kiiig","kjiig","liiig","ljiig","miiig","mjiig","niiig","njiig","oiiig","ojiig","piiig","pjiig","jjiig","jjaig","jjajg","jjbig","jjbjg","jjcig","jjcjg","jjdig","jjdjg","jjeig","jjejg","jjfig","jjfjg","jjgig","jjgjg","jjhig","jjhjg","jjkig","jjkjg","jjlig","jjljg","jjmig","jjmjg","jjnig","jjnjg","jjoig","jjojg","jjpig","jjpjg","jjjjg","ajjjg","akjjg","bjjjg","bkjjg","cjjjg","ckjjg","djjjg","dkjjg","ejjjg","ekjjg","fjjjg","fkjjg","gjjjg","gkjjg","hjjjg","hkjjg","ijjjg","ikjjg","ljjjg","lkjjg","mjjjg","mkjjg","njjjg","nkjjg","ojjjg","okjjg","pjjjg","pkjjg","kkjjg","kkajg","kkakg","kkbjg","kkbkg","kkcjg","kkckg","kkdjg","kkdkg","kkejg","kkekg","kkfjg","kkfkg","kkgjg","kkgkg","kkhjg","kkhkg","kkijg","kkikg","kkljg","kklkg","kkmjg","kkmkg","kknjg","kknkg","kkojg","kkokg","kkpjg","kkpkg","kkkkg","ggggx","gggxx","ggxxx","gxxxx","xxxxx","xxxxy","xxxyy","xxyyy","xyyyy","yyyyy","yyyyw","yyyww","yywww","ywwww","wwwww","wwvww","wvvww","vvvww","vvvwz","avvwz","aavwz","aaawz","aaaaz"}))
	
	// [[hit hot dot dog cog] [hit hot lot log cog]]
	fmt.Println(findLadders2("hit","cog",[]string{"hot","dot","dog","lot","log","cog"}))
	// []
	fmt.Println(findLadders2("hit","cog",[]string{"hot","dot","dog","lot","log"}))
	// [[aaaaa aaaaz aaawz aavwz avvwz vvvwz vvvww wvvww wwvww wwwww ywwww yywww yyyww yyyyw yyyyy xyyyy xxyyy xxxyy xxxxy xxxxx gxxxx ggxxx gggxx ggggx ggggg]]
	fmt.Println(findLadders2("aaaaa","ggggg",[]string{"aaaaa","caaaa","cbaaa","daaaa","dbaaa","eaaaa","ebaaa","faaaa","fbaaa","gaaaa","gbaaa","haaaa","hbaaa","iaaaa","ibaaa","jaaaa","jbaaa","kaaaa","kbaaa","laaaa","lbaaa","maaaa","mbaaa","naaaa","nbaaa","oaaaa","obaaa","paaaa","pbaaa","bbaaa","bbcaa","bbcba","bbdaa","bbdba","bbeaa","bbeba","bbfaa","bbfba","bbgaa","bbgba","bbhaa","bbhba","bbiaa","bbiba","bbjaa","bbjba","bbkaa","bbkba","bblaa","bblba","bbmaa","bbmba","bbnaa","bbnba","bboaa","bboba","bbpaa","bbpba","bbbba","abbba","acbba","dbbba","dcbba","ebbba","ecbba","fbbba","fcbba","gbbba","gcbba","hbbba","hcbba","ibbba","icbba","jbbba","jcbba","kbbba","kcbba","lbbba","lcbba","mbbba","mcbba","nbbba","ncbba","obbba","ocbba","pbbba","pcbba","ccbba","ccaba","ccaca","ccdba","ccdca","cceba","cceca","ccfba","ccfca","ccgba","ccgca","cchba","cchca","cciba","ccica","ccjba","ccjca","cckba","cckca","cclba","cclca","ccmba","ccmca","ccnba","ccnca","ccoba","ccoca","ccpba","ccpca","cccca","accca","adcca","bccca","bdcca","eccca","edcca","fccca","fdcca","gccca","gdcca","hccca","hdcca","iccca","idcca","jccca","jdcca","kccca","kdcca","lccca","ldcca","mccca","mdcca","nccca","ndcca","occca","odcca","pccca","pdcca","ddcca","ddaca","ddada","ddbca","ddbda","ddeca","ddeda","ddfca","ddfda","ddgca","ddgda","ddhca","ddhda","ddica","ddida","ddjca","ddjda","ddkca","ddkda","ddlca","ddlda","ddmca","ddmda","ddnca","ddnda","ddoca","ddoda","ddpca","ddpda","dddda","addda","aedda","bddda","bedda","cddda","cedda","fddda","fedda","gddda","gedda","hddda","hedda","iddda","iedda","jddda","jedda","kddda","kedda","lddda","ledda","mddda","medda","nddda","nedda","oddda","oedda","pddda","pedda","eedda","eeada","eeaea","eebda","eebea","eecda","eecea","eefda","eefea","eegda","eegea","eehda","eehea","eeida","eeiea","eejda","eejea","eekda","eekea","eelda","eelea","eemda","eemea","eenda","eenea","eeoda","eeoea","eepda","eepea","eeeea","ggggg","agggg","ahggg","bgggg","bhggg","cgggg","chggg","dgggg","dhggg","egggg","ehggg","fgggg","fhggg","igggg","ihggg","jgggg","jhggg","kgggg","khggg","lgggg","lhggg","mgggg","mhggg","ngggg","nhggg","ogggg","ohggg","pgggg","phggg","hhggg","hhagg","hhahg","hhbgg","hhbhg","hhcgg","hhchg","hhdgg","hhdhg","hhegg","hhehg","hhfgg","hhfhg","hhigg","hhihg","hhjgg","hhjhg","hhkgg","hhkhg","hhlgg","hhlhg","hhmgg","hhmhg","hhngg","hhnhg","hhogg","hhohg","hhpgg","hhphg","hhhhg","ahhhg","aihhg","bhhhg","bihhg","chhhg","cihhg","dhhhg","dihhg","ehhhg","eihhg","fhhhg","fihhg","ghhhg","gihhg","jhhhg","jihhg","khhhg","kihhg","lhhhg","lihhg","mhhhg","mihhg","nhhhg","nihhg","ohhhg","oihhg","phhhg","pihhg","iihhg","iiahg","iiaig","iibhg","iibig","iichg","iicig","iidhg","iidig","iiehg","iieig","iifhg","iifig","iighg","iigig","iijhg","iijig","iikhg","iikig","iilhg","iilig","iimhg","iimig","iinhg","iinig","iiohg","iioig","iiphg","iipig","iiiig","aiiig","ajiig","biiig","bjiig","ciiig","cjiig","diiig","djiig","eiiig","ejiig","fiiig","fjiig","giiig","gjiig","hiiig","hjiig","kiiig","kjiig","liiig","ljiig","miiig","mjiig","niiig","njiig","oiiig","ojiig","piiig","pjiig","jjiig","jjaig","jjajg","jjbig","jjbjg","jjcig","jjcjg","jjdig","jjdjg","jjeig","jjejg","jjfig","jjfjg","jjgig","jjgjg","jjhig","jjhjg","jjkig","jjkjg","jjlig","jjljg","jjmig","jjmjg","jjnig","jjnjg","jjoig","jjojg","jjpig","jjpjg","jjjjg","ajjjg","akjjg","bjjjg","bkjjg","cjjjg","ckjjg","djjjg","dkjjg","ejjjg","ekjjg","fjjjg","fkjjg","gjjjg","gkjjg","hjjjg","hkjjg","ijjjg","ikjjg","ljjjg","lkjjg","mjjjg","mkjjg","njjjg","nkjjg","ojjjg","okjjg","pjjjg","pkjjg","kkjjg","kkajg","kkakg","kkbjg","kkbkg","kkcjg","kkckg","kkdjg","kkdkg","kkejg","kkekg","kkfjg","kkfkg","kkgjg","kkgkg","kkhjg","kkhkg","kkijg","kkikg","kkljg","kklkg","kkmjg","kkmkg","kknjg","kknkg","kkojg","kkokg","kkpjg","kkpkg","kkkkg","ggggx","gggxx","ggxxx","gxxxx","xxxxx","xxxxy","xxxyy","xxyyy","xyyyy","yyyyy","yyyyw","yyyww","yywww","ywwww","wwwww","wwvww","wvvww","vvvww","vvvwz","avvwz","aavwz","aaawz","aaaaz"}))

	// [[hit hot dot dog cog] [hit hot lot log cog]]
	fmt.Println(findLadders3("hit","cog",[]string{"hot","dot","dog","lot","log","cog"}))
	// []
	fmt.Println(findLadders3("hit","cog",[]string{"hot","dot","dog","lot","log"}))
	// [[aaaaa aaaaz aaawz aavwz avvwz vvvwz vvvww wvvww wwvww wwwww ywwww yywww yyyww yyyyw yyyyy xyyyy xxyyy xxxyy xxxxy xxxxx gxxxx ggxxx gggxx ggggx ggggg]]
	fmt.Println(findLadders3("aaaaa","ggggg",[]string{"aaaaa","caaaa","cbaaa","daaaa","dbaaa","eaaaa","ebaaa","faaaa","fbaaa","gaaaa","gbaaa","haaaa","hbaaa","iaaaa","ibaaa","jaaaa","jbaaa","kaaaa","kbaaa","laaaa","lbaaa","maaaa","mbaaa","naaaa","nbaaa","oaaaa","obaaa","paaaa","pbaaa","bbaaa","bbcaa","bbcba","bbdaa","bbdba","bbeaa","bbeba","bbfaa","bbfba","bbgaa","bbgba","bbhaa","bbhba","bbiaa","bbiba","bbjaa","bbjba","bbkaa","bbkba","bblaa","bblba","bbmaa","bbmba","bbnaa","bbnba","bboaa","bboba","bbpaa","bbpba","bbbba","abbba","acbba","dbbba","dcbba","ebbba","ecbba","fbbba","fcbba","gbbba","gcbba","hbbba","hcbba","ibbba","icbba","jbbba","jcbba","kbbba","kcbba","lbbba","lcbba","mbbba","mcbba","nbbba","ncbba","obbba","ocbba","pbbba","pcbba","ccbba","ccaba","ccaca","ccdba","ccdca","cceba","cceca","ccfba","ccfca","ccgba","ccgca","cchba","cchca","cciba","ccica","ccjba","ccjca","cckba","cckca","cclba","cclca","ccmba","ccmca","ccnba","ccnca","ccoba","ccoca","ccpba","ccpca","cccca","accca","adcca","bccca","bdcca","eccca","edcca","fccca","fdcca","gccca","gdcca","hccca","hdcca","iccca","idcca","jccca","jdcca","kccca","kdcca","lccca","ldcca","mccca","mdcca","nccca","ndcca","occca","odcca","pccca","pdcca","ddcca","ddaca","ddada","ddbca","ddbda","ddeca","ddeda","ddfca","ddfda","ddgca","ddgda","ddhca","ddhda","ddica","ddida","ddjca","ddjda","ddkca","ddkda","ddlca","ddlda","ddmca","ddmda","ddnca","ddnda","ddoca","ddoda","ddpca","ddpda","dddda","addda","aedda","bddda","bedda","cddda","cedda","fddda","fedda","gddda","gedda","hddda","hedda","iddda","iedda","jddda","jedda","kddda","kedda","lddda","ledda","mddda","medda","nddda","nedda","oddda","oedda","pddda","pedda","eedda","eeada","eeaea","eebda","eebea","eecda","eecea","eefda","eefea","eegda","eegea","eehda","eehea","eeida","eeiea","eejda","eejea","eekda","eekea","eelda","eelea","eemda","eemea","eenda","eenea","eeoda","eeoea","eepda","eepea","eeeea","ggggg","agggg","ahggg","bgggg","bhggg","cgggg","chggg","dgggg","dhggg","egggg","ehggg","fgggg","fhggg","igggg","ihggg","jgggg","jhggg","kgggg","khggg","lgggg","lhggg","mgggg","mhggg","ngggg","nhggg","ogggg","ohggg","pgggg","phggg","hhggg","hhagg","hhahg","hhbgg","hhbhg","hhcgg","hhchg","hhdgg","hhdhg","hhegg","hhehg","hhfgg","hhfhg","hhigg","hhihg","hhjgg","hhjhg","hhkgg","hhkhg","hhlgg","hhlhg","hhmgg","hhmhg","hhngg","hhnhg","hhogg","hhohg","hhpgg","hhphg","hhhhg","ahhhg","aihhg","bhhhg","bihhg","chhhg","cihhg","dhhhg","dihhg","ehhhg","eihhg","fhhhg","fihhg","ghhhg","gihhg","jhhhg","jihhg","khhhg","kihhg","lhhhg","lihhg","mhhhg","mihhg","nhhhg","nihhg","ohhhg","oihhg","phhhg","pihhg","iihhg","iiahg","iiaig","iibhg","iibig","iichg","iicig","iidhg","iidig","iiehg","iieig","iifhg","iifig","iighg","iigig","iijhg","iijig","iikhg","iikig","iilhg","iilig","iimhg","iimig","iinhg","iinig","iiohg","iioig","iiphg","iipig","iiiig","aiiig","ajiig","biiig","bjiig","ciiig","cjiig","diiig","djiig","eiiig","ejiig","fiiig","fjiig","giiig","gjiig","hiiig","hjiig","kiiig","kjiig","liiig","ljiig","miiig","mjiig","niiig","njiig","oiiig","ojiig","piiig","pjiig","jjiig","jjaig","jjajg","jjbig","jjbjg","jjcig","jjcjg","jjdig","jjdjg","jjeig","jjejg","jjfig","jjfjg","jjgig","jjgjg","jjhig","jjhjg","jjkig","jjkjg","jjlig","jjljg","jjmig","jjmjg","jjnig","jjnjg","jjoig","jjojg","jjpig","jjpjg","jjjjg","ajjjg","akjjg","bjjjg","bkjjg","cjjjg","ckjjg","djjjg","dkjjg","ejjjg","ekjjg","fjjjg","fkjjg","gjjjg","gkjjg","hjjjg","hkjjg","ijjjg","ikjjg","ljjjg","lkjjg","mjjjg","mkjjg","njjjg","nkjjg","ojjjg","okjjg","pjjjg","pkjjg","kkjjg","kkajg","kkakg","kkbjg","kkbkg","kkcjg","kkckg","kkdjg","kkdkg","kkejg","kkekg","kkfjg","kkfkg","kkgjg","kkgkg","kkhjg","kkhkg","kkijg","kkikg","kkljg","kklkg","kkmjg","kkmkg","kknjg","kknkg","kkojg","kkokg","kkpjg","kkpkg","kkkkg","ggggx","gggxx","ggxxx","gxxxx","xxxxx","xxxxy","xxxyy","xxyyy","xyyyy","yyyyy","yyyyw","yyyww","yywww","ywwww","wwwww","wwvww","wvvww","vvvww","vvvwz","avvwz","aavwz","aaawz","aaaaz"}))
}