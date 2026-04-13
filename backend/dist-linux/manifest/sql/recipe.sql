-- 食谱表
CREATE TABLE IF NOT EXISTS `recipe` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '食谱ID',
  `name` varchar(200) NOT NULL DEFAULT '' COMMENT '食谱名称',
  `subtitle` varchar(500) DEFAULT '' COMMENT '副标题',
  `description` text COMMENT '简介',
  `image` varchar(500) DEFAULT '' COMMENT '主图片',
  `images` text COMMENT '详情图片（JSON数组）',
  `ingredients` text COMMENT '食材列表（JSON数组）',
  `content` longtext COMMENT '制作步骤（富文本HTML）',
  `cooking_time` int(11) DEFAULT 0 COMMENT '烹饪时间（分钟）',
  `difficulty` tinyint(1) DEFAULT 1 COMMENT '难度等级（1简单 2中等 3困难）',
  `servings` int(11) DEFAULT 1 COMMENT '份量（几人份）',
  `product_ids` varchar(500) DEFAULT '' COMMENT '关联产品ID（逗号分隔）',
  `tags` varchar(500) DEFAULT '' COMMENT '标签（逗号分隔）',
  `view_count` int(11) DEFAULT 0 COMMENT '浏览次数',
  `like_count` int(11) DEFAULT 0 COMMENT '点赞次数',
  `sort_order` int(11) DEFAULT 0 COMMENT '排序',
  `status` tinyint(1) DEFAULT 1 COMMENT '状态（0下架 1上架）',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_sort_order` (`sort_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='食谱表';

INSERT INTO `recipe` (`name`, `subtitle`, `description`, `image`, `ingredients`, `content`, `cooking_time`, `difficulty`, `servings`, `product_ids`, `tags`, `sort_order`, `status`)
VALUES 
('喜马拉雅盐烤鸡', '外酥里嫩，盐香四溢', '使用喜马拉雅粉盐腌制的烤鸡，肉质鲜嫩多汁，带有独特的矿物质风味', '/uploads/images/recipe-chicken.jpg', 
'[{"name":"整鸡","amount":"1只"},{"name":"喜马拉雅粉盐","amount":"30g"},{"name":"黑胡椒","amount":"5g"},{"name":"橄榄油","amount":"20ml"},{"name":"迷迭香","amount":"适量"}]',
'<h3>步骤1：准备食材</h3><p>将整鸡清洗干净，沥干水分。</p><h3>步骤2：腌制</h3><p>用喜马拉雅粉盐、黑胡椒均匀涂抹在鸡身内外，腌制2小时。</p><h3>步骤3：烤制</h3><p>烤箱预热至200度，烤制60分钟，期间刷油2-3次。</p><h3>步骤4：完成</h3><p>取出后静置10分钟，切块装盘即可。</p>',
90, 2, 4, '1,2', '烤鸡,主菜,家常菜', 1, 1),

('海盐柠檬烤鱼', '清新爽口，营养健康', '采用天然海盐和新鲜柠檬调味的烤鱼，保留鱼肉的鲜美', '/uploads/images/recipe-fish.jpg',
'[{"name":"鲈鱼","amount":"1条"},{"name":"海盐","amount":"20g"},{"name":"柠檬","amount":"2个"},{"name":"大蒜","amount":"5瓣"},{"name":"香菜","amount":"适量"}]',
'<h3>步骤1：处理鱼</h3><p>鲈鱼去鳞去内脏，清洗干净，在鱼身两侧划几刀。</p><h3>步骤2：调味</h3><p>用海盐、柠檬汁、蒜末腌制30分钟。</p><h3>步骤3：烤制</h3><p>烤箱180度烤25分钟，翻面再烤10分钟。</p><h3>步骤4：装饰</h3><p>撒上香菜和柠檬片即可。</p>',
60, 1, 2, '3', '烤鱼,海鲜,健康', 2, 1);
