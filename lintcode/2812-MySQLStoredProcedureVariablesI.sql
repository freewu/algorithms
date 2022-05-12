-- 2812 · MySQL Stored Procedure Variables I
-- Description
-- Write an SQL statement to create a procedure called GetTotalTeacher, 
-- declare a variable totalTeacher with a default value of 0, 
-- and assign the number of teachers in the teachers table to this variable

-- Table definition: teachers (teachers table)
-- columns_name	type explaination
-- id	int	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- Example
-- Input：

-- teachers：

-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Linghu Chong'	None	18	'CN'
-- Return：

-- totalTeacher
-- 5

-- Write your SQL Query here --
CREATE PROCEDURE GetTotalTeacher(
)
BEGIN
	DECLARE totalTeacher INT DEFAULT 0;
	SET totalTeacher = (SELECT COUNT(*) FROM `teachers`);
	SELECT totalTeacher;
END;