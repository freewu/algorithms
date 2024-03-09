-- 1853. Convert Date Format
-- Table: Days
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | day         | date |
-- +-------------+------+
-- day is the column with unique values for this table.
-- Write a solution to convert each date in Days into a string formatted as "day_name, month_name day, year".
-- Return the result table in any order.
-- The result format is in the following example.

-- Example 1:
-- Input: 
-- Days table:
-- +------------+
-- | day        |
-- +------------+
-- | 2022-04-12 |
-- | 2021-08-09 |
-- | 2020-06-26 |
-- +------------+
-- Output: 
-- +-------------------------+
-- | day                     |
-- +-------------------------+
-- | Tuesday, April 12, 2022 |
-- | Monday, August 9, 2021  |
-- | Friday, June 26, 2020   |
-- +-------------------------+
-- Explanation: Please note that the output is case-sensitive.

-- Create table If Not Exists Days (day date)
-- Truncate table Days
-- insert into Days (day) values ('2022-04-12')
-- insert into Days (day) values ('2021-08-09')
-- insert into Days (day) values ('2020-06-26')

SELECT
    DATE_FORMAT(day,"%W, %M %e, %Y") AS day
FROM 
    Days

-- DATE_FORMAT(date,format)
--     date: 合法的日期
--     format: 规定日期/时间的输出格式 
-- DATE_FORMAT函数可以使用的参数格式
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