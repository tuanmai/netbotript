package markets

import (
	"fmt"
	"log"
	"os"

	"github.com/go-numb/go-ftx/auth"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/public/markets"
)

func GetSpots() []markets.Market {
	client := rest.New(auth.New(os.Getenv("FTX_KEY"), os.Getenv("FTX_SECRET")))
        p, err := client.Markets(&markets.RequestForMarkets{})
        products := *p
	if err != nil {
		log.Fatal("Cannot call ftx api", err)
	}
        var spots []markets.Market
        for _, product := range products {
          if product.Type == "spot" {
            spots = append(spots, product)
          }

        }
        fmt.Println(len(spots))
	return spots
}
