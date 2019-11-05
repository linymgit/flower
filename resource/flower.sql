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
  `avatar_url` varchar(255) NOT NULL DEFAULT '0' COMMENT '头像',
  `name` varchar(50) NOT NULL DEFAULT '0',
  `password` varchar(50) NOT NULL DEFAULT '0',
  `save_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='账户';

-- 数据导出被取消选择。


-- 导出  表 flower.ad 结构
CREATE TABLE IF NOT EXISTS `ad` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `slogan` varchar(50) NOT NULL DEFAULT '0' COMMENT '广告语',
  `pic_url` varchar(50) NOT NULL DEFAULT '0',
  `postion_id` int(11) NOT NULL DEFAULT '0' COMMENT '广告位置',
  `goto_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '跳转类型 0:url 1:to product',
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
  `expression` varchar(50) NOT NULL DEFAULT '0' COMMENT '位置表达式',
  `name` varchar(50) NOT NULL DEFAULT '0' COMMENT '广告位',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 数据导出被取消选择。


-- 导出  表 flower.article 结构
CREATE TABLE IF NOT EXISTS `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type_id` int(11) NOT NULL DEFAULT '0',
  `title` varchar(50) NOT NULL DEFAULT '0',
  `author` varchar(50) NOT NULL DEFAULT '0',
  `source` varchar(50) NOT NULL DEFAULT '0',
  `source_url` varchar(50) NOT NULL DEFAULT '0',
  `preview` varchar(50) NOT NULL DEFAULT '0',
  `key_word` varchar(50) NOT NULL DEFAULT '0',
  `summary` varchar(50) NOT NULL DEFAULT '0',
  `content` varchar(50) NOT NULL DEFAULT '0',
  `clicks` int(11) NOT NULL DEFAULT '0',
  `states` tinyint(4) NOT NULL DEFAULT '0',
  `weight` int(11) NOT NULL DEFAULT '0',
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章';

-- 数据导出被取消选择。


-- 导出  表 flower.article_type 结构
CREATE TABLE IF NOT EXISTS `article_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type_name` varchar(50) DEFAULT NULL,
  `weight` int(11) DEFAULT NULL,
  `parent_id` bigint(20) DEFAULT NULL,
  `save_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章类别';

-- 数据导出被取消选择。


-- 导出  表 flower.crm 结构
CREATE TABLE IF NOT EXISTS `crm` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `phone` varchar(20) NOT NULL DEFAULT '0',
  `official_web` varchar(50) NOT NULL DEFAULT '0' COMMENT '官网',
  `company` varchar(50) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL COMMENT '删除状态 1删除 0未删除',
  `message` varchar(50) NOT NULL DEFAULT '0' COMMENT '留言',
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='客户管理';

-- 数据导出被取消选择。


-- 导出  表 flower.partner 结构
CREATE TABLE IF NOT EXISTS `partner` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `logo` varchar(50) NOT NULL DEFAULT '0',
  `enterprise_name` varchar(50) NOT NULL DEFAULT '0' COMMENT '企业名称',
  `intro` varchar(50) NOT NULL DEFAULT '0',
  `weight` int(11) NOT NULL DEFAULT '0',
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='合作商';

-- 数据导出被取消选择。


-- 导出  表 flower.product 结构
CREATE TABLE IF NOT EXISTS `product` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '0' COMMENT '商品名称',
  `intro` varchar(50) NOT NULL DEFAULT '0' COMMENT '简介',
  `summary` varchar(50) NOT NULL DEFAULT '0',
  `states` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0在线 1下线',
  `index_show` tinyint(4) NOT NULL DEFAULT '0' COMMENT '首页推荐 1是 0否',
  `details_pic_url` varchar(50) NOT NULL DEFAULT '0',
  `cover_url` varchar(50) NOT NULL DEFAULT '0',
  `price` decimal(10,0) NOT NULL DEFAULT '0',
  `heat` int(11) NOT NULL DEFAULT '0' COMMENT '热度',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '类目id',
  `author_id` int(11) NOT NULL DEFAULT '0' COMMENT '作者id',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='产品';

-- 数据导出被取消选择。


-- 导出  表 flower.product_category 结构
CREATE TABLE IF NOT EXISTS `product_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '0',
  `desc` varchar(50) NOT NULL DEFAULT '0',
  `states` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0显示 1不显示',
  `level` int(11) NOT NULL DEFAULT '0' COMMENT '类目级别',
  `weight` int(11) NOT NULL DEFAULT '0',
  `save_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='产品类目';

-- 数据导出被取消选择。


-- 导出  表 flower.web_setting 结构
CREATE TABLE IF NOT EXISTS `web_setting` (
  `id` int(11) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL COMMENT '公司名称',
  `url` varchar(50) DEFAULT NULL COMMENT '官网地址',
  `rectangle_logo` varchar(50) DEFAULT NULL COMMENT '长方形的logo',
  `square_logo` varchar(50) DEFAULT NULL COMMENT '正方形的logo',
  `address` varchar(50) DEFAULT NULL COMMENT '公司地址',
  `enterprise_email` varchar(50) DEFAULT NULL COMMENT '企业邮箱',
  `hotline` varchar(50) DEFAULT NULL COMMENT '服务热线',
  `icp` varchar(50) DEFAULT NULL COMMENT '网站ICP备案号',
  `save_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='网站设置';

-- 数据导出被取消选择。
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
