-- 1. Create the posts table
CREATE TABLE posts (
    post_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    description VARCHAR(500),
    image_link TEXT
);

-- 2. Create the comments table
CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,
    post_id INT NOT NULL REFERENCES posts(post_id) ON DELETE CASCADE,
    content VARCHAR(500) NOT NULL,
    user_id INT NOT NULL
);

-- 3. Create the likes table
CREATE TABLE likes (
    post_id INT NOT NULL REFERENCES posts(post_id) ON DELETE CASCADE,
    user_id INT NOT NULL,
    PRIMARY KEY (post_id, user_id) -- Ensures a user can't like the same post multiple times
);

-- Index on user_id in posts to speed up queries filtering by user
CREATE INDEX idx_posts_user_id ON posts(user_id);

-- Index on post_id in comments to optimize joins and lookups
CREATE INDEX idx_comments_post_id ON comments(post_id);

-- Index on post_id in likes for faster counting of likes per post
CREATE INDEX idx_likes_post_id ON likes(post_id);

-- Index on user_id in likes for efficiently finding posts liked by a user
CREATE INDEX idx_likes_user_id ON likes(user_id);
