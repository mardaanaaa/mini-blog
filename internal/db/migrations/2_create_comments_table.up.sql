CREATE TABLE comments (
                          id SERIAL PRIMARY KEY,
                          post_id INT REFERENCES posts(id),
                          content TEXT NOT NULL,
                          user_id INT,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
