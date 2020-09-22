package domain_test

import (
	"testing"

	"github.com/resotto/goilerplate/cmd/app/domain"
)

func TestParameter(t *testing.T) {
	tests := []struct {
		name                       string
		funds, btc                 int
		expectedfunds, expectedbtc int
	}{
		{"more funds than btc", 1000, 0, 1000, 0},
		{"same amount", 100, 100, 100, 100},
		{"much more funds than btc", 100000, 20, 100000, 20},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			parameter := domain.Parameter{
				Funds: tt.funds,
				Btc:   tt.btc,
			}
			if parameter.Funds != tt.expectedfunds {
				t.Errorf("got %q, want %q", parameter.Funds, tt.expectedfunds)
			}
			if parameter.Btc != tt.expectedbtc {
				t.Errorf("got %q, want %q", parameter.Btc, tt.expectedbtc)
			}
		})
	}
}
