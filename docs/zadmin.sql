/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : zadmin

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 06/11/2020 16:26:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for zs_files
-- ----------------------------
DROP TABLE IF EXISTS `zs_files`;
CREATE TABLE `zs_files`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片名称',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片地址',
  `size` double(20, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '图片大小',
  `type` tinyint(2) UNSIGNED NOT NULL DEFAULT 0 COMMENT '使用状态：0未使用，1使用中',
  `ctime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文件管理器' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for zs_role
-- ----------------------------
DROP TABLE IF EXISTS `zs_role`;
CREATE TABLE `zs_role`  (
  `id` smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `role_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `role_ctime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色创建时间',
  `role_utime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色修改时间',
  `role_desc` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `role_name`(`role_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for zs_user
-- ----------------------------
DROP TABLE IF EXISTS `zs_user`;
CREATE TABLE `zs_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'zadmin' COMMENT '混淆值',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `rname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `ctime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `utime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `ltime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录时间',
  `is_del` tinyint(2) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0未删除，1已删除',
  `dtime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  `openid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信openid',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `state` tinyint(2) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态：0关闭，1启用',
  `desc` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  UNIQUE INDEX `email`(`email`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for zs_user_role
-- ----------------------------
DROP TABLE IF EXISTS `zs_user_role`;
CREATE TABLE `zs_user_role`  (
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `role_id` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色id',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户-角色表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
