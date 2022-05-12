-- 2809 · MySQL Stored Procedure OUT Parameters II
-- # Description
-- Write an SQL statement to create a procedure that returns the number of courses taught by a teacher by teacher id. 
-- The procedure has two parameters.
--      Teacher: is an IN parameter that specifies the teacher id to be returned
--      total: is the OUT parameter that stores the number of courses taught by the teacher
-- Call this procedure to find the number of courses taught by a teacher with teacher id 3

-- Table definition : courses
-- columns_name	type	explaination
-- id	int unsigned	primary key
-- name	varchar	courses' name
-- student_count	int	number of students
-- created_at	date	course creation time
-- teacher_id	int	teacher id

-- Example
-- Input：

-- courses：
-- id	name	student_count	created_at	teacher_id
-- 1	'Advanced Algorithms'	880	'2020-6-1'	4
-- 2	'System Design'	1350	'2020-7-18'	3
-- 3	'Django'	780	'2020-2-29'	3
-- 4	'Web'	340	'2020-4-22'	4
-- 5	'Big Data'	700	'2020-9-11'	1
-- 6	'Artificial Intelligence'	1660	'2018-5-13'	3
-- 7	'Java P6+'	780	'2019-1-19'	3
-- 8	'Data Analysis'	500	'2019-7-12'	1
-- 10	'Object Oriented Design'	300	'2020-8-8'	4
-- 12	'Dynamic Programming'	2000	'2018-8-18'	1

-- Return：
-- @total
-- 4

-- Write your SQL Query here --
CREATE PROCEDURE GetCoursesByTeacher (
	IN Teacher INT,
	OUT total INT
)
BEGIN
    SET total =  (
        SELECT
            	COUNT(*)
        FROM
            	`courses`
        WHERE
            	teacher_id = Teacher
    );
END;

CALL GetCoursesByTeacher(3,@total);

SELECT @total;