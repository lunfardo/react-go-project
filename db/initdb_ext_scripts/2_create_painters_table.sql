\c my_arts_database

CREATE TABLE painters (
    Id              SERIAL PRIMARY KEY,
    Fullname        VARCHAR(255),
    CityOfOrigin    VARCHAR(255),
    DiedRich        BOOLEAN  DEFAULT false,
    Image           bytea
);

