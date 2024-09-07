-- +goose Up
-- +goose StatementBegin
CREATE TABLE saved (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	snippetsID INTEGER NOT NULL,
	usersID INTEGER NOT NULL,
	FOREIGN KEY (snippetsID) REFERENCES snippets(id),
	FOREIGN KEY (usersID) REFERENCES users(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE saved;
-- +goose StatementEnd
