-- 1955 · Query teachers who do not meet the conditions
-- # Description
-- Please write a SQL statement to query the information of all teachers who are not Chinese and not older than 20 years old(excluding 20) in the table teachers .

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Tip:

-- Not older than 20 years old, including 20 years old
-- Nationality is China'cn'
-- Notice that it is 'not' in the question stem
-- If the query cannot find the result, nothing is returned
-- Example
-- Example 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- Example 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	22	CN
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	CN
-- 4	Southern Emperor	southern.emperor@qq.com	21	CN
-- 5	Linghu Chong		21	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- Because there is no data to be queried in the data, only the title is displayed here, no data
SELECT
	*
FROM
	teachers
WHERE
	id NOT IN ( -- 找出 大于 20 岁的中国人 其它就是结果了
		SELECT 
			id
		FROM
			teachers
		WHERE
			country = 'CN' AND
			age > 20
	)