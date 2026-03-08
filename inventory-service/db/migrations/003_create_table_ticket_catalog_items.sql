-- Migration: 003_create_table_ticket_catalog_items
-- Description: Create ticket catalog items table to manage ticket types and inventory
-- Created: 2024-03-07

CREATE TABLE IF NOT EXISTS ticket_catalog_items (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    event_id BIGINT NOT NULL,
    type VARCHAR(50) NOT NULL COMMENT 'Ticket type: INTEIRA, MEIA, VIP, etc',
    price BIGINT NOT NULL COMMENT 'Price in cents',
    total_quantity INT NOT NULL,
    sold_quantity INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_ticket_catalog_items_event_id FOREIGN KEY (event_id)
        REFERENCES events(id) ON DELETE CASCADE ON UPDATE CASCADE,
    
    UNIQUE KEY uk_event_id_type (event_id, type),
    INDEX idx_event_id (event_id),
    INDEX idx_type (type),
    INDEX idx_sold_quantity (sold_quantity)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Stores ticket catalog items with inventory control';
