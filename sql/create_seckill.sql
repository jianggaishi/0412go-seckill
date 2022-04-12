/*
 Navicat Premium Data Transfer

 Source Server         : 阿里云mysql
 Source Server Type    : MySQL
 Source Server Version : 50651
 Source Host           : 39.99.214.230:3306
 Source Schema         : seckill

 Target Server Type    : MySQL
 Target Server Version : 50651
 File Encoding         : 65001

 Date: 20/02/2022 00:07:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE `seckill` CHARACTER SET 'utf8';

use seckill;
-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity` (
  `activity_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '活动Id',
  `activity_name` varchar(50) NOT NULL DEFAULT '' COMMENT '活动名称',
  `product_id` int(11) unsigned NOT NULL COMMENT '商品Id',
  `start_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '活动开始时间',
  `end_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '活动结束时间',
  `total` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '活动状态',
  `sec_speed` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '每秒限制多少个商品售出',
  `buy_limit` int(5) unsigned NOT NULL COMMENT '购买限制',
  `buy_rate` decimal(2,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '购买限制',
  PRIMARY KEY (`activity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='@活动数据表';

-- ----------------------------
-- Records of activity
-- ----------------------------
-- BEGIN;
-- INSERT INTO `activity` (`activity_id`, `activity_name`, `product_id`, `start_time`, `end_time`, `total`, `status`, `sec_speed`, `buy_limit`, `buy_rate`) VALUES (1, '香蕉大甩卖', 1, 530871061, 530872061, 20, 0, 1, 1, 0.20);
-- INSERT INTO `activity` (`activity_id`, `activity_name`, `product_id`, `start_time`, `end_time`, `total`, `status`, `sec_speed`, `buy_limit`, `buy_rate`) VALUES (2, '苹果大甩卖', 2, 530871061, 530872061, 20, 0, 1, 1, 0.20);
-- INSERT INTO `activity` (`activity_id`, `activity_name`, `product_id`, `start_time`, `end_time`, `total`, `status`, `sec_speed`, `buy_limit`, `buy_rate`) VALUES (3, '桃子大甩卖', 3, 1530928052, 1530989052, 20, 0, 1, 1, 0.20);
-- INSERT INTO `activity` (`activity_id`, `activity_name`, `product_id`, `start_time`, `end_time`, `total`, `status`, `sec_speed`, `buy_limit`, `buy_rate`) VALUES (4, '梨子大甩卖', 4, 1530928052, 1530989052, 20, 0, 1, 1, 0.20);
-- COMMIT;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `product_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品Id',
  `product_name` varchar(50) NOT NULL DEFAULT '' COMMENT '商品名称',
  `total` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '商品状态',
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='@商品数据表';

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` (`product_id`, `product_name`, `total`, `status`) VALUES (1, '香蕉', 100, 1);
INSERT INTO `product` (`product_id`, `product_name`, `total`, `status`) VALUES (2, '苹果', 100, 1);
INSERT INTO `product` (`product_id`, `product_name`, `total`, `status`) VALUES (3, '桃子', 100, 1);
INSERT INTO `product` (`product_id`, `product_name`, `total`, `status`) VALUES (4, '梨子', 100, 1);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` int(10) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `age` int(10) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`user_id`, `user_name`, `password`, `age`) VALUES (1, 'zp', 'zhangpeng', 25);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
