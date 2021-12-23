package markets

import (
	"errors"
	"netbotript/markets/mocks"
	"testing"

	"github.com/go-numb/go-ftx/rest/public/markets"
)

func TestSmoke(t *testing.T) {
  _, err := GetSpots(CreateFTXClient())
  
  if err != nil {
    t.Fatalf(`Cannot call ftx api: %q`, err)
  }
}

func TestSpotsNotEmpty(t *testing.T) {
  spots, _ := GetSpots(CreateFTXClient())
  if len(spots) == 0 {
    t.Errorf("Spots is empty")
  }
}

func TestErrorWhenApiFail(t *testing.T) {
  mockClient := &mocks.FXTClient{}
  marketsResponse := new(markets.ResponseForMarkets)
  mockClient.On("Markets", &markets.RequestForMarkets{}).Return(marketsResponse, errors.New("Api failed"))
  _, err := GetSpots(mockClient)


  if err == nil {
    t.Error("Expect return error when api failed")
  }
}
