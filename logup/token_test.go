package logup

import (
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	now := time.Now()
	token := GenToken(&now)
	t.Log(token)

}
