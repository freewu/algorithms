-- 2034 · Check the average age of teachers at the end of the specified mailbox
-- # Description
-- Please write a SQL statement to query the average age of teachers who use QQ email(email end with @qq.com), and the returned result column named 'average_teacher_age'.

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The column names that return statistical results need to be changed
-- If the statistics are null, 0 is returned
-- Example
-- Example 1：

-- Table content : teachers
-- ​

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
				
-- After running your SQL statement, the table should return :				
			
-- average_teacher_age				
-- -----------------				
-- 21.0	


-- Example 2：

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com		UK
-- 2	Southern Emperor	southern.emperor@gmail.com		JP
-- 3	Linghu Chong	NULL	NULL	USA
-- After running your SQL statement, the table should return :

-- average_teacher_age
-- All the teachers' ages in sample 2 are empty and no data is entered, so it returns null .

SELECT
	SUM(age) / COUNT(*) AS average_teacher_age
FROM
	teachers
WHERE
	email LIKE '%@qq.com'