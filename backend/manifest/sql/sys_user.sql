-- 系统用户表
CREATE TABLE IF NOT EXISTS `sys_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `user_name` varchar(30) NOT NULL COMMENT '用户账号',
  `nick_name` varchar(30) DEFAULT NULL COMMENT '用户昵称',
  `email` varchar(50) DEFAULT NULL COMMENT '用户邮箱',
  `phonenumber` varchar(11) DEFAULT NULL COMMENT '手机号码',
  `sex` char(1) DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) DEFAULT NULL COMMENT '头像地址',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `status` char(1) DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `idx_user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统用户表';

-- 插入测试用户（用户名: admin, 密码: admin）
INSERT INTO `sys_user` (`user_name`, `nick_name`, `email`, `phonenumber`, `sex`, `password`, `status`, `del_flag`)
VALUES ('admin', '管理员', 'admin@example.com', '13800138000', '0', 'admin', '0', '0')
ON DUPLICATE KEY UPDATE `user_name` = `user_name`;
