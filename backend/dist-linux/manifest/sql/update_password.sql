-- 更新现有用户的密码为加密后的哈希值
-- 原密码: admin
-- 加密后: $2a$10$gYn0EynARjJebHThLfUSVOypTKjYpyucFIiC0bRJTxOeAPFQ.hJFm

UPDATE `sys_user` 
SET `password` = '$2a$10$gYn0EynARjJebHThLfUSVOypTKjYpyucFIiC0bRJTxOeAPFQ.hJFm'
WHERE `user_name` = 'admin';
