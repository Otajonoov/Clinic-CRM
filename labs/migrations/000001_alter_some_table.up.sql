CREATE TABLE IF NOT EXISTS "labs"(
    "id" uuid NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "price" double precision NOT NULL,
    "type" VARCHAR(255) NOT NULL,
    "sub_category_id" uuid NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "aparats"(
    "id" uuid NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "price" double precision NOT NULL,
    "type" VARCHAR(255) NOT NULL,
    "sub_category_id" uuid NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "lab_category"(
    "id" uuid NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "aparat_category"(
    "id" uuid NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "lab_sub_category"(
    "id" uuid PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "category_id" uuid NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "aparat_sub_category"(
    "id" uuid PRIMARY KEY NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "category_id" uuid NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "aparat_analysis"(
    "id" uuid NOT NULL,
    "client_id" INTEGER NOT NULL,
    "aparat_id" uuid NOT NULL,
    "analysis_url" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "lab_analysis"(
    "id" uuid NOT NULL,
    "client_id" INTEGER NOT NULL,
    "aparat_id" uuid NOT NULL,
    "analysis_url" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

