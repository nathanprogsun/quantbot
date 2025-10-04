package core

import "context"

type Strategy interface {
    Init(ctx context.Context, ex ExchangeConnector) error
    OnTicker(t Ticker)
    OnOrderUpdate(o Order)
    Run(ctx context.Context)
}
