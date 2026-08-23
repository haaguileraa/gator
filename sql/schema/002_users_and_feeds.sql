-- +goose Up
CREATE TABLE feeds(
	user_id UUID NOT NULL,
	name TEXT NOT NULL,
	url TEXT NOT NULL UNIQUE,
	FOREIGN KEY(user_id)
	REFERENCES users(id)
	ON DELETE CASCADE
); 

-- +goose Down
DROP TABLE feeds;
