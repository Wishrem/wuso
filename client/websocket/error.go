package websocket

type ErrCode uint

const (
	ErrRecvWrong = iota
)

func toBytes(code ErrCode) []byte {
	buf := make([]byte, 1)
	header := (ACKType << 4) & code
	buf[0] = byte(header)
	return buf
}

var (
	ErrRecvWrongBytes = toBytes(ErrRecvWrong)
)
