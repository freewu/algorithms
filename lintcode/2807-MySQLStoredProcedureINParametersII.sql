-- 2807 · MySQL Stored Procedure IN Parameters II
-- # Description
-- Write an SQL statement to create a procedure that finds all teachers 
-- with the age specified by the IN parameter teacherAge and use the procedure to query for teachers with the age of 21

-- Table definition: teachers (teachers table)

-- columns_name	type	explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- Example
-- Sample 1:

-- Table content : teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong		18	CN
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- Example 2:

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	23	JP
-- After running your SQL statement, the table should return.

-- id	name	email	age	country
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN

-- Write your SQL Query here --
CREATE PROCEDURE GetTeachersByAge (
	IN teacherAge  INT
)
BEGIN
	SELECT
		*
	FROM
		`teachers`
	WHERE
		age = teacherAge;
END;

CALL GetTeachersByAge(21);