CREATE TABLE comic(
    id INT PRIMARY KEY     NOT NULL,
    title CHAR(50) NOT NULL,
    date DATE NOT NULL,
    image_url CHAR(100) NOT NULL,
    description CHAR(100) NOT NULL DEFAULT "",
);
