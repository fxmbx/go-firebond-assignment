-- name: InsertExchangeRate :one
INSERT INTO "exchange_rate" (
  "crypto_currency_Id",
  fiat_currency_acronym,
  exchanage_rate
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCurrentCryptoExchangeRate :many
SELECT
	subquery.id,
	subquery.created_at,
	subquery. "crypto_currency_Id",
	subquery.fiat_currency_acronym,
	subquery.exchanage_rate,
	row_num
FROM (
	SELECT
		id,
		created_at,
		"crypto_currency_Id",
		fiat_currency_acronym,
		exchanage_rate,
		ROW_NUMBER() OVER (PARTITION BY "crypto_currency_Id",
			fiat_currency_acronym ORDER BY created_at DESC) AS row_num
	FROM
		"exchange_rate") subquery
WHERE
	row_num = 1
	AND "crypto_currency_Id" = $1;

-- name: GetCurrentCryptoFiatRate :one
SELECT
		subquery.id, subquery.created_at, subquery."crypto_currency_Id", subquery.fiat_currency_acronym, subquery.exchanage_rate, row_num
	FROM (
		SELECT
			id,
			created_at,
			"crypto_currency_Id",
			fiat_currency_acronym,
			exchanage_rate,
			ROW_NUMBER() OVER (PARTITION BY "crypto_currency_Id",
				fiat_currency_acronym ORDER BY created_at DESC) AS row_num
		FROM
			"exchange_rate") subquery
WHERE
	subquery.row_num = 1 AND "crypto_currency_Id" = $1 AND fiat_currency_acronym = $2;


-- name: GetCurrentExchangeRates :many
SELECT
		subquery.id, subquery.created_at, subquery."crypto_currency_Id", subquery.fiat_currency_acronym, subquery.exchanage_rate, row_num
	FROM (
		SELECT
			id,
			created_at,
			"crypto_currency_Id",
			fiat_currency_acronym,
			exchanage_rate,
			ROW_NUMBER() OVER (PARTITION BY "crypto_currency_Id",
				fiat_currency_acronym ORDER BY created_at DESC) AS row_num
		FROM
			"exchange_rate") subquery
WHERE
	subquery.row_num = 1;

-- name: GetExchangeRateHistory :many
SELECT
		subquery.id, subquery.created_at, subquery."crypto_currency_Id", subquery.fiat_currency_acronym, subquery.exchanage_rate, row_num
	FROM (
		SELECT
			id,
			created_at,
			"crypto_currency_Id",
			fiat_currency_acronym,
			exchanage_rate,
			ROW_NUMBER() OVER (PARTITION BY "crypto_currency_Id",
				fiat_currency_acronym ORDER BY created_at DESC) AS row_num
		FROM
			"exchange_rate"
      WHERE created_at >= NOW() - INTERVAL '24 hours'
      ) subquery
WHERE
  "crypto_currency_Id" = $1 AND fiat_currency_acronym = $2;
