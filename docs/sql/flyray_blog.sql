/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 8.0.15 : Database - flyray_blog
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`flyray_blog` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `flyray_blog`;

/*Table structure for table `blog_article` */

DROP TABLE IF EXISTS `blog_article`;

CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '博客内容(html)',
  `md_content` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '博客内容(markdown)',
  `create_time` int(11) DEFAULT NULL COMMENT '创建时间',
  `create_by` int(100) DEFAULT NULL COMMENT '创建人',
  `update_time` int(11) DEFAULT NULL COMMENT '修改时间',
  `update_by` int(255) DEFAULT NULL COMMENT '修改人',
  `delete_time` int(11) DEFAULT NULL COMMENT '删除时间',
  `delete_by` int(255) DEFAULT NULL COMMENT '删除人',
  `delete_flag` int(10) unsigned DEFAULT '0' COMMENT '硬删除标志',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  `cover_image_url` varchar(255) DEFAULT NULL COMMENT '封面图片地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COMMENT='文章管理';

/*Data for the table `blog_article` */

insert  into `blog_article`(`id`,`tag_id`,`title`,`desc`,`content`,`md_content`,`create_time`,`create_by`,`update_time`,`update_by`,`delete_time`,`delete_by`,`delete_flag`,`state`,`cover_image_url`) values (1,1,'33','put article','233242','',1565626691,0,1565627361,12,NULL,NULL,1565627558,0,'22'),(13,1,'33','ee','233242','123456',1565626744,0,1565626744,0,NULL,NULL,1565627607,0,'22'),(14,1,'33','ee','233242','123456',1565627289,0,1565627289,0,NULL,NULL,0,0,'22'),(15,1,'33','ee','233242','123456',1565627290,0,1565627290,0,NULL,NULL,0,0,'22'),(16,1,'33','ee','233242','123456',1565627291,0,1565627291,0,NULL,NULL,0,0,'22');

/*Table structure for table `blog_auth` */

DROP TABLE IF EXISTS `blog_auth`;

CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `blog_auth` */

insert  into `blog_auth`(`id`,`username`,`password`) values (1,'test','test123456');

/*Table structure for table `blog_tag` */

DROP TABLE IF EXISTS `blog_tag`;

CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `create_time` int(11) DEFAULT NULL COMMENT '创建时间',
  `create_by` int(100) DEFAULT '0' COMMENT '创建人',
  `update_time` int(11) DEFAULT NULL COMMENT '修改时间',
  `update_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '0' COMMENT '修改人',
  `delete_time` int(11) DEFAULT '0' COMMENT '删除时间',
  `delete_by` int(11) DEFAULT '0' COMMENT '删除人',
  `delete_flag` int(2) unsigned DEFAULT '0' COMMENT '删除标识',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

/*Data for the table `blog_tag` */

insert  into `blog_tag`(`id`,`name`,`create_time`,`create_by`,`update_time`,`update_by`,`delete_time`,`delete_by`,`delete_flag`,`state`) values (1,'put',NULL,0,1565625141,'1',NULL,0,0,0),(9,'hello',1565621574,10,1565621574,'0',0,0,0,1),(10,'world',1565622467,10,1565622467,'0',0,0,0,1),(11,'creat tag',1565625188,112,1565625188,'0',0,0,0,1);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
