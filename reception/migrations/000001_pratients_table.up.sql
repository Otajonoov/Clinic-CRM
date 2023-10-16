CREATE TABLE IF NOT EXISTS "patients"(
    "id" UUID NOT NULL,
    "client_id" INTEGER NOT NULL,
    "doctor_id" UUID NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "patronymic" VARCHAR(255) NOT NULL,
    "date_of_birth" DATE NOT NULL,
    "main_phone_number" VARCHAR(255) NOT NULL,
    "other_phone_number" VARCHAR(255) NOT NULL,
    "advertising_channel" VARCHAR(255) NOT NULL,
    "respublic" VARCHAR(255) NOT NULL,
    "region" VARCHAR(255) NOT NULL,
    "district" VARCHAR(255) NOT NULL,
    "passport_info" VARCHAR(255) NOT NULL,
    "discount" VARCHAR(255) NOT NULL,
    "condition" VARCHAR(255) NOT NULL,
    "gender" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "analysiss" (
    "id" UUID NOT NULL PRIMARY KEY,
    "client_phone_number" VARCHAR(255) NOT NULL,
    "analysiss_name" VARCHAR(255) NOT NULL,
    "analysiss_url" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);   

CREATE TABLE IF NOT EXISTS "lab_analysis"(
    "id" UUID NOT NULL,
    "client_id" INTEGER NOT NULL,
    "lab_id" UUID NOT NULL,
    "analysis_url" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "queues"(
    "id" UUID NOT NULL UNIQUE,
    "client_id" INTEGER NOT NULL,
    "queue_number" INTEGER,
    "service_id" UUID NOT NULL,
    "service_type" VARCHAR(255) NOT NULL,
    "turn_passed" BOOLEAN NOT NULL DEFAULT FALSE,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "cashbox"(
    "id" UUID NOT NULL,
    "client_id" INTEGER NOT NULL,
    "summa" INTEGER DEFAULT 0,
    "is_payed" BOOLEAN DEFAULT FALSE,
    "cash_count" INTEGER,
    "payment_type" VARCHAR(255) DEFAULT '',
    "doctors_ids" VARCHAR{},
    "labs_ids" VARCHAR{},
    "aparats_ids" VARCHAR{},
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "payment_history"(
    "id" UUID NOT NULL,
    "client_id" INTEGER NOT NULL,
    "summa" INTEGER DEFAULT 0,
    "payment_type" VARCHAR(255) DEFAULT '',
    "cashbox_id" UUID NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);






