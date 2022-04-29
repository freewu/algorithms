-- 1996 · Determining whether a teacher has a mailbox
-- # Description
-- Write an SQL statement to determine if a teacher in the teachers table teachers has a mailbox, 
-- and finally return the teacher's name and email, and the result of the determination using the functions ISNULL, IFNULL, and COALESCE.

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher‘s name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher’s nationality
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)

-- Note the use of different NULL functions
-- Different NULL functions return different values for null values, so pay attention to learning the NULL functions by comparison

-- Eample 1 ：

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return :

-- name	email	ISNULL(email)	IFNULL(email,0)	COALESCE(email,0)
-- Eastern Heretic	eastern.heretic@gmail.com	0	eastern.heretic@gmail.com	eastern.heretic@gmail.com
-- Northern Beggar	northern.beggar@qq.com	0	northern.beggar@qq.com	northern.beggar@qq.com
-- Western Venom	western.venom@163.com	0	western.venom@163.com	western.venom@163.com
-- Southern Emperor	southern.emperor@qq.com	0	southern.emperor@qq.com	southern.emperor@qq.com
-- Linghu Chong		1	0	0
-- Eample 2 ：

-- Table content : teachers

-- id	name	email	age	country
-- 1	White	north.heretic@gmail.com	20	UK
-- 3	Emperor	qw.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return :

-- name	email	ISNULL(email)	IFNULL(email,0)	COALESCE(email,0)
-- White	north.heretic@gmail.com	0	north.heretic@gmail.com	north.heretic@gmail.com
-- Emperor	qw.emperor@qq.com	0	qw.emperor@qq.com	qw.emperor@qq.com
-- Linghu Chong		1	0	0
SELECT
	name,
	email,
	ISNULL(email) AS isnull_email,
	IFNULL(email,0) AS ifnull_email,
	COALESCE(email,0) AS coalesce_email
FROM
	teachers

-- COALESCE(expression_1, expression_2, ...,expression_n)
-- 依次参考各参数表达式，遇到非null值即停止并返回该值。如果所有的表达式都是空值，最终将返回一个空值。使用COALESCE在于大部分包含空值的表达式最终将返回空值。