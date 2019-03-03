package coinbase

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/preichenberger/go-coinbasepro/v2"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

// mockNextResponses allows mocking the following count responses with content
// negative value = forever
func mockNextResponses(c *http.Client, newRes *http.Response, count int) {
	currentRT := c.Transport
	c.Transport = roundTripperFunc(func(req *http.Request) (res *http.Response, err error) {
		if count == 0 {
			c.Transport = currentRT
			return currentRT.RoundTrip(req)
		}
		res = newRes
		if count > 0 {
			count--
		}
		return
	})
}

func TestGetTicker(t *testing.T) {
	type response struct {
		httpCode int
		body     string
		mock     bool
	}
	tests := []struct {
		name         string
		product      string
		mockResponse *http.Response
		wantTicker   *coinbasepro.Ticker
		wantRes      response
		wantErr      bool
	}{
		{
			name:    "unit_valid",
			product: "ETH-BTC",
			mockResponse: &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString(`
			{"trade_id":95416,"price":"0.98996000","size":"0.01000000","time":"2019-03-02T07:53:27.602Z",
			"bid":"0.98995","ask":"0.98996","volume":"0.03000000"}`))},
			wantTicker: &coinbasepro.Ticker{
				TradeID: 95416,
				Price:   "0.98996000",
				Size:    "0.01000000",
				Bid:     "0.98995",
				Ask:     "0.98996",
				Volume:  "0.03000000",
			},
		},
		{
			name:    "unit_invalid",
			product: "doesn't exist",
			mockResponse: &http.Response{StatusCode: 404, Body: ioutil.NopCloser(bytes.NewBufferString(
				`{"message":"NotFound"}`))},
			wantErr: true,
		},
		{
			name:    "integration_valid",
			product: "ETH-BTC",
		},
		{
			name:    "integration_invalid",
			product: "doesn't exist",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{Client: *coinbasepro.NewClient()}

			if testing.Short() && tt.mockResponse == nil {
				t.Skip("skipping integration tests in short mode")
			} else {
				c.RetryCount = 3
			}

			if tt.mockResponse != nil {
				// mock requests when running locally
				mockNextResponses(c.HTTPClient, tt.mockResponse, 1)
			}

			gotTicker, _, err := c.GetTicker(tt.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetTicker() error = '%v', wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && gotTicker == nil {
				t.Error("ticker shouldn't be nil if there was no error")
				return
			}

			// ticker data changes frequently when doing live reqs, so only checking for err or missing ticker
			if tt.mockResponse == nil || gotTicker == nil {
				return
			}

			tt.wantTicker.Time = gotTicker.Time // todo compare times too
			if !reflect.DeepEqual(gotTicker, tt.wantTicker) {
				t.Errorf("GetTicker() got \n%+v\n, want\n%+v", gotTicker, tt.wantTicker)
			}
		})
	}
}
