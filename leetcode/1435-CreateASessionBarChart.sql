-- 1435. Create a Session Bar Chart
-- Table: Sessions
-- +---------------------+---------+
-- | Column Name         | Type    |
-- +---------------------+---------+
-- | session_id          | int     |
-- | duration            | int     |
-- +---------------------+---------+
-- session_id is the column of unique values for this table.
-- duration is the time in seconds that a user has visited the application.
 
-- You want to know how long a user visits your application. You decided to create bins of "[0-5>", "[5-10>", "[10-15>", and "15 minutes or more" and count the number of sessions on it.
-- Write a solution to report the (bin, total).
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Sessions table:
-- +-------------+---------------+
-- | session_id  | duration      |
-- +-------------+---------------+
-- | 1           | 30            |
-- | 2           | 199           |
-- | 3           | 299           |
-- | 4           | 580           |
-- | 5           | 1000          |
-- +-------------+---------------+
-- Output: 
-- +--------------+--------------+
-- | bin          | total        |
-- +--------------+--------------+
-- | [0-5>        | 3            |
-- | [5-10>       | 1            |
-- | [10-15>      | 0            |
-- | 15 or more   | 1            |
-- +--------------+--------------+
-- Explanation: 
-- For session_id 1, 2, and 3 have a duration greater or equal than 0 minutes and less than 5 minutes.
-- For session_id 4 has a duration greater or equal than 5 minutes and less than 10 minutes.
-- There is no session with a duration greater than or equal to 10 minutes and less than 15 minutes.
-- For session_id 5 has a duration greater than or equal to 15 minutes.

-- Create table If Not Exists Sessions (session_id int, duration int)
-- Truncate table Sessions
-- insert into Sessions (session_id, duration) values ('1', '30')
-- insert into Sessions (session_id, duration) values ('2', '199')
-- insert into Sessions (session_id, duration) values ('3', '299')
-- insert into Sessions (session_id, duration) values ('4', '580')
-- insert into Sessions (session_id, duration) values ('5', '1000')

(-- 统计  0 <= (duration / 60) < 5 的数据
    SELECT 
        '[0-5>' AS bin, 
        SUM(CASE 
                WHEN duration / 60 >= 0 AND duration / 60 < 5  THEN 1 
                ELSE 0 
        END) AS total 
    FROM Sessions
) 
UNION ALL 
(-- 统计  5 <= (duration / 60) < 10 的数据
    SELECT 
        '[5-10>' AS bin, 
        SUM(CASE 
                WHEN duration / 60 >= 5 AND duration / 60 < 10  THEN 1 
                ELSE 0 
        END) AS total 
    FROM Sessions
)
UNION ALL 
(-- 统计  10 <= (duration / 60) < 15 的数据
    SELECT 
        '[10-15>' AS bin, 
        SUM(CASE 
                WHEN duration / 60 >= 10 AND duration / 60 < 15  THEN 1 
                ELSE 0 
        END) AS total 
    FROM Sessions
)
UNION ALL 
(-- 统计  15 <= (duration / 60) 的数据
    SELECT 
        '15 or more' AS bin, 
        SUM(CASE 
                WHEN duration / 60 >= 15  THEN 1 
                ELSE 0 
        END) AS total 
    FROM Sessions
)

-- best solution
SELECT 
    bin,
    SUM(total) AS total 
FROM 
(
    -- 按区间统计分隔
    SELECT
        CASE 
            WHEN duration < 300 THEN '[0-5>'
            WHEN duration < 600 THEN '[5-10>'
            WHEN duration < 900 THEN '[10-15>'
            ELSE '15 or more' 
        END AS bin,
        COUNT(*) as total
    FROM 
        Sessions
    GROUP BY 
        bin
    -- 每个分区统一加上 一条为 0 的数据 处理没有相关数据空缺的问题
    UNION SELECT '[0-5>' bin,0 total 
    UNION SELECT '[5-10>' bin,0 total 
    UNION SELECT '[10-15>' bin,0 total 
    UNION SELECT '15 or more' bin,0 total       
) AS a
GROUP BY bin;