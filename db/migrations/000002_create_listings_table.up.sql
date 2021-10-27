CREATE TABLE IF NOT EXISTS listings
(
    id          SERIAL PRIMARY KEY,
    mls_number  INTEGER      NOT NULL,
    address     VARCHAR(255) NOT NULL,
    address_2   VARCHAR(255),
    city        VARCHAR(255) NOT NULL,
    postal_code VARCHAR(6)   NOT NULL,
    state       VARCHAR(255) NOT NULL,
    country     VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);