CREATE TABLE `ob_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(30) NOT NULL COMMENT '登录账户',
  `password` varchar(30) NOT NULL COMMENT '登录密码',
  `nike_name` varchar(50) NOT NULL COMMENT '昵称',
  `reg_from` varchar(20) NOT NULL COMMENT '注册来源',
  `reg_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '注册时间',
  `login_time` int(11) DEFAULT '0' COMMENT '登录时间',
  `head_img` varchar(150) NOT NULL COMMENT '头像地址',
  `summary` varchar(250) DEFAULT NULL COMMENT '个性签名',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=100002 DEFAULT CHARSET=utf8;


CREATE TABLE `ob_user_content` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户uid',
  `content` tinytext CHARACTER SET utf8mb4 NOT NULL COMMENT '发表内容',
  `imglist` tinytext COMMENT '图片列表',
  `pubtime` int(10) unsigned NOT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;