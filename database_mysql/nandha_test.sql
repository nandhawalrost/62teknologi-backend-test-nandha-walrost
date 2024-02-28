-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 28, 2024 at 01:36 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 7.4.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `nandha_test`
--

-- --------------------------------------------------------

--
-- Table structure for table `businesses`
--

CREATE TABLE `businesses` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `location` longtext DEFAULT NULL,
  `latitude` bigint(20) UNSIGNED DEFAULT NULL,
  `longitude` bigint(20) UNSIGNED DEFAULT NULL,
  `term` longtext DEFAULT NULL,
  `radius` bigint(20) UNSIGNED DEFAULT NULL,
  `categories` longtext DEFAULT NULL,
  `locale` longtext DEFAULT NULL,
  `price` bigint(20) UNSIGNED DEFAULT NULL,
  `open_now` tinyint(1) DEFAULT NULL,
  `open_at` bigint(20) UNSIGNED DEFAULT NULL,
  `attributes` longtext DEFAULT NULL,
  `sort_by` longtext DEFAULT NULL,
  `device_platform` longtext DEFAULT NULL,
  `reservation_date` longtext DEFAULT NULL,
  `reservation_time` longtext DEFAULT NULL,
  `reservation_covers` bigint(20) UNSIGNED DEFAULT NULL,
  `matches_party_size_param` tinyint(1) DEFAULT NULL,
  `limit` bigint(20) UNSIGNED DEFAULT NULL,
  `offset` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `businesses`
--

INSERT INTO `businesses` (`id`, `created_at`, `updated_at`, `deleted_at`, `location`, `latitude`, `longitude`, `term`, `radius`, `categories`, `locale`, `price`, `open_now`, `open_at`, `attributes`, `sort_by`, `device_platform`, `reservation_date`, `reservation_time`, `reservation_covers`, `matches_party_size_param`, `limit`, `offset`) VALUES
(1, NULL, NULL, NULL, 'jakarta', 12345, 54321, 'testterm', 333, 'testcat', 'testlocale', 200000, 1, 5431123, 'afasvda', 'asvdav', 'asvadva', '2024-09-08', '20:04', 2, 1, 20, 1),
(8, '2024-02-27 09:01:14.636', '2024-02-27 09:34:20.425', NULL, 'tangerang', 12349, 54329, 'testterm2', 339, 'cattest', 'localtest', 200009, 0, 5431129, 'afasvd9', 'asvda9', 'asvadv9', '2024-09-01', '20:11', 1, 0, 29, 2),
(9, '2024-02-27 09:01:14.636', '2024-02-27 09:34:20.425', '2024-02-27 09:35:32.607', 'tangerang', 12349, 54329, 'testterm2', 339, 'cattest', 'localtest', 200009, 0, 5431129, 'afasvd9', 'asvda9', 'asvadv9', '2024-09-01', '20:11', 1, 0, 29, 2);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `businesses`
--
ALTER TABLE `businesses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_businesses_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `businesses`
--
ALTER TABLE `businesses`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
