package rest

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/fundingtrade"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
)

// KeepFundingRequest - data structure for constructing keep funding request payload
type KeepFundingRequest struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}

// FundingService manages the Funding endpoint.
type FundingService struct {
	requestFactory
	Synchronous
}

// Retreive all of the active fundign offers
// see https://docs.bitfinex.com/reference#rest-auth-funding-offers for more info
func (fs *FundingService) Offers(symbol string) (*bitfinex.FundingOfferSnapshot, error) {
	req, err := fs.requestFactory.NewAuthenticatedRequest(bitfinex.PermissionRead, path.Join("funding/offers", symbol))
	if err != nil {
		return nil, err
	}
	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}
	offers, err := bitfinex.NewFundingOfferSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

// Retreive all of the past in-active funding offers
// see https://docs.bitfinex.com/reference#rest-auth-funding-offers-hist for more info
func (fs *FundingService) OfferHistory(symbol string) (*bitfinex.FundingOfferSnapshot, error) {
	req, err := fs.requestFactory.NewAuthenticatedRequest(bitfinex.PermissionRead, path.Join("funding/offers", symbol, "hist"))
	if err != nil {
		return nil, err
	}
	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}
	offers, err := bitfinex.NewFundingOfferSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

// Retreive all of the active funding loans
// see https://docs.bitfinex.com/reference#rest-auth-funding-loans for more info
func (fs *FundingService) Loans(symbol string) (*bitfinex.FundingLoanSnapshot, error) {
	req, err := fs.requestFactory.NewAuthenticatedRequest(bitfinex.PermissionRead, path.Join("funding/loans", symbol))
	if err != nil {
		return nil, err
	}
	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}
	loans, err := bitfinex.NewFundingLoanSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

// Retreive all of the past in-active funding loans
// see https://docs.bitfinex.com/reference#rest-auth-funding-loans-hist for more info
func (fs *FundingService) LoansHistory(symbol string) (*bitfinex.FundingLoanSnapshot, error) {
	req, err := fs.requestFactory.NewAuthenticatedRequest(bitfinex.PermissionRead, path.Join("funding/loans", symbol, "hist"))
	if err != nil {
		return nil, err
	}
	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}
	loans, err := bitfinex.NewFundingLoanSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

// Retreive all of the active credits used in positions
// see https://docs.bitfinex.com/reference#rest-auth-funding-credits for more info
func (fs *FundingService) Credits(symbol string) (*bitfinex.FundingCreditSnapshot, error) {
	req, err := fs.requestFactory.NewAuthenticatedRequest(bitfinex.PermissionRead, path.Join("funding/credits", symbol))
	if err != nil {
		return nil, err
	}
	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}
	loans, err := bitfinex.NewFundingCreditSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

// Retreive all of the past in-active credits used in positions
// see https://docs.bitfinex.com/reference#rest-auth-funding-credits-hist for more info
func (fs *FundingService) CreditsHistory(symbol string) (*bitfinex.FundingCreditSnapshot, error) {
	req, err := fs.requestFactory.NewAuthenticatedRequest(bitfinex.PermissionRead, path.Join("funding/credits", symbol, "hist"))
	if err != nil {
		return nil, err
	}
	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}
	loans, err := bitfinex.NewFundingCreditSnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}
	return loans, nil
}

// Retreive all of the matched funding trades
// see https://docs.bitfinex.com/reference#rest-auth-funding-trades-hist for more info
func (fs *FundingService) Trades(symbol string) (*fundingtrade.Snapshot, error) {
	req, err := fs.requestFactory.NewAuthenticatedRequest(bitfinex.PermissionRead, path.Join("funding/trades", symbol, "hist"))
	if err != nil {
		return nil, err
	}

	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}

	fts, err := fundingtrade.SnapshotFromRaw(raw)
	if err != nil {
		return nil, err
	}

	return fts, nil
}

// Submits a request to create a new funding offer
// see https://docs.bitfinex.com/reference#submit-funding-offer for more info
func (fs *FundingService) SubmitOffer(fo *bitfinex.FundingOfferRequest) (*bitfinex.Notification, error) {
	bytes, err := fo.ToJSON()
	if err != nil {
		return nil, err
	}
	req, err := fs.requestFactory.NewAuthenticatedRequestWithBytes(bitfinex.PermissionWrite, path.Join("funding/offer/submit"), bytes)
	if err != nil {
		return nil, err
	}
	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}
	return bitfinex.NewNotificationFromRaw(raw)
}

// Submits a request to cancel the given offer
// see https://docs.bitfinex.com/reference#cancel-funding-offer for more info
func (fs *FundingService) CancelOffer(fc *bitfinex.FundingOfferCancelRequest) (*bitfinex.Notification, error) {
	bytes, err := fc.ToJSON()
	if err != nil {
		return nil, err
	}
	req, err := fs.requestFactory.NewAuthenticatedRequestWithBytes(bitfinex.PermissionWrite, "funding/offer/cancel", bytes)
	if err != nil {
		return nil, err
	}
	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}
	return bitfinex.NewNotificationFromRaw(raw)
}

// KeepFunding - toggle to keep funding taken. Specify loan for unused funding and credit for used funding.
// see https://docs.bitfinex.com/reference#rest-auth-keep-funding for more info
func (fs *FundingService) KeepFunding(args KeepFundingRequest) (*bitfinex.Notification, error) {
	if args.Type != "credit" && args.Type != "loan" {
		return nil, fmt.Errorf("Expected type: credit or loan, got: %s", args.Type)
	}

	bytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	req, err := fs.requestFactory.NewAuthenticatedRequestWithBytes(
		bitfinex.PermissionWrite,
		path.Join("funding", "keep"),
		bytes,
	)
	if err != nil {
		return nil, err
	}

	raw, err := fs.Request(req)
	if err != nil {
		return nil, err
	}

	return bitfinex.NewNotificationFromRaw(raw)
}
