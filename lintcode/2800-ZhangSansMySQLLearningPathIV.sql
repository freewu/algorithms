-- 2800 · Zhang San's MySQL Learning Path (IV)
-- # Description
-- "The moon is so beautiful tonight~", Zhang San was enjoying the moon alone, 
-- this wonderful moment can only be felt after the end of the study, 
-- today Zhang San just started to learn the storage process, he found that this is a good thing ah, 
-- after learning and then enjoy a moon is really beautiful. At this time, a teacher gave him a phone call,
-- directly ruined his wonderful night. A teacher said to help him the whole live, 
-- the department is now ready to reduce the pressure on older teachers, 
-- on more than 1 course (at least 2 courses) is the active teacher,
-- but now some young teachers like to be lazy, only 1 course or no teaching,
-- so to will be the oldest teacher in the active teacher to teach all the courses taught 
-- by the youngest teachers (In order of name if age is the same) inactive teachers,
-- until the whole teacher The team looks younger and reduce the pressure on older teachers, 
-- and the criterion is that the average age of active teachers should not be greater than the average age of all teachers,
-- and this work will be done regularly, because of the busy business, 
-- and then throw Zhang San 2 tables and ask Zhang San to help him do, teachers table records the information of teachers,
-- and courses table records the course information, associated id for teacher_id. 
-- Zhang San wants to use a SQL Procedure to solve this problem, how should he write it?

-- Table Definition 1: teachers (teachers table)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	Lecturer's name
-- email	varchar	Instructor's email
-- age	int	lecturer's age
-- country	varchar	tutor's nationality

-- Table Definition 2: courses (Course List)
-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	course_name
-- student_count	int	Total number of students attending a class
-- created_at	date	Course creation time
-- teacher_id	int unsigned	instructor v_best_teachers

-- Example
-- Input data.

-- courses table:
-- id	name	student_count	created_at	teacher_id
-- 1	'Advanced Algorithms'	380	'2020-6-1 09:03:12'	7
-- 2	'System Design'	1350	'2020-7-18 10:03:12'	5
-- 3	'Django'	480	'2020-2-29 12:03:12'	5
-- 4	'Web'	940	'2020-4-22 13:03:12'	6
-- 5	'Big Data'	2103	'2020-9-11 16:03:12'	9

-- teachers table:
-- id	name	email	age	country
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	20	'UK'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	21	'CN'
-- 3	'Western Venom'	'western.venom@163.com'	28	'USA'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	21	'JP'
-- 5	'Jiminer'	None	48	'RS'
-- 6	'Beggar'	None	35	'CN'
-- 7	'Jack Turism'	None	26	'RS'
-- 8	'Linghu Chong'	None	25	'JP'
-- 9	'Tom Hertic'	None	46	'UK'
-- 10	'Youker'	None	28	'CN'

-- Return results.
-- course_name	teacher_id	age
-- 'Advanced Algorithms'	7	26
-- 'Big Data'	9	46
-- 'Django'	1	20
-- 'System Design'	1	20
-- 'Web'	6	35

-- Write your SQL here --
create procedure youngTeachers ()
begin
    declare oldt int default 0;
    declare tid int default 0;
    declare actnum int default 0;
    declare notactnum int default 0;

    loop1:loop
        --活跃教师的平均年龄 
        select avg(age) into actnum from teachers as t 
        left join (select teacher_id,count(*) as num from courses group by teacher_id) as c 
        on t.id=c.teacher_id where c.num>1;
        --所有教师的平均年龄
        select avg(age) into notactnum from teachers;
        
        --判断年龄是否符合要求
        if actnum<=notactnum then
            leave loop1;
        end if;
        
        --找出不活跃教师中的最年轻的
        select  t.id into oldt from courses as c 
        left join   teachers as t 
        on t.id=c.teacher_id
        order by t.age desc
        limit 1;
        
        --找出年龄最大的老师
        select  t.id into tid
        from    teachers as t 
        left join    (select teacher_id,count(*) as num from courses group by teacher_id) as c
        on t.id=c.teacher_id 
        where c.num<=1 or c.num is null
        order by age asc,t.name asc
        limit 1;
        --更新数据
        update courses 
        set teacher_id=tid where teacher_id=oldt;
    end loop;
end;
call youngTeachers ();