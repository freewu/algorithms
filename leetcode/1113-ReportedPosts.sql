-- 1113. Reported Posts
-- Table: Actions

-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | user_id       | int     |
-- | post_id       | int     |
-- | action_date   | date    | 
-- | action        | enum    |
-- | extra         | varchar |
-- +---------------+---------+
-- This table may have duplicate rows.
-- The action column is an ENUM (category) type of ('view', 'like', 'reaction', 'comment', 'report', 'share').
-- The extra column has optional information about the action, such as a reason for the report or a type of reaction.
-- extra is never NULL.
 
-- Write a solution to report the number of posts reported yesterday for each report reason. Assume today is 2019-07-05.
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:

-- Input: 
-- Actions table:
-- +---------+---------+-------------+--------+--------+
-- | user_id | post_id | action_date | action | extra  |
-- +---------+---------+-------------+--------+--------+
-- | 1       | 1       | 2019-07-01  | view   | null   |
-- | 1       | 1       | 2019-07-01  | like   | null   |
-- | 1       | 1       | 2019-07-01  | share  | null   |
-- | 2       | 4       | 2019-07-04  | view   | null   |
-- | 2       | 4       | 2019-07-04  | report | spam   |
-- | 3       | 4       | 2019-07-04  | view   | null   |
-- | 3       | 4       | 2019-07-04  | report | spam   |
-- | 4       | 3       | 2019-07-02  | view   | null   |
-- | 4       | 3       | 2019-07-02  | report | spam   |
-- | 5       | 2       | 2019-07-04  | view   | null   |
-- | 5       | 2       | 2019-07-04  | report | racism |
-- | 5       | 5       | 2019-07-04  | view   | null   |
-- | 5       | 5       | 2019-07-04  | report | racism |
-- +---------+---------+-------------+--------+--------+
-- Output: 
-- +---------------+--------------+
-- | report_reason | report_count |
-- +---------------+--------------+
-- | spam          | 1            |
-- | racism        | 2            |
-- +---------------+--------------+
-- Explanation: Note that we only care about report reasons with non-zero number of reports.

-- 所需知识
-- 获取当前日期：
-- SELECT CURRENT_DATE();

-- 在指定日期上加或减天数：
-- 在当前日期上加7天
-- SELECT DATE_ADD(CURRENT_DATE(), INTERVAL 7 DAY);
-- 在指定日期'2023-04-15'上减去3天
-- SELECT DATE_SUB('2023-04-15', INTERVAL 3 DAY);

-- 计算两个日期之间的天数差：
-- 计算从今天到明年元旦（2024-01-01）的天数差
-- SELECT TIMESTAMPDIFF(DAY, CURRENT_DATE(), '2024-01-01');

-- Write your MySQL query statement below
-- SELECT 
--     extra AS report_reason,
--     COUNT(DISTINCT post_id) AS report_count
-- FROM
--     Actions 
-- WHERE 
--     action_date =  DATE_SUB(CURRENT_DATE(), INTERVAL 1 DAY) AND
--     action = 'report'
-- GROUP BY
--     report_reason

SELECT 
    extra AS report_reason,
    COUNT(DISTINCT post_id) AS report_count
FROM
    Actions 
WHERE 
    action_date =  '2019-07-04' AND
    action = 'report'
GROUP BY
    report_reason