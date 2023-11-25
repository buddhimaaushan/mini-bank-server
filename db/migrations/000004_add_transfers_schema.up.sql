CREATE TABLE
    "transfers" (
        "id" bigserial PRIMARY KEY,
        "from_account_id" bigint NOT NULL,
        "to_account_id" bigint NOT NULL,
        "amount" bigint NOT NULL,
        "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
    );

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

CREATE INDEX ON "transfers" ("amount");

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

ALTER TABLE "transfers"
ADD
    FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
ADD
    FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");