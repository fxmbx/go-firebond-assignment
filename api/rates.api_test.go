package api

import (
	"database/sql"
	"go-firebond-assignment/common"
	mockdb "go-firebond-assignment/db/mock"
	db "go-firebond-assignment/db/sqlc"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetCurrentCryptoToFiatRate(t *testing.T) {
	type Query struct {
		CryptoCurrencyId    sql.NullString `json:"crypto_currency_Id"`
		FiatCurrencyAcronym sql.NullString `json:"fiat_currency_acronym"`
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetCurrentCryptoFiatRateParams{
					CryptoCurrencyId: sql.NullString{
						String: common.RadomCrypto(),
						Valid:  true,
					},
					FiatCurrencyAcronym: sql.NullString{
						String: common.RadomFiat(),
						Valid:  true,
					},
				}
				store.EXPECT().GetCurrentCryptoFiatRate(gomock.Any(), gomock.Eq(arg)).AnyTimes().Return(db.GetCurrentCryptoFiatRateRow{}, nil)
			},
		},
		{
			name: "InternalServerErro",
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetCurrentCryptoFiatRateParams{
					CryptoCurrencyId: sql.NullString{
						String: common.RadomCrypto(),
						Valid:  false,
					},
					FiatCurrencyAcronym: sql.NullString{
						String: common.RadomFiat(),
						Valid:  false,
					},
				}
				store.EXPECT().GetCurrentCryptoFiatRate(gomock.Any(), gomock.Eq(arg)).AnyTimes().Return(db.GetCurrentCryptoFiatRateRow{}, sql.ErrConnDone)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
		})
	}
}
