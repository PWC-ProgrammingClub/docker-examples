CREATE DATABASE IF NOT EXISTS pwpc;
USE pwpc;

-- --------------------------------------------------------


CREATE TABLE IF NOT EXISTS `ideas` (
  `id` int(11) UNIQUE NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `category` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;


INSERT INTO `ideas` (`id`, `category`, `description`) VALUES 
(1, 'projects', 'make more docker containers!');
