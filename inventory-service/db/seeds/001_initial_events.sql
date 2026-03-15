-- Seed: Initial Events and Ticket Catalog
-- Description: Development seed data with 3 events and their ticket types
-- Environment: local, dev

INSERT INTO addresses (street, number, city, state, zip_code) VALUES
('Avenida Atlântica', 500, 'Rio de Janeiro', 'RJ', '22010-000'),
('Rua Augusta', 2000, 'São Paulo', 'SP', '01305-100'),
('Rua Oscar Freire', 200, 'São Paulo', 'SP', '01426-100');

INSERT INTO events (name, starts_at, ends_at, venue, address_id, open_sales_at, close_sales_at) VALUES
(
    'Coldplay - Music of the Spheres Tour 2024',
    '2024-09-15 20:00:00',
    '2024-09-15 23:30:00',
    'Estádio do Morumbi',
    2,
    '2024-01-01 00:00:00',
    '2024-09-14 23:59:59'
),
(
    'Festival Back Beats 2024',
    '2024-10-20 14:00:00',
    '2024-10-20 23:59:59',
    'Parque Villa Lobos',
    2,
    '2024-02-01 00:00:00',
    '2024-10-19 23:59:59'
),
(
    'Praia Venue Summer Concert',
    '2024-12-25 18:00:00',
    '2024-12-25 22:00:00',
    'Copacabana Beach',
    1,
    '2024-03-01 00:00:00',
    '2024-12-24 23:59:59'
);

INSERT INTO ticket_catalog_items (event_id, type, price, total_quantity, sold_quantity) VALUES
-- Coldplay - Music of the Spheres
(1, 'GENERAL', 50000, 2000, 150),      -- R$ 500.00
(1, 'VIP', 100000, 200, 20),           -- R$ 1000.00

-- Festival Back Beats
(2, 'GENERAL', 35000, 3000, 280),      -- R$ 350.00
(2, 'VIP', 75000, 300, 50),       -- R$ 750.00

-- Praia Venue Summer Concert
(3, 'GENERAL', 40000, 1500, 200),      -- R$ 400.00
(3, 'VIP', 90000, 150, 25);            -- R$ 900.00
