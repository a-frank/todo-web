CREATE TABLE IF NOT EXISTS Todo(
	id SERIAL PRIMARY KEY,
	todo TEXT,
	done boolean
);