# ************************************************************
# Sequel Pro SQL dump
# Version 4096
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Hôte: localhost (MySQL 5.5.34)
# Base de données: cme_test
# Temps de génération: 2014-05-06 03:30:15 PM +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Affichage de la table forum
# ------------------------------------------------------------

DROP TABLE IF EXISTS `forum`;

CREATE TABLE `forum` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `forum_category_id` int(10) unsigned NOT NULL DEFAULT '5',
  `title` varchar(100) NOT NULL,
  `text` text NOT NULL,
  `keywords` varchar(200) NOT NULL,
  `is_solved` smallint(6) NOT NULL DEFAULT '0',
  `is_online` smallint(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

LOCK TABLES `forum` WRITE;
/*!40000 ALTER TABLE `forum` DISABLE KEYS */;

INSERT INTO `forum` (`id`, `user_id`, `forum_category_id`, `title`, `text`, `keywords`, `is_solved`, `is_online`, `created_at`, `updated_at`)
VALUES
	(1,0,5,'Titre A','Titre A : Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',1,1,'2014-04-15 09:05:12','2014-04-15 09:05:12'),
	(2,0,4,'Titre B','Doudoud didi Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',0,0,'2014-04-15 09:06:09','2014-04-15 09:06:09'),
	(3,0,4,'Titre C','objet C Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',1,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(4,0,4,'Titre D','coucou de la forêt','',0,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(5,0,3,'Titre E','objet E Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',1,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(6,0,2,'Titre F','objet F Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',0,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(7,0,1,'Titre G','objet G Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',0,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(8,0,1,'Titre H','objet H Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',0,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(9,0,1,'Comment faire pour rien faire ?','objet I Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',1,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(10,0,1,'Titre J','objet J Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',1,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(11,0,1,'Titre K','objet K Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',0,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(12,0,1,'Titre L','objet L Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.','',0,1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(13,0,0,'qsdfqsdf','<p>qsdfqsdf</p>\r\n','',0,0,'2014-04-17 16:08:30','2014-04-17 16:08:30'),
	(14,0,2,'Mon super titre','<p>La question de la mort : &nbsp;Comment je fait pour r&eacute;cup&eacute;rer un post d&#39;un formulaire ?</p>\r\n','',0,0,'2014-04-17 16:16:50','2014-04-17 16:16:50'),
	(15,0,2,'titi','<p>coucou comment sa vas ?</p>\r\n','',0,1,'2014-04-17 17:27:19','2014-04-17 17:27:19'),
	(16,0,2,'test coucou','<p>ceci est un autre test&nbsp;</p>\r\n','',0,1,'2014-04-22 09:03:38','2014-04-22 09:03:38'),
	(17,0,1,'coucou hu hu','<p>ha ha ho ho</p>\r\n','',0,1,'2014-04-23 08:49:23','2014-04-23 08:49:23'),
	(18,0,3,'qsfqsdfqs','<p>sdfqsdfqsdf</p>\r\n','',0,1,'2014-04-23 16:24:27','2014-04-23 16:24:27'),
	(19,0,3,'coucouc','<p>hihi&nbsp;</p>\r\n','',0,1,'2014-04-24 09:07:26','2014-04-24 09:07:26'),
	(20,0,1,'Bonjour les amis','<p>Comment allez-vous?</p>\r\n','',0,1,'2014-04-25 12:13:27','2014-04-25 12:13:27'),
	(21,0,1,'Bonjour les amis','<p>Comment allez-vous?</p>\r\n','',0,1,'2014-04-25 12:16:00','2014-04-25 12:16:00');

/*!40000 ALTER TABLE `forum` ENABLE KEYS */;
UNLOCK TABLES;


# Affichage de la table forum_category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `forum_category`;

CREATE TABLE `forum_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

LOCK TABLES `forum_category` WRITE;
/*!40000 ALTER TABLE `forum_category` DISABLE KEYS */;

INSERT INTO `forum_category` (`id`, `title`, `created_at`, `updated_at`)
VALUES
	(1,'Packaging','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(2,'Web','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(3,'Rough','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(4,'Cours','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(5,'Autre','0000-00-00 00:00:00','0000-00-00 00:00:00');

/*!40000 ALTER TABLE `forum_category` ENABLE KEYS */;
UNLOCK TABLES;


# Affichage de la table forum_post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `forum_post`;

CREATE TABLE `forum_post` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `forum_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `text` text NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

LOCK TABLES `forum_post` WRITE;
/*!40000 ALTER TABLE `forum_post` DISABLE KEYS */;

INSERT INTO `forum_post` (`id`, `forum_id`, `user_id`, `text`, `is_online`, `created_at`, `updated_at`)
VALUES
	(1,1,1,'voici ma réponse A',1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(2,1,2,'voici ma réponse B',1,'0000-00-00 00:00:00','0000-00-00 00:00:00');

/*!40000 ALTER TABLE `forum_post` ENABLE KEYS */;
UNLOCK TABLES;


# Affichage de la table news
# ------------------------------------------------------------

DROP TABLE IF EXISTS `news`;

CREATE TABLE `news` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `id_user` int(10) unsigned NOT NULL,
  `news_category_id` int(10) unsigned NOT NULL,
  `title` varchar(100) NOT NULL,
  `text` text NOT NULL,
  `keywords` varchar(200) NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

LOCK TABLES `news` WRITE;
/*!40000 ALTER TABLE `news` DISABLE KEYS */;

INSERT INTO `news` (`id`, `id_user`, `news_category_id`, `title`, `text`, `keywords`, `is_online`, `created_at`, `updated_at`)
VALUES
	(1,1,0,'1','Aenean lacinia bibendum nulla sed consectetur. Curabitur blandit tempus porttitor. Vestibulum id ligula porta felis euismod semper. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Donec sed odio dui. Donec sed odio dui.\r\rNullam id dolor id nibh ultricies vehicula ut id elit. Maecenas sed diam eget risus varius blandit sit amet non magna. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Etiam porta sem malesuada magna mollis euismod. Donec ullamcorper nulla non metus auctor fringilla.','',1,'2014-04-28 10:30:00','0000-00-00 00:00:00'),
	(2,1,0,'2','Aenean lacinia bibendum nulla sed consectetur. Curabitur blandit tempus porttitor. Vestibulum id ligula porta felis euismod semper. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Donec sed odio dui. Donec sed odio dui.\r\rNullam id dolor id nibh ultricies vehicula ut id elit. Maecenas sed diam eget risus varius blandit sit amet non magna. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Etiam porta sem malesuada magna mollis euismod. Donec ullamcorper nulla non metus auctor fringilla.','',1,'2014-04-28 10:30:00','0000-00-00 00:00:00'),
	(3,1,0,'3','Aenean lacinia bibendum nulla sed consectetur. Curabitur blandit tempus porttitor. Vestibulum id ligula porta felis euismod semper. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Donec sed odio dui. Donec sed odio dui.\r\rNullam id dolor id nibh ultricies vehicula ut id elit. Maecenas sed diam eget risus varius blandit sit amet non magna. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Etiam porta sem malesuada magna mollis euismod. Donec ullamcorper nulla non metus auctor fringilla.','',1,'2014-04-28 10:30:00','0000-00-00 00:00:00'),
	(4,1,0,'4','Aenean lacinia bibendum nulla sed consectetur. Curabitur blandit tempus porttitor. Vestibulum id ligula porta felis euismod semper. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Donec sed odio dui. Donec sed odio dui.\r\rNullam id dolor id nibh ultricies vehicula ut id elit. Maecenas sed diam eget risus varius blandit sit amet non magna. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Etiam porta sem malesuada magna mollis euismod. Donec ullamcorper nulla non metus auctor fringilla.','',1,'2014-04-28 12:27:00','0000-00-00 00:00:00'),
	(5,1,0,'5','Aenean lacinia bibendum nulla sed consectetur. Curabitur blandit tempus porttitor. Vestibulum id ligula porta felis euismod semper. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Donec sed odio dui. Donec sed odio dui.\r\rNullam id dolor id nibh ultricies vehicula ut id elit. Maecenas sed diam eget risus varius blandit sit amet non magna. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Etiam porta sem malesuada magna mollis euismod. Donec ullamcorper nulla non metus auctor fringilla.','',1,'2014-04-28 12:28:00','0000-00-00 00:00:00'),
	(6,1,0,'6','Aenean lacinia bibendum nulla sed consectetur. Curabitur blandit tempus porttitor. Vestibulum id ligula porta felis euismod semper. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Donec sed odio dui. Donec sed odio dui.\r\rNullam id dolor id nibh ultricies vehicula ut id elit. Maecenas sed diam eget risus varius blandit sit amet non magna. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Etiam porta sem malesuada magna mollis euismod. Donec ullamcorper nulla non metus auctor fringilla.','',1,'2014-04-28 12:29:00','0000-00-00 00:00:00'),
	(7,1,0,'7','Aenean lacinia bibendum nulla sed consectetur. Curabitur blandit tempus porttitor. Vestibulum id ligula porta felis euismod semper. Integer posuere erat a ante venenatis dapibus posuere velit aliquet. Donec sed odio dui. Donec sed odio dui.\r\rNullam id dolor id nibh ultricies vehicula ut id elit. Maecenas sed diam eget risus varius blandit sit amet non magna. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Etiam porta sem malesuada magna mollis euismod. Donec ullamcorper nulla non metus auctor fringilla.','',1,'2014-04-28 12:30:00','0000-00-00 00:00:00');

/*!40000 ALTER TABLE `news` ENABLE KEYS */;
UNLOCK TABLES;


# Affichage de la table news_category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `news_category`;

CREATE TABLE `news_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Affichage de la table news_post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `news_post`;

CREATE TABLE `news_post` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `news_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `text` text NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Affichage de la table tag
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tag`;

CREATE TABLE `tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(25) NOT NULL,
  `is_online` int(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


# Affichage de la table page_simple
# ------------------------------------------------------------

DROP TABLE IF EXISTS `page_simple`;

CREATE TABLE `page_simple` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(100) NOT NULL,
  `title` varchar(100) NOT NULL,
  `text` text NOT NULL,
  `is_online` int(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=3 ;


LOCK TABLES `page_simple` WRITE;
/*!40000 ALTER TABLE `page_simple` DISABLE KEYS */;

INSERT INTO `page_simple` (`id`, `url`, `title`, `text`, `is_online`, `created_at`, `updated_at`) VALUES
(1, 'pourquoi-une-association', 'Pourquoi une association ?', '<img class="img-info" src="/img/pq_une_asso_large.jpg" alt="Qui sommes nous ?">\r\n\r\n                            <p>Nullam quis risus eget urna mollis ornare vel eu leo. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Donec id elit non mi porta gravida at eget metus.</p>\r\n                            <p>Nullam quis risus eget urna mollis ornare vel eu leo. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Donec id elit non mi porta gravida at eget metus.</p>\r\n                            <p>Nullam quis risus eget urna mollis ornare vel eu leo. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Donec id elit non mi porta gravida at eget metus.</p>', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 'qui-sommes-nous', 'Qui sommes-nous ?', '<img class="img-info" src="/img/qui_sommes_nous_large.jpg" alt="Qui sommes nous ?">\r\n\r\n                            <p>Nullam quis risus eget urna mollis ornare vel eu leo. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Donec id elit non mi porta gravida at eget metus.</p>\r\n                            <p>Nullam quis risus eget urna mollis ornare vel eu leo. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Donec id elit non mi porta gravida at eget metus.</p>\r\n                            <p>Nullam quis risus eget urna mollis ornare vel eu leo. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Donec id elit non mi porta gravida at eget metus.</p>', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00');

/*!40000 ALTER TABLE `page_simple` ENABLE KEYS */;
UNLOCK TABLES;


# Affichage de la table tutorial
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tutorial`;

CREATE TABLE `tutorial` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `tutorial_category_id` int(10) unsigned NOT NULL,
  `title` varchar(100) NOT NULL,
  `text` text NOT NULL,
  `keywords` varchar(200) NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

LOCK TABLES `tutorial` WRITE;
/*!40000 ALTER TABLE `tutorial` DISABLE KEYS */;

INSERT INTO `tutorial` (`id`, `user_id`, `tutorial_category_id`, `title`, `text`, `keywords`, `is_online`, `created_at`, `updated_at`)
VALUES
	(1,0,1,'Comment créer un effet néon sur du lettrage','Un petit tutoriel très simple qui explique comment générer un effet de lettrage néon, de style un peu rétro, avec les effets de Photoshop.','',1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(2,0,2,'Comment utiliser le Pathfinder dans Illustrator','Outil indispensable à tous les utilisateurs de Illustrator! Je vous explique ici comment maîtriser le Pathfinder.','',1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(3,0,3,'Comment générer du CSS avec SASS/SCSS dans Sublime Text 2','SASS/SCSS s\'incruste dans les bonnes habitudes des intégrateurs de ce monde à cause de ses fonctionnalités très utiles. Je vous montre ici comment générer du CSS compressé à partir de SASS avec Sublime Text 2.','',1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(4,0,2,'coucou voici mon super tuto','<p>Illustrator est un outils super simpa super tuto hihi</p>\r\n','',1,'2014-04-24 13:47:11','2014-04-24 13:47:11');

/*!40000 ALTER TABLE `tutorial` ENABLE KEYS */;
UNLOCK TABLES;


# Affichage de la table tutorial_category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tutorial_category`;

CREATE TABLE `tutorial_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8 NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf32;

LOCK TABLES `tutorial_category` WRITE;
/*!40000 ALTER TABLE `tutorial_category` DISABLE KEYS */;

INSERT INTO `tutorial_category` (`id`, `title`, `created_at`, `updated_at`)
VALUES
	(1,'Photoshop','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(2,'Illustator','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(3,'Sublime Text','0000-00-00 00:00:00','0000-00-00 00:00:00');

/*!40000 ALTER TABLE `tutorial_category` ENABLE KEYS */;
UNLOCK TABLES;


# Affichage de la table tutorial_post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tutorial_post`;

CREATE TABLE `tutorial_post` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tutorial_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `text` text NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Affichage de la table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `text` text NOT NULL,
  `email` varchar(100) NOT NULL,
  `pass` varchar(200) NOT NULL,
  `Facebook` text,
  `Twitter` text,
  `LinkedIn` text,
  `keywords` varchar(200) NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`id`, `first_name`, `last_name`, `text`, `email`, `pass`, `Facebook`, `Twitter`, `LinkedIn`, `keywords`, `is_online`, `created_at`, `updated_at`)
VALUES
	(1,'Damien','Weil','texte de Damien Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.','damien.weil@gmail.com','',NULL,NULL,NULL,'',1,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(2,'René','Grossi','texte de rené Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.','rene.grossi@gmail.com','',NULL,NULL,NULL,'',1,'2014-04-07 16:35:14','2014-04-07 16:35:14'),
	(3,'Antoine','Chameau','texte d\'antoine Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.','antoine.chameau@gmail.com','',NULL,NULL,NULL,'',1,'2014-04-07 16:35:38','2014-04-07 16:35:38'),
	(4,'Henri','Lepic','texte d\'Henri Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.','henrilepic@gmail.com','',NULL,NULL,NULL,'',1,'2014-04-07 16:36:41','2014-04-07 16:36:41'),
	(5,'Amélie','Lepic','texte Amélie Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.','amelie.lepic@gmail.com','',NULL,NULL,NULL,'',0,'2014-04-07 17:04:15','2014-04-07 17:04:15'),
	(6,'Hugues','Lepic','texte d\'hugues Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.','h.lepic@gmail.com','',NULL,NULL,NULL,'',1,'2014-04-07 17:08:35','2014-04-07 17:08:35'),
	(7,'Antoine','Lord','Texte d\'Antoine Lord - Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Nulla vitae elit libero, a pharetra augue. Donec id elit non mi porta gravida at eget metus. Cras justo odio, dapibus ac facilisis in, egestas eget quam.\r\rCras justo odio, dapibus ac facilisis in, egestas eget quam. Donec id elit non mi porta gravida at eget metus. Donec id elit non mi porta gravida at eget metus. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Aenean lacinia bibendum nulla sed consectetur. Aenean lacinia bibendum nulla sed consectetur.','antoinelord@gmail.com','','http://fb.com/antoine.lord.50','http://twitter.com/antoinelord',NULL,'',1,'2014-05-06 00:00:00','0000-00-00 00:00:00');

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;


# Affichage de la table user_image
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_image`;

CREATE TABLE `user_image` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `user_image_category_id` int(10) unsigned NOT NULL,
  `title` varchar(100) NOT NULL,
  `url` varchar(200) NOT NULL,
  `description` text NOT NULL,
  `is_featured` int(1) unsigned NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `Year` int(4) NOT NULL,
  `Month` int(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8;

LOCK TABLES `user_image` WRITE;
/*!40000 ALTER TABLE `user_image` DISABLE KEYS */;

INSERT INTO `user_image` (`id`, `user_id`, `user_image_category_id`, `title`, `url`, `description`, `is_featured`, `is_online`, `created_at`, `updated_at`, `Year`, `Month`)
VALUES
	(1,7,0,'Antoine Lord','http://google.fr','Cras mattis consectetur purus sit amet fermentum. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aenean lacinia bibendum nulla sed consectetur.',0,1,'2014-04-07 17:04:15','2014-04-07 17:04:15',2014,5),
	(2,6,0,'Test titre...','','',0,1,'2014-04-07 17:08:35','2014-04-07 17:08:35',2014,5),
	(3,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(4,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(5,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(6,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(7,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(8,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(9,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(10,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(11,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(12,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(13,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(14,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(15,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(16,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(17,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(18,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(19,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(20,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(21,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(22,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(23,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(24,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(25,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(26,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(27,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(28,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(29,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5),
	(30,1,0,'Justo Consectetur','/','Etiam porta sem malesuada magna mollis euismod. Duis mollis, est non commodo luctus, nisi erat porttitor ligula, eget lacinia odio sem nec elit. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Maecenas sed diam eget risus varius blandit sit amet non magna. Maecenas faucibus mollis interdum. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer posuere erat a ante venenatis dapibus posuere velit aliquet.',0,1,'2014-05-00 00:00:00','0000-00-00 00:00:00',2014,5);

/*!40000 ALTER TABLE `user_image` ENABLE KEYS */;
UNLOCK TABLES;


# Affichage de la table user_image_category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_image_category`;

CREATE TABLE `user_image_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

LOCK TABLES `user_image_category` WRITE;
/*!40000 ALTER TABLE `user_image_category` DISABLE KEYS */;

INSERT INTO `user_image_category` (`id`, `title`, `created_at`, `updated_at`)
VALUES
	(1,' GRAPHISME','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(2,'PACKAGING','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(3,'DESSINS ET CROQUIS','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(4,'COULEUR','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(5,'PERSPECTIVE','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(6,'ROUGH ET STORYBOARD','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(7,'PUBLICITÉ','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(8,'TYPOGRAPHIE','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(9,'IDENTITÉ VISUELLE','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(10,'ÉDITION','0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(11,'PROJETS DE DIPLÔMES','0000-00-00 00:00:00','0000-00-00 00:00:00');

/*!40000 ALTER TABLE `user_image_category` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
