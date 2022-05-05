-- 2024 · Query the course name and year of creation date of all course schedules
-- # Description
-- Write a SQL statement to query the name and creation year of the course from the course table courses and alias created_at to created_year.

-- Table definition: courses（courses table）

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	datetime	start time
-- teacher_id	int	teacher id

-- The column names returned by the query need to be the same case as the sample output.

-- Example
-- Example 1

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01 09:10:12	4
-- 2	System Design	1350	2020-07-18 10:11:12	3
-- 3	Django	780	2020-02-29 12:10:12	3
-- 4	Web	340	2020-04-22 13:01:12	4
-- 5	Big Data	700	2020-09-11 16:01:12	1
-- 6	Artificial Intelligence	1660	2018-05-13 18:12:30	3
-- 7	Java P6+	780	2019-01-19 13:31:12	3
-- 8	Data Analysis	500	2019-07-12 13:01:12	1
-- 10	Object Oriented Design	300	2020-08-08 13:01:12	4
-- 12	Dynamic Programming	2000	2018-08-18 20:01:12	1
-- After running your SQL statement, the table should return.

-- name	created_year
-- Senior Algorithm	2020
-- System Design	2020
-- Django	2020
-- Web	2020
-- Big Data	2020
-- Artificial Intelligence	2018
-- Java P6+	2019
-- Data Analysis	2019
-- Object Oriented Design	2020
-- Dynamic Programming	2018
-- Example 2

-- Table Contents : courses

-- id	name	student_count	created_at	teacher_id
-- 1	null	880	null	4
-- 2	null	1350	null	3
-- 3	null	780	null	3
-- 4	null	340	null	4
-- 5	null	700	null	1
-- 6	null	1660	null	3
-- 7	null	780	null	3
-- 8	null	500	null	1
-- 10	'IDE'	300	null	4
-- 12	null	2000	'2018-08-18 20:01:12'	1
-- After running your SQL statement, the table should return.

-- name	created_year
-- null	null
-- null	null
-- null	null
-- null	null
-- null	null
-- null	null
-- null	null
-- null	null
-- 'IDE'	null
-- null	2018
-- 样例二中数值如果有 null 值，则会返回 null 值。

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

-- DATE_FORMAT(NOW(),'%b %d %Y %h:%i %p') // Dec 29 2008 11:45 PM
-- DATE_FORMAT(NOW(),'%m-%d-%Y') // 12-29-2008
-- DATE_FORMAT(NOW(),'%d %b %y') // 29 Dec 08
-- DATE_FORMAT(NOW(),'%d %b %Y %T:%f') // 29 Dec 2008 16:25:46.635

-- use substr
SELECT
	name,
	CAST(SUBSTR(created_at,1,4) AS SIGNED) AS created_year
FROM
	courses

-- use date_format
SELECT
	name,
	CAST(DATE_FORMAT(created_at,'%Y') AS SIGNED) AS created_year
FROM
	courses

