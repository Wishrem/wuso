package message

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigFastest

type Message struct {
}

func SendMsg(s string) error {
	// TODO
	return m.SendMsg(nil...)
}

func RecvMsg(bytes []byte) {
	m.RecvMsg(bytes)
}

func RecvCmdCode(code uint8) {

}

func GetMsgs() ([]string, error) {
	return nil, nil
}

func (m *Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}
