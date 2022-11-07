# 用户表
CREATE TABLE IF NOT EXISTS users
(
    `id`         bigint unsigned not null AUTO_INCREMENT,
    `nickname`   varchar(50)     not null comment '昵称',
    `username`   varchar(50)     not null comment '账号',
    `password`   varchar(255)    not null comment '密码',
    `avatar`     varchar(255)    NULL     DEFAULT NULL comment '头像',
    `status`     tinyint         not null default 1 COMMENT '1 启用 0禁用',
    `login_ip`   varchar(100)    NULL COMMENT '登录IP',
    `login_at`   timestamp       NULL COMMENT '登录时间',
    `created_at` timestamp       NULL     default NULL COMMENT '创建时间',
    `updated_at` timestamp       NULL     default NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `users_username_unique` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

# 角色表
CREATE TABLE IF NOT EXISTS roles
(
    `id`         bigint unsigned not null auto_increment,
    `name`       varchar(50)     not null comment '角色名',
    `status`     tinyint              default 1 comment '角色状态 1启用 0禁用',
    `created_at` timestamp       null default null comment '创建时间',
    `updated_at` timestamp       null default null comment '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `roles_name_unique` (`name`) USING HASH
) ENGINE = InnoDB
  DEFAULT CHARSET utf8mb4
  COLLATE = utf8mb4_general_ci
    COMMENT '角色表';

# 角色用户中间表
CREATE TABLE IF NOT EXISTS user_role
(
    `user_id` bigint unsigned not null,
    `role_id` bigint unsigned not null,
    UNIQUE KEY `role_user_unique` (`user_id`, `role_id`) USING HASH
) ENGINE = InnoDB
  DEFAULT CHARSET utf8mb4
  COLLATE = utf8mb4_general_ci
    COMMENT '用户角色中间表';


# 菜单表
CREATE TABLE IF NOT EXISTS menus
(
    `id`   bigint unsigned not null auto_increment,
    `pid`  bigint unsigned default 0 comment '父级ID',
    `name` varchar(50)     not null comment '显示名称'
) ENGINE = InnoDB
  DEFAULT CHARSET utf8mb4
    COMMENT '菜单表';