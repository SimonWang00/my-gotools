package msg

import (
	"my-gotools/35.工具库-基于go实现的websocket/pb"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	NewMsgProtocol(true)
	os.Exit(m.Run())
}

func TestNewMsgProtocol(t *testing.T) {
	p := GetMsgProtocol()
	p.Register(&pb.Ping{}, 1)
	data, err := p.Marshal(&pb.Ping{Times: 1})
	if err != nil {
		t.Error(err)
	}
	info, err := p.Unmarshal(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(info.(*pb.Ping))
}
