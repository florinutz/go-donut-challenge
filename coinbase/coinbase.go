package coinbase

import (
	"fmt"
	"net/http"

	"github.com/preichenberger/go-coinbasepro/v2"
)

// Client offers access to the responses
type Client struct {
	coinbasepro.Client
}

func (c *Client) GetTicker(product string) (ticker *coinbasepro.Ticker, res *http.Response, err error) {
	res, err = c.Request("GET", fmt.Sprintf("/products/%s/ticker", product), nil, &ticker)
	return
}

func (c *Client) CreateOrder(newOrder *coinbasepro.Order) (created *coinbasepro.Order, res *http.Response, err error) {
	if len(newOrder.Type) == 0 {
		newOrder.Type = "limit"
	}
	res, err = c.Request("POST", fmt.Sprintf("/orders"), newOrder, &created)
	return
}
