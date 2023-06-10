package investec

import (
	"fmt"
	"net/http"
)

type Account struct {
	AccountID     string `json:"accountId"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	ReferenceName string `json:"referenceName"`
	ProductName   string `json:"productName"`
	KYCCompliant  bool   `json:"kycCompliant"`
	ProfileID     string `json:"profileId"`
}

type AccountBalance struct {
	AccountID        string  `json:"accountId"`
	CurrentBalance   float64 `json:"currentBalance"`
	AvailableBalance float64 `json:"availableBalance"`
	Currency         string  `json:"currency"`
}

// EqualAccounts is a basic comparison function to ensure that two slices of
// accounts are equal. That is, they are equal in length and the order in which
// accounts are indexed are the same.
func EqualAccounts(a, b *[]Account) bool {
	// both are nil
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		// either a or b is nil
		return false
	}
	// a and b are not nil
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

// GetAccounts fetches all the accounts that are accessible based on an access token.
func (s *Service) GetAccounts() (*[]Account, error) {
	// set request path
	s.URL.Path = "/za/pb/v1/accounts"
	req, err := http.NewRequest(http.MethodGet, s.URL.String(), nil)
	if err != nil {
		return nil, err
	}
	// add request headers
	req.Header.Set("authorization", "Bearer"+s.Token)
	req.Header.Set("accept", "application/json")

	res, err := s.DoRequest(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Error %d %s", res.StatusCode, res.Status)
	}

	// define the data structure expected from Investec.
	type Data struct {
		Accounts *[]Account `json:"accounts"`
	}
	resp := struct {
		Data `json:"data"`
	}{}
	err = s.MarshalResponseJSON(res, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data.Accounts, err
}

// GetAccountBalance fetches the balance of a specific account.
func (s *Service) GetAccountBalance(accountID string) (AccountBalance, error) {
	// set request path
	s.URL.Path = "/za/pb/v1/accounts/" + accountID + "/balance"
	req, err := http.NewRequest(http.MethodGet, s.URL.String(), nil)
	if err != nil {
		return AccountBalance{}, err
	}
	// add request headers
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", s.Token))
	req.Header.Set("accept", "application/json")

	// make the exchange
	res, err := s.DoRequest(req)
	if err != nil {
		return AccountBalance{}, err
	}

	if res.StatusCode != 200 {
		return AccountBalance{}, fmt.Errorf("HTTP Error %d %s", res.StatusCode, res.Status)
	}

	// define the data structure expected from Investec.
	resp := struct {
		AccountBalance AccountBalance `json:"data"`
	}{}
	err = s.MarshalResponseJSON(res, &resp)
	if err != nil {
		return AccountBalance{}, err
	}

	return resp.AccountBalance, nil
}
