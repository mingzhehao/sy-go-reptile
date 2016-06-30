CREATE TABLE `lagou_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `positionType` varchar(255) DEFAULT NULL,
  `positionName` varchar(255) DEFAULT NULL,
  `workYear` varchar(255) DEFAULT NULL,
  `salary` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `language` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `city` (`city`),
  KEY `language` (`positionName`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8
