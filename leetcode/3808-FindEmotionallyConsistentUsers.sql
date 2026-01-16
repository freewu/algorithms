-- 3808. Find Emotionally Consistent Users
-- Table: reactions
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | user_id      | int     |
-- | content_id   | int     |
-- | reaction     | varchar |
-- +--------------+---------+
-- (user_id, content_id) is the primary key (unique value) for this table.
-- Each row represents a reaction given by a user to a piece of content.
-- Write a solution to identify emotionally consistent users based on the following requirements:

-- For each user, count the total number of reactions they have given.
-- Only include users who have reacted to at least 5 different content items.
-- A user is considered emotionally consistent if at least 60% of their reactions are of the same type.
-- Return the result table ordered by reaction_ratio in descending order and then by user_id in ascending order.

-- Note:
--     1. reaction_ratio should be rounded to 2 decimal places
--     2. The result format is in the following example.

-- Example:
-- Input:
-- reactions table:
-- +---------+------------+----------+
-- | user_id | content_id | reaction |
-- +---------+------------+----------+
-- | 1       | 101        | like     |
-- | 1       | 102        | like     |
-- | 1       | 103        | like     |
-- | 1       | 104        | wow      |
-- | 1       | 105        | like     |
-- | 2       | 201        | like     |
-- | 2       | 202        | wow      |
-- | 2       | 203        | sad      |
-- | 2       | 204        | like     |
-- | 2       | 205        | wow      |
-- | 3       | 301        | love     |
-- | 3       | 302        | love     |
-- | 3       | 303        | love     |
-- | 3       | 304        | love     |
-- | 3       | 305        | love     |
-- +---------+------------+----------+
-- Output:
-- +---------+-------------------+----------------+
-- | user_id | dominant_reaction | reaction_ratio |
-- +---------+-------------------+----------------+
-- | 3       | love              | 1.00           |
-- | 1       | like              | 0.80           |
-- +---------+-------------------+----------------+
-- Explanation:
-- User 1:
-- Total reactions = 5
-- like appears 4 times
-- reaction_ratio = 4 / 5 = 0.80
-- Meets the 60% consistency requirement
-- User 2:
-- Total reactions = 5
-- Most frequent reaction appears only 2 times
-- reaction_ratio = 2 / 5 = 0.40
-- Does not meet the consistency requirement
-- User 3:
-- Total reactions = 5
-- 'love' appears 5 times
-- reaction_ratio = 5 / 5 = 1.00
-- Meets the consistency requirement
-- The Results table is ordered by reaction_ratio in descending order, then by user_id in ascending order.

-- Write your MySQL query statement below
WITH t1 AS ( -- 统计每个用户每个 reaction 的数量
    SELECT
        user_id,
        reaction,
        count(*) AS count
    FROM reactions
    GROUP BY user_id, reaction
),
t2 AS ( -- 总 reaction 数量 大于 5个 的用户
    SELECT
        user_id,
        count(*) AS total_reaction
    FROM reactions
    GROUP BY user_id
    HAVING total_reaction >= 5 -- Only include users who have reacted to at least 5 different content items.
),
-- 连接 t2 和 t1
t3 AS (
    SELECT
        t1.user_id,
        t1.reaction,
        t1.count,
        t2.total_reaction,
        -- 编号找每个用户数量最多的reaction
        row_number() OVER (PARTITION BY t1.user_id ORDER BY count DESC) AS rk
    FROM
        t2 
    LEFT JOIN t1 on t2.user_id = t1.user_id
)
-- 按要求整理结果
SELECT
    user_id,
    reaction AS dominant_reaction,
    round(count / total_reaction, 2) AS reaction_ratio
FROM
    t3
WHERE 
    rk = 1 AND count / total_reaction >= 0.6 -- A user is considered emotionally consistent if at least 60% of their reactions are of the same type.
ORDER BY 
    reaction_ratio DESC, user_id; -- Results table is ordered by reaction_ratio in descending order, then by user_id in ascending order.
