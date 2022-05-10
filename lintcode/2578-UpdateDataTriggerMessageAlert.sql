-- 2578 · Update data trigger message alert
-- -- # Description
-- Create an UPDATE trigger for a table named members that requires the data to be modified to record 
-- the modified message memberId = [id], message = 'Update {[old]} To {[new]}' to the reminders table

-- [id] is the primary key id of the modified record in `members
-- [old] and [new]: only the field that has been modified is displayed, e.g. name is updated from Tom to Jack, [old] is displayed as [name=tom], [new] is displayed as [name=jack]
-- Table Definition 1:members (members table)

-- column name	type	comment
-- id	int unsigned	primary key
-- name	varchar	member name
-- email	varchar	member email
-- birthDate	date	date of birth
-- Table Definition 2: reminders(reminders table)

-- column name	type	comment
-- id	int unsigned	primary key
-- message	varchar	message content
-- memberId	int	member id

-- please note that there is a space after [].

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

-- | id | memberId | message                                                                                                                                                       |
-- | :- | :------- | :-------------------------------------------------------------------------------------------------------------------- --------------------------------------- |
-- | 1  | 4        | 'Update {[name=Southern Emperor] [email=southern.emperor@qq.com] [birthDate=2014-12-01] } To {[name=tommer] [email=tommer@136.com] [ birthDate=2011-02-15] }' |
-- | 2  | 3        | 'Update {[name=Western Venom] [birthDate=2012-04-12] } To {[name=july] [birthDate=2001-02-15] }'                                                              |
-- | 3  | 2        | 'Update {[birthDate=2005-02-05] } To {[birthDate=1999-01-17] }'                                                                                               |
-- | 4  | 5        | 'Update {[birthDate=2018-11-12] } To {[birthDate=2000-01-23] }'                                                                                               |

-- Write your SQL here --
DROP TRIGGER IF EXISTS `after_members_update`;
CREATE TRIGGER `after_members_update`
AFTER UPDATE ON `members`
FOR EACH ROW
BEGIN
	-- 没有对象的概念只能拼字符串了
	DECLARE new_str varchar(1024);
	DECLARE old_str varchar(1024);
	SET new_str = ''; -- 一定要做初始化 上面只是声明变量
	SET old_str = '';

	-- name 有变动
	IF old.name != new.name THEN
		SET new_str = CONCAT(new_str, '[name=', new.name ,'] ');
		SET old_str = CONCAT(old_str, '[name=', old.name ,'] ');
	END IF;
	-- email 有变动
	IF old.email != new.email THEN
		SET new_str = CONCAT(new_str, '[email=', new.email, '] ');
		SET old_str = CONCAT(old_str, '[email=', old.email, '] ');
	END IF;
	-- birthDate 有变动
	IF old.birthDate != new.birthDate THEN
		SET new_str = CONCAT(new_str, '[birthDate=', new.birthDate,'] ');
		SET old_str = CONCAT(old_str, '[birthDate=', old.birthDate,'] ');
	END IF;

	-- 添加 提醒消息
	--IF new_str IS NOT NULL THEN
		INSERT INTO
			`reminders`(
				`memberId`,
				`message`
			)
		VALUES (
			new.id,
			CONCAT('Update {', old_str ,' } To {', new_str, '}')
		);
	--END IF;
END