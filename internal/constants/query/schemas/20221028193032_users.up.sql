CREATE TABLE users (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "first_name" varchar      NOT NULL,
    "middle_name" varchar      NOT NULL,
    "last_name" varchar      NOT NULL,
    "email" varchar NOT NULL,
    "status" status NOT NULL DEFAULT 'ACTIVE',
    "created_at" timestamptz NOT NULL default now(),
    "updated_at" timestamptz NOT NULL default now(),
    "deleted_at" timestamptz
);

create unique index users_email_deleted_at_key on users(email,deleted_at)where deleted_at IS NULL;