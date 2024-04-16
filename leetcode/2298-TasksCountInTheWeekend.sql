-- 2298. Tasks Count in the Weekend
-- Table: Tasks
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | task_id     | int  |
-- | assignee_id | int  |
-- | submit_date | date |
-- +-------------+------+
-- task_id is the primary key (column with unique values) for this table.
-- Each row in this table contains the ID of a task, the id of the assignee, and the submission date.

-- Write a solution to report:
--     the number of tasks that were submitted during the weekend (Saturday, Sunday) as weekend_cnt, and
--     the number of tasks that were submitted during the working days as working_cnt.

-- Return the result table in any order.
-- The result format is shown in the following example.

-- Example 1:
-- Input: 
-- Tasks table:
-- +---------+-------------+-------------+
-- | task_id | assignee_id | submit_date |
-- +---------+-------------+-------------+
-- | 1       | 1           | 2022-06-13  |
-- | 2       | 6           | 2022-06-14  |
-- | 3       | 6           | 2022-06-15  |
-- | 4       | 3           | 2022-06-18  |
-- | 5       | 5           | 2022-06-19  |
-- | 6       | 7           | 2022-06-19  |
-- +---------+-------------+-------------+
-- Output: 
-- +-------------+-------------+
-- | weekend_cnt | working_cnt |
-- +-------------+-------------+
-- | 3           | 3           |
-- +-------------+-------------+
-- Explanation: 
-- Task 1 was submitted on Monday.
-- Task 2 was submitted on Tuesday.
-- Task 3 was submitted on Wednesday.
-- Task 4 was submitted on Saturday.
-- Task 5 was submitted on Sunday.
-- Task 6 was submitted on Sunday.
-- 3 tasks were submitted during the weekend.
-- 3 tasks were submitted during the working days.

-- Create table If Not Exists Tasks (task_id int, assignee_id int, submit_date date)
-- Truncate table Tasks
-- insert into Tasks (task_id, assignee_id, submit_date) values ('1', '1', '2022-06-13')
-- insert into Tasks (task_id, assignee_id, submit_date) values ('2', '6', '2022-06-14')
-- insert into Tasks (task_id, assignee_id, submit_date) values ('3', '6', '2022-06-15')
-- insert into Tasks (task_id, assignee_id, submit_date) values ('4', '3', '2022-06-18')
-- insert into Tasks (task_id, assignee_id, submit_date) values ('5', '5', '2022-06-19')
-- insert into Tasks (task_id, assignee_id, submit_date) values ('6', '7', '2022-06-19')

-- DATE_FORMAT(date,format)
-- date 参数是合法的日期。format 规定日期/时间的输出格式。

-- 可以使用的格式有：
-- 格式	描述
-- %a	缩写星期名
-- %b	缩写月名
-- %c	月，数值
-- %D	带有英文前缀的月中的天
-- %d	月的天，数值(00-31)
-- %e	月的天，数值(0-31)
-- %f	微秒
-- %H	小时 (00-23)
-- %h	小时 (01-12)
-- %I	小时 (01-12)
-- %i	分钟，数值(00-59)
-- %j	年的天 (001-366)
-- %k	小时 (0-23)
-- %l	小时 (1-12)
-- %M	月名
-- %m	月，数值(00-12)
-- %p	AM 或 PM
-- %r	时间，12-小时（hh:mm:ss AM 或 PM）
-- %S	秒(00-59)
-- %s	秒(00-59)
-- %T	时间, 24-小时 (hh:mm:ss)
-- %U	周 (00-53) 星期日是一周的第一天
-- %u	周 (00-53) 星期一是一周的第一天
-- %V	周 (01-53) 星期日是一周的第一天，与 %X 使用
-- %v	周 (01-53) 星期一是一周的第一天，与 %x 使用
-- %W	星期名
-- %w	周的天 （0=星期日, 6=星期六）
-- %X	年，其中的星期日是周的第一天，4 位，与 %V 使用
-- %x	年，其中的星期一是周的第一天，4 位，与 %v 使用
-- %Y	年，4 位
-- %y	年，2 位

-- use date_format
SELECT
    COUNT(IF(w =  0 OR  w = 6,  1, NULL)) AS weekend_cnt,
    COUNT(IF(w != 0 AND w != 6, 1, NULL)) AS working_cnt
FROM 
    (
        SELECT 
            DATE_FORMAT(submit_date,"%w") AS w
        FROM
            Tasks
    ) AS t

-- WEEKDAY
SELECT 
    SUM(WEEKDAY(submit_date) BETWEEN 5 AND 6) AS weekend_cnt,
    SUM(WEEKDAY(submit_date) BETWEEN 0 AND 4) AS working_cnt
FROM 
    Tasks
    