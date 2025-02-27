CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    hashed_password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 🔹 Case-insensitive unique index on username
CREATE UNIQUE INDEX idx_users_username_ci ON users(LOWER(username));

-- 🔹 Check constraint to enforce username length
ALTER TABLE users ADD CONSTRAINT chk_username_length CHECK (LENGTH(username) >= 3);

-- 🔹 Ensure hashed_password is not empty
ALTER TABLE users ADD CONSTRAINT chk_password_not_empty CHECK (LENGTH(hashed_password) > 0);

