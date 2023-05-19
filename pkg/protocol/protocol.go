package protocol

import (
	"bytes"
	"encoding/binary"

	"github.com/Wishrem/wuso/pkg/errno"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigFastest

type DataType int

const (
	ACKType DataType = iota
	ErrsType
	MsgType
)

type Parser struct {
	r             Receiver
	withoutClient bool
}

func NewParser(r Receiver, withoutClient bool) *Parser {
	return &Parser{
		r,
		withoutClient,
	}
}

func (pa *Parser) Parse(msgType int, p []byte) error {
	switch msgType {
	case websocket.BinaryMessage:
		d := new(data)
		if err := d.Unmarshal(p); err != nil {
			return err
		}
		switch d.Type {
		case ACKType:
			id, err := parseACKId(d.Data)
			if err != nil {
				return err
			}
			pa.r.ReceiveACK(id)
		case ErrsType:
			es := new(errs)
			if err := es.Unmarshal(d.Data); err != nil {
				return err
			}
			pa.r.ReceiveError(es.Id, es.Code)
		case MsgType:
			m := new(msg)
			if err := m.Unmarshal(d.Data); err != nil {
				return err
			}
			pa.r.ReceiveText(m.Id, m.From, m.To, m.Text, !pa.withoutClient)
		}
	case websocket.TextMessage:
		return errno.UnknownMessageType
	default:
		return errno.UnknownMessageType
	}

	return nil
}

func buildBytes(dataType DataType, b []byte) ([]byte, error) {
	d := &data{
		Type: dataType,
		Data: b,
	}
	return d.Marshal()
}

func BuildMsgBytes(id, from, to int64, text string) ([]byte, error) {
	m := &msg{
		Id:   id,
		From: from,
		To:   to,
		Text: text,
	}
	b, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	return buildBytes(MsgType, b)
}

func BuildErrsBytes(id int64, code ErrCode) ([]byte, error) {
	e := &errs{
		Id:   id,
		Code: code,
	}
	b, err := e.Marshal()
	if err != nil {
		return nil, err
	}
	return buildBytes(MsgType, b)
}

func BuildACKBytes(id int64) ([]byte, error) {
	b := make([]byte, 0)
	binary.AppendVarint(b, id)
	return buildBytes(MsgType, b)
}

func parseACKId(data []byte) (int64, error) {
	reader := bytes.NewReader(data)
	return binary.ReadVarint(reader)
}
