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

 Date: 12/11/2020 11:31:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for zs_article
-- ----------------------------
DROP TABLE IF EXISTS `zs_article`;
CREATE TABLE `zs_article`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `title` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `short_title` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '简短副标题',
  `category_id` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类id',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '内容',
  `remark` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `ctime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `utime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '编辑时间',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '作者id',
  `click` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击量',
  `like` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '喜欢量',
  `type` tinyint(2) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0普通文章，1单页',
  `state` tinyint(2) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态：0隐藏，1显示',
  `seo_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'seo标题',
  `seo_kwds` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'seo关键字',
  `seo_desc` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'seo描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `title`(`title`) USING BTREE,
  INDEX `short_title`(`short_title`) USING BTREE,
  INDEX `category_id`(`category_id`) USING BTREE,
  CONSTRAINT `category_id` FOREIGN KEY (`category_id`) REFERENCES `zs_category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for zs_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `zs_article_tag`;
CREATE TABLE `zs_article_tag`  (
  `article_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章id',
  `tag_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '标签id',
  PRIMARY KEY (`article_id`, `tag_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章-标签' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for zs_category
-- ----------------------------
DROP TABLE IF EXISTS `zs_category`;
CREATE TABLE `zs_category`  (
  `id` smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `pid` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级分类',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `ename` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类英文',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类图标',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类描述',
  `ctime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `utime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '编辑时间',
  `seo_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'seo标题',
  `seo_kwds` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'seo关键字',
  `seo_desc` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'seo描述',
  `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类图片',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `name`(`name`) USING BTREE,
  INDEX `ename`(`ename`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '栏目分类' ROW_FORMAT = Dynamic;

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
-- Table structure for zs_menu
-- ----------------------------
DROP TABLE IF EXISTS `zs_menu`;
CREATE TABLE `zs_menu`  (
  `id` smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单地址',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '附加参数',
  `pid` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级菜单id',
  `ordernum` int(10) NOT NULL DEFAULT 1000 COMMENT '排序',
  `ctime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `utime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `mname` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `type` tinyint(2) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0菜单，1地址',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `url`(`url`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `mname`(`mname`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单表' ROW_FORMAT = Dynamic;

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
-- Table structure for zs_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `zs_role_menu`;
CREATE TABLE `zs_role_menu`  (
  `role_id` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色id',
  `menu_id` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单id',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色-菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for zs_tag
-- ----------------------------
DROP TABLE IF EXISTS `zs_tag`;
CREATE TABLE `zs_tag`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '标签id',
  `tag_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签名称',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `tag_name`(`tag_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签' ROW_FORMAT = Dynamic;

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
  `role_id` smallint(6) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  UNIQUE INDEX `email`(`email`) USING BTREE,
  INDEX `role_id`(`role_id`) USING BTREE,
  CONSTRAINT `role_id` FOREIGN KEY (`role_id`) REFERENCES `zs_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
