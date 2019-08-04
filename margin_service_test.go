package binance

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginTestSuite struct {
	baseTestSuite
}

func TestMarginAccountService(t *testing.T) {
	suite.Run(t, new(marginTestSuite))
}

func (s *marginTestSuite) TestTransfer() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	transferType := MarginTransferTypeToMargin
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
			"type":   transferType,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginTransferService().Asset(asset).
		Amount(amount).Type(transferType).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) assertTransactionResponseEqual(a, e *TransactionResponse) {
	s.r().Equal(a.TranID, e.TranID, "TranID")
}

func (s *marginTestSuite) TestLoan() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginLoanService().Asset(asset).
		Amount(amount).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) TestRepay() {
	data := []byte(`{
		"tranId": 100000001
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BTC"
	amount := "1.000"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setFormParams(params{
			"asset":  asset,
			"amount": amount,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewMarginRepayService().Asset(asset).
		Amount(amount).Do(newContext())
	s.r().NoError(err)
	e := &TransactionResponse{
		TranID: 100000001,
	}
	s.assertTransactionResponseEqual(e, res)
}

func (s *marginTestSuite) TestListMarginLoans() {
	data := []byte(`{
		"rows": [
		  {
			"asset": "BNB",
			"principal": "0.84624403",
			"timestamp": 1555056425000,
			"status": "CONFIRMED"
		  }
		],
		"total": 1
	  }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BNB"
	txID := int64(1)
	startTime := int64(1555056425000)
	endTime := int64(1555056425001)
	current := int64(1)
	size := int64(10)
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"asset":     asset,
			"txId":      txID,
			"startTime": startTime,
			"endTime":   endTime,
			"current":   current,
			"size":      size,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListMarginLoansService().Asset(asset).
		TxID(txID).StartTime(startTime).EndTime(endTime).
		Current(current).Size(size).Do(newContext())
	s.r().NoError(err)
	e := &MarginLoanResponse{
		Rows: []MarginLoan{
			{
				Asset:     asset,
				Principal: "0.84624403",
				Timestamp: 1555056425000,
				Status:    MarginLoanStatusTypeConfirmed,
			},
		},
		Total: 1,
	}
	s.assertMarginLoanResponseEqual(e, res)
}

func (s *marginTestSuite) assertMarginLoanResponseEqual(e, a *MarginLoanResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	r.Len(a.Rows, len(e.Rows), "Rows")
	for i := 0; i < len(e.Rows); i++ {
		s.assertMarginLoanEqual(&e.Rows[i], &a.Rows[i])
	}
}

func (s *marginTestSuite) assertMarginLoanEqual(e, a *MarginLoan) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Principal, a.Principal, "Principal")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.Status, a.Status, "Status")
}

