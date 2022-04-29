-- 1985 · Number of teachers aged 20 to 28 who are Chinese and British nationals
-- # Description
-- Write an SQL statement to count the number of teachers who are Chinese or British with the age between 20 and 28 years old in the teachers table, and return the statistics with the column namedteacher_count.

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- Age between 20 and 28 years old, including 20 and 28 years old
-- The column name that returns the statistics needs to be changed
-- Returns 0 if no data is counted

-- Example 1：
-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return :

-- teacher_count
-- 2

-- Example 2：
-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	29	UK
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return :

-- teacher_count
-- 0
-- There are no eligible data in sample 2, so the statistic result is 0

SELECT
	COUNT(*) AS teacher_count
FROM
	teachers
WHERE
	age BETWEEN 20 AND 28 AND
	country IN ('CN','UK')