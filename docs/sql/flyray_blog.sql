/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.7.27-log : Database - flyray_blog
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`flyray_blog` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `flyray_blog`;

/*Table structure for table `blog_article` */

DROP TABLE IF EXISTS `blog_article`;

CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章主键id',
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '文章标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text COMMENT '博客内容(html)',
  `md_content` text COMMENT '博客内容(markdown)',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `create_by` int(100) DEFAULT NULL COMMENT '创建人',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '修改时间',
  `update_by` int(255) DEFAULT NULL COMMENT '修改人',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `delete_by` int(255) DEFAULT NULL COMMENT '删除人',
  `delete_flag` int(10) DEFAULT '0' COMMENT '硬删除标志',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  `cover_image_url` varchar(255) DEFAULT NULL COMMENT '封面图片地址',
  `enable_comment` tinyint(4) DEFAULT '0' COMMENT '0-允许评论 1-不允许评论',
  `views` bigint(20) DEFAULT '0' COMMENT '阅读量',
  `category_name` varchar(200) DEFAULT NULL COMMENT '博客分类(冗余字段)',
  `category_id` int(11) DEFAULT NULL COMMENT '博客分类ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8 COMMENT='文章管理';

/*Data for the table `blog_article` */

insert  into `blog_article`(`id`,`tag_id`,`title`,`desc`,`content`,`md_content`,`create_time`,`create_by`,`update_time`,`update_by`,`delete_time`,`delete_by`,`delete_flag`,`state`,`cover_image_url`,`enable_comment`,`views`,`category_name`,`category_id`) values (35,1,'33','ee','<h3 id=\"h3--editor-md\"><a name=\"关于 Editor.md\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>关于 Editor.md</h3><pre><code>                    **Editor.md** 是一款开源的、可嵌入的 Markdown 在线编辑器（组件），基于 CodeMirror、jQuery 和 Marked 构建。\n</code></pre>','### 关于 Editor.md\n						**Editor.md** 是一款开源的、可嵌入的 Markdown 在线编辑器（组件），基于 CodeMirror、jQuery 和 Marked 构建。\n				','2019-08-25 20:57:04',0,'2019-08-25 20:57:04',0,NULL,NULL,0,0,'22',0,0,NULL,NULL);

/*Table structure for table `blog_auth` */

DROP TABLE IF EXISTS `blog_auth`;

CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `nickname` varchar(50) DEFAULT '' COMMENT '昵称',
  `email` varchar(50) DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(20) DEFAULT '' COMMENT '手机号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `blog_auth` */

insert  into `blog_auth`(`id`,`username`,`password`,`nickname`,`email`,`mobile`) values (1,'test','test123456','','','');

/*Table structure for table `blog_category` */

DROP TABLE IF EXISTS `blog_category`;

CREATE TABLE `blog_category` (
  `id` bigint(20) NOT NULL DEFAULT '0' COMMENT '文章分类ID',
  `name` varchar(100) DEFAULT '' COMMENT '分类名称',
  `create_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者',
  `delete_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 0-未删除 1-已删除',
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否审核通过 0-未审核 1-审核通过',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `blog_category` */

/*Table structure for table `blog_comment` */

DROP TABLE IF EXISTS `blog_comment`;

CREATE TABLE `blog_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `article_id` bigint(20) DEFAULT NULL COMMENT '关联的文章ID',
  `create_by` bigint(20) DEFAULT NULL COMMENT '评论者',
  `email` varchar(100) DEFAULT NULL COMMENT '评论人的邮箱',
  `content` text COMMENT '评论内容',
  `reply_content` text COMMENT '回复内容',
  `delete_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 0-未删除 1-已删除',
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否审核通过 0-未审核 1-审核通过',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
  `reply_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '回复时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `blog_comment` */

insert  into `blog_comment`(`id`,`article_id`,`create_by`,`email`,`content`,`reply_content`,`delete_flag`,`state`,`create_time`,`reply_time`,`update_time`) values (1,1,1,NULL,'hello comment',NULL,0,1,'2019-08-25 23:21:21','2019-08-25 23:11:18','2019-08-25 23:11:19'),(2,2,2,NULL,'hello comment',NULL,0,1,'2019-08-25 23:21:22','2019-08-25 23:14:04','2019-08-25 23:14:05'),(3,3,3,NULL,'hello comment',NULL,0,1,'2019-08-25 23:21:23','2019-08-25 23:19:56','2019-08-25 23:19:56');

/*Table structure for table `blog_tag` */

DROP TABLE IF EXISTS `blog_tag`;

CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '标签表主键id',
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `create_by` int(100) DEFAULT '0' COMMENT '创建人',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '修改时间',
  `update_by` varchar(100) DEFAULT '0' COMMENT '修改人',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `delete_by` int(10) DEFAULT '0' COMMENT '删除人',
  `delete_flag` int(10) DEFAULT '0' COMMENT '是否删除 0=否 1=是',
  `state` tinyint(10) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

/*Data for the table `blog_tag` */

insert  into `blog_tag`(`id`,`name`,`create_time`,`create_by`,`update_time`,`update_by`,`delete_time`,`delete_by`,`delete_flag`,`state`) values (1,'puttt',NULL,112,'2019-08-18 08:17:41','112',NULL,0,0,0);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
