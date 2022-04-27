-- 1928 · Analysis of Online Class I
-- # Description
-- The online_class_situations table shows the behavioral activities of some students in online classes.
-- Each row of data records the number of courses (may be 0) that a student has listened to after logging in to the course with the same device on the same day before quitting the online course.
-- Write a SQL statement to query the date each student first logged into the platform to attend a class.

-- Table definition: online_class_situations

-- columns_name	type	explaination
-- student_id	int	student's id
-- device_id	int	device's id
-- date	date	Class date of the course
-- course_number	int	course number
-- The primary key of the table is (student_id, date) combined primary key
-- Please note that the returned result column name is: student_id, earliest_course_date

-- Example
-- Example 1:

-- Table content: online_class_situations

-- student_id	device_id	date	course_number
-- 1	2	2020-03-01	5
-- 1	2	2020-04-02	6
-- 2	3	2020-05-25	1
-- 3	1	2020-03-02	0
-- 3	4	2020-12-03	5
-- After running your SQL statement, the table should return:

-- student_id	earliest_course_date
-- 1	2020-03-01
-- 2	2020-05-25
-- 3	2020-12-03
-- Example 2:

-- Table content: online_class_situations

-- student_id	device_id	date	course_number
-- 1	2	2020-03-01	5
-- 1	2	2020-04-02	6
-- 3	1	2020-03-02	3
-- 2	4	2020-12-19	2
-- After running your SQL statement, the table should return:

-- student_id	earliest_course_date
-- 1	2020-03-01
-- 2	2020-12-19
-- 3	2020-03-02

SELECT 
	r.student_id AS student_id,
	MIN(r.date) AS earliest_course_date
FROM 
	online_class_situations AS r
WHERE 
	r.course_number > 0
GROUP BY r.student_id