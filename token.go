package push

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"time"
)

func GenToken(t *time.Time) string {
	h := md5.New()
	buf := bytes.NewBufferString(t.String())
	h.Write(buf.Bytes())
	b := h.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(b)
}
