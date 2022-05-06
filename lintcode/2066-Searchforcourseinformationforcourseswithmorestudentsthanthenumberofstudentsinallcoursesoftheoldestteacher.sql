-- 2066 · Search for course information for courses with more students than the number of students in all courses of the oldest teacher
-- # Description
-- Write an SQL statement that queries the student count of courses taught by the oldest teacher 
-- from the teacher table teachers and the course table courses, 
-- and finally returns the information of courses in which the number of students exceeds all those courses.

-- Table Definition 1: courses

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course name
-- student_count	int	total students
-- created_at	date	creation time
-- teacher_id	int	instructor id
-- Table Definition 2: teachers

-- column_name	type	comment
-- id	int unsigned	primary key
-- name	varchar	teacher's name
-- email	varchar	teacher's email
-- age	int	teacher's age
-- country	varchar	teacher's nationality

-- The query returns column names that need to be case-sensitive to the sample output.
-- There may be more than one oldest teacher.
-- If the number of students present in the input data is NULL, the data is skipped.
-- If the query does not return any results, nothing is returned.
-- Example
-- Example I:

-- Table Contents 1: courses
-- | id   | name                    | student_count | created_at | teacher_id |
-- | ---- | ----------------------- | ------------- | ---------- | ---------- |
-- | 1    | Senior Algorithm        | 880           | 2020-06-01 | 4          |
-- | 2    | System Design           | 1350          | 2020-07-18 | 3          |
-- | 3    | Django                  | 780           | 2020-02-29 | 3          |
-- | 4    | Web                     | 340           | 2020-04-22 | 4          |
-- | 5    | Big Data                | 700           | 2020-09-11 | 1          |
-- | 6    | Artificial Intelligence | 1660          | 2018-05-13 | 3          |
-- | 7    | Java P6+                | 780           | 2019-01-19 | 3          |
-- | 8    | Data Analysis           | 500           | 2019-07-12 | 1          |
-- | 10   | Object Oriented Design  | 300           | 2020-08-08 | 4          |
-- | 12   | Dynamic Programming     | 2000          | 2018-08-18 | 1          |

-- Table of Contents 2: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the courses table should return.

-- id	name	student_count	created_at	teacher_id
-- 12	Dynamic Programming	2000	2018-08-18	1
-- Example 2:				
-- Table Contents 1: courses

-- id	name	student_count	created_at	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- Table of Contents 2: teachers

-- id	name	email	age	country
-- 1	Eastern Heretic	eastern.heretic@gmail.com	20	UK
-- 2	Northern Beggar	northern.beggar@qq.com	21	CN
-- 3	Western Venom	western.venom@163.com	28	USA
-- 4	Southern Emperor	southern.emperor@qq.com	21	JP
-- 5	Linghu Chong	NULL	18	CN
-- After running your SQL statement, the courses table should return.

-- id	name	student_count	created_at	teacher_id
-- Because there is no eligible data in the input sample, only the title is shown here, no data.

SELECT
	*
FROM
	courses 
WHERE
	student_count > (
		SELECT -- 最大的学生数
			MAX(c.student_count) 
		FROM
			courses AS c,
		 	teachers AS t
		WHERE
			c.teacher_id = t.id AND 
			t.age = ( -- 最大的年龄
				SELECT
					MAX(age)
				FROM
					teachers
			)
	)