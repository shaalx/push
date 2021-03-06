package main

import (
	"expvar"
	"fmt"
	"github.com/toukii/push/logup"
	"net/http"

	"github.com/toukii/goutils"
	"path/filepath"
	"time"
)

var (
	count  *expvar.Int
	tokenM map[string]interface{}
)

func init() {
	count = expvar.NewInt("count")
	tokenM = make(map[string]interface{})
}

func main() {
	http.HandleFunc("/", logupHandler)
	http.HandleFunc("/token/", checktokenHandler)
	http.ListenAndServe(":80", nil)
}

func logupHandler(rw http.ResponseWriter, req *http.Request) {
	count.Add(1)
	now := time.Now()
	token := logup.GenToken(&now)
	fmt.Println(token)
	tokenM[token] = req.RemoteAddr + "[" + count.String() + "]"
	rw.Write(goutils.ToByte(fmt.Sprintf("http://localhost:80/token/%s", token)))
}

func checktokenHandler(rw http.ResponseWriter, req *http.Request) {
	tokenURI := req.RequestURI
	token := filepath.Base(tokenURI)
	fmt.Println(token)
	remote, ok := tokenM[token]
	if ok {
		rw.Write(goutils.ToByte(fmt.Sprintf("current user IP:%v", remote)))
	} else {
		rw.Write(goutils.ToByte("wrong token."))
	}
}
