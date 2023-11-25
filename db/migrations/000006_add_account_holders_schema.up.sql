CREATE TABLE
    "account_holders" (
        "acc_id" bigint NOT NULL REFERENCES "accounts" ("id"),
        "user_id" bigint NOT NULL REFERENCES "users" ("id"),
        "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now ()),
        PRIMARY KEY ("acc_id", "user_id")
    );

CREATE INDEX ON "account_holders" ("acc_id","user_id");

CREATE INDEX ON "account_holders" ("acc_id");

CREATE INDEX ON "account_holders" ("user_id");