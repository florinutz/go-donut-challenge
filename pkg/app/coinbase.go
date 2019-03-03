package app

import (
	"github.com/florinutz/go-donut-challenge/coinbase"
	"github.com/preichenberger/go-coinbasepro/v2"
)

func (*App) GetTicker(product string) (ticker *coinbasepro.Ticker, err error) {
	c := &coinbase.Client{Client: *coinbasepro.NewClient()}
	ticker, _, err = c.GetTicker(product)
	return
}

func (*App) GetProducts() (products []coinbasepro.Product, err error) {
	c := &coinbase.Client{Client: *coinbasepro.NewClient()}
	return c.GetProducts()
}

func (*App) PlaceOrder(newOrder *coinbasepro.Order) (placedOrder *coinbasepro.Order, err error) {
	c := &coinbase.Client{Client: *coinbasepro.NewClient()}
	placedOrder, _, err = c.CreateOrder(newOrder)
	return
}
