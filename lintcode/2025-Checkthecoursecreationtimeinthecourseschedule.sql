-- 2025 · Check the course creation time in the course schedule
-- # Description
-- Question Description: Write a SQL statement to query the creation time of courses in the course table which output in 'hour:minute:second' format, and the returned column named created_at.

-- Table definition: courses (course table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	student count
-- created_at	datetime	course creation time
-- teacher_id	int unsigned	teacher id

-- The column names returned by the query need to match the case of the column names output by the sample
-- If there is a null value in the SELECT value, NULL will be returned.
-- Example
-- Sample I

-- Table content: courses

-- | id   | name                    | student_count | created_at          | teacher_id |
-- | ---- | ----------------------- | ------------- | ------------------- | ---------- |
-- | 1    | Senior Algorithm        | 880           | 2020-06-01 09:03:12 | 4          |
-- | 2    | System Design           | 1350          | 2020-07-18 10:03:12 | 3          |
-- | 3    | Django                  | 780           | 2020-02-29 12:03:12 | 3          |
-- | 4    | Web                     | 340           | 2020-04-22 13:03:12 | 4          |
-- | 5    | Big Data                | 700           | 2020-09-11 16:03:12 | 1          |
-- | 6    | Artificial Intelligence | 1660          | 2018-05-13 18:03:30 | 3          |
-- | 7    | Java P6+                | 780           | 2019-01-19 13:03:12 | 3          |
-- | 8    | Data Analysis           | 500           | 2019-07-12 13:03:12 | 1          |
-- | 10   | Object Oriented Design  | 300           | 2020-08-08 13:03:12 | 4          |
-- | 12   | Dynamic Programming     | 2000          | 2018-08-18 20:03:12 | 1          |

-- After running your SQL statement, the table should return.

-- created_at
-- 09:03:12
-- 10:03:12
-- 12:03:12
-- 13:03:12
-- 16:03:12
-- 18:03:12
-- 13:03:12
-- 13:03:12
-- 13:03:12
-- 20:03:12
-- 样例二

-- 表内容：courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	NULL	4
-- 2	System Design	1350	NULL	3
-- 3	Django	780	NULL	3
-- 4	Web	340	NULL	4
-- 5	Big Data	700	NULL	1
-- 6	Artificial Intelligence	1660	NULL	3
-- 7	Java P6+	780	NULL	3
-- 8	Data Analysis	500	NULL	1
-- 10	Object Oriented Design	300	NULL	4
-- 12	Dynamic Programming	2000	NULL	1
-- 在运行你的 SQL 语句之后，表应返回：

-- created_at
-- NULL
-- NULL
-- NULL
-- NULL
-- NULL
-- NULL
-- NULL
-- NULL
-- NULL
-- NULL
-- NULL
-- NULL

SELECT
	DATE_FORMAT(created_at,'%H:%i:%s') AS created_at 
FROM
	courses

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