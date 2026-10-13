CREATE TABLE `example`
(
  `id`          bigint unsigned  NOT NULL,
  `name`        varchar(255)     NOT NULL DEFAULT '' COMMENT '名称',
  `type`        tinyint unsigned NOT NULL DEFAULT 0 COMMENT '类型',
  `create_time` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted`     bool             NOT NULL DEFAULT FALSE,
  `_version`    int              NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
);
