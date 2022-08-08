package stock

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
)

const (
	SymbolPosition = 0
	OpenPosition   = 3
)

type StockServiceModule struct {
	externalServiceUrl string
}

func (*StockServiceModule) ProvideStockServiceModule() StockService {
	return &stockService{}
}

type StockService interface {
	GetStockByCode(code string) ([]string, error)
}

type stockService struct {
}

func (s *stockService) GetStockByCode(code string) ([]string, error) {
	resp, err := http.Get(parseUrl(code))
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(resp.Body)
	csvLines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	for key, line := range csvLines {
		if key == 0 {
			continue
		}

		for _, col := range line {
			if col == "N/D" {
				return nil, &StockNotFoundError{}
			}
		}

		return line, nil
	}

	return nil, errors.New("symbol not found")
}

func parseUrl(code string) string {
	url := fmt.Sprintf("https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv", code)

	return url
}
