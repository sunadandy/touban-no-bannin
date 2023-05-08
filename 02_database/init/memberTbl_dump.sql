-- MySQL dump 10.13  Distrib 5.7.42, for Linux (x86_64)
--
-- Host: localhost    Database: toubanDB
-- ------------------------------------------------------
-- Server version	5.7.42

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `memberTbl`
--

DROP TABLE IF EXISTS `memberTbl`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `memberTbl` (
  `touban_id` int(11) DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL,
  `employee_number` varchar(10) DEFAULT NULL,
  `order_number` smallint(5) unsigned DEFAULT NULL,
  `last` date DEFAULT NULL,
  `next` date DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `affiliation` varchar(10) DEFAULT NULL,
  KEY `link_touban` (`touban_id`),
  CONSTRAINT `link_touban` FOREIGN KEY (`touban_id`) REFERENCES `toubanTbl` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `memberTbl`
--

LOCK TABLES `memberTbl` WRITE;
/*!40000 ALTER TABLE `memberTbl` DISABLE KEYS */;
INSERT INTO `memberTbl` VALUES (10,'賀治 宏亮','E920401',1,'2100-01-01','2023-05-01','','FSE'),(10,'堀田 尊之','E80994',2,'2023-05-07','2023-05-07','','FSE'),(10,'市川 誠','E138420',3,'2100-01-01','2023-06-07','','FSE');
/*!40000 ALTER TABLE `memberTbl` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-05-08  6:29:25
