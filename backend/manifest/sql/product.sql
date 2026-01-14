CREATE TABLE `product_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '分类名称',
  `slug` varchar(100) NOT NULL COMMENT '分类标识',
  `description` text COMMENT '分类描述',
  `image` varchar(500) DEFAULT '' COMMENT '分类图片',
  `sort_order` int(11) DEFAULT 0 COMMENT '排序',
  `status` tinyint(1) DEFAULT 1 COMMENT '状态：1启用 0禁用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品分类表';

CREATE TABLE `product` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '产品ID',
  `category_id` int(11) unsigned NOT NULL COMMENT '分类ID',
  `category_ids` varchar(255) DEFAULT '' COMMENT '分类ID列表(逗号分隔)',
  `name` varchar(200) NOT NULL DEFAULT '' COMMENT '产品名称',
  `name_en` varchar(200) DEFAULT '' COMMENT '产品英文名称',
  `subtitle` varchar(500) DEFAULT '' COMMENT '副标题',
  `description` text COMMENT '产品描述',
  `image` varchar(500) DEFAULT '' COMMENT '主图片',
  `images` text COMMENT '产品图片集(JSON数组)',
  `price` decimal(10,2) DEFAULT 0.00 COMMENT '价格',
  `stock` int(11) DEFAULT 0 COMMENT '库存',
  `weight` varchar(50) DEFAULT '' COMMENT '重量/规格',
  `ingredients` text COMMENT '成分说明',
  `nutrition` text COMMENT '营养信息(JSON)',
  `usage` text COMMENT '使用方法',
  `features` text COMMENT '产品特点(JSON数组)',
  `organic_cert` varchar(500) DEFAULT '' COMMENT '有机认证图标',
  `recycling_info` text COMMENT '回收信息',
  `purchase_link` varchar(500) DEFAULT '' COMMENT '购买链接',
  `allergen_info` text COMMENT '过敏原信息',
  `storage_info` varchar(500) DEFAULT '' COMMENT '储存信息',
  `origin` varchar(200) DEFAULT '' COMMENT '产地',
  `sort_order` int(11) DEFAULT 0 COMMENT '排序',
  `status` tinyint(1) DEFAULT 1 COMMENT '状态：1上架 0下架',
  `view_count` int(11) DEFAULT 0 COMMENT '浏览次数',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品表';

-- 添加新字段的 ALTER TABLE 语句（如果表已存在）
 

INSERT INTO `product_category` (`name`, `slug`, `description`, `sort_order`, `status`) VALUES
('调味盐系列', 'seasoned-salt', 'BBQ调味盐、香草调味盐等特色调味盐产品', 1, 1),
('阿尔卑斯盐', 'alpine-salt', '来自巴伐利亚阿尔卑斯山的纯净盐产品', 2, 1),
('食用盐', 'table-salt', '日常食用盐系列产品', 3, 1),
('特色盐', 'specialty-salt', '喜马拉雅粉盐、黑盐等特色盐产品', 4, 1);

INSERT INTO `product` (`category_id`, `name`, `subtitle`, `description`, `image`, `price`, `stock`, `weight`, `sort_order`, `status`) VALUES
(2, '阿尔卑斯粗盐', '经典粗盐，适合烹饪和调味', '巴特赖兴哈勒阿尔卑斯粗盐产自纯净的高山盐水，保留了天然的矿物质和微量元素。粗盐颗粒适合烹饪时使用，能够更好地渗透食材，带来纯正的咸味。', '@/assets/images/02.png', 15.90, 1000, '500g', 1, 1),
(2, '阿尔卑斯细盐', '细腻口感，日常调味首选', '细盐颗粒更加细腻，易于溶解，适合日常烹饪和餐桌调味使用。采用传统蒸发工艺制成，保持了阿尔卑斯盐的纯正风味。', '@/assets/images/02.png', 12.90, 1500, '250g', 2, 1),
(2, 'AlpenJodSalz + 碘化物', '添加碘元素，呵护甲状腺健康', '在纯正阿尔卑斯盐的基础上添加碘元素，有助于预防碘缺乏症，保护甲状腺健康。适合全家日常使用。', '@/assets/images/02.png', 13.90, 1200, '500g', 3, 1),
(2, 'AlpenJodSalz + 氟化物 + Folsäure', '多重营养，全面呵护', '添加氟化物和叶酸的阿尔卑斯盐，氟化物有助于牙齿健康，叶酸对孕妇和备孕人群尤为重要。', '@/assets/images/02.png', 16.90, 800, '500g', 4, 1),
(2, '阿尔卑斯盐 + 碘', '经典加碘配方', '经典的加碘食用盐，满足日常碘元素需求，适合各种烹饪场景。', '@/assets/images/02.png', 11.90, 2000, '500g', 5, 1),
(2, '阿尔卑斯盐袋装', '经济实惠，家庭装', '大容量袋装设计，经济实惠，适合家庭长期使用。保持了阿尔卑斯盐的优质品质。', '@/assets/images/02.png', 25.90, 500, '1kg', 6, 1),
(1, 'BBQ调味盐 - 禽肉专用', '专为禽肉调制的BBQ风味', '精选香料与阿尔卑斯盐完美融合，专为鸡肉、鸭肉等禽肉烧烤调制。带来浓郁的BBQ风味。', 'https://via.placeholder.com/300x350?text=BBQ+Poultry', 18.90, 600, '200g', 1, 1),
(1, 'BBQ调味盐 - 牛肉专用', '牛肉烧烤的最佳搭档', '专为牛肉烧烤设计的调味盐，香料配比经过精心调制，能够激发牛肉的鲜美。', 'https://via.placeholder.com/300x350?text=BBQ+Beef', 19.90, 500, '200g', 2, 1),
(1, '地中海香草盐', '地中海风情，香草芬芳', '融合迷迭香、百里香等地中海香草，为菜肴带来独特的地中海风味。', 'https://via.placeholder.com/300x350?text=Mediterranean', 17.90, 700, '150g', 3, 1),
(1, '大蒜香草盐', '蒜香浓郁，提味增香', '精选大蒜与多种香草调配，适合各种肉类和蔬菜的调味，提升菜肴层次。', 'https://via.placeholder.com/300x350?text=Garlic+Herb', 16.90, 800, '150g', 4, 1),
(3, '精制食用盐', '纯净精制，日常必备', '经过精制工艺处理的食用盐，纯度高，颗粒均匀，是厨房必备的基础调味品。', 'https://via.placeholder.com/300x350?text=Table+Salt', 8.90, 3000, '500g', 1, 1),
(3, '加碘食用盐', '补充碘元素，健康之选', '添加适量碘元素的食用盐，有助于预防碘缺乏，适合全家日常使用。', 'https://via.placeholder.com/300x350?text=Iodized', 9.90, 2500, '500g', 2, 1),
(3, '低钠食用盐', '减钠配方，关爱健康', '采用减钠配方，适合需要控制钠摄入的人群，不影响咸味口感。', 'https://via.placeholder.com/300x350?text=Low+Sodium', 14.90, 1000, '500g', 3, 1),
(3, '海盐', '天然海盐，矿物质丰富', '来自纯净海域的天然海盐，富含多种矿物质和微量元素。', 'https://via.placeholder.com/300x350?text=Sea+Salt', 13.90, 1200, '500g', 4, 1),
(4, '喜马拉雅粉盐', '粉红色泽，矿物质丰富', '来自喜马拉雅山脉的天然粉盐，富含84种矿物质和微量元素，呈现独特的粉红色泽。', 'https://via.placeholder.com/300x350?text=Himalayan', 29.90, 400, '500g', 1, 1),
(4, '黑盐', '独特风味，印度特色', '印度传统黑盐，带有独特的硫磺味，常用于印度料理和素食烹饪。', 'https://via.placeholder.com/300x350?text=Black+Salt', 24.90, 300, '250g', 2, 1),
(4, '烟熏盐', '烟熏风味，提升层次', '经过天然木材烟熏处理的盐，带有浓郁的烟熏香气，为菜肴增添独特风味。', 'https://via.placeholder.com/300x350?text=Smoked', 26.90, 350, '200g', 3, 1);

