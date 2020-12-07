CREATE TABLE `ob_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `device_id` varchar(30) NOT NULL COMMENT '登录设备',
  `password` varchar(30) default '' COMMENT '登录密码',
  `nike_name` varchar(50) NOT NULL COMMENT '昵称',
  `head_img` varchar(150) NOT NULL COMMENT '头像地址',
  `summary` varchar(250) DEFAULT '' COMMENT '个性签名',
  `reg_from` varchar(20) NOT NULL COMMENT '注册来源(wx,qq)',
  `unionid` varchar(20) defalut '' COMMENT 'wx用户统一标识',
  `reg_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '注册时间',
  `login_time` int(11) DEFAULT '0' COMMENT '登录时间'
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8;


CREATE TABLE `ob_user_content` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户uid',
  `content` tinytext CHARACTER SET utf8mb4 NOT NULL COMMENT '发表内容',
  `img_list` tinytext COMMENT '图片列表',
  `video_url` varchar(255) NOT NULL default '' COMMENT '短视频地址',
  `addr` varchar(50) NOT NULL default '' COMMENT '地理位置',
  `pubtime` int(10) unsigned NOT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


CREATE TABLE `ob_pet` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户uid',
  `head_img` varchar(150) NOT NULL COMMENT '头像地址',
  `pet_name` varchar(50) NOT NULL COMMENT '名字',
  `pet_type` varchar(50) NOT NULL COMMENT '品种',
  `pet_age` tinyint(2) unsigned NOT NULL COMMENT '年龄',
  `pet_sex` varchar(50) NOT NULL COMMENT '性别',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;