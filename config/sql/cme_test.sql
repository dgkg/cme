-- phpMyAdmin SQL Dump
-- version 4.1.12
-- http://www.phpmyadmin.net
--
-- Client :  localhost:8889
-- Généré le :  Mar 06 Mai 2014 à 09:43
-- Version du serveur :  5.5.34
-- Version de PHP :  5.5.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Base de données :  `cme_test`
--

-- --------------------------------------------------------

--
-- Structure de la table `news`
--

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
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=5 ;

--
-- Contenu de la table `news`
--

INSERT INTO `news` (`id`, `id_user`, `news_category_id`, `title`, `text`, `keywords`, `is_online`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'Lancement du site CME', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse odio magna, lobortis ac dolor vitae, feugiat rhoncus justo. Maecenas quis augue leo. Praesent bibendum eget nunc non interdum. Pellentesque ipsum nisi, convallis ut tempor nec, lobortis sit amet sapien. Etiam auctor blandit quam, a commodo nibh imperdiet ac. Suspendisse imperdiet odio non nibh rhoncus pharetra. Pellentesque pulvinar, elit eu condimentum tempus, dui metus adipiscing elit, vel sollicitudin risus sem ac ligula. In vitae sagittis tortor. Donec nec turpis viverra, lobortis elit quis, malesuada velit.', '', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 2, 2, 'Coucou actu 2', 'Autre actu ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse odio magna, lobortis ac dolor vitae, feugiat rhoncus justo. Maecenas quis augue leo. Praesent bibendum eget nunc non interdum. Pellentesque ipsum nisi, convallis ut tempor nec, lobortis sit amet sapien. Etiam auctor blandit quam, a commodo nibh imperdiet ac. Suspendisse imperdiet odio non nibh rhoncus pharetra. Pellentesque pulvinar, elit eu condimentum tempus, dui metus adipiscing elit, vel sollicitudin risus sem ac ligula. In vitae sagittis tortor. Donec nec turpis viverra, lobortis elit quis, malesuada velit.', '', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(3, 3, 2, 'Autre titre de news 3', 'Quisque iaculis turpis eget semper porta. Proin vestibulum tincidunt turpis at consequat. Maecenas dignissim, nunc a tempor venenatis, risus mauris bibendum nibh, et consequat eros mauris sit amet massa. Pellentesque sodales tortor a felis ornare tincidunt. Phasellus interdum ante mi, nec rutrum sem mollis eget. Nullam ac malesuada metus. Donec erat mauris, cursus id elementum vestibulum, mollis sed mi. Morbi hendrerit nisl et accumsan tristique. Duis fermentum est in viverra mollis. Phasellus pellentesque gravida scelerisque. Ut feugiat porttitor suscipit.', '', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(4, 4, 1, 'Nouveau Titre 4', 'Morbi vehicula magna vitae arcu suscipit congue. Proin porttitor tellus vitae feugiat cursus. Cras leo eros, varius ut molestie id, dapibus eu mi. Nulla blandit mi id odio feugiat, eu consequat magna iaculis. Mauris orci velit, accumsan nec sapien nec, elementum posuere urna. Cras lorem massa, mollis vel euismod id, rhoncus quis nulla. Interdum et malesuada fames ac ante ipsum primis in faucibus. Praesent tempor ultricies iaculis. Fusce venenatis ut metus ut faucibus. Praesent tempor facilisis rutrum.', '', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00');
