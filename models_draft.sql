-- Table for storing user information
CREATE TABLE "User" (
    user_id SERIAL PRIMARY KEY,          -- Unique identifier for each user
    username VARCHAR(50) UNIQUE,        -- Username must be unique, optional for OAuth users
    email VARCHAR(100) UNIQUE,          -- Email must be unique, optional for OAuth users
    password_hash VARCHAR(255),         -- For securely storing user passwords, null for OAuth users
    oauth_provider VARCHAR(50),         -- OAuth2 provider (e.g., Google, Facebook), null for email/password users
    oauth_provider_id VARCHAR(255),              -- Unique identifier from the OAuth2 provider, null for email/password users
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for when the user was created
); 

-- Table for storing habit information
CREATE TABLE "Habits" (
    habit_id SERIAL PRIMARY KEY,         -- Unique identifier for each habit
    user_id INT NOT NULL,                -- Foreign key referencing the User table
    title VARCHAR(100) NOT NULL,         -- Title of the habit
    description TEXT,                    -- Optional description of the habit
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the habit was created
    FOREIGN KEY (user_id) REFERENCES "User" (user_id) ON DELETE CASCADE -- Cascade delete when a user is removed
);

-- Table for tracking daily progress for each habit
CREATE TABLE "Progress" (
    progress_id SERIAL PRIMARY KEY,      -- Unique identifier for each progress entry
    habit_id INT NOT NULL,               -- Foreign key referencing the Habit table
    progress_date DATE NOT NULL,         -- Date of progress
    is_completed BOOLEAN NOT NULL,       -- Whether the habit was completed on this date
    UNIQUE (habit_id, progress_date),    -- Ensure one entry per habit per day
    FOREIGN KEY (habit_id) REFERENCES "Habit" (habit_id) ON DELETE CASCADE -- Cascade delete when a habit is removed
);

-- Index for optimizing sorting progress rows
CREATE INDEX idx_progress_date ON "Progress" (progress_date);

-- Query for acquiring latest progress row
SELECT *
FROM "Progress"
WHERE habit_id = <HABIT_ID>
ORDER BY progress_date DESC
LIMIT 1;

-- Query for creating a new progress row if the current date is not equal to progress_date of last row
INSERT INTO "Progress" (habit_id, progress_date, is_completed, created_at)
VALUES (<HABIT_ID>, CURRENT_DATE, FALSE, CURRENT_TIMESTAMP);
