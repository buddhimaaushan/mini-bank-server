CREATE TABLE
    "accounts" (
        "id" bigserial PRIMARY KEY,
        "type" varchar NOT NULL,
        "balance" bigint NOT NULL,
        "acc_status" status NOT NULL DEFAULT 'inactive',
        "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
    );

CREATE INDEX ON "accounts" ("type");

CREATE INDEX ON "accounts" ("balance");