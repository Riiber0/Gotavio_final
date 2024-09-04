-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions (
	token TEXT PRIMARY KEY,
	data BLOB NOT NULL,
	expiry REAL NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions(expiry);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP sessions
-- +goose StatementEnd
