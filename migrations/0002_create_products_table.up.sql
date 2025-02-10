CREATE TABLE IF NOT EXISTS public.products (
    id bigserial PRIMARY KEY,
	product_code varchar(20) NOT NULL,
	product_name varchar(200) NOT NULL,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(36) NULL,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(36) NULL,
    deleted_at timestamp NULL,
    deleted_by varchar(36) NULL
);