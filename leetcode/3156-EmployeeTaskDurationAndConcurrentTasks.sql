-- 3156. Employee Task Duration and Concurrent Tasks
-- Table: Tasks
-- +---------------+----------+
-- | Column Name   | Type     |
-- +---------------+----------+
-- | task_id       | int      |
-- | employee_id   | int      |
-- | start_time    | datetime |
-- | end_time      | datetime |
-- +---------------+----------+
-- (task_id, employee_id) is the primary key for this table.
-- Each row in this table contains the task identifier, the employee identifier, and the start and end times of each task.
-- Write a solution to find the total duration of tasks for each employee and the maximum number of concurrent tasks an employee handled at any point in time. The total duration should be rounded down to the nearest number of full hours.

-- Return the result table ordered by employee_id ascending order.
-- The result format is in the following example.

-- Example:
-- Input:
-- Tasks table:
-- +---------+-------------+---------------------+---------------------+
-- | task_id | employee_id | start_time          | end_time            |
-- +---------+-------------+---------------------+---------------------+
-- | 1       | 1001        | 2023-05-01 08:00:00 | 2023-05-01 09:00:00 |
-- | 2       | 1001        | 2023-05-01 08:30:00 | 2023-05-01 10:30:00 |
-- | 3       | 1001        | 2023-05-01 11:00:00 | 2023-05-01 12:00:00 |
-- | 7       | 1001        | 2023-05-01 13:00:00 | 2023-05-01 15:30:00 |
-- | 4       | 1002        | 2023-05-01 09:00:00 | 2023-05-01 10:00:00 |
-- | 5       | 1002        | 2023-05-01 09:30:00 | 2023-05-01 11:30:00 |
-- | 6       | 1003        | 2023-05-01 14:00:00 | 2023-05-01 16:00:00 |
-- +---------+-------------+---------------------+---------------------+
-- Output:
-- +-------------+------------------+----------------------+
-- | employee_id | total_task_hours | max_concurrent_tasks |
-- +-------------+------------------+----------------------+
-- | 1001        | 6                | 2                    |
-- | 1002        | 2                | 2                    |
-- | 1003        | 2                | 1                    |
-- +-------------+------------------+----------------------+
-- Explanation:
-- For employee ID 1001:
-- Task 1 and Task 2 overlap from 08:30 to 09:00 (30 minutes).
-- Task 7 has a duration of 150 minutes (2 hours and 30 minutes).
-- Total task time: 60 (Task 1) + 120 (Task 2) + 60 (Task 3) + 150 (Task 7) - 30 (overlap) = 360 minutes = 6 hours.
-- Maximum concurrent tasks: 2 (during the overlap period).
-- For employee ID 1002:
-- Task 4 and Task 5 overlap from 09:30 to 10:00 (30 minutes).
-- Total task time: 60 (Task 4) + 120 (Task 5) - 30 (overlap) = 150 minutes = 2 hours and 30 minutes.
-- Total task hours (rounded down): 2 hours.
-- Maximum concurrent tasks: 2 (during the overlap period).
-- For employee ID 1003:
-- No overlapping tasks.
-- Total task time: 120 minutes = 2 hours.
-- Maximum concurrent tasks: 1.
-- Note: Output table is ordered by employee_id in ascending order.

-- CREATE TABLE  if Not exists Tasks (
--     task_id INT ,
--     employee_id INT,
--     start_time DATETIME,
--     end_time DATETIME
-- );
-- Truncate table Tasks;
-- insert into Tasks (task_id, employee_id, start_time, end_time) values ('1', '1001', '2023-05-01 08:00:00', '2023-05-01 09:00:00');
-- insert into Tasks (task_id, employee_id, start_time, end_time) values ('2', '1001', '2023-05-01 08:30:00', '2023-05-01 10:30:00');
-- insert into Tasks (task_id, employee_id, start_time, end_time) values ('3', '1001', '2023-05-01 11:00:00', '2023-05-01 12:00:00');
-- insert into Tasks (task_id, employee_id, start_time, end_time) values ('7', '1001', '2023-05-01 13:00:00', '2023-05-01 15:30:00');
-- insert into Tasks (task_id, employee_id, start_time, end_time) values ('4', '1002', '2023-05-01 09:00:00', '2023-05-01 10:00:00');
-- insert into Tasks (task_id, employee_id, start_time, end_time) values ('5', '1002', '2023-05-01 09:30:00', '2023-05-01 11:30:00');
-- insert into Tasks (task_id, employee_id, start_time, end_time) values ('6', '1003', '2023-05-01 14:00:00', '2023-05-01 16:00:00');

