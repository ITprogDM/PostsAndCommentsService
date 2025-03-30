CREATE TABLE IF NOT EXISTS Posts (
    id serial PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    name VARCHAR(100),
    content VARCHAR(2000),
    author VARCHAR(100),
    comments_allowed BOOLEAN
);

CREATE TABLE IF NOT EXISTS Comments (
    id serial PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    content VARCHAR(2000),
    author VARCHAR(100),
    post INT NOT NULL,
    FOREIGN KEY (post) REFERENCES Posts(id) ON DELETE CASCADE ON UPDATE CASCADE ,
    reply_to INT,
    FOREIGN KEY (reply_to) REFERENCES Comments(id) ON DELETE SET NULL ON UPDATE CASCADE
);