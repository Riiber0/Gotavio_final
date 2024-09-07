-- +goose Up
-- +goose StatementBegin
CREATE TABLE snippets (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	title VARCHAR(100) NOT NULL,
	author INTEGER NOT NULL,
	content TEXT NOT NULL,
	created DATETIME NOT NULL,
	expires DATETIME NOT NULL,
	FOREIGN KEY (author) REFERENCES users(id)
);

CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO snippets (title, author, content, created, expires) VALUES (
	'An old silent pond',
	1,
	'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
	DATETIME('now'),
	DATETIME(DATETIME('now'), '+1 month')
);

INSERT INTO snippets (title, author, content, created, expires) VALUES (
	'Over the wintry forest',
	1,
	'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
	DATETIME('now'),
	DATETIME(DATETIME('now'), '+1 month')
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE snippets;
-- +goose StatementEnd
