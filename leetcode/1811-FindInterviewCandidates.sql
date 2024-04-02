-- 1811. Find Interview Candidates
-- Table: Contests
-- +--------------+------+
-- | Column Name  | Type |
-- +--------------+------+
-- | contest_id   | int  |
-- | gold_medal   | int  |
-- | silver_medal | int  |
-- | bronze_medal | int  |
-- +--------------+------+
-- contest_id is the column with unique values for this table.
-- This table contains the LeetCode contest ID and the user IDs of the gold, silver, and bronze medalists.
-- It is guaranteed that any consecutive contests have consecutive IDs and that no ID is skipped.

-- Table: Users
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | user_id     | int     |
-- | mail        | varchar |
-- | name        | varchar |
-- +-------------+---------+
-- user_id is the column with unique values for this table.
-- This table contains information about the users.
 
-- Write a solution to report the name and the mail of all interview candidates. 
-- A user is an interview candidate if at least one of these two conditions is true:
--     The user won any medal in three or more consecutive contests.
--     The user won the gold medal in three or more different contests (not necessarily consecutive).

-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Contests table:
-- +------------+------------+--------------+--------------+
-- | contest_id | gold_medal | silver_medal | bronze_medal |
-- +------------+------------+--------------+--------------+
-- | 190        | 1          | 5            | 2            |
-- | 191        | 2          | 3            | 5            |
-- | 192        | 5          | 2            | 3            |
-- | 193        | 1          | 3            | 5            |
-- | 194        | 4          | 5            | 2            |
-- | 195        | 4          | 2            | 1            |
-- | 196        | 1          | 5            | 2            |
-- +------------+------------+--------------+--------------+
-- Users table:
-- +---------+--------------------+-------+
-- | user_id | mail               | name  |
-- +---------+--------------------+-------+
-- | 1       | sarah@leetcode.com | Sarah |
-- | 2       | bob@leetcode.com   | Bob   |
-- | 3       | alice@leetcode.com | Alice |
-- | 4       | hercy@leetcode.com | Hercy |
-- | 5       | quarz@leetcode.com | Quarz |
-- +---------+--------------------+-------+
-- Output: 
-- +-------+--------------------+
-- | name  | mail               |
-- +-------+--------------------+
-- | Sarah | sarah@leetcode.com |
-- | Bob   | bob@leetcode.com   |
-- | Alice | alice@leetcode.com |
-- | Quarz | quarz@leetcode.com |
-- +-------+--------------------+
-- Explanation: 
-- Sarah won 3 gold medals (190, 193, and 196), so we include her in the result table.
-- Bob won a medal in 3 consecutive contests (190, 191, and 192), so we include him in the result table.
--     - Note that he also won a medal in 3 other consecutive contests (194, 195, and 196).
-- Alice won a medal in 3 consecutive contests (191, 192, and 193), so we include her in the result table.
-- Quarz won a medal in 5 consecutive contests (190, 191, 192, 193, and 194), so we include them in the result table.
 
-- Follow up:
--     What if the first condition changed to be "any medal in n or more consecutive contests"? How would you change your solution to get the interview candidates? Imagine that n is the parameter of a stored procedure.
--     Some users may not participate in every contest but still perform well in the ones they do. How would you change your solution to only consider contests where the user was a participant? Suppose the registered users for each contest are given in another table.

-- Create table If Not Exists Contests (contest_id int, gold_medal int, silver_medal int, bronze_medal int)
-- Create table If Not Exists Users (user_id int, mail varchar(50), name varchar(30))
-- Truncate table Contests
-- insert into Contests (contest_id, gold_medal, silver_medal, bronze_medal) values ('190', '1', '5', '2')
-- insert into Contests (contest_id, gold_medal, silver_medal, bronze_medal) values ('191', '2', '3', '5')
-- insert into Contests (contest_id, gold_medal, silver_medal, bronze_medal) values ('192', '5', '2', '3')
-- insert into Contests (contest_id, gold_medal, silver_medal, bronze_medal) values ('193', '1', '3', '5')
-- insert into Contests (contest_id, gold_medal, silver_medal, bronze_medal) values ('194', '4', '5', '2')
-- insert into Contests (contest_id, gold_medal, silver_medal, bronze_medal) values ('195', '4', '2', '1')
-- insert into Contests (contest_id, gold_medal, silver_medal, bronze_medal) values ('196', '1', '5', '2')
-- Truncate table Users
-- insert into Users (user_id, mail, name) values ('1', 'sarah@leetcode.com', 'Sarah')
-- insert into Users (user_id, mail, name) values ('2', 'bob@leetcode.com', 'Bob')
-- insert into Users (user_id, mail, name) values ('3', 'alice@leetcode.com', 'Alice')
-- insert into Users (user_id, mail, name) values ('4', 'hercy@leetcode.com', 'Hercy')
-- insert into Users (user_id, mail, name) values ('5', 'quarz@leetcode.com', 'Quarz')

-- union all
WITH t1 AS (
    SELECT 
        *, 
        ROW_NUMBER() OVER(PARTITION BY id ORDER BY contest_id) AS rk
    FROM (
        (SELECT contest_id, gold_medal AS id FROM contests) -- 金牌用户
        UNION ALL
        (SELECT contest_id, silver_medal AS id FROM contests) -- 银牌用户
        UNION ALL
        (SELECT contest_id, bronze_medal AS id FROM contests) -- 铜牌用户
        ORDER BY  id, contest_id
    ) AS t
)
SELECT 
    name, 
    mail
FROM 
    Users
WHERE 
    user_id IN (
        (SELECT id FROM t1 GROUP BY id, contest_id - rk HAVING COUNT(*) >= 3) -- 连续三次获得奖牌以上
        UNION
        (SELECT gold_medal AS id FROM Contests GROUP BY gold_medal HAVING COUNT(*) >= 3)-- 获得三次金牌以上
    )

-- lead
WITH cte as (
    select contest_id, gold_medal as user_id from Contests
    union all
    select contest_id, silver_medal from Contests
    union all
    select contest_id, bronze_medal from Contests
)
select 
    name,mail 
from (                  
    select
        user_id
    from (
        select 
            contest_id,user_id,
            lead(contest_id,1) over(partition by user_id order by contest_id) as second,
            lead(contest_id,2) over(partition by user_id order by contest_id) as third
        from cte
    ) t
    where contest_id+1=second AND contest_id+2=third
    UNION 
    SELECT gold_medal
    from Contests
    group by gold_medal
    having count(*)>=3
) t2
join Users using (user_id)