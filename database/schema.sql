-- users table structure
CREATE TABLE Users(
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`userName` VARCHAR(255) UNIQUE NOT NULL,
	`password` VARCHAR(255) NOT NULL,
	`isAdmin` BOOL DEFAULT false
);
-- notifications table structure
CREATE TABLE Notifications(
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`userID` INTEGER NOT NULL,
	`message` VARCHAR(510),
	`isRead` Bool DEFAULT false,
	`created_at` DATE default CURRENT_DATE
);

-- activities table structure
CREATE TABLE Activities (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`name` VARCHAR(255) NOT NULL,
	`userID` INTEGER NOT NULL,
	`duration` INTEGER NOT NULL,
	`intensity` VARCHAR(255) NOT NULL,
	`date` DATE default CURRENT_DATE
);
