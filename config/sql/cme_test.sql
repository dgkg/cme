-- phpMyAdmin SQL Dump
-- version 3.5.7
-- http://www.phpmyadmin.net
--
-- Client: localhost
-- Généré le: Jeu 24 Avril 2014 à 15:52
-- Version du serveur: 5.5.29
-- Version de PHP: 5.4.10

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Base de données: `cme_test`
--

-- --------------------------------------------------------

--
-- Structure de la table `forum`
--

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
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=20 ;

--
-- Contenu de la table `forum`
--

INSERT INTO `forum` (`id`, `user_id`, `forum_category_id`, `title`, `text`, `keywords`, `is_solved`, `is_online`, `created_at`, `updated_at`) VALUES
(1, 0, 5, 'Titre A', 'Titre A : Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 1, 1, '2014-04-15 09:05:12', '2014-04-15 09:05:12'),
(2, 0, 4, 'Titre B', 'Doudoud didi Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 0, 0, '2014-04-15 09:06:09', '2014-04-15 09:06:09'),
(3, 0, 4, 'Titre C', 'objet C Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 1, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(4, 0, 4, 'Titre D', 'coucou de la forêt', '', 0, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(5, 0, 3, 'Titre E', 'objet E Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 1, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(6, 0, 2, 'Titre F', 'objet F Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 0, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(7, 0, 1, 'Titre G', 'objet G Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 0, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(8, 0, 1, 'Titre H', 'objet H Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 0, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(9, 0, 1, 'Comment faire pour rien faire ?', 'objet I Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 1, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(10, 0, 1, 'Titre J', 'objet J Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 1, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(11, 0, 1, 'Titre K', 'objet K Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 0, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(12, 0, 1, 'Titre L', 'objet L Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris pellentesque eget lorem ut cursus. Sed erat neque, sodales sit amet auctor id, consequat vel erat. Phasellus urna ligula, egestas ut lacus non, dapibus eleifend nulla. Vestibulum accumsan tellus nibh, non rhoncus nisi dictum quis. Suspendisse lobortis eros ac lacus eleifend commodo. Suspendisse lacinia lacus id eros aliquet lacinia. Vestibulum vel ligula pellentesque tortor eleifend cursus id eget massa. Duis eleifend pharetra nunc, vel dignissim eros. Vestibulum sed bibendum erat.', '', 0, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(13, 0, 0, 'qsdfqsdf', '<p>qsdfqsdf</p>\r\n', '', 0, 0, '2014-04-17 16:08:30', '2014-04-17 16:08:30'),
(14, 0, 2, 'Mon super titre', '<p>La question de la mort : &nbsp;Comment je fait pour r&eacute;cup&eacute;rer un post d&#39;un formulaire ?</p>\r\n', '', 0, 0, '2014-04-17 16:16:50', '2014-04-17 16:16:50'),
(15, 0, 2, 'titi', '<p>coucou comment sa vas ?</p>\r\n', '', 0, 1, '2014-04-17 17:27:19', '2014-04-17 17:27:19'),
(16, 0, 2, 'test coucou', '<p>ceci est un autre test&nbsp;</p>\r\n', '', 0, 1, '2014-04-22 09:03:38', '2014-04-22 09:03:38'),
(17, 0, 1, 'coucou hu hu', '<p>ha ha ho ho</p>\r\n', '', 0, 1, '2014-04-23 08:49:23', '2014-04-23 08:49:23'),
(18, 0, 3, 'qsfqsdfqs', '<p>sdfqsdfqsdf</p>\r\n', '', 0, 1, '2014-04-23 16:24:27', '2014-04-23 16:24:27'),
(19, 0, 3, 'coucouc', '<p>hihi&nbsp;</p>\r\n', '', 0, 1, '2014-04-24 09:07:26', '2014-04-24 09:07:26');

-- --------------------------------------------------------

--
-- Structure de la table `forum_category`
--

CREATE TABLE `forum_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=6 ;

--
-- Contenu de la table `forum_category`
--

INSERT INTO `forum_category` (`id`, `title`, `created_at`, `updated_at`) VALUES
(1, 'Packaging', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 'Web', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(3, 'Rough', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(4, 'Cours', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(5, 'Autre', '0000-00-00 00:00:00', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Structure de la table `forum_post`
--

CREATE TABLE `forum_post` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `forum_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `text` text NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=3 ;

--
-- Contenu de la table `forum_post`
--

INSERT INTO `forum_post` (`id`, `forum_id`, `user_id`, `text`, `is_online`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'voici ma réponse A', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 1, 2, 'voici ma réponse B', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00');

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Structure de la table `news_category`
--

CREATE TABLE `news_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Structure de la table `news_post`
--

CREATE TABLE `news_post` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `news_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `text` text NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Structure de la table `tag`
--

CREATE TABLE `tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(25) NOT NULL,
  `is_online` int(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Structure de la table `tutorial`
--

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
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=5 ;

--
-- Contenu de la table `tutorial`
--

INSERT INTO `tutorial` (`id`, `user_id`, `tutorial_category_id`, `title`, `text`, `keywords`, `is_online`, `created_at`, `updated_at`) VALUES
(1, 0, 1, 'Comment créer un effet néon sur du lettrage', 'Un petit tutoriel très simple qui explique comment générer un effet de lettrage néon, de style un peu rétro, avec les effets de Photoshop.', '', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 0, 2, 'Comment utiliser le Pathfinder dans Illustrator', 'Outil indispensable à tous les utilisateurs de Illustrator! Je vous explique ici comment maîtriser le Pathfinder.', '', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(3, 0, 3, 'Comment générer du CSS avec SASS/SCSS dans Sublime Text 2', 'SASS/SCSS s''incruste dans les bonnes habitudes des intégrateurs de ce monde à cause de ses fonctionnalités très utiles. Je vous montre ici comment générer du CSS compressé à partir de SASS avec Sublime Text 2.', '', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(4, 0, 2, 'coucou voici mon super tuto', '<p>Illustrator est un outils super simpa super tuto hihi</p>\r\n', '', 1, '2014-04-24 13:47:11', '2014-04-24 13:47:11');

-- --------------------------------------------------------

--
-- Structure de la table `tutorial_category`
--

CREATE TABLE `tutorial_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8 NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf32 AUTO_INCREMENT=4 ;

--
-- Contenu de la table `tutorial_category`
--

INSERT INTO `tutorial_category` (`id`, `title`, `created_at`, `updated_at`) VALUES
(1, 'Photoshop', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 'Illustator', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(3, 'Sublime Text', '0000-00-00 00:00:00', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Structure de la table `tutorial_post`
--

CREATE TABLE `tutorial_post` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tutorial_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `text` text NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Structure de la table `user`
--

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `text` text NOT NULL,
  `email` varchar(100) NOT NULL,
  `pass` varchar(200) NOT NULL,
  `keywords` varchar(200) NOT NULL,
  `is_online` int(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=7 ;

--
-- Contenu de la table `user`
--

INSERT INTO `user` (`id`, `first_name`, `last_name`, `text`, `email`, `pass`, `keywords`, `is_online`, `created_at`, `updated_at`) VALUES
(1, 'Damien', 'Weil', 'texte de Damien Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.', 'damien.weil@gmail.com', '', '', 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 'René', 'Grossi', 'texte de rené Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.', 'rene.grossi@gmail.com', '', '', 1, '2014-04-07 16:35:14', '2014-04-07 16:35:14'),
(3, 'Antoine', 'Chameau', 'texte d''antoine Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.', 'antoine.chameau@gmail.com', '', '', 1, '2014-04-07 16:35:38', '2014-04-07 16:35:38'),
(4, 'Henri', 'Lepic', 'texte d''Henri Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.', 'henrilepic@gmail.com', '', '', 1, '2014-04-07 16:36:41', '2014-04-07 16:36:41'),
(5, 'Amélie', 'Lepic', 'texte Amélie Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.', 'amelie.lepic@gmail.com', '', '', 0, '2014-04-07 17:04:15', '2014-04-07 17:04:15'),
(6, 'Hugues', 'Lepic', 'texte d''hugues Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc sit amet rhoncus augue, at mollis lacus. Phasellus eget consequat tellus. Donec sed turpis nunc. Nam id dolor porttitor, tristique tortor eget, iaculis sem. Donec nec tincidunt erat. Curabitur volutpat convallis porttitor. Nunc eleifend sem sem, quis ultricies augue porttitor ac. Nullam vitae ante nunc. Suspendisse nisi augue, condimentum eu nulla a, cursus condimentum mauris. Integer tristique pharetra tempor. Donec pharetra vestibulum risus sit amet dictum.', 'h.lepic@gmail.com', '', '', 1, '2014-04-07 17:08:35', '2014-04-07 17:08:35');

-- --------------------------------------------------------

--
-- Structure de la table `user_image`
--

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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=3 ;

--
-- Contenu de la table `user_image`
--

INSERT INTO `user_image` (`id`, `user_id`, `user_image_category_id`, `title`, `url`, `description`, `is_featured`, `is_online`, `created_at`, `updated_at`) VALUES
(1, 0, 0, 'Test titre...', '', '', 0, 0, '2014-04-07 17:04:15', '2014-04-07 17:04:15'),
(2, 6, 0, 'Test titre...', '', '', 0, 0, '2014-04-07 17:08:35', '2014-04-07 17:08:35');

-- --------------------------------------------------------

--
-- Structure de la table `user_image_category`
--

CREATE TABLE `user_image_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=12 ;

--
-- Contenu de la table `user_image_category`
--

INSERT INTO `user_image_category` (`id`, `title`, `created_at`, `updated_at`) VALUES
(1, ' GRAPHISME', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 'PACKAGING', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(3, 'DESSINS ET CROQUIS', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(4, 'COULEUR', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(5, 'PERSPECTIVE', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(6, 'ROUGH ET STORYBOARD', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(7, 'PUBLICITÉ', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(8, 'TYPOGRAPHIE', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(9, 'IDENTITÉ VISUELLE', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(10, 'ÉDITION', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(11, 'PROJETS DE DIPLÔMES', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
