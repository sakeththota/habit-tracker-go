CREATE TABLE IF NOT EXISTS progress (
    progress_id SERIAL PRIMARY KEY,
    habit_id INT NOT NULL,
    progress_date DATE NOT NULL,
    created_at DATE NOT NULL,
    UNIQUE (habit_id, progress_date),
    FOREIGN KEY (habit_id) REFERENCES habits (habit_id) ON DELETE CASCADE
);
