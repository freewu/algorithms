-- 2723 · Showing the definition of a view
-- # Description
-- We have a view v_teachers, but now we want to display the definition of this view, write SQL statements to do so

-- Example
-- The result should be returned after executing your SQL statement:

-- View	Create View	character_set_client	collation_connection
-- 'v_teachers'	"CREATE ALGORITHM=UNDEFINED DEFINER=lintcode@% SQL SECURITY DEFINER VIEW v_teachers AS select teachers.name AS name,teachers.age AS age,teachers.country AS country from teachers where (teachers.country in ('CN','USA','UK'))"	'utf8mb4'	'utf8mb4_general_ci'

-- Write your SQL here --

-- show create view
SHOW CREATE VIEW `v_teachers`;

-- select information_schema.views
--SELECT * FROM information_schema.views  WHERE TABLE_NAME='v_teachers';

-- # 查看视图
-- ## 1 使用show create view语句查看视图信息 [ show create view ]
-- mysql> show create view v_F_players\G;
-- *************************** 1. row ***************************
--                 View: v_F_players
--          Create View: CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `v_F_players` AS select `PLAYERS`.`PLAYERNO` AS `编号`,`PLAYERS`.`NAME` AS `名字`,`PLAYERS`.`SEX` AS `性别`,`PLAYERS`.`PHONENO` AS `电话` from `PLAYERS` where (`PLAYERS`.`SEX` = 'F') WITH CASCADED CHECK OPTION
-- character_set_client: utf8
-- collation_connection: utf8_general_ci
-- 1 row in set (0.00 sec)
---
-- ## 2 视图一旦创建完毕，就可以像一个普通表那样使用，视图主要用来查询
-- mysql> select * from view_name;
--
-- ## 3 有关视图的信息记录在information_schema数据库中的 views 表中 [ select information_schema.views ]
-- mysql> select * from information_schema.views 
--     -> where TABLE_NAME='v_F_players'\G;
-- *************************** 1. row ***************************
--        TABLE_CATALOG: def
--         TABLE_SCHEMA: TENNIS
--           TABLE_NAME: v_F_players
--      VIEW_DEFINITION: select `TENNIS`.`PLAYERS`.`PLAYERNO` AS `编号`,`TENNIS`.`PLAYERS`.`NAME` AS `名字`,`TENNIS`.`PLAYERS`.`SEX` AS `性别`,`TENNIS`.`PLAYERS`.`PHONENO` AS `电话` from `TENNIS`.`PLAYERS` where (`TENNIS`.`PLAYERS`.`SEX` = 'F')
--         CHECK_OPTION: CASCADED
--         IS_UPDATABLE: YES
--              DEFINER: root@localhost
--        SECURITY_TYPE: DEFINER
-- CHARACTER_SET_CLIENT: utf8
-- COLLATION_CONNECTION: utf8_general_ci
-- 1 row in set (0.00 sec)    