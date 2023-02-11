package investec

import (
	"fmt"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_GetTransactions(t *testing.T) {
	tests := []struct {
		name         string
		exchange     *microtest.Exchange
		transactions []Transaction
		err          error
	}{
		{
			name: "401 Unauthorised",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 401,
				},
			},
			transactions: []Transaction{},
			err:          fmt.Errorf("HTTP Error 401 Unauthorized"),
		},
		{
			name: "200 Success",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body: `{
						"data": {
							"transactions": [
								{
									"accountId": "091298347856091298341011",
									"type": "DEBIT",
									"transactionType": "CardPurchases",
									"status": "POSTED",
									"description": "ABOUT CATS AND DOGS STELLENBOSCH ZA",
									"cardNumber": "402111xxxxxx1001",
									"postedOrder": 1234,
									"postingDate": "2023-02-07",
									"valueDate": "2023-02-15",
									"actionDate": "2023-02-11",
									"transactionDate": "2023-02-04",
									"amount": 3868.8,
									"runningBalance": 4615.83
								},
								{
									"accountId": "091298347856091298341011",
									"type": "CREDIT",
									"transactionType": "CardPurchases",
									"status": "POSTED",
									"description": "KWIKSPAR PARADYSKLOOF WESTERN CAPE ZA",
									"cardNumber": "402111xxxxxx1001",
									"postedOrder": 1237,
									"postingDate": "2023-02-06",
									"valueDate": "2023-02-15",
									"actionDate": "2023-02-11",
									"transactionDate": "2023-02-05",
									"amount": 68.78,
									"runningBalance": 935.01
								}
							]
						},
						"links": {
							"self": "https://openapi.investec.com/za/pb/v1/accounts/5912592322510190321754460/transactions?fromDate=2023-02-05&toDate=2023-02-11"
						},
						"meta": {
							"totalPages": 1
						}
					}`,
				},
			},
			transactions: []Transaction{
				{
					AccountID:       "091298347856091298341011",
					Type:            "DEBIT",
					TransactionType: "CardPurchases",
					Status:          "POSTED",
					Description:     "ABOUT CATS AND DOGS STELLENBOSCH ZA",
					CardNumber:      "402111xxxxxx1001",
					PostedOrder:     1234,
					PostingDate:     "2023-02-07",
					ValueDate:       "2023-02-15",
					ActionDate:      "2023-02-11",
					TransactionDate: "2023-02-04",
					Amount:          3868.8,
					RunningBalance:  4615.83,
				},
				{
					AccountID:       "091298347856091298341011",
					Type:            "CREDIT",
					TransactionType: "CardPurchases",
					Status:          "POSTED",
					Description:     "KWIKSPAR PARADYSKLOOF WESTERN CAPE ZA",
					CardNumber:      "402111xxxxxx1001",
					PostedOrder:     1237,
					PostingDate:     "2023-02-06",
					ValueDate:       "2023-02-15",
					ActionDate:      "2023-02-11",
					TransactionDate: "2023-02-05",
					Amount:          68.78,
					RunningBalance:  935.01,
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
			// call the method
			xt, err := s.GetTransactions("", "123456")
			if !equalError(tc.err, err) {
				t.Errorf("expected error %v got %v", tc.err, err)
			}
			if !EqualTransactions(tc.transactions, xt) {
				t.Errorf("expected transactions %v got %v", tc.transactions, xt)
			}
		})
	}
}

func TestEqualTransactions(t *testing.T) {
	a := []Transaction{
		{
			AccountID:       "091298347856091298341011",
			Type:            "CREDIT",
			TransactionType: "CardPurchases",
			Status:          "POSTED",
			Description:     "KWIKSPAR PARADYSKLOOF WESTERN CAPE ZA",
			CardNumber:      "402111xxxxxx1001",
			PostedOrder:     1237,
			PostingDate:     "2023-02-06",
			ValueDate:       "2023-02-15",
			ActionDate:      "2023-02-11",
			TransactionDate: "2023-02-05",
			Amount:          68.78,
			RunningBalance:  935.01,
		},
	}
	b := []Transaction{
		{
			AccountID:       "091298347856091298341011",
			Type:            "CREDIT",
			TransactionType: "CardPurchases",
			Status:          "POSTED",
			Description:     "KWIKSPAR PARADYSKLOOF WESTERN CAPE ZA",
			CardNumber:      "402111xxxxxx1001",
			PostedOrder:     1237,
			PostingDate:     "2023-02-06",
			ValueDate:       "2023-02-15",
			ActionDate:      "2023-02-11",
			TransactionDate: "2023-02-05",
			Amount:          68.78,
			RunningBalance:  935.01,
		},
	}

	if !EqualTransactions(a, b) {
		t.Errorf("expected equal transactions %v got %v", a, b)
	}
}
