CREATE TABLE IF NOT EXISTS public.users (
    id bigserial PRIMARY KEY,
	name varchar(200) NOT NULL,
    email varchar(100) NOT NULL,
	phone_number varchar(20) NULL,
	password text NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(36) NULL,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(36) NULL,
    deleted_at timestamp NULL,
    deleted_by varchar(36) NULL
);