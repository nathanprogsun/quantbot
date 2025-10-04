package binance

import (
	"context"
	"errors"

	"github.com/nathanprogsun/quantbot/quantbot/core"
	"github.com/shopspring/decimal"
)

type Client struct{}

func NewClient(apiKey, apiSecret string) *Client {
    return &Client{}
}

func (c *Client) Name() string { return "binance-spot" }

func (c *Client) PlaceOrder(ctx context.Context, o core.Order) (core.Order, error) {
    return core.Order{}, errors.New("not implemented")
}

func (c *Client) CancelOrder(ctx context.Context, symbol, orderID string) error {
    return errors.New("not implemented")
}

func (c *Client) GetOrder(ctx context.Context, symbol, orderID string) (core.Order, error) {
    return core.Order{}, errors.New("not implemented")
}

func (c *Client) GetBalance(ctx context.Context) (map[string]decimal.Decimal, error) {
    return nil, errors.New("not implemented")
}

func (c *Client) SubscribeTicker(ctx context.Context, symbol string) (<-chan core.Ticker, <-chan error) {
    ticks := make(chan core.Ticker)
    errs := make(chan error, 1)
    // TODO: implement WS @bookTicker
    return ticks, errs
}

func (c *Client) SubscribeUserStream(ctx context.Context) (<-chan core.Order, <-chan map[string]decimal.Decimal, <-chan error) {
    orders := make(chan core.Order)
    bal := make(chan map[string]decimal.Decimal)
    errs := make(chan error, 1)
    // TODO: listenKey + ws user stream (executionReport / outboundAccountPosition)
    return orders, bal, errs
}

func (c *Client) Close() error { return nil }
