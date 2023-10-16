CREATE TABLE IF NOT EXISTS "doctors"(
    "id" UUID NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "gender" VARCHAR(255) NOT NULL,
    "work_time" VARCHAR(255) NOT NULL,
    "price" double precision NOT NULL,
    "cpecialety" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "room_number" VARCHAR(255) NOT NULL,        
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);  

CREATE TABLE IF NOT EXISTS "doctor_reports"(
    "id" UUID NOT NULL,
    "client_id" UUID NOT NULL,
    "doctor_id" UUID NOT NULL,
    "text" TEXT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "sqlad"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "count" INTEGER NOT NULL,
    "price" double precision NOT NULL,
    "low_stock" INTEGER NOT NULL,
    "expiration_date" DATE NOT NULL,
    "provider" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);