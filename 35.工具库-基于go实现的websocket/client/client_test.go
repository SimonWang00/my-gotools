package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"my-gotools/websocket/gateway/msg"
	"my-gotools/websocket/pb"
	"net/url"
	"os"
	"testing"
	"time"
)

var conn *websocket.Conn

func TestMain(m *testing.M) {
	msg.NewMsgProtocol(true)
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8888", Path: "/connect"}
	var dialer websocket.Dialer
	var err error
	conn, _, err = dialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	if err = msg.GetMsgProtocol().Register(&pb.Ping{}, 1); err != nil {
		panic(err)
	}
	if err = msg.GetMsgProtocol().Register(&pb.Pong{}, 2); err != nil {
		panic(err)
	}
	go func() {
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				_ = conn.Close()
				panic(err)
			}
			info, err := msg.GetMsgProtocol().Unmarshal(data)
			if err != nil {
				_ = conn.Close()
				panic(err)
			}
			fmt.Println(info)
		}
	}()
	os.Exit(m.Run())
}

func TestPing(t *testing.T) {
	for {
		time.Sleep(time.Second * 1)
		data, err := msg.GetMsgProtocol().Marshal(&pb.Ping{
			Times: time.Now().Unix(),
		})
		if err != nil {
			_ = conn.Close()
			panic(err)
		}
		_ = conn.WriteMessage(websocket.BinaryMessage, data)
	}
}
