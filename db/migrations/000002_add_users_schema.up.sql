CREATE TABLE
    "users" (
        "id" bigserial PRIMARY KEY,
        "first_name" varchar NOT NULL,
        "last_name" varchar NOT NULL,
        "username" varchar UNIQUE NOT NULL,
        "nic" varchar UNIQUE NOT NULL,
        "hashed_password" varchar NOT NULL,
        "password_changed_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
        "email" varchar UNIQUE NOT NULL,
        "is_email_verified" BOOLEAN NOT NULL DEFAULT false,
        "email_changed_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
        "phone" varchar UNIQUE NOT NULL,
        "is_phone_verified" BOOLEAN NOT NULL DEFAULT false,
        "phone_changed_at" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
        "acc_status" status NOT NULL DEFAULT 'inactive',
        "customer_rank" rank NOT NULL DEFAULT 'bronze',
        "is_an_employee" BOOLEAN NOT NULL DEFAULT false,
        "is_a_customer" BOOLEAN NOT NULL DEFAULT false,
        "role" varchar,
        "department" varchar,
        "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now ())
    );

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("nic");

CREATE INDEX ON "users" ("phone");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("role");

CREATE INDEX ON "users" ("department");

CREATE INDEX ON "users" ("first_name");

CREATE INDEX ON "users" ("last_name");

CREATE INDEX ON "users" ("acc_status");