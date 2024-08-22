/*
 Navicat Premium Data Transfer

 Source Server         : rds_admin
 Source Server Type    : MySQL
 Source Server Version : 50743 (5.7.43-log)
 Source Host           : rm-bp19sh9ez6j720szy6o.mysql.rds.aliyuncs.com:3306
 Source Schema         : qm_admin

 Target Server Type    : MySQL
 Target Server Version : 50743 (5.7.43-log)
 File Encoding         : 65001

 Date: 22/08/2024 16:52:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for qm_short_urls
-- ----------------------------
DROP TABLE IF EXISTS `qm_short_urls`;
CREATE TABLE `qm_short_urls`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sha1` char(40) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `create_at` int(11) NOT NULL,
  `creator` int(11) NOT NULL DEFAULT 0,
  `count` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `sha1`(`sha1`) USING BTREE,
  INDEX `create_at`(`create_at`) USING BTREE,
  INDEX `creator`(`creator`) USING BTREE,
  INDEX `count`(`count`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 163936 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci COMMENT = '短邮部短链接表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
