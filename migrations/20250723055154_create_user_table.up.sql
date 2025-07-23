CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email TEXT UNIQUE NOT NULL,
        name VARCHAR UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT now (),
            updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT now ()
    );