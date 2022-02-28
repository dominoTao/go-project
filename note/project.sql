CREATE TABLE `user`
(
    `id`              bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `sex`             tinyint(2) NOT NULL DEFAULT '0' COMMENT '性别;0:保密,1:男,2:女',
    `birthday`        int(11) NOT NULL DEFAULT '0' COMMENT '生日',
    `last_login_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
    `create_time`     int(11) NOT NULL DEFAULT '0' COMMENT '注册时间',
    `user_status`     tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '用户状态;0:禁用,1:正常',
    `user_login`      varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
    `user_pass`       varchar(64)                                                  NOT NULL COMMENT '登录密码',
    `user_nickname`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户昵称',
    `user_email`      varchar(100)                                                 NOT NULL COMMENT '用户登录邮箱',
    `user_url`        varchar(100)                                                 NOT NULL COMMENT '用户个人网址',
    `avatar`          varchar(255)                                                 NOT NULL COMMENT '用户头像',
    `signature`       varchar(255)                                                 NOT NULL COMMENT '个性签名',
    `last_login_ip`   varchar(15)                                                  NOT NULL COMMENT '最后登录ip',
    `mobile`          varchar(20)                                                  NOT NULL COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
    `more`            text                                                         NOT NULL COMMENT '扩展属性',
    PRIMARY KEY (`id`),
    KEY               `user_login` (`user_login`),
    KEY               `user_nickname` (`user_nickname`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';


CREATE TABLE `role`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `status`      tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态;0:禁用;1:正常',
    `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
    `list_order`  float                                                        NOT NULL DEFAULT '0' COMMENT '排序',
    `name`        varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
    `remark`      varchar(255)                                                 NOT NULL DEFAULT '' COMMENT '备注',
    PRIMARY KEY (`id`),
    KEY           `parent_id` (`parent_id`),
    KEY           `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';


CREATE TABLE `role_user`
(
    `id`      bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '角色 id',
    `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    PRIMARY KEY (`id`),
    KEY       `role_id` (`role_id`),
    KEY       `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户角色对应表';


CREATE TABLE `admin_menu`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单id',
    `parent_id`  int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父菜单id',
    `type`       tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '菜单类型;1:有界面可访问菜单,2:无界面可访问菜单,0:只作为菜单',
    `status`     tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态;1:显示,0:不显示',
    `list_order` float                                                         NOT NULL DEFAULT '10000' COMMENT '排序',
    `app`        varchar(40) CHARACTER SET utf8                                NOT NULL DEFAULT '' COMMENT '应用名',
    `controller` varchar(30) CHARACTER SET utf8                                NOT NULL DEFAULT '' COMMENT '路由',
    `name`       varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '菜单名称',
    `icon`       varchar(20) CHARACTER SET utf8                                NOT NULL DEFAULT '' COMMENT '菜单图标',
    `remark`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
    PRIMARY KEY (`id`),
    KEY          `status` (`status`),
    KEY          `parent_id` (`parent_id`),
    KEY          `controller` (`controller`)
) ENGINE=InnoDB AUTO_INCREMENT=168 DEFAULT CHARSET=utf8mb4 COMMENT='后台菜单表';