package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-firebond-assignment/common"
	"go-firebond-assignment/config"
	db "go-firebond-assignment/db/sqlc"
	"log"
	"net/http"
	"strings"
	"sync"
)

type workerService struct {
	config config.Config
	store  db.Store
}

// Worker implements IWorkerService.
func (ws *workerService) Worker() {
	apiUrl := ws.config.CoinGeckoBaseUrl + "/simple/price"

	params := map[string]string{
		"ids":           strings.Join(common.Cryptocurrencies, ","),
		"vs_currencies": strings.Join(common.FiatCurrencies, ","),
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Printf("error making new request :: %v \n", err)
		return
	}
	query := req.URL.Query()
	for k, v := range params {
		query.Add(k, v)
	}

	req.URL.RawQuery = query.Encode()

	response, err := client.Do(req)
	if err != nil {
		log.Printf("error sending request :: %v \n", err)
		return
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("error processing request ... \n statusCode - %v \n", response.StatusCode)
		return
	}

	var result map[string]map[string]float64
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Printf("error parsing response :: %v \n", err)
		return
	}

	var wg sync.WaitGroup

	for _, crypto := range common.Cryptocurrencies {
		exchangeRates := result[crypto]
		for _, fiat := range common.FiatCurrencies {
			wg.Add(1)

			go func(crypto, fiat string) {
				defer wg.Done()
				arg := db.InsertExchangeRateParams{
					CryptoCurrencyId: sql.NullString{
						String: crypto,
						Valid:  true,
					},
					FiatCurrencyAcronym: sql.NullString{
						String: fmt.Sprintf("%.2f", exchangeRates[fiat]),
						Valid:  true,
					},
				}
				_, err = ws.store.InsertExchangeRate(context.TODO(), arg)
				if err != nil {
					log.Printf("error inserting %v - %v \n\n %v \n\n", fiat, crypto, err)
				}
			}(crypto, fiat)
		}
	}

	wg.Wait()
}

type IWorkerService interface {
	Worker()
}

func NewWorkerService(config config.Config, store db.Store) IWorkerService {
	return &workerService{
		store:  store,
		config: config,
	}
}
