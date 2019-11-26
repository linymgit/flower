-- --------------------------------------------------------
-- 主机:                           123.207.1.119
-- 服务器版本:                        8.0.13 - MySQL Community Server - GPL
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  9.3.0.4984
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 flower 的数据库结构
CREATE DATABASE IF NOT EXISTS `flower` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */;
USE `flower`;


-- 导出  表 flower.account 结构
CREATE TABLE IF NOT EXISTS `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `avatar_url` varchar(255) DEFAULT '0' COMMENT '头像',
  `name` varchar(50) NOT NULL DEFAULT '0',
  `password` varchar(50) NOT NULL DEFAULT '0',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='账户';

-- 数据导出被取消选择。


-- 导出  表 flower.ad 结构
CREATE TABLE IF NOT EXISTS `ad` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `slogan` varchar(255) NOT NULL DEFAULT '0' COMMENT '广告语',
  `pic_url` varchar(255) NOT NULL DEFAULT '0',
  `postion_id` int(11) NOT NULL DEFAULT '0' COMMENT '广告位置',
  `goto_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '跳转类型 0:url 1:to product',
  `ad_link` varchar(255) NOT NULL DEFAULT '0' COMMENT '广告链接',
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0在线 1下线',
  `clicks` int(11) NOT NULL DEFAULT '0' COMMENT '点击数',
  `weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重用于排序，数值越大权重越大',
  `start_time` timestamp NOT NULL,
  `end_time` timestamp NOT NULL,
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='广告';

-- 数据导出被取消选择。


-- 导出  表 flower.ad_position 结构
CREATE TABLE IF NOT EXISTS `ad_position` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `expression` varchar(255) NOT NULL DEFAULT '0' COMMENT '位置表达式',
  `name` varchar(50) NOT NULL DEFAULT '0' COMMENT '广告位',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 数据导出被取消选择。


-- 导出  表 flower.article 结构
CREATE TABLE IF NOT EXISTS `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type_id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `author` varchar(50) NOT NULL,
  `source` varchar(255) NOT NULL,
  `source_url` varchar(255) NOT NULL,
  `preview` varchar(50) NOT NULL,
  `key_word` varchar(50) NOT NULL,
  `summary` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `clicks` int(11) NOT NULL,
  `states` tinyint(4) NOT NULL,
  `sort` int(11) NOT NULL,
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ids_clicks` (`clicks`),
  KEY `ids_type_id` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章';

-- 数据导出被取消选择。


-- 导出  表 flower.article_type 结构
CREATE TABLE IF NOT EXISTS `article_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type_name` varchar(50) DEFAULT NULL,
  `sort` int(11) DEFAULT NULL,
  `level` int(11) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `save_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章类别';

-- 数据导出被取消选择。


-- 导出  表 flower.business_partners 结构
CREATE TABLE IF NOT EXISTS `business_partners` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `logo` varchar(256) NOT NULL DEFAULT '0' COMMENT 'logo url',
  `business_name` varchar(64) NOT NULL DEFAULT '0' COMMENT '企业名称',
  `intro` varchar(256) NOT NULL DEFAULT '0' COMMENT '企业介绍',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `save_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商业合作伙伴';

-- 数据导出被取消选择。


-- 导出  表 flower.crm 结构
CREATE TABLE IF NOT EXISTS `crm` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(32) NOT NULL,
  `official_web` varchar(255) NOT NULL COMMENT '官网',
  `company` varchar(255) NOT NULL,
  `deleted` tinyint(4) NOT NULL COMMENT '删除状态 1删除 0未删除',
  `message` varchar(1024) NOT NULL COMMENT '留言',
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='客户管理';

-- 数据导出被取消选择。


-- 导出  表 flower.news_time_nav 结构
CREATE TABLE IF NOT EXISTS `news_time_nav` (
  `year` smallint(5) unsigned NOT NULL,
  `moth` tinyint(2) unsigned NOT NULL,
  `count` int(11) NOT NULL,
  `article_titles` text NOT NULL COMMENT '文章标题列表',
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`year`,`moth`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='新闻发布时间导航';

-- 数据导出被取消选择。


-- 导出  表 flower.partner 结构
CREATE TABLE IF NOT EXISTS `partner` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `logo` varchar(255) NOT NULL,
  `enterprise_name` varchar(255) NOT NULL COMMENT '企业名称',
  `intro` varchar(255) NOT NULL,
  `weight` int(11) NOT NULL,
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='合作商';

-- 数据导出被取消选择。


-- 导出  表 flower.product 结构
CREATE TABLE IF NOT EXISTS `product` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '商品名称',
  `intro` varchar(1024) NOT NULL COMMENT '简介',
  `summary` varchar(1024) NOT NULL,
  `states` tinyint(4) NOT NULL COMMENT '0在线 1下线',
  `index_show` tinyint(4) NOT NULL COMMENT '首页推荐 1是 0否',
  `details_pic_url` varchar(1024) DEFAULT NULL,
  `cover_url` varchar(255) NOT NULL,
  `price` decimal(12,2) NOT NULL,
  `heat` int(11) NOT NULL COMMENT '热度',
  `category_id` int(11) NOT NULL COMMENT '类目id',
  `author_id` bigint(20) NOT NULL COMMENT '作者id',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='产品';

-- 数据导出被取消选择。


-- 导出  表 flower.product_category 结构
CREATE TABLE IF NOT EXISTS `product_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '上级类目id',
  `name` varchar(50) NOT NULL,
  `desc` varchar(255) NOT NULL,
  `states` tinyint(4) NOT NULL COMMENT '0显示 1不显示',
  `level` int(11) NOT NULL COMMENT '类目级别',
  `sort` int(11) NOT NULL,
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='产品类目';

-- 数据导出被取消选择。


-- 导出  表 flower.product_uv 结构
CREATE TABLE IF NOT EXISTS `product_uv` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `p_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '产品id',
  `ip` varchar(32) NOT NULL DEFAULT '0' COMMENT '客户端ip',
  `access_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '访问时间',
  PRIMARY KEY (`id`),
  KEY `p_id` (`p_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='统计访问网站的ip';

-- 数据导出被取消选择。


-- 导出  表 flower.web_setting 结构
CREATE TABLE IF NOT EXISTS `web_setting` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '公司名称',
  `url` varchar(255) NOT NULL COMMENT '官网地址',
  `rectangle_logo` varchar(255) NOT NULL COMMENT '长方形的logo',
  `square_logo` varchar(255) NOT NULL COMMENT '正方形的logo',
  `address` varchar(255) NOT NULL COMMENT '公司地址',
  `enterprise_email` varchar(255) NOT NULL COMMENT '企业邮箱',
  `hotline` varchar(32) NOT NULL COMMENT '服务热线',
  `icp` varchar(32) NOT NULL COMMENT '网站ICP备案号',
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='网站设置';

-- 数据导出被取消选择。
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
