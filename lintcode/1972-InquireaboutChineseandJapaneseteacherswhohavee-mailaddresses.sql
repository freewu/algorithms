-- 1972 · Inquire about Chinese and Japanese teachers who have e-mail addresses
-- # Description
-- Please write a SQL statement to query the information of all teachers whose nationality is 'CN' or 'JP' and the emailemail is not empty in the teacher table teachers.

-- Table definition : teachers

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- The inquiries are for teachers from China or Japan
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
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- Example 2:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	18	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	USA
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	UK
-- 5	Linghu Chong		18	USA
-- After running your SQL statement, the table should return:

-- id	name	email	age	country
-- Because there is no data that meets the conditions in the input sample, only the title is shown here, and there is no data.

SELECT
	*
FROM
	teachers
WHERE
	country IN ('CN','JP') AND
	email IS NOT NULL