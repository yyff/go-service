create database if not exists go_service DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

use go_service;

create table if not exists post (
   id BIGINT unsigned NOT NULL COMMENT 'ID' AUTO_INCREMENT,
   userid BIGINT unsigned NOT NULL COMMENT 'user ID',
   title VARCHAR(64) NOT NULL COMMENT 'title',
   content VARCHAR(64) NOT NULL COMMENT 'content',
   PRIMARY KEY (id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='帖子表';