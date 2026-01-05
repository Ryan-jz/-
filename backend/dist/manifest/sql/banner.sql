-- 轮播图表
CREATE TABLE IF NOT EXISTS `banner` (
  `banner_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '轮播图ID',
  `title` varchar(100) DEFAULT '' COMMENT '标题',
  `description` varchar(500) DEFAULT '' COMMENT '描述',
  `media_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '媒体类型（1图片 2视频）',
  `media_url` varchar(500) NOT NULL COMMENT '媒体URL',
  `button_text` varchar(50) DEFAULT '' COMMENT '按钮文字',
  `button_link` varchar(500) DEFAULT '' COMMENT '按钮链接',
  `sort_order` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态（0停用 1启用）',
  `position` varchar(50) DEFAULT 'home' COMMENT '位置（home首页 product产品页）',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`banner_id`),
  KEY `idx_position_status` (`position`, `status`),
  KEY `idx_sort_order` (`sort_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='轮播图表';

INSERT INTO `banner` (`title`, `description`, `media_type`, `media_url`, `button_text`, `button_link`, `sort_order`, `status`, `position`)
VALUES 
('纯净阿尔卑斯盐', '源自2.5亿年前的原始海洋，深藏于巴伐利亚阿尔卑斯山', 1, 'https://images.unsplash.com/photo-1474440692490-2e83ae13ba29?w=1920', '了解更多', '#alpine-salt', 1, 1, 'home'),
('天然调味盐系列', '精选香草与香料，为您的美食增添独特风味', 1, 'https://images.unsplash.com/photo-1505935428862-770b6f24f629?w=1920', '探索产品', '#seasoned-salt', 2, 1, 'home'),
('可持续发展承诺', '保护环境，传承自然的馈赠', 1, 'https://images.unsplash.com/photo-1518843875459-f738682238a6?w=1920', '了解详情', '#sustainability', 3, 1, 'home');
