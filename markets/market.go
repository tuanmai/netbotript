package markets

import (
	"fmt"
	"log"
	"os"

	"github.com/go-numb/go-ftx/auth"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/public/markets"
)

func GetSpots() *markets.ResponseForMarkets {
	client := rest.New(auth.New(os.Getenv("FTX_KEY"), os.Getenv("FTX_SECRET")))
        products, err := client.Markets(&markets.RequestForMarkets{})
	if err != nil {
		log.Fatal("Cannot call ftx api", err)
	}
        fmt.Println(products)
	return products
}
