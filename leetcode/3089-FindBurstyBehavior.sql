-- 3089. Find Bursty Behavior
-- Table: Posts
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | post_id     | int     |
-- | user_id     | int     |
-- | post_date   | date    |
-- +-------------+---------+
-- post_id is the primary key (column with unique values) for this table.
-- Each row of this table contains post_id, user_id, and post_date.
-- Write a solution to find users who demonstrate bursty behavior in their posting patterns during February 2024. Bursty behavior is defined as any period of 7 consecutive days where a user's posting frequency is at least twice to their average weekly posting frequency for February 2024.

-- Note: Only include the dates from February 1 to February 28 in your analysis, which means you should count February as having exactly 4 weeks.
-- Return the result table orderd by user_id in ascending order.
-- The result format is in the following example.


-- Example:
-- Input:
-- Posts table:
-- +---------+---------+------------+
-- | post_id | user_id | post_date  |
-- +---------+---------+------------+
-- | 1       | 1       | 2024-02-27 |
-- | 2       | 5       | 2024-02-06 |
-- | 3       | 3       | 2024-02-25 |
-- | 4       | 3       | 2024-02-14 |
-- | 5       | 3       | 2024-02-06 |
-- | 6       | 2       | 2024-02-25 |
-- +---------+---------+------------+
-- Output:
-- +---------+----------------+------------------+
-- | user_id | max_7day_posts | avg_weekly_posts |
-- +---------+----------------+------------------+
-- | 1       | 1              | 0.2500           |
-- | 2       | 1              | 0.2500           |
-- | 5       | 1              | 0.2500           |
-- +---------+----------------+------------------+
-- Explanation:
-- User 1: Made only 1 post in February, resulting in an average of 0.25 posts per week and a max of 1 post in any 7-day period.
-- User 2: Also made just 1 post, with the same average and max 7-day posting frequency as User 1.
-- User 5: Like Users 1 and 2, User 5 made only 1 post throughout February, leading to the same average and max 7-day posting metrics.
-- User 3: Although User 3 made more posts than the others (3 posts), they did not reach twice the average weekly posts in their consecutive 7-day window, so they are not listed in the output.
-- Note: Output table is ordered by user_id in ascending order.

-- Create table If Not Exists Posts (post_id int, user_id int,  post_date date)
-- Truncate table Posts
-- insert into Posts (post_id, user_id, post_date) values ('1', '1', '2024-02-27')
-- insert into Posts (post_id, user_id, post_date) values ('2', '5', '2024-02-06')
-- insert into Posts (post_id, user_id, post_date) values ('3', '3', '2024-02-25')
-- insert into Posts (post_id, user_id, post_date) values ('4', '3', '2024-02-14')
-- insert into Posts (post_id, user_id, post_date) values ('5', '3', '2024-02-06')
-- insert into Posts (post_id, user_id, post_date) values ('6', '2', '2024-02-25')

-- SELECT 
--     user_id,
--     MAX(cnt) AS max_7day_posts,
--     SUM(cnt) / 4 AS avg_weekly_posts
-- FROM 
-- (
--     (
--         SELECT 
--             COUNT(*) AS cnt,
--             user_id
--         FROM
--             Posts
--         WHERE
--             post_date BETWEEN '2024-02-01' AND '2024-02-06'
--         GROUP BY
--             user_id
--     ) 
--     UNION ALL
--     (
--         SELECT 
--             COUNT(*) AS cnt,
--             user_id
--         FROM
--             Posts
--         WHERE
--             post_date BETWEEN '2024-02-07' AND '2024-02-13'
--         GROUP BY
--             user_id
--     )
--     UNION ALL
--     (
--         SELECT 
--             COUNT(*) AS cnt,
--             user_id
--         FROM
--             Posts
--         WHERE
--             post_date BETWEEN '2024-02-14' AND '2024-02-20'
--         GROUP BY
--             user_id
--     )
--     UNION ALL
--     (
--         SELECT 
--             COUNT(*) AS cnt,
--             user_id
--         FROM
--             Posts
--         WHERE
--             post_date BETWEEN '2024-02-21' AND '2024-02-28'
--         GROUP BY
--             user_id
--     )
-- ) AS t
-- GROUP BY 
--     user_id
-- HAVING 
--     COUNT(*) <= 2
-- ORDER BY
--     user_id

