CREATE TABLE IF NOT EXISTS providers(
    id serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS tokens(
    id serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS operations(
    id serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS histories(
    id serial PRIMARY KEY,

    provider_id serial NOT NULL,
    token_id serial NOT NULL,

    operation_id serial NOT NULL,
    apr real NOT NULL,

    created_at timestamp default (timezone('utc', now())) NOT NULL,

    FOREIGN KEY(provider_id) REFERENCES providers(id),
    FOREIGN KEY(token_id) REFERENCES tokens(id),
    FOREIGN KEY(operation_id) REFERENCES operations(id)
);

CREATE TABLE IF NOT EXISTS liquidity_pool_histories(
    id serial PRIMARY KEY,

    provider_id serial NOT NULL,
    token_a_id serial NOT NULL,
    token_b_id serial NOT NULL,

    apr real NOT NULL,

    created_at timestamp default (timezone('utc', now())) NOT NULL,

    FOREIGN KEY(provider_id) REFERENCES providers(id),
    FOREIGN KEY(token_a_id) REFERENCES tokens(id),
    FOREIGN KEY(token_b_id) REFERENCES tokens(id)
);
