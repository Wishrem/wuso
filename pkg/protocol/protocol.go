package protocol

import (
	"errors"
)

type DataType uint

// 4bit Type + 4bit CMD code + msg
const (
	MsgType DataType = 0x1 << iota
	CMDType
	ACKType
	Reserve
)

const mask = 0b1111

func Decode(data []byte) (DataType, CMDCode, []byte, error) {
	if len(data) < 2 {
		return 0, 0, nil, errors.New("length of data is too short")
	}

	dt := DataType(data[0] >> 4 & mask)
	code := CMDCode(data[0] & mask)

	return dt, code, data[1:], nil
}
