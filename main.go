package main

import (
	"fmt"
	"github.com/everfore/exc"
	"github.com/toukii/oschat/peers"
	"net"
	"time"
)

var Conn_Keeper *peers.ConnKeeper

func init() {
	Conn_Keeper = peers.NewConnKeeper()
}

func screen() {
	cmd := exc.NewCMD("xwd -root -out screen.xwd").Debug()
	cmd.ExecuteAfter(5)
	cmd.Reset("convert screen.xwd screen.png").Execute()
}

func main() {
	screen()
	return
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 8080})
	if exc.Checkerr(err) {
		return
	}
	go func() {
		for {

			time.Sleep(3e9)
			for k, v := range Conn_Keeper.Conns {
				fmt.Printf("send msg to %s\n", k)
				_, err = v.Write([]byte("hello,push msg."))
				fmt.Println(err)
				if exc.Checkerr(err) {
					Conn_Keeper.Delete(v.RemoteAddr().String())
				}
			}
		}
	}()
	for {
		conn, err := listener.AcceptTCP()
		if exc.Checkerr(err) {
			continue
		}
		Conn_Keeper.Set(conn.RemoteAddr().String(), conn)
	}
}
