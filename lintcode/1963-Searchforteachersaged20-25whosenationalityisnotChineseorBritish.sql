-- 1963 · Search for teachers aged 20~25 whose nationality is not Chinese or British
-- # Description
-- Write an SQL statement to query the teachers between the ages of 20 (including) and 25 (including) who are not Chinese or British in the teacher table teachers , and finally return all the information of the queried teachers.

-- The teachers table is defined as follows：

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Tip：

-- Between the ages of 20 and 25, including 20 and 25
-- Nationality is neither Chinese nor British
-- If there is no query result, nothing will be returned
-- Example
-- Example 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return :

-- id	name	email	age	country
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- Example 2:
-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Linghu Chong		18	CN
-- After running your SQL statement, the table should return :

-- id	name	email	age	country
-- Because there are no teachers aged 20 to 25 whose nationality is not Chinese or British, Sample Output 2 is an empty table, and only the headings are shown here.

-- use BETWEEN
SELECT
	* 
FROM
	teachers 
WHERE
	age BETWEEN 20 AND 25 AND
	country NOT IN ('CN','UK')

-- use <= & >=
SELECT
	* 
FROM
	teachers 
WHERE
	age >= 20 AND 
	age <= 25 AND
	country NOT IN ('CN','UK')