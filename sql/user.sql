CREATE TABLE `user`  (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `passwd` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sex` tinyint(4) NOT NULL DEFAULT 0,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `idx_username` (`username`) USING BTREE
)ENGINE =innoDB default charset = utf8mb4 collate = utf8mb4_general_ci;