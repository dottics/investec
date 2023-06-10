package investec

import (
	"fmt"
	"net/http"
	"net/url"
)

type Transaction struct {
	AccountID       string  `json:"accountId"`
	Type            string  `json:"type"`
	TransactionType string  `json:"transactionType"`
	Status          string  `json:"status"`
	Description     string  `json:"description"`
	CardNumber      string  `json:"cardNumber"`
	PostedOrder     int     `json:"postedOrder"`
	PostingDate     string  `json:"postingDate"`
	ValueDate       string  `json:"valueDate"`
	ActionDate      string  `json:"actionDate"`
	TransactionDate string  `json:"transactionDate"`
	Amount          float32 `json:"amount"`
	RunningBalance  float32 `json:"runningBalance"`
}

type TransactionQueryParameters struct {
	FromDate        string // 2006-01-02
	ToDate          string // 2006-01-02
	TransactionType string // "FeesAndInterest"
}

// EqualTransactions is a basic comparison that returns a boolean if two slices
// of Transaction are equal, in length and each indexed element is equal.
func EqualTransactions(a, b *[]Transaction) bool {
	if a == nil && b == nil {
		// both are nil
		return true
	} else if a == nil || b == nil {
		// either a or be are nil, not both
		return false
	}
	// a and b are both pointers
	if len(*a) != len(*b) {
		return false
	}
	for i := 0; i < len(*a); i++ {
		if (*a)[i] != (*b)[i] {
			return false
		}
	}
	return true
}

// GetTransactions fetches all the transactions from an account.
func (s *Service) GetTransactions(accountID string, options TransactionQueryParameters) (*[]Transaction, error) {
	// https://openapi.investec.com/za/pb/v1/accounts/{accountId}/transactions?fromDate={fromDate}&toDate={toDate}&transactionType={transactionType}
	// set the path
	s.URL.Path = "/za/pb/v1/accounts/" + accountID + "/transactions"
	// set the query parameters
	qs := url.Values{}
	if options.FromDate != "" {
		qs.Set("fromDate", options.FromDate)
	}
	if options.ToDate != "" {
		qs.Set("toDate", options.ToDate)
	}
	if options.TransactionType != "" {
		qs.Set("transactionType", options.TransactionType)
	}
	// add and encode the query parameters
	s.URL.RawQuery = qs.Encode()
	// make request
	req, err := http.NewRequest(http.MethodGet, s.URL.String(), nil)
	if err != nil {
		return nil, err
	}
	// set the request headers
	req.Header.Set("authorization", "Bearer"+s.Token)
	// do the request
	res, err := s.DoRequest(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Error %s", res.Status)
	}
	// response data structure
	type Data struct {
		Transactions *[]Transaction `json:"transactions"`
	}
	resp := struct {
		Data `json:"data"`
	}{}
	err = s.MarshalResponseJSON(res, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data.Transactions, nil
}
