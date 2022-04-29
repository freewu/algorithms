-- 1971 · Search for teachers whose email is empty
-- Description
-- Please write a SQL statement to query the information of teachers whose email is null in the teacher table teachers.

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality



-- The difference between NULL and an empty string: null means that the email data is empty.

-- Example
-- Example 1:

-- Table content：teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- 5	Linghu Chong		18	CN
-- Example 2:

-- Table content: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	linghu.chong@outlook.com	18	CN
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- There is no record of a teacher with an empty email in output of Example 2, so only the title is shown here, no data.
SELECT	
	*
FROM
	teachers
WHERE
	email IS NULL