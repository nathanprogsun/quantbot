package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func sign(secret, payload string) string { 
    m := hmac.New(sha256.New, []byte(secret))
    m.Write([]byte(payload))
    return hex.EncodeToString(m.Sum(nil))
}
