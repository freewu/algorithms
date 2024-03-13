-- 608. Tree Node
--Table: Tree
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | id          | int  |
-- | p_id        | int  |
-- +-------------+------+
-- id is the primary key column for this table.
-- Each row of this table contains information about the id of a node and the id of its parent node in a tree.
-- The given structure is always a valid tree.

-- Each node in the tree can be one of three types:
--      "Leaf": if the node is a leaf node.
--      "Root": if the node is the root of the tree.
--      "Inner": If the node is neither a leaf node nor a root node.

-- Write an SQL query to report the type of each node in the tree.
-- Return the result table ordered by id in ascending order.
-- The query result format is in the following example.
--
-- Example 1:
-- Input:
-- Tree table:
-- +----+------+
-- | id | p_id |
-- +----+------+
-- | 1  | null |
-- | 2  | 1    |
-- | 3  | 1    |
-- | 4  | 2    |
-- | 5  | 2    |
-- +----+------+
-- Output:
-- +----+-------+
-- | id | type  |
-- +----+-------+
-- | 1  | Root  |
-- | 2  | Inner |
-- | 3  | Leaf  |
-- | 4  | Leaf  |
-- | 5  | Leaf  |
-- +----+-------+
-- Explanation:
-- Node 1 is the root node because its parent node is null and it has child nodes 2 and 3.
-- Node 2 is an inner node because it has parent node 1 and child node 4 and 5.
-- Nodes 3, 4, and 5 are leaf nodes because they have parent nodes and they do not have child nodes.

-- Example 2:
-- Input:
-- Tree table:
-- +----+------+
-- | id | p_id |
-- +----+------+
-- | 1  | null |
-- +----+------+
-- Output:
-- +----+-------+
-- | id | type  |
-- +----+-------+
-- | 1  | Root  |
-- +----+-------+
-- Explanation: If there is only one node on the tree, you only need to output its root attributes.

-- Create table If Not Exists Tree (id int, p_id int)
-- Truncate table Tree
-- insert into Tree (id, p_id) values ('1', 'None')
-- insert into Tree (id, p_id) values ('2', '1')
-- insert into Tree (id, p_id) values ('3', '1')
-- insert into Tree (id, p_id) values ('4', '2')
-- insert into Tree (id, p_id) values ('5', '2')

-- CASE THEN
SELECT
    id,
    CASE
        WHEN tree.p_id IS NULL THEN 'Root' -- "Root": if the node is the root of the tree.
        WHEN tree.id IN (SELECT s.p_id FROM tree AS s) THEN 'Inner' -- "Inner": If the node is neither a leaf node nor a root node.
        ELSE 'Leaf' -- "Leaf": if the node is a leaf node.
    END AS `type`
FROM
    tree
ORDER BY 
    id

-- IF
SELECT
    id,
    IF(ISNULL(p_id),
        'Root', -- p_id 为 NULL 没有父节点 就是 Root
        IF(id IN (SELECT p_id FROM tree), 'Inner','Leaf')) AS `type` -- 如果不存在 p_id 中 就是 Leaf
FROM
    tree
ORDER BY 
    id