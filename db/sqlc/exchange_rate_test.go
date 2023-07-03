package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-firebond-assignment/common"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func createdRandomExchangeRate(t *testing.T) ExchangeRate {
	arg := InsertExchangeRateParams{
		CryptoCurrencyId: sql.NullString{
			Valid:  true,
			String: common.RadomCrypto(),
		},
		FiatCurrencyAcronym: sql.NullString{
			Valid:  true,
			String: common.RadomFiat(),
		},
		ExchanageRate: strconv.Itoa(int(common.RandomInt(1, 1000))),
	}

	ex, err := testQueries.InsertExchangeRate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ex)
	require.NotZero(t, ex.ID)
	return ex
}
func TestInsertExchangeRate(t *testing.T) {
	arg := InsertExchangeRateParams{
		CryptoCurrencyId: sql.NullString{
			Valid:  true,
			String: common.RadomCrypto(),
		},
		FiatCurrencyAcronym: sql.NullString{
			Valid:  true,
			String: common.RadomFiat(),
		},
		ExchanageRate: strconv.Itoa(int(common.RandomInt(1, 1000))),
	}

	ex, err := testQueries.InsertExchangeRate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ex)
	require.NotZero(t, ex.ID)

}

func TestListCurrentCryptoFiatRate(t *testing.T) {
	var exRate ExchangeRate
	for i := 0; i < 5; i++ {
		exRate = createdRandomExchangeRate(t)
	}
	arg := GetCurrentCryptoFiatRateParams{
		CryptoCurrencyId:    exRate.CryptoCurrencyId,
		FiatCurrencyAcronym: exRate.FiatCurrencyAcronym,
	}

	ex, err := testQueries.GetCurrentCryptoFiatRate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ex)
	require.NotZero(t, ex.ID)
	fmt.Printf(" stuff from query :: %v \n", ex)

}

func TestGetCurrentCryptoExchangeRate(t *testing.T) {
	var exRate ExchangeRate
	for i := 0; i < 5; i++ {
		exRate = createdRandomExchangeRate(t)
	}
	currencyId := exRate.CryptoCurrencyId
	rates, err := testQueries.GetCurrentCryptoExchangeRate(context.Background(), currencyId)
	require.NoError(t, err)
	require.NotEmpty(t, rates)
	require.NotZero(t, rates[0].ID)

}

func TestGetCurrentExchangeRates(t *testing.T) {
	for i := 0; i < 5; i++ {
		createdRandomExchangeRate(t)
	}
	rates, err := testQueries.GetCurrentExchangeRates(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, rates)
	for _, v := range rates {
		require.NotZero(t, v.ID)
	}
}

func TestGetExchangeRateHistory(t *testing.T) {
	var exRate ExchangeRate
	for i := 0; i < 5; i++ {
		exRate = createdRandomExchangeRate(t)
	}
	arg := GetExchangeRateHistoryParams{
		CryptoCurrencyId:    exRate.CryptoCurrencyId,
		FiatCurrencyAcronym: exRate.FiatCurrencyAcronym,
	}
	history, err := testQueries.GetExchangeRateHistory(context.Background(), arg)
	require.NoError(t, err)
	for _, v := range history {
		require.NotZero(t, v.ID)
		fmt.Printf("%v - %v :: %v  \n", v.CryptoCurrencyId.String, v.FiatCurrencyAcronym.String, v.ExchanageRate)
	}
}
