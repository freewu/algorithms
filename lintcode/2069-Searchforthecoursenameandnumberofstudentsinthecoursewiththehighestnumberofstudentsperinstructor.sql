-- 2069 · Search for the course name and number of students in the course with the highest number of students per instructor
-- # Description
-- Write an SQL statement to query the course name name 
-- and the number of students student_count for the course with the most students taught by each teacher from the course table courses.

-- Table Definition: courses (Course Table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	course creation time
-- teacher_id	int unsigned	teacher id

-- The column name returned by the query needs to be the same as the case of the column name output by the sample.
-- If the input data is null, NULL is returned.
-- If there are multiple courses of one same teacher with the most students , return all of them.
-- Example
-- Example I

-- Table Contents: courses (Course Table)

-- id	name	student_count	created_at	teacher_id
-- 1	Advanced Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Inteliggence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, the table should return.

-- name	student_count
-- Advanced Algorithms	880
-- Artificial Intelligence	1660
-- Dynamic Programming	2000
-- Example II

-- Table Contents: courses (Course Table)

-- id	name	student_count	created_at	teacher_id
-- 1	Web	800	2019-8-9	3
-- 2	Database	1440	2018-10-8	1
-- 3	cloud computing	850	2020-4-5	2
-- 4	C++	970	2020-5-28	4
-- 5	virtual reality	970	2020-11-21	4
-- After running your SQL statement, the table should return.

-- name	student_count
-- Web	800
-- Database	1440
-- cloud computing	850
-- C++	970
-- virtual reality	970
-- The instructor with teacher id 4 in Example 2 has an equal and maximum number of students in both courses, so both courses return

SELECT
	c.name,
	c.student_count
FROM
	courses AS c,
	(-- 取出最大的学生数 & 老师编号
		SELECT
			teacher_id,
			MAX(student_count) AS student_count
		FROM
			courses
		GROUP BY 
			teacher_id
	) AS p
WHERE
	c.teacher_id = p.teacher_id AND
	c.student_count = p.student_count