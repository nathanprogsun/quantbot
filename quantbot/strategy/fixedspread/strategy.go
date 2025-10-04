package fixedspread

import (
	"context"

	"github.com/nathanprogsun/quantbot/quantbot/core"
)

type Config struct {
    Symbol           string
    SpreadBps        float64
    OrderSize        string // use decimal later
    RefreshInterval  int64  // ms
}

type Strategy struct {
    cfg Config
    ex  core.ExchangeConnector
}

func New(cfg Config) *Strategy {
    return &Strategy{cfg: cfg}
}

func (s *Strategy) Init(ctx context.Context, ex core.ExchangeConnector) error {
    s.ex = ex
    return nil
}

func (s *Strategy) OnTicker(t core.Ticker) {
    // TODO: compute quotes and place/replace orders
}

func (s *Strategy) OnOrderUpdate(o core.Order) {
    // TODO: react on partial/filled/canceled
}

func (s *Strategy) Run(ctx context.Context) {
    // TODO: main loop, timers, risk checks
}
