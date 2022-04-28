-- 1965 · Search for information on teachers who are not within the age range of 20 to 30 years
-- # Description
-- Please write an SQL statement to query the information of teachers who are not between the ages of 20 and 30 in the teacher table teachers.

-- Table definition: Teachers

-- Column name	type	comment
-- ID	int	primary key
-- Name	varchar	lecturer's name
-- Email	varchar	lecturer‘s email
-- Age	int	age of lecturer
-- Country	varchar	lecturer's nationality
-- 1、NOT BETWEEN AND statement is not including the boundary values
-- 2、If there is no query results, nothing will be returned

-- Example
-- Example 1:
-- Table content: Teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic @ gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar @ qq.com	21	CN
-- 3	Western Venom	western.venom @163.com	38	USA
-- 4	Southern Emperor	southern.emperor @ qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- 3	Western Venom	western.venom @163.com	38	USA
-- 5	Linghu Chong		18	CN
-- Example 2:

-- Table content: Teachers

-- id	name	email	age	country
-- 2	Northern Beggar	northern.beggar @ qq.com	21	CN
-- 4	Southern Emperor	southern.emperor @ qq.com	21	JP
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- Because the teachers in the input sample are between 20 and 30 years old, no data are returned.

-- use not between
SELECT
	*
FROM
	teachers
WHERE
	age NOT BETWEEN 20 AND 30

-- use > & <
SELECT
	*
FROM
	teachers
WHERE
	age < 20 OR 
	age > 30