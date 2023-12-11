create database if not exists todolist;
CREATE TABLE todolist.`user`
(
    `uid`        bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `username`   varchar(50)     NOT NULL DEFAULT '' COMMENT 'User name',
    `password`   varchar(50)     not null comment 'password',
    `created_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User information create time',
    `updated_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User information update time',
    `deleted_at` timestamp       NULL     DEFAULT NULL COMMENT 'User information delete time',
    constraint `id`
        PRIMARY KEY (`uid`),
    KEY `idx_name` (`username`, `deleted_at`) COMMENT 'User name index'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

CREATE TABLE todolist.`task`
(
    `id`         bigint       not null auto_increment comment 'pk',
    `uid`        bigint       not null,
    `title`      varchar(255)          default 'default' comment 'title',
    `content`    varchar(255) null comment 'content',
    `status`     int          not null,
    `start_at`   timestamp    not null,
    `end_at`     timestamp    not null,
    `created_at` timestamp    not null default current_timestamp comment 'user task create time',
    `updated_at` timestamp    not null default current_timestamp on update current_timestamp comment 'user task update time',
    `deleted_at` timestamp    null     default null comment 'user task delete time',
    constraint `id`
        primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci