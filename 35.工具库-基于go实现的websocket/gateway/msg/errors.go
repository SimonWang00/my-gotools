package msg

import "errors"

var (
	ErrMsgNotProto = errors.New("msg not proto")
	ErrProtocol    = errors.New("protocol too much")
	ErrNotRegister = errors.New("protocol not register")
	ErrMsgShort    = errors.New("msg too short")
	ErrExceedMaxConn = errors.New("maximum connections exceeded")
)
