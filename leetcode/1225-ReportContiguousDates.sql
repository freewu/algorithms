-- 1225. Report Contiguous Dates
-- Table: Failed
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | fail_date    | date    |
-- +--------------+---------+
-- fail_date is the primary key (column with unique values) for this table.
-- This table contains the days of failed tasks.
 
-- Table: Succeeded
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | success_date | date    |
-- +--------------+---------+
-- success_date is the primary key (column with unique values) for this table.
-- This table contains the days of succeeded tasks.
 
-- A system is running one task every day. Every task is independent of the previous tasks. 
-- The tasks can fail or succeed.

-- Write a solution to report the period_state for each continuous interval of days in the period from 2019-01-01 to 2019-12-31.
-- period_state is 'failed' if tasks in this interval failed or 'succeeded' if tasks in this interval succeeded. 
-- Interval of days are retrieved as start_date and end_date.
-- Return the result table ordered by start_date.
-- The result format is in the following example.
 
-- Example 1:
-- Input: 
-- Failed table:
-- +-------------------+
-- | fail_date         |
-- +-------------------+
-- | 2018-12-28        |
-- | 2018-12-29        |
-- | 2019-01-04        |
-- | 2019-01-05        |
-- +-------------------+
-- Succeeded table:
-- +-------------------+
-- | success_date      |
-- +-------------------+
-- | 2018-12-30        |
-- | 2018-12-31        |
-- | 2019-01-01        |
-- | 2019-01-02        |
-- | 2019-01-03        |
-- | 2019-01-06        |
-- +-------------------+
-- Output: 
-- +--------------+--------------+--------------+
-- | period_state | start_date   | end_date     |
-- +--------------+--------------+--------------+
-- | succeeded    | 2019-01-01   | 2019-01-03   |
-- | failed       | 2019-01-04   | 2019-01-05   |
-- | succeeded    | 2019-01-06   | 2019-01-06   |
-- +--------------+--------------+--------------+
-- Explanation: 
-- The report ignored the system state in 2018 as we care about the system in the period 2019-01-01 to 2019-12-31.
-- From 2019-01-01 to 2019-01-03 all tasks succeeded and the system state was "succeeded".
-- From 2019-01-04 to 2019-01-05 all tasks failed and the system state was "failed".
-- From 2019-01-06 to 2019-01-06 all tasks succeeded and the system state was "succeeded".

-- Create table If Not Exists Failed (fail_date date)
-- Create table If Not Exists Succeeded (success_date date)
-- Truncate table Failed
-- insert into Failed (fail_date) values ('2018-12-28')
-- insert into Failed (fail_date) values ('2018-12-29')
-- insert into Failed (fail_date) values ('2019-01-04')
-- insert into Failed (fail_date) values ('2019-01-05')
-- Truncate table Succeeded
-- insert into Succeeded (success_date) values ('2018-12-30')
-- insert into Succeeded (success_date) values ('2018-12-31')
-- insert into Succeeded (success_date) values ('2019-01-01')
-- insert into Succeeded (success_date) values ('2019-01-02')
-- insert into Succeeded (success_date) values ('2019-01-03')
-- insert into Succeeded (success_date) values ('2019-01-06')

-- row_number
SELECT 
    state AS period_state, 
    min(date) AS start_date, 
    max(date) AS end_date
FROM 
(
    SELECT 
        *, 
        row_number() OVER (PARTITION BY state ORDER BY date ASC) AS rk1, -- 按 date ASC 排序 根据状态编号 
        row_number() OVER (ORDER BY date ASC) AS rk2
    FROM
    ( -- 把 Failed & Succeeded 两张表合并起来
        SELECT fail_date AS 'date', 'failed' AS state FROM failed
        UNION ALL
        SELECT success_date AS 'date', 'succeeded' AS state FROM succeeded 
    ) AS t
) AS t2
WHERE 
    date BETWEEN '2019-01-01' AND '2019-12-31' -- 找出 2019-01-01 到 2019-12-31 期间任务连续同状态
GROUP BY 
    state, rk2 - rk1


-- SELECT 
--     *, 
--     row_number() OVER (PARTITION BY state ORDER BY date ASC) rk1, -- 按 date ASC 排序 根据状态编号 
--     row_number() OVER (ORDER BY date ASC) rk2
-- FROM
-- ( -- 把 Failed & Succeeded 两张表合并起来
--     SELECT fail_date AS 'date', 'failed' AS state FROM failed
--     UNION ALL
--     SELECT success_date AS 'date', 'succeeded' AS state FROM succeeded 
-- ) AS t

-- | date       | state     | rk1 | rk2 |
-- | ---------- | --------- | --- | --- |
-- | 2018-12-28 | failed    | 1   | 1   |
-- | 2018-12-29 | failed    | 2   | 2   |
-- | 2018-12-30 | succeeded | 1   | 3   | 
-- | 2018-12-31 | succeeded | 2   | 4   |
-- | 2019-01-01 | succeeded | 3   | 5   |
-- | 2019-01-02 | succeeded | 4   | 6   |
-- | 2019-01-03 | succeeded | 5   | 7   |
-- | 2019-01-04 | failed    | 3   | 8   |
-- | 2019-01-05 | failed    | 4   | 9   |
-- | 2019-01-06 | succeeded | 6   | 10  |

-- SELECT
--     *,
--     rk2 - rk1 
-- FROM
-- (
--     SELECT 
--         *, 
--         row_number() OVER (PARTITION BY state ORDER BY date ASC) AS rk1, -- 按 date ASC 排序 根据状态编号 
--         row_number() OVER (ORDER BY date ASC) AS rk2
--     FROM
--     ( -- 把 Failed & Succeeded 两张表合并起来
--         SELECT fail_date AS 'date', 'failed' AS state FROM failed
--         UNION ALL
--         SELECT success_date AS 'date', 'succeeded' AS state FROM succeeded 
--     ) AS t
-- ) AS a
-- WHERE 
--     date BETWEEN '2019-01-01' AND '2019-12-31'

-- | date       | state     | rk1 | rk2 | rk2 - rk1 |
-- | ---------- | --------- | --- | --- | --------- |
-- | 2019-01-01 | succeeded | 3   | 5   | 2         |
-- | 2019-01-02 | succeeded | 4   | 6   | 2         |
-- | 2019-01-03 | succeeded | 5   | 7   | 2         |
-- | 2019-01-04 | failed    | 3   | 8   | 5         |
-- | 2019-01-05 | failed    | 4   | 9   | 5         |
-- | 2019-01-06 | succeeded | 6   | 10  | 4         |

-- best solution
SELECT 
    period_state,
    MIN(date) AS start_date,
    MAX(date) AS end_date
FROM 
(
    (
        SELECT 
            'succeeded' AS period_state,
            success_date AS date,
            DATE_SUB(success_date,INTERVAL ROW_NUMBER() OVER(ORDER BY success_date ASC) DAY) AS diff
        FROM Succeeded
        WHERE YEAR(success_date) = '2019'
    )
    UNION ALL
    (
        SELECT 
            'failed' AS period_state,
            fail_date AS date,
            DATE_SUB(fail_date,INTERVAL ROW_NUMBER() OVER(ORDER BY fail_date ASC) DAY) AS diff
        FROM Failed
        WHERE YEAR(fail_date) = '2019'
    )
) AS t
GROUP BY period_state, diff 
ORDER BY start_date