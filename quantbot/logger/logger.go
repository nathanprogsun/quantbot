package logger

import "go.uber.org/zap"

var L *zap.Logger

func Init() error {
    l, err := zap.NewProduction()
    if err != nil {
        return err
    }
    L = l
    return nil
}

func Sync() {
    if L != nil {
        _ = L.Sync()
    }
}