-- WITH t AS (-- 得到一张能有时间关联的任务表
--     SELECT
--         t1.*,
--         t2.start_time as t2s,
--         t2.end_time as t2e,
--         t2.task_id as t2id
--     FROM
--         Tasks AS t1
--     LEFT JOIN 
--         Tasks AS t2
--     ON 
--         t1.task_id != t2.task_id AND 
--         t1.employee_id = t2.employee_id AND
--         t1.start_time <= t2.start_time AND t1.end_time > t2.start_time
-- )
-- -- select * from t;
-- -- SELECT t2id FROM t;
-- SELECT
--     employee_id,
--     FLOOR(
--         SUM(
--             (TIMESTAMPDIFF(SECOND, t.start_time, IFNULL(t.t2e, t.end_time)) / 3600)
--         )
--     ) AS total_task_hours,
--     COUNT(DISTINCT t2id) + 1 AS max_concurrent_tasks 
-- FROM 
--     t
-- WHERE 
--     task_id NOT IN (SELECT t2id FROM t WHERE t2id IS NOT NULL)
-- GROUP BY 
--     employee_id 


-- Write your MySQL query statement below
WITH t AS ( -- 取出开始时间&结束时间
    SELECT DISTINCT employee_id, start_time AS st FROM Tasks
    UNION DISTINCT
    SELECT DISTINCT employee_id, end_time AS st   FROM Tasks
),
p AS (
    SELECT
        *,
        LEAD(st) OVER ( PARTITION BY employee_id ORDER BY st ) AS ed
    FROM 
        t
),
s AS (
    SELECT
        p.*,
        COUNT(1) AS concurrent_count
    FROM
        p
    INNER JOIN 
        Tasks AS t
    USING (employee_id)
    WHERE 
        p.st >= t.start_time AND p.ed <= t.end_time
    GROUP BY 
        p.employee_id, p.st , p.ed
)
-- -- select * from p;
-- | employee_id | st                  | ed                  |
-- | ----------- | ------------------- | ------------------- |
-- | 1001        | 2023-05-01 08:00:00 | 2023-05-01 08:30:00 |
-- | 1001        | 2023-05-01 08:30:00 | 2023-05-01 09:00:00 |
-- | 1001        | 2023-05-01 09:00:00 | 2023-05-01 10:30:00 |
-- | 1001        | 2023-05-01 10:30:00 | 2023-05-01 11:00:00 |
-- | 1001        | 2023-05-01 11:00:00 | 2023-05-01 12:00:00 |
-- | 1001        | 2023-05-01 12:00:00 | 2023-05-01 13:00:00 |
-- | 1001        | 2023-05-01 13:00:00 | 2023-05-01 15:30:00 |
-- | 1001        | 2023-05-01 15:30:00 | null                |
-- | 1002        | 2023-05-01 09:00:00 | 2023-05-01 09:30:00 |
-- | 1002        | 2023-05-01 09:30:00 | 2023-05-01 10:00:00 |
-- | 1002        | 2023-05-01 10:00:00 | 2023-05-01 11:30:00 |
-- | 1002        | 2023-05-01 11:30:00 | null                |
-- | 1003        | 2023-05-01 14:00:00 | 2023-05-01 16:00:00 |
-- | 1003        | 2023-05-01 16:00:00 | null                |

--select * from s;
-- | employee_id | st                  | ed                  | concurrent_count |
-- | ----------- | ------------------- | ------------------- | ---------------- |
-- | 1001        | 2023-05-01 08:00:00 | 2023-05-01 08:30:00 | 1                |
-- | 1001        | 2023-05-01 08:30:00 | 2023-05-01 09:00:00 | 2                |
-- | 1001        | 2023-05-01 09:00:00 | 2023-05-01 10:30:00 | 1                |
-- | 1001        | 2023-05-01 11:00:00 | 2023-05-01 12:00:00 | 1                |
-- | 1001        | 2023-05-01 13:00:00 | 2023-05-01 15:30:00 | 1                |
-- | 1002        | 2023-05-01 09:00:00 | 2023-05-01 09:30:00 | 1                |
-- | 1002        | 2023-05-01 09:30:00 | 2023-05-01 10:00:00 | 2                |
-- | 1002        | 2023-05-01 10:00:00 | 2023-05-01 11:30:00 | 1                |
-- | 1003        | 2023-05-01 14:00:00 | 2023-05-01 16:00:00 | 1                |

SELECT
    employee_id,
    FLOOR(SUM(TIMESTAMPDIFF(SECOND, st, ed) / 3600)) AS total_task_hours, 
    MAX(concurrent_count) AS max_concurrent_tasks
FROM 
    s
GROUP BY 
    employee_id
ORDER BY 
    employee_id
