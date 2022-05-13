-- 2810 · MySQL Stored Procedure INOUT Parameters I
-- Description
-- Write an SQL statement to create a procedure called UpdateTeacherAge to update the age of Linghu Chong.
-- The procedure has two parameters.
--      age: is the INOUT parameter that specifies the age of the teacher
--      inc: is the IN parameter, specifying the number of years to add
-- We will call this procedure to add 10 years to the teacher's age

-- Table definition: teachers (teachers table)
-- columns_name	type	explaination
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
-- @age
-- 28

-- Write your SQL Query here --
CREATE PROCEDURE UpdateTeacherAge(
	INOUT age  INT,
	IN inc  INT -- 增加年岁数
)
BEGIN
	SET age = age + inc;
	UPDATE 
		`teachers`
	SET
		age = age
	WHERE
		name = 'Linghu Chong';
END;