-- 2583 · Delete data trigger message alert
-- # Description
-- Create a DELETE trigger for a table named members that requires an alert message memberId = [id] and message = 'Delete members {[old]}' to be inserted into the reminders table when a piece of data is deleted

-- [id] is the primary key id of the modified record in members
-- [old] in the parameter description: the deleted data, e.g. '[id=1] [name=jack] [email=null] [birthDate=2000-1-1]', please note that there is a space between []
-- Table Definition 1:members (Member Table)

-- Column Name	Type	Comments
-- id	int unsigned	primary key
-- name	varchar	member name
-- email	varchar	member email
-- birthDate	date	date of birth
-- Table Definition 2: reminders(reminders table)

-- column name	type	comment
-- id	int unsigned	primary key
-- message	varchar	message content
-- memberId	int	member id
-- Contact me on wechat to get Amazon、Google requent Interview questions . (wechat id : jiuzhang15)


-- Example
-- Enter data:

-- members table.

-- id	name	email	birthDate
-- 1	'Eastern heretic'	'eastern.heretic@gmail.com'	'2001-02-05'
-- 2	'Northern Beggar'	'northern.beggar@qq.com'	'2005-02-05'
-- 3	'Western Venom'	'western.venom@163.com'	'2012-04-12'
-- 4	'Southern Emperor'	'southern.emperor@qq.com'	'2014-12-01'
-- 5	'Linghu Chong'	None	'2018-11-12'
-- reminders Table.

-- id	message	memberId
-- After running your SQL statement, the table should return.

-- id	memberId	message
-- 1	1	'Delete members {[id=1] [name=Eastern heretic] [email=eastern.heretic@gmail.com] [birthDate=2001-02-05]}'
-- 2	2	'Delete members {[id=2] [name=Northern Beggar] [email=northern.beggar@qq.com] [birthDate=2005-02-05}'
-- 3	3	'Delete members {[id=3] [name=Western Venom] [email=western.venom@163.com] [birthDate=2012-04-12}'
-- 4	4	'Delete members {[id=4] [name=Southern Emperor] [email=southern.emperor@qq.com] [birthDate=2014-12-01}'
-- 5	5	'Delete members {[id=5] [name=Linghu Chong] [email=null] [birthDate=2018-11-12}'

-- Write your SQL here --
DROP TRIGGER IF EXISTS `after_members_delete`;
CREATE TRIGGER `after_members_delete`
AFTER DELETE ON `members`
FOR EACH ROW
BEGIN
    -- 写提醒表 reminders
    INSERT INTO 
        `reminders` (
			`memberId`,
			`message`
        )
    VALUES (
	     old.id,
		-- CONCAT('Delete members {[id=', old.id, '] [name=', old.name, '] [email=', old.email, '] [birthDate=', old.birthDate, ']}')
		-- CONCAT 如果有值为 null 会整体返回 null 
		-- Column 'message' cannot be null
		-- 所以要使用 IFNULL 判断处理
		CONCAT('Delete members {[id=', old.id, '] [name=', old.name, '] [email=', IFNULL(old.email,'null'), '] [birthDate=', old.birthDate, ']}')
    );
END