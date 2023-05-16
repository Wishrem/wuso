package protocol

import (
	"encoding/binary"
	"errors"
)

type CMDCode uint

const (
	CMDDeleteSingleMsg CMDCode = iota
	CMDDeleteSpecificTx
	CMDDeleteAllMsg
	CMDPullMsg
)

func cmdCodeToByte(code CMDCode) byte {
	header := CMDCode(CMDType<<4) | code
	return byte(header)
}

var (
	headerCMDDeleteSingleMsg = cmdCodeToByte(CMDDeleteSingleMsg)
	headerCMDDeleteAllMsg    = cmdCodeToByte(CMDDeleteAllMsg)
	headerCMDPullMsg         = cmdCodeToByte(CMDPullMsg) // TODO
)

func EncodeCMDDeleteSingleMsg(to int64, msgId int64) ([]byte, error) {
	return encodeCMD(headerCMDDeleteSingleMsg, to, msgId)
}
func DecodeCMDDeleteSingleMsg(msg []byte) (int64, int64, error) {
	if len(msg) != 16 {
		return 0, 0, errors.New("length of msg requires 16 bytes")
	}

	x, err := decodeCMD(msg)
	if err != nil {
		return 0, 0, err
	}
	return x[0], x[1], nil
}

func EncodeCMDDeleteAllMsg(to int64) ([]byte, error) {
	return encodeCMD(headerCMDDeleteAllMsg, to)
}
func DecodeCMDDeleteAllMsg(msg []byte) (int64, error) {
	if len(msg) != 8 {
		return 0, errors.New("length of msg requires 24 bytes")
	}
	x, err := decodeCMD(msg)
	if err != nil {
		return 0, err
	}
	return x[0], nil
}

func encodeCMD(header byte, id ...int64) ([]byte, error) {
	buf := make([]byte, len(id)<<3+1)
	buf[0] = header
	var n int
	for i, x := range id {
		n = binary.PutVarint(buf[1+8*i:9+8*i], x)
		if n != 8 {
			return nil, errors.New("need to write 8 bytes")
		}
	}
	return buf, nil
}
func decodeCMD(msg []byte) ([]int64, error) {
	if len(msg) == 0 || len(msg)%8 != 0 {
		return nil, errors.New("wrong length of msg")
	}

	n := len(msg) / 8
	x := make([]int64, n)
	var y int
	for i := 0; i < n; i++ {
		x[i], y = binary.Varint(msg[i*8 : 8+i*8])
		if y != 8 {
			return nil, errors.New("need to read 8 bytes")
		}
	}
	return x, nil
}
