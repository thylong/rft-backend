CREATE TABLE okrs (
    id UUID NOT NULL DEFAULT (uuid_generate_v4()),
    name TEXT NOT NULL,
    number INTEGER NOT NULL,
    year INTEGER NOT NULL,
    description TEXT NOT NULL
);


CREATE TABLE okr_krs (
    id UUID NOT NULL DEFAULT (uuid_generate_v4()),
    okr_id UUID NOT NULL,
    name TEXT NOT NULL,
    number INTEGER NOT NULL,
    description TEXT NOT NULL,
    sponsor TEXT NOT NULL,
    kpis TEXT NOT NULL,
    scope TEXT NOT NULL,
    initiatives TEXT NOT NULL
);
