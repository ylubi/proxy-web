-- -------------------------------------------
-- proxy参数表
-- -------------------------------------------

DROP TABLE IF EXISTS `parameter`;
CREATE TABLE `parameter` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `protocol` varchar(10) NOT NULL DEFAULT '',
  `proxy_level` tinyint(1) NOT NULL DEFAULT '1',
  `proxy_ip` varchar(25) NOT NULL DEFAULT '',
  `superior_proxy_ip` varchar(25) NOT NULL DEFAULT '',
  `superior` tinyint(1) NOT NULL DEFAULT '1',
  `encryption_condition` varchar(255) NOT NULL DEFAULT '',
  `process_id` integer NOT NULL DEFAULT '0',
  `local` tinyint(1) NOT NULL DEFAULT '1'
);

