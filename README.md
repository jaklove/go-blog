# go-blog
1.global全局变量 setting.go中的setting.go中存储对应的ServerSetting、AppSetting、DataBaseSetting、Logger全局变量配置，在

docker run -v /workspace/skeleton:/data/project -p 9501:9501 -it 


CREATE TABLE `blog_auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_key` varchar(100) DEFAULT '' COMMENT 'key',
  `app_secret` varchar(100) DEFAULT '' COMMENT 'secret',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

INSERT into `blog_service` ( `id`, `created_on`, `created_by`, `modified_on`, `modified_by`, `created_on`, `is_del` )
VALUES
	( 1, 0, 'zhourenjie', 0, '', 0, 0 )
	
insert INTO `blog_auth` (`app_key`,`app_secret`) VALUES ('zhourenjie','go-blog')