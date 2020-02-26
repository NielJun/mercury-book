/*
 Navicat MySQL Data Transfer

 Source Server         : golang
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : mercury

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 26/02/2020 20:20:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for answer
-- ----------------------------
DROP TABLE IF EXISTS `answer`;
CREATE TABLE `answer` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `answer_id` bigint unsigned NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `comment_count` int unsigned NOT NULL,
  `voteup_count` int NOT NULL,
  `author_id` bigint NOT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '1',
  `can_comment` tinyint unsigned NOT NULL DEFAULT '1',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_answer_id` (`answer_id`),
  KEY `idx_author_Id` (`author_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of answer
-- ----------------------------
BEGIN;
INSERT INTO `answer` VALUES (1, 12333, '你睡觉打卡', 1, 1, 290520267376033793, 1, 1, '2020-02-26 20:12:06', '2020-02-26 20:12:06');
INSERT INTO `answer` VALUES (2, 123332, '阿巴斯对不对', 1, 1, 290520267376033793, 1, 1, '2020-02-26 20:12:26', '2020-02-26 20:12:26');
INSERT INTO `answer` VALUES (3, 22, '阿时间快点快点吧', 1, 1, 290520267376033793, 1, 1, '2020-02-26 20:12:41', '2020-02-26 20:12:41');
COMMIT;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `category_id` int unsigned NOT NULL,
  `category_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_category_id` (`category_id`),
  UNIQUE KEY `idx_category_name` (`category_name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1, 1, '技术', '2019-01-01 08:30:40', '2019-01-01 08:30:40');
INSERT INTO `category` VALUES (2, 2, '情感', '2019-01-01 08:31:07', '2019-01-01 08:31:07');
INSERT INTO `category` VALUES (3, 3, '王者荣耀', '2019-01-01 08:31:25', '2019-01-01 08:31:25');
INSERT INTO `category` VALUES (4, 4, '吃鸡', '2019-01-01 15:45:13', '2019-01-01 15:45:13');
INSERT INTO `category` VALUES (5, 5, '科幻', '2019-01-05 23:02:43', '2019-01-05 23:02:43');
COMMIT;

-- ----------------------------
-- Table structure for question
-- ----------------------------
DROP TABLE IF EXISTS `question`;
CREATE TABLE `question` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `question_id` bigint NOT NULL COMMENT '问题id',
  `caption` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '问题标题',
  `content` varchar(8192) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '问题内容',
  `author_id` bigint NOT NULL COMMENT '作者的用户id',
  `category_id` bigint NOT NULL COMMENT '所属栏目',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '问题状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_question_id` (`question_id`),
  KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of question
-- ----------------------------
BEGIN;
INSERT INTO `question` VALUES (5, 290520334266793985, 'Mac版小黑屋提示无法确认开发者身份的解决办法', '安装Mac版小黑屋之后，竟然提示\n\n无法打开“LonelyWriter”，因为无法确认开发者的身份。\n\n您的安全性偏好设置仅允许安装来自 Mac App Store 和被认可的开发者的应用程序。\n\n“Safari”于昨天 下午7:13 从“www.mochiwang.com”下载了此文件。', 290520267376033793, 2, 1, '2020-02-26 13:00:13', '2020-02-26 13:00:13');
INSERT INTO `question` VALUES (7, 290563705467305985, '阿斯顿毕卡索的', '离马路上都可能没那么傻的难', 290520267376033793, 2, 1, '2020-02-26 20:11:04', '2020-02-26 20:11:04');
COMMIT;

-- ----------------------------
-- Table structure for question_answer_rel
-- ----------------------------
DROP TABLE IF EXISTS `question_answer_rel`;
CREATE TABLE `question_answer_rel` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `question_id` bigint NOT NULL,
  `answer_id` bigint NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_question_answer` (`question_id`,`answer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of question_answer_rel
-- ----------------------------
BEGIN;
INSERT INTO `question_answer_rel` VALUES (1, 290520334266793985, 123332, '2020-02-26 18:41:19');
INSERT INTO `question_answer_rel` VALUES (2, 290520334266793985, 22, '2020-02-26 18:41:29');
INSERT INTO `question_answer_rel` VALUES (3, 290520334266793985, 12333, '2020-02-26 20:13:44');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sex` tinyint NOT NULL DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (29, 290520267376033793, 'daniel', '烈火讽刺', '81ff5158ef0538a6b725c456e4d6d85f', '769288695@qq.com', 1, '2020-02-26 12:59:33', '2020-02-26 12:59:33');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
