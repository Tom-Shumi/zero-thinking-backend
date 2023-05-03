CREATE TABLE IF NOT EXISTS thinking_tree (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `thinking_tree` text NOT NULL,
  `insert_date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3834 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci