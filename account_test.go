package investec

import (
	"fmt"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_GetAccounts(t *testing.T) {
	tests := []struct {
		name     string
		exchange *microtest.Exchange
		accounts *[]Account
		err      error
	}{
		{
			name: "401 Forbidden",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 401,
					Body:   `{}`,
				},
			},
			accounts: nil,
			err:      nil,
		},
		{
			name: "200 Success",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					// this data also serves as a visual example of data to be returned.
					Body: `{
						"data":{
							"accounts":[
								{
        							"accountId": "091298347856091298341011",
									"accountNumber": "10011238899",
        							"accountName": "Mr J Bond",
        							"referenceName": "Mr J Bond",
        							"productName": "Private Bank Account",
        							"kycCompliant": true,
        							"profileId": "10190921254450"
								}, 
								{
        							"accountId": "091298347856091298341012",
        							"accountNumber": "110065431234",
        							"accountName": "Mr J Bond",
        							"referenceName": "Mr J Bond",
        							"productName": "PrimeSaver",
        							"kycCompliant": true,
        							"profileId": "10190921254450"
      							}
							]
						},
						"links": {
							"self": "https://openapi.investec.com/za/pb/v1/accounts"
						},
						"meta": {
							"totalPages": 1
						}
					}`,
				},
			},
			accounts: &[]Account{
				{
					AccountID:     "091298347856091298341011",
					AccountNumber: "10011238899",
					AccountName:   "Mr J Bond",
					ReferenceName: "Mr J Bond",
					ProductName:   "Private Bank Account",
					KYCCompliant:  true,
					ProfileID:     "10190921254450",
				},
				{
					AccountID:     "091298347856091298341012",
					AccountNumber: "110065431234",
					AccountName:   "Mr J Bond",
					ReferenceName: "Mr J Bond",
					ProductName:   "PrimeSaver",
					KYCCompliant:  true,
					ProfileID:     "10190921254450",
				},
			},
			err: nil,
		},
	}

	s := NewService("investec")
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

	for i, tc := range tests {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			xa, _ := s.GetAccounts("")
			if !EqualAccounts(tc.accounts, xa) {
				t.Errorf("expected accounts %v got %v", tc.accounts, xa)
			}
		})
	}
}
