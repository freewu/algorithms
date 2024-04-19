-- 2394. Employees With Deductions
-- Table: Employees
-- +--------------+------+
-- | Column Name  | Type |
-- +--------------+------+
-- | employee_id  | int  |
-- | needed_hours | int  |
-- +--------------+------+
-- employee_id is column with unique values for this table.
-- Each row contains the id of an employee and the minimum number of hours needed for them to work to get their salary.

-- Table: Logs
-- +-------------+----------+
-- | Column Name | Type     |
-- +-------------+----------+
-- | employee_id | int      |
-- | in_time     | datetime |
-- | out_time    | datetime |
-- +-------------+----------+
-- (employee_id, in_time, out_time) is the primary key (combination of columns with unique values) for this table.
-- Each row of this table shows the time stamps for an employee. in_time is the time the employee started to work, and out_time is the time the employee ended work.
-- All the times are in October 2022. out_time can be one day after in_time which means the employee worked after the midnight.
 
-- In a company, each employee must work a certain number of hours every month. Employees work in sessions. The number of hours an employee worked can be calculated from the sum of the number of minutes the employee worked in all of their sessions. The number of minutes in each session is rounded up.

-- For example, if the employee worked for 51 minutes and 2 seconds in a session, we consider it 52 minutes.
-- Write a solution to report the IDs of the employees that will be deducted. In other words, report the IDs of the employees that did not work the needed hours.

-- Return the result table in any order.
-- The result format is in the following example.


-- Example 1:
-- Input: 
-- Employees table:
-- +-------------+--------------+
-- | employee_id | needed_hours |
-- +-------------+--------------+
-- | 1           | 20           |
-- | 2           | 12           |
-- | 3           | 2            |
-- +-------------+--------------+
-- Logs table:
-- +-------------+---------------------+---------------------+
-- | employee_id | in_time             | out_time            |
-- +-------------+---------------------+---------------------+
-- | 1           | 2022-10-01 09:00:00 | 2022-10-01 17:00:00 |
-- | 1           | 2022-10-06 09:05:04 | 2022-10-06 17:09:03 |
-- | 1           | 2022-10-12 23:00:00 | 2022-10-13 03:00:01 |
-- | 2           | 2022-10-29 12:00:00 | 2022-10-29 23:58:58 |
-- +-------------+---------------------+---------------------+
-- Output: 
-- +-------------+
-- | employee_id |
-- +-------------+
-- | 2           |
-- | 3           |
-- +-------------+
-- Explanation: 
-- Employee 1:
--  - Worked for three sessions:
--     - On 2022-10-01, they worked for 8 hours.
--     - On 2022-10-06, they worked for 8 hours and 4 minutes.
--     - On 2022-10-12, they worked for 4 hours and 1 minute. Note that they worked through midnight.
--  - Employee 1 worked a total of 20 hours and 5 minutes across sessions and will not be deducted.
-- Employee 2:
--  - Worked for one session:
--     - On 2022-10-29, they worked for 11 hours and 59 minutes.
--  - Employee 2 did not work their hours and will be deducted.
-- Employee 3:
--  - Did not work any session.
--  - Employee 3 did not work their hours and will be deducted.

-- Create table If Not Exists Employees (employee_id int, needed_hours int)
-- Create table If Not Exists Logs (employee_id int, in_time datetime, out_time datetime)
-- Truncate table Employees
-- insert into Employees (employee_id, needed_hours) values ('1', '20')
-- insert into Employees (employee_id, needed_hours) values ('2', '12')
-- insert into Employees (employee_id, needed_hours) values ('3', '2')
-- Truncate table Logs
-- insert into Logs (employee_id, in_time, out_time) values ('1', '2022-10-01 09:00:00', '2022-10-01 17:00:00')
-- insert into Logs (employee_id, in_time, out_time) values ('1', '2022-10-06 09:05:04', '2022-10-06 17:09:03')
-- insert into Logs (employee_id, in_time, out_time) values ('1', '2022-10-12 23:00:00', '2022-10-13 03:00:01')
-- insert into Logs (employee_id, in_time, out_time) values ('2', '2022-10-29 12:00:00', '2022-10-29 23:58:58')

-- DATEDIFF函数与TIMESTAMPDIFF函数的区别 
-- 1.DATEDIFF函数仅用于返回两个日期的天数，TIMESTAMPDIFF函数用于返回计算两个日期指定单位的时间差(指定单位可以是年，季度，月，星期，天数，小时，分钟，秒等等)
-- 2.对日期差值的计算方式相反
-- DATEDIFF函数的语法格式: DATEDIFF(start,end)
--     DATEDIFF函数返回start - end的计算结果

-- TIMESTAMPDIFF函数的语法格式: TIMESTAMPDIFF(DAY,start,end)
--     TIMESTAMPDIFF函数返回end - start的计算结果

-- -- -2
-- SELECT DATEDIFF('2022-04-28', '2022-04-30');
-- -- 2
-- SELECT TIMESTAMPDIFF(DAY,'2022-04-28', '2022-04-30');

-- CEIL(x)	向上取整
-- FLOOR(x)	向下取整


-- 员工工作时长(小时)
-- SELECT 
--     employee_id,
--     SUM(CEIL(TIMESTAMPDIFF(MINUTE,in_time, out_time)) / 60)  AS m -- 分钟数
-- FROM
--     Logs
-- GROUP BY
--     employee_id

-- Write your MySQL query statement below
-- -- 员工工作时长(小时)
-- SELECT 
--     employee_id,
--     SUM(CEIL(TIMESTAMPDIFF(SECOND,in_time, out_time) / 60) / 60)  AS m -- 分钟数
-- FROM
--     Logs
-- GROUP BY
--     employee_id

-- SELECT
--     e.employee_id 
-- FROM
--     Employees AS e
-- LEFT JOIN
--     (
--         -- 员工工作时长(小时)
--         SELECT 
--             employee_id,
--             SUM(TIMESTAMPDIFF(MINUTE,in_time, out_time)) / 60 AS m -- 分钟数
--         FROM
--             Logs
--         GROUP BY
--             employee_id
--     ) AS l 
-- ON
--     e.employee_id = l.employee_id
-- WHERE
--     l.m IS NULL OR -- 没有工作记录的人
--     l.m < e.needed_hours -- 工作时间长不足


SELECT
    e.employee_id 
FROM
    Employees AS e
LEFT JOIN
    (
        -- 员工工作时长(小时)
        SELECT 
            employee_id,
            -- 如果员工在一个时间段中工作了 51 分 2 秒，我们就认为它是 52 分钟 只对秒级做向上取正的处理，即:
            -- CEIL(TIMESTAMPDIFF(SECOND,in_time, out_time) / 60)
            SUM(CEIL(TIMESTAMPDIFF(SECOND,in_time, out_time) / 60) / 60)  AS m -- 用时
        FROM
            Logs
        GROUP BY
            employee_id
    ) AS l 
ON
    e.employee_id = l.employee_id
WHERE
    l.m IS NULL OR -- 没有工作记录的人
    l.m < e.needed_hours -- 工作时间长不足