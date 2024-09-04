-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	hashed_password CHAR(60) NOT NULL,
	created DATETIME NOT NULL
);

INSERT INTO users (name, email, hashed_password, created) VALUES (
	'user1',
	'user1@email.com',
	'$2a$12$SWjnakVw5tiX0A450H73z.XgWfC6hfK4uchbMtKp81tEVwYRd8Mcu',
	DATETIME('now')
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
