-- 3268. Find Overlapping Shifts II
-- Table: EmployeeShifts
-- +------------------+----------+
-- | Column Name      | Type     |
-- +------------------+----------+
-- | employee_id      | int      |
-- | start_time       | datetime |
-- | end_time         | datetime |
-- +------------------+----------+
-- (employee_id, start_time) is the unique key for this table.
-- This table contains information about the shifts worked by employees, including the start time, and end time.
-- Write a solution to analyze overlapping shifts for each employee. Two shifts are considered overlapping if they occur on the same date and one shift's end_time is later than another shift's start_time.

-- For each employee, calculate the following:
--     The maximum number of shifts that overlap at any given time.
--     The total duration of all overlaps in minutes.

-- Return the result table ordered by employee_id in ascending order.
-- The query result format is in the following example.

-- Example:
-- Input:
-- EmployeeShifts table:
-- +-------------+---------------------+---------------------+
-- | employee_id | start_time          | end_time            |
-- +-------------+---------------------+---------------------+
-- | 1           | 2023-10-01 09:00:00 | 2023-10-01 17:00:00 |
-- | 1           | 2023-10-01 15:00:00 | 2023-10-01 23:00:00 |
-- | 1           | 2023-10-01 16:00:00 | 2023-10-02 00:00:00 |
-- | 2           | 2023-10-01 09:00:00 | 2023-10-01 17:00:00 |
-- | 2           | 2023-10-01 11:00:00 | 2023-10-01 19:00:00 |
-- | 3           | 2023-10-01 09:00:00 | 2023-10-01 17:00:00 |
-- +-------------+---------------------+---------------------+
-- Output:
-- +-------------+---------------------------+------------------------+
-- | employee_id | max_overlapping_shifts    | total_overlap_duration |
-- +-------------+---------------------------+------------------------+
-- | 1           | 3                         | 600                    |
-- | 2           | 2                         | 360                    |
-- | 3           | 1                         | 0                      |
-- +-------------+---------------------------+------------------------+
-- Explanation:
--     Employee 1 has 3 shifts:
--         2023-10-01 09:00:00 to 2023-10-01 17:00:00
--         2023-10-01 15:00:00 to 2023-10-01 23:00:00
--         2023-10-01 16:00:00 to 2023-10-02 00:00:00
--     The maximum number of overlapping shifts is 3 (from 16:00 to 17:00). 
--     The total overlap duration is: - 2 hours (15:00-17:00) between 1st and 2nd shifts - 1 hour (16:00-17:00) between 1st and 3rd shifts - 7 hours (16:00-23:00) between 2nd and 3rd shifts Total: 10 hours = 600 minutes
--     Employee 2 has 2 shifts:
--         2023-10-01 09:00:00 to 2023-10-01 17:00:00
--         2023-10-01 11:00:00 to 2023-10-01 19:00:00
--     The maximum number of overlapping shifts is 2. The total overlap duration is 6 hours (11:00-17:00) = 360 minutes.
--     Employee 3 has only 1 shift, so there are no overlaps.
-- The output table contains the employee_id, the maximum number of simultaneous overlaps, and the total overlap duration in minutes for each employee, ordered by employee_id in ascending order.

-- Create table if not exists EmployeeShifts(employee_id int, start_time datetime, end_time datetime)
-- Truncate table EmployeeShifts
-- insert into EmployeeShifts (employee_id, start_time, end_time) values ('1', '2023-10-01 09:00:00', '2023-10-01 17:00:00')
-- insert into EmployeeShifts (employee_id, start_time, end_time) values ('1', '2023-10-01 15:00:00', '2023-10-01 23:00:00')
-- insert into EmployeeShifts (employee_id, start_time, end_time) values ('1', '2023-10-01 16:00:00', '2023-10-02 00:00:00')
-- insert into EmployeeShifts (employee_id, start_time, end_time) values ('2', '2023-10-01 09:00:00', '2023-10-01 17:00:00')
-- insert into EmployeeShifts (employee_id, start_time, end_time) values ('2', '2023-10-01 11:00:00', '2023-10-01 19:00:00')
-- insert into EmployeeShifts (employee_id, start_time, end_time) values ('3', '2023-10-01 09:00:00', '2023-10-01 17:00:00')

# Write your MySQL query statement below
WITH  t AS ( -- 将开始时间和结束时间合并成一张表
    (
        SELECT 
            DISTINCT employee_id,
            start_time AS st
        FROM 
            EmployeeShifts
    ) 
    UNION DISTINCT
    (
        SELECT 
            DISTINCT employee_id, 
            end_time AS st
        FROM 
            EmployeeShifts
    )
),
t1 AS (
    SELECT
        *,
        LEAD(st) OVER ( PARTITION BY employee_id ORDER BY st ) AS ed
    FROM 
        t
),
t2 AS (
    SELECT
        t1.*,
        COUNT(1) AS concurrent_count
    FROM
        t1
    INNER JOIN 
        EmployeeShifts AS e 
    ON 
        t1.employee_id = e.employee_id
    WHERE 
        t1.st >= e.start_time AND t1.ed <= e.end_time
    GROUP BY 
        1, 2, 3
),
t3 AS (
    SELECT
        e1.employee_id,
        SUM(TIMESTAMPDIFF(MINUTE, e2.start_time, LEAST(e1.end_time, e2.end_time))) AS total_overlap_duration
    FROM
        EmployeeShifts AS e1
    JOIN 
        EmployeeShifts AS e2
    ON 
        e1.employee_id = e2.employee_id AND e1.start_time < e2.start_time AND e1.end_time > e2.start_time
    GROUP BY 
        e1.employee_id
)

SELECT
    t2.employee_id AS employee_id,
    MAX(concurrent_count) AS max_overlapping_shifts,
    IFNULL(AVG(total_overlap_duration), 0) AS total_overlap_duration
FROM
    t2
LEFT JOIN 
    t3 
ON 
    t2.employee_id = t3.employee_id
GROUP BY 
    employee_id
ORDER BY 
    employee_id;