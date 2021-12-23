package markets

import (
	"testing"
)

func TestCanGetMarkets(t *testing.T) {
  _, err := GetSpots()
  
  if err != nil {
    t.Fatalf(`Cannot call ftx api: %q`, err)
  }
}

func TestSpotsNotEmpty(t *testing.T) {
  spots, _ := GetSpots()
  if len(spots) > 0 {
    t.Errorf("Spots is empty")
  }
}
