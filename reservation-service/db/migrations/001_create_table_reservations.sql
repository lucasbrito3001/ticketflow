-- Migration: Create reservations and reservation_tickets tables
-- Description: Creates the baseline schema for the reservation service

CREATE TABLE reservations (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    event_id BIGINT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE reservation_tickets (
    reservation_id BIGINT NOT NULL,
    ticket_type VARCHAR(50) NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (reservation_id, ticket_type),
    FOREIGN KEY (reservation_id) REFERENCES reservations(id) ON DELETE CASCADE
);
