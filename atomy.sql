-- MySQL dump 10.13  Distrib 5.7.28, for Linux (aarch64)
--
-- Host: localhost    Database: atomy
-- ------------------------------------------------------
-- Server version	5.7.28-0ubuntu0.18.04.4

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
-- Table structure for table `carts`
--

DROP TABLE IF EXISTS `carts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `carts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` varchar(100) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `carts`
--

LOCK TABLES `carts` WRITE;
/*!40000 ALTER TABLE `carts` DISABLE KEYS */;
/*!40000 ALTER TABLE `carts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `detail` text NOT NULL,
  `total_price` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `products` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `price` varchar(100) DEFAULT NULL,
  `point` int(100) unsigned DEFAULT NULL,
  `pic` longtext,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=260 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'艾多美 艾不釋手經典系列','13670',250000,NULL,'2019-12-22 14:08:21'),(2,'艾多美 托特包','246',0,'https://www.atomy.com/tw/shopping/p_img/100/92434_2.jpg','2020-01-01 06:12:37'),(3,'艾多美 物理性防曬膏','380',5000,'https://www.atomy.com/tw/shopping/p_img/100/00780_2.jpg','2020-01-01 06:12:37'),(4,'艾多美 新春紅包袋','20',0,'https://www.atomy.com/tw/shopping/p_img/100/02704_2.jpg','2020-01-01 06:12:37'),(5,'艾多美 輕柔棉18cm','72',450,'https://www.atomy.com/tw/shopping/p_img/100/90574_2.jpg','2020-01-01 06:12:37'),(6,'艾多美 水嫩亮白組','1399',24700,'https://www.atomy.com/tw/admin/userfiles/201910/hbCare-1_03_2.jpg','2020-01-01 06:12:37'),(7,'艾多美 經典營養霜','688',11700,'https://www.atomy.com/tw/shopping/p_img/100/00198_2.jpg','2020-01-01 06:12:37'),(8,'艾多美 經典乳液','688',11700,'https://www.atomy.com/tw/shopping/p_img/100/00196_2.jpg','2020-01-01 06:12:37'),(9,'艾多美 經典精華液','730',12480,'https://www.atomy.com/tw/shopping/p_img/100/00194_2.jpg','2020-01-01 06:12:37'),(10,'艾多美 經典眼霜','688',11700,'https://www.atomy.com/tw/shopping/p_img/100/00192_2.jpg','2020-01-01 06:12:37'),(11,'艾多美 經典化妝水','688',11700,'https://www.atomy.com/tw/shopping/p_img/100/00190_2.jpg','2020-01-01 06:12:37'),(12,'艾多美 雙色空氣粉餅 No.21','699',8800,'https://www.atomy.com/tw/shopping/p_img/100/00396_2.jpg','2020-01-01 06:12:37'),(13,'艾多美 月桂保濕噴霧','388',5000,'https://www.atomy.com/tw/shopping/p_img/100/00483_2.jpg','2020-01-01 06:12:37'),(14,'艾多美 膠囊葉黃素30','1499',23400,'https://www.atomy.com/tw/shopping/p_img/100/00174_2.gif','2020-01-01 06:12:37'),(15,'艾多美 柔護寶貝三件組','1390',19600,'https://www.atomy.com/tw/shopping/p_img/100/00475_2.jpg','2020-01-01 06:12:37'),(16,'艾多美 銀離子柔濕巾','520',3000,'https://www.atomy.com/tw/shopping/p_img/100/90872_2.jpg','2020-01-01 06:12:37'),(17,'艾多美 養髮液','370',5200,'https://www.atomy.com/tw/shopping/p_img/100/00620_2.jpg','2020-01-01 06:12:37'),(18,'艾多美 好纖果乾-芒果綜合','320',2000,'https://www.atomy.com/tw/shopping/p_img/100/90942_2.jpg','2020-01-01 06:12:37'),(19,'艾多美 凝萃煥髮洗髮乳','525',6750,'https://www.atomy.com/tw/shopping/p_img/100/00694_2.jpg','2020-01-01 06:12:37'),(20,'艾多美 兒童薄膜口罩','250',1100,'https://www.atomy.com/tw/shopping/p_img/100/90809_2.jpg','2020-01-01 06:12:37'),(21,'艾多美 凝萃煥白精華液','1280',23400,'https://www.atomy.com/tw/shopping/p_img/100/00238_2.jpg','2020-01-01 06:12:37'),(22,'艾多美 抽取式廚房紙巾','88',360,'https://www.atomy.com/tw/shopping/p_img/100/90878_2.jpg','2020-01-01 06:12:37'),(23,'艾多美 凝萃煥髮護髮霜','350',4700,'https://www.atomy.com/tw/shopping/p_img/100/00696_2.JPG','2020-01-01 06:12:37'),(24,'艾多美 馬克杯組','380',0,'https://www.atomy.com/tw/shopping/p_img/100/02705_2.jpg','2020-01-01 06:12:37'),(49,'Hi Atomy','55',0,'https://www.atomy.com/tw/shopping/p_img/100/02703_2.jpg','2020-01-01 06:12:37'),(50,'艾多美 凝萃煥髪三部曲','1360',18200,'https://www.atomy.com/tw/shopping/p_img/100/00657_2.jpg','2020-01-01 06:12:37'),(51,'艾多美 香烤海苔(大片裝)','216',1300,'https://www.atomy.com/tw/shopping/p_img/100/90906_2.jpg','2020-01-01 06:12:37'),(52,'艾多美 凝萃煥膚防曬乳','355',4560,'https://www.atomy.com/tw/shopping/p_img/100/00276_2.jpg','2020-01-01 06:12:37'),(53,'艾多美 長效睫毛膏','370',5160,'https://www.atomy.com/tw/shopping/p_img/100/01422_2.jpg','2020-01-01 06:12:37'),(54,'艾多美 Daily沁涼舒緩面膜','326',4200,'https://www.atomy.com/tw/shopping/p_img/100/00787_2.jpg','2020-01-01 06:12:37'),(55,'艾多美 Daily亮白水嫩面膜','326',4200,'https://www.atomy.com/tw/shopping/p_img/100/00786_2.jpg','2020-01-01 06:12:37'),(56,'艾多美 Daily緊緻保濕面膜','326',4200,'https://www.atomy.com/tw/shopping/p_img/100/00785_2.jpg','2020-01-01 06:12:37'),(57,'艾多美 氣墊粉撲','80',600,'https://www.atomy.com/tw/shopping/p_img/100/01499_2.jpg','2020-01-01 06:12:37'),(58,'艾多美 經典保養五件組','3188',57200,'https://www.atomy.com/tw/shopping/p_img/100/00003_2.jpg','2020-01-01 06:12:37'),(59,'艾多美 深層卸妝油','520',6800,'https://www.atomy.com/tw/shopping/p_img/100/00482_2.jpg','2020-01-01 06:12:37'),(60,'艾多美 蜂蜜飲(Atomy HemoHIM) 4組','11280',198000,'https://www.atomy.com/tw/shopping/p_img/100/00002_2.jpg','2020-01-01 06:12:37'),(61,'艾多美 蜂蜜飲(Atomy HemoHIM) 1組','3260',54600,'https://www.atomy.com/tw/shopping/p_img/100/00001_2.jpg','2020-01-01 06:12:37'),(62,'艾多美 夾心海苔燒(2盒裝)','900',8000,'https://www.atomy.com/tw/shopping/p_img/100/90905_2.jpg','2020-01-01 06:12:37'),(63,'艾多美 維他命C粉','680',8400,'https://www.atomy.com/tw/shopping/p_img/100/00121_2.jpg','2020-01-01 06:12:37'),(64,'艾多美 空汙防護口罩','480',3000,'https://www.atomy.com/tw/shopping/p_img/100/90808_2.jpg','2020-01-01 06:12:37'),(65,'艾多美 幸福堅果','1350',4500,'https://www.atomy.com/tw/shopping/p_img/100/00982_2.jpg','2020-01-01 06:12:37'),(66,'艾多美 凝萃煥膚六部曲','6800',130000,'https://www.atomy.com/tw/shopping/p_img/100/00207_2.jpg','2020-01-01 06:12:37'),(67,'艾多美 潤色護唇膏','385',4920,'https://www.atomy.com/tw/shopping/p_img/100/00400_2.jpg','2020-01-01 06:12:37'),(68,'艾多美 隨行袖珍包','72',250,'https://www.atomy.com/tw/shopping/p_img/100/90879_2.jpg','2020-01-01 06:12:37'),(69,'艾多美 珠光隔離乳','488',5000,'https://www.atomy.com/tw/shopping/p_img/100/00266_2.jpg','2020-01-01 06:12:37'),(70,'艾多美 保溫隨行杯','650',4500,'https://www.atomy.com/tw/shopping/p_img/100/01812_2.jpg','2020-01-01 06:12:37'),(71,'艾多美 阿拉比卡三合一即溶咖啡50入','280',1300,'https://www.atomy.com/tw/shopping/p_img/100/00973_2.jpg','2020-01-01 06:12:37'),(72,'艾多美 抽取式洗碗布','120',980,'https://www.atomy.com/tw/shopping/p_img/100/00836_2.jpg','2020-01-01 06:12:37'),(97,'艾多美 即溶普洱茶','920',11000,'https://www.atomy.com/tw/shopping/p_img/100/00180_2.jpg','2020-01-01 06:12:37'),(98,'艾多美 螺旋藻膠囊 100%','780',8000,'https://www.atomy.com/tw/shopping/p_img/100/00178_2.jpg','2020-01-01 06:12:37'),(99,'艾多美 304不鏽鋼絲球','85',750,'https://www.atomy.com/tw/shopping/p_img/100/00835_2.jpg','2020-01-01 06:12:37'),(100,'艾多美 乳膠手套(M)','150',1300,'https://www.atomy.com/tw/admin/userfiles/201803/lglovesL_750_02_1.jpg','2020-01-01 06:12:37'),(101,'艾多美 洋薊膠囊','1080',9000,'https://www.atomy.com/tw/shopping/p_img/100/90174_2.jpg','2020-01-01 06:12:37'),(102,'艾多美 凝萃煥膚 營養霜','1125',20800,NULL,'2019-12-22 14:08:22'),(103,'艾多美 凝萃煥膚 眼霜','1265',23400,'https://www.atomy.com/tw/shopping/p_img/100/00257_2.jpg','2020-01-01 06:12:37'),(104,'艾多美 凝萃煥膚 乳液','1125',20800,'https://www.atomy.com/tw/shopping/p_img/100/00247_2.jpg','2020-01-01 06:12:37'),(105,'艾多美 凝萃煥膚 精華液','1265',23400,'https://www.atomy.com/tw/shopping/p_img/100/00237_2.jpg','2020-01-01 06:12:37'),(106,'艾多美 凝萃煥膚 化妝水','1125',20800,'https://www.atomy.com/tw/shopping/p_img/100/00217_2.jpg','2020-01-01 06:12:37'),(107,'艾多美 鳳梨綜合酵素粉','920',11000,'https://www.atomy.com/tw/shopping/p_img/100/00179_2.jpg','2020-01-01 06:12:37'),(108,'艾多美 凝萃煥膚 安瓶','1365',26000,'https://www.atomy.com/tw/shopping/p_img/100/00227_2.jpg','2020-01-01 06:12:37'),(109,'艾多美 精油 貼布 1盒(10包*5片)','780',10000,'https://www.atomy.com/tw/shopping/p_img/100/00591_2.gif','2020-01-01 06:12:37'),(110,'艾多美 花漾口紅-粉紅絲絨','460',6100,'https://www.atomy.com/tw/shopping/p_img/100/00413_2.jpg','2020-01-01 06:12:37'),(111,'艾多美 頭皮舒爽洗髮露','425',5600,'https://www.atomy.com/tw/shopping/p_img/100/00691_2.jpg','2020-01-01 06:12:37'),(112,'艾多美 男士精華液 1瓶','430',3750,'https://www.atomy.com/tw/shopping/p_img/100/01231_2.jpg','2020-01-01 06:12:37'),(113,'艾多美 男士乳液 1瓶','425',4300,NULL,'2019-12-22 14:08:22'),(114,'艾多美 男士化妝水 1瓶','420',4300,'https://www.atomy.com/tw/shopping/p_img/100/01211_2.jpg','2020-01-01 06:12:37'),(115,'艾多美 316不鏽鋼湯鍋5.5公升','5480',82000,'https://www.atomy.com/tw/shopping/p_img/100/00884_2.jpg','2020-01-01 06:12:37'),(116,'艾多美 316不鏽鋼湯鍋3.3公升','5200',70000,'https://www.atomy.com/tw/shopping/p_img/100/00883_2.jpg','2020-01-01 06:12:37'),(117,'艾多美 316不鏽鋼湯鍋2.2公升','4520',57000,'https://www.atomy.com/tw/shopping/p_img/100/00882_2.jpg','2020-01-01 06:12:37'),(118,'艾多美 316不鏽鋼單柄鍋1.8公升','3820',51000,'https://www.atomy.com/tw/shopping/p_img/100/00881_2.jpg','2020-01-01 06:12:37'),(119,'艾多美 柔護寶貝洗髮沐浴露','460',6200,'https://www.atomy.com/tw/shopping/p_img/100/00474_2.jpg','2020-01-01 06:12:37'),(120,'艾多美 柔護寶貝保濕乳霜','480',6500,'https://www.atomy.com/tw/shopping/p_img/100/00473_2.jpg','2020-01-01 06:12:37'),(145,'艾多美 柔護寶貝潤膚乳','520',6900,'https://www.atomy.com/tw/shopping/p_img/100/00472_2.jpg','2020-01-01 06:12:37'),(146,'艾多美 香烤海苔(小片裝) 1箱 (1箱4盒)','998',6000,'https://www.atomy.com/tw/shopping/p_img/100/00904_2.jpg','2020-01-01 06:12:37'),(147,'艾多美 輕柔棉28cm','238',1800,'https://www.atomy.com/tw/shopping/p_img/100/00564_2.jpg','2020-01-01 06:12:37'),(148,'艾多美 輕柔棉24cm','188',1500,'https://www.atomy.com/tw/shopping/p_img/100/00554_2.jpg','2020-01-01 06:12:37'),(149,'艾多美 溫和潔面慕斯','288',4500,'https://www.atomy.com/tw/shopping/p_img/100/00471_2.jpg','2020-01-01 06:12:37'),(150,'艾多美 洗手乳','145',1000,'https://www.atomy.com/tw/shopping/p_img/100/00541_2.jpg','2020-01-01 06:12:37'),(151,'艾多美 小麥馬鈴薯杯麵','980',3100,NULL,'2019-12-22 14:08:22'),(152,'艾多美 蜂膠牙膏50g 1組(4條)','150',3000,'https://www.atomy.com/tw/shopping/p_img/100/00521_2.jpg','2020-01-01 06:12:37'),(153,'艾多美 牙刷(小型刷頭) 1組 (8支)','260',3900,'https://www.atomy.com/tw/shopping/p_img/100/00511_2.jpg','2020-01-01 06:12:37'),(154,'艾多美 牙刷(大型刷頭) 1組 (8支)','260',5000,'https://www.atomy.com/tw/shopping/p_img/100/00510_2.jpg','2020-01-01 06:12:37'),(155,'艾多美 蜂膠牙膏200g(5條)','485',3250,'https://www.atomy.com/tw/shopping/p_img/100/00505_2.jpg','2020-01-01 06:12:37'),(156,'艾多美 溫和卸妝水','420',6300,'https://www.atomy.com/tw/shopping/p_img/100/00481_2.jpg','2020-01-01 06:12:37'),(157,'艾多美 水嫩潤唇膏','200',2200,'https://www.atomy.com/tw/shopping/p_img/100/00460_2.jpg','2020-01-01 06:12:37'),(158,'艾多美 潤絲乳','335',5200,'https://www.atomy.com/tw/shopping/p_img/100/00665_2.jpg','2020-01-01 06:12:37'),(159,'艾多美 身體乳液 1瓶','285',3200,'https://www.atomy.com/tw/shopping/p_img/100/00645_2.jpg','2020-01-01 06:12:37'),(160,'艾多美 沐浴乳','270',3200,'https://www.atomy.com/tw/shopping/p_img/100/00635_2.jpg','2020-01-01 06:12:37'),(161,'艾多美 洗髮乳','335',5200,'https://www.atomy.com/tw/shopping/p_img/100/00605_2.jpg','2020-01-01 06:12:37'),(162,'艾多美 愛丹','178',2000,'https://www.atomy.com/tw/shopping/p_img/100/00535_2.JPG','2020-01-01 06:12:37'),(163,'艾多美 氣墊粉底 NO.23 SPF50+ PA+++','880',15600,'https://www.atomy.com/tw/shopping/p_img/100/01405_2.jpg','2020-01-01 06:12:37'),(164,'艾多美 氣墊粉底 NO.21 SPF50+ PA+++','880',15600,'https://www.atomy.com/tw/shopping/p_img/100/01404_2.jpg','2020-01-01 06:12:37'),(165,'艾多美 阿拉比卡即溶黑咖啡','460',3900,'https://www.atomy.com/tw/shopping/p_img/100/00975_2.jpg','2020-01-01 06:12:37'),(166,'艾多美 阿拉比卡三合一即溶咖啡','980',3900,'https://www.atomy.com/tw/shopping/p_img/100/00971_2.jpg','2020-01-01 06:12:37'),(167,'艾多美 濃縮衣物洗潔液','390',4400,'https://www.atomy.com/tw/shopping/p_img/100/00861_2.jpg','2020-01-01 06:12:37'),(168,'艾多美 濃縮洗衣粉','420',3600,'https://www.atomy.com/tw/shopping/p_img/100/00821_2.jpg','2020-01-01 06:12:37'),(193,'艾多美 碗盤蔬果洗潔液','275',3100,'https://www.atomy.com/tw/shopping/p_img/100/00801_2.jpg','2020-01-01 06:12:37'),(194,'艾多美 清潔護膚旅行組 1組','400',3500,'https://www.atomy.com/tw/shopping/p_img/100/00553_2.jpg','2020-01-01 06:12:37'),(195,'艾多美 肌膚保養旅行組 1組','890',8000,'https://www.atomy.com/tw/shopping/p_img/100/00552_2.jpg','2020-01-01 06:12:37'),(196,'艾多美 護手霜四件組 1組(4條)','480',6500,'https://www.atomy.com/tw/shopping/p_img/100/00545_2.jpg','2020-01-01 06:12:37'),(197,'艾多美 清潔護膚四件組','1020',13000,'https://www.atomy.com/tw/shopping/p_img/100/00355_2.jpg','2020-01-01 06:12:37'),(198,'艾多美 炸醬麵*1箱(16包)','1010',3900,'https://www.atomy.com/tw/shopping/p_img/100/00932_2.jpg','2020-01-01 06:12:37'),(199,'艾多美 濃縮衣物柔軟精','310',3400,'https://www.atomy.com/tw/shopping/p_img/100/00871_2.jpg','2020-01-01 06:12:37'),(200,'艾多美 貢布黑胡椒粒','280',1300,NULL,'2019-12-22 14:08:23'),(201,'艾多美 馬鈴薯蔬菜拉麵*1箱','1380',5200,'https://www.atomy.com/tw/shopping/p_img/100/00930_2.jpg','2020-01-01 06:12:37'),(202,'艾多美 316不鏽鋼炒鍋5.4公升','5980',87000,'https://www.atomy.com/tw/shopping/p_img/100/00880_2.jpg','2020-01-01 06:12:37'),(203,'艾多美 卸妝乳','255',3250,'https://www.atomy.com/tw/shopping/p_img/100/00325_2.jpg','2020-01-01 06:12:37'),(204,'艾多美 洗面乳','255',3250,'https://www.atomy.com/tw/shopping/p_img/100/00305_2.jpg','2020-01-01 06:12:37'),(205,'艾多美 玫瑰舒緩噴霧 1瓶','325',4200,'https://www.atomy.com/tw/shopping/p_img/100/00295_2.jpg','2020-01-01 06:12:37'),(206,'艾多美 防曬霜(裸膚)','255',3250,'https://www.atomy.com/tw/shopping/p_img/100/00285_2.jpg','2020-01-01 06:12:37'),(207,'艾多美 防曬霜(白皙)','255',3250,'https://www.atomy.com/tw/shopping/p_img/100/00275_2.jpg','2020-01-01 06:12:37'),(208,'艾多美 BB霜','255',2500,'https://www.atomy.com/tw/shopping/p_img/100/00265_2.jpg','2020-01-01 06:12:37'),(209,'艾多美 護髮油','360',5200,'https://www.atomy.com/tw/shopping/p_img/100/00681_2.jpg','2020-01-01 06:12:37'),(210,'艾多美7Solutions水凝膠面膜','1000',12600,'https://www.atomy.com/tw/shopping/p_img/100/00461_2.jpg','2020-01-01 06:12:37'),(211,'艾多美 剝離式面膜','255',3250,'https://www.atomy.com/tw/shopping/p_img/100/00335_2.jpg','2020-01-01 06:12:37'),(212,'艾多美 去角質凝膠','255',3250,'https://www.atomy.com/tw/shopping/p_img/100/00315_2.jpg','2020-01-01 06:12:37'),(213,'艾多美 攜帶式口腔保健組 1組 (4個)','400',6250,'https://www.atomy.com/tw/shopping/p_img/100/00520_2.jpg','2020-01-01 06:12:37'),(214,'艾多美口腔護理禮盒','680',8400,'https://www.atomy.com/tw/shopping/p_img/100/00514_2.jpg','2020-01-01 06:12:37'),(215,'艾多美 綜合穀物飲','1500',14300,'https://www.atomy.com/tw/shopping/p_img/100/00965_2.jpg','2020-01-01 06:12:37'),(216,'艾多美 舒緩修護三件組','1790',29000,NULL,'2019-12-22 14:08:23'),(241,'艾多美 控油調理三件組','1790',29000,'https://www.atomy.com/tw/shopping/p_img/100/00765_2.jpg','2020-01-01 06:12:37'),(242,'艾多美 舒緩修護乳液','545',9000,NULL,'2020-01-01 05:30:32'),(243,'艾多美 控油調理乳液','655',10000,'https://www.atomy.com/tw/shopping/p_img/100/00735_2.jpg','2020-01-01 06:12:37'),(244,'艾多美 控油調理精華','545',9000,'https://www.atomy.com/tw/shopping/p_img/100/00725_2.jpg','2020-01-01 06:12:37'),(245,'艾多美 舒緩修護精華','660',10000,NULL,'2019-12-22 14:08:24'),(246,'艾多美 修護化妝水','595',10000,NULL,'2019-12-22 14:08:24'),(247,'艾多美 環保衛生紙','688',3000,'https://www.atomy.com/tw/shopping/p_img/100/90877_2.jpg','2020-01-01 06:00:57'),(248,'艾多美 益生菌(Probiotics10+)','2088',26000,'https://www.atomy.com/tw/shopping/p_img/100/00164_2.jpg','2020-01-01 06:12:37'),(249,'艾多美 葉黃素膠囊','1080',12000,'https://www.atomy.com/tw/Home/Product/ProductView?GdsCode=W00173','2020-01-01 06:12:37'),(250,'艾多美 高濃度魚油膠囊','825',8300,'https://www.atomy.com/tw/shopping/p_img/100/00111_2.jpg','2020-01-01 06:12:37'),(251,'艾多美 男士保養三件組','1140',12350,'https://www.atomy.com/tw/shopping/p_img/100/01205_2.jpg','2020-01-01 06:12:37'),(252,'會員加入申請書1ea','50',0,'https://www.atomy.com/tw/shopping/p_img/100/02406_2.jpg','2020-01-01 06:12:37'),(253,'艾多美 產品介紹摺頁(10張)','30',0,'https://www.atomy.com/tw/shopping/p_img/100/02618_2.jpg','2020-01-01 06:12:37'),(254,'艾多美產品型錄*1本','30',0,NULL,'2019-12-22 14:08:24'),(255,'塑膠購物袋(大)*10個','140',0,'https://www.atomy.com/tw/shopping/p_img/100/02438_2.jpg','2020-01-01 06:12:37'),(256,'塑膠購物袋(大)*1個','14',0,'https://www.atomy.com/tw/shopping/p_img/100/02431_2.jpg','2020-01-01 06:12:37'),(257,'紙購物袋 *10個','100',0,'https://www.atomy.com/tw/shopping/p_img/100/02414_2.jpg','2020-01-01 06:12:37'),(258,'紙購物袋 *1個','10',0,'https://www.atomy.com/tw/shopping/p_img/100/02404_2.jpg','2020-01-01 06:12:37'),(259,'艾多美 紅蔘 1盒','3080',33000,'https://www.atomy.com/tw/admin/userfiles/201904/redginseng_750_01.gif','2020-01-01 06:12:37');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-01-05  8:29:46
