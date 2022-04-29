-- 1991 · Count the total number of students for teacher #3
-- # Description
-- Write an SQL statement to count the total number of students with teacher_id of 3 in the courses table courses, and use select_student_sum as the result set column name.

-- Table Definition: courses (courses table)

-- column name	type	comment
-- id	int	primary key
-- name	varchar	course name
-- student_count	int	number of students
-- created_at	datetime	course start time
-- teacher_id	int	teacher id
-- Contact me on wechat to get more FLAMG requent Interview questions . (wechat id : jiuzhang15)


-- The result set column name needs to be renamed.
-- If there is no valid headcount data, NULL is returned.
-- Example
-- Sample 1:

-- Table Contents : courses

-- id	name	student_count	create_time	teacher_id
-- 1	Senior Algorithm	880	2020-06-01	4
-- 2	System Design	1350	2020-07-18	3
-- 3	Django	780	2020-02-29	3
-- 4	Web	340	2020-04-22	4
-- 5	Big Data	700	2020-09-11	1
-- 6	Artificial Intelligence	1660	2018-05-13	3
-- 7	Java P6+	780	2019-01-19	3
-- 8	Data Analysis	500	2019-07-12	1
-- 10	Object Oriented Design	300	2020-08-08	4
-- 12	Dynamic Programming	2000	2018-08-18	1
-- After running your SQL statement, the table should return.

-- select_student_sum
-- 4570
-- Sample 2:

-- Table Contents : courses

-- | id | name | student_count | created_at | teacher_id |
-- | :----: | :----: | :----: | :----: | :----: | :----: |
-- | 2 | Web | 20 | 2021-03-05 | 2 |
-- | 3 | Python | 30 | 2021-08-03 | 2 |
-- | 4 | MySQL | 40 | 2021-06-01 | 4 |

-- After running your SQL statement, the table should return.

-- select_student_sum
-- NULL
-- Sample 2 does not have a teacher_id of 3, so it will return null.
SELECT
	SUM(student_count) AS select_student_sum
FROM
	courses
WHERE
	teacher_id = 3