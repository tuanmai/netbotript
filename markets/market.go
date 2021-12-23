package markets

import (
	"os"

	"github.com/go-numb/go-ftx/auth"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/public/markets"
)

type FXTClient interface {
  Markets(* markets.RequestForMarkets) (*markets.ResponseForMarkets, error)
}

func GetSpots(client FXTClient) ([]markets.Market, error) {
	p, err := client.Markets(&markets.RequestForMarkets{})
	products := *p
	if err != nil {
		return make([]markets.Market, 0), err
	}
	var spots []markets.Market
	for _, product := range products {
		if product.Type == "spot" {
			spots = append(spots, product)
		}

	}
	return spots, nil
}

func CreateFTXClient() FXTClient {
  client := rest.New(auth.New(os.Getenv("FTX_KEY"), os.Getenv("FTX_SECRET")))
  return client
}
