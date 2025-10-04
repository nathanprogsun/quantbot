package core

import (
	"context"

	"github.com/shopspring/decimal"
)

type Order struct {
    ID       string
    Symbol   string
    Side     string // BUY/SELL
    Type     string // LIMIT/MARKET
    Price    decimal.Decimal
    Quantity decimal.Decimal
    Status   string
    Ts       int64
}

type Ticker struct {
    Symbol string
    Bid    decimal.Decimal
    Ask    decimal.Decimal
    Ts     int64
}

type ExchangeConnector interface {
    Name() string
    PlaceOrder(ctx context.Context, o Order) (Order, error)
    CancelOrder(ctx context.Context, symbol, orderID string) error
    GetOrder(ctx context.Context, symbol, orderID string) (Order, error)
    GetBalance(ctx context.Context) (map[string]decimal.Decimal, error)

    SubscribeTicker(ctx context.Context, symbol string) (<-chan Ticker, <-chan error)
    SubscribeUserStream(ctx context.Context) (<-chan Order, <-chan map[string]decimal.Decimal, <-chan error)

    Close() error
}
