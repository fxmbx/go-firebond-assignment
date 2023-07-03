-- Drop indexes
DROP INDEX IF EXISTS "fiat_currency_currency_acronym_idx";
DROP INDEX IF EXISTS "crypto_currency_currencyId_idx";
DROP INDEX IF EXISTS "exchange_rate_created_at_idx";

-- Drop tables
DROP TABLE IF EXISTS "exchange_rate";