-- 3054. Binary Tree Nodes
-- Table: Tree
-- +-------------+------+ 
-- | Column Name | Type | 
-- +-------------+------+ 
-- | N           | int  | 
-- | P           | int  |
-- +-------------+------+
-- N is the column of unique values for this table.
-- Each row includes N and P, where N represents the value of a node in Binary Tree, and P is the parent of N.
-- Write a solution to find the node type of the Binary Tree. Output one of the following for each node:
--     Root: if the node is the root node.
--     Leaf: if the node is the leaf node.
--     Inner: if the node is neither root nor leaf node.

-- Return the result table ordered by node value in ascending order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Tree table:
-- +---+------+
-- | N | P    | 
-- +---+------+
-- | 1 | 2    |
-- | 3 | 2    | 
-- | 6 | 8    | 
-- | 9 | 8    | 
-- | 2 | 5    | 
-- | 8 | 5    | 
-- | 5 | null | 
-- +---+------+
-- Output: 
-- +---+-------+
-- | N | Type  | 
-- +---+-------+
-- | 1 | Leaf  | 
-- | 2 | Inner |
-- | 3 | Leaf  |
-- | 5 | Root  |
-- | 6 | Leaf  |
-- | 8 | Inner |
-- | 9 | Leaf  |    
-- +---+-------+
-- Explanation: 
-- - Node 5 is the root node since it has no parent node.
-- - Nodes 1, 3, 6, and 8 are leaf nodes because they don't have any child nodes.
-- - Nodes 2, 4, and 7 are inner nodes as they serve as parents to some of the nodes in the structure.

-- Create table If Not Exists Tree (N int,P int)
-- Truncate table Tree
-- insert into Tree (N, P) values ('1', '2')
-- insert into Tree (N, P) values ('3', '2')
-- insert into Tree (N, P) values ('6', '8')
-- insert into Tree (N, P) values ('9', '8')
-- insert into Tree (N, P) values ('2', '5')
-- insert into Tree (N, P) values ('8', '5')
-- insert into Tree (N, P) values ('5', 'None')

-- Write your MySQL query statement below
SELECT 
    t.N,
    CASE 
        WHEN COUNT(p.N) = 0 THEN "Root" -- 没有父类为 root
        WHEN COUNT(c.N) = 0 THEN "Leaf" -- 没有儿子的为 leaf
        ELSE "Inner" -- 有父有子的
    END AS Type
FROM
    Tree AS t
LEFT JOIN 
    Tree AS p -- parent
ON 
    t.P = p.N
LEFT JOIN 
    Tree AS c -- child
ON 
    t.N = c.P
GROUP BY
    t.N
ORDER BY 
    t.N 

SELECT 
    t.N,
    CASE 
        WHEN t.P IS NULL THEN "Root" -- 没有父类为 root
        WHEN COUNT(c.N) = 0 THEN "Leaf" -- 没有儿子的为 leaf
        ELSE "Inner" -- 有父有子的
    END AS Type
FROM
    Tree AS t
LEFT JOIN 
    Tree AS c -- child
ON 
    t.N = c.P
GROUP BY
    t.N
ORDER BY 
    t.N 

WITH t AS (
    SELECT 
        a.*, 
        b.N as S
    FROM 
        tree AS a 
    LEFT JOIN 
        tree AS b
    ON 
        b.P = a.N
)
SELECT 
    DISTINCT N,
    CASE 
        WHEN P IS Null THEN 'Root'
        WHEN S IS Null THEN 'Leaf'
        ELSE 'Inner' 
    END AS Type
FROM 
    t
ORDER BY 
    t.N 