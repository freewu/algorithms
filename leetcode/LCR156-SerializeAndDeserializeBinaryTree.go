package main

// LCR 156. 序列化与反序列化二叉树
// 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

// 请设计一个算法来实现二叉树的序列化与反序列化。
// 这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。
// 你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/09/15/serdeser.jpg" />
// 输入：root = [1,2,3,null,null,4,5]
// 输出：[1,2,3,null,null,4,5]

// 示例 2：
// 输入：root = []
// 输出：[]

// 示例 3：
// 输入：root = [1]
// 输出：[1]

// 示例 4：
// 输入：root = [1,2]
// 输出：[1,2]

// 提示：
//     树中结点数在范围 [0, 10^4] 内
//     -1000 <= Node.val <= 1000

import "fmt"
import "bytes"
import "strconv"
import "strings"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {

}
 
func Constructor() Codec {
    return Codec{}
}

// dfs
func (this *Codec) serialize(root *TreeNode) string {
    var buffer bytes.Buffer
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            buffer.WriteString("#,")
        } else {
            buffer.WriteString(strconv.Itoa(node.Val))
            buffer.WriteString(",")
            dfs(node.Left)
            dfs(node.Right)
        }
    }
    dfs(root)
    return buffer.String()
}

// dfs
func (this *Codec) deserialize(data string) *TreeNode {    
    tokens := strings.Split(data, ",")
    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        token := tokens[0]
        tokens = tokens[1:]
        if token == "#" { return nil }
        val, _ := strconv.Atoi(token)
        return &TreeNode{val, dfs(), dfs() }
    }
    return dfs()
}
 
 /**
  * Your Codec object will be instantiated and called as such:
  * ser := Constructor();
  * deser := Constructor();
  * data := ser.serialize(root);
  * ans := deser.deserialize(data);
  */
// bfs
type Codec1 struct {
}

func Constructor1() Codec1 {
	return Codec1{}
}

// Serializes a tree to a single string.
func (this *Codec1) serialize(root *TreeNode) string {
    res := ""
    if root == nil {
        return res
    }
    queue := []*TreeNode{}
    queue = append(queue, root)
    for len(queue) > 0 {
        temp := queue[0]
        queue = queue[1:]
        if temp == nil {
            res += "#,"
        } else {
            res += strconv.Itoa(temp.Val) + ","
        }
        if temp != nil{
            queue = append(queue, temp.Left)
            queue = append(queue, temp.Right)
        }
    }
    return res
}

// Deserializes your encoded data to tree.
func (this *Codec1) deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    nodes := strings.Split(data, ",")
    rootVal, _ := strconv.Atoi(nodes[0])
    root, queue, index := &TreeNode{Val: rootVal}, []*TreeNode{}, 1
    queue = append(queue, root)

    for len(queue) > 0 {
        temp := queue[0]
        queue = queue[1:]
        if nodes[index] != "#" {
            val, _ := strconv.Atoi(nodes[index])
            temp.Left = &TreeNode{Val: val}
            queue = append(queue, temp.Left)
        }
        index++
        if nodes[index] != "#" {
            val, _ := strconv.Atoi(nodes[index])
            temp.Right = &TreeNode{Val: val}
            queue = append(queue, temp.Right)
        }
        index++
    }
    return root
}

// class Codec {
//     struct Node {
//         int val, left, right;
//         Node(int val, int left, int right) : val(val), left(left), right(right) {}
//     };
// public:
//     // Encodes a tree to a single string.
//     string serialize(TreeNode* root) {
//         vector<Node> nodes;
//         function<int(TreeNode*)> postOrder = [&](TreeNode *p) {
//             if (!p) return -1;
//             nodes.emplace_back(p->val, postOrder(p->left), postOrder(p->right));
//             return int(nodes.size()) - 1;
//         };
//         postOrder(root);
//         int n = nodes.size();
//         string s(sizeof(int) + n * sizeof(Node), ' ');
//         char *buffer = s.data();
//         memcpy(buffer, (char*)(&n), sizeof(int));
//         if (n) memcpy(buffer + sizeof(int), (char*)(&nodes[0]), n * sizeof(Node));
//         return s;
//     }

//     // Decodes your encoded data to tree.
//     TreeNode* deserialize(string data) {
//         char *buffer = data.data();
//         int n = *((int*)buffer);
//         if (!n) return nullptr;
//         Node *nodes = (Node*)(buffer + sizeof(int));
//         TreeNode *treeNodes[n];
//         for (int i = 0; i < n; i++) {
//             treeNodes[i] = new TreeNode(nodes[i].val);
//         }
//         for (int i = 0; i < n; i++) {
//             if (nodes[i].left != -1) treeNodes[i]->left = treeNodes[nodes[i].left];
//             if (nodes[i].right != -1) treeNodes[i]->right = treeNodes[nodes[i].right];
//         }
//         return treeNodes[n-1];
//     }
// };

func main() {
    tree1 := &TreeNode {
        1,
        &TreeNode{2, nil, nil},
        &TreeNode {
            3,
            &TreeNode{4, nil, nil},
            &TreeNode{5, nil, nil},
        },
    }
    obj := Constructor()
    data := obj.serialize(tree1)
    fmt.Println(data)
    t := obj.deserialize(data)
    fmt.Println(t.Val)
    fmt.Println(t.Left.Val)
    fmt.Println(t.Right.Val)

    obj1 := Constructor1()
    data1 := obj1.serialize(tree1)
    fmt.Println(data1)
    t1 := obj1.deserialize(data1)
    fmt.Println(t1.Val)
    fmt.Println(t1.Left.Val)
    fmt.Println(t1.Right.Val)
}