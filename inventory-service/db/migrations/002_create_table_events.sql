-- Migration: 002_create_table_events
-- Description: Create events table to store event details and sale windows
-- Created: 2024-03-07

CREATE TABLE IF NOT EXISTS events (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    starts_at DATETIME NOT NULL,
    ends_at DATETIME NOT NULL,
    venue VARCHAR(255) NOT NULL,
    address_id BIGINT NOT NULL,
    open_sales_at DATETIME NOT NULL,
    close_sales_at DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT fk_events_address_id FOREIGN KEY (address_id)
        REFERENCES addresses(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    
    INDEX idx_starts_at (starts_at),
    INDEX idx_ends_at (ends_at),
    INDEX idx_open_sales_at (open_sales_at),
    INDEX idx_close_sales_at (close_sales_at),
    INDEX idx_venue (venue),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Stores event information including sales windows and venue';
