package markets

import (
	"testing"
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
}