func (s *marginTestSuite) TestListMarginRepays() {
	data := []byte(`{
		"rows": [
			{
				"amount": "14.00000000",
				"asset": "BNB",
				"interest": "0.01866667",
				"principal": "13.98133333",
				"status": "CONFIRMED",
				"timestamp": 1563438204000,
				"txId": 2970933056
			}
		],
		"total": 1
   	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := "BNB"
	txID := int64(1)
	startTime := int64(1563438204000)
	endTime := int64(1563438204001)
	current := int64(1)
	size := int64(10)
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"asset":     asset,
			"txId":      txID,
			"startTime": startTime,
			"endTime":   endTime,
			"current":   current,
			"size":      size,
		})
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewListMarginRepaysService().Asset(asset).
		TxID(txID).StartTime(startTime).EndTime(endTime).
		Current(current).Size(size).Do(newContext())
	s.r().NoError(err)
	e := &MarginRepayResponse{
		Rows: []MarginRepay{
			{
				Asset:     asset,
				Amount:    "14.00000000",
				Interest:  "0.01866667",
				Principal: "13.98133333",
				Timestamp: 1563438204000,
				Status:    MarginRepayStatusTypeConfirmed,
				TxID:      2970933056,
			},
		},
		Total: 1,
	}
	s.assertMarginRepayResponseEqual(e, res)
}

func (s *marginTestSuite) assertMarginRepayResponseEqual(e, a *MarginRepayResponse) {
	r := s.r()
	r.Equal(e.Total, a.Total, "Total")
	r.Len(a.Rows, len(e.Rows), "Rows")
	for i := 0; i < len(e.Rows); i++ {
		s.assertMarginRepayEqual(&e.Rows[i], &a.Rows[i])
	}
}

func (s *marginTestSuite) assertMarginRepayEqual(e, a *MarginRepay) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Interest, a.Interest, "Interest")
	r.Equal(e.Principal, a.Principal, "Principal")
	r.Equal(e.Timestamp, a.Timestamp, "Timestamp")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TxID, a.TxID, "TxID")
}

func (s *marginTestSuite) TestGetMarginAccount() {
	data := []byte(`{
		"borrowEnabled": true,
		"marginLevel": "11.64405625",
		"totalAssetOfBtc": "6.82728457",
		"totalLiabilityOfBtc": "0.58633215",
		"totalNetAssetOfBtc": "6.24095242",
		"tradeEnabled": true,
		"transferEnabled": true,
		"userAssets": [
			{
				"asset": "BTC",
				"borrowed": "0.00000000",
				"free": "0.00499500",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "0.00499500"
			},
			{
				"asset": "BNB",
				"borrowed": "201.66666672",
				"free": "2346.50000000",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "2144.83333328"
			},
			{
				"asset": "ETH",
				"borrowed": "0.00000000",
				"free": "0.00000000",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "0.00000000"
			},
			{
				"asset": "USDT",
				"borrowed": "0.00000000",
				"free": "0.00000000",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "0.00000000"
			}
		]
  	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewGetMarginAccountService().Do(newContext())
	s.r().NoError(err)
	e := &MarginAccount{
		BorrowEnabled:       true,
		MarginLevel:         "11.64405625",
		TotalAssetOfBTC:     "6.82728457",
		TotalLiabilityOfBTC: "0.58633215",
		TotalNetAssetOfBTC:  "6.24095242",
		TradeEnabled:        true,
		TransferEnabled:     true,
		UserAssets: []UserAsset{
			{
				Asset:    "BTC",
				Borrowed: "0.00000000",
				Free:     "0.00499500",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "0.00499500",
			},
			{
				Asset:    "BNB",
				Borrowed: "201.66666672",
				Free:     "2346.50000000",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "2144.83333328",
			},
			{
				Asset:    "ETH",
				Borrowed: "0.00000000",
				Free:     "0.00000000",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "0.00000000",
			},
			{
				Asset:    "USDT",
				Borrowed: "0.00000000",
				Free:     "0.00000000",
				Interest: "0.00000000",
				Locked:   "0.00000000",
				NetAsset: "0.00000000",
			},
		},
	}
	s.assertMarginAccountEqual(e, res)
}

func (s *marginTestSuite) assertMarginAccountEqual(e, a *MarginAccount) {
	r := s.r()
	r.Equal(e.BorrowEnabled, a.BorrowEnabled, "BorrowEnabled")
	r.Equal(e.MarginLevel, a.MarginLevel, "MarginLevel")
	r.Equal(e.TotalAssetOfBTC, a.TotalAssetOfBTC, "TotalAssetOfBTC")
	r.Equal(e.TotalLiabilityOfBTC, a.TotalLiabilityOfBTC, "TotalLiabilityOfBTC")
	r.Equal(e.TotalNetAssetOfBTC, a.TotalNetAssetOfBTC, "TotalNetAssetOfBTC")
	r.Equal(e.TradeEnabled, a.TradeEnabled, "TradeEnabled")
	r.Equal(e.TransferEnabled, a.TransferEnabled, "TransferEnabled")
	r.Len(a.UserAssets, len(e.UserAssets), "UserAssets")
	for i := 0; i < len(a.UserAssets); i++ {
		s.assertUserAssetEqual(e.UserAssets[i], a.UserAssets[i])
	}
}

func (s *marginTestSuite) assertUserAssetEqual(e, a UserAsset) {
	r := s.r()
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Borrowed, a.Borrowed, "Borrowed")
	r.Equal(e.Free, a.Free, "Free")
	r.Equal(e.Interest, a.Interest, "Interest")
	r.Equal(e.Locked, a.Locked, "Locked")
	r.Equal(e.NetAsset, a.NetAsset, "NetAsset")
}
