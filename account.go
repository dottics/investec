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

// EqualAccounts is a basic comparison function to ensure that two slices of
// accounts are equal. That is, they are equal in length and the order in which
// accounts are indexed are the same.
func EqualAccounts(a, b []Account) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// GetAccounts fetches all the accounts that are accessible based on an access token.
func (s *Service) GetAccounts(token string) ([]Account, error) {
	s.URL.Path = "/za/pb/v1/accounts"
	req, err := http.NewRequest(http.MethodGet, s.URL.String(), nil)
	if err != nil {
		return []Account{}, err
	}
	res, err := s.DoRequest(req)
	if err != nil {
		return []Account{}, err
	}

	if res.StatusCode != 200 {
		return []Account{}, fmt.Errorf("HTTP Error: %d %s", res.StatusCode, res.Status)
	}

	// define the data structure expected from Investec.
	type Data struct {
		Accounts []Account `json:"accounts"`
	}
	resp := struct {
		Data `json:"data"`
	}{}
	err = s.MarshalResponseJSON(res, &resp)
	if err != nil {
		return []Account{}, err
	}
	return resp.Data.Accounts, err
}
