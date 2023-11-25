CREATE TABLE
    "entries" (
        "id" bigserial PRIMARY KEY,
        "account_id" bigint NOT NULL,
        "amount" bigint NOT NULL,
        "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
    );

CREATE INDEX ON "entries" ("account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

ALTER TABLE "entries"
ADD
    FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");