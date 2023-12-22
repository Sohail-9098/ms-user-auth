CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(50),
    password VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (user_name, password) VALUES ('Chek', '65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5');