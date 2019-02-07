\c my_arts_database

CREATE TABLE paints (
    Id                  SERIAL PRIMARY KEY,
    PainterId           INTEGER REFERENCES painters(Id),
    Name                VARCHAR(255),
    CurrentLocation     VARCHAR(255)
);

