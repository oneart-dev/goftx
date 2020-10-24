package goftx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/grishinsana/goftx/models"
)

const (
	apiGetMarkets          = "/markets"
	apiGetOrderBook        = "/markets/%s/orderbook"
	apiGetTrades           = "/markets/%s/trades"
	apiGetHistoricalPrices = "/markets/%s/candles"
)

type Markets struct {
	client *Client
}

func (m *Markets) GetMarkets() ([]*models.Market, error) {
	request, err := m.client.prepareRequest(Request{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, apiGetMarkets),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := m.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.Market
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (m *Markets) GetMarketByName(name string) (*models.Market, error) {
	request, err := m.client.prepareRequest(Request{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s/%s", apiUrl, apiGetMarkets, name),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := m.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result models.Market
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &result, nil
}

func (m *Markets) GetOrderBook(marketName string, depth *int) (*models.OrderBook, error) {
	params := map[string]string{}
	if depth != nil {
		params["depth"] = fmt.Sprintf("%d", *depth)
	}

	path := fmt.Sprintf(apiGetOrderBook, marketName)

	request, err := m.client.prepareRequest(Request{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, path),
		Params: params,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := m.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result models.OrderBook
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &result, nil
}

type GetTradesParams struct {
	Limit     *int `json:"limit"`
	StartTime *int `json:"start_time"`
	EndTime   *int `json:"end_time"`
}

func (m *Markets) GetTrades(marketName string, params *GetTradesParams) ([]*models.Trade, error) {
	queryParams, err := PrepareQueryParams(params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	path := fmt.Sprintf(apiGetTrades, marketName)
	request, err := m.client.prepareRequest(Request{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, path),
		Params: queryParams,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := m.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.Trade
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

type GetHistoricalPricesParams struct {
	Resolution models.Resolution `json:"resolution"`
	Limit      *int              `json:"limit"`
	StartTime  *int              `json:"start_time"`
	EndTime    *int              `json:"end_time"`
}

func (m *Markets) GetHistoricalPrices(marketName string, params *GetHistoricalPricesParams) ([]*models.HistoricalPrice, error) {
	queryParams, err := PrepareQueryParams(params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	path := fmt.Sprintf(apiGetHistoricalPrices, marketName)
	request, err := m.client.prepareRequest(Request{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", apiUrl, path),
		Params: queryParams,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := m.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.HistoricalPrice
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}