-- SELECT 
--     t.*,
--     rank() OVER(ORDER BY week) AS "rk"
-- FROM
-- (
--     (
--         SELECT 
--             '1' AS week,
--             COUNT(*) AS cnt,
--             user_id
--         FROM
--             Posts
--         WHERE
--             post_date BETWEEN '2024-02-01' AND '2024-02-06'
--         GROUP BY
--             user_id
--     ) 
--     UNION ALL
--     (
--         SELECT 
--             '2' AS week,
--             COUNT(*) AS cnt,
--             user_id
--         FROM
--             Posts
--         WHERE
--             post_date BETWEEN '2024-02-07' AND '2024-02-13'
--         GROUP BY
--             user_id
--     )
--     UNION ALL
--     (
--         SELECT 
--             '3' AS week,
--             COUNT(*) AS cnt,
--             user_id
--         FROM
--             Posts
--         WHERE
--             post_date BETWEEN '2024-02-14' AND '2024-02-20'
--         GROUP BY
--             user_id
--     )
--     UNION ALL
--     (
--         SELECT 
--             '4' AS week,
--             COUNT(*) AS cnt,
--             user_id
--         FROM
--             Posts
--         WHERE
--             post_date BETWEEN '2024-02-21' AND '2024-02-28'
--         GROUP BY
--             user_id
--     )
-- ) AS t


WITH 
P AS ( -- 用户任何 连续 7 天 的时段发帐数
    SELECT 
        p1.user_id AS user_id, 
        COUNT(1) AS cnt
    FROM
        Posts AS p1
    JOIN 
        Posts AS p2
    ON 
        p1.user_id = p2.user_id AND 
        -- 任何 连续 7 天 的时段
        p2.post_date BETWEEN p1.post_date AND DATE_ADD(p1.post_date, INTERVAL 6 DAY)
    GROUP BY 
        p1.user_id, p1.post_date
),
T AS ( -- 2月平均每周发贴数
    SELECT 
        user_id, 
        COUNT(1) / 4 AS avg_weekly_posts
    FROM 
        Posts
    WHERE 
        post_date BETWEEN '2024-02-01' AND '2024-02-28'
    GROUP BY 
        user_id
)

SELECT 
    user_id, 
    MAX(cnt) AS max_7day_posts, 
    avg_weekly_posts
FROM
    P
JOIN T USING (user_id)
GROUP BY 
    user_id
HAVING 
    max_7day_posts >= avg_weekly_posts * 2 -- 任何 连续 7 天 的时段中发帖频率是其 平均 每周发帖频率的 至少两倍。
ORDER BY 
    user_id;

# Write your MySQL query statement below
-- WITH
--     P AS (
--         SELECT p1.user_id AS user_id, COUNT(1) AS cnt
--         FROM
--             Posts AS p1
--             JOIN Posts AS p2
--                 ON p1.user_id = p2.user_id
--                 AND p2.post_date BETWEEN p1.post_date AND DATE_ADD(p1.post_date, INTERVAL 6 DAY)
--         GROUP BY p1.user_id, p1.post_id
--     ),
--     T AS (
--         SELECT user_id, COUNT(1) / 4 AS avg_weekly_posts
--         FROM Posts
--         WHERE post_date BETWEEN '2024-02-01' AND '2024-02-28'
--         GROUP BY user_id
--     )
-- SELECT user_id, MAX(cnt) AS max_7day_posts, avg_weekly_posts
-- FROM
--     P
--     JOIN T USING (user_id)
-- GROUP BY user_id
-- HAVING max_7day_posts >= avg_weekly_posts * 2
-- ORDER BY user_id;