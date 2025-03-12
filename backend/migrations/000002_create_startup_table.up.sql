CREATE TABLE IF NOT EXISTS startup (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    author_id INT REFERENCES author(id) ON DELETE CASCADE,
    views INT DEFAULT 0,
    description TEXT NOT NULL,
    category VARCHAR(20) NOT NULL CHECK (LENGTH(category) BETWEEN 1 AND 20),
    image_url TEXT NOT NULL,
    pitch_markdown TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);
