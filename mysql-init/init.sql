-- 🔥 IMPROVED: Complete database schema initialization

CREATE DATABASE IF NOT EXISTS userdb;
USE userdb;
CREATE TABLE IF NOT EXISTS users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    alamat TEXT,
    preferensi VARCHAR(255),
    role VARCHAR(50) DEFAULT 'customer',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ===========================
-- ORDER DATABASE
-- ===========================
CREATE DATABASE IF NOT EXISTS orderdb;
USE orderdb;
CREATE TABLE IF NOT EXISTS orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    resi VARCHAR(100) UNIQUE NOT NULL,
    nama_barang VARCHAR(255),
    berat INT,
    dimensi VARCHAR(100),
    jenis VARCHAR(100),
    alamat_pengirim TEXT,
    alamat_penerima TEXT,
    nama_penerima VARCHAR(255),
    no_telp_penerima VARCHAR(50),
    status VARCHAR(50) DEFAULT 'created',
    eta VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ===========================
-- PAYMENT DATABASE
-- ===========================
CREATE DATABASE IF NOT EXISTS paymentdb;

-- ===========================
-- TRACKING DATABASE
-- ===========================
CREATE DATABASE IF NOT EXISTS trackingdb;
USE trackingdb;
CREATE TABLE IF NOT EXISTS tracking_events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    resi VARCHAR(100) NOT NULL,
    lokasi VARCHAR(255),
    event VARCHAR(255),
    timestamp DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_resi (resi)
);

-- ===========================
-- GUDANG DATABASE
-- ===========================
CREATE DATABASE IF NOT EXISTS gudangdb;
USE gudangdb;
CREATE TABLE IF NOT EXISTS packages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    resi VARCHAR(100) UNIQUE NOT NULL,
    nama_barang VARCHAR(255),
    berat INT,
    dimensi VARCHAR(100),
    jenis VARCHAR(100),
    alamat_pengirim TEXT,
    alamat_penerima TEXT,
    nama_penerima VARCHAR(255),
    no_telp_penerima VARCHAR(50),
    warehouse_zone VARCHAR(50),
    status VARCHAR(50) DEFAULT 'created',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    sorted_at TIMESTAMP NULL,
    INDEX idx_resi (resi)
);

CREATE TABLE IF NOT EXISTS outbox_events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    event_type VARCHAR(100),
    payload LONGTEXT,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    sent_at TIMESTAMP NULL,
    INDEX idx_status (status)
);

-- ===========================
-- COURIER DATABASE
-- ===========================
CREATE DATABASE IF NOT EXISTS courierdb;
USE courierdb;
CREATE TABLE IF NOT EXISTS deliveries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    resi VARCHAR(100) NOT NULL,
    courier_id INT,
    assigned_zone VARCHAR(50),
    nama_penerima VARCHAR(255),
    no_telp_penerima VARCHAR(50),
    alamat_penerima TEXT,
    berat INT,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    delivered_at TIMESTAMP NULL,
    INDEX idx_resi (resi),
    INDEX idx_courier (courier_id),
    INDEX idx_status (status)
);

-- ===========================
-- REPORT DATABASE
-- ===========================
CREATE DATABASE IF NOT EXISTS reportdb;
USE reportdb;
CREATE TABLE IF NOT EXISTS packages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    resi VARCHAR(100),
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_resi (resi),
    INDEX idx_status (status)
);

CREATE TABLE IF NOT EXISTS deliveries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    courier_id INT,
    resi VARCHAR(100),
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_courier (courier_id),
    INDEX idx_resi (resi)
);