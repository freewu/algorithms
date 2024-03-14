-- 1767. Find the Subtasks That Did Not Execute
-- Table: Tasks
-- +----------------+---------+
-- | Column Name    | Type    |
-- +----------------+---------+
-- | task_id        | int     |
-- | subtasks_count | int     |
-- +----------------+---------+
-- task_id is the column with unique values for this table.
-- Each row in this table indicates that task_id was divided into subtasks_count subtasks labeled from 1 to subtasks_count.
-- It is guaranteed that 2 <= subtasks_count <= 20.
 
-- Table: Executed
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | task_id       | int     |
-- | subtask_id    | int     |
-- +---------------+---------+
-- (task_id, subtask_id) is the combination of columns with unique values for this table.
-- Each row in this table indicates that for the task task_id, the subtask with ID subtask_id was executed successfully.
-- It is guaranteed that subtask_id <= subtasks_count for each task_id.
 
-- Write a solution to report the IDs of the missing subtasks for each task_id.
-- Return the result table in any order.
-- The result format is in the following example.
 
-- Example 1:
-- Input: 
-- Tasks table:
-- +---------+----------------+
-- | task_id | subtasks_count |
-- +---------+----------------+
-- | 1       | 3              |
-- | 2       | 2              |
-- | 3       | 4              |
-- +---------+----------------+
-- Executed table:
-- +---------+------------+
-- | task_id | subtask_id |
-- +---------+------------+
-- | 1       | 2          |
-- | 3       | 1          |
-- | 3       | 2          |
-- | 3       | 3          |
-- | 3       | 4          |
-- +---------+------------+
-- Output: 
-- +---------+------------+
-- | task_id | subtask_id |
-- +---------+------------+
-- | 1       | 1          |
-- | 1       | 3          |
-- | 2       | 1          |
-- | 2       | 2          |
-- +---------+------------+
-- Explanation: 
-- Task 1 was divided into 3 subtasks (1, 2, 3). Only subtask 2 was executed successfully, so we include (1, 1) and (1, 3) in the answer.
-- Task 2 was divided into 2 subtasks (1, 2). No subtask was executed successfully, so we include (2, 1) and (2, 2) in the answer.
-- Task 3 was divided into 4 subtasks (1, 2, 3, 4). All of the subtasks were executed successfully.

-- Create table If Not Exists Tasks (task_id int, subtasks_count int)
-- Create table If Not Exists Executed (task_id int, subtask_id int)
-- Truncate table Tasks
-- insert into Tasks (task_id, subtasks_count) values ('1', '3')
-- insert into Tasks (task_id, subtasks_count) values ('2', '2')
-- insert into Tasks (task_id, subtasks_count) values ('3', '4')
-- Truncate table Executed
-- insert into Executed (task_id, subtask_id) values ('1', '2')
-- insert into Executed (task_id, subtask_id) values ('3', '1')
-- insert into Executed (task_id, subtask_id) values ('3', '2')
-- insert into Executed (task_id, subtask_id) values ('3', '3')
-- insert into Executed (task_id, subtask_id) values ('3', '4')

-- 使用 recursive 生成任务
WITH RECURSIVE t AS 
(
    SELECT 1 AS subtask_id
    UNION ALL
    SELECT subtask_id + 1 FROM t WHERE subtask_id <= 20 -- 保证 2 <= subtasks_count <= 20
)

SELECT 
    task_id,
    subtask_id
FROM 
    Tasks, t
WHERE 
    (task_id, subtask_id) NOT IN ( SELECT * FROM Executed ) AND 
    subtask_id <= subtasks_count -- 对于每一个task_id，subtask_id <= subtasks_count

-- SELECT 
--     t.task_id,
--     s.subtask_id
-- FROM 
--     Tasks AS t 
-- CROSS JOIN
--     (
--         SELECT 1  AS subtask_id UNION
--         SELECT 2  AS subtask_id UNION
--         SELECT 3  AS subtask_id UNION
--         SELECT 4  AS subtask_id UNION
--         SELECT 5  AS subtask_id UNION
--         SELECT 6  AS subtask_id UNION
--         SELECT 7  AS subtask_id UNION
--         SELECT 8  AS subtask_id UNION
--         SELECT 9  AS subtask_id UNION
--         SELECT 10 AS subtask_id UNION
--         SELECT 11 AS subtask_id UNION
--         SELECT 12 AS subtask_id UNION
--         SELECT 13 AS subtask_id UNION
--         SELECT 14 AS subtask_id UNION
--         SELECT 15 AS subtask_id UNION
--         SELECT 16 AS subtask_id UNION
--         SELECT 17 AS subtask_id UNION
--         SELECT 18 AS subtask_id UNION
--         SELECT 19 AS subtask_id UNION
--         SELECT 20 AS subtask_id
--     ) AS s
-- WHERE 
--     t.subtasks_count >= subtask_id
-- ORDER BY 
--     task_id, subtask_id
-- | task_id | subtask_id |
-- | ------- | ---------- |
-- | 1       | 1          |
-- | 1       | 2          |
-- | 1       | 3          |
-- | 2       | 1          |
-- | 2       | 2          |
-- | 3       | 1          |
-- | 3       | 2          |
-- | 3       | 3          |
-- | 3       | 4          |

-- use union
SELECT 
    *
FROM 
(-- 生成一张完整的 task_id + subtask_id 表
    SELECT 
        t.task_id,
        s.subtask_id
    FROM 
        Tasks AS t 
    CROSS JOIN
        (
            SELECT 1  AS subtask_id UNION
            SELECT 2  AS subtask_id UNION
            SELECT 3  AS subtask_id UNION
            SELECT 4  AS subtask_id UNION
            SELECT 5  AS subtask_id UNION
            SELECT 6  AS subtask_id UNION
            SELECT 7  AS subtask_id UNION
            SELECT 8  AS subtask_id UNION
            SELECT 9  AS subtask_id UNION
            SELECT 10 AS subtask_id UNION
            SELECT 11 AS subtask_id UNION
            SELECT 12 AS subtask_id UNION
            SELECT 13 AS subtask_id UNION
            SELECT 14 AS subtask_id UNION
            SELECT 15 AS subtask_id UNION
            SELECT 16 AS subtask_id UNION
            SELECT 17 AS subtask_id UNION
            SELECT 18 AS subtask_id UNION
            SELECT 19 AS subtask_id UNION
            SELECT 20 AS subtask_id
        ) AS s
    WHERE 
        t.subtasks_count >= subtask_id
    ORDER BY 
        task_id, subtask_id
) AS a
WHERE 
    (task_id,subtask_id) NOT IN ( SELECT * FROM Executed)