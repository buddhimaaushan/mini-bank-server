CREATE TABLE
    "verify_phones" (
        "id" bigserial PRIMARY KEY,
        "username" varchar NOT NULL,
        "phone" varchar NOT NULL,
        "secret_code" varchar NOT NULL,
        "is_used" bool NOT NULL DEFAULT false,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
    );

ALTER TABLE "verify_phones"
ADD
    FOREIGN KEY ("username") REFERENCES "users" ("username");