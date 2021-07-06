/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : 127.0.0.1:3306
 Source Schema         : hauth

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 06/07/2021 15:44:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user_info
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_info`;
CREATE TABLE `sys_user_info` (
  `user_id` varchar(30) NOT NULL,
  `user_name` varchar(300) DEFAULT NULL,
  `user_create_date` datetime DEFAULT NULL,
  `user_owner` varchar(30) DEFAULT NULL,
  `user_email` varchar(30) DEFAULT NULL,
  `user_phone` decimal(15,0) DEFAULT NULL,
  `org_unit_id` varchar(66) DEFAULT NULL,
  `user_maintance_date` datetime DEFAULT NULL,
  `user_maintance_user` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `fk_sys_user_org_idx` (`org_unit_id`),
  CONSTRAINT `fk_sys_user_org` FOREIGN KEY (`org_unit_id`) REFERENCES `sys_org_info` (`org_unit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_user_info
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_info` VALUES ('431243', '2143214', '2017-06-28 10:51:09', 'admin', 'hzwy23@163.com', 18986110550, 'vertex_root_join_vertex_root', '2017-06-28 10:51:09', 'admin');
INSERT INTO `sys_user_info` VALUES ('admin', '超级管理员', '2016-01-01 00:00:00', 'sys', 'hzwy23@163.com', 18986110550, 'vertex_root_join_vertex_root', '2017-06-28 10:54:54', 'admin');
INSERT INTO `sys_user_info` VALUES ('caadmin', 'CA业务管理员', '2017-03-18 14:32:22', 'admin', 'hzwy23@163.com', 18986110550, 'mas_join_34124', '2017-03-18 14:32:22', 'admin');
INSERT INTO `sys_user_info` VALUES ('demo', '演示账号', '2017-03-01 21:21:38', 'admin', 'hzwy23@sina.com', 18986110551, 'mas_join_4542346', '2017-04-24 20:53:42', 'admin');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
