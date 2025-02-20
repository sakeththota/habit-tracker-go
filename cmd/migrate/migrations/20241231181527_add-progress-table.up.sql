CREATE TABLE IF NOT EXISTS progress (
    id SERIAL PRIMARY KEY,
    habit_id INT NOT NULL,
    date DATE NOT NULL,
    created_at DATE NOT NULL,
    UNIQUE (habit_id, date),
    FOREIGN KEY (habit_id) REFERENCES habits (habit_id) ON DELETE CASCADE
);
