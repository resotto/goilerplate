package usecase_test

import (
	"testing"

	"github.com/resotto/goilerplate/internal/app/application/usecase"
	"github.com/resotto/goilerplate/internal/app/domain/valueobject"
	"github.com/resotto/goilerplate/testdata"
)

func TestTicker(t *testing.T) {
	tests := []struct {
		name              string
		pair              valueobject.Pair
		expectedsell      string
		expectedbuy       string
		expectedhigh      string
		expectedlow       string
		expectedlast      string
		expectedvol       string
		expectedtimestamp string
	}{
		{"btcjpy", valueobject.BtcJpy, "1000", "1000", "2000", "500", "1200", "20", "1600769562"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mexchange := testdata.MExchange{} // using Mock
			result := usecase.Ticker(mexchange, tt.pair)
			if result.Sell != tt.expectedsell {
				t.Errorf("got %q, want %q", result.Sell, tt.expectedsell)
			}
			if result.Buy != tt.expectedbuy {
				t.Errorf("got %q, want %q", result.Buy, tt.expectedbuy)
			}
			if result.High != tt.expectedhigh {
				t.Errorf("got %q, want %q", result.High, tt.expectedhigh)
			}
			if result.Low != tt.expectedlow {
				t.Errorf("got %q, want %q", result.Low, tt.expectedlow)
			}
			if result.Last != tt.expectedlast {
				t.Errorf("got %q, want %q", result.Last, tt.expectedlast)
			}
			if result.Vol != tt.expectedvol {
				t.Errorf("got %q, want %q", result.Vol, tt.expectedvol)
			}
			if result.Timestamp != tt.expectedtimestamp {
				t.Errorf("got %q, want %q", result.Timestamp, tt.expectedtimestamp)
			}
		})
	}
}
