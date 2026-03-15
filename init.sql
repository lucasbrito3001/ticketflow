CREATE DATABASE IF NOT EXISTS ticketflow_inventory;
CREATE DATABASE IF NOT EXISTS ticketflow_reservation;

GRANT ALL PRIVILEGES ON ticketflow_inventory.* TO 'app'@'%';
GRANT ALL PRIVILEGES ON ticketflow_reservation.* TO 'app'@'%';

FLUSH PRIVILEGES;
