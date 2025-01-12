CREATE TABLE kpis (
    id UUID NOT NULL DEFAULT (uuid_generate_v4()),
    name TEXT NOT NULL,
    value FLOAT NOT NULL,
    target FLOAT NOT NULL,
    day DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
