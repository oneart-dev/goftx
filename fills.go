package goftx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/grishinsana/goftx/models"
)

const (
	apiFills           = "/fills"
	apiFundingPayments = "/funding_payments"
)

type Fills struct {
	client *Client
}

func (f *Fills) GetFills(params *models.GetFillsParams) ([]*models.Fill, error) {
	queryParams, err := PrepareQueryParams(params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	request, err := f.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", f.client.apiURL, apiFills),
		Params: queryParams,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := f.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.Fill
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}

func (f *Fills) GetFundingPayments(params *models.GetFundingPaymentParams) ([]*models.FundingPayment, error) {
	queryParams, err := PrepareQueryParams(params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	request, err := f.client.prepareRequest(Request{
		Auth:   true,
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s%s", f.client.apiURL, apiFundingPayments),
		Params: queryParams,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := f.client.do(request)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []*models.FundingPayment
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}