CREATE TABLE `borrowers` ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, `first_name` TEXT NOT NULL, `middle_name` TEXT, `last_name` TEXT NOT NULL, `email` TEXT NOT NULL, `guid` TEXT NOT NULL, `score` INTEGER NOT NULL, `created_at` TEXT NOT NULL, `updated_at` TEXT NOT NULL )


CREATE TABLE `lender_loans` ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, `lender_id` INTEGER NOT NULL, `loan_id` INTEGER NOT NULL, `created_at` TEXT NOT NULL, `updated_at` TEXT NOT NULL )

CREATE TABLE `lenders` ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, `email` TEXT NOT NULL UNIQUE, `first_name` TEXT NOT NULL, `middle_name` TEXT, `last_name` TEXT NOT NULL, `guid` TEXT NOT NULL, `created_at` TEXT NOT NULL, `updated_at` TEXT NOT NULL )

CREATE TABLE `loans` ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, `name` TEXT NOT NULL, `description` TEXT, `borrower_id` INTEGER NOT NULL, `guid` TEXT NOT NULL, `amount` NUMERIC NOT NULL, `starts_at` TEXT NOT NULL, `ends_at` TEXT NOT NULL, `created_at` TEXT NOT NULL, `updated_at` TEXT NOT NULL )

CREATE TABLE `payments` ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, `loan_id` INTEGER NOT NULL, `guid` TEXT NOT NULL, `iso_currency_code` INTEGER NOT NULL, `amount` NUMERIC NOT NULL, `status` INTEGER NOT NULL, `payment_at` TEXT NOT NULL, `borrower_id` INTEGER NOT NULL, `lender_id` INTEGER NOT NULL, `created_at` TEXT NOT NULL, `updated_at` TEXT NOT NULL )



CREATE INDEX `idx_loans_borrower_id` ON `loans` ( `borrower_id` )
CREATE INDEX `idx_payments_borrower_id` ON `payments` ( `borrower_id` )
CREATE INDEX `idx_payments_lender_id` ON `payments` ( `lender_id` )
CREATE INDEX `idx_payments_loan_id` ON `payments` ( `loan_id` )
CREATE UNIQUE INDEX `uq_idx_lender_loan_lender_id_loan_id` ON `lender_loans` ( `loan_id`, `lender_id` )