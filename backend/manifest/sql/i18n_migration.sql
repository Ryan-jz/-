-- ========================================
-- 国际化数据库迁移脚本
-- 支持中文(zh-CN)、英文(en-US)、德文(de-DE)
-- ========================================

-- 产品分类国际化表
CREATE TABLE IF NOT EXISTS `product_category_i18n` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `category_id` int(11) unsigned NOT NULL COMMENT '分类ID',
  `locale` varchar(10) NOT NULL COMMENT '语言代码: zh-CN, en-US, de-DE',
  `name` varchar(100) NOT NULL COMMENT '分类名称',
  `description` text COMMENT '分类描述',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_category_locale` (`category_id`, `locale`),
  KEY `idx_locale` (`locale`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品分类国际化表';

-- 产品国际化表
CREATE TABLE IF NOT EXISTS `product_i18n` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `product_id` int(11) unsigned NOT NULL COMMENT '产品ID',
  `locale` varchar(10) NOT NULL COMMENT '语言代码: zh-CN, en-US, de-DE',
  `name` varchar(200) NOT NULL COMMENT '产品名称',
  `subtitle` varchar(500) DEFAULT NULL COMMENT '副标题',
  `description` text COMMENT '产品描述',
  `ingredients` text COMMENT '成分说明',
  `usage` text COMMENT '使用方法',
  `features` text COMMENT '产品特点(JSON数组)',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_product_locale` (`product_id`, `locale`),
  KEY `idx_locale` (`locale`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品国际化表';

-- 食谱国际化表
CREATE TABLE IF NOT EXISTS `recipe_i18n` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `recipe_id` int(11) unsigned NOT NULL COMMENT '食谱ID',
  `locale` varchar(10) NOT NULL COMMENT '语言代码: zh-CN, en-US, de-DE',
  `name` varchar(200) NOT NULL COMMENT '食谱名称',
  `subtitle` varchar(500) DEFAULT NULL COMMENT '副标题',
  `description` text COMMENT '简介',
  `ingredients` text COMMENT '食材列表(JSON数组)',
  `content` longtext COMMENT '制作步骤(富文本HTML)',
  `tags` varchar(500) DEFAULT NULL COMMENT '标签（逗号分隔）',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_recipe_locale` (`recipe_id`, `locale`),
  KEY `idx_locale` (`locale`),
  KEY `idx_recipe_id` (`recipe_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='食谱国际化表';

-- ========================================
-- 迁移现有数据到国际化表
-- ========================================

-- 迁移产品分类数据（中文）
INSERT INTO `product_category_i18n` (`category_id`, `locale`, `name`, `description`)
SELECT `id`, 'zh-CN', `name`, `description`
FROM `product_category`
WHERE NOT EXISTS (
  SELECT 1 FROM `product_category_i18n` 
  WHERE `category_id` = `product_category`.`id` AND `locale` = 'zh-CN'
);

-- 迁移产品分类数据（英文）
INSERT INTO `product_category_i18n` (`category_id`, `locale`, `name`, `description`)
SELECT `id`, 'en-US', 
  COALESCE(`name_en`, `name`), 
  `description`
FROM `product_category`
WHERE NOT EXISTS (
  SELECT 1 FROM `product_category_i18n` 
  WHERE `category_id` = `product_category`.`id` AND `locale` = 'en-US'
);

-- 迁移产品数据（中文）
INSERT INTO `product_i18n` (`product_id`, `locale`, `name`, `subtitle`, `description`, `ingredients`, `usage`, `features`)
SELECT `id`, 'zh-CN', `name`, `subtitle`, `description`, `ingredients`, `usage`, `features`
FROM `product`
WHERE NOT EXISTS (
  SELECT 1 FROM `product_i18n` 
  WHERE `product_id` = `product`.`id` AND `locale` = 'zh-CN'
);

-- 迁移产品数据（英文）
INSERT INTO `product_i18n` (`product_id`, `locale`, `name`, `subtitle`, `description`, `ingredients`, `usage`, `features`)
SELECT `id`, 'en-US', 
  COALESCE(`name_en`, `name`), 
  `subtitle`, 
  `description`, 
  `ingredients`, 
  `usage`, 
  `features`
FROM `product`
WHERE NOT EXISTS (
  SELECT 1 FROM `product_i18n` 
  WHERE `product_id` = `product`.`id` AND `locale` = 'en-US'
);

-- 迁移食谱数据（中文）
INSERT INTO `recipe_i18n` (`recipe_id`, `locale`, `name`, `subtitle`, `description`, `ingredients`, `content`, `tags`)
SELECT `id`, 'zh-CN', `name`, `subtitle`, `description`, `ingredients`, `content`, `tags`
FROM `recipe`
WHERE NOT EXISTS (
  SELECT 1 FROM `recipe_i18n` 
  WHERE `recipe_id` = `recipe`.`id` AND `locale` = 'zh-CN'
);

-- 迁移食谱数据（英文）
INSERT INTO `recipe_i18n` (`recipe_id`, `locale`, `name`, `subtitle`, `description`, `ingredients`, `content`, `tags`)
SELECT `id`, 'en-US', 
  COALESCE(`name_en`, `name`), 
  `subtitle`, 
  `description`, 
  `ingredients`, 
  `content`, 
  `tags`
FROM `recipe`
WHERE NOT EXISTS (
  SELECT 1 FROM `recipe_i18n` 
  WHERE `recipe_id` = `recipe`.`id` AND `locale` = 'en-US'
);

-- ========================================
-- 插入德语翻译示例数据
-- ========================================

-- 产品分类德语翻译
INSERT INTO `product_category_i18n` (`category_id`, `locale`, `name`, `description`) VALUES
(1, 'de-DE', 'Gewürzsalz-Serie', 'BBQ-Gewürzsalz, Kräutergewürzsalz und andere spezielle Gewürzsalzprodukte'),
(2, 'de-DE', 'Alpensalz', 'Reine Salzprodukte aus den bayerischen Alpen'),
(3, 'de-DE', 'Speisesalz', 'Tägliche Speisesalz-Serie'),
(4, 'de-DE', 'Spezialsalz', 'Himalaya-Rosa-Salz, schwarzes Salz und andere Spezialsalzprodukte');

-- 产品德语翻译（示例）
INSERT INTO `product_i18n` (`product_id`, `locale`, `name`, `subtitle`, `description`) VALUES
(1, 'de-DE', 'Alpines Grobes Salz', 'Klassisches grobes Salz, geeignet zum Kochen und Würzen', 'Bad Reichenhaller Alpines Grobes Salz stammt aus reinem Hochgebirgssalzwasser und bewahrt natürliche Mineralien und Spurenelemente. Die groben Salzkörner eignen sich zum Kochen und können besser in die Zutaten eindringen, um einen reinen salzigen Geschmack zu bringen.'),
(2, 'de-DE', 'Alpines Feines Salz', 'Feine Textur, erste Wahl für die tägliche Würzung', 'Feine Salzkörner sind feiner und leichter löslich, geeignet für das tägliche Kochen und Würzen am Tisch. Hergestellt mit traditionellem Verdampfungsverfahren, behält den reinen Geschmack von Alpensalz.'),
(3, 'de-DE', 'AlpenJodSalz + Jodid', 'Mit Jod angereichert, schützt die Schilddrüsengesundheit', 'Auf der Basis von reinem Alpensalz wird Jod hinzugefügt, um Jodmangel vorzubeugen und die Schilddrüsengesundheit zu schützen. Geeignet für den täglichen Gebrauch der ganzen Familie.');

-- 食谱德语翻译（示例）
INSERT INTO `recipe_i18n` (`recipe_id`, `locale`, `name`, `subtitle`, `description`, `content`, `tags`) VALUES
(1, 'de-DE', 'Himalaya-Salz-Brathähnchen', 'Knusprig außen, zart innen, salziger Duft', 'Mit Himalaya-Rosa-Salz mariniertes Brathähnchen, zartes und saftiges Fleisch mit einzigartigem Mineralgeschmack', 
'<h3>Schritt 1: Zutaten vorbereiten</h3><p>Das ganze Hähnchen waschen und abtropfen lassen.</p><h3>Schritt 2: Marinieren</h3><p>Mit Himalaya-Rosa-Salz und schwarzem Pfeffer gleichmäßig innen und außen einreiben, 2 Stunden marinieren.</p><h3>Schritt 3: Braten</h3><p>Ofen auf 200 Grad vorheizen, 60 Minuten braten, dabei 2-3 Mal mit Öl bestreichen.</p><h3>Schritt 4: Fertig</h3><p>Nach dem Herausnehmen 10 Minuten ruhen lassen, in Stücke schneiden und servieren.</p>',
'Brathähnchen,Hauptgericht,Hausmannskost'),
(2, 'de-DE', 'Meersalz-Zitronen-Grillfisch', 'Frisch und erfrischend, nahrhaft und gesund', 'Mit natürlichem Meersalz und frischer Zitrone gewürzter Grillfisch, bewahrt die Frische des Fisches', 
'<h3>Schritt 1: Fisch vorbereiten</h3><p>Barsch entschuppen und ausnehmen, waschen, auf beiden Seiten einschneiden.</p><h3>Schritt 2: Würzen</h3><p>Mit Meersalz, Zitronensaft und Knoblauch 30 Minuten marinieren.</p><h3>Schritt 3: Grillen</h3><p>Im Ofen bei 180 Grad 25 Minuten grillen, wenden und weitere 10 Minuten grillen.</p><h3>Schritt 4: Garnieren</h3><p>Mit Koriander und Zitronenscheiben garnieren.</p>',
'Grillfisch,Meeresfrüchte,Gesund');